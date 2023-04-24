package projector

import (
	"encoding/json"
	"os"
	"path"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config *Config
	data   *Data
}

func CreateProjector(config *Config, data *Data) *Projector {
	return &Projector{
		config: config,
		data:   data,
	}
}

func (p *Projector) GetValue(key string) (string, bool) {

	curr := p.config.Pwd
	prev := ""

	for curr != prev {
		if dir, ok := p.data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				return value, true
			}
		}
		prev = curr
		curr = path.Dir(curr)
	}

	return "", false
}

func (p *Projector) GetValues() map[string]string {

	values := map[string]string{}
	paths := []string{}

	curr := p.config.Pwd
	prev := ""

	for curr != prev {
		if _, ok := p.data.Projector[curr]; ok {
			paths = append(paths, curr)
		}
		prev = curr
		curr = path.Dir(curr)
	}

	for i := len(paths) - 1; i >= 0; i-- {
		for k, v := range p.data.Projector[paths[i]] {
			values[k] = v
		}
	}

	return values

}
func (p *Projector) SetValue(key, value string) {
	pwd := p.config.Pwd

	if _, ok := p.data.Projector[pwd]; !ok {
		p.data.Projector[pwd] = map[string]string{}

	}

	p.data.Projector[pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	delete(p.data.Projector[p.config.Pwd], key)
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		config: config,
		data: &Data{
			Projector: map[string]map[string]string{},
		},
	}
}

func (p *Projector) Save() error {
	dir := path.Dir(p.config.Config)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)

		if err != nil {
			return err
		}
	}

	jsonStr, err := json.Marshal(p.data)

	if err != nil {
		return err
	}

	os.WriteFile(p.config.Config, jsonStr, 0755)

	return nil
}

func NewProjector(config *Config) *Projector {
	if _, err := os.Stat(config.Config); err == nil {
		contents, err := os.ReadFile(config.Config)
		if err != nil {
			return defaultProjector(config)
		}

		var data Data

		err = json.Unmarshal(contents, &data)

		if err != nil {
			return defaultProjector(config)
		}

		return &Projector{
			config: config,
			data:   &data,
		}
	}

	return defaultProjector(config)
}
