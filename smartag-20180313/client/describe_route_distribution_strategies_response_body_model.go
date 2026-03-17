// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteDistributionStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteDistributionStrategiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteDistributionStrategiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteDistributionStrategiesResponseBody
	GetRequestId() *string
	SetStrategies(v *DescribeRouteDistributionStrategiesResponseBodyStrategies) *DescribeRouteDistributionStrategiesResponseBody
	GetStrategies() *DescribeRouteDistributionStrategiesResponseBodyStrategies
	SetTotalCount(v int32) *DescribeRouteDistributionStrategiesResponseBody
	GetTotalCount() *int32
}

type DescribeRouteDistributionStrategiesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 944C2533-1BB7-4578-B6EB-DA05BB61C02A
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Strategies *DescribeRouteDistributionStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Struct"`
	// The total number of routes.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteDistributionStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteDistributionStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteDistributionStrategiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteDistributionStrategiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteDistributionStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteDistributionStrategiesResponseBody) GetStrategies() *DescribeRouteDistributionStrategiesResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeRouteDistributionStrategiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteDistributionStrategiesResponseBody) SetPageNumber(v int32) *DescribeRouteDistributionStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBody) SetPageSize(v int32) *DescribeRouteDistributionStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBody) SetRequestId(v string) *DescribeRouteDistributionStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBody) SetStrategies(v *DescribeRouteDistributionStrategiesResponseBodyStrategies) *DescribeRouteDistributionStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBody) SetTotalCount(v int32) *DescribeRouteDistributionStrategiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBody) Validate() error {
	if s.Strategies != nil {
		if err := s.Strategies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteDistributionStrategiesResponseBodyStrategies struct {
	Strategy []*DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Repeated"`
}

func (s DescribeRouteDistributionStrategiesResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteDistributionStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategies) GetStrategy() []*DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	return s.Strategy
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategies) SetStrategy(v []*DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) *DescribeRouteDistributionStrategiesResponseBodyStrategies {
	s.Strategy = v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategies) Validate() error {
	if s.Strategy != nil {
		for _, item := range s.Strategy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy struct {
	ConflictInfo          *string `json:"ConflictInfo,omitempty" xml:"ConflictInfo,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DestCidrBlock         *string `json:"DestCidrBlock,omitempty" xml:"DestCidrBlock,omitempty"`
	HcInstanceId          *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	IsConflict            *bool   `json:"IsConflict,omitempty" xml:"IsConflict,omitempty"`
	RouteDistribution     *string `json:"RouteDistribution,omitempty" xml:"RouteDistribution,omitempty"`
	RouteSource           *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	SmartAGId             *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	SourceType            *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StrategyPublishStatus *string `json:"StrategyPublishStatus,omitempty" xml:"StrategyPublishStatus,omitempty"`
}

func (s DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GoString() string {
	return s.String()
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetConflictInfo() *string {
	return s.ConflictInfo
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetDestCidrBlock() *string {
	return s.DestCidrBlock
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetIsConflict() *bool {
	return s.IsConflict
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetRouteDistribution() *string {
	return s.RouteDistribution
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetRouteSource() *string {
	return s.RouteSource
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) GetStrategyPublishStatus() *string {
	return s.StrategyPublishStatus
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetConflictInfo(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.ConflictInfo = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetCreateTime(v int64) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.CreateTime = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetDestCidrBlock(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.DestCidrBlock = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetHcInstanceId(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.HcInstanceId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetIsConflict(v bool) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.IsConflict = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetRouteDistribution(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.RouteDistribution = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetRouteSource(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.RouteSource = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetSmartAGId(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.SmartAGId = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetSourceType(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.SourceType = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetStatus(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.Status = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) SetStrategyPublishStatus(v string) *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy {
	s.StrategyPublishStatus = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponseBodyStrategiesStrategy) Validate() error {
	return dara.Validate(s)
}
