// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeAndPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDisposeAndPlaybookRequest
	GetCurrentPage() *int32
	SetEntityType(v string) *DescribeDisposeAndPlaybookRequest
	GetEntityType() *string
	SetEntityUuid(v string) *DescribeDisposeAndPlaybookRequest
	GetEntityUuid() *string
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
	// 列表当前页号， 大于等于 1。
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The entity type. Valid values:
	//
	// - `ip`: IP address
	//
	// - `process`: process
	//
	// - `file`: file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 实体 uuid。
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The incident UUID.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// 列表每页条数， 最大不超过 100。
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 威胁分析的数据管理中心所在地。您需要根据资产所在地域，选择管理中心所在地。取值：
	//
	// - cn-hangzhou：资产属于中国内地与中国香港
	//
	// - ap-southeast-1：资产属于海外地域
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 管理员切换成其他成员视角的用户 ID。
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// 视图类型。
	//
	// - 0：当前阿里云账号视图。
	//
	// - 1：企业下所有账号的视图。
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

func (s *DescribeDisposeAndPlaybookRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDisposeAndPlaybookRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeDisposeAndPlaybookRequest) GetEntityUuid() *string {
	return s.EntityUuid
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
