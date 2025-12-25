// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DetailSceneResponseBodyAccessDeniedDetail) *DetailSceneResponseBody
	GetAccessDeniedDetail() *DetailSceneResponseBodyAccessDeniedDetail
	SetCaptures(v []*DetailSceneResponseBodyCaptures) *DetailSceneResponseBody
	GetCaptures() []*DetailSceneResponseBodyCaptures
	SetCode(v int64) *DetailSceneResponseBody
	GetCode() *int64
	SetCoverUrl(v string) *DetailSceneResponseBody
	GetCoverUrl() *string
	SetFloorPlans(v []*DetailSceneResponseBodyFloorPlans) *DetailSceneResponseBody
	GetFloorPlans() []*DetailSceneResponseBodyFloorPlans
	SetGmtCreate(v int64) *DetailSceneResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *DetailSceneResponseBody
	GetGmtModified() *int64
	SetId(v string) *DetailSceneResponseBody
	GetId() *string
	SetMessage(v string) *DetailSceneResponseBody
	GetMessage() *string
	SetName(v string) *DetailSceneResponseBody
	GetName() *string
	SetPreviewToken(v string) *DetailSceneResponseBody
	GetPreviewToken() *string
	SetPublished(v bool) *DetailSceneResponseBody
	GetPublished() *bool
	SetRequestId(v string) *DetailSceneResponseBody
	GetRequestId() *string
	SetSourceNum(v int64) *DetailSceneResponseBody
	GetSourceNum() *int64
	SetStatus(v string) *DetailSceneResponseBody
	GetStatus() *string
	SetStatusName(v string) *DetailSceneResponseBody
	GetStatusName() *string
	SetSubSceneNum(v int64) *DetailSceneResponseBody
	GetSubSceneNum() *int64
	SetSuccess(v bool) *DetailSceneResponseBody
	GetSuccess() *bool
	SetType(v string) *DetailSceneResponseBody
	GetType() *string
}

type DetailSceneResponseBody struct {
	AccessDeniedDetail *DetailSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Captures           []*DetailSceneResponseBodyCaptures         `json:"Captures,omitempty" xml:"Captures,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// www.example.com/xxx/xxx.jpg
	CoverUrl   *string                              `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	FloorPlans []*DetailSceneResponseBodyFloorPlans `json:"FloorPlans,omitempty" xml:"FloorPlans,omitempty" type:"Repeated"`
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// d989623696ab4f87a80b8d5b0b0****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// example:
	//
	// false
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb130936****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// MODEL_3D
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetailSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBody) GetAccessDeniedDetail() *DetailSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DetailSceneResponseBody) GetCaptures() []*DetailSceneResponseBodyCaptures {
	return s.Captures
}

func (s *DetailSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DetailSceneResponseBody) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *DetailSceneResponseBody) GetFloorPlans() []*DetailSceneResponseBodyFloorPlans {
	return s.FloorPlans
}

func (s *DetailSceneResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DetailSceneResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DetailSceneResponseBody) GetId() *string {
	return s.Id
}

func (s *DetailSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetailSceneResponseBody) GetName() *string {
	return s.Name
}

func (s *DetailSceneResponseBody) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *DetailSceneResponseBody) GetPublished() *bool {
	return s.Published
}

func (s *DetailSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailSceneResponseBody) GetSourceNum() *int64 {
	return s.SourceNum
}

func (s *DetailSceneResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DetailSceneResponseBody) GetStatusName() *string {
	return s.StatusName
}

func (s *DetailSceneResponseBody) GetSubSceneNum() *int64 {
	return s.SubSceneNum
}

func (s *DetailSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetailSceneResponseBody) GetType() *string {
	return s.Type
}

func (s *DetailSceneResponseBody) SetAccessDeniedDetail(v *DetailSceneResponseBodyAccessDeniedDetail) *DetailSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DetailSceneResponseBody) SetCaptures(v []*DetailSceneResponseBodyCaptures) *DetailSceneResponseBody {
	s.Captures = v
	return s
}

