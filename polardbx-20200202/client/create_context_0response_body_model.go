// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContext0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateContext0ResponseBodyAccessDeniedDetail) *CreateContext0ResponseBody
	GetAccessDeniedDetail() *CreateContext0ResponseBodyAccessDeniedDetail
	SetData(v *CreateContext0ResponseBodyData) *CreateContext0ResponseBody
	GetData() *CreateContext0ResponseBodyData
	SetRequestId(v string) *CreateContext0ResponseBody
	GetRequestId() *string
}

type CreateContext0ResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateContext0ResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The instance data.
	Data *CreateContext0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContext0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContext0ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContext0ResponseBody) GetAccessDeniedDetail() *CreateContext0ResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateContext0ResponseBody) GetData() *CreateContext0ResponseBodyData {
	return s.Data
}

func (s *CreateContext0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContext0ResponseBody) SetAccessDeniedDetail(v *CreateContext0ResponseBodyAccessDeniedDetail) *CreateContext0ResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateContext0ResponseBody) SetData(v *CreateContext0ResponseBodyData) *CreateContext0ResponseBody {
	s.Data = v
	return s
}

func (s *CreateContext0ResponseBody) SetRequestId(v string) *CreateContext0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContext0ResponseBody) Validate() error {
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

type CreateContext0ResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateContext0ResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateContext0ResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateContext0ResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateContext0ResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateContext0ResponseBodyData struct {
	// The administrator key of the context service.
	//
	// example:
	//
	// admin-key-example
	Context0AdminKey *string `json:"Context0AdminKey,omitempty" xml:"Context0AdminKey,omitempty"`
	// The name of the context service instance.
	//
	// example:
	//
	// context0-example
	Context0InstanceName *string `json:"Context0InstanceName,omitempty" xml:"Context0InstanceName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the context service dashboard instance.
	//
	// example:
	//
	// pxt-*********-d
	DashboardReplicaSetName *string `json:"DashboardReplicaSetName,omitempty" xml:"DashboardReplicaSetName,omitempty"`
	// The name of the PolarDB-X Search instance.
	//
	// example:
	//
	// pxs-*********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
	// The name of the context service instance.
	//
	// example:
	//
	// pxt-*********-s
	ServiceReplicaSetName *string `json:"ServiceReplicaSetName,omitempty" xml:"ServiceReplicaSetName,omitempty"`
	// The backend task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateContext0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContext0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContext0ResponseBodyData) GetContext0AdminKey() *string {
	return s.Context0AdminKey
}

func (s *CreateContext0ResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *CreateContext0ResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateContext0ResponseBodyData) GetDashboardReplicaSetName() *string {
	return s.DashboardReplicaSetName
}

func (s *CreateContext0ResponseBodyData) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *CreateContext0ResponseBodyData) GetServiceReplicaSetName() *string {
	return s.ServiceReplicaSetName
}

func (s *CreateContext0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateContext0ResponseBodyData) SetContext0AdminKey(v string) *CreateContext0ResponseBodyData {
	s.Context0AdminKey = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetContext0InstanceName(v string) *CreateContext0ResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetDBInstanceName(v string) *CreateContext0ResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetDashboardReplicaSetName(v string) *CreateContext0ResponseBodyData {
	s.DashboardReplicaSetName = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetOpenSearchInstanceName(v string) *CreateContext0ResponseBodyData {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetServiceReplicaSetName(v string) *CreateContext0ResponseBodyData {
	s.ServiceReplicaSetName = &v
	return s
}

func (s *CreateContext0ResponseBodyData) SetTaskId(v int32) *CreateContext0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateContext0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
