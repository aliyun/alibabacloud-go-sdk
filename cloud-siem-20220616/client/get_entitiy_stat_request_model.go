// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntitiyStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *GetEntitiyStatRequest
	GetAssetName() *string
	SetAssetUuid(v string) *GetEntitiyStatRequest
	GetAssetUuid() *string
	SetEntityName(v string) *GetEntitiyStatRequest
	GetEntityName() *string
	SetEntityType(v string) *GetEntitiyStatRequest
	GetEntityType() *string
	SetEntityUuid(v string) *GetEntitiyStatRequest
	GetEntityUuid() *string
	SetIncidentUuid(v string) *GetEntitiyStatRequest
	GetIncidentUuid() *string
	SetIsAsset(v string) *GetEntitiyStatRequest
	GetIsAsset() *string
	SetIsMalwareEntity(v string) *GetEntitiyStatRequest
	GetIsMalwareEntity() *string
	SetRegionId(v string) *GetEntitiyStatRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetEntitiyStatRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *GetEntitiyStatRequest
	GetRoleType() *int32
	SetTags(v string) *GetEntitiyStatRequest
	GetTags() *string
}

type GetEntitiyStatRequest struct {
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	AssetUuid *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The incident ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The asset ID associated with the incident.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	IsAsset *string `json:"IsAsset,omitempty" xml:"IsAsset,omitempty"`
	// The sort order of the incident list. Valid values:
	//
	// - desc: descending order.
	//
	// - asc: ascending order.
	//
	// example:
	//
	// desc
	IsMalwareEntity *string `json:"IsMalwareEntity,omitempty" xml:"IsMalwareEntity,omitempty"`
	// The region where the threat detection and response data management center resides. Select the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: the asset belongs to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: the asset belongs to a region outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 1637941677243702
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: single-account logon.
	//
	// - 1: global view.
	//
	// - 2: switched view.
	//
	// - 3: partial view.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The entity tags. The value is a JSON array string in the following format: \\"[{\\"tagKey1\\":\\"tagValue1\\"},{\\"tagKey2\\":\\"tagValue2\\"}]\\"
	//
	// example:
	//
	// sys:agent:dispose
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetEntitiyStatRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEntitiyStatRequest) GoString() string {
	return s.String()
}

func (s *GetEntitiyStatRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *GetEntitiyStatRequest) GetAssetUuid() *string {
	return s.AssetUuid
}

func (s *GetEntitiyStatRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *GetEntitiyStatRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetEntitiyStatRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *GetEntitiyStatRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *GetEntitiyStatRequest) GetIsAsset() *string {
	return s.IsAsset
}

func (s *GetEntitiyStatRequest) GetIsMalwareEntity() *string {
	return s.IsMalwareEntity
}

func (s *GetEntitiyStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEntitiyStatRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetEntitiyStatRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *GetEntitiyStatRequest) GetTags() *string {
	return s.Tags
}

func (s *GetEntitiyStatRequest) SetAssetName(v string) *GetEntitiyStatRequest {
	s.AssetName = &v
	return s
}

func (s *GetEntitiyStatRequest) SetAssetUuid(v string) *GetEntitiyStatRequest {
	s.AssetUuid = &v
	return s
}

func (s *GetEntitiyStatRequest) SetEntityName(v string) *GetEntitiyStatRequest {
	s.EntityName = &v
	return s
}

func (s *GetEntitiyStatRequest) SetEntityType(v string) *GetEntitiyStatRequest {
	s.EntityType = &v
	return s
}

func (s *GetEntitiyStatRequest) SetEntityUuid(v string) *GetEntitiyStatRequest {
	s.EntityUuid = &v
	return s
}

func (s *GetEntitiyStatRequest) SetIncidentUuid(v string) *GetEntitiyStatRequest {
	s.IncidentUuid = &v
	return s
}

func (s *GetEntitiyStatRequest) SetIsAsset(v string) *GetEntitiyStatRequest {
	s.IsAsset = &v
	return s
}

func (s *GetEntitiyStatRequest) SetIsMalwareEntity(v string) *GetEntitiyStatRequest {
	s.IsMalwareEntity = &v
	return s
}

func (s *GetEntitiyStatRequest) SetRegionId(v string) *GetEntitiyStatRequest {
	s.RegionId = &v
	return s
}

func (s *GetEntitiyStatRequest) SetRoleFor(v int64) *GetEntitiyStatRequest {
	s.RoleFor = &v
	return s
}

func (s *GetEntitiyStatRequest) SetRoleType(v int32) *GetEntitiyStatRequest {
	s.RoleType = &v
	return s
}

func (s *GetEntitiyStatRequest) SetTags(v string) *GetEntitiyStatRequest {
	s.Tags = &v
	return s
}

func (s *GetEntitiyStatRequest) Validate() error {
	return dara.Validate(s)
}
