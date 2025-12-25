// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScenePreviewInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetScenePreviewInfoResponseBodyAccessDeniedDetail) *GetScenePreviewInfoResponseBody
	GetAccessDeniedDetail() *GetScenePreviewInfoResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetScenePreviewInfoResponseBody
	GetCode() *int64
	SetData(v *GetScenePreviewInfoResponseBodyData) *GetScenePreviewInfoResponseBody
	GetData() *GetScenePreviewInfoResponseBodyData
	SetMessage(v string) *GetScenePreviewInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScenePreviewInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScenePreviewInfoResponseBody
	GetSuccess() *bool
}

type GetScenePreviewInfoResponseBody struct {
	AccessDeniedDetail *GetScenePreviewInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 0：成功，其他：失败
	Code *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScenePreviewInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBody) GetAccessDeniedDetail() *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetScenePreviewInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetScenePreviewInfoResponseBody) GetData() *GetScenePreviewInfoResponseBodyData {
	return s.Data
}

func (s *GetScenePreviewInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScenePreviewInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScenePreviewInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScenePreviewInfoResponseBody) SetAccessDeniedDetail(v *GetScenePreviewInfoResponseBodyAccessDeniedDetail) *GetScenePreviewInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetCode(v int64) *GetScenePreviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetData(v *GetScenePreviewInfoResponseBodyData) *GetScenePreviewInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetMessage(v string) *GetScenePreviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetRequestId(v string) *GetScenePreviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetSuccess(v bool) *GetScenePreviewInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScenePreviewInfoResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetScenePreviewInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetScenePreviewInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetScenePreviewInfoResponseBodyData struct {
	// example:
	//
	// https://www.aliyundoc.com/sgm/A.e.QRQRLWYEHIUEYLYW/A.e.QRQRLWYEHIUEYLYW****.sgm
	ModelPath        *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	PanoList         *string `json:"PanoList,omitempty" xml:"PanoList,omitempty"`
	TextureModelPath *string `json:"TextureModelPath,omitempty" xml:"TextureModelPath,omitempty"`
	TexturePanoPath  *string `json:"TexturePanoPath,omitempty" xml:"TexturePanoPath,omitempty"`
}

func (s GetScenePreviewInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScenePreviewInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBodyData) GetModelPath() *string {
	return s.ModelPath
}

func (s *GetScenePreviewInfoResponseBodyData) GetPanoList() *string {
	return s.PanoList
}

func (s *GetScenePreviewInfoResponseBodyData) GetTextureModelPath() *string {
	return s.TextureModelPath
}

func (s *GetScenePreviewInfoResponseBodyData) GetTexturePanoPath() *string {
	return s.TexturePanoPath
}

func (s *GetScenePreviewInfoResponseBodyData) SetModelPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.ModelPath = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetPanoList(v string) *GetScenePreviewInfoResponseBodyData {
	s.PanoList = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetTextureModelPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.TextureModelPath = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetTexturePanoPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.TexturePanoPath = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
