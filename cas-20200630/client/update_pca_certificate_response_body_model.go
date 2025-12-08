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
	AccessDeniedDetail *UpdatePcaCertificateResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// Id of the request
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
	// example:
	//
	// yundun-cert:XXX
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// RoleSessionName
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// AuthPrincipalOwnerId
	//
	// example:
	//
	// 186XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaEjlETkzRkQ5QjVELTI3NTEtM0I2Ni1BM0E1LThBQUYzMkJBNEJCQg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
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
