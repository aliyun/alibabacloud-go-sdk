// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteContextDBResponseBodyAccessDeniedDetail) *DeleteContextDBResponseBody
	GetAccessDeniedDetail() *DeleteContextDBResponseBodyAccessDeniedDetail
	SetData(v *DeleteContextDBResponseBodyData) *DeleteContextDBResponseBody
	GetData() *DeleteContextDBResponseBodyData
	SetRequestId(v string) *DeleteContextDBResponseBody
	GetRequestId() *string
}

type DeleteContextDBResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteContextDBResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The parameter details.
	Data *DeleteContextDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContextDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextDBResponseBody) GetAccessDeniedDetail() *DeleteContextDBResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteContextDBResponseBody) GetData() *DeleteContextDBResponseBodyData {
	return s.Data
}

func (s *DeleteContextDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextDBResponseBody) SetAccessDeniedDetail(v *DeleteContextDBResponseBodyAccessDeniedDetail) *DeleteContextDBResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteContextDBResponseBody) SetData(v *DeleteContextDBResponseBodyData) *DeleteContextDBResponseBody {
	s.Data = v
	return s
}

func (s *DeleteContextDBResponseBody) SetRequestId(v string) *DeleteContextDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextDBResponseBody) Validate() error {
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

type DeleteContextDBResponseBodyAccessDeniedDetail struct {
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

func (s DeleteContextDBResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDBResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteContextDBResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteContextDBResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteContextDBResponseBodyData struct {
	// The context service instance name.
	//
	// example:
	//
	// pxt-*********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The context service dashboard instance name.
	//
	// example:
	//
	// pxt-*********-d
	DashboardReplicaSetName *string `json:"DashboardReplicaSetName,omitempty" xml:"DashboardReplicaSetName,omitempty"`
	// The context service service instance name.
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

func (s DeleteContextDBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteContextDBResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *DeleteContextDBResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteContextDBResponseBodyData) GetDashboardReplicaSetName() *string {
	return s.DashboardReplicaSetName
}

func (s *DeleteContextDBResponseBodyData) GetServiceReplicaSetName() *string {
	return s.ServiceReplicaSetName
}

func (s *DeleteContextDBResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteContextDBResponseBodyData) SetContextDBInstanceName(v string) *DeleteContextDBResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *DeleteContextDBResponseBodyData) SetDBInstanceName(v string) *DeleteContextDBResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteContextDBResponseBodyData) SetDashboardReplicaSetName(v string) *DeleteContextDBResponseBodyData {
	s.DashboardReplicaSetName = &v
	return s
}

func (s *DeleteContextDBResponseBodyData) SetServiceReplicaSetName(v string) *DeleteContextDBResponseBodyData {
	s.ServiceReplicaSetName = &v
	return s
}

func (s *DeleteContextDBResponseBodyData) SetTaskId(v int32) *DeleteContextDBResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteContextDBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
