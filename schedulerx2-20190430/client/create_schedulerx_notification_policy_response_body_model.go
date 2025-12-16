// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *CreateSchedulerxNotificationPolicyResponseBody
	GetAccessDeniedDetail() *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail
	SetCode(v int32) *CreateSchedulerxNotificationPolicyResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateSchedulerxNotificationPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSchedulerxNotificationPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSchedulerxNotificationPolicyResponseBody
	GetSuccess() *bool
}

type CreateSchedulerxNotificationPolicyResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information. Returned only when an error occurs.
	//
	// example:
	//
	// Invalid parameter: Notification policy already exists: test-weekdays
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchedulerxNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) GetAccessDeniedDetail() *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) SetAccessDeniedDetail(v *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *CreateSchedulerxNotificationPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) SetCode(v int32) *CreateSchedulerxNotificationPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) SetMessage(v string) *CreateSchedulerxNotificationPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) SetRequestId(v string) *CreateSchedulerxNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) SetSuccess(v bool) *CreateSchedulerxNotificationPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:CreateSchedulerxNotificationPolicy
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The account of the principal.
	//
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFn/cLPZ/3Cz0YxQkZBMjVGLTY0REUtNTlGNS05NzUwLTgyMUE4M0MwMTFDRQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The permission denial type.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
