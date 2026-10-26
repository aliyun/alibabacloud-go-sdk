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
	SetIsDefault(v string) *ModifyInstanceADAuthServerRequest
	GetIsDefault() *string
	SetIsSSL(v string) *ModifyInstanceADAuthServerRequest
	GetIsSSL() *string
	SetIsSyncEmailAttr(v string) *ModifyInstanceADAuthServerRequest
	GetIsSyncEmailAttr() *string
	SetIsSyncMobileAttr(v string) *ModifyInstanceADAuthServerRequest
	GetIsSyncMobileAttr() *string
	SetIsSyncNameAttr(v string) *ModifyInstanceADAuthServerRequest
	GetIsSyncNameAttr() *string
	SetIsSyncOuAsUserGroup(v string) *ModifyInstanceADAuthServerRequest
	GetIsSyncOuAsUserGroup() *string
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
	SetServerName(v string) *ModifyInstanceADAuthServerRequest
	GetServerName() *string
	SetStandbyServer(v string) *ModifyInstanceADAuthServerRequest
	GetStandbyServer() *string
	SetSyncInterval(v string) *ModifyInstanceADAuthServerRequest
	GetSyncInterval() *string
	SetUserSourceId(v string) *ModifyInstanceADAuthServerRequest
	GetUserSourceId() *string
}

type ModifyInstanceADAuthServerRequest struct {
	// The server account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn=Manager,dc=test,dc=com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The base distinguished name (Base DN).
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
	// The server attribute name that maps to the user email address.
	//
	// example:
	//
	// emailAttr
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The filter condition for users.
	//
	// example:
	//
	// (objectClass=top)
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the AD authentication server is the default server. A value of true indicates that the server is the default AD authentication server. A value of false indicates that the server is not the default AD authentication server.
	//
	// example:
	//
	// true
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether SSL is supported. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	//
	// example:
	//
	// true
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// Specifies whether to synchronize the email address. A value of true indicates that the email address is synchronized. A value of false indicates that the email address is not synchronized.
	//
	// example:
	//
	// true
	IsSyncEmailAttr *string `json:"IsSyncEmailAttr,omitempty" xml:"IsSyncEmailAttr,omitempty"`
	// Specifies whether to synchronize the mobile phone number. A value of true indicates that the mobile phone number is synchronized. A value of false indicates that the mobile phone number is not synchronized.
	//
	// example:
	//
	// true
	IsSyncMobileAttr *string `json:"IsSyncMobileAttr,omitempty" xml:"IsSyncMobileAttr,omitempty"`
	// Specifies whether to synchronize the display name. A value of true indicates that the display name is synchronized. A value of false indicates that the display name is not synchronized.
	//
	// example:
	//
	// true
	IsSyncNameAttr *string `json:"IsSyncNameAttr,omitempty" xml:"IsSyncNameAttr,omitempty"`
	// Specifies whether to synchronize the organization to which the user belongs as a user group. A value of true indicates that the organization is synchronized as a user group. A value of false indicates that the organization is not synchronized as a user group.
	//
	// example:
	//
	// true
	IsSyncOuAsUserGroup *string `json:"IsSyncOuAsUserGroup,omitempty" xml:"IsSyncOuAsUserGroup,omitempty"`
	// The server attribute name that maps to the user mobile phone number.
	//
	// example:
	//
	// mobileAttr
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The server attribute name that maps to the username.
	//
	// example:
	//
	// nameAttr
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// The password of the server account.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port used to access the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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
	// The name of the AD authentication server.
	//
	// example:
	//
	// AD Server
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// The address of the standby server.
	//
	// example:
	//
	// 192.168.XX.XX
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
	// The interval at which user snapshots are automatically synchronized, in hours. Valid values: 0, and 4 to 168. A value of 0 indicates that automatic synchronization is disabled.
	//
	// example:
	//
	// 12
	SyncInterval *string `json:"SyncInterval,omitempty" xml:"SyncInterval,omitempty"`
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
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

func (s *ModifyInstanceADAuthServerRequest) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSSL() *string {
	return s.IsSSL
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSyncEmailAttr() *string {
	return s.IsSyncEmailAttr
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSyncMobileAttr() *string {
	return s.IsSyncMobileAttr
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSyncNameAttr() *string {
	return s.IsSyncNameAttr
}

func (s *ModifyInstanceADAuthServerRequest) GetIsSyncOuAsUserGroup() *string {
	return s.IsSyncOuAsUserGroup
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

func (s *ModifyInstanceADAuthServerRequest) GetServerName() *string {
	return s.ServerName
}

func (s *ModifyInstanceADAuthServerRequest) GetStandbyServer() *string {
	return s.StandbyServer
}

func (s *ModifyInstanceADAuthServerRequest) GetSyncInterval() *string {
	return s.SyncInterval
}

func (s *ModifyInstanceADAuthServerRequest) GetUserSourceId() *string {
	return s.UserSourceId
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

func (s *ModifyInstanceADAuthServerRequest) SetIsDefault(v string) *ModifyInstanceADAuthServerRequest {
	s.IsDefault = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSSL(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSyncEmailAttr(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSyncEmailAttr = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSyncMobileAttr(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSyncMobileAttr = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSyncNameAttr(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSyncNameAttr = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSyncOuAsUserGroup(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSyncOuAsUserGroup = &v
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

func (s *ModifyInstanceADAuthServerRequest) SetServerName(v string) *ModifyInstanceADAuthServerRequest {
	s.ServerName = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetStandbyServer(v string) *ModifyInstanceADAuthServerRequest {
	s.StandbyServer = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetSyncInterval(v string) *ModifyInstanceADAuthServerRequest {
	s.SyncInterval = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetUserSourceId(v string) *ModifyInstanceADAuthServerRequest {
	s.UserSourceId = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) Validate() error {
	return dara.Validate(s)
}
