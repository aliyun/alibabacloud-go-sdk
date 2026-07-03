// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolardbxSupabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) *DeletePolardbxSupabaseInstanceResponseBody
	GetAccessDeniedDetail() *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail
	SetData(v *DeletePolardbxSupabaseInstanceResponseBodyData) *DeletePolardbxSupabaseInstanceResponseBody
	GetData() *DeletePolardbxSupabaseInstanceResponseBodyData
	SetRequestId(v string) *DeletePolardbxSupabaseInstanceResponseBody
	GetRequestId() *string
}

type DeletePolardbxSupabaseInstanceResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The deletion result.
	Data *DeletePolardbxSupabaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolardbxSupabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolardbxSupabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) GetAccessDeniedDetail() *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) GetData() *DeletePolardbxSupabaseInstanceResponseBodyData {
	return s.Data
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) SetAccessDeniedDetail(v *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) *DeletePolardbxSupabaseInstanceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) SetData(v *DeletePolardbxSupabaseInstanceResponseBodyData) *DeletePolardbxSupabaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) SetRequestId(v string) *DeletePolardbxSupabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBody) Validate() error {
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

type DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail struct {
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

func (s DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeletePolardbxSupabaseInstanceResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxsp-xxxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeletePolardbxSupabaseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeletePolardbxSupabaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyData) SetDBInstanceName(v string) *DeletePolardbxSupabaseInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyData) SetTaskId(v int64) *DeletePolardbxSupabaseInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
