// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContextDBInfoResponseBodyAccessDeniedDetail) *DescribeContextDBInfoResponseBody
	GetAccessDeniedDetail() *DescribeContextDBInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribeContextDBInfoResponseBodyData) *DescribeContextDBInfoResponseBody
	GetData() *DescribeContextDBInfoResponseBodyData
	SetRequestId(v string) *DescribeContextDBInfoResponseBody
	GetRequestId() *string
}

type DescribeContextDBInfoResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeContextDBInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The instance data.
	Data *DescribeContextDBInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1A586DCB-39A6-4050-81CC-C7BD4CCDB49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContextDBInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBody) GetAccessDeniedDetail() *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContextDBInfoResponseBody) GetData() *DescribeContextDBInfoResponseBodyData {
	return s.Data
}

func (s *DescribeContextDBInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContextDBInfoResponseBody) SetAccessDeniedDetail(v *DescribeContextDBInfoResponseBodyAccessDeniedDetail) *DescribeContextDBInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContextDBInfoResponseBody) SetData(v *DescribeContextDBInfoResponseBodyData) *DescribeContextDBInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContextDBInfoResponseBody) SetRequestId(v string) *DescribeContextDBInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBody) Validate() error {
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

type DescribeContextDBInfoResponseBodyAccessDeniedDetail struct {
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

func (s DescribeContextDBInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContextDBInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContextDBInfoResponseBodyData struct {
	// The instance information.
	Instance *DescribeContextDBInfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
}

func (s DescribeContextDBInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyData) GetInstance() *DescribeContextDBInfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribeContextDBInfoResponseBodyData) SetInstance(v *DescribeContextDBInfoResponseBodyDataInstance) *DescribeContextDBInfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribeContextDBInfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContextDBInfoResponseBodyDataInstance struct {
	// The instance specifications.
	//
	// example:
	//
	// mysql.x2.large.2c
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The list of endpoints.
	ConnAddrs []*DescribeContextDBInfoResponseBodyDataInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
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
	// The VPC instance ID of the replica set node.
	//
	// > This parameter is returned only when the network type of the instance is VPC.
	ReplicaSets []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSets `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Repeated"`
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

func (s DescribeContextDBInfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetConnAddrs() []*DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetReplicaSets() []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	return s.ReplicaSets
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetClassCode(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.ClassCode = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetConnAddrs(v []*DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) *DescribeContextDBInfoResponseBodyDataInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetCreateTime(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetDBInstanceName(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetInstanceId(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetNodeCount(v int32) *DescribeContextDBInfoResponseBodyDataInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetOpenSearchInstanceName(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetRegionId(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetReplicaSets(v []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) *DescribeContextDBInfoResponseBodyDataInstance {
	s.ReplicaSets = v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetStatus(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetStorageType(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetVPCId(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetVSwitchId(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) SetZoneId(v string) *DescribeContextDBInfoResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstance) Validate() error {
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

type DescribeContextDBInfoResponseBodyDataInstanceConnAddrs struct {
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

func (s DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetConnectionString(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetNodeType(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.NodeType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetPort(v int32) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetType(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetVPCId(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetVSwitchId(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeContextDBInfoResponseBodyDataInstanceReplicaSets struct {
	// The instance specifications.
	//
	// example:
	//
	// pg.x2.13large.2c
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The endpoint type.
	ConnAddrs []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
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

func (s DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetConnAddrs() []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetStatus() *string {
	return s.Status
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetClassCode(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.ClassCode = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetConnAddrs(v []*DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.ConnAddrs = v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetCreateTime(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.CreateTime = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetInstanceId(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.InstanceId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetNodeCount(v int32) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.NodeCount = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetNodeType(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.NodeType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetStatus(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.Status = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetStorageType(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.StorageType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) SetZoneId(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets {
	s.ZoneId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSets) Validate() error {
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

type DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs struct {
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
	// - **ReadWrite**: primary instance.
	//
	// - **ReadOnly**: read-only instance.
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

func (s DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetConnectionString(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetNodeType(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.NodeType = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetPort(v int32) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetType(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVPCId(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVSwitchId(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) SetVpcInstanceId(v string) *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeContextDBInfoResponseBodyDataInstanceReplicaSetsConnAddrs) Validate() error {
	return dara.Validate(s)
}
