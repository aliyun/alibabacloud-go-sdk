// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *DeleteSchedulerxNotificationPolicyResponseBody
	GetAccessDeniedDetail() *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail
	SetCode(v int32) *DeleteSchedulerxNotificationPolicyResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteSchedulerxNotificationPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSchedulerxNotificationPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSchedulerxNotificationPolicyResponseBody
	GetSuccess() *bool
}

type DeleteSchedulerxNotificationPolicyResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. Returned only when an error occurs.
	//
	// example:
	//
	// Invalid parameter: policyName cannot be null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// C8E5FB4A-6D8D-424D-9AAA-4FE06BB74FF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSchedulerxNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) GetAccessDeniedDetail() *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) SetAccessDeniedDetail(v *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) *DeleteSchedulerxNotificationPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) SetCode(v int32) *DeleteSchedulerxNotificationPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) SetMessage(v string) *DeleteSchedulerxNotificationPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) SetRequestId(v string) *DeleteSchedulerxNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) SetSuccess(v bool) *DeleteSchedulerxNotificationPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:DeleteSchedulerxNotificationPolicy
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
	// AQFoiYKaaImCmkI3QjgzNzM1LTIzQTUtNURENy1COUQ3LTBFOTIxRkRFOEM3NQ==
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

func (s DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
