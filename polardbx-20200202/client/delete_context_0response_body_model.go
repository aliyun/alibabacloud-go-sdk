// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContext0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteContext0ResponseBodyAccessDeniedDetail) *DeleteContext0ResponseBody
	GetAccessDeniedDetail() *DeleteContext0ResponseBodyAccessDeniedDetail
	SetData(v *DeleteContext0ResponseBodyData) *DeleteContext0ResponseBody
	GetData() *DeleteContext0ResponseBodyData
	SetRequestId(v string) *DeleteContext0ResponseBody
	GetRequestId() *string
}

type DeleteContext0ResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteContext0ResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DeleteContext0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6352AC16-76BF-5135-B1EA-ED49293526E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContext0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContext0ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContext0ResponseBody) GetAccessDeniedDetail() *DeleteContext0ResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteContext0ResponseBody) GetData() *DeleteContext0ResponseBodyData {
	return s.Data
}

func (s *DeleteContext0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContext0ResponseBody) SetAccessDeniedDetail(v *DeleteContext0ResponseBodyAccessDeniedDetail) *DeleteContext0ResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteContext0ResponseBody) SetData(v *DeleteContext0ResponseBodyData) *DeleteContext0ResponseBody {
	s.Data = v
	return s
}

func (s *DeleteContext0ResponseBody) SetRequestId(v string) *DeleteContext0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContext0ResponseBody) Validate() error {
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

type DeleteContext0ResponseBodyAccessDeniedDetail struct {
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
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteContext0ResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteContext0ResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteContext0ResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteContext0ResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteContext0ResponseBodyData struct {
	// The name of the context service instance.
	//
	// example:
	//
	// context0-example
	Context0InstanceName *string `json:"Context0InstanceName,omitempty" xml:"Context0InstanceName,omitempty"`
	// The name of the instance.
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
	// The name of the context service instance.
	//
	// example:
	//
	// pxt-*********-s
	ServiceReplicaSetName *string `json:"ServiceReplicaSetName,omitempty" xml:"ServiceReplicaSetName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteContext0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteContext0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteContext0ResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *DeleteContext0ResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteContext0ResponseBodyData) GetDashboardReplicaSetName() *string {
	return s.DashboardReplicaSetName
}

func (s *DeleteContext0ResponseBodyData) GetServiceReplicaSetName() *string {
	return s.ServiceReplicaSetName
}

func (s *DeleteContext0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteContext0ResponseBodyData) SetContext0InstanceName(v string) *DeleteContext0ResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *DeleteContext0ResponseBodyData) SetDBInstanceName(v string) *DeleteContext0ResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteContext0ResponseBodyData) SetDashboardReplicaSetName(v string) *DeleteContext0ResponseBodyData {
	s.DashboardReplicaSetName = &v
	return s
}

func (s *DeleteContext0ResponseBodyData) SetServiceReplicaSetName(v string) *DeleteContext0ResponseBodyData {
	s.ServiceReplicaSetName = &v
	return s
}

func (s *DeleteContext0ResponseBodyData) SetTaskId(v int32) *DeleteContext0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteContext0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
