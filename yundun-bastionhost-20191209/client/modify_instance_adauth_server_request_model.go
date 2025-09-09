// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceADAuthServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyInstanceADAuthServerRequest
	GetAccount() *string
	SetBaseDN(v string) *ModifyInstanceADAuthServerRequest
	GetBaseDN() *string
	SetDomain(v string) *ModifyInstanceADAuthServerRequest
	GetDomain() *string
	SetEmailMapping(v string) *ModifyInstanceADAuthServerRequest
	GetEmailMapping() *string
	SetFilter(v string) *ModifyInstanceADAuthServerRequest
	GetFilter() *string
	SetInstanceId(v string) *ModifyInstanceADAuthServerRequest
	GetInstanceId() *string
	SetIsSSL(v string) *ModifyInstanceADAuthServerRequest
	GetIsSSL() *string
	SetMobileMapping(v string) *ModifyInstanceADAuthServerRequest
	GetMobileMapping() *string
	SetNameMapping(v string) *ModifyInstanceADAuthServerRequest
	GetNameMapping() *string
	SetPassword(v string) *ModifyInstanceADAuthServerRequest
	GetPassword() *string
	SetPort(v string) *ModifyInstanceADAuthServerRequest
	GetPort() *string
	SetRegionId(v string) *ModifyInstanceADAuthServerRequest
	GetRegionId() *string
	SetServer(v string) *ModifyInstanceADAuthServerRequest
	GetServer() *string
	SetStandbyServer(v string) *ModifyInstanceADAuthServerRequest
	GetStandbyServer() *string
}

type ModifyInstanceADAuthServerRequest struct {
	// The username of the account that is used for the AD server.
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
	// The domain on the AD server.
	//
	// This parameter is required.
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
	// Specifies whether SSL is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
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
	// The password of the account that is used for the AD server.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to access the server.
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
	// The address of the AD server.
	//
	// This parameter is required.
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

func (s ModifyInstanceADAuthServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyInstanceADAuthServerRequest) GetBaseDN() *string {
	return s.BaseDN
}

func (s *ModifyInstanceADAuthServerRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyInstanceADAuthServerRequest) GetEmailMapping() *string {
	return s.EmailMapping
}

func (s *ModifyInstanceADAuthServerRequest) GetFilter() *string {
	return s.Filter
}

func (s *ModifyInstanceADAuthServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSSL() *string {
	return s.IsSSL
}

func (s *ModifyInstanceADAuthServerRequest) GetMobileMapping() *string {
	return s.MobileMapping
}

func (s *ModifyInstanceADAuthServerRequest) GetNameMapping() *string {
	return s.NameMapping
}

func (s *ModifyInstanceADAuthServerRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceADAuthServerRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyInstanceADAuthServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceADAuthServerRequest) GetServer() *string {
	return s.Server
}

func (s *ModifyInstanceADAuthServerRequest) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *ModifyInstanceADAuthServerRequest) SetAccount(v string) *ModifyInstanceADAuthServerRequest {
	s.Account = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetBaseDN(v string) *ModifyInstanceADAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetDomain(v string) *ModifyInstanceADAuthServerRequest {
	s.Domain = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetEmailMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.EmailMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetFilter(v string) *ModifyInstanceADAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetInstanceId(v string) *ModifyInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSSL(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetMobileMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.MobileMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetNameMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.NameMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetPassword(v string) *ModifyInstanceADAuthServerRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetPort(v string) *ModifyInstanceADAuthServerRequest {
	s.Port = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetRegionId(v string) *ModifyInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetServer(v string) *ModifyInstanceADAuthServerRequest {
	s.Server = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetStandbyServer(v string) *ModifyInstanceADAuthServerRequest {
	s.StandbyServer = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
