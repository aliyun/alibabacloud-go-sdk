// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseRuleStatisticsResponseBody
	GetRequestId() *string
	SetStatisticsInfos(v []*DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) *DescribeDefenseRuleStatisticsResponseBody
	GetStatisticsInfos() []*DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos
}

type DescribeDefenseRuleStatisticsResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId       *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatisticsInfos []*DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos `json:"StatisticsInfos,omitempty" xml:"StatisticsInfos,omitempty" type:"Repeated"`
}

func (s DescribeDefenseRuleStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseRuleStatisticsResponseBody) GetStatisticsInfos() []*DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	return s.StatisticsInfos
}

func (s *DescribeDefenseRuleStatisticsResponseBody) SetRequestId(v string) *DescribeDefenseRuleStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBody) SetStatisticsInfos(v []*DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) *DescribeDefenseRuleStatisticsResponseBody {
	s.StatisticsInfos = v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBody) Validate() error {
	if s.StatisticsInfos != nil {
		for _, item := range s.StatisticsInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos struct {
	// example:
	//
	// 27
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// action
	FourthValue *string `json:"FourthValue,omitempty" xml:"FourthValue,omitempty"`
	// example:
	//
	// sytem
	PrimaryValue *string `json:"PrimaryValue,omitempty" xml:"PrimaryValue,omitempty"`
	// example:
	//
	// block
	SecondaryValue *string `json:"SecondaryValue,omitempty" xml:"SecondaryValue,omitempty"`
	// example:
	//
	// 1
	ThirdValue *string `json:"ThirdValue,omitempty" xml:"ThirdValue,omitempty"`
}

func (s DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GetFourthValue() *string {
	return s.FourthValue
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GetPrimaryValue() *string {
	return s.PrimaryValue
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GetSecondaryValue() *string {
	return s.SecondaryValue
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) GetThirdValue() *string {
	return s.ThirdValue
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) SetCount(v int64) *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	s.Count = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) SetFourthValue(v string) *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	s.FourthValue = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) SetPrimaryValue(v string) *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	s.PrimaryValue = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) SetSecondaryValue(v string) *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	s.SecondaryValue = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) SetThirdValue(v string) *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos {
	s.ThirdValue = &v
	return s
}

func (s *DescribeDefenseRuleStatisticsResponseBodyStatisticsInfos) Validate() error {
	return dara.Validate(s)
}
