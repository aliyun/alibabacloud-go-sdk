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
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The type of the alert.
	//
	// example:
	//
	// scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
