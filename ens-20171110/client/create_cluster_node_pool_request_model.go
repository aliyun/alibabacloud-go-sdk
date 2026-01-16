// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterNodePoolRequest
	GetClusterId() *string
	SetKubernetesConfig(v *CreateClusterNodePoolRequestKubernetesConfig) *CreateClusterNodePoolRequest
	GetKubernetesConfig() *CreateClusterNodePoolRequestKubernetesConfig
	SetNodepoolInfo(v *CreateClusterNodePoolRequestNodepoolInfo) *CreateClusterNodePoolRequest
	GetNodepoolInfo() *CreateClusterNodePoolRequestNodepoolInfo
	SetScalingGroup(v *CreateClusterNodePoolRequestScalingGroup) *CreateClusterNodePoolRequest
	GetScalingGroup() *CreateClusterNodePoolRequestScalingGroup
}

type CreateClusterNodePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId        *string                                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfig *CreateClusterNodePoolRequestKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// This parameter is required.
	NodepoolInfo *CreateClusterNodePoolRequestNodepoolInfo `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty" type:"Struct"`
	// This parameter is required.
	ScalingGroup *CreateClusterNodePoolRequestScalingGroup `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterNodePoolRequest) GetKubernetesConfig() *CreateClusterNodePoolRequestKubernetesConfig {
	return s.KubernetesConfig
}

func (s *CreateClusterNodePoolRequest) GetNodepoolInfo() *CreateClusterNodePoolRequestNodepoolInfo {
	return s.NodepoolInfo
}

func (s *CreateClusterNodePoolRequest) GetScalingGroup() *CreateClusterNodePoolRequestScalingGroup {
	return s.ScalingGroup
}

func (s *CreateClusterNodePoolRequest) SetClusterId(v string) *CreateClusterNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetKubernetesConfig(v *CreateClusterNodePoolRequestKubernetesConfig) *CreateClusterNodePoolRequest {
	s.KubernetesConfig = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetNodepoolInfo(v *CreateClusterNodePoolRequestNodepoolInfo) *CreateClusterNodePoolRequest {
	s.NodepoolInfo = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetScalingGroup(v *CreateClusterNodePoolRequestScalingGroup) *CreateClusterNodePoolRequest {
	s.ScalingGroup = v
	return s
}

func (s *CreateClusterNodePoolRequest) Validate() error {
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
	return nil
}

type CreateClusterNodePoolRequestKubernetesConfig struct {
	// example:
	//
	// ZWNobyBvawo=
	PreUserData *string `json:"PreUserData,omitempty" xml:"PreUserData,omitempty"`
	// example:
	//
	// ZWNobyBvawo=
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateClusterNodePoolRequestKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestKubernetesConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetPreUserData(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetUserData(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestNodepoolInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterNodePoolRequestNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodepoolInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) SetName(v string) *CreateClusterNodePoolRequestNodepoolInfo {
	s.Name = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroup struct {
	DataDisks []*CreateClusterNodePoolRequestScalingGroupDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	DesiredSize *int32 `json:"DesiredSize,omitempty" xml:"DesiredSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-6734fzvcwcv2o8jj201cpa7ox
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// This parameter is required.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// key-pair-name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// example:
	//
	// some-test-password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// This parameter is required.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s CreateClusterNodePoolRequestScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDataDisks() []*CreateClusterNodePoolRequestScalingGroupDataDisks {
	return s.DataDisks
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDesiredSize() *int32 {
	return s.DesiredSize
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetPassword() *string {
	return s.Password
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDataDisks(v []*CreateClusterNodePoolRequestScalingGroupDataDisks) *CreateClusterNodePoolRequestScalingGroup {
	s.DataDisks = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDesiredSize(v int32) *CreateClusterNodePoolRequestScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetImageId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.ImageId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstanceChargeType(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstanceTypes(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetKeyPairName(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetPassword(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.Password = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskCategory(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskSize(v int32) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetVswitchIds(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) Validate() error {
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

type CreateClusterNodePoolRequestScalingGroupDataDisks struct {
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupDataDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupDataDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupDataDisks) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterNodePoolRequestScalingGroupDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterNodePoolRequestScalingGroupDataDisks) SetCategory(v string) *CreateClusterNodePoolRequestScalingGroupDataDisks {
	s.Category = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupDataDisks) SetSize(v int32) *CreateClusterNodePoolRequestScalingGroupDataDisks {
	s.Size = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupDataDisks) Validate() error {
	return dara.Validate(s)
}
