// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListProjectResponseBodyAccessDeniedDetail) *ListProjectResponseBody
	GetAccessDeniedDetail() *ListProjectResponseBodyAccessDeniedDetail
	SetCode(v int64) *ListProjectResponseBody
	GetCode() *int64
	SetCount(v int64) *ListProjectResponseBody
	GetCount() *int64
	SetCurrentPage(v int64) *ListProjectResponseBody
	GetCurrentPage() *int64
	SetHasNext(v bool) *ListProjectResponseBody
	GetHasNext() *bool
	SetList(v []*ListProjectResponseBodyList) *ListProjectResponseBody
	GetList() []*ListProjectResponseBodyList
	SetMessage(v string) *ListProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectResponseBody
	GetSuccess() *bool
	SetTotalPage(v int64) *ListProjectResponseBody
	GetTotalPage() *int64
}

type ListProjectResponseBody struct {
	AccessDeniedDetail *ListProjectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// count
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// true
	HasNext *bool                          `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List    []*ListProjectResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) GetAccessDeniedDetail() *ListProjectResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListProjectResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListProjectResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProjectResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListProjectResponseBody) GetList() []*ListProjectResponseBodyList {
	return s.List
}

func (s *ListProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListProjectResponseBody) SetAccessDeniedDetail(v *ListProjectResponseBodyAccessDeniedDetail) *ListProjectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListProjectResponseBody) SetCode(v int64) *ListProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetCurrentPage(v int64) *ListProjectResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListProjectResponseBody) SetHasNext(v bool) *ListProjectResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListProjectResponseBody) SetList(v []*ListProjectResponseBodyList) *ListProjectResponseBody {
	s.List = v
	return s
}

func (s *ListProjectResponseBody) SetMessage(v string) *ListProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) SetSuccess(v bool) *ListProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectResponseBody) SetTotalPage(v int64) *ListProjectResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListProjectResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListProjectResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListProjectResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListProjectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListProjectResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyList struct {
	// example:
	//
	// 5244****
	BusinessId   *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 123123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123214
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Token
	//
	// example:
	//
	// d989623696ab4f87a80b8d5b0b0****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListProjectResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyList) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListProjectResponseBodyList) GetBusinessName() *string {
	return s.BusinessName
}

func (s *ListProjectResponseBodyList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListProjectResponseBodyList) GetId() *string {
	return s.Id
}

func (s *ListProjectResponseBodyList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListProjectResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListProjectResponseBodyList) GetToken() *string {
	return s.Token
}

func (s *ListProjectResponseBodyList) SetBusinessId(v int64) *ListProjectResponseBodyList {
	s.BusinessId = &v
	return s
}

func (s *ListProjectResponseBodyList) SetBusinessName(v string) *ListProjectResponseBodyList {
	s.BusinessName = &v
	return s
}

func (s *ListProjectResponseBodyList) SetCreateTime(v int64) *ListProjectResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyList) SetId(v string) *ListProjectResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListProjectResponseBodyList) SetModifiedTime(v int64) *ListProjectResponseBodyList {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectResponseBodyList) SetName(v string) *ListProjectResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListProjectResponseBodyList) SetToken(v string) *ListProjectResponseBodyList {
	s.Token = &v
	return s
}

func (s *ListProjectResponseBodyList) Validate() error {
	return dara.Validate(s)
}
