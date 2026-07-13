// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAtiRegistrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) *DeleteAtiRegistrantResponseBody
	GetAccessDeniedDetail() *DeleteAtiRegistrantResponseBodyAccessDeniedDetail
	SetRequestId(v string) *DeleteAtiRegistrantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAtiRegistrantResponseBody
	GetSuccess() *bool
}

type DeleteAtiRegistrantResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *DeleteAtiRegistrantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 853805EA-3D47-47D5-9A1A-A45C24313ABD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation to delete the domain instance configuration was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAtiRegistrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiRegistrantResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAtiRegistrantResponseBody) GetAccessDeniedDetail() *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteAtiRegistrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAtiRegistrantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAtiRegistrantResponseBody) SetAccessDeniedDetail(v *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) *DeleteAtiRegistrantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteAtiRegistrantResponseBody) SetRequestId(v string) *DeleteAtiRegistrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBody) SetSuccess(v bool) *DeleteAtiRegistrantResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAtiRegistrantResponseBodyAccessDeniedDetail struct {
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
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit denial.
	//
	// - ImplicitDeny: implicit denial.
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

func (s DeleteAtiRegistrantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteAtiRegistrantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteAtiRegistrantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
