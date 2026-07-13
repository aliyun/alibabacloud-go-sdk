// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAtiRegistrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) *RevokeAtiRegistrantResponseBody
	GetAccessDeniedDetail() *RevokeAtiRegistrantResponseBodyAccessDeniedDetail
	SetRequestId(v string) *RevokeAtiRegistrantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeAtiRegistrantResponseBody
	GetSuccess() *bool
}

type RevokeAtiRegistrantResponseBody struct {
	// The details about the access denial. This parameter is returned only when the RAM authentication fails.
	AccessDeniedDetail *RevokeAtiRegistrantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
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

func (s RevokeAtiRegistrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiRegistrantResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAtiRegistrantResponseBody) GetAccessDeniedDetail() *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RevokeAtiRegistrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeAtiRegistrantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeAtiRegistrantResponseBody) SetAccessDeniedDetail(v *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) *RevokeAtiRegistrantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RevokeAtiRegistrantResponseBody) SetRequestId(v string) *RevokeAtiRegistrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBody) SetSuccess(v bool) *RevokeAtiRegistrantResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevokeAtiRegistrantResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// UpdateRspDomainServerProhibitStatusForGatewayOte
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
	// 1046973331XXXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted complete diagnostic message.
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

func (s RevokeAtiRegistrantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RevokeAtiRegistrantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RevokeAtiRegistrantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
