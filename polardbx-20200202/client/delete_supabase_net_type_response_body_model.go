// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseNetTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) *DeleteSupabaseNetTypeResponseBody
	GetAccessDeniedDetail() *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail
	SetData(v *DeleteSupabaseNetTypeResponseBodyData) *DeleteSupabaseNetTypeResponseBody
	GetData() *DeleteSupabaseNetTypeResponseBodyData
	SetRequestId(v string) *DeleteSupabaseNetTypeResponseBody
	GetRequestId() *string
}

type DeleteSupabaseNetTypeResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *DeleteSupabaseNetTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSupabaseNetTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseNetTypeResponseBody) GetAccessDeniedDetail() *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteSupabaseNetTypeResponseBody) GetData() *DeleteSupabaseNetTypeResponseBodyData {
	return s.Data
}

func (s *DeleteSupabaseNetTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSupabaseNetTypeResponseBody) SetAccessDeniedDetail(v *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) *DeleteSupabaseNetTypeResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBody) SetData(v *DeleteSupabaseNetTypeResponseBodyData) *DeleteSupabaseNetTypeResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBody) SetRequestId(v string) *DeleteSupabaseNetTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBody) Validate() error {
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

type DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail struct {
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
	// The encoded diagnostic information.
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

func (s DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteSupabaseNetTypeResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteSupabaseNetTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseNetTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseNetTypeResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteSupabaseNetTypeResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteSupabaseNetTypeResponseBodyData) SetDBInstanceName(v string) *DeleteSupabaseNetTypeResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyData) SetTaskId(v int32) *DeleteSupabaseNetTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
