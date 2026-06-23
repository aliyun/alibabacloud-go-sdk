// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceADAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *VerifyInstanceADAuthServerRequest
	GetAccount() *string
	SetBaseDN(v string) *VerifyInstanceADAuthServerRequest
	GetBaseDN() *string
	SetDomain(v string) *VerifyInstanceADAuthServerRequest
	GetDomain() *string
	SetFilter(v string) *VerifyInstanceADAuthServerRequest
	GetFilter() *string
	SetInstanceId(v string) *VerifyInstanceADAuthServerRequest
	GetInstanceId() *string
	SetIsSSL(v string) *VerifyInstanceADAuthServerRequest
	GetIsSSL() *string
	SetPassword(v string) *VerifyInstanceADAuthServerRequest
	GetPassword() *string
	SetPort(v string) *VerifyInstanceADAuthServerRequest
	GetPort() *string
	SetRegionId(v string) *VerifyInstanceADAuthServerRequest
	GetRegionId() *string
	SetServer(v string) *VerifyInstanceADAuthServerRequest
	GetServer() *string
	SetStandbyServer(v string) *VerifyInstanceADAuthServerRequest
	GetStandbyServer() *string
}

type VerifyInstanceADAuthServerRequest struct {
	// The account of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base DN.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc=test,dc=com
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The domain on the Active Directory server.
	//
	// This parameter is required.
	//
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The filter condition for users.
	//
	// example:
	//
	// (objectClass=top)
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the Bastionhost instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether SSL is supported. Valid values:
	//
	// - **true**: supported
	//
	// - **false**: not supported
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The password of the server account.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the Bastionhost instance. For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The address of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The address of the standby server.
	//
	// example:
	//
	// 192.168.XX.XX
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s VerifyInstanceADAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerRequest) GetAccount() *string {
	return s.Account
}

func (s *VerifyInstanceADAuthServerRequest) GetBaseDN() *string {
	return s.BaseDN
}

func (s *VerifyInstanceADAuthServerRequest) GetDomain() *string {
	return s.Domain
}

func (s *VerifyInstanceADAuthServerRequest) GetFilter() *string {
	return s.Filter
}

func (s *VerifyInstanceADAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *VerifyInstanceADAuthServerRequest) GetIsSSL() *string {
	return s.IsSSL
}

func (s *VerifyInstanceADAuthServerRequest) GetPassword() *string {
	return s.Password
}

func (s *VerifyInstanceADAuthServerRequest) GetPort() *string {
	return s.Port
}

func (s *VerifyInstanceADAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VerifyInstanceADAuthServerRequest) GetServer() *string {
	return s.Server
}

func (s *VerifyInstanceADAuthServerRequest) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *VerifyInstanceADAuthServerRequest) SetAccount(v string) *VerifyInstanceADAuthServerRequest {
	s.Account = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetBaseDN(v string) *VerifyInstanceADAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetDomain(v string) *VerifyInstanceADAuthServerRequest {
	s.Domain = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetFilter(v string) *VerifyInstanceADAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetInstanceId(v string) *VerifyInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetIsSSL(v string) *VerifyInstanceADAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetPassword(v string) *VerifyInstanceADAuthServerRequest {
	s.Password = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetPort(v string) *VerifyInstanceADAuthServerRequest {
	s.Port = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetRegionId(v string) *VerifyInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetServer(v string) *VerifyInstanceADAuthServerRequest {
	s.Server = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetStandbyServer(v string) *VerifyInstanceADAuthServerRequest {
	s.StandbyServer = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
