// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) *ModifyOpenSearchClassResponseBody
	GetAccessDeniedDetail() *ModifyOpenSearchClassResponseBodyAccessDeniedDetail
	SetData(v *ModifyOpenSearchClassResponseBodyData) *ModifyOpenSearchClassResponseBody
	GetData() *ModifyOpenSearchClassResponseBodyData
	SetRequestId(v string) *ModifyOpenSearchClassResponseBody
	GetRequestId() *string
}

type ModifyOpenSearchClassResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyOpenSearchClassResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned result.
	Data *ModifyOpenSearchClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOpenSearchClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchClassResponseBody) GetAccessDeniedDetail() *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyOpenSearchClassResponseBody) GetData() *ModifyOpenSearchClassResponseBodyData {
	return s.Data
}

func (s *ModifyOpenSearchClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOpenSearchClassResponseBody) SetAccessDeniedDetail(v *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) *ModifyOpenSearchClassResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyOpenSearchClassResponseBody) SetData(v *ModifyOpenSearchClassResponseBodyData) *ModifyOpenSearchClassResponseBody {
	s.Data = v
	return s
}

func (s *ModifyOpenSearchClassResponseBody) SetRequestId(v string) *ModifyOpenSearchClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOpenSearchClassResponseBodyAccessDeniedDetail struct {
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
	// The identity type used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user.
	//
	// - AssumedRoleUser: RAM role.
	//
	// - Federated: SSO federated identity.
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
	// The type of permission denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyOpenSearchClassResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyOpenSearchClassResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyOpenSearchClassResponseBodyData struct {
	// The order ID.
	//
	// example:
	//
	// 265325896860727
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyOpenSearchClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchClassResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyOpenSearchClassResponseBodyData) SetOrderId(v string) *ModifyOpenSearchClassResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyOpenSearchClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
