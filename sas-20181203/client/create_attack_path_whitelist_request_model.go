// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackPathWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathAssetList(v []*CreateAttackPathWhitelistRequestAttackPathAssetList) *CreateAttackPathWhitelistRequest
	GetAttackPathAssetList() []*CreateAttackPathWhitelistRequestAttackPathAssetList
	SetPathName(v string) *CreateAttackPathWhitelistRequest
	GetPathName() *string
	SetPathType(v string) *CreateAttackPathWhitelistRequest
	GetPathType() *string
	SetRemark(v string) *CreateAttackPathWhitelistRequest
	GetRemark() *string
	SetWhitelistName(v string) *CreateAttackPathWhitelistRequest
	GetWhitelistName() *string
	SetWhitelistType(v string) *CreateAttackPathWhitelistRequest
	GetWhitelistType() *string
}

type CreateAttackPathWhitelistRequest struct {
	// List of cloud product assets in the attack path.
	AttackPathAssetList []*CreateAttackPathWhitelistRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// Path name.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path name.
	//
	// example:
	//
	// ecs_get_credential_by_create_login_profile
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// Path type.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path type.
	//
	// This parameter is required.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// Remark information.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Whitelist name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	WhitelistName *string `json:"WhitelistName,omitempty" xml:"WhitelistName,omitempty"`
	// Whitelist type. Values:
	//
	// - **ALL_ASSET**: All assets
	//
	// - **PART_ASSET**: Partial assets
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL_ASSET
	WhitelistType *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
}

func (s CreateAttackPathWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateAttackPathWhitelistRequest) GetAttackPathAssetList() []*CreateAttackPathWhitelistRequestAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *CreateAttackPathWhitelistRequest) GetPathName() *string {
	return s.PathName
}

func (s *CreateAttackPathWhitelistRequest) GetPathType() *string {
	return s.PathType
}

func (s *CreateAttackPathWhitelistRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateAttackPathWhitelistRequest) GetWhitelistName() *string {
	return s.WhitelistName
}

func (s *CreateAttackPathWhitelistRequest) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *CreateAttackPathWhitelistRequest) SetAttackPathAssetList(v []*CreateAttackPathWhitelistRequestAttackPathAssetList) *CreateAttackPathWhitelistRequest {
	s.AttackPathAssetList = v
	return s
}

func (s *CreateAttackPathWhitelistRequest) SetPathName(v string) *CreateAttackPathWhitelistRequest {
	s.PathName = &v
	return s
}

func (s *CreateAttackPathWhitelistRequest) SetPathType(v string) *CreateAttackPathWhitelistRequest {
	s.PathType = &v
	return s
}

func (s *CreateAttackPathWhitelistRequest) SetRemark(v string) *CreateAttackPathWhitelistRequest {
	s.Remark = &v
	return s
}

func (s *CreateAttackPathWhitelistRequest) SetWhitelistName(v string) *CreateAttackPathWhitelistRequest {
	s.WhitelistName = &v
	return s
}

func (s *CreateAttackPathWhitelistRequest) SetWhitelistType(v string) *CreateAttackPathWhitelistRequest {
	s.WhitelistType = &v
	return s
}

func (s *CreateAttackPathWhitelistRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAttackPathWhitelistRequestAttackPathAssetList struct {
	// Subtype of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the subtype of the cloud product asset.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// Type of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the type of the cloud product asset.
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Cloud product asset instance ID.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the cloud product asset instance ID.
	//
	// example:
	//
	// i-8vb0e8qdaj0yyxjo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Node type, with values:
	//
	// - **start**: Start point.
	//
	// - **end**: End point.
	//
	// example:
	//
	// start
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Region ID of the cloud product asset instance.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the region ID of the cloud product asset instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Vendor of the cloud product asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the vendor of the cloud product asset.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateAttackPathWhitelistRequestAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathWhitelistRequestAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetAssetSubType(v int32) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetAssetType(v int32) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetInstanceId(v string) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.InstanceId = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetNodeType(v string) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.NodeType = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetRegionId(v string) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.RegionId = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) SetVendor(v int32) *CreateAttackPathWhitelistRequestAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *CreateAttackPathWhitelistRequestAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
