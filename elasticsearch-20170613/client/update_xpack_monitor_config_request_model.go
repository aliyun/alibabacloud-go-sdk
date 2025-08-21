// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateXpackMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateXpackMonitorConfigRequest
	GetClientToken() *string
	SetEnable(v bool) *UpdateXpackMonitorConfigRequest
	GetEnable() *bool
	SetEndpoints(v []*string) *UpdateXpackMonitorConfigRequest
	GetEndpoints() []*string
	SetPassword(v string) *UpdateXpackMonitorConfigRequest
	GetPassword() *string
	SetUserName(v string) *UpdateXpackMonitorConfigRequest
	GetUserName() *string
}

type UpdateXpackMonitorConfigRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	Enable    *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	Endpoints []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// if can be null:
	// true
	//
	// example:
	//
	// ******
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateXpackMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateXpackMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateXpackMonitorConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateXpackMonitorConfigRequest) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *UpdateXpackMonitorConfigRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateXpackMonitorConfigRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateXpackMonitorConfigRequest) SetClientToken(v string) *UpdateXpackMonitorConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateXpackMonitorConfigRequest) SetEnable(v bool) *UpdateXpackMonitorConfigRequest {
	s.Enable = &v
	return s
}

func (s *UpdateXpackMonitorConfigRequest) SetEndpoints(v []*string) *UpdateXpackMonitorConfigRequest {
	s.Endpoints = v
	return s
}

func (s *UpdateXpackMonitorConfigRequest) SetPassword(v string) *UpdateXpackMonitorConfigRequest {
	s.Password = &v
	return s
}

func (s *UpdateXpackMonitorConfigRequest) SetUserName(v string) *UpdateXpackMonitorConfigRequest {
	s.UserName = &v
	return s
}

func (s *UpdateXpackMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
