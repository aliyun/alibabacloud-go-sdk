// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContext0SecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) *ModifyContext0SecurityIpsResponseBody
	GetAccessDeniedDetail() *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *ModifyContext0SecurityIpsResponseBodyData) *ModifyContext0SecurityIpsResponseBody
	GetData() *ModifyContext0SecurityIpsResponseBodyData
	SetRequestId(v string) *ModifyContext0SecurityIpsResponseBody
	GetRequestId() *string
}

type ModifyContext0SecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data.
	Data *ModifyContext0SecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContext0SecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContext0SecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContext0SecurityIpsResponseBody) GetAccessDeniedDetail() *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyContext0SecurityIpsResponseBody) GetData() *ModifyContext0SecurityIpsResponseBodyData {
	return s.Data
}

func (s *ModifyContext0SecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContext0SecurityIpsResponseBody) SetAccessDeniedDetail(v *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) *ModifyContext0SecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBody) SetData(v *ModifyContext0SecurityIpsResponseBodyData) *ModifyContext0SecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBody) SetRequestId(v string) *ModifyContext0SecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBody) Validate() error {
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

type ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail struct {
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
	// The diagnostic information.
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
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyContext0SecurityIpsResponseBodyData struct {
	// The name of the context service instance.
	//
	// example:
	//
	// context0-example
	Context0InstanceName *string `json:"Context0InstanceName,omitempty" xml:"Context0InstanceName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxsp-xxxxxxxxxx
	DBInstanceId *int32 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The list of task IDs in the task group.
	TaskIds []*int32 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s ModifyContext0SecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyContext0SecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyContext0SecurityIpsResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *ModifyContext0SecurityIpsResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *ModifyContext0SecurityIpsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyContext0SecurityIpsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyContext0SecurityIpsResponseBodyData) GetTaskIds() []*int32 {
	return s.TaskIds
}

func (s *ModifyContext0SecurityIpsResponseBodyData) SetContext0InstanceName(v string) *ModifyContext0SecurityIpsResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyData) SetDBInstanceId(v int32) *ModifyContext0SecurityIpsResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyData) SetDBInstanceName(v string) *ModifyContext0SecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyData) SetTaskId(v int32) *ModifyContext0SecurityIpsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyData) SetTaskIds(v []*int32) *ModifyContext0SecurityIpsResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *ModifyContext0SecurityIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
