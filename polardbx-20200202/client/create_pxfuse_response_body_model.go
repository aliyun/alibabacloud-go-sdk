// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfuseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePxfuseResponseBodyAccessDeniedDetail) *CreatePxfuseResponseBody
	GetAccessDeniedDetail() *CreatePxfuseResponseBodyAccessDeniedDetail
	SetData(v *CreatePxfuseResponseBodyData) *CreatePxfuseResponseBody
	GetData() *CreatePxfuseResponseBodyData
	SetRequestId(v string) *CreatePxfuseResponseBody
	GetRequestId() *string
}

type CreatePxfuseResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePxfuseResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The policy details.
	Data *CreatePxfuseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePxfuseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfuseResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePxfuseResponseBody) GetAccessDeniedDetail() *CreatePxfuseResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePxfuseResponseBody) GetData() *CreatePxfuseResponseBodyData {
	return s.Data
}

func (s *CreatePxfuseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePxfuseResponseBody) SetAccessDeniedDetail(v *CreatePxfuseResponseBodyAccessDeniedDetail) *CreatePxfuseResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePxfuseResponseBody) SetData(v *CreatePxfuseResponseBodyData) *CreatePxfuseResponseBody {
	s.Data = v
	return s
}

func (s *CreatePxfuseResponseBody) SetRequestId(v string) *CreatePxfuseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePxfuseResponseBody) Validate() error {
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

type CreatePxfuseResponseBodyAccessDeniedDetail struct {
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

func (s CreatePxfuseResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfuseResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePxfuseResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePxfuseResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePxfuseResponseBodyData struct {
	// Indicates whether the instance already exists.
	//
	// example:
	//
	// False
	AlreadyExists *bool `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the Agent observability service instance.
	//
	// example:
	//
	// pxc-*********-pxf
	PxfuseInstanceName *string `json:"PxfuseInstanceName,omitempty" xml:"PxfuseInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreatePxfuseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfuseResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePxfuseResponseBodyData) GetAlreadyExists() *bool {
	return s.AlreadyExists
}

func (s *CreatePxfuseResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfuseResponseBodyData) GetPxfuseInstanceName() *string {
	return s.PxfuseInstanceName
}

func (s *CreatePxfuseResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreatePxfuseResponseBodyData) SetAlreadyExists(v bool) *CreatePxfuseResponseBodyData {
	s.AlreadyExists = &v
	return s
}

func (s *CreatePxfuseResponseBodyData) SetDBInstanceName(v string) *CreatePxfuseResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfuseResponseBodyData) SetPxfuseInstanceName(v string) *CreatePxfuseResponseBodyData {
	s.PxfuseInstanceName = &v
	return s
}

func (s *CreatePxfuseResponseBodyData) SetTaskId(v int32) *CreatePxfuseResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreatePxfuseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
