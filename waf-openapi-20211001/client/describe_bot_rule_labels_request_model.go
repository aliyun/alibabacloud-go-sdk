// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotRuleLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBotRuleLabelsRequest
	GetInstanceId() *string
	SetLabelType(v string) *DescribeBotRuleLabelsRequest
	GetLabelType() *string
	SetMaxResults(v int32) *DescribeBotRuleLabelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeBotRuleLabelsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeBotRuleLabelsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeBotRuleLabelsRequest
	GetResourceManagerResourceGroupId() *string
	SetSubScene(v string) *DescribeBotRuleLabelsRequest
	GetSubScene() *string
}

type DescribeBotRuleLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-53y4******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// human_machine_challenge
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// app
	SubScene *string `json:"SubScene,omitempty" xml:"SubScene,omitempty"`
}

func (s DescribeBotRuleLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotRuleLabelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBotRuleLabelsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBotRuleLabelsRequest) GetLabelType() *string {
	return s.LabelType
}

func (s *DescribeBotRuleLabelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeBotRuleLabelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBotRuleLabelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBotRuleLabelsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeBotRuleLabelsRequest) GetSubScene() *string {
	return s.SubScene
}

func (s *DescribeBotRuleLabelsRequest) SetInstanceId(v string) *DescribeBotRuleLabelsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetLabelType(v string) *DescribeBotRuleLabelsRequest {
	s.LabelType = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetMaxResults(v int32) *DescribeBotRuleLabelsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetNextToken(v string) *DescribeBotRuleLabelsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetRegionId(v string) *DescribeBotRuleLabelsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetResourceManagerResourceGroupId(v string) *DescribeBotRuleLabelsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) SetSubScene(v string) *DescribeBotRuleLabelsRequest {
	s.SubScene = &v
	return s
}

func (s *DescribeBotRuleLabelsRequest) Validate() error {
	return dara.Validate(s)
}
