// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiRegistrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) *UpdateAtiRegistrantResponseBody
	GetAccessDeniedDetail() *UpdateAtiRegistrantResponseBodyAccessDeniedDetail
	SetRequestId(v string) *UpdateAtiRegistrantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAtiRegistrantResponseBody
	GetSuccess() *bool
}

type UpdateAtiRegistrantResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *UpdateAtiRegistrantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAtiRegistrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiRegistrantResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAtiRegistrantResponseBody) GetAccessDeniedDetail() *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateAtiRegistrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAtiRegistrantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAtiRegistrantResponseBody) SetAccessDeniedDetail(v *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) *UpdateAtiRegistrantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateAtiRegistrantResponseBody) SetRequestId(v string) *UpdateAtiRegistrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBody) SetSuccess(v bool) *UpdateAtiRegistrantResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAtiRegistrantResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// AddRspDomainServerHoldStatusForGateway
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
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
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

func (s UpdateAtiRegistrantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateAtiRegistrantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
