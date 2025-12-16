// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *ManageSchedulerxNotificationPolicyResponseBody
	GetAccessDeniedDetail() *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail
	SetCode(v int32) *ManageSchedulerxNotificationPolicyResponseBody
	GetCode() *int32
	SetMessage(v string) *ManageSchedulerxNotificationPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManageSchedulerxNotificationPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ManageSchedulerxNotificationPolicyResponseBody
	GetSuccess() *bool
}

type ManageSchedulerxNotificationPolicyResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message, returned only when an error occurs.
	//
	// example:
	//
	// Invalid parameter: policyName cannot be null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManageSchedulerxNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) GetAccessDeniedDetail() *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) SetAccessDeniedDetail(v *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *ManageSchedulerxNotificationPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) SetCode(v int32) *ManageSchedulerxNotificationPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) SetMessage(v string) *ManageSchedulerxNotificationPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) SetRequestId(v string) *ManageSchedulerxNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) SetSuccess(v bool) *ManageSchedulerxNotificationPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:ManageSchedulerxNotificationPolicy
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

func (s ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
