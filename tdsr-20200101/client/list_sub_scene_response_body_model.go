// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListSubSceneResponseBodyAccessDeniedDetail) *ListSubSceneResponseBody
	GetAccessDeniedDetail() *ListSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *ListSubSceneResponseBody
	GetCode() *int64
	SetCount(v int64) *ListSubSceneResponseBody
	GetCount() *int64
	SetCurrentPage(v int64) *ListSubSceneResponseBody
	GetCurrentPage() *int64
	SetHasNext(v bool) *ListSubSceneResponseBody
	GetHasNext() *bool
	SetList(v []*ListSubSceneResponseBodyList) *ListSubSceneResponseBody
	GetList() []*ListSubSceneResponseBodyList
	SetMessage(v string) *ListSubSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSubSceneResponseBody
	GetSuccess() *bool
	SetTotalPage(v int64) *ListSubSceneResponseBody
	GetTotalPage() *int64
}

type ListSubSceneResponseBody struct {
	AccessDeniedDetail *ListSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	HasNext *bool                           `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List    []*ListSubSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
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

func (s ListSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBody) GetAccessDeniedDetail() *ListSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListSubSceneResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListSubSceneResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSubSceneResponseBody) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListSubSceneResponseBody) GetList() []*ListSubSceneResponseBodyList {
	return s.List
}

func (s *ListSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubSceneResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListSubSceneResponseBody) SetAccessDeniedDetail(v *ListSubSceneResponseBodyAccessDeniedDetail) *ListSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListSubSceneResponseBody) SetCode(v int64) *ListSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCount(v int64) *ListSubSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCurrentPage(v int64) *ListSubSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSubSceneResponseBody) SetHasNext(v bool) *ListSubSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSubSceneResponseBody) SetList(v []*ListSubSceneResponseBodyList) *ListSubSceneResponseBody {
	s.List = v
	return s
}

func (s *ListSubSceneResponseBody) SetMessage(v string) *ListSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubSceneResponseBody) SetRequestId(v string) *ListSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubSceneResponseBody) SetSuccess(v bool) *ListSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubSceneResponseBody) SetTotalPage(v int64) *ListSubSceneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSubSceneResponseBody) Validate() error {
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

type ListSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListSubSceneResponseBodyList struct {
	// example:
	//
	// https:/image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	BaseImageUrl *string `json:"BaseImageUrl,omitempty" xml:"BaseImageUrl,omitempty"`
	// example:
	//
	// https:/image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	CoverUrl    *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CubemapPath *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
	// example:
	//
	// true
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
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
	// xsfwsddd==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {}
	LayoutData *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	// example:
	//
	// c俄式
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https:/image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	OriginUrl *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	// example:
	//
	// 2345****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// xxxx.jpg
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https:/image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSubSceneResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSubSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBodyList) GetBaseImageUrl() *string {
	return s.BaseImageUrl
}

func (s *ListSubSceneResponseBodyList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListSubSceneResponseBodyList) GetCubemapPath() *string {
	return s.CubemapPath
}

func (s *ListSubSceneResponseBodyList) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListSubSceneResponseBodyList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSubSceneResponseBodyList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListSubSceneResponseBodyList) GetId() *string {
	return s.Id
}

func (s *ListSubSceneResponseBodyList) GetLayoutData() *string {
	return s.LayoutData
}

func (s *ListSubSceneResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListSubSceneResponseBodyList) GetOriginUrl() *string {
	return s.OriginUrl
}

func (s *ListSubSceneResponseBodyList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListSubSceneResponseBodyList) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListSubSceneResponseBodyList) GetStatus() *int64 {
	return s.Status
}

func (s *ListSubSceneResponseBodyList) GetType() *string {
	return s.Type
}

func (s *ListSubSceneResponseBodyList) GetUrl() *string {
	return s.Url
}

func (s *ListSubSceneResponseBodyList) SetBaseImageUrl(v string) *ListSubSceneResponseBodyList {
	s.BaseImageUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetCoverUrl(v string) *ListSubSceneResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetCubemapPath(v string) *ListSubSceneResponseBodyList {
	s.CubemapPath = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetDeleted(v bool) *ListSubSceneResponseBodyList {
	s.Deleted = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetGmtCreate(v int64) *ListSubSceneResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetGmtModified(v int64) *ListSubSceneResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetId(v string) *ListSubSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetLayoutData(v string) *ListSubSceneResponseBodyList {
	s.LayoutData = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetName(v string) *ListSubSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetOriginUrl(v string) *ListSubSceneResponseBodyList {
	s.OriginUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetResourceId(v string) *ListSubSceneResponseBodyList {
	s.ResourceId = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetResourceName(v string) *ListSubSceneResponseBodyList {
	s.ResourceName = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetStatus(v int64) *ListSubSceneResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetType(v string) *ListSubSceneResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetUrl(v string) *ListSubSceneResponseBodyList {
	s.Url = &v
	return s
}

func (s *ListSubSceneResponseBodyList) Validate() error {
	return dara.Validate(s)
}
