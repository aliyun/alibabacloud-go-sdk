// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutOpenSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) *ScaleOutOpenSearchResponseBody
	GetAccessDeniedDetail() *ScaleOutOpenSearchResponseBodyAccessDeniedDetail
	SetOrderId(v string) *ScaleOutOpenSearchResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ScaleOutOpenSearchResponseBody
	GetRequestId() *string
}

type ScaleOutOpenSearchResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ScaleOutOpenSearchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The order ID.
	//
	// example:
	//
	// 12345
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleOutOpenSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutOpenSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutOpenSearchResponseBody) GetAccessDeniedDetail() *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ScaleOutOpenSearchResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ScaleOutOpenSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleOutOpenSearchResponseBody) SetAccessDeniedDetail(v *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) *ScaleOutOpenSearchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ScaleOutOpenSearchResponseBody) SetOrderId(v string) *ScaleOutOpenSearchResponseBody {
	s.OrderId = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBody) SetRequestId(v string) *ScaleOutOpenSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleOutOpenSearchResponseBodyAccessDeniedDetail struct {
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
	// The encoded diagnostic message.
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

func (s ScaleOutOpenSearchResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ScaleOutOpenSearchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ScaleOutOpenSearchResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
