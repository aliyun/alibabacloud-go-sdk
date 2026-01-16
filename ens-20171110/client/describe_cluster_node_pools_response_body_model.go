// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodePools(v []*DescribeClusterNodePoolsResponseBodyNodePools) *DescribeClusterNodePoolsResponseBody
	GetNodePools() []*DescribeClusterNodePoolsResponseBodyNodePools
	SetPage(v *DescribeClusterNodePoolsResponseBodyPage) *DescribeClusterNodePoolsResponseBody
	GetPage() *DescribeClusterNodePoolsResponseBodyPage
	SetRequestId(v string) *DescribeClusterNodePoolsResponseBody
	GetRequestId() *string
}

type DescribeClusterNodePoolsResponseBody struct {
	NodePools []*DescribeClusterNodePoolsResponseBodyNodePools `json:"NodePools,omitempty" xml:"NodePools,omitempty" type:"Repeated"`
	Page      *DescribeClusterNodePoolsResponseBodyPage        `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBody) GetNodePools() []*DescribeClusterNodePoolsResponseBodyNodePools {
	return s.NodePools
}

func (s *DescribeClusterNodePoolsResponseBody) GetPage() *DescribeClusterNodePoolsResponseBodyPage {
	return s.Page
}

func (s *DescribeClusterNodePoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterNodePoolsResponseBody) SetNodePools(v []*DescribeClusterNodePoolsResponseBodyNodePools) *DescribeClusterNodePoolsResponseBody {
	s.NodePools = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBody) SetPage(v *DescribeClusterNodePoolsResponseBodyPage) *DescribeClusterNodePoolsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBody) SetRequestId(v string) *DescribeClusterNodePoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBody) Validate() error {
	if s.NodePools != nil {
		for _, item := range s.NodePools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodePoolsResponseBodyNodePools struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId        *string                                                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfig *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	NodepoolInfo     *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo     `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty" type:"Struct"`
	ScalingGroup     *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup     `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
	Status           *DescribeClusterNodePoolsResponseBodyNodePoolsStatus           `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolsResponseBodyNodePools) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePools) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) GetKubernetesConfig() *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig {
	return s.KubernetesConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) GetNodepoolInfo() *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo {
	return s.NodepoolInfo
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) GetScalingGroup() *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	return s.ScalingGroup
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) GetStatus() *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	return s.Status
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) SetClusterId(v string) *DescribeClusterNodePoolsResponseBodyNodePools {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) SetKubernetesConfig(v *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) *DescribeClusterNodePoolsResponseBodyNodePools {
	s.KubernetesConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) SetNodepoolInfo(v *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) *DescribeClusterNodePoolsResponseBodyNodePools {
	s.NodepoolInfo = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) SetScalingGroup(v *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) *DescribeClusterNodePoolsResponseBodyNodePools {
	s.ScalingGroup = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) SetStatus(v *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) *DescribeClusterNodePoolsResponseBodyNodePools {
	s.Status = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePools) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig struct {
	// example:
	//
	// ZWNobyBvawo=
	PreUserData *string `json:"PreUserData,omitempty" xml:"PreUserData,omitempty"`
	// example:
	//
	// ZWNobyBvawo=
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) SetPreUserData(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) SetUserData(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo struct {
	// example:
	//
	// cn-fuzhou-23
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// eck-node-pool-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) SetEnsRegionId(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) SetName(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) SetNodepoolId(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup struct {
	DataDisks []*DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// m-673f5z4h69ibwtallg6zmcaxr
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
	// ************
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

func (s DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetDataDisks() []*DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks {
	return s.DataDisks
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetPassword() *string {
	return s.Password
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetDataDisks(v []*DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.DataDisks = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetImageId(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetInstanceChargeType(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetInstanceTypes(v []*string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetKeyPairName(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetPassword(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.Password = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetSystemDiskCategory(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetSystemDiskSize(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) SetVswitchIds(v []*string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroup) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks struct {
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) SetCategory(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks {
	s.Category = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) SetSize(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks {
	s.Size = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsScalingGroupDataDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodePoolsStatus struct {
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

func (s DescribeClusterNodePoolsResponseBodyNodePoolsStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetDesiredNodes() *int32 {
	return s.DesiredNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetInitialNodes() *int32 {
	return s.InitialNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetRemovingNodes() *int32 {
	return s.RemovingNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetServingNodes() *int32 {
	return s.ServingNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) GetTotalNodes() *int32 {
	return s.TotalNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetDesiredNodes(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.DesiredNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetInitialNodes(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.InitialNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetRemovingNodes(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.RemovingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetServingNodes(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.ServingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetState(v string) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.State = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) SetTotalNodes(v int32) *DescribeClusterNodePoolsResponseBodyNodePoolsStatus {
	s.TotalNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodePoolsStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterNodePoolsResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterNodePoolsResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClusterNodePoolsResponseBodyPage) SetPageNumber(v int32) *DescribeClusterNodePoolsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyPage) SetPageSize(v int32) *DescribeClusterNodePoolsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyPage) SetTotalCount(v int32) *DescribeClusterNodePoolsResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
