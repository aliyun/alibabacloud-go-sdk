// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PublishSceneResponseBodyAccessDeniedDetail) *PublishSceneResponseBody
	GetAccessDeniedDetail() *PublishSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *PublishSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *PublishSceneResponseBody
	GetMessage() *string
	SetPreviewUrl(v string) *PublishSceneResponseBody
	GetPreviewUrl() *string
	SetRequestId(v string) *PublishSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishSceneResponseBody
	GetSuccess() *bool
}

type PublishSceneResponseBody struct {
	AccessDeniedDetail *PublishSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// preview-lyj.aliyuncs.com/preview/xxx****
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishSceneResponseBody) GoString() string {
	return s.String()
}

func (s *PublishSceneResponseBody) GetAccessDeniedDetail() *PublishSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PublishSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PublishSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishSceneResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *PublishSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishSceneResponseBody) SetAccessDeniedDetail(v *PublishSceneResponseBodyAccessDeniedDetail) *PublishSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PublishSceneResponseBody) SetCode(v int64) *PublishSceneResponseBody {
	s.Code = &v
	return s
}

func (s *PublishSceneResponseBody) SetMessage(v string) *PublishSceneResponseBody {
	s.Message = &v
	return s
}

func (s *PublishSceneResponseBody) SetPreviewUrl(v string) *PublishSceneResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PublishSceneResponseBody) SetRequestId(v string) *PublishSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishSceneResponseBody) SetSuccess(v bool) *PublishSceneResponseBody {
	s.Success = &v
	return s
}

func (s *PublishSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PublishSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PublishSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PublishSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PublishSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
