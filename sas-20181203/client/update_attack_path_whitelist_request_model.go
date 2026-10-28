// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttackPathWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathAssetList(v []*UpdateAttackPathWhitelistRequestAttackPathAssetList) *UpdateAttackPathWhitelistRequest
	GetAttackPathAssetList() []*UpdateAttackPathWhitelistRequestAttackPathAssetList
	SetAttackPathWhitelistId(v string) *UpdateAttackPathWhitelistRequest
	GetAttackPathWhitelistId() *string
	SetPathName(v string) *UpdateAttackPathWhitelistRequest
	GetPathName() *string
	SetPathType(v string) *UpdateAttackPathWhitelistRequest
	GetPathType() *string
	SetRemark(v string) *UpdateAttackPathWhitelistRequest
	GetRemark() *string
	SetWhitelistName(v string) *UpdateAttackPathWhitelistRequest
	GetWhitelistName() *string
	SetWhitelistType(v string) *UpdateAttackPathWhitelistRequest
	GetWhitelistType() *string
}

type UpdateAttackPathWhitelistRequest struct {
	// The list of cloud assets in the attack path.
	AttackPathAssetList []*UpdateAttackPathWhitelistRequestAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// The ID of the attack path whitelist.
	//
	// > You can call [ListAttackPathWhitelist](~~ListAttackPathWhitelist~~) to query the ID of the attack path whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// apwl-b33dec0acf9b42aabde032d656c0****
	AttackPathWhitelistId *string `json:"AttackPathWhitelistId,omitempty" xml:"AttackPathWhitelistId,omitempty"`
	// The path name.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path name.
	//
	// example:
	//
	// ecs_get_credential_by_create_login_profile
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// The path type.
	//
	// > You can call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query the path type.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The remark.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// test
	WhitelistName *string `json:"WhitelistName,omitempty" xml:"WhitelistName,omitempty"`
	// The whitelist type. Valid values:
	//
	// - **ALL_ASSET**: all assets.
	//
	// - **PART_ASSET**: partial assets.
	//
	// example:
	//
	// ALL_ASSET
	WhitelistType *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
}

func (s UpdateAttackPathWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathWhitelistRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathWhitelistRequest) GetAttackPathAssetList() []*UpdateAttackPathWhitelistRequestAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *UpdateAttackPathWhitelistRequest) GetAttackPathWhitelistId() *string {
	return s.AttackPathWhitelistId
}

func (s *UpdateAttackPathWhitelistRequest) GetPathName() *string {
	return s.PathName
}

func (s *UpdateAttackPathWhitelistRequest) GetPathType() *string {
	return s.PathType
}

func (s *UpdateAttackPathWhitelistRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateAttackPathWhitelistRequest) GetWhitelistName() *string {
	return s.WhitelistName
}

func (s *UpdateAttackPathWhitelistRequest) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *UpdateAttackPathWhitelistRequest) SetAttackPathAssetList(v []*UpdateAttackPathWhitelistRequestAttackPathAssetList) *UpdateAttackPathWhitelistRequest {
	s.AttackPathAssetList = v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetAttackPathWhitelistId(v string) *UpdateAttackPathWhitelistRequest {
	s.AttackPathWhitelistId = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetPathName(v string) *UpdateAttackPathWhitelistRequest {
	s.PathName = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetPathType(v string) *UpdateAttackPathWhitelistRequest {
	s.PathType = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetRemark(v string) *UpdateAttackPathWhitelistRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetWhitelistName(v string) *UpdateAttackPathWhitelistRequest {
	s.WhitelistName = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) SetWhitelistType(v string) *UpdateAttackPathWhitelistRequest {
	s.WhitelistType = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequest) Validate() error {
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

type UpdateAttackPathWhitelistRequestAttackPathAssetList struct {
	// The subtype of the cloud asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the subtype of the cloud asset.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the type of the cloud asset.
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Instance ID of the cloud asset instance.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query instance ID of the cloud asset instance.
	//
	// example:
	//
	// i-8vb0e8qdaj0yyxjo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node type. Valid values:
	//
	// - **start**: start node.
	//
	// - **end**: end node.
	//
	// example:
	//
	// start
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region ID of the cloud asset instance.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the region ID of the cloud asset instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vendor of the cloud asset.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the vendor of the cloud asset.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s UpdateAttackPathWhitelistRequestAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathWhitelistRequestAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetNodeType() *string {
	return s.NodeType
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetAssetSubType(v int32) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetAssetType(v int32) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetInstanceId(v string) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.InstanceId = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetNodeType(v string) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.NodeType = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetRegionId(v string) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.RegionId = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) SetVendor(v int32) *UpdateAttackPathWhitelistRequestAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *UpdateAttackPathWhitelistRequestAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
