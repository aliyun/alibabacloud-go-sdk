// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) *DescribeSupabaseApiKeyResponseBody
	GetAccessDeniedDetail() *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail
	SetData(v *DescribeSupabaseApiKeyResponseBodyData) *DescribeSupabaseApiKeyResponseBody
	GetData() *DescribeSupabaseApiKeyResponseBodyData
	SetRequestId(v string) *DescribeSupabaseApiKeyResponseBody
	GetRequestId() *string
}

type DescribeSupabaseApiKeyResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The API key information.
	Data *DescribeSupabaseApiKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupabaseApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseApiKeyResponseBody) GetAccessDeniedDetail() *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeSupabaseApiKeyResponseBody) GetData() *DescribeSupabaseApiKeyResponseBodyData {
	return s.Data
}

func (s *DescribeSupabaseApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupabaseApiKeyResponseBody) SetAccessDeniedDetail(v *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) *DescribeSupabaseApiKeyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBody) SetData(v *DescribeSupabaseApiKeyResponseBodyData) *DescribeSupabaseApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBody) SetRequestId(v string) *DescribeSupabaseApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBody) Validate() error {
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

type DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail struct {
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
	// The authentication principal type.
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSupabaseApiKeyResponseBodyData struct {
	// The anonymous access key.
	//
	// example:
	//
	// *********
	AnonKey *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The service role key.
	//
	// example:
	//
	// *********
	ServiceRoleKey *string `json:"ServiceRoleKey,omitempty" xml:"ServiceRoleKey,omitempty"`
}

func (s DescribeSupabaseApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseApiKeyResponseBodyData) GetAnonKey() *string {
	return s.AnonKey
}

func (s *DescribeSupabaseApiKeyResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseApiKeyResponseBodyData) GetServiceRoleKey() *string {
	return s.ServiceRoleKey
}

func (s *DescribeSupabaseApiKeyResponseBodyData) SetAnonKey(v string) *DescribeSupabaseApiKeyResponseBodyData {
	s.AnonKey = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyData) SetDBInstanceName(v string) *DescribeSupabaseApiKeyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyData) SetServiceRoleKey(v string) *DescribeSupabaseApiKeyResponseBodyData {
	s.ServiceRoleKey = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
