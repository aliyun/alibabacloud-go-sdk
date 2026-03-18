// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDnsGtmAccessStrategiesResponseBody
	GetRequestId() *string
	SetStrategies(v *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) *DescribeDnsGtmAccessStrategiesResponseBody
	GetStrategies() *DescribeDnsGtmAccessStrategiesResponseBodyStrategies
	SetTotalItems(v int32) *DescribeDnsGtmAccessStrategiesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDnsGtmAccessStrategiesResponseBody
	GetTotalPages() *int32
}

type DescribeDnsGtmAccessStrategiesResponseBody struct {
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0CCC9971-CEC9-4132-824B-4AE611C07623
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Strategies *DescribeDnsGtmAccessStrategiesResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Struct"`
	// The total number of entries returned on all pages.
	//
	// example:
	//
	// 11
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 11
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetStrategies() *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetStrategies(v *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetTotalItems(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) SetTotalPages(v int32) *DescribeDnsGtmAccessStrategiesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBody) Validate() error {
	if s.Strategies != nil {
		if err := s.Strategies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategies struct {
	Strategy []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) GetStrategy() []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	return s.Strategy
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) SetStrategy(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) *DescribeDnsGtmAccessStrategiesResponseBodyStrategies {
	s.Strategy = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategies) Validate() error {
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

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy struct {
	CreateTime                 *string                                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp            *int64                                                                          `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EffectiveAddrPoolGroupType *string                                                                         `json:"EffectiveAddrPoolGroupType,omitempty" xml:"EffectiveAddrPoolGroupType,omitempty"`
	EffectiveAddrPoolType      *string                                                                         `json:"EffectiveAddrPoolType,omitempty" xml:"EffectiveAddrPoolType,omitempty"`
	EffectiveAddrPools         *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools `json:"EffectiveAddrPools,omitempty" xml:"EffectiveAddrPools,omitempty" type:"Struct"`
	EffectiveLbaStrategy       *string                                                                         `json:"EffectiveLbaStrategy,omitempty" xml:"EffectiveLbaStrategy,omitempty"`
	Lines                      *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines              `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	StrategyId                 *string                                                                         `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName               *string                                                                         `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetEffectiveAddrPoolGroupType() *string {
	return s.EffectiveAddrPoolGroupType
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetEffectiveAddrPoolType() *string {
	return s.EffectiveAddrPoolType
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetEffectiveAddrPools() *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools {
	return s.EffectiveAddrPools
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetEffectiveLbaStrategy() *string {
	return s.EffectiveLbaStrategy
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetLines() *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines {
	return s.Lines
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetCreateTime(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetCreateTimestamp(v int64) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetEffectiveAddrPoolGroupType(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.EffectiveAddrPoolGroupType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetEffectiveAddrPoolType(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.EffectiveAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetEffectiveAddrPools(v *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.EffectiveAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetEffectiveLbaStrategy(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.EffectiveLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetLines(v *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.Lines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetStrategyId(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.StrategyId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) SetStrategyName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy {
	s.StrategyName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategy) Validate() error {
	if s.EffectiveAddrPools != nil {
		if err := s.EffectiveAddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.Lines != nil {
		if err := s.Lines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools struct {
	EffectiveAddrPool []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool `json:"EffectiveAddrPool,omitempty" xml:"EffectiveAddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) GetEffectiveAddrPool() []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool {
	return s.EffectiveAddrPool
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) SetEffectiveAddrPool(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools {
	s.EffectiveAddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPools) Validate() error {
	if s.EffectiveAddrPool != nil {
		for _, item := range s.EffectiveAddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool struct {
	AddrCount *int32  `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LbaWeight *int32  `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) SetId(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) SetName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyEffectiveAddrPoolsEffectiveAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines struct {
	Line []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) GetLine() []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	return s.Line
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) SetLine(v []*DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines {
	s.Line = v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLines) Validate() error {
	if s.Line != nil {
		for _, item := range s.Line {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine struct {
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LineCode  *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	LineName  *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetGroupCode(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetGroupName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetLineCode(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) SetLineName(v string) *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesResponseBodyStrategiesStrategyLinesLine) Validate() error {
	return dara.Validate(s)
}
