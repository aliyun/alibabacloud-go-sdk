// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateContextDBResponseBodyAccessDeniedDetail) *CreateContextDBResponseBody
	GetAccessDeniedDetail() *CreateContextDBResponseBodyAccessDeniedDetail
	SetData(v *CreateContextDBResponseBodyData) *CreateContextDBResponseBody
	GetData() *CreateContextDBResponseBodyData
	SetRequestId(v string) *CreateContextDBResponseBody
	GetRequestId() *string
}

type CreateContextDBResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateContextDBResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned result.
	Data *CreateContextDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContextDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDBResponseBody) GetAccessDeniedDetail() *CreateContextDBResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateContextDBResponseBody) GetData() *CreateContextDBResponseBodyData {
	return s.Data
}

func (s *CreateContextDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDBResponseBody) SetAccessDeniedDetail(v *CreateContextDBResponseBodyAccessDeniedDetail) *CreateContextDBResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateContextDBResponseBody) SetData(v *CreateContextDBResponseBodyData) *CreateContextDBResponseBody {
	s.Data = v
	return s
}

func (s *CreateContextDBResponseBody) SetRequestId(v string) *CreateContextDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDBResponseBody) Validate() error {
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

type CreateContextDBResponseBodyAccessDeniedDetail struct {
	// As described above.
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

func (s CreateContextDBResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDBResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateContextDBResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateContextDBResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateContextDBResponseBodyData struct {
	// The administrator key of the context service.
	//
	// example:
	//
	// ctx-admin-***
	ContextDBAdminKey *string `json:"ContextDBAdminKey,omitempty" xml:"ContextDBAdminKey,omitempty"`
	// The name of the context service instance.
	//
	// example:
	//
	// pxt-*********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The instance ID.
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
	// The name of the context service service instance.
	//
	// example:
	//
	// pxt-*********-s
	ServiceReplicaSetName *string `json:"ServiceReplicaSetName,omitempty" xml:"ServiceReplicaSetName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateContextDBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContextDBResponseBodyData) GetContextDBAdminKey() *string {
	return s.ContextDBAdminKey
}

func (s *CreateContextDBResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *CreateContextDBResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateContextDBResponseBodyData) GetDashboardReplicaSetName() *string {
	return s.DashboardReplicaSetName
}

func (s *CreateContextDBResponseBodyData) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *CreateContextDBResponseBodyData) GetServiceReplicaSetName() *string {
	return s.ServiceReplicaSetName
}

func (s *CreateContextDBResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateContextDBResponseBodyData) SetContextDBAdminKey(v string) *CreateContextDBResponseBodyData {
	s.ContextDBAdminKey = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetContextDBInstanceName(v string) *CreateContextDBResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetDBInstanceName(v string) *CreateContextDBResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetDashboardReplicaSetName(v string) *CreateContextDBResponseBodyData {
	s.DashboardReplicaSetName = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetOpenSearchInstanceName(v string) *CreateContextDBResponseBodyData {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetServiceReplicaSetName(v string) *CreateContextDBResponseBodyData {
	s.ServiceReplicaSetName = &v
	return s
}

func (s *CreateContextDBResponseBodyData) SetTaskId(v int32) *CreateContextDBResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateContextDBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
