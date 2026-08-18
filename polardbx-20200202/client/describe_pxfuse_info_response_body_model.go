// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfuseInfoResponseBodyAccessDeniedDetail) *DescribePxfuseInfoResponseBody
	GetAccessDeniedDetail() *DescribePxfuseInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfuseInfoResponseBodyData) *DescribePxfuseInfoResponseBody
	GetData() *DescribePxfuseInfoResponseBodyData
	SetRequestId(v string) *DescribePxfuseInfoResponseBody
	GetRequestId() *string
}

type DescribePxfuseInfoResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribePxfuseInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *DescribePxfuseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfuseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponseBody) GetAccessDeniedDetail() *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfuseInfoResponseBody) GetData() *DescribePxfuseInfoResponseBodyData {
	return s.Data
}

func (s *DescribePxfuseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfuseInfoResponseBody) SetAccessDeniedDetail(v *DescribePxfuseInfoResponseBodyAccessDeniedDetail) *DescribePxfuseInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfuseInfoResponseBody) SetData(v *DescribePxfuseInfoResponseBodyData) *DescribePxfuseInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfuseInfoResponseBody) SetRequestId(v string) *DescribePxfuseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBody) Validate() error {
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

type DescribePxfuseInfoResponseBodyAccessDeniedDetail struct {
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
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The authentication principal type.
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
	// The type of missing permission.
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

func (s DescribePxfuseInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfuseInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfuseInfoResponseBodyData struct {
	// The instance information.
	Instance *DescribePxfuseInfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
}

func (s DescribePxfuseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponseBodyData) GetInstance() *DescribePxfuseInfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribePxfuseInfoResponseBodyData) SetInstance(v *DescribePxfuseInfoResponseBodyDataInstance) *DescribePxfuseInfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribePxfuseInfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePxfuseInfoResponseBodyDataInstance struct {
	// The instance specifications.
	//
	// example:
	//
	// mysql.x2.large.2c
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The list of endpoints.
	ConnAddrs []*DescribePxfuseInfoResponseBodyDataInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2026-02-17T02:00:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-spsil01pww4hfz-mem
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// example:
	//
	// local_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-t4n4hf9xey7ea3lp4bwwx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-t4ny14pr37spmjsbv5dc2
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// ap-southeast-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePxfuseInfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetConnAddrs() []*DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetClassCode(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.ClassCode = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetConnAddrs(v []*DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) *DescribePxfuseInfoResponseBodyDataInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetCreateTime(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetInstanceId(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetNodeCount(v int32) *DescribePxfuseInfoResponseBodyDataInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetRegionId(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetStatus(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetStorageType(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.StorageType = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetVPCId(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.VPCId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetVSwitchId(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) SetZoneId(v string) *DescribePxfuseInfoResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePxfuseInfoResponseBodyDataInstanceConnAddrs struct {
	// The endpoint.
	//
	// example:
	//
	// pxc-spsil01pww4hfz.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The endpoint type.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-t4n4hf9xey7ea3lp4bwwx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-t4ny14pr37spmjsbv5dc2
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC-connected instance ID.
	//
	// example:
	//
	// pxc-spsil01pww4hfzjayd-cn-20251013180429
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetConnectionString(v string) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetPort(v int32) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetType(v string) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetVPCId(v string) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetVSwitchId(v string) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) SetVpcInstanceId(v string) *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribePxfuseInfoResponseBodyDataInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}
