// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePcaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdatePcaCertificateResponseBodyAccessDeniedDetail) *UpdatePcaCertificateResponseBody
	GetAccessDeniedDetail() *UpdatePcaCertificateResponseBodyAccessDeniedDetail
	SetRequestId(v string) *UpdatePcaCertificateResponseBody
	GetRequestId() *string
}

type UpdatePcaCertificateResponseBody struct {
	// The error details of the authorization.
	AccessDeniedDetail *UpdatePcaCertificateResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePcaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePcaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePcaCertificateResponseBody) GetAccessDeniedDetail() *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdatePcaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePcaCertificateResponseBody) SetAccessDeniedDetail(v *UpdatePcaCertificateResponseBodyAccessDeniedDetail) *UpdatePcaCertificateResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdatePcaCertificateResponseBody) SetRequestId(v string) *UpdatePcaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePcaCertificateResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePcaCertificateResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that you attempted to perform.
	//
	// example:
	//
	// yundun-cert:XXX
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity that is used for authorization in the request. Valid values:
	//
	// - RAM user: UID of the RAM user
	//
	// - RAM role: RoleName:RoleSessionName
	//
	// - Federated user: ProviderType/ProviderName
	//
	// example:
	//
	// RoleSessionName
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account to which the authorized principal belongs.
	//
	// example:
	//
	// 186XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the identity.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The complete diagnostic information that is encrypted.
	//
	// example:
	//
	// AQEAAAAAaEjlETkzRkQ5QjVELTI3NTEtM0I2Ni1BM0E1LThBQUYzMkJBNEJCQg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason why the authorization failed. Valid values: ExplicitDeny: The authorization is explicitly denied. ImplicitDeny: The authorization is implicitly denied.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s UpdatePcaCertificateResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdatePcaCertificateResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdatePcaCertificateResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdatePcaCertificateResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
