// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RevokeTokenRequest
	GetInstanceId() *string
	SetToken(v string) *RevokeTokenRequest
	GetToken() *string
}

type RevokeTokenRequest struct {
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details*	- page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com/).
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that you want to revoke.
	//
	// This parameter is required.
	//
	// example:
	//
	// LzMT+XLFl5s/YWJ/MlDz4t/Lq5HC1iGU1P28HAMaxYxn8aQbALNtml7QZKl9L9kPe6LqUb95tEVo+zUqOogs9+jZwDUSzsd4X4qaD3n2TrBEuMOqKkk1Xdrvu9VBQQvIYbz7MJWZDYC3DlW7gLEr33Cuj54iIhagtBi3epStJitsssWs7otY9zhKOSZxhr49G3d0bh35mwyP18EMvDas8UlzeSozsSrujNUqZXOGK0PEBSd+rWMGDJlCt6GFmJgm2JFY7PJwf/7OOSmUYIYFs5o/PuPpoTMF+hcVXMs+0yDukIMTOzG9m3t8k36PVrghFmnK6pC3Rt3mibjW****ng==
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeTokenRequest) GetToken() *string {
	return s.Token
}

func (s *RevokeTokenRequest) SetInstanceId(v string) *RevokeTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeTokenRequest) Validate() error {
	return dara.Validate(s)
}
