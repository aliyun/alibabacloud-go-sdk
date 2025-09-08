// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFourthKey(v string) *DescribeDefenseRuleStatisticsRequest
	GetFourthKey() *string
	SetInstanceId(v string) *DescribeDefenseRuleStatisticsRequest
	GetInstanceId() *string
	SetPrimaryKey(v string) *DescribeDefenseRuleStatisticsRequest
	GetPrimaryKey() *string
	SetRegionId(v string) *DescribeDefenseRuleStatisticsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseRuleStatisticsRequest
	GetResourceManagerResourceGroupId() *string
	SetSecondaryKey(v string) *DescribeDefenseRuleStatisticsRequest
	GetSecondaryKey() *string
	SetTemplateId(v int64) *DescribeDefenseRuleStatisticsRequest
	GetTemplateId() *int64
	SetThirdKey(v string) *DescribeDefenseRuleStatisticsRequest
	GetThirdKey() *string
}

type DescribeDefenseRuleStatisticsRequest struct {
	// example:
	//
	// riskLevel
	FourthKey *string `json:"FourthKey,omitempty" xml:"FourthKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scene
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// action
	SecondaryKey *string `json:"SecondaryKey,omitempty" xml:"SecondaryKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239136
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// status
	ThirdKey *string `json:"ThirdKey,omitempty" xml:"ThirdKey,omitempty"`
}

func (s DescribeDefenseRuleStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleStatisticsRequest) GetFourthKey() *string {
	return s.FourthKey
}

func (s *DescribeDefenseRuleStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseRuleStatisticsRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *DescribeDefenseRuleStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseRuleStatisticsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseRuleStatisticsRequest) GetSecondaryKey() *string {
	return s.SecondaryKey
}

func (s *DescribeDefenseRuleStatisticsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseRuleStatisticsRequest) GetThirdKey() *string {
	return s.ThirdKey
}

func (s *DescribeDefenseRuleStatisticsRequest) SetFourthKey(v string) *DescribeDefenseRuleStatisticsRequest {
	s.FourthKey = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetInstanceId(v string) *DescribeDefenseRuleStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetPrimaryKey(v string) *DescribeDefenseRuleStatisticsRequest {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetRegionId(v string) *DescribeDefenseRuleStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseRuleStatisticsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetSecondaryKey(v string) *DescribeDefenseRuleStatisticsRequest {
	s.SecondaryKey = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetTemplateId(v int64) *DescribeDefenseRuleStatisticsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) SetThirdKey(v string) *DescribeDefenseRuleStatisticsRequest {
	s.ThirdKey = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
