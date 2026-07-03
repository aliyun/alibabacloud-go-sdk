// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) *RestartSupabaseInstanceResponseBody
	GetAccessDeniedDetail() *RestartSupabaseInstanceResponseBodyAccessDeniedDetail
	SetData(v *RestartSupabaseInstanceResponseBodyData) *RestartSupabaseInstanceResponseBody
	GetData() *RestartSupabaseInstanceResponseBodyData
	SetRequestId(v string) *RestartSupabaseInstanceResponseBody
	GetRequestId() *string
}

type RestartSupabaseInstanceResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *RestartSupabaseInstanceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *RestartSupabaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartSupabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartSupabaseInstanceResponseBody) GetAccessDeniedDetail() *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RestartSupabaseInstanceResponseBody) GetData() *RestartSupabaseInstanceResponseBodyData {
	return s.Data
}

func (s *RestartSupabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartSupabaseInstanceResponseBody) SetAccessDeniedDetail(v *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) *RestartSupabaseInstanceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RestartSupabaseInstanceResponseBody) SetData(v *RestartSupabaseInstanceResponseBodyData) *RestartSupabaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RestartSupabaseInstanceResponseBody) SetRequestId(v string) *RestartSupabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBody) Validate() error {
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

type RestartSupabaseInstanceResponseBodyAccessDeniedDetail struct {
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
	// The diagnostic information.
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

func (s RestartSupabaseInstanceResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RestartSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RestartSupabaseInstanceResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s RestartSupabaseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartSupabaseInstanceResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartSupabaseInstanceResponseBodyData) SetDBInstanceName(v string) *RestartSupabaseInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *RestartSupabaseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
