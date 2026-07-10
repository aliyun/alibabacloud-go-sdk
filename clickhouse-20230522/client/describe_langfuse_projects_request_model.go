// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLangfuseProjectsRequest
	GetDBInstanceId() *string
	SetOrganizationId(v string) *DescribeLangfuseProjectsRequest
	GetOrganizationId() *string
	SetRegionId(v string) *DescribeLangfuseProjectsRequest
	GetRegionId() *string
}

type DescribeLangfuseProjectsRequest struct {
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLangfuseProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLangfuseProjectsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLangfuseProjectsRequest) SetDBInstanceId(v string) *DescribeLangfuseProjectsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLangfuseProjectsRequest) SetOrganizationId(v string) *DescribeLangfuseProjectsRequest {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseProjectsRequest) SetRegionId(v string) *DescribeLangfuseProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLangfuseProjectsRequest) Validate() error {
	return dara.Validate(s)
}
