// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseCountStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest
	GetResourceGroupId() *string
}

type DescribeDefenseCountStatisticsRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDefenseCountStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
