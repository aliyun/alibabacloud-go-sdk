// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *DescribeWhiteRuleListRequest
	GetAlertName() *string
	SetAlertType(v string) *DescribeWhiteRuleListRequest
	GetAlertType() *string
	SetCurrentPage(v int32) *DescribeWhiteRuleListRequest
	GetCurrentPage() *int32
	SetIncidentUuid(v string) *DescribeWhiteRuleListRequest
	GetIncidentUuid() *string
	SetPageSize(v int32) *DescribeWhiteRuleListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeWhiteRuleListRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeWhiteRuleListRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeWhiteRuleListRequest
	GetRoleType() *int32
}

type DescribeWhiteRuleListRequest struct {
	// The alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The alert type.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The page number. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The universally unique identifier (UUID) of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. The maximum value is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of Threat Analysis is deployed. You must select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are deployed in the Chinese mainland or in the China (Hong Kong) region.
	//
	// - ap-southeast-1: Your assets are deployed in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member. This parameter is used in a multi-account management scenario. An administrator can specify this parameter to query the data of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that are managed by the administrator account.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeWhiteRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteRuleListRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeWhiteRuleListRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeWhiteRuleListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteRuleListRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeWhiteRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWhiteRuleListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeWhiteRuleListRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeWhiteRuleListRequest) SetAlertName(v string) *DescribeWhiteRuleListRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetAlertType(v string) *DescribeWhiteRuleListRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetCurrentPage(v int32) *DescribeWhiteRuleListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetIncidentUuid(v string) *DescribeWhiteRuleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetPageSize(v int32) *DescribeWhiteRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRegionId(v string) *DescribeWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRoleFor(v int64) *DescribeWhiteRuleListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) SetRoleType(v int32) *DescribeWhiteRuleListRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeWhiteRuleListRequest) Validate() error {
	return dara.Validate(s)
}
