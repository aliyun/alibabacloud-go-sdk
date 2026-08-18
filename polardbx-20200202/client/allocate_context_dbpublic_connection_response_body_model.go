// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContextDBPublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) *AllocateContextDBPublicConnectionResponseBody
	GetAccessDeniedDetail() *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *AllocateContextDBPublicConnectionResponseBodyData) *AllocateContextDBPublicConnectionResponseBody
	GetData() *AllocateContextDBPublicConnectionResponseBodyData
	SetRequestId(v string) *AllocateContextDBPublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateContextDBPublicConnectionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The task details.
	Data *AllocateContextDBPublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateContextDBPublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateContextDBPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateContextDBPublicConnectionResponseBody) GetAccessDeniedDetail() *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AllocateContextDBPublicConnectionResponseBody) GetData() *AllocateContextDBPublicConnectionResponseBodyData {
	return s.Data
}

func (s *AllocateContextDBPublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateContextDBPublicConnectionResponseBody) SetAccessDeniedDetail(v *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) *AllocateContextDBPublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBody) SetData(v *AllocateContextDBPublicConnectionResponseBodyData) *AllocateContextDBPublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBody) SetRequestId(v string) *AllocateContextDBPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBody) Validate() error {
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

type AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AllocateContextDBPublicConnectionResponseBodyData struct {
	// The endpoint.
	//
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The name of the context service instance.
	//
	// example:
	//
	// pxt-********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxsp-xxxxxxxxxx
	DBInstanceId *int32 `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The network type of the endpoint.
	//
	// example:
	//
	// 1
	DBInstanceNetType *int32 `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The type of the target node. Valid values: service and dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port of the endpoint.
	//
	// example:
	//
	// 3300
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The backend task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The IP address of the Anti-DDoS Proxy instance protected by the policy.
	//
	// example:
	//
	// https://anchashi.aliyun-inc.coM
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
}

func (s AllocateContextDBPublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AllocateContextDBPublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetDBInstanceNetType() *int32 {
	return s.DBInstanceNetType
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetPort() *string {
	return s.Port
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) GetVip() *string {
	return s.Vip
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetConnectionString(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetContextDBInstanceName(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetDBInstanceId(v int32) *AllocateContextDBPublicConnectionResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetDBInstanceName(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetDBInstanceNetType(v int32) *AllocateContextDBPublicConnectionResponseBodyData {
	s.DBInstanceNetType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetNodeType(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetPort(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.Port = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetTaskId(v int32) *AllocateContextDBPublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) SetVip(v string) *AllocateContextDBPublicConnectionResponseBodyData {
	s.Vip = &v
	return s
}

func (s *AllocateContextDBPublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
