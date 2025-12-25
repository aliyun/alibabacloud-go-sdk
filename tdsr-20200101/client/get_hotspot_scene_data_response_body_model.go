// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotSceneDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetHotspotSceneDataResponseBodyAccessDeniedDetail) *GetHotspotSceneDataResponseBody
	GetAccessDeniedDetail() *GetHotspotSceneDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetHotspotSceneDataResponseBody
	GetCode() *int64
	SetData(v *GetHotspotSceneDataResponseBodyData) *GetHotspotSceneDataResponseBody
	GetData() *GetHotspotSceneDataResponseBodyData
	SetMessage(v string) *GetHotspotSceneDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotSceneDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotSceneDataResponseBody
	GetSuccess() *bool
}

type GetHotspotSceneDataResponseBody struct {
	AccessDeniedDetail *GetHotspotSceneDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHotspotSceneDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotSceneDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotSceneDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBody) GetAccessDeniedDetail() *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetHotspotSceneDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetHotspotSceneDataResponseBody) GetData() *GetHotspotSceneDataResponseBodyData {
	return s.Data
}

func (s *GetHotspotSceneDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotSceneDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotSceneDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotSceneDataResponseBody) SetAccessDeniedDetail(v *GetHotspotSceneDataResponseBodyAccessDeniedDetail) *GetHotspotSceneDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetCode(v int64) *GetHotspotSceneDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetData(v *GetHotspotSceneDataResponseBodyData) *GetHotspotSceneDataResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetMessage(v string) *GetHotspotSceneDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetRequestId(v string) *GetHotspotSceneDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetSuccess(v bool) *GetHotspotSceneDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) Validate() error {
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

type GetHotspotSceneDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetHotspotSceneDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotSceneDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetHotspotSceneDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetHotspotSceneDataResponseBodyData struct {
	// example:
	//
	// A.e.QRQRLWYEHIUE****
	ModelToken  *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	PreviewData *string `json:"PreviewData,omitempty" xml:"PreviewData,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb130936****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// example:
	//
	// MODEL_3D
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetHotspotSceneDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotSceneDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBodyData) GetModelToken() *string {
	return s.ModelToken
}

func (s *GetHotspotSceneDataResponseBodyData) GetPreviewData() *string {
	return s.PreviewData
}

func (s *GetHotspotSceneDataResponseBodyData) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetHotspotSceneDataResponseBodyData) GetSceneType() *string {
	return s.SceneType
}

func (s *GetHotspotSceneDataResponseBodyData) SetModelToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.ModelToken = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewData(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewData = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetSceneType(v string) *GetHotspotSceneDataResponseBodyData {
	s.SceneType = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
