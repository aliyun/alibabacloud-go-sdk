// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodepool(v *CreateClusterNodePoolResponseBodyNodepool) *CreateClusterNodePoolResponseBody
	GetNodepool() *CreateClusterNodePoolResponseBodyNodepool
	SetRequestId(v string) *CreateClusterNodePoolResponseBody
	GetRequestId() *string
}

type CreateClusterNodePoolResponseBody struct {
	// The node pool configurations.
	Nodepool *CreateClusterNodePoolResponseBodyNodepool `json:"Nodepool,omitempty" xml:"Nodepool,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBody) GetNodepool() *CreateClusterNodePoolResponseBodyNodepool {
	return s.Nodepool
}

func (s *CreateClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterNodePoolResponseBody) SetNodepool(v *CreateClusterNodePoolResponseBodyNodepool) *CreateClusterNodePoolResponseBody {
	s.Nodepool = v
	return s
}

func (s *CreateClusterNodePoolResponseBody) SetRequestId(v string) *CreateClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBody) Validate() error {
	if s.Nodepool != nil {
		if err := s.Nodepool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolResponseBodyNodepool struct {
	// The instance ID.
	//
	// example:
	//
	// eck-d2666c5f
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The Kubernetes-related configurations.
	KubernetesConfig *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// The node pool information.
	NodepoolInfo *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty" type:"Struct"`
	// The ID of the resource group to which the newly created scaling group belongs. You can assign the newly created scaling group to a specified resource group. If you do not specify this parameter, the scaling group is assigned to the default resource group.
	ScalingGroup *CreateClusterNodePoolResponseBodyNodepoolScalingGroup `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
	// The status information of the node pool.
	Status *CreateClusterNodePoolResponseBodyNodepoolStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolResponseBodyNodepool) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepool) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepool) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterNodePoolResponseBodyNodepool) GetKubernetesConfig() *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig {
	return s.KubernetesConfig
}

func (s *CreateClusterNodePoolResponseBodyNodepool) GetNodepoolInfo() *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo {
	return s.NodepoolInfo
}

func (s *CreateClusterNodePoolResponseBodyNodepool) GetScalingGroup() *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	return s.ScalingGroup
}

func (s *CreateClusterNodePoolResponseBodyNodepool) GetStatus() *CreateClusterNodePoolResponseBodyNodepoolStatus {
	return s.Status
}

func (s *CreateClusterNodePoolResponseBodyNodepool) SetClusterId(v string) *CreateClusterNodePoolResponseBodyNodepool {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepool) SetKubernetesConfig(v *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) *CreateClusterNodePoolResponseBodyNodepool {
	s.KubernetesConfig = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepool) SetNodepoolInfo(v *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) *CreateClusterNodePoolResponseBodyNodepool {
	s.NodepoolInfo = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepool) SetScalingGroup(v *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) *CreateClusterNodePoolResponseBodyNodepool {
	s.ScalingGroup = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepool) SetStatus(v *CreateClusterNodePoolResponseBodyNodepoolStatus) *CreateClusterNodePoolResponseBodyNodepool {
	s.Status = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepool) Validate() error {
	if s.KubernetesConfig != nil {
		if err := s.KubernetesConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodepoolInfo != nil {
		if err := s.NodepoolInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingGroup != nil {
		if err := s.ScalingGroup.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig struct {
	// The script to be executed after the instance is created and before it is added to the cluster. The content must be Base64-encoded.
	//
	// example:
	//
	// ZWNobyBvawo=
	PreUserData *string `json:"PreUserData,omitempty" xml:"PreUserData,omitempty"`
	// The script to be executed after the instance is created and added to the cluster. The content must be Base64-encoded.
	//
	// example:
	//
	// ZWNobyBvawo=
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) SetPreUserData(v string) *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) SetUserData(v string) *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo struct {
	// The ENS region ID of the cluster.
	//
	// example:
	//
	// cn-guiyang-14
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np1176939548ac49e59febe82dcbf1ad56
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) SetEnsRegionId(v string) *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo {
	s.EnsRegionId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) SetName(v string) *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo {
	s.Name = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) SetNodepoolId(v string) *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo {
	s.NodepoolId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolResponseBodyNodepoolScalingGroup struct {
	// The data cloud disk information of the nodes.
	DataDisks []*CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The image ID of the control plane virtual machine.
	//
	// example:
	//
	// m-67bdk1kpu1ylmt7k33h5dbiov
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method of the instance.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The list of instance types for the nodes in the node pool.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The name of the key pair.
	//
	// example:
	//
	// key-pair-name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The logon password of the instance.
	//
	// example:
	//
	// **********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The system cloud disk category of the node.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The size of the system cloud disk. Unit: GB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The IDs of the vSwitches where the cluster nodes reside.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s CreateClusterNodePoolResponseBodyNodepoolScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetDataDisks() []*CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks {
	return s.DataDisks
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetPassword() *string {
	return s.Password
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetDataDisks(v []*CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.DataDisks = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetImageId(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.ImageId = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetInstanceChargeType(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetInstanceTypes(v []*string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetKeyPairName(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetPassword(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.Password = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetSystemDiskCategory(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetSystemDiskSize(v int32) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) SetVswitchIds(v []*string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroup) Validate() error {
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks struct {
	// The disk category.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The size of the data cloud disk. Unit: GB.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) SetCategory(v string) *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks {
	s.Category = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) SetSize(v int32) *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks {
	s.Size = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolResponseBodyNodepoolStatus struct {
	// The expected number of instances.
	//
	// example:
	//
	// 10
	DesiredNodes *int32 `json:"DesiredNodes,omitempty" xml:"DesiredNodes,omitempty"`
	// The nodes that failed to be created in the node pool.
	//
	// example:
	//
	// 0
	FailedNodes *int32 `json:"FailedNodes,omitempty" xml:"FailedNodes,omitempty"`
	// The number of instances that are being created.
	//
	// example:
	//
	// 0
	InitialNodes *int32 `json:"InitialNodes,omitempty" xml:"InitialNodes,omitempty"`
	// The number of instances that are being deleted.
	//
	// example:
	//
	// 0
	RemovingNodes *int32 `json:"RemovingNodes,omitempty" xml:"RemovingNodes,omitempty"`
	// The number of instances that are in use.
	//
	// example:
	//
	// 10
	ServingNodes *int32 `json:"ServingNodes,omitempty" xml:"ServingNodes,omitempty"`
	// The state of the node pool. Valid values:
	//
	// - active: The desired state is reached.
	//
	// - scaling: The node pool is scaling out.
	//
	// - removing_nodes: The node pool is scaling in.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of instances, including instances that are being created, in use, and being deleted.
	//
	// example:
	//
	// 10
	TotalNodes *int32 `json:"TotalNodes,omitempty" xml:"TotalNodes,omitempty"`
}

func (s CreateClusterNodePoolResponseBodyNodepoolStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponseBodyNodepoolStatus) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetDesiredNodes() *int32 {
	return s.DesiredNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetFailedNodes() *int32 {
	return s.FailedNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetInitialNodes() *int32 {
	return s.InitialNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetRemovingNodes() *int32 {
	return s.RemovingNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetServingNodes() *int32 {
	return s.ServingNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetState() *string {
	return s.State
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) GetTotalNodes() *int32 {
	return s.TotalNodes
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetDesiredNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.DesiredNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetFailedNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.FailedNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetInitialNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.InitialNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetRemovingNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.RemovingNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetServingNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.ServingNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetState(v string) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.State = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) SetTotalNodes(v int32) *CreateClusterNodePoolResponseBodyNodepoolStatus {
	s.TotalNodes = &v
	return s
}

func (s *CreateClusterNodePoolResponseBodyNodepoolStatus) Validate() error {
	return dara.Validate(s)
}
