// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContext0PublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) *AllocateContext0PublicConnectionResponseBody
	GetAccessDeniedDetail() *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *AllocateContext0PublicConnectionResponseBodyData) *AllocateContext0PublicConnectionResponseBody
	GetData() *AllocateContext0PublicConnectionResponseBodyData
	SetRequestId(v string) *AllocateContext0PublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateContext0PublicConnectionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The task details.
	Data *AllocateContext0PublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateContext0PublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateContext0PublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateContext0PublicConnectionResponseBody) GetAccessDeniedDetail() *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AllocateContext0PublicConnectionResponseBody) GetData() *AllocateContext0PublicConnectionResponseBodyData {
	return s.Data
}

func (s *AllocateContext0PublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateContext0PublicConnectionResponseBody) SetAccessDeniedDetail(v *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) *AllocateContext0PublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBody) SetData(v *AllocateContext0PublicConnectionResponseBodyData) *AllocateContext0PublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBody) SetRequestId(v string) *AllocateContext0PublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBody) Validate() error {
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

type AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail struct {
	// The description is the same as above.
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
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The description is the same as above.
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

func (s AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AllocateContext0PublicConnectionResponseBodyData struct {
	// The endpoint.
	//
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The context service instance name.
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
	// The instance ID.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The network type of the endpoint.
	//
	// example:
	//
	// 1
	DBInstanceNetType *int32 `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The target node type: service or dashboard.
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
	// The task ID.
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

func (s AllocateContext0PublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AllocateContext0PublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetDBInstanceNetType() *int32 {
	return s.DBInstanceNetType
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetPort() *string {
	return s.Port
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *AllocateContext0PublicConnectionResponseBodyData) GetVip() *string {
	return s.Vip
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetConnectionString(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetContext0InstanceName(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetDBInstanceId(v int32) *AllocateContext0PublicConnectionResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetDBInstanceName(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetDBInstanceNetType(v int32) *AllocateContext0PublicConnectionResponseBodyData {
	s.DBInstanceNetType = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetNodeType(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetPort(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.Port = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetTaskId(v int32) *AllocateContext0PublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) SetVip(v string) *AllocateContext0PublicConnectionResponseBodyData {
	s.Vip = &v
	return s
}

func (s *AllocateContext0PublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
