// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContextDBSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) *ModifyContextDBSecurityIpsResponseBody
	GetAccessDeniedDetail() *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *ModifyContextDBSecurityIpsResponseBodyData) *ModifyContextDBSecurityIpsResponseBody
	GetData() *ModifyContextDBSecurityIpsResponseBodyData
	SetRequestId(v string) *ModifyContextDBSecurityIpsResponseBody
	GetRequestId() *string
}

type ModifyContextDBSecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *ModifyContextDBSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContextDBSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContextDBSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContextDBSecurityIpsResponseBody) GetAccessDeniedDetail() *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyContextDBSecurityIpsResponseBody) GetData() *ModifyContextDBSecurityIpsResponseBodyData {
	return s.Data
}

func (s *ModifyContextDBSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContextDBSecurityIpsResponseBody) SetAccessDeniedDetail(v *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) *ModifyContextDBSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBody) SetData(v *ModifyContextDBSecurityIpsResponseBodyData) *ModifyContextDBSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBody) SetRequestId(v string) *ModifyContextDBSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBody) Validate() error {
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

type ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail struct {
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

func (s ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyContextDBSecurityIpsResponseBodyData struct {
	// The context service instance name.
	//
	// example:
	//
	// pxt-*********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxxxx
	DBInstanceId *int32 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-supabase-001
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

func (s ModifyContextDBSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyContextDBSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) GetTaskIds() []*int32 {
	return s.TaskIds
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) SetContextDBInstanceName(v string) *ModifyContextDBSecurityIpsResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) SetDBInstanceId(v int32) *ModifyContextDBSecurityIpsResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) SetDBInstanceName(v string) *ModifyContextDBSecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) SetTaskId(v int32) *ModifyContextDBSecurityIpsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) SetTaskIds(v []*int32) *ModifyContextDBSecurityIpsResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *ModifyContextDBSecurityIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
