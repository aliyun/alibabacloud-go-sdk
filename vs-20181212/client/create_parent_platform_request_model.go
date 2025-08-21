// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStart(v bool) *CreateParentPlatformRequest
	GetAutoStart() *bool
	SetClientAuth(v bool) *CreateParentPlatformRequest
	GetClientAuth() *bool
	SetClientPassword(v string) *CreateParentPlatformRequest
	GetClientPassword() *string
	SetClientUsername(v string) *CreateParentPlatformRequest
	GetClientUsername() *string
	SetDescription(v string) *CreateParentPlatformRequest
	GetDescription() *string
	SetGbId(v string) *CreateParentPlatformRequest
	GetGbId() *string
	SetIp(v string) *CreateParentPlatformRequest
	GetIp() *string
	SetName(v string) *CreateParentPlatformRequest
	GetName() *string
	SetOwnerId(v int64) *CreateParentPlatformRequest
	GetOwnerId() *int64
	SetPort(v int64) *CreateParentPlatformRequest
	GetPort() *int64
	SetProtocol(v string) *CreateParentPlatformRequest
	GetProtocol() *string
}

type CreateParentPlatformRequest struct {
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// example:
	//
	// true
	ClientAuth *bool `json:"ClientAuth,omitempty" xml:"ClientAuth,omitempty"`
	// example:
	//
	// admin123
	ClientPassword *string `json:"ClientPassword,omitempty" xml:"ClientPassword,omitempty"`
	// example:
	//
	// user01
	ClientUsername *string `json:"ClientUsername,omitempty" xml:"ClientUsername,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 31000*****2170123451
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5060
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateParentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParentPlatformRequest) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *CreateParentPlatformRequest) GetClientAuth() *bool {
	return s.ClientAuth
}

func (s *CreateParentPlatformRequest) GetClientPassword() *string {
	return s.ClientPassword
}

func (s *CreateParentPlatformRequest) GetClientUsername() *string {
	return s.ClientUsername
}

func (s *CreateParentPlatformRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParentPlatformRequest) GetGbId() *string {
	return s.GbId
}

func (s *CreateParentPlatformRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateParentPlatformRequest) GetName() *string {
	return s.Name
}

func (s *CreateParentPlatformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateParentPlatformRequest) GetPort() *int64 {
	return s.Port
}

func (s *CreateParentPlatformRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateParentPlatformRequest) SetAutoStart(v bool) *CreateParentPlatformRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientAuth(v bool) *CreateParentPlatformRequest {
	s.ClientAuth = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientPassword(v string) *CreateParentPlatformRequest {
	s.ClientPassword = &v
	return s
}

func (s *CreateParentPlatformRequest) SetClientUsername(v string) *CreateParentPlatformRequest {
	s.ClientUsername = &v
	return s
}

func (s *CreateParentPlatformRequest) SetDescription(v string) *CreateParentPlatformRequest {
	s.Description = &v
	return s
}

func (s *CreateParentPlatformRequest) SetGbId(v string) *CreateParentPlatformRequest {
	s.GbId = &v
	return s
}

func (s *CreateParentPlatformRequest) SetIp(v string) *CreateParentPlatformRequest {
	s.Ip = &v
	return s
}

func (s *CreateParentPlatformRequest) SetName(v string) *CreateParentPlatformRequest {
	s.Name = &v
	return s
}

func (s *CreateParentPlatformRequest) SetOwnerId(v int64) *CreateParentPlatformRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateParentPlatformRequest) SetPort(v int64) *CreateParentPlatformRequest {
	s.Port = &v
	return s
}

func (s *CreateParentPlatformRequest) SetProtocol(v string) *CreateParentPlatformRequest {
	s.Protocol = &v
	return s
}

func (s *CreateParentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
