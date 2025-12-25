// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveModelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SaveModelConfigResponseBodyAccessDeniedDetail) *SaveModelConfigResponseBody
	GetAccessDeniedDetail() *SaveModelConfigResponseBodyAccessDeniedDetail
	SetCode(v int64) *SaveModelConfigResponseBody
	GetCode() *int64
	SetMessage(v string) *SaveModelConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveModelConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveModelConfigResponseBody
	GetSuccess() *bool
}

type SaveModelConfigResponseBody struct {
	AccessDeniedDetail *SaveModelConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B28A2ECB-AB29-1E01****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveModelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveModelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveModelConfigResponseBody) GetAccessDeniedDetail() *SaveModelConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SaveModelConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SaveModelConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveModelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveModelConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveModelConfigResponseBody) SetAccessDeniedDetail(v *SaveModelConfigResponseBodyAccessDeniedDetail) *SaveModelConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SaveModelConfigResponseBody) SetCode(v int64) *SaveModelConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetMessage(v string) *SaveModelConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetRequestId(v string) *SaveModelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetSuccess(v bool) *SaveModelConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveModelConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveModelConfigResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SaveModelConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveModelConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SaveModelConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SaveModelConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
