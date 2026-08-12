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
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of your WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-53y4******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of bot rule tag to query.
	//
	// example:
	//
	// human_machine_challenge
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 200. Default value: 20. This parameter is used for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. If a value is returned for this parameter, the next page exists.
	//
	// > If this parameter has a return value, the next page exists. Use the returned NextToken value as a request parameter to retrieve the next page of data. Repeat until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of bot management protection scenario.
	//
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
