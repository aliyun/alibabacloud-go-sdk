// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *TempPreviewResponseBodyAccessDeniedDetail) *TempPreviewResponseBody
	GetAccessDeniedDetail() *TempPreviewResponseBodyAccessDeniedDetail
	SetCode(v int64) *TempPreviewResponseBody
	GetCode() *int64
	SetMessage(v string) *TempPreviewResponseBody
	GetMessage() *string
	SetPreviewUrl(v string) *TempPreviewResponseBody
	GetPreviewUrl() *string
	SetRequestId(v string) *TempPreviewResponseBody
	GetRequestId() *string
	SetSceneId(v string) *TempPreviewResponseBody
	GetSceneId() *string
	SetSuccess(v bool) *TempPreviewResponseBody
	GetSuccess() *bool
}

type TempPreviewResponseBody struct {
	AccessDeniedDetail *TempPreviewResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://preview-lyj.aliyuncs.com/preview/temp/xxx****
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TempPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewResponseBody) GetAccessDeniedDetail() *TempPreviewResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *TempPreviewResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TempPreviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TempPreviewResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *TempPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempPreviewResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *TempPreviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TempPreviewResponseBody) SetAccessDeniedDetail(v *TempPreviewResponseBodyAccessDeniedDetail) *TempPreviewResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *TempPreviewResponseBody) SetCode(v int64) *TempPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewResponseBody) SetMessage(v string) *TempPreviewResponseBody {
	s.Message = &v
	return s
}

func (s *TempPreviewResponseBody) SetPreviewUrl(v string) *TempPreviewResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *TempPreviewResponseBody) SetRequestId(v string) *TempPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewResponseBody) SetSceneId(v string) *TempPreviewResponseBody {
	s.SceneId = &v
	return s
}

func (s *TempPreviewResponseBody) SetSuccess(v bool) *TempPreviewResponseBody {
	s.Success = &v
	return s
}

func (s *TempPreviewResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TempPreviewResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s TempPreviewResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetAuthAction(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) SetPolicyType(v string) *TempPreviewResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *TempPreviewResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
