// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAtiAgentRegisterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *RevokeAtiAgentRegisterInfoResponseBody
	GetAccessDeniedDetail() *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail
	SetRequestId(v string) *RevokeAtiAgentRegisterInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeAtiAgentRegisterInfoResponseBody
	GetSuccess() *bool
}

type RevokeAtiAgentRegisterInfoResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeAtiAgentRegisterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiAgentRegisterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) GetAccessDeniedDetail() *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) SetAccessDeniedDetail(v *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *RevokeAtiAgentRegisterInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) SetRequestId(v string) *RevokeAtiAgentRegisterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) SetSuccess(v bool) *RevokeAtiAgentRegisterInfoResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGatewayOte
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorized principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorized principal.
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
	// - ExplicitDeny: Explicit deny.
	//
	// - ImplicitDeny: Implicit deny.
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

func (s RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
