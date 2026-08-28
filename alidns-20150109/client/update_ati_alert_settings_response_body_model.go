// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiAlertSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) *UpdateAtiAlertSettingsResponseBody
	GetAccessDeniedDetail() *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail
	SetRequestId(v string) *UpdateAtiAlertSettingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAtiAlertSettingsResponseBody
	GetSuccess() *bool
}

type UpdateAtiAlertSettingsResponseBody struct {
	// The details about the access denial. This parameter is returned only when RAM authentication fails.
	AccessDeniedDetail *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s UpdateAtiAlertSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAlertSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAtiAlertSettingsResponseBody) GetAccessDeniedDetail() *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateAtiAlertSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAtiAlertSettingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAtiAlertSettingsResponseBody) SetAccessDeniedDetail(v *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) *UpdateAtiAlertSettingsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBody) SetRequestId(v string) *UpdateAtiAlertSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBody) SetSuccess(v bool) *UpdateAtiAlertSettingsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGateway
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the authorization principal owner.
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

func (s UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
