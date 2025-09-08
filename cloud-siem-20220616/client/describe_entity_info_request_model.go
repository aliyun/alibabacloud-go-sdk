// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEntityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *DescribeEntityInfoRequest
	GetEntityId() *int64
	SetEntityIdentity(v string) *DescribeEntityInfoRequest
	GetEntityIdentity() *string
	SetIncidentUuid(v string) *DescribeEntityInfoRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *DescribeEntityInfoRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeEntityInfoRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeEntityInfoRequest
	GetRoleType() *int32
	SetSophonTaskId(v string) *DescribeEntityInfoRequest
	GetSophonTaskId() *string
}

type DescribeEntityInfoRequest struct {
	// The logical ID of the entity.
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The feature value of the entity. Fuzzy match is supported.
	//
	// example:
	//
	// test22.php
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s DescribeEntityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEntityInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeEntityInfoRequest) GetEntityIdentity() *string {
	return s.EntityIdentity
}

func (s *DescribeEntityInfoRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeEntityInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEntityInfoRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeEntityInfoRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeEntityInfoRequest) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *DescribeEntityInfoRequest) SetEntityId(v int64) *DescribeEntityInfoRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetEntityIdentity(v string) *DescribeEntityInfoRequest {
	s.EntityIdentity = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetIncidentUuid(v string) *DescribeEntityInfoRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRegionId(v string) *DescribeEntityInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRoleFor(v int64) *DescribeEntityInfoRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRoleType(v int32) *DescribeEntityInfoRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetSophonTaskId(v string) *DescribeEntityInfoRequest {
	s.SophonTaskId = &v
	return s
}

func (s *DescribeEntityInfoRequest) Validate() error {
	return dara.Validate(s)
}
