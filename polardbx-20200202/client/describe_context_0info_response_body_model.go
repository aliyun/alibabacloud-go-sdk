// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0InfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContext0InfoResponseBodyAccessDeniedDetail) *DescribeContext0InfoResponseBody
	GetAccessDeniedDetail() *DescribeContext0InfoResponseBodyAccessDeniedDetail
	SetData(v *DescribeContext0InfoResponseBodyData) *DescribeContext0InfoResponseBody
	GetData() *DescribeContext0InfoResponseBodyData
	SetRequestId(v string) *DescribeContext0InfoResponseBody
	GetRequestId() *string
}

type DescribeContext0InfoResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DescribeContext0InfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The parameter details.
	Data *DescribeContext0InfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContext0InfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBody) GetAccessDeniedDetail() *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContext0InfoResponseBody) GetData() *DescribeContext0InfoResponseBodyData {
	return s.Data
}

func (s *DescribeContext0InfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContext0InfoResponseBody) SetAccessDeniedDetail(v *DescribeContext0InfoResponseBodyAccessDeniedDetail) *DescribeContext0InfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContext0InfoResponseBody) SetData(v *DescribeContext0InfoResponseBodyData) *DescribeContext0InfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContext0InfoResponseBody) SetRequestId(v string) *DescribeContext0InfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContext0InfoResponseBody) Validate() error {
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

type DescribeContext0InfoResponseBodyAccessDeniedDetail struct {
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
	// The authentication principal type.
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

func (s DescribeContext0InfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContext0InfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContext0InfoResponseBodyData struct {
	// The instance information.
	Instance *DescribeContext0InfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
}

func (s DescribeContext0InfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyData) GetInstance() *DescribeContext0InfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribeContext0InfoResponseBodyData) SetInstance(v *DescribeContext0InfoResponseBodyDataInstance) *DescribeContext0InfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribeContext0InfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContext0InfoResponseBodyDataInstance struct {
	// The instance specifications.
	//
	// example:
	//
	// mysql.x2.large.2c
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The list of endpoints.
	ConnAddrs []*DescribeContext0InfoResponseBodyDataInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2026-02-17T02:00:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database instance name.
	//
	// example:
	//
	// pxc-hzr9qzafkeury3
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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
	// The PolarDB-X Search instance name.
	//
	// example:
	//
	// pxs-********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoints of the nodes.
	ReplicaSets []*DescribeContext0InfoResponseBodyDataInstanceReplicaSets `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Repeated"`
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

func (s DescribeContext0InfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetConnAddrs() []*DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetReplicaSets() []*DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	return s.ReplicaSets
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContext0InfoResponseBodyDataInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetClassCode(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.ClassCode = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetConnAddrs(v []*DescribeContext0InfoResponseBodyDataInstanceConnAddrs) *DescribeContext0InfoResponseBodyDataInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetCreateTime(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetDBInstanceName(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetInstanceId(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetNodeCount(v int32) *DescribeContext0InfoResponseBodyDataInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetOpenSearchInstanceName(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetRegionId(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetReplicaSets(v []*DescribeContext0InfoResponseBodyDataInstanceReplicaSets) *DescribeContext0InfoResponseBodyDataInstance {
	s.ReplicaSets = v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetStatus(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetStorageType(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetVPCId(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetVSwitchId(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) SetZoneId(v string) *DescribeContext0InfoResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReplicaSets != nil {
		for _, item := range s.ReplicaSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeContext0InfoResponseBodyDataInstanceConnAddrs struct {
	// The endpoint.
	//
	// example:
	//
	// pxc-spsil01pww4hfz.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The target node type: service or dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
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

func (s DescribeContext0InfoResponseBodyDataInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetConnectionString(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetNodeType(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.NodeType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetPort(v int32) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetType(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetVPCId(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetVSwitchId(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeContext0InfoResponseBodyDataInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeContext0InfoResponseBodyDataInstanceReplicaSets struct {
	// The instance specifications.
	//
	// example:
	//
	// pg.x2.13large.2c
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The vSwitch ID.
	ConnAddrs []*DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2026-07-28T02:01:13Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-hzrbqgiocrpu8a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The target node type: service or dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The node status. Valid values:
	//
	// - **0**: Running.
	//
	// - **1**: Creating.
	//
	// - **2**: Abnormal.
	//
	// - **3**: Expired.
	//
	// - **4**: Releasing.
	//
	// - **5**: Released.
	//
	// - **6**: Locked.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContext0InfoResponseBodyDataInstanceReplicaSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetConnAddrs() []*DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetStatus() *string {
	return s.Status
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetClassCode(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.ClassCode = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetConnAddrs(v []*DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.ConnAddrs = v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetCreateTime(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.CreateTime = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetInstanceId(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.InstanceId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetNodeCount(v int32) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.NodeCount = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetNodeType(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.NodeType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetStatus(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.Status = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetStorageType(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.StorageType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) SetZoneId(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSets {
	s.ZoneId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSets) Validate() error {
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

type DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs struct {
	// The endpoint.
	//
	// example:
	//
	// pxc-hzrlz8e3khuaoz.polarx.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The target node type: service or dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The instance type. Valid values:
	//
	// - **ReadWrite**: Primary instance.
	//
	// - **ReadOnly**: Read-only instance.
	//
	// example:
	//
	// RemoveHeader
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1550umsomy2mw24vhwl
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zedp17pfss1133bvdizl
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC-connected instance ID.
	//
	// example:
	//
	// vpc-8vbdw66evguopfcfvieoi
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetConnectionString(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetNodeType(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.NodeType = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetPort(v int32) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetType(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVPCId(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVSwitchId(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVpcInstanceId(v string) *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeContext0InfoResponseBodyDataInstanceReplicaSetsConnAddrs) Validate() error {
	return dara.Validate(s)
}
