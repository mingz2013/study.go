package conf

import "fmt"

func GetSDKAddr() (ServerConfig, error) {

	var c ServerConfig

	err := LoadJsonFile("config/sdk/0.json", &c)
	if err != nil {
		return c, err
	}
	fmt.Println(c)
	return c, nil
}
