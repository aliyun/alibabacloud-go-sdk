// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMinimapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SaveMinimapResponseBodyAccessDeniedDetail) *SaveMinimapResponseBody
	GetAccessDeniedDetail() *SaveMinimapResponseBodyAccessDeniedDetail
	SetCode(v int64) *SaveMinimapResponseBody
	GetCode() *int64
	SetMessage(v string) *SaveMinimapResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveMinimapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveMinimapResponseBody
	GetSuccess() *bool
}

type SaveMinimapResponseBody struct {
	AccessDeniedDetail *SaveMinimapResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *int64                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveMinimapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveMinimapResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMinimapResponseBody) GetAccessDeniedDetail() *SaveMinimapResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SaveMinimapResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SaveMinimapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveMinimapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveMinimapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveMinimapResponseBody) SetAccessDeniedDetail(v *SaveMinimapResponseBodyAccessDeniedDetail) *SaveMinimapResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SaveMinimapResponseBody) SetCode(v int64) *SaveMinimapResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMinimapResponseBody) SetMessage(v string) *SaveMinimapResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMinimapResponseBody) SetRequestId(v string) *SaveMinimapResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMinimapResponseBody) SetSuccess(v bool) *SaveMinimapResponseBody {
	s.Success = &v
	return s
}

func (s *SaveMinimapResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveMinimapResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SaveMinimapResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveMinimapResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SaveMinimapResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SaveMinimapResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
