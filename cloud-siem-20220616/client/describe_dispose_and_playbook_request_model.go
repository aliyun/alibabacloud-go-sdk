// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeAndPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableOnly(v bool) *DescribeDisposeAndPlaybookRequest
	GetAvailableOnly() *bool
	SetCurrentPage(v int32) *DescribeDisposeAndPlaybookRequest
	GetCurrentPage() *int32
	SetEntityType(v string) *DescribeDisposeAndPlaybookRequest
	GetEntityType() *string
	SetEntityUuid(v string) *DescribeDisposeAndPlaybookRequest
	GetEntityUuid() *string
	SetEntityUuidList(v string) *DescribeDisposeAndPlaybookRequest
	GetEntityUuidList() *string
	SetIncidentUuid(v string) *DescribeDisposeAndPlaybookRequest
	GetIncidentUuid() *string
	SetPageSize(v int32) *DescribeDisposeAndPlaybookRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDisposeAndPlaybookRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeDisposeAndPlaybookRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeDisposeAndPlaybookRequest
	GetRoleType() *int32
}

type DescribeDisposeAndPlaybookRequest struct {
	// Specifies whether to return only available playbooks.
	AvailableOnly *bool `json:"AvailableOnly,omitempty" xml:"AvailableOnly,omitempty"`
	// The current page number. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The entity type. Valid values:
	//
	// - ip: IP address.
	//
	// - process: Process.
	//
	// - file: File.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The UUID of the entity.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The list of entity UUIDs. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****,6c740667-80b2-476d-8924-2e706feb****
	EntityUuidList *string `json:"EntityUuidList,omitempty" xml:"EntityUuidList,omitempty"`
	// The UUID of the incident.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: Current Alibaba Cloud account view.
	//
	// - 1: View of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeDisposeAndPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookRequest) GetAvailableOnly() *bool {
	return s.AvailableOnly
}

func (s *DescribeDisposeAndPlaybookRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDisposeAndPlaybookRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeDisposeAndPlaybookRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *DescribeDisposeAndPlaybookRequest) GetEntityUuidList() *string {
	return s.EntityUuidList
}

func (s *DescribeDisposeAndPlaybookRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeDisposeAndPlaybookRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisposeAndPlaybookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDisposeAndPlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeDisposeAndPlaybookRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeDisposeAndPlaybookRequest) SetAvailableOnly(v bool) *DescribeDisposeAndPlaybookRequest {
	s.AvailableOnly = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityType(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityUuid(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityUuidList(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityUuidList = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetIncidentUuid(v string) *DescribeDisposeAndPlaybookRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetPageSize(v int32) *DescribeDisposeAndPlaybookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRegionId(v string) *DescribeDisposeAndPlaybookRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRoleFor(v int64) *DescribeDisposeAndPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRoleType(v int32) *DescribeDisposeAndPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
