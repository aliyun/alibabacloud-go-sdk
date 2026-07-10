// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectMembershipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseProjectMembershipsRequest
	GetDBInstanceId() *string
	SetOrganizationId(v string) *DescribeLangfuseProjectMembershipsRequest
	GetOrganizationId() *string
	SetPageNumber(v int64) *DescribeLangfuseProjectMembershipsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLangfuseProjectMembershipsRequest
	GetPageSize() *int64
	SetProjectId(v string) *DescribeLangfuseProjectMembershipsRequest
	GetProjectId() *string
	SetRegionId(v string) *DescribeLangfuseProjectMembershipsRequest
	GetRegionId() *string
}

type DescribeLangfuseProjectMembershipsRequest struct {
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
	// The page number.
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
	// The Langfuse project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseProjectMembershipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectMembershipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeLangfuseProjectMembershipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetDBInstanceId(v string) *DescribeLangfuseProjectMembershipsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetOrganizationId(v string) *DescribeLangfuseProjectMembershipsRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetPageNumber(v int64) *DescribeLangfuseProjectMembershipsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetPageSize(v int64) *DescribeLangfuseProjectMembershipsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetProjectId(v string) *DescribeLangfuseProjectMembershipsRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) SetRegionId(v string) *DescribeLangfuseProjectMembershipsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsRequest) Validate() error {
	return dara.Validate(s)
}
