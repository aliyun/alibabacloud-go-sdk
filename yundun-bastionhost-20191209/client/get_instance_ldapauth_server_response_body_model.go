// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLDAPAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLDAP(v *GetInstanceLDAPAuthServerResponseBodyLDAP) *GetInstanceLDAPAuthServerResponseBody
	GetLDAP() *GetInstanceLDAPAuthServerResponseBodyLDAP
	SetRequestId(v string) *GetInstanceLDAPAuthServerResponseBody
	GetRequestId() *string
}

type GetInstanceLDAPAuthServerResponseBody struct {
	// The settings of LDAP authentication.
	LDAP *GetInstanceLDAPAuthServerResponseBodyLDAP `json:"LDAP,omitempty" xml:"LDAP,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1C60E741-102D-5E8F-9710-B06D3F0183FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceLDAPAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponseBody) GetLDAP() *GetInstanceLDAPAuthServerResponseBodyLDAP {
	return s.LDAP
}

func (s *GetInstanceLDAPAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLDAPAuthServerResponseBody) SetLDAP(v *GetInstanceLDAPAuthServerResponseBodyLDAP) *GetInstanceLDAPAuthServerResponseBody {
	s.LDAP = v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *GetInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBody) Validate() error {
	if s.LDAP != nil {
		if err := s.LDAP.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceLDAPAuthServerResponseBodyLDAP struct {
	// The account of the LDAP server.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base distinguished name (DN).
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
	// (&(objectClass=top))
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Indicates whether passwords are required. Valid values:
	//
	// 	- **true**: required
	//
	// 	- **false**: not required
	//
	// example:
	//
	// true
	HasPassword *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// Indicates whether SSL is supported. Valid values:
	//
	// 	- **true**: supported
	//
	// 	- **false**: not supported
	//
	// example:
	//
	// true
	IsSSL *bool `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
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
	// The port that is used to access the LDAP server.
	//
	// example:
	//
	// 389
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The address of the LDAP server.
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

func (s GetInstanceLDAPAuthServerResponseBodyLDAP) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponseBodyLDAP) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetAccount() *string {
	return s.Account
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetBaseDN() *string {
	return s.BaseDN
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetEmailMapping() *string {
	return s.EmailMapping
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetFilter() *string {
	return s.Filter
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetHasPassword() *string {
	return s.HasPassword
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetIsSSL() *bool {
	return s.IsSSL
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetLoginNameMapping() *string {
	return s.LoginNameMapping
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetMobileMapping() *string {
	return s.MobileMapping
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetNameMapping() *string {
	return s.NameMapping
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetPort() *int64 {
	return s.Port
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetServer() *string {
	return s.Server
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetAccount(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Account = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetBaseDN(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.BaseDN = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetEmailMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.EmailMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetFilter(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Filter = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetHasPassword(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.HasPassword = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetIsSSL(v bool) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.IsSSL = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetLoginNameMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.LoginNameMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetMobileMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.MobileMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetNameMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.NameMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetPort(v int64) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Port = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetServer(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Server = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetStandbyServer(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.StandbyServer = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) Validate() error {
	return dara.Validate(s)
}
