// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceLDAPAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyInstanceLDAPAuthServerRequest
	GetAccount() *string
	SetBaseDN(v string) *ModifyInstanceLDAPAuthServerRequest
	GetBaseDN() *string
	SetEmailMapping(v string) *ModifyInstanceLDAPAuthServerRequest
	GetEmailMapping() *string
	SetFilter(v string) *ModifyInstanceLDAPAuthServerRequest
	GetFilter() *string
	SetInstanceId(v string) *ModifyInstanceLDAPAuthServerRequest
	GetInstanceId() *string
	SetIsSSL(v string) *ModifyInstanceLDAPAuthServerRequest
	GetIsSSL() *string
	SetLoginNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest
	GetLoginNameMapping() *string
	SetMobileMapping(v string) *ModifyInstanceLDAPAuthServerRequest
	GetMobileMapping() *string
	SetNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest
	GetNameMapping() *string
	SetPassword(v string) *ModifyInstanceLDAPAuthServerRequest
	GetPassword() *string
	SetPort(v string) *ModifyInstanceLDAPAuthServerRequest
	GetPort() *string
	SetRegionId(v string) *ModifyInstanceLDAPAuthServerRequest
	GetRegionId() *string
	SetServer(v string) *ModifyInstanceLDAPAuthServerRequest
	GetServer() *string
	SetStandbyServer(v string) *ModifyInstanceLDAPAuthServerRequest
	GetStandbyServer() *string
}

type ModifyInstanceLDAPAuthServerRequest struct {
	// The username of the account that is used for the LDAP server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base distinguished name (DN).
	//
	// This parameter is required.
	//
	// example:
	//
	// dc=test,dc=com
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The field that is used to indicate the email address of a user on the LDAP server.
	//
	// example:
	//
	// emailAttr
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The condition that is used to filter users.
	//
	// example:
	//
	// (objectClass=top)
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to support SSL. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The field that is used to indicate the logon name of a user on the LDAP server.
	//
	// example:
	//
	// userNameAttr
	LoginNameMapping *string `json:"LoginNameMapping,omitempty" xml:"LoginNameMapping,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the LDAP server.
	//
	// example:
	//
	// mobileAttr
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The field that is used to indicate the name of a user on the LDAP server.
	//
	// example:
	//
	// nameAttr
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// The password of the account that is used for the LDAP server. You must configure a password when you configure LDAP authentication. If you leave this parameter empty when you modify the settings of LDAP authentication, the current password is used.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to access the LDAP server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The address of the LDAP server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The address of the secondary LDAP server.
	//
	// example:
	//
	// 192.168.XX.XX
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s ModifyInstanceLDAPAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetBaseDN() *string {
	return s.BaseDN
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetEmailMapping() *string {
	return s.EmailMapping
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetFilter() *string {
	return s.Filter
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetIsSSL() *string {
	return s.IsSSL
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetLoginNameMapping() *string {
	return s.LoginNameMapping
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetMobileMapping() *string {
	return s.MobileMapping
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetNameMapping() *string {
	return s.NameMapping
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetServer() *string {
	return s.Server
}

func (s *ModifyInstanceLDAPAuthServerRequest) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetAccount(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Account = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetBaseDN(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetEmailMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.EmailMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetFilter(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetInstanceId(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetIsSSL(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetLoginNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.LoginNameMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetMobileMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.MobileMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.NameMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetPassword(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetPort(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Port = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetRegionId(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetServer(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Server = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetStandbyServer(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.StandbyServer = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
