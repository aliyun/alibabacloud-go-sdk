// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *RestartInstanceRequest
	GetBody() *string
	SetClientToken(v string) *RestartInstanceRequest
	GetClientToken() *string
	SetForce(v bool) *RestartInstanceRequest
	GetForce() *bool
}

type RestartInstanceRequest struct {
	// example:
	//
	// 1. restart nodes example:
	//
	// {
	//
	//   "restartType": "nodeIp",
	//
	//   "nodes": [
	//
	//     "es-cn-xx-data-j-0"
	//
	//   ],
	//
	//   "blueGreenDep": false
	//
	// }
	//
	// 2. restart instance example:
	//
	// {
	//
	//   "restartType": "instance",
	//
	//   "blueGreenDep": false
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// A unique token used to ensure the idempotence of the request. The client generates this value. The value must be unique among different requests and can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to forcefully restart the cluster regardless of the cluster status.
	//
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetBody() *string {
	return s.Body
}

func (s *RestartInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *RestartInstanceRequest) SetBody(v string) *RestartInstanceRequest {
	s.Body = &v
	return s
}

func (s *RestartInstanceRequest) SetClientToken(v string) *RestartInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartInstanceRequest) SetForce(v bool) *RestartInstanceRequest {
	s.Force = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