func (s *DetailSceneResponseBody) SetCode(v int64) *DetailSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSceneResponseBody) SetCoverUrl(v string) *DetailSceneResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *DetailSceneResponseBody) SetFloorPlans(v []*DetailSceneResponseBodyFloorPlans) *DetailSceneResponseBody {
	s.FloorPlans = v
	return s
}

func (s *DetailSceneResponseBody) SetGmtCreate(v int64) *DetailSceneResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailSceneResponseBody) SetGmtModified(v int64) *DetailSceneResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailSceneResponseBody) SetId(v string) *DetailSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSceneResponseBody) SetMessage(v string) *DetailSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSceneResponseBody) SetName(v string) *DetailSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSceneResponseBody) SetPreviewToken(v string) *DetailSceneResponseBody {
	s.PreviewToken = &v
	return s
}

func (s *DetailSceneResponseBody) SetPublished(v bool) *DetailSceneResponseBody {
	s.Published = &v
	return s
}

func (s *DetailSceneResponseBody) SetRequestId(v string) *DetailSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSceneResponseBody) SetSourceNum(v int64) *DetailSceneResponseBody {
	s.SourceNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetStatus(v string) *DetailSceneResponseBody {
	s.Status = &v
	return s
}

func (s *DetailSceneResponseBody) SetStatusName(v string) *DetailSceneResponseBody {
	s.StatusName = &v
	return s
}

func (s *DetailSceneResponseBody) SetSubSceneNum(v int64) *DetailSceneResponseBody {
	s.SubSceneNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetSuccess(v bool) *DetailSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSceneResponseBody) SetType(v string) *DetailSceneResponseBody {
	s.Type = &v
	return s
}

func (s *DetailSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Captures != nil {
		for _, item := range s.Captures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FloorPlans != nil {
		for _, item := range s.FloorPlans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetailSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetailSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DetailSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DetailSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DetailSceneResponseBodyCaptures struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.aliyun.com/test1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetailSceneResponseBodyCaptures) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneResponseBodyCaptures) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBodyCaptures) GetTitle() *string {
	return s.Title
}

func (s *DetailSceneResponseBodyCaptures) GetUrl() *string {
	return s.Url
}

func (s *DetailSceneResponseBodyCaptures) SetTitle(v string) *DetailSceneResponseBodyCaptures {
	s.Title = &v
	return s
}

func (s *DetailSceneResponseBodyCaptures) SetUrl(v string) *DetailSceneResponseBodyCaptures {
	s.Url = &v
	return s
}

func (s *DetailSceneResponseBodyCaptures) Validate() error {
	return dara.Validate(s)
}

type DetailSceneResponseBodyFloorPlans struct {
	ColorMapUrl *string `json:"ColorMapUrl,omitempty" xml:"ColorMapUrl,omitempty"`
	FloorLabel  *string `json:"FloorLabel,omitempty" xml:"FloorLabel,omitempty"`
	FloorName   *string `json:"FloorName,omitempty" xml:"FloorName,omitempty"`
	MiniMapUrl  *string `json:"MiniMapUrl,omitempty" xml:"MiniMapUrl,omitempty"`
}

func (s DetailSceneResponseBodyFloorPlans) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneResponseBodyFloorPlans) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBodyFloorPlans) GetColorMapUrl() *string {
	return s.ColorMapUrl
}

func (s *DetailSceneResponseBodyFloorPlans) GetFloorLabel() *string {
	return s.FloorLabel
}

func (s *DetailSceneResponseBodyFloorPlans) GetFloorName() *string {
	return s.FloorName
}

func (s *DetailSceneResponseBodyFloorPlans) GetMiniMapUrl() *string {
	return s.MiniMapUrl
}

func (s *DetailSceneResponseBodyFloorPlans) SetColorMapUrl(v string) *DetailSceneResponseBodyFloorPlans {
	s.ColorMapUrl = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetFloorLabel(v string) *DetailSceneResponseBodyFloorPlans {
	s.FloorLabel = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetFloorName(v string) *DetailSceneResponseBodyFloorPlans {
	s.FloorName = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetMiniMapUrl(v string) *DetailSceneResponseBodyFloorPlans {
	s.MiniMapUrl = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) Validate() error {
	return dara.Validate(s)
}
