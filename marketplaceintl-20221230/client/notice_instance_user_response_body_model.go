// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNoticeInstanceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *NoticeInstanceUserResponseBodyAccessDeniedDetail) *NoticeInstanceUserResponseBody
	GetAccessDeniedDetail() *NoticeInstanceUserResponseBodyAccessDeniedDetail
	SetCode(v string) *NoticeInstanceUserResponseBody
	GetCode() *string
	SetMessage(v string) *NoticeInstanceUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *NoticeInstanceUserResponseBody
	GetRequestId() *string
	SetResult(v bool) *NoticeInstanceUserResponseBody
	GetResult() *bool
	SetSuccess(v bool) *NoticeInstanceUserResponseBody
	GetSuccess() *bool
}

type NoticeInstanceUserResponseBody struct {
	AccessDeniedDetail *NoticeInstanceUserResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result             *bool                                             `json:"Result,omitempty" xml:"Result,omitempty"`
	Success            *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NoticeInstanceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NoticeInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponseBody) GetAccessDeniedDetail() *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *NoticeInstanceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *NoticeInstanceUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *NoticeInstanceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NoticeInstanceUserResponseBody) GetResult() *bool {
	return s.Result
}

func (s *NoticeInstanceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *NoticeInstanceUserResponseBody) SetAccessDeniedDetail(v *NoticeInstanceUserResponseBodyAccessDeniedDetail) *NoticeInstanceUserResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetCode(v string) *NoticeInstanceUserResponseBody {
	s.Code = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetMessage(v string) *NoticeInstanceUserResponseBody {
	s.Message = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetRequestId(v string) *NoticeInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetResult(v bool) *NoticeInstanceUserResponseBody {
	s.Result = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) SetSuccess(v bool) *NoticeInstanceUserResponseBody {
	s.Success = &v
	return s
}

func (s *NoticeInstanceUserResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NoticeInstanceUserResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s NoticeInstanceUserResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s NoticeInstanceUserResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthAction(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) SetPolicyType(v string) *NoticeInstanceUserResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *NoticeInstanceUserResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
