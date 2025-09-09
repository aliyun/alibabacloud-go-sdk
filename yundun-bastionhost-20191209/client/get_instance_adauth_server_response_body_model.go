// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceADAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAD(v *GetInstanceADAuthServerResponseBodyAD) *GetInstanceADAuthServerResponseBody
	GetAD() *GetInstanceADAuthServerResponseBodyAD
	SetRequestId(v string) *GetInstanceADAuthServerResponseBody
	GetRequestId() *string
}

type GetInstanceADAuthServerResponseBody struct {
	// The settings of AD authentication.
	AD *GetInstanceADAuthServerResponseBodyAD `json:"AD,omitempty" xml:"AD,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 89398CFB-4EB6-4C7E-BB3C-EF213AC8FA50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceADAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponseBody) GetAD() *GetInstanceADAuthServerResponseBodyAD {
	return s.AD
}

func (s *GetInstanceADAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceADAuthServerResponseBody) SetAD(v *GetInstanceADAuthServerResponseBodyAD) *GetInstanceADAuthServerResponseBody {
	s.AD = v
	return s
}

func (s *GetInstanceADAuthServerResponseBody) SetRequestId(v string) *GetInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceADAuthServerResponseBodyAD struct {
	// The distinguished name (DN) of the AD server account.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base DN of the AD server.
	//
	// example:
	//
	// dc=test,dc=com
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The domain on the AD server.
	//
	// example:
	//
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The field that is used to indicate the email address of a user on the AD server.
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
	// 	- **true**:
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// Indicates whether SSL is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsSSL *bool `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the AD server.
	//
	// example:
	//
	// mobileAttr
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The field that is used to indicate the name of a user on the AD server.
	//
	// example:
	//
	// nameAttr
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// The port that is used to access the AD server.
	//
	// example:
	//
	// 389
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The address of the AD server.
	//
	// example:
	//
	// 192.168.XX.XX
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The address of the secondary AD server.
	//
	// example:
	//
	// 192.168.XX.XX
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s GetInstanceADAuthServerResponseBodyAD) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceADAuthServerResponseBodyAD) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetAccount() *string {
	return s.Account
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetBaseDN() *string {
	return s.BaseDN
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetEmailMapping() *string {
	return s.EmailMapping
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetFilter() *string {
	return s.Filter
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetIsSSL() *bool {
	return s.IsSSL
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetMobileMapping() *string {
	return s.MobileMapping
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetNameMapping() *string {
	return s.NameMapping
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetPort() *int64 {
	return s.Port
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetServer() *string {
	return s.Server
}

func (s *GetInstanceADAuthServerResponseBodyAD) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetAccount(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Account = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetBaseDN(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.BaseDN = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetDomain(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Domain = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetEmailMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.EmailMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetFilter(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Filter = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetHasPassword(v bool) *GetInstanceADAuthServerResponseBodyAD {
	s.HasPassword = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetIsSSL(v bool) *GetInstanceADAuthServerResponseBodyAD {
	s.IsSSL = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetMobileMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.MobileMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetNameMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.NameMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetPort(v int64) *GetInstanceADAuthServerResponseBodyAD {
	s.Port = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetServer(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Server = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetStandbyServer(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.StandbyServer = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) Validate() error {
	return dara.Validate(s)
}
