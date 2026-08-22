// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateOpenSearchResponseBodyAccessDeniedDetail) *CreateOpenSearchResponseBody
	GetAccessDeniedDetail() *CreateOpenSearchResponseBodyAccessDeniedDetail
	SetData(v *CreateOpenSearchResponseBodyData) *CreateOpenSearchResponseBody
	GetData() *CreateOpenSearchResponseBodyData
	SetRequestId(v string) *CreateOpenSearchResponseBody
	GetRequestId() *string
}

type CreateOpenSearchResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateOpenSearchResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *CreateOpenSearchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpenSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchResponseBody) GetAccessDeniedDetail() *CreateOpenSearchResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateOpenSearchResponseBody) GetData() *CreateOpenSearchResponseBodyData {
	return s.Data
}

func (s *CreateOpenSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSearchResponseBody) SetAccessDeniedDetail(v *CreateOpenSearchResponseBodyAccessDeniedDetail) *CreateOpenSearchResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateOpenSearchResponseBody) SetData(v *CreateOpenSearchResponseBodyData) *CreateOpenSearchResponseBody {
	s.Data = v
	return s
}

func (s *CreateOpenSearchResponseBody) SetRequestId(v string) *CreateOpenSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSearchResponseBody) Validate() error {
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

type CreateOpenSearchResponseBodyAccessDeniedDetail struct {
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
	// The ID of the Alibaba Cloud account to which the authentication principal belongs.
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
	// NoPermissionType
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

func (s CreateOpenSearchResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateOpenSearchResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateOpenSearchResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateOpenSearchResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 265325896860727
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateOpenSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateOpenSearchResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateOpenSearchResponseBodyData) SetDBInstanceName(v string) *CreateOpenSearchResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateOpenSearchResponseBodyData) SetOrderId(v string) *CreateOpenSearchResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateOpenSearchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
