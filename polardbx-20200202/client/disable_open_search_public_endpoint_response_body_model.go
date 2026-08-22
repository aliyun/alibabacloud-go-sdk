// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOpenSearchPublicEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) *DisableOpenSearchPublicEndpointResponseBody
	GetAccessDeniedDetail() *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail
	SetMessage(v string) *DisableOpenSearchPublicEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableOpenSearchPublicEndpointResponseBody
	GetRequestId() *string
}

type DisableOpenSearchPublicEndpointResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableOpenSearchPublicEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableOpenSearchPublicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DisableOpenSearchPublicEndpointResponseBody) GetAccessDeniedDetail() *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DisableOpenSearchPublicEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableOpenSearchPublicEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableOpenSearchPublicEndpointResponseBody) SetAccessDeniedDetail(v *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) *DisableOpenSearchPublicEndpointResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBody) SetMessage(v string) *DisableOpenSearchPublicEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBody) SetRequestId(v string) *DisableOpenSearchPublicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The description is the same as above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
