// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseContext0PublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) *ReleaseContext0PublicConnectionResponseBody
	GetAccessDeniedDetail() *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail
	SetData(v *ReleaseContext0PublicConnectionResponseBodyData) *ReleaseContext0PublicConnectionResponseBody
	GetData() *ReleaseContext0PublicConnectionResponseBodyData
	SetRequestId(v string) *ReleaseContext0PublicConnectionResponseBody
	GetRequestId() *string
}

type ReleaseContext0PublicConnectionResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The task details.
	Data *ReleaseContext0PublicConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseContext0PublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContext0PublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseContext0PublicConnectionResponseBody) GetAccessDeniedDetail() *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReleaseContext0PublicConnectionResponseBody) GetData() *ReleaseContext0PublicConnectionResponseBodyData {
	return s.Data
}

func (s *ReleaseContext0PublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseContext0PublicConnectionResponseBody) SetAccessDeniedDetail(v *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) *ReleaseContext0PublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBody) SetData(v *ReleaseContext0PublicConnectionResponseBodyData) *ReleaseContext0PublicConnectionResponseBody {
	s.Data = v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBody) SetRequestId(v string) *ReleaseContext0PublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBody) Validate() error {
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

type ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail struct {
	// The description is as above.
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
	// The description is as above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The description is as above.
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
	// The policy type.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReleaseContext0PublicConnectionResponseBodyData struct {
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
	// The instance name.
	//
	// example:
	//
	// pxc-hzjasd****
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

func (s ReleaseContext0PublicConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReleaseContext0PublicConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetDBInstanceId() *int32 {
	return s.DBInstanceId
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetNetType() *int32 {
	return s.NetType
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetOldConnectionString() *string {
	return s.OldConnectionString
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetOldPort() *string {
	return s.OldPort
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetContext0InstanceName(v string) *ReleaseContext0PublicConnectionResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetDBInstanceId(v int32) *ReleaseContext0PublicConnectionResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetDBInstanceName(v string) *ReleaseContext0PublicConnectionResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetNetType(v int32) *ReleaseContext0PublicConnectionResponseBodyData {
	s.NetType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetNodeType(v string) *ReleaseContext0PublicConnectionResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetOldConnectionString(v string) *ReleaseContext0PublicConnectionResponseBodyData {
	s.OldConnectionString = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetOldPort(v string) *ReleaseContext0PublicConnectionResponseBodyData {
	s.OldPort = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) SetTaskId(v int32) *ReleaseContext0PublicConnectionResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ReleaseContext0PublicConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
