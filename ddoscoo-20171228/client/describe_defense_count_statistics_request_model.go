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
	SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest
	GetSourceIp() *string
}

type DescribeDefenseCountStatisticsRequest struct {
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *DescribeDefenseCountStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
