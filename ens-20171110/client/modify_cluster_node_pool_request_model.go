// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterNodePoolRequest
	GetClusterId() *string
	SetKubernetesConfig(v *ModifyClusterNodePoolRequestKubernetesConfig) *ModifyClusterNodePoolRequest
	GetKubernetesConfig() *ModifyClusterNodePoolRequestKubernetesConfig
	SetNodepoolInfo(v *ModifyClusterNodePoolRequestNodepoolInfo) *ModifyClusterNodePoolRequest
	GetNodepoolInfo() *ModifyClusterNodePoolRequestNodepoolInfo
	SetScalingGroup(v *ModifyClusterNodePoolRequestScalingGroup) *ModifyClusterNodePoolRequest
	GetScalingGroup() *ModifyClusterNodePoolRequestScalingGroup
}

type ModifyClusterNodePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId        *string                                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	KubernetesConfig *ModifyClusterNodePoolRequestKubernetesConfig `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty" type:"Struct"`
	// This parameter is required.
	NodepoolInfo *ModifyClusterNodePoolRequestNodepoolInfo `json:"NodepoolInfo,omitempty" xml:"NodepoolInfo,omitempty" type:"Struct"`
	ScalingGroup *ModifyClusterNodePoolRequestScalingGroup `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
}

func (s ModifyClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterNodePoolRequest) GetKubernetesConfig() *ModifyClusterNodePoolRequestKubernetesConfig {
	return s.KubernetesConfig
}

func (s *ModifyClusterNodePoolRequest) GetNodepoolInfo() *ModifyClusterNodePoolRequestNodepoolInfo {
	return s.NodepoolInfo
}

func (s *ModifyClusterNodePoolRequest) GetScalingGroup() *ModifyClusterNodePoolRequestScalingGroup {
	return s.ScalingGroup
}

func (s *ModifyClusterNodePoolRequest) SetClusterId(v string) *ModifyClusterNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetKubernetesConfig(v *ModifyClusterNodePoolRequestKubernetesConfig) *ModifyClusterNodePoolRequest {
	s.KubernetesConfig = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetNodepoolInfo(v *ModifyClusterNodePoolRequestNodepoolInfo) *ModifyClusterNodePoolRequest {
	s.NodepoolInfo = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetScalingGroup(v *ModifyClusterNodePoolRequestScalingGroup) *ModifyClusterNodePoolRequest {
	s.ScalingGroup = v
	return s
}

func (s *ModifyClusterNodePoolRequest) Validate() error {
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

type ModifyClusterNodePoolRequestKubernetesConfig struct {
	// example:
	//
	// ZWNobyBvawo=
	PreUserData *string `json:"PreUserData,omitempty" xml:"PreUserData,omitempty"`
	// example:
	//
	// ZWNobyBvawo=
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ModifyClusterNodePoolRequestKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestKubernetesConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetPreUserData(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetUserData(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestNodepoolInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s ModifyClusterNodePoolRequestNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestNodepoolInfo) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) SetNodepoolId(v string) *ModifyClusterNodePoolRequestNodepoolInfo {
	s.NodepoolId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestScalingGroup struct {
	DataDisks []*ModifyClusterNodePoolRequestScalingGroupDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// m-68bysm5a8gkm7oik7ngacf0in
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string   `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceTypes      []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// tf-testAcc
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// example:
	//
	// @AN5Ejx1YyzzK
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

func (s ModifyClusterNodePoolRequestScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroup) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetDataDisks() []*ModifyClusterNodePoolRequestScalingGroupDataDisks {
	return s.DataDisks
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetPassword() *string {
	return s.Password
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetDataDisks(v []*ModifyClusterNodePoolRequestScalingGroupDataDisks) *ModifyClusterNodePoolRequestScalingGroup {
	s.DataDisks = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetImageId(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.ImageId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInstanceChargeType(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInstanceTypes(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetKeyPairName(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.KeyPairName = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetPassword(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.Password = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskCategory(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskSize(v int32) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetVswitchIds(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) Validate() error {
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

type ModifyClusterNodePoolRequestScalingGroupDataDisks struct {
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyClusterNodePoolRequestScalingGroupDataDisks) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroupDataDisks) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroupDataDisks) GetCategory() *string {
	return s.Category
}

func (s *ModifyClusterNodePoolRequestScalingGroupDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *ModifyClusterNodePoolRequestScalingGroupDataDisks) SetCategory(v string) *ModifyClusterNodePoolRequestScalingGroupDataDisks {
	s.Category = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupDataDisks) SetSize(v int32) *ModifyClusterNodePoolRequestScalingGroupDataDisks {
	s.Size = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupDataDisks) Validate() error {
	return dara.Validate(s)
}
