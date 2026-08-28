// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAtiAgentRegisterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *SubmitAtiAgentRegisterInfoResponseBody
	GetAccessDeniedDetail() *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail
	SetRequestId(v string) *SubmitAtiAgentRegisterInfoResponseBody
	GetRequestId() *string
	SetStatus(v bool) *SubmitAtiAgentRegisterInfoResponseBody
	GetStatus() *bool
}

type SubmitAtiAgentRegisterInfoResponseBody struct {
	// The access denied details. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The unique request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Agent status. Valid values:
	//
	// - Draft: The Agent registration form is being filled in and has not been formally submitted. In the Draft state, only modification and detail viewing operations are supported. Other operations are not supported.
	//
	// - Private CA Pending Issuance: The Agent registration has been formally submitted. Alibaba Cloud has completed the ACME DNS-01 pre-check and submitted the registration information and generated DNS records to CNNIC. The system is waiting for CNNIC to approve and issue the Private CA and complete TL sealing.
	//
	// - DNS Pending Verification: CNNIC has approved the request, issued the Private CA certificate, and completed TL sealing, but the DNS records of the user have not been verified. The user needs to add the corresponding DNS records in domain name resolution and complete verification.
	//
	// - Active: All processes are complete. The Private CA certificate has been issued, TL has been sealed, and DNS records have been verified. The Agent is activated and can be discovered and trusted across the network.
	//
	// - Expired: The Agent identity certificate has expired, and the user did not complete certificate renewal within the validity period.
	//
	// - Revoked: The Agent certificate has been revoked, DNS records have been cleaned up, and the Agent cannot be discovered or trusted. It cannot be restored to the Active state.
	//
	// example:
	//
	// Private CA 待签发
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitAtiAgentRegisterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAtiAgentRegisterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) GetAccessDeniedDetail() *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) SetAccessDeniedDetail(v *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *SubmitAtiAgentRegisterInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) SetRequestId(v string) *SubmitAtiAgentRegisterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) SetStatus(v bool) *SubmitAtiAgentRegisterInfoResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail struct {
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
	// The encrypted complete diagnostic message.
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

func (s SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
