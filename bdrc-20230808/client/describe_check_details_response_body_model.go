// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCheckDetailsResponseBodyData) *DescribeCheckDetailsResponseBody
	GetData() *DescribeCheckDetailsResponseBodyData
	SetRequestId(v string) *DescribeCheckDetailsResponseBody
	GetRequestId() *string
}

type DescribeCheckDetailsResponseBody struct {
	// The data returned.
	Data *DescribeCheckDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// 92793A50-0B97-59F1-BAEA-EAED83BA1998
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBody) GetData() *DescribeCheckDetailsResponseBodyData {
	return s.Data
}

func (s *DescribeCheckDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckDetailsResponseBody) SetData(v *DescribeCheckDetailsResponseBodyData) *DescribeCheckDetailsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCheckDetailsResponseBody) SetRequestId(v string) *DescribeCheckDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCheckDetailsResponseBodyData struct {
	// The collection of records returned by this request.
	Content []*DescribeCheckDetailsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of entries returned in this response.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this parameter is empty, it indicates that all data has been retrieved.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAOTzWWYAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM4NzA3NTcwMzY2MjMwNzY2ODcyMzAzMTY2Nzg3ODY5MzY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries that meet the query conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCheckDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBodyData) GetContent() []*DescribeCheckDetailsResponseBodyDataContent {
	return s.Content
}

func (s *DescribeCheckDetailsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCheckDetailsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCheckDetailsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCheckDetailsResponseBodyData) SetContent(v []*DescribeCheckDetailsResponseBodyDataContent) *DescribeCheckDetailsResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetMaxResults(v int32) *DescribeCheckDetailsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetNextToken(v string) *DescribeCheckDetailsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetTotalCount(v int64) *DescribeCheckDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCheckDetailsResponseBodyDataContent struct {
	// The check status. Valid values: NOT_CHECKED, PASSED, FAILED, CHECKING, and CHECK_FAILED.
	//
	// example:
	//
	// PASSED
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The time when the check was performed.
	//
	// example:
	//
	// 1701725715
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// The check details.
	//
	// example:
	//
	// {"ecsAutoSnapshotPolicyIds":[],"hbrBackupPlans":[{"planId":"po-xxxxxxxx","sourceType":"UDM_ECS"}]}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The globally unique Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The unique ID of the resource.
	//
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the resource owner.
	//
	// example:
	//
	// 123***7890
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The unique ID of the rule.
	//
	// example:
	//
	// rule-xxxxxxxx
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule template.
	//
	// example:
	//
	// ecs-backup
	RuleTemplate *string `json:"RuleTemplate,omitempty" xml:"RuleTemplate,omitempty"`
}

func (s DescribeCheckDetailsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckDetailsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetCheckTime() *int64 {
	return s.CheckTime
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetDetail() *string {
	return s.Detail
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCheckDetailsResponseBodyDataContent) GetRuleTemplate() *string {
	return s.RuleTemplate
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetCheckStatus(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.CheckStatus = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetCheckTime(v int64) *DescribeCheckDetailsResponseBodyDataContent {
	s.CheckTime = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetDetail(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.Detail = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetProductType(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceArn(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceId(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceName(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceOwnerId(v int64) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceType(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetRuleId(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.RuleId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetRuleTemplate(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.RuleTemplate = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
