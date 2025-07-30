// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayNameStartsWith(v string) *ListUsersRequest
	GetDisplayNameStartsWith() *string
	SetEmail(v string) *ListUsersRequest
	GetEmail() *string
	SetInstanceId(v string) *ListUsersRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *ListUsersRequest
	GetOrganizationalUnitId() *string
	SetPageNumber(v int64) *ListUsersRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUsersRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *ListUsersRequest
	GetPhoneNumber() *string
	SetPhoneRegion(v string) *ListUsersRequest
	GetPhoneRegion() *string
	SetStatus(v string) *ListUsersRequest
	GetStatus() *string
	SetUserExternalId(v string) *ListUsersRequest
	GetUserExternalId() *string
	SetUserIds(v []*string) *ListUsersRequest
	GetUserIds() []*string
	SetUserSourceId(v string) *ListUsersRequest
	GetUserSourceId() *string
	SetUserSourceType(v string) *ListUsersRequest
	GetUserSourceType() *string
	SetUsernameStartsWith(v string) *ListUsersRequest
	GetUsernameStartsWith() *string
}

type ListUsersRequest struct {
	// Displayname
	//
	// example:
	//
	// name_001
	DisplayNameStartsWith *string `json:"DisplayNameStartsWith,omitempty" xml:"DisplayNameStartsWith,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- enabled: The account is enabled.
	//
	// 	- disabled: The account is disabled.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The external ID of the account. The external ID can be used by external data to map the data of the account in IDaaS EIAM.
	//
	// For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// User ID set
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// Username
	//
	// example:
	//
	// name_001
	UsernameStartsWith *string `json:"UsernameStartsWith,omitempty" xml:"UsernameStartsWith,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetDisplayNameStartsWith() *string {
	return s.DisplayNameStartsWith
}

func (s *ListUsersRequest) GetEmail() *string {
	return s.Email
}

func (s *ListUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListUsersRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListUsersRequest) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *ListUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUsersRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *ListUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListUsersRequest) GetUserSourceId() *string {
	return s.UserSourceId
}

func (s *ListUsersRequest) GetUserSourceType() *string {
	return s.UserSourceType
}

func (s *ListUsersRequest) GetUsernameStartsWith() *string {
	return s.UsernameStartsWith
}

func (s *ListUsersRequest) SetDisplayNameStartsWith(v string) *ListUsersRequest {
	s.DisplayNameStartsWith = &v
	return s
}

func (s *ListUsersRequest) SetEmail(v string) *ListUsersRequest {
	s.Email = &v
	return s
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetOrganizationalUnitId(v string) *ListUsersRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int64) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPhoneNumber(v string) *ListUsersRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersRequest) SetPhoneRegion(v string) *ListUsersRequest {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) SetUserExternalId(v string) *ListUsersRequest {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersRequest) SetUserIds(v []*string) *ListUsersRequest {
	s.UserIds = v
	return s
}

func (s *ListUsersRequest) SetUserSourceId(v string) *ListUsersRequest {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersRequest) SetUserSourceType(v string) *ListUsersRequest {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersRequest) SetUsernameStartsWith(v string) *ListUsersRequest {
	s.UsernameStartsWith = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
