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
	// The display name prefix. A left-match query is used.
	//
	// example:
	//
	// name
	DisplayNameStartsWith *string `json:"DisplayNameStartsWith,omitempty" xml:"DisplayNameStartsWith,omitempty"`
	// The email address of the account.
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
	// The number of entries per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NTxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizational unit ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number of the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The phone region code. Example: The region code for the Chinese mainland is 86, without the 00 or + prefix.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The account status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The external ID, which is used to associate external data with IDaaS accounts.
	//
	// Note: The external ID must be unique within the same source type and source ID.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The list of account IDs.
	//
	// example:
	//
	// 20
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// The source ID of the account.
	//
	// For self-built accounts, the default value is the instance ID. For other types, the value corresponds to the enterprise ID of the respective source. For example, for a DingTalk source, the value corresponds to the corpId of the DingTalk enterprise.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// - build_in: self-built.
	//
	// - ding_talk: imported from DingTalk.
	//
	// - ad: imported from AD.
	//
	// - ldap: imported from LDAP.
	//
	// - we_com: imported from WeCom.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The username prefix. A left-match query is used.
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
