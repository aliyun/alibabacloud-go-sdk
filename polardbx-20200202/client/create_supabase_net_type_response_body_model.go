// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseNetTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) *CreateSupabaseNetTypeResponseBody
	GetAccessDeniedDetail() *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail
	SetData(v *CreateSupabaseNetTypeResponseBodyData) *CreateSupabaseNetTypeResponseBody
	GetData() *CreateSupabaseNetTypeResponseBodyData
	SetRequestId(v string) *CreateSupabaseNetTypeResponseBody
	GetRequestId() *string
}

type CreateSupabaseNetTypeResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *CreateSupabaseNetTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupabaseNetTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupabaseNetTypeResponseBody) GetAccessDeniedDetail() *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateSupabaseNetTypeResponseBody) GetData() *CreateSupabaseNetTypeResponseBodyData {
	return s.Data
}

func (s *CreateSupabaseNetTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSupabaseNetTypeResponseBody) SetAccessDeniedDetail(v *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) *CreateSupabaseNetTypeResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateSupabaseNetTypeResponseBody) SetData(v *CreateSupabaseNetTypeResponseBodyData) *CreateSupabaseNetTypeResponseBody {
	s.Data = v
	return s
}

func (s *CreateSupabaseNetTypeResponseBody) SetRequestId(v string) *CreateSupabaseNetTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBody) Validate() error {
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

type CreateSupabaseNetTypeResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateSupabaseNetTypeResponseBodyData struct {
	// The public endpoint.
	//
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSupabaseNetTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseNetTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSupabaseNetTypeResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateSupabaseNetTypeResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateSupabaseNetTypeResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateSupabaseNetTypeResponseBodyData) SetConnectionString(v string) *CreateSupabaseNetTypeResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyData) SetDBInstanceName(v string) *CreateSupabaseNetTypeResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyData) SetTaskId(v int32) *CreateSupabaseNetTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateSupabaseNetTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
