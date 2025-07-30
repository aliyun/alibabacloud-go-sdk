// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMonitorRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeJobMonitorRuleRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DescribeJobMonitorRuleRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeJobMonitorRuleRequest
	GetResourceGroupId() *string
}

type DescribeJobMonitorRuleRequest struct {
	// The ID of the data migration, data synchronization, or change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta7w132u12h****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzydi675xfea
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeJobMonitorRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeJobMonitorRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeJobMonitorRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeJobMonitorRuleRequest) SetDtsJobId(v string) *DescribeJobMonitorRuleRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeJobMonitorRuleRequest) SetRegionId(v string) *DescribeJobMonitorRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeJobMonitorRuleRequest) SetResourceGroupId(v string) *DescribeJobMonitorRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeJobMonitorRuleRequest) Validate() error {
	return dara.Validate(s)
}
