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
	// The IDs of the deployment sets. The value can be a JSON array that consists of deployment set IDs in the format of `["ds-xxxxxxxxx", "ds-yyyyyyyyy", ... "ds-zzzzzzzzz"]`. You can specify up to 100 deployment set IDs in each request. Separate the deployment set IDs with commas (,).
	//
	// example:
	//
	// ["ds-2zeeuw16zo2gr9e6****"]
	DeploymentSetIds *string `json:"DeploymentSetIds,omitempty" xml:"DeploymentSetIds,omitempty"`
	// The deployment set name. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, letters, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// deployment_test
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment strategy. Valid values:
	//
	// 	- **Availability**: high availability strategy
	//
	// 	- **AvailabilityGroup**: high availability group strategy
	//
	// Default value: Availability.
	//
	// example:
	//
	// Availability
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
