// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDeploymentSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetIds(v string) *DescribeRCDeploymentSetsRequest
	GetDeploymentSetIds() *string
	SetDeploymentSetName(v string) *DescribeRCDeploymentSetsRequest
	GetDeploymentSetName() *string
	SetPageNumber(v int32) *DescribeRCDeploymentSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCDeploymentSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCDeploymentSetsRequest
	GetRegionId() *string
	SetStrategy(v string) *DescribeRCDeploymentSetsRequest
	GetStrategy() *string
	SetTag(v string) *DescribeRCDeploymentSetsRequest
	GetTag() *string
}

type DescribeRCDeploymentSetsRequest struct {
	DeploymentSetIds  *string `json:"DeploymentSetIds,omitempty" xml:"DeploymentSetIds,omitempty"`
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeRCDeploymentSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsRequest) GetDeploymentSetIds() *string {
	return s.DeploymentSetIds
}

func (s *DescribeRCDeploymentSetsRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *DescribeRCDeploymentSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCDeploymentSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCDeploymentSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCDeploymentSetsRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeRCDeploymentSetsRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeRCDeploymentSetsRequest) SetDeploymentSetIds(v string) *DescribeRCDeploymentSetsRequest {
	s.DeploymentSetIds = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetDeploymentSetName(v string) *DescribeRCDeploymentSetsRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetPageNumber(v int32) *DescribeRCDeploymentSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetPageSize(v int32) *DescribeRCDeploymentSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetRegionId(v string) *DescribeRCDeploymentSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetStrategy(v string) *DescribeRCDeploymentSetsRequest {
	s.Strategy = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) SetTag(v string) *DescribeRCDeploymentSetsRequest {
	s.Tag = &v
	return s
}

func (s *DescribeRCDeploymentSetsRequest) Validate() error {
	return dara.Validate(s)
}
