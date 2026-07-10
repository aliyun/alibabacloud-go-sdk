// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgMembershipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseOrgMembershipsRequest
	GetDBInstanceId() *string
	SetOrganizationId(v string) *DescribeLangfuseOrgMembershipsRequest
	GetOrganizationId() *string
	SetPageNumber(v int64) *DescribeLangfuseOrgMembershipsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLangfuseOrgMembershipsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLangfuseOrgMembershipsRequest
	GetRegionId() *string
}

type DescribeLangfuseOrgMembershipsRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The Langfuse organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseOrgMembershipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgMembershipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgMembershipsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseOrgMembershipsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseOrgMembershipsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseOrgMembershipsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseOrgMembershipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseOrgMembershipsRequest) SetDBInstanceId(v string) *DescribeLangfuseOrgMembershipsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsRequest) SetOrganizationId(v string) *DescribeLangfuseOrgMembershipsRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsRequest) SetPageNumber(v int64) *DescribeLangfuseOrgMembershipsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsRequest) SetPageSize(v int64) *DescribeLangfuseOrgMembershipsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsRequest) SetRegionId(v string) *DescribeLangfuseOrgMembershipsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsRequest) Validate() error {
	return dara.Validate(s)
}
