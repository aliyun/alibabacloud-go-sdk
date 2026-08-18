// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContextDBPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) *ReleaseContextDBPublicConnectionResponseBody
	GetAccessDeniedDetail() *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *ReleaseContextDBPublicConnectionResponseBodyData) *ReleaseContextDBPublicConnectionResponseBody
	GetData() *ReleaseContextDBPublicConnectionResponseBodyData
	SetRequestId(v string) *ReleaseContextDBPublicConnectionResponseBody
	GetRequestId() *string
}

type ReleaseContextDBPublicConnectionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *ReleaseContextDBPublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseContextDBPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContextDBPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseContextDBPublicConnectionResponseBody) GetAccessDeniedDetail() *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReleaseContextDBPublicConnectionResponseBody) GetData() *ReleaseContextDBPublicConnectionResponseBodyData {
	return s.Data
}

func (s *ReleaseContextDBPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseContextDBPublicConnectionResponseBody) SetAccessDeniedDetail(v *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) *ReleaseContextDBPublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBody) SetData(v *ReleaseContextDBPublicConnectionResponseBodyData) *ReleaseContextDBPublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBody) SetRequestId(v string) *ReleaseContextDBPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBody) Validate() error {
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

type ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The type of the authentication principal.
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
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReleaseContextDBPublicConnectionResponseBodyData struct {
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
	// pxc-xxxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The network type.
	//
	// example:
	//
	// 0
	NetType *int32 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The query node type. Valid values:
	//
	// - service
	//
	// - dashboard
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The database endpoint before the switchover.
	//
	// example:
	//
	// ***.polarxcontextdb.rds.aliyuncs.com
	OldConnectionString *string `json:"OldConnectionString,omitempty" xml:"OldConnectionString,omitempty"`
	// The previous port value.
	//
	// example:
	//
	// 8080
	OldPort *string `json:"OldPort,omitempty" xml:"OldPort,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReleaseContextDBPublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContextDBPublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetNetType() *int32 {
	return s.NetType
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetOldConnectionString() *string {
	return s.OldConnectionString
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetOldPort() *string {
	return s.OldPort
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetContextDBInstanceName(v string) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetDBInstanceId(v int32) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetDBInstanceName(v string) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetNetType(v int32) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.NetType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetNodeType(v string) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetOldConnectionString(v string) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.OldConnectionString = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetOldPort(v string) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.OldPort = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) SetTaskId(v int32) *ReleaseContextDBPublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ReleaseContextDBPublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
