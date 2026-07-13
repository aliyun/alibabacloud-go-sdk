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
	AccessDeniedDetail *UpdateAtiAlertSettingsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGateway
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
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
