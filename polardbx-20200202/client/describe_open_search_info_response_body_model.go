// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchInfoResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchInfoResponseBodyData) *DescribeOpenSearchInfoResponseBody
	GetData() *DescribeOpenSearchInfoResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchInfoResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchInfoResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned result set.
	Data *DescribeOpenSearchInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchInfoResponseBody) GetData() *DescribeOpenSearchInfoResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchInfoResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchInfoResponseBody) SetData(v *DescribeOpenSearchInfoResponseBodyData) *DescribeOpenSearchInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchInfoResponseBody) SetRequestId(v string) *DescribeOpenSearchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBody) Validate() error {
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

type DescribeOpenSearchInfoResponseBodyAccessDeniedDetail struct {
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

func (s DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchInfoResponseBodyData struct {
	// The instance information.
	Instance *DescribeOpenSearchInfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The specifications.
	Spec *DescribeOpenSearchInfoResponseBodyDataSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s DescribeOpenSearchInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponseBodyData) GetInstance() *DescribeOpenSearchInfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribeOpenSearchInfoResponseBodyData) GetSpec() *DescribeOpenSearchInfoResponseBodyDataSpec {
	return s.Spec
}

func (s *DescribeOpenSearchInfoResponseBodyData) SetInstance(v *DescribeOpenSearchInfoResponseBodyDataInstance) *DescribeOpenSearchInfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyData) SetSpec(v *DescribeOpenSearchInfoResponseBodyDataSpec) *DescribeOpenSearchInfoResponseBodyData {
	s.Spec = v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchInfoResponseBodyDataInstance struct {
	// The billing method. Valid values:
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// - **PREPAY**: subscription.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The OpenSearch-compatible version.
	//
	// example:
	//
	// 2.0
	CompatibleVersion *string `json:"CompatibleVersion,omitempty" xml:"CompatibleVersion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-02-17T02:00:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deployment mode. Valid values:
	//
	// - multiple: multi-zone deployment.
	//
	// - single: single-zone deployment.
	//
	// example:
	//
	// NORMAL
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The DPI engine version. Default value: 2.0.
	//
	// example:
	//
	// anchashi
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-01-27T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-spsil01pww4hfz-mem
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-bjr76v5uw7z4f5fs-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type of the connection string. Valid values:
	//
	// 	- **Public**: public endpoint.
	//
	// 	- **Private**: private endpoint.
	//
	// 	- **Inner**: private endpoint (classic network).
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
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
	// The time when the last task was updated (in timestamp format).
	//
	// example:
	//
	// 2025-09-02T16:01:51Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the access endpoint resides.
	//
	// example:
	//
	// vpc-2ze99u5upo8zxyf5dlfl5
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID. This parameter is required when you create a DRDS instance of the VPC network type.
	//
	// example:
	//
	// vsw-2zes4ojp6ygziyvq3vhd2
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeOpenSearchInfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetCompatibleVersion() *string {
	return s.CompatibleVersion
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetDeployMode() *string {
	return s.DeployMode
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetNetType() *string {
	return s.NetType
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetChargeType(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetCompatibleVersion(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.CompatibleVersion = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetCreateTime(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetDeployMode(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.DeployMode = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetEngineVersion(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetExpireTime(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetInstanceId(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetInstanceName(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetNetType(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.NetType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetRegionId(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetStatus(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetUpdateTime(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.UpdateTime = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetVpcId(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) SetVswitchId(v string) *DescribeOpenSearchInfoResponseBodyDataInstance {
	s.VswitchId = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchInfoResponseBodyDataSpec struct {
	// The number of coordinator nodes.
	//
	// example:
	//
	// 3
	CoordinatorNodeCount *int32 `json:"CoordinatorNodeCount,omitempty" xml:"CoordinatorNodeCount,omitempty"`
	// The number of CPU cores of a single coordinator node.
	//
	// example:
	//
	// 4
	CoordinatorNodeCpu *int32 `json:"CoordinatorNodeCpu,omitempty" xml:"CoordinatorNodeCpu,omitempty"`
	// Indicates whether coordinator nodes are enabled.
	CoordinatorNodeEnabled *bool `json:"CoordinatorNodeEnabled,omitempty" xml:"CoordinatorNodeEnabled,omitempty"`
	// The memory size of a single coordinator node. Unit: GB.
	//
	// example:
	//
	// 16
	CoordinatorNodeMemoryGB *int32 `json:"CoordinatorNodeMemoryGB,omitempty" xml:"CoordinatorNodeMemoryGB,omitempty"`
	// The number of data nodes.
	//
	// example:
	//
	// 3
	DataNodeCount *int32 `json:"DataNodeCount,omitempty" xml:"DataNodeCount,omitempty"`
	// The number of CPU cores of a single data node.
	//
	// example:
	//
	// 8
	DataNodeCpu *int32 `json:"DataNodeCpu,omitempty" xml:"DataNodeCpu,omitempty"`
	// The memory size of a single data node. Unit: GB.
	//
	// example:
	//
	// 32
	DataNodeMemoryGB *int32 `json:"DataNodeMemoryGB,omitempty" xml:"DataNodeMemoryGB,omitempty"`
	// The master node type. Valid values:
	//
	// - **0**: The master node is a single node.
	//
	// - **2**: The master node is in Cluster Edition.
	//
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// The number of CPU cores of a single dedicated master node.
	//
	// example:
	//
	// 4
	MasterNodeCpu *int32 `json:"MasterNodeCpu,omitempty" xml:"MasterNodeCpu,omitempty"`
	// Indicates whether dedicated master nodes are enabled.
	MasterNodeEnabled *bool `json:"MasterNodeEnabled,omitempty" xml:"MasterNodeEnabled,omitempty"`
	// The memory size of a single dedicated master node. Unit: GB.
	//
	// example:
	//
	// 16
	MasterNodeMemoryGB *int32 `json:"MasterNodeMemoryGB,omitempty" xml:"MasterNodeMemoryGB,omitempty"`
	// The number of replica nodes in the primary zone.
	//
	// > The **ReplicaCount*	- and **SlaveReplicaCount*	- parameters apply only to cloud-native instances. If the instance uses a cluster architecture, these parameters indicate the number of replica nodes of a **single shard*	- in the primary and secondary zones.
	//
	// example:
	//
	// 0
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The storage size of a single data node. Unit: GB.
	//
	// example:
	//
	// 500
	StorageSizeGB *int32 `json:"StorageSizeGB,omitempty" xml:"StorageSizeGB,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeOpenSearchInfoResponseBodyDataSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponseBodyDataSpec) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetCoordinatorNodeCount() *int32 {
	return s.CoordinatorNodeCount
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetCoordinatorNodeCpu() *int32 {
	return s.CoordinatorNodeCpu
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetCoordinatorNodeEnabled() *bool {
	return s.CoordinatorNodeEnabled
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetCoordinatorNodeMemoryGB() *int32 {
	return s.CoordinatorNodeMemoryGB
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetDataNodeCount() *int32 {
	return s.DataNodeCount
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetDataNodeCpu() *int32 {
	return s.DataNodeCpu
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetDataNodeMemoryGB() *int32 {
	return s.DataNodeMemoryGB
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetMasterNodeCount() *int32 {
	return s.MasterNodeCount
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetMasterNodeCpu() *int32 {
	return s.MasterNodeCpu
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetMasterNodeEnabled() *bool {
	return s.MasterNodeEnabled
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetMasterNodeMemoryGB() *int32 {
	return s.MasterNodeMemoryGB
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetStorageSizeGB() *int32 {
	return s.StorageSizeGB
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetCoordinatorNodeCount(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.CoordinatorNodeCount = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetCoordinatorNodeCpu(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.CoordinatorNodeCpu = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetCoordinatorNodeEnabled(v bool) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.CoordinatorNodeEnabled = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetCoordinatorNodeMemoryGB(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.CoordinatorNodeMemoryGB = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetDataNodeCount(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.DataNodeCount = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetDataNodeCpu(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.DataNodeCpu = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetDataNodeMemoryGB(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.DataNodeMemoryGB = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetMasterNodeCount(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetMasterNodeCpu(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.MasterNodeCpu = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetMasterNodeEnabled(v bool) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.MasterNodeEnabled = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetMasterNodeMemoryGB(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.MasterNodeMemoryGB = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetReplicaCount(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.ReplicaCount = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetStorageSizeGB(v int32) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.StorageSizeGB = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) SetStorageType(v string) *DescribeOpenSearchInfoResponseBodyDataSpec {
	s.StorageType = &v
	return s
}

func (s *DescribeOpenSearchInfoResponseBodyDataSpec) Validate() error {
	return dara.Validate(s)
}
