// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInOpenSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ScaleInOpenSearchResponseBodyAccessDeniedDetail) *ScaleInOpenSearchResponseBody
	GetAccessDeniedDetail() *ScaleInOpenSearchResponseBodyAccessDeniedDetail
	SetOrderId(v string) *ScaleInOpenSearchResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ScaleInOpenSearchResponseBody
	GetRequestId() *string
}

type ScaleInOpenSearchResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ScaleInOpenSearchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The order ID.
	//
	// example:
	//
	// 20211103105558
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleInOpenSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleInOpenSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleInOpenSearchResponseBody) GetAccessDeniedDetail() *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ScaleInOpenSearchResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ScaleInOpenSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleInOpenSearchResponseBody) SetAccessDeniedDetail(v *ScaleInOpenSearchResponseBodyAccessDeniedDetail) *ScaleInOpenSearchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ScaleInOpenSearchResponseBody) SetOrderId(v string) *ScaleInOpenSearchResponseBody {
	s.OrderId = &v
	return s
}

func (s *ScaleInOpenSearchResponseBody) SetRequestId(v string) *ScaleInOpenSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleInOpenSearchResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleInOpenSearchResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
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
	// The encrypted complete diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the permission denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ScaleInOpenSearchResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ScaleInOpenSearchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ScaleInOpenSearchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ScaleInOpenSearchResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
