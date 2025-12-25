// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListSceneResponseBodyAccessDeniedDetail) *ListSceneResponseBody
	GetAccessDeniedDetail() *ListSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *ListSceneResponseBody
	GetCode() *int64
	SetCount(v int64) *ListSceneResponseBody
	GetCount() *int64
	SetCurrentPage(v int64) *ListSceneResponseBody
	GetCurrentPage() *int64
	SetHasNext(v bool) *ListSceneResponseBody
	GetHasNext() *bool
	SetList(v []*ListSceneResponseBodyList) *ListSceneResponseBody
	GetList() []*ListSceneResponseBodyList
	SetMessage(v string) *ListSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSceneResponseBody
	GetSuccess() *bool
	SetTotalPage(v int64) *ListSceneResponseBody
	GetTotalPage() *int64
}

type ListSceneResponseBody struct {
	AccessDeniedDetail *ListSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	HasNext *bool                        `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List    []*ListSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBody) GetAccessDeniedDetail() *ListSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListSceneResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListSceneResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSceneResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListSceneResponseBody) GetList() []*ListSceneResponseBodyList {
	return s.List
}

func (s *ListSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSceneResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListSceneResponseBody) SetAccessDeniedDetail(v *ListSceneResponseBodyAccessDeniedDetail) *ListSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListSceneResponseBody) SetCode(v int64) *ListSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSceneResponseBody) SetCount(v int64) *ListSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSceneResponseBody) SetCurrentPage(v int64) *ListSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSceneResponseBody) SetHasNext(v bool) *ListSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSceneResponseBody) SetList(v []*ListSceneResponseBodyList) *ListSceneResponseBody {
	s.List = v
	return s
}

func (s *ListSceneResponseBody) SetMessage(v string) *ListSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSceneResponseBody) SetRequestId(v string) *ListSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneResponseBody) SetSuccess(v bool) *ListSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSceneResponseBody) SetTotalPage(v int64) *ListSceneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSceneResponseBody) Validate() error {
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

type ListSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListSceneResponseBodyList struct {
	// example:
	//
	// www.example.com/xxxx/xxx.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 1621236933677
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1621236933677
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 厨房
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// d989623696ab4f87a80b8d5b0b00****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// example:
	//
	// false
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// 20
	SourceNum *int64 `json:"SourceNum,omitempty" xml:"SourceNum,omitempty"`
	// example:
	//
	// init
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// example:
	//
	// 20
	SubSceneNum *int64 `json:"SubSceneNum,omitempty" xml:"SubSceneNum,omitempty"`
	// example:
	//
	// MODEL_3D
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSceneResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBodyList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListSceneResponseBodyList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSceneResponseBodyList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListSceneResponseBodyList) GetId() *string {
	return s.Id
}

func (s *ListSceneResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListSceneResponseBodyList) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *ListSceneResponseBodyList) GetPublished() *bool {
	return s.Published
}

func (s *ListSceneResponseBodyList) GetSourceNum() *int64 {
	return s.SourceNum
}

func (s *ListSceneResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListSceneResponseBodyList) GetStatusName() *string {
	return s.StatusName
}

func (s *ListSceneResponseBodyList) GetSubSceneNum() *int64 {
	return s.SubSceneNum
}

func (s *ListSceneResponseBodyList) GetType() *string {
	return s.Type
}

func (s *ListSceneResponseBodyList) SetCoverUrl(v string) *ListSceneResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListSceneResponseBodyList) SetGmtCreate(v int64) *ListSceneResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListSceneResponseBodyList) SetGmtModified(v int64) *ListSceneResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListSceneResponseBodyList) SetId(v string) *ListSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSceneResponseBodyList) SetName(v string) *ListSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSceneResponseBodyList) SetPreviewToken(v string) *ListSceneResponseBodyList {
	s.PreviewToken = &v
	return s
}

func (s *ListSceneResponseBodyList) SetPublished(v bool) *ListSceneResponseBodyList {
	s.Published = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSourceNum(v int64) *ListSceneResponseBodyList {
	s.SourceNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetStatus(v string) *ListSceneResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSceneResponseBodyList) SetStatusName(v string) *ListSceneResponseBodyList {
	s.StatusName = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSubSceneNum(v int64) *ListSceneResponseBodyList {
	s.SubSceneNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetType(v string) *ListSceneResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListSceneResponseBodyList) Validate() error {
	return dara.Validate(s)
}
