// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathWhitelist(v *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) *GetAttackPathWhitelistResponseBody
	GetAttackPathWhitelist() *GetAttackPathWhitelistResponseBodyAttackPathWhitelist
	SetRequestId(v string) *GetAttackPathWhitelistResponseBody
	GetRequestId() *string
}

type GetAttackPathWhitelistResponseBody struct {
	// The attack path whitelist.
	AttackPathWhitelist *GetAttackPathWhitelistResponseBodyAttackPathWhitelist `json:"AttackPathWhitelist,omitempty" xml:"AttackPathWhitelist,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttackPathWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackPathWhitelistResponseBody) GetAttackPathWhitelist() *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	return s.AttackPathWhitelist
}

func (s *GetAttackPathWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttackPathWhitelistResponseBody) SetAttackPathWhitelist(v *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) *GetAttackPathWhitelistResponseBody {
	s.AttackPathWhitelist = v
	return s
}

func (s *GetAttackPathWhitelistResponseBody) SetRequestId(v string) *GetAttackPathWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBody) Validate() error {
	if s.AttackPathWhitelist != nil {
		if err := s.AttackPathWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAttackPathWhitelistResponseBodyAttackPathWhitelist struct {
	// The list of cloud service assets in the attack path.
	AttackPathAssetList []*GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList `json:"AttackPathAssetList,omitempty" xml:"AttackPathAssetList,omitempty" type:"Repeated"`
	// The attack path whitelist ID.
	//
	// example:
	//
	// apwl-b33dec0acf9b42aabde032d656c0****
	AttackPathWhitelistId *string `json:"AttackPathWhitelistId,omitempty" xml:"AttackPathWhitelistId,omitempty"`
	// The timestamp of the last modification, in milliseconds.
	//
	// example:
	//
	// 1743004587000
	LastModifiedTimestamp *int64 `json:"LastModifiedTimestamp,omitempty" xml:"LastModifiedTimestamp,omitempty"`
	// The path name.
	//
	// example:
	//
	// ecs_get_credential_by_create_login_profile
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// The path type.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The whitelist name.
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

func (s GetAttackPathWhitelistResponseBodyAttackPathWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GoString() string {
	return s.String()
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetAttackPathAssetList() []*GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	return s.AttackPathAssetList
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetAttackPathWhitelistId() *string {
	return s.AttackPathWhitelistId
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetLastModifiedTimestamp() *int64 {
	return s.LastModifiedTimestamp
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetPathName() *string {
	return s.PathName
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetPathType() *string {
	return s.PathType
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetRemark() *string {
	return s.Remark
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetWhitelistName() *string {
	return s.WhitelistName
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetAttackPathAssetList(v []*GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.AttackPathAssetList = v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetAttackPathWhitelistId(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.AttackPathWhitelistId = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetLastModifiedTimestamp(v int64) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.LastModifiedTimestamp = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetPathName(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.PathName = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetPathType(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.PathType = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetRemark(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.Remark = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetWhitelistName(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.WhitelistName = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) SetWhitelistType(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.WhitelistType = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelist) Validate() error {
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

type GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList struct {
	// The subtype of the cloud service asset.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud service asset.
	//
	// example:
	//
	// 1
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The instance ID of the cloud service asset.
	//
	// example:
	//
	// i-8vb0e8qdaj0yyxjo****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// xwl
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	// The region ID of the cloud service asset instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vendor of the cloud service asset.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GoString() string {
	return s.String()
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetAssetSubType(v int32) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.AssetSubType = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetAssetType(v int32) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.AssetType = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetInstanceId(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.InstanceId = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetInstanceName(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.InstanceName = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetNodeType(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.NodeType = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetRegionId(v string) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.RegionId = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) SetVendor(v int32) *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList {
	s.Vendor = &v
	return s
}

func (s *GetAttackPathWhitelistResponseBodyAttackPathWhitelistAttackPathAssetList) Validate() error {
	return dara.Validate(s)
}
