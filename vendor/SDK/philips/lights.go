package lights

import (
	"path/filepath"
	"encoding/json"
	"fmt"
	"github.com/Ascension/main/common/util"
)

const (
	baseURL = "http://localhost/api/newdeveloper/lights"
)

// GetAll gets all the lights on the network
func GetAll() ([]ColorHue, error) {
	fmt.Println("base: ", baseURL)
	data, err := util.Get(baseURL)
	if err != nil {
		return nil, err
	}
	lightMap := map[string]ColorHue{}
	err = json.Unmarshal(data, &lightMap)
	if err != nil {
		return nil, err
	}

	lights := []ColorHue{}
	for index, light := range lightMap {
		light.Index = index
		lights = append(lights, light)
	}
	return lights, nil
}

func GetNew() {

}

func SearchForNew() {

}

func GetStateAttrs() {

}

// SetAttrs renames the light.
func SetAttrs() {

}

// SetState sets the state of a light based on it'd id.
func SetState(light ColorHue) error {
	stateURI := baseURL+"/"+filepath.Join(light.Index,"state")
	bytes, err := json.Marshal(light.State)
	if err != nil {
		return err
	}
	res, err := util.Put(stateURI, bytes)
	fmt.Println("Response: ", res)
	return nil
}

func Delete() {

}