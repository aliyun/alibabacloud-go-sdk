// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CancelIProductionJobResponseBodyAccessDeniedDetail) *CancelIProductionJobResponseBody
	GetAccessDeniedDetail() *CancelIProductionJobResponseBodyAccessDeniedDetail
	SetMessage(v string) *CancelIProductionJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelIProductionJobResponseBody
	GetRequestId() *string
}

type CancelIProductionJobResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *CancelIProductionJobResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobResponseBody) GetAccessDeniedDetail() *CancelIProductionJobResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CancelIProductionJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelIProductionJobResponseBody) SetAccessDeniedDetail(v *CancelIProductionJobResponseBodyAccessDeniedDetail) *CancelIProductionJobResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CancelIProductionJobResponseBody) SetMessage(v string) *CancelIProductionJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelIProductionJobResponseBody) SetRequestId(v string) *CancelIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelIProductionJobResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelIProductionJobResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// ice:CancelIProductionJob
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request. Valid values:
	//
	// - RAM user: RAM user UID
	//
	// - RAM role: RoleName:RoleSessionName
	//
	// - Federated: ProviderType/ProviderName
	//
	// example:
	//
	// ****4522705967****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The account to which the authenticate principal belongs.
	//
	// example:
	//
	// ****82303720****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the identity used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user
	//
	// - AssumedRoleUser: RAM role
	//
	// - Federated: SSO federated identity
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// ******AAZ/h8jzNEODc5QUUyLUZCOTAtNUQyQy1BMEFBLUUzODQxODUx******==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of denial by the access policy. Valid values:
	//
	// - **ImplicitDeny**: The resource owner has not configured a relevant permission policy for the current user. Access to unauthorized operations is denied by default.
	//
	// - **ExplicitDeny**: The RAM policy configured by the resource owner explicitly denies the current user access to the corresponding resource.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The type of the policy that caused the access denial. Valid values:
	//
	// - **ControlPolicy**: control policy.
	//
	// - **SessionPolicy**: an additional permission policy attached to a temporary token.
	//
	// - **AssumeRolePolicy**: the trust policy of a RAM role.
	//
	// - **AccountLevelIdentityBasedPolicy**: an identity-access policy at the account authorization scope, including custom policies and system policies.
	//
	// - **ResourceGroupLevelIdentityBasedPolicy**: an identity-access policy at the resource group authorization scope, including custom policies and system policies.
	//
	// example:
	//
	// AssumeRolePolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CancelIProductionJobResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
