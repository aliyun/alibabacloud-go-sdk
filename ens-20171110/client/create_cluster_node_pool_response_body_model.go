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
	Nodepool *CreateClusterNodePoolResponseBodyNodepool `json:"Nodepool,omitempty" xml:"Nodepool,omitempty" type:"Struct"`
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
	// example:
	//
	// eck-d2666c5f
	ClusterId        *string                                                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfig *CreateClusterNodePoolResponseBodyNodepoolKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	NodepoolInfo     *CreateClusterNodePoolResponseBodyNodepoolNodepoolInfo     `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty" type:"Struct"`
	ScalingGroup     *CreateClusterNodePoolResponseBodyNodepoolScalingGroup     `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
	Status           *CreateClusterNodePoolResponseBodyNodepoolStatus           `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
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
	// example:
	//
	// ZWNobyBvawo=
	PreUserData *string `json:"PreUserData,omitempty" xml:"PreUserData,omitempty"`
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
	// example:
	//
	// cn-guiyang-14
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	DataDisks []*CreateClusterNodePoolResponseBodyNodepoolScalingGroupDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// m-67bdk1kpu1ylmt7k33h5dbiov
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PrePaid
	InstanceChargeType *string   `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceTypes      []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// key-pair-name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// example:
	//
	// **********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// 80
	SystemDiskSize *int32    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	VswitchIds     []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
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
	// example:
	//
	// 10
	DesiredNodes *int32 `json:"DesiredNodes,omitempty" xml:"DesiredNodes,omitempty"`
	// example:
	//
	// 0
	InitialNodes *int32 `json:"InitialNodes,omitempty" xml:"InitialNodes,omitempty"`
	// example:
	//
	// 0
	RemovingNodes *int32 `json:"RemovingNodes,omitempty" xml:"RemovingNodes,omitempty"`
	// example:
	//
	// 10
	ServingNodes *int32 `json:"ServingNodes,omitempty" xml:"ServingNodes,omitempty"`
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
