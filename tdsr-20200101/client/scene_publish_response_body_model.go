// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScenePublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ScenePublishResponseBodyAccessDeniedDetail) *ScenePublishResponseBody
	GetAccessDeniedDetail() *ScenePublishResponseBodyAccessDeniedDetail
	SetCode(v int64) *ScenePublishResponseBody
	GetCode() *int64
	SetMessage(v string) *ScenePublishResponseBody
	GetMessage() *string
	SetPreviewUrl(v string) *ScenePublishResponseBody
	GetPreviewUrl() *string
	SetRequestId(v string) *ScenePublishResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ScenePublishResponseBody
	GetSuccess() *bool
}

type ScenePublishResponseBody struct {
	AccessDeniedDetail *ScenePublishResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// https://lyj.aliyun.com/xxx
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb1309366493
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScenePublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScenePublishResponseBody) GoString() string {
	return s.String()
}

func (s *ScenePublishResponseBody) GetAccessDeniedDetail() *ScenePublishResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ScenePublishResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ScenePublishResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScenePublishResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *ScenePublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScenePublishResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ScenePublishResponseBody) SetAccessDeniedDetail(v *ScenePublishResponseBodyAccessDeniedDetail) *ScenePublishResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ScenePublishResponseBody) SetCode(v int64) *ScenePublishResponseBody {
	s.Code = &v
	return s
}

func (s *ScenePublishResponseBody) SetMessage(v string) *ScenePublishResponseBody {
	s.Message = &v
	return s
}

func (s *ScenePublishResponseBody) SetPreviewUrl(v string) *ScenePublishResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *ScenePublishResponseBody) SetRequestId(v string) *ScenePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScenePublishResponseBody) SetSuccess(v bool) *ScenePublishResponseBody {
	s.Success = &v
	return s
}

func (s *ScenePublishResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScenePublishResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ScenePublishResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ScenePublishResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ScenePublishResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ScenePublishResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
