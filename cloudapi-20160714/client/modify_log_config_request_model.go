// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogType(v string) *ModifyLogConfigRequest
	GetLogType() *string
	SetSecurityToken(v string) *ModifyLogConfigRequest
	GetSecurityToken() *string
	SetSlsLogStore(v string) *ModifyLogConfigRequest
	GetSlsLogStore() *string
	SetSlsProject(v string) *ModifyLogConfigRequest
	GetSlsProject() *string
}

type ModifyLogConfigRequest struct {
	// The log type. Valid values:
	//
	// 	- **log**: other logs
	//
	// 	- **survey**: inspection logs
	//
	// Enumeration value:
	//
	// 	- PROVIDER
	//
	// example:
	//
	// PROVIDER
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// logs-gateway
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	// The name of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// ford-api-gateway-log
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s ModifyLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigRequest) GetLogType() *string {
	return s.LogType
}

func (s *ModifyLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyLogConfigRequest) GetSlsLogStore() *string {
	return s.SlsLogStore
}

func (s *ModifyLogConfigRequest) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ModifyLogConfigRequest) SetLogType(v string) *ModifyLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSecurityToken(v string) *ModifyLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSlsLogStore(v string) *ModifyLogConfigRequest {
	s.SlsLogStore = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSlsProject(v string) *ModifyLogConfigRequest {
	s.SlsProject = &v
	return s
}

func (s *ModifyLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
