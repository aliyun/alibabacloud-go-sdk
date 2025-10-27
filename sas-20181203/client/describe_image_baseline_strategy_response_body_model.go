// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeImageBaselineStrategyResponseBody
	GetRequestId() *string
	SetStrategy(v *DescribeImageBaselineStrategyResponseBodyStrategy) *DescribeImageBaselineStrategyResponseBody
	GetStrategy() *DescribeImageBaselineStrategyResponseBodyStrategy
}

type DescribeImageBaselineStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F85AC10-A1FE-54D7-935A-F28D5256****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the baseline check policy.
	Strategy *DescribeImageBaselineStrategyResponseBodyStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeImageBaselineStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBaselineStrategyResponseBody) GetStrategy() *DescribeImageBaselineStrategyResponseBodyStrategy {
	return s.Strategy
}

func (s *DescribeImageBaselineStrategyResponseBody) SetRequestId(v string) *DescribeImageBaselineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBody) SetStrategy(v *DescribeImageBaselineStrategyResponseBodyStrategy) *DescribeImageBaselineStrategyResponseBody {
	s.Strategy = v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBody) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBaselineStrategyResponseBodyStrategy struct {
	// The baseline check policy for agentless detection.
	//
	// example:
	//
	// hc_win2008_cis_rules
	BaselineItem *string `json:"BaselineItem,omitempty" xml:"BaselineItem,omitempty"`
	// An array that contains the baselines.
	BaselineItemList []*DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList `json:"BaselineItemList,omitempty" xml:"BaselineItemList,omitempty" type:"Repeated"`
	// The number of selected baseline check items.
	//
	// example:
	//
	// 10
	SelectedItemCount *int32 `json:"SelectedItemCount,omitempty" xml:"SelectedItemCount,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 8257
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the baseline check policy.
	//
	// example:
	//
	// default
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The total number of baseline check items.
	//
	// example:
	//
	// 100
	TotalItemCount *int32 `json:"TotalItemCount,omitempty" xml:"TotalItemCount,omitempty"`
	// The type of the baseline check policy. Valid values:
	//
	// 	- **default**: the default policy
	//
	// 	- **full**: a policy that uses all baselines
	//
	// 	- **normal**: a policy that uses general baselines
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageBaselineStrategyResponseBodyStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineStrategyResponseBodyStrategy) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetBaselineItem() *string {
	return s.BaselineItem
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetBaselineItemList() []*DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList {
	return s.BaselineItemList
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetSelectedItemCount() *int32 {
	return s.SelectedItemCount
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetTotalItemCount() *int32 {
	return s.TotalItemCount
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetBaselineItem(v string) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.BaselineItem = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetBaselineItemList(v []*DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.BaselineItemList = v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetSelectedItemCount(v int32) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.SelectedItemCount = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetStrategyId(v int64) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.StrategyId = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetStrategyName(v string) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.StrategyName = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetTotalItemCount(v int32) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.TotalItemCount = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) SetType(v string) *DescribeImageBaselineStrategyResponseBodyStrategy {
	s.Type = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategy) Validate() error {
	if s.BaselineItemList != nil {
		for _, item := range s.BaselineItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList struct {
	// The key of the baseline type.
	//
	// example:
	//
	// identification
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The key of the baseline check item.
	//
	// example:
	//
	// duplicate_pwd_hash
	ItemKey *string `json:"ItemKey,omitempty" xml:"ItemKey,omitempty"`
	// The key of the name for the baseline.
	//
	// example:
	//
	// identification
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
}

func (s DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) GetClassKey() *string {
	return s.ClassKey
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) GetItemKey() *string {
	return s.ItemKey
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) GetNameKey() *string {
	return s.NameKey
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) SetClassKey(v string) *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList {
	s.ClassKey = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) SetItemKey(v string) *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList {
	s.ItemKey = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) SetNameKey(v string) *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList {
	s.NameKey = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponseBodyStrategyBaselineItemList) Validate() error {
	return dara.Validate(s)
}
