// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeProductsResponseBodyData) *DescribeProductsResponseBody
	GetData() *DescribeProductsResponseBodyData
	SetRequestId(v string) *DescribeProductsResponseBody
	GetRequestId() *string
}

type DescribeProductsResponseBody struct {
	// The returned data.
	Data *DescribeProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30FB202A-1D22-5394-AB02-4477CDFCF51F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBody) GetData() *DescribeProductsResponseBodyData {
	return s.Data
}

func (s *DescribeProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductsResponseBody) SetData(v *DescribeProductsResponseBodyData) *DescribeProductsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProductsResponseBody) SetRequestId(v string) *DescribeProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProductsResponseBodyData struct {
	// A list of cloud products and their data protection status.
	Content []*DescribeProductsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. If this parameter is absent from the response, all results have been retrieved.
	//
	// example:
	//
	// b4fd3cffcacafd65e3818a0b9b2ff9a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries that match the query. This parameter is not returned by default.
	//
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyData) GetContent() []*DescribeProductsResponseBodyDataContent {
	return s.Content
}

func (s *DescribeProductsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProductsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeProductsResponseBodyData) SetContent(v []*DescribeProductsResponseBodyDataContent) *DescribeProductsResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeProductsResponseBodyData) SetMaxResults(v int32) *DescribeProductsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductsResponseBodyData) SetNextToken(v string) *DescribeProductsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeProductsResponseBodyData) SetTotalCount(v int64) *DescribeProductsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductsResponseBodyData) Validate() error {
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

type DescribeProductsResponseBodyDataContent struct {
	// The count of failed check items.
	//
	// example:
	//
	// 1
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// The count of resources that failed the check.
	//
	// example:
	//
	// 1
	CheckFailedResourceCount *int64 `json:"CheckFailedResourceCount,omitempty" xml:"CheckFailedResourceCount,omitempty"`
	// The count of resources for which the check is disabled.
	//
	// example:
	//
	// 1
	DisableCheckResourceCount *int64 `json:"DisableCheckResourceCount,omitempty" xml:"DisableCheckResourceCount,omitempty"`
	// Indicates whether the data protection score is enabled for the cloud product.
	//
	// example:
	//
	// true
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// The cloud product type, such as `ECS` and `OSS`.
	//
	// example:
	//
	// oss
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The data protection score, ranging from 0 to 100.
	//
	// example:
	//
	// 90
	ProtectionScore *int32 `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	// The distribution of resources across different score ranges.
	ProtectionScoreDistribution []*DescribeProductsResponseBodyDataContentProtectionScoreDistribution `json:"ProtectionScoreDistribution,omitempty" xml:"ProtectionScoreDistribution,omitempty" type:"Repeated"`
	// The UNIX timestamp of the last data protection score update.
	//
	// example:
	//
	// 1726036498
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// The count of risky check items.
	//
	// example:
	//
	// 1
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The count of risky resources.
	//
	// example:
	//
	// 1
	RiskyResourceCount *int64 `json:"RiskyResourceCount,omitempty" xml:"RiskyResourceCount,omitempty"`
	// The total count of resources for the cloud product.
	//
	// example:
	//
	// 100
	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
	// The count of resources pending a check.
	//
	// example:
	//
	// 1
	WaitForCheckResourceCount *int64 `json:"WaitForCheckResourceCount,omitempty" xml:"WaitForCheckResourceCount,omitempty"`
}

func (s DescribeProductsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContent) GetCheckFailedCount() *int64 {
	return s.CheckFailedCount
}

func (s *DescribeProductsResponseBodyDataContent) GetCheckFailedResourceCount() *int64 {
	return s.CheckFailedResourceCount
}

func (s *DescribeProductsResponseBodyDataContent) GetDisableCheckResourceCount() *int64 {
	return s.DisableCheckResourceCount
}

func (s *DescribeProductsResponseBodyDataContent) GetEnableCheck() *bool {
	return s.EnableCheck
}

func (s *DescribeProductsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeProductsResponseBodyDataContent) GetProtectionScore() *int32 {
	return s.ProtectionScore
}

func (s *DescribeProductsResponseBodyDataContent) GetProtectionScoreDistribution() []*DescribeProductsResponseBodyDataContentProtectionScoreDistribution {
	return s.ProtectionScoreDistribution
}

func (s *DescribeProductsResponseBodyDataContent) GetProtectionScoreUpdatedTime() *int64 {
	return s.ProtectionScoreUpdatedTime
}

func (s *DescribeProductsResponseBodyDataContent) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeProductsResponseBodyDataContent) GetRiskyResourceCount() *int64 {
	return s.RiskyResourceCount
}

func (s *DescribeProductsResponseBodyDataContent) GetTotalResourceCount() *int64 {
	return s.TotalResourceCount
}

func (s *DescribeProductsResponseBodyDataContent) GetWaitForCheckResourceCount() *int64 {
	return s.WaitForCheckResourceCount
}

func (s *DescribeProductsResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetCheckFailedResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.CheckFailedResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetDisableCheckResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.DisableCheckResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetEnableCheck(v bool) *DescribeProductsResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProductType(v string) *DescribeProductsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScore(v int32) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScoreDistribution(v []*DescribeProductsResponseBodyDataContentProtectionScoreDistribution) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScoreDistribution = v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetRiskCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetRiskyResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.RiskyResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetTotalResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.TotalResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetWaitForCheckResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.WaitForCheckResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) Validate() error {
	if s.ProtectionScoreDistribution != nil {
		for _, item := range s.ProtectionScoreDistribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProductsResponseBodyDataContentProtectionScoreDistribution struct {
	// The count of resources within this score range.
	//
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The score range.
	Range *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistribution) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistribution) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) GetCount() *int64 {
	return s.Count
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) GetRange() *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange {
	return s.Range
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) SetCount(v int64) *DescribeProductsResponseBodyDataContentProtectionScoreDistribution {
	s.Count = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) SetRange(v *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) *DescribeProductsResponseBodyDataContentProtectionScoreDistribution {
	s.Range = v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) Validate() error {
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange struct {
	// The lower bound of the score range, inclusive.
	//
	// example:
	//
	// 0
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The upper bound of the score range, inclusive.
	//
	// example:
	//
	// 60
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) GetFrom() *int32 {
	return s.From
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) GetTo() *int32 {
	return s.To
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) SetFrom(v int32) *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange {
	s.From = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) SetTo(v int32) *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange {
	s.To = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) Validate() error {
	return dara.Validate(s)
}
