// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStart(v bool) *ModifyParentPlatformRequest
	GetAutoStart() *bool
	SetClientAuth(v bool) *ModifyParentPlatformRequest
	GetClientAuth() *bool
	SetClientPassword(v string) *ModifyParentPlatformRequest
	GetClientPassword() *string
	SetClientUsername(v string) *ModifyParentPlatformRequest
	GetClientUsername() *string
	SetDescription(v string) *ModifyParentPlatformRequest
	GetDescription() *string
	SetGbId(v string) *ModifyParentPlatformRequest
	GetGbId() *string
	SetId(v string) *ModifyParentPlatformRequest
	GetId() *string
	SetIp(v string) *ModifyParentPlatformRequest
	GetIp() *string
	SetName(v string) *ModifyParentPlatformRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyParentPlatformRequest
	GetOwnerId() *int64
	SetPort(v int64) *ModifyParentPlatformRequest
	GetPort() *int64
}

type ModifyParentPlatformRequest struct {
	// Specifies whether to automatically enable the platform. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// Specifies whether to enable local authentication. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	ClientAuth *bool `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	// The local password.
	//
	// example:
	//
	// admin123
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	// The local username.
	//
	// example:
	//
	// user01
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	// The description of the parent platform.
	//
	// example:
	//
	// 国标级联修改测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The GB ID of the parent platform.
	//
	// example:
	//
	// 31000*****2170123451
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// The ID of the parent platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The SIP service IP of the parent platform.
	//
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the parent platform.
	//
	// example:
	//
	// 国标级联修改测试
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The SIP service port of the parent platform.
	//
	// example:
	//
	// 5060
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyParentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *ModifyParentPlatformRequest) GetClientAuth() *bool {
	return s.ClientAuth
}

func (s *ModifyParentPlatformRequest) GetClientPassword() *string {
	return s.ClientPassword
}

func (s *ModifyParentPlatformRequest) GetClientUsername() *string {
	return s.ClientUsername
}

func (s *ModifyParentPlatformRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyParentPlatformRequest) GetGbId() *string {
	return s.GbId
}

func (s *ModifyParentPlatformRequest) GetId() *string {
	return s.Id
}

func (s *ModifyParentPlatformRequest) GetIp() *string {
	return s.Ip
}

func (s *ModifyParentPlatformRequest) GetName() *string {
	return s.Name
}

func (s *ModifyParentPlatformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParentPlatformRequest) GetPort() *int64 {
	return s.Port
}

func (s *ModifyParentPlatformRequest) SetAutoStart(v bool) *ModifyParentPlatformRequest {
	s.AutoStart = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientAuth(v bool) *ModifyParentPlatformRequest {
	s.ClientAuth = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientPassword(v string) *ModifyParentPlatformRequest {
	s.ClientPassword = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetClientUsername(v string) *ModifyParentPlatformRequest {
	s.ClientUsername = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetDescription(v string) *ModifyParentPlatformRequest {
	s.Description = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetGbId(v string) *ModifyParentPlatformRequest {
	s.GbId = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetId(v string) *ModifyParentPlatformRequest {
	s.Id = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetIp(v string) *ModifyParentPlatformRequest {
	s.Ip = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetName(v string) *ModifyParentPlatformRequest {
	s.Name = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetOwnerId(v int64) *ModifyParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParentPlatformRequest) SetPort(v int64) *ModifyParentPlatformRequest {
	s.Port = &v
	return s
}

func (s *ModifyParentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
