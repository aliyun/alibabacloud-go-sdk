// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DetailSubSceneResponseBodyAccessDeniedDetail) *DetailSubSceneResponseBody
	GetAccessDeniedDetail() *DetailSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *DetailSubSceneResponseBody
	GetCode() *int64
	SetCoverUrl(v string) *DetailSubSceneResponseBody
	GetCoverUrl() *string
	SetCubemapPath(v string) *DetailSubSceneResponseBody
	GetCubemapPath() *string
	SetGmtCreate(v int64) *DetailSubSceneResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *DetailSubSceneResponseBody
	GetGmtModified() *int64
	SetId(v string) *DetailSubSceneResponseBody
	GetId() *string
	SetImageUrl(v string) *DetailSubSceneResponseBody
	GetImageUrl() *string
	SetLayoutData(v string) *DetailSubSceneResponseBody
	GetLayoutData() *string
	SetMessage(v string) *DetailSubSceneResponseBody
	GetMessage() *string
	SetName(v string) *DetailSubSceneResponseBody
	GetName() *string
	SetOriginUrl(v string) *DetailSubSceneResponseBody
	GetOriginUrl() *string
	SetPosition(v string) *DetailSubSceneResponseBody
	GetPosition() *string
	SetRequestId(v string) *DetailSubSceneResponseBody
	GetRequestId() *string
	SetResourceId(v string) *DetailSubSceneResponseBody
	GetResourceId() *string
	SetStatus(v int64) *DetailSubSceneResponseBody
	GetStatus() *int64
	SetSuccess(v bool) *DetailSubSceneResponseBody
	GetSuccess() *bool
	SetType(v string) *DetailSubSceneResponseBody
	GetType() *string
	SetUrl(v string) *DetailSubSceneResponseBody
	GetUrl() *string
}

type DetailSubSceneResponseBody struct {
	AccessDeniedDetail *DetailSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// https://image-demo.oss-cn-hangzhou.aliyuncs.com/cubemap/****
	CubemapPath *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
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
	// 1234***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://www.aliyun.com/test1.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// {}
	LayoutData *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http://www.aliyun.com/test.jpg
	OriginUrl *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	// example:
	//
	// [-0.8928,-0.21467,0.39603]
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://image-demo.oss-cn-hangzhou.aliyuncs.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetailSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponseBody) GetAccessDeniedDetail() *DetailSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DetailSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DetailSubSceneResponseBody) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *DetailSubSceneResponseBody) GetCubemapPath() *string {
	return s.CubemapPath
}

func (s *DetailSubSceneResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DetailSubSceneResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DetailSubSceneResponseBody) GetId() *string {
	return s.Id
}

func (s *DetailSubSceneResponseBody) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetailSubSceneResponseBody) GetLayoutData() *string {
	return s.LayoutData
}

func (s *DetailSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetailSubSceneResponseBody) GetName() *string {
	return s.Name
}

func (s *DetailSubSceneResponseBody) GetOriginUrl() *string {
	return s.OriginUrl
}

func (s *DetailSubSceneResponseBody) GetPosition() *string {
	return s.Position
}

func (s *DetailSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailSubSceneResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DetailSubSceneResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *DetailSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetailSubSceneResponseBody) GetType() *string {
	return s.Type
}

func (s *DetailSubSceneResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DetailSubSceneResponseBody) SetAccessDeniedDetail(v *DetailSubSceneResponseBodyAccessDeniedDetail) *DetailSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DetailSubSceneResponseBody) SetCode(v int64) *DetailSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCoverUrl(v string) *DetailSubSceneResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCubemapPath(v string) *DetailSubSceneResponseBody {
	s.CubemapPath = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetGmtCreate(v int64) *DetailSubSceneResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetGmtModified(v int64) *DetailSubSceneResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetId(v string) *DetailSubSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetImageUrl(v string) *DetailSubSceneResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetLayoutData(v string) *DetailSubSceneResponseBody {
	s.LayoutData = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetMessage(v string) *DetailSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetName(v string) *DetailSubSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetOriginUrl(v string) *DetailSubSceneResponseBody {
	s.OriginUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetPosition(v string) *DetailSubSceneResponseBody {
	s.Position = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetRequestId(v string) *DetailSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetResourceId(v string) *DetailSubSceneResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetStatus(v int64) *DetailSubSceneResponseBody {
	s.Status = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetSuccess(v bool) *DetailSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetType(v string) *DetailSubSceneResponseBody {
	s.Type = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetUrl(v string) *DetailSubSceneResponseBody {
	s.Url = &v
	return s
}

func (s *DetailSubSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetailSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetailSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DetailSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DetailSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DetailSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
