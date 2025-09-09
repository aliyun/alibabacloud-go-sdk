// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceLDAPAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *VerifyInstanceLDAPAuthServerRequest
	GetAccount() *string
	SetBaseDN(v string) *VerifyInstanceLDAPAuthServerRequest
	GetBaseDN() *string
	SetFilter(v string) *VerifyInstanceLDAPAuthServerRequest
	GetFilter() *string
	SetInstanceId(v string) *VerifyInstanceLDAPAuthServerRequest
	GetInstanceId() *string
	SetIsSSL(v string) *VerifyInstanceLDAPAuthServerRequest
	GetIsSSL() *string
	SetPassword(v string) *VerifyInstanceLDAPAuthServerRequest
	GetPassword() *string
	SetPort(v string) *VerifyInstanceLDAPAuthServerRequest
	GetPort() *string
	SetRegionId(v string) *VerifyInstanceLDAPAuthServerRequest
	GetRegionId() *string
	SetServer(v string) *VerifyInstanceLDAPAuthServerRequest
	GetServer() *string
	SetStandbyServer(v string) *VerifyInstanceLDAPAuthServerRequest
	GetStandbyServer() *string
}

type VerifyInstanceLDAPAuthServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dc=test,dc=com
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// example:
	//
	// (objectClass=top)
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s VerifyInstanceLDAPAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetAccount() *string {
	return s.Account
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetBaseDN() *string {
	return s.BaseDN
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetFilter() *string {
	return s.Filter
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetIsSSL() *string {
	return s.IsSSL
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetPassword() *string {
	return s.Password
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetPort() *string {
	return s.Port
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetServer() *string {
	return s.Server
}

func (s *VerifyInstanceLDAPAuthServerRequest) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetAccount(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Account = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetBaseDN(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetFilter(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetInstanceId(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetIsSSL(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetPassword(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Password = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetPort(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Port = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetRegionId(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetServer(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Server = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetStandbyServer(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.StandbyServer = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
