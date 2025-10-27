// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttackPathSensitiveAssetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathAssetList(v []*UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) *UpdateAttackPathSensitiveAssetConfigRequest
	GetAttackPathAssetList() []*UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList
	SetAttackPathSensitiveAssetConfigId(v string) *UpdateAttackPathSensitiveAssetConfigRequest
	GetAttackPathSensitiveAssetConfigId() *string
}

type UpdateAttackPathSensitiveAssetConfigRequest struct {
	// List of cloud product assets in the attack path.
	//
	// This parameter is required.
	AttackPathAssetList []*UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// ID of the sensitive asset setting for the attack path.
	//
	// This parameter is required.
	//
	// example:
	//
	// apsac-123
	AttackPathSensitiveAssetConfigId *string `json:"AttackPathSensitiveAssetConfigId,omitempty" xml:"AttackPathSensitiveAssetConfigId,omitempty"`
}

func (s UpdateAttackPathSensitiveAssetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathSensitiveAssetConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathSensitiveAssetConfigRequest) GetAttackPathAssetList() []*UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *UpdateAttackPathSensitiveAssetConfigRequest) GetAttackPathSensitiveAssetConfigId() *string {
	return s.AttackPathSensitiveAssetConfigId
}

func (s *UpdateAttackPathSensitiveAssetConfigRequest) SetAttackPathAssetList(v []*UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) *UpdateAttackPathSensitiveAssetConfigRequest {
	s.AttackPathAssetList = v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequest) SetAttackPathSensitiveAssetConfigId(v string) *UpdateAttackPathSensitiveAssetConfigRequest {
	s.AttackPathSensitiveAssetConfigId = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequest) Validate() error {
	if s.AttackPathAssetList != nil {
		for _, item := range s.AttackPathAssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList struct {
	// Subtype of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the subtype of the cloud product asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// Type of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the type of the cloud product asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Cloud product asset instance ID.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the cloud product asset instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze357b4mrkwi7tq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID of the cloud product asset instance.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the region ID of the cloud product asset instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Vendor of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the vendor of the cloud product asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetAssetSubType(v int32) *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetAssetType(v int32) *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetInstanceId(v string) *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetRegionId(v string) *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.RegionId = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetVendor(v int32) *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
