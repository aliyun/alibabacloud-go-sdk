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
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersRequest
	GetNextToken() *string
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
	// The prefix of the display name. The query is performed based on the prefix.
	//
	// example:
	//
	// name
	DisplayNameStartsWith *string `json:"DisplayNameStartsWith,omitempty" xml:"DisplayNameStartsWith,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// NTxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The default value is 20. The maximum value is 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The country calling code. For example, the country calling code of China is `86`. Do not add `00` or `+` to the country calling code.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The status of the user. Valid values:
	//
	// - `enabled`: The user is enabled.
	//
	// - `disabled`: The user is disabled.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The external ID of the user. The external ID can be used to associate the user with a user in an external system.
	//
	// > The external ID must be unique within the same source type and source ID.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The list of user IDs.
	//
	// example:
	//
	// 20
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// The source ID of the user.
	//
	// If the user is created in EIAM, the value of this parameter is the ID of the EIAM instance. If the user is imported from an external system, the value of this parameter is the enterprise ID of the user in the external system. For example, if the user is imported from DingTalk, the value of this parameter is the `corpId` of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the user. Valid values:
	//
	// - `build_in`: The user is created in EIAM.
	//
	// - `ding_talk`: The user is imported from DingTalk.
	//
	// - `ad`: The user is imported from Active Directory (AD).
	//
	// - `ldap`: The user is imported from a Lightweight Directory Access Protocol (LDAP) directory.
	//
	// - `we_com`: The user is imported from WeCom.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The prefix of the username. The query is performed based on the prefix.
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

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetNextToken() *string {
	return s.NextToken
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

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
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
