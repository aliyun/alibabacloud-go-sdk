// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackPathSensitiveAssetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathAssetList(v []*CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) *CreateAttackPathSensitiveAssetConfigRequest
	GetAttackPathAssetList() []*CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList
	SetConfigType(v string) *CreateAttackPathSensitiveAssetConfigRequest
	GetConfigType() *string
}

type CreateAttackPathSensitiveAssetConfigRequest struct {
	// The list of cloud service assets.
	//
	// This parameter is required.
	AttackPathAssetList []*CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// The configuration type. Valid values:
	//
	// - asset_instance: asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// asset_instance
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
}

func (s CreateAttackPathSensitiveAssetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathSensitiveAssetConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAttackPathSensitiveAssetConfigRequest) GetAttackPathAssetList() []*CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *CreateAttackPathSensitiveAssetConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *CreateAttackPathSensitiveAssetConfigRequest) SetAttackPathAssetList(v []*CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) *CreateAttackPathSensitiveAssetConfigRequest {
	s.AttackPathAssetList = v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequest) SetConfigType(v string) *CreateAttackPathSensitiveAssetConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequest) Validate() error {
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

type CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList struct {
	// The subtype of the cloud service asset.
	//
	// > You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the subtype of the cloud service asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud service asset.
	//
	// > You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the type of the cloud service asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The instance ID of the cloud service asset.
	//
	// > You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the instance ID of the cloud service asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-8vb0e8qdaj0yyxjo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the cloud service asset instance.
	//
	// > You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the region ID of the cloud service asset instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vendor of the cloud service asset.
	//
	// > You can call the [ListCloudAssetInstances](~~ListCloudAssetInstances~~) operation to query the vendor of the cloud service asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetAssetSubType(v int32) *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetAssetType(v int32) *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetInstanceId(v string) *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.InstanceId = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetRegionId(v string) *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.RegionId = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) SetVendor(v int32) *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigRequestAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
