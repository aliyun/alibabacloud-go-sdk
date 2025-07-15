// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSlr(v bool) *CreateLogConfigRequest
	GetCreateSlr() *bool
	SetLogType(v string) *CreateLogConfigRequest
	GetLogType() *string
	SetSecurityToken(v string) *CreateLogConfigRequest
	GetSecurityToken() *string
	SetSlsLogStore(v string) *CreateLogConfigRequest
	GetSlsLogStore() *string
	SetSlsProject(v string) *CreateLogConfigRequest
	GetSlsProject() *string
}

type CreateLogConfigRequest struct {
	// Specifies to create a service-linked role.
	//
	// example:
	//
	// true
	CreateSlr *bool `json:"CreateSlr,omitempty" xml:"CreateSlr,omitempty"`
	// The log type.
	//
	// Valid values:
	//
	// 	- PROVIDER
	//
	// example:
	//
	// PROVIDER
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// slslogstore
	//
	// This parameter is required.
	//
	// example:
	//
	// api-gateway
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	// The name of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// rec-lq-sls
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s CreateLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLogConfigRequest) GetCreateSlr() *bool {
	return s.CreateSlr
}

func (s *CreateLogConfigRequest) GetLogType() *string {
	return s.LogType
}

func (s *CreateLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateLogConfigRequest) GetSlsLogStore() *string {
	return s.SlsLogStore
}

func (s *CreateLogConfigRequest) GetSlsProject() *string {
	return s.SlsProject
}

func (s *CreateLogConfigRequest) SetCreateSlr(v bool) *CreateLogConfigRequest {
	s.CreateSlr = &v
	return s
}

func (s *CreateLogConfigRequest) SetLogType(v string) *CreateLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *CreateLogConfigRequest) SetSecurityToken(v string) *CreateLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLogConfigRequest) SetSlsLogStore(v string) *CreateLogConfigRequest {
	s.SlsLogStore = &v
	return s
}

func (s *CreateLogConfigRequest) SetSlsProject(v string) *CreateLogConfigRequest {
	s.SlsProject = &v
	return s
}

func (s *CreateLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
