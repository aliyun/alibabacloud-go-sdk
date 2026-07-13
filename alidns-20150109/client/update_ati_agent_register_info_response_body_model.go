// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiAgentRegisterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *UpdateAtiAgentRegisterInfoResponseBody
	GetAccessDeniedDetail() *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail
	SetRequestId(v string) *UpdateAtiAgentRegisterInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAtiAgentRegisterInfoResponseBody
	GetSuccess() *bool
}

type UpdateAtiAgentRegisterInfoResponseBody struct {
	// The access denial details. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAtiAgentRegisterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAgentRegisterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) GetAccessDeniedDetail() *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) SetAccessDeniedDetail(v *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *UpdateAtiAgentRegisterInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) SetRequestId(v string) *UpdateAtiAgentRegisterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) SetSuccess(v bool) *UpdateAtiAgentRegisterInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason why the authentication failed. Valid values:
	//
	// - ExplicitDeny: Explicit denial.
	//
	// - ImplicitDeny: Implicit denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
