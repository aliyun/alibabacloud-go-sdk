// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListEntitiesRequest
	GetCurrentPage() *int32
	SetEntityName(v string) *ListEntitiesRequest
	GetEntityName() *string
	SetEntityType(v string) *ListEntitiesRequest
	GetEntityType() *string
	SetEntityUuid(v string) *ListEntitiesRequest
	GetEntityUuid() *string
	SetEntityUuids(v string) *ListEntitiesRequest
	GetEntityUuids() *string
	SetIncidentUuid(v string) *ListEntitiesRequest
	GetIncidentUuid() *string
	SetIsMalwareEntity(v string) *ListEntitiesRequest
	GetIsMalwareEntity() *string
	SetMalwareType(v string) *ListEntitiesRequest
	GetMalwareType() *string
	SetPageSize(v int32) *ListEntitiesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListEntitiesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListEntitiesRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListEntitiesRequest
	GetRoleType() *int32
	SetTags(v string) *ListEntitiesRequest
	GetTags() *string
}

type ListEntitiesRequest struct {
	// The current page number, which must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The entity name.
	//
	// example:
	//
	// host1****
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The entity type. Valid values:
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity UUID.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The list of entity UUIDs.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****,6c740667-80b2-476d-8924-2e706feb****
	EntityUuids *string `json:"EntityUuids,omitempty" xml:"EntityUuids,omitempty"`
	// The incident ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Specifies whether the entity is malicious. Valid values:
	//
	// example:
	//
	// 1
	IsMalwareEntity *string `json:"IsMalwareEntity,omitempty" xml:"IsMalwareEntity,omitempty"`
	// The malware entity type.
	//
	// example:
	//
	// aliyun.siem.sas.alert_tag.miner_software
	MalwareType *string `json:"MalwareType,omitempty" xml:"MalwareType,omitempty"`
	// The number of entries per page, up to a maximum of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of threat analysis is located. Select the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets reside in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The entity tags. The value is a JSON array string:
	//
	// `"[{"tagKey1":"tagValue1"},{"tagKey2":"tagValue2"}]"`
	//
	// example:
	//
	// [{"tagKey1":"tagValue1"},{"tagKey2":"tagValue2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListEntitiesRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *ListEntitiesRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListEntitiesRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *ListEntitiesRequest) GetEntityUuids() *string {
	return s.EntityUuids
}

func (s *ListEntitiesRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListEntitiesRequest) GetIsMalwareEntity() *string {
	return s.IsMalwareEntity
}

func (s *ListEntitiesRequest) GetMalwareType() *string {
	return s.MalwareType
}

func (s *ListEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEntitiesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListEntitiesRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListEntitiesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListEntitiesRequest) SetCurrentPage(v int32) *ListEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityName(v string) *ListEntitiesRequest {
	s.EntityName = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityType(v string) *ListEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityUuid(v string) *ListEntitiesRequest {
	s.EntityUuid = &v
	return s
}

func (s *ListEntitiesRequest) SetEntityUuids(v string) *ListEntitiesRequest {
	s.EntityUuids = &v
	return s
}

func (s *ListEntitiesRequest) SetIncidentUuid(v string) *ListEntitiesRequest {
	s.IncidentUuid = &v
	return s
}

func (s *ListEntitiesRequest) SetIsMalwareEntity(v string) *ListEntitiesRequest {
	s.IsMalwareEntity = &v
	return s
}

func (s *ListEntitiesRequest) SetMalwareType(v string) *ListEntitiesRequest {
	s.MalwareType = &v
	return s
}

func (s *ListEntitiesRequest) SetPageSize(v int32) *ListEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesRequest) SetRegionId(v string) *ListEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEntitiesRequest) SetRoleFor(v int64) *ListEntitiesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListEntitiesRequest) SetRoleType(v int32) *ListEntitiesRequest {
	s.RoleType = &v
	return s
}

func (s *ListEntitiesRequest) SetTags(v string) *ListEntitiesRequest {
	s.Tags = &v
	return s
}

func (s *ListEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
