// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOrganizationMembersResponseBody
	GetCode() *string
	SetData(v []*ListOrganizationMembersResponseBodyData) *ListOrganizationMembersResponseBody
	GetData() []*ListOrganizationMembersResponseBodyData
	SetMessage(v string) *ListOrganizationMembersResponseBody
	GetMessage() *string
	SetPageNo(v int32) *ListOrganizationMembersResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListOrganizationMembersResponseBody
	GetPageSize() *int32
	SetSuccess(v bool) *ListOrganizationMembersResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListOrganizationMembersResponseBody
	GetTotal() *int32
}

type ListOrganizationMembersResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data []*ListOrganizationMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 18
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOrganizationMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOrganizationMembersResponseBody) GetData() []*ListOrganizationMembersResponseBodyData {
	return s.Data
}

func (s *ListOrganizationMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOrganizationMembersResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListOrganizationMembersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationMembersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOrganizationMembersResponseBody) SetCode(v string) *ListOrganizationMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetData(v []*ListOrganizationMembersResponseBodyData) *ListOrganizationMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetMessage(v string) *ListOrganizationMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetPageNo(v int32) *ListOrganizationMembersResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetPageSize(v int32) *ListOrganizationMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetSuccess(v bool) *ListOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetTotal(v int32) *ListOrganizationMembersResponseBody {
	s.Total = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyData struct {
	// The member business ID.
	//
	// example:
	//
	// 112233
	AccountBizId *string `json:"AccountBizId,omitempty" xml:"AccountBizId,omitempty"`
	// The ID of the member account.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account.
	//
	// example:
	//
	// test_001
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// API Key ID
	//
	// example:
	//
	// key_123456789
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The email address of the member.
	//
	// example:
	//
	// test@email.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The time when the member joined.
	//
	// example:
	//
	// 2026-06-10T11:57:42.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The masked API key.
	//
	// example:
	//
	// prefix.abc****456
	MaskedApiKey *string `json:"MaskedApiKey,omitempty" xml:"MaskedApiKey,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_123456789
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The list of member roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The ID used to allocate the seat resource.
	//
	// example:
	//
	// seat_123456
	SeatId *string `json:"SeatId,omitempty" xml:"SeatId,omitempty"`
	// The seat specification type. Valid values:
	//
	// - standard: Standard seat.
	//
	// - pro: Pro seat.
	//
	// - max: Premium seat.
	//
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The member status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOrganizationMembersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountBizId() *string {
	return s.AccountBizId
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListOrganizationMembersResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListOrganizationMembersResponseBodyData) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ListOrganizationMembersResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ListOrganizationMembersResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListOrganizationMembersResponseBodyData) GetMaskedApiKey() *string {
	return s.MaskedApiKey
}

func (s *ListOrganizationMembersResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *ListOrganizationMembersResponseBodyData) GetRoles() []*string {
	return s.Roles
}

func (s *ListOrganizationMembersResponseBodyData) GetSeatId() *string {
	return s.SeatId
}

func (s *ListOrganizationMembersResponseBodyData) GetSpecType() *string {
	return s.SpecType
}

func (s *ListOrganizationMembersResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountBizId(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountBizId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountId(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetAccountName(v string) *ListOrganizationMembersResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetApiKeyId(v string) *ListOrganizationMembersResponseBodyData {
	s.ApiKeyId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetEmail(v string) *ListOrganizationMembersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetGmtCreate(v string) *ListOrganizationMembersResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetMaskedApiKey(v string) *ListOrganizationMembersResponseBodyData {
	s.MaskedApiKey = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetOrgId(v string) *ListOrganizationMembersResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetRoles(v []*string) *ListOrganizationMembersResponseBodyData {
	s.Roles = v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetSeatId(v string) *ListOrganizationMembersResponseBodyData {
	s.SeatId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetSpecType(v string) *ListOrganizationMembersResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) SetStatus(v string) *ListOrganizationMembersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
