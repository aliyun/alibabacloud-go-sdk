// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePxfsResponseBodyAccessDeniedDetail) *CreatePxfsResponseBody
	GetAccessDeniedDetail() *CreatePxfsResponseBodyAccessDeniedDetail
	SetData(v *CreatePxfsResponseBodyData) *CreatePxfsResponseBody
	GetData() *CreatePxfsResponseBodyData
	SetRequestId(v string) *CreatePxfsResponseBody
	GetRequestId() *string
}

type CreatePxfsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePxfsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *CreatePxfsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePxfsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePxfsResponseBody) GetAccessDeniedDetail() *CreatePxfsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePxfsResponseBody) GetData() *CreatePxfsResponseBodyData {
	return s.Data
}

func (s *CreatePxfsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePxfsResponseBody) SetAccessDeniedDetail(v *CreatePxfsResponseBodyAccessDeniedDetail) *CreatePxfsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePxfsResponseBody) SetData(v *CreatePxfsResponseBodyData) *CreatePxfsResponseBody {
	s.Data = v
	return s
}

func (s *CreatePxfsResponseBody) SetRequestId(v string) *CreatePxfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePxfsResponseBody) Validate() error {
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

type CreatePxfsResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the security principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the security principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the security principal.
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

func (s CreatePxfsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePxfsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePxfsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePxfsResponseBodyData struct {
	// Indicates whether the instance already exists.
	//
	// example:
	//
	// false
	AlreadyExists *bool `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	// The PolarDB-X instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The Agent file service instance ID.
	//
	// example:
	//
	// pxfs-xxxxxxxx
	PxfsInstanceId *string `json:"PxfsInstanceId,omitempty" xml:"PxfsInstanceId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 123456789
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreatePxfsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePxfsResponseBodyData) GetAlreadyExists() *bool {
	return s.AlreadyExists
}

func (s *CreatePxfsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfsResponseBodyData) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *CreatePxfsResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreatePxfsResponseBodyData) SetAlreadyExists(v bool) *CreatePxfsResponseBodyData {
	s.AlreadyExists = &v
	return s
}

func (s *CreatePxfsResponseBodyData) SetDBInstanceName(v string) *CreatePxfsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfsResponseBodyData) SetPxfsInstanceId(v string) *CreatePxfsResponseBodyData {
	s.PxfsInstanceId = &v
	return s
}

func (s *CreatePxfsResponseBodyData) SetTaskId(v int64) *CreatePxfsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreatePxfsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
