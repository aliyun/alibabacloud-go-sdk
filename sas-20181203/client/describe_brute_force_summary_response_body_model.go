// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBruteForceSummary(v *DescribeBruteForceSummaryResponseBodyBruteForceSummary) *DescribeBruteForceSummaryResponseBody
	GetBruteForceSummary() *DescribeBruteForceSummaryResponseBodyBruteForceSummary
	SetRequestId(v string) *DescribeBruteForceSummaryResponseBody
	GetRequestId() *string
}

type DescribeBruteForceSummaryResponseBody struct {
	// The statistics of IP address blocking policies.
	BruteForceSummary *DescribeBruteForceSummaryResponseBodyBruteForceSummary `json:"BruteForceSummary,omitempty" xml:"BruteForceSummary,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AE60EAE3-ABD0-897C-B0F16CAC6C7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBruteForceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponseBody) GetBruteForceSummary() *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	return s.BruteForceSummary
}

func (s *DescribeBruteForceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBruteForceSummaryResponseBody) SetBruteForceSummary(v *DescribeBruteForceSummaryResponseBodyBruteForceSummary) *DescribeBruteForceSummaryResponseBody {
	s.BruteForceSummary = v
	return s
}

func (s *DescribeBruteForceSummaryResponseBody) SetRequestId(v string) *DescribeBruteForceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBody) Validate() error {
	if s.BruteForceSummary != nil {
		if err := s.BruteForceSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBruteForceSummaryResponseBodyBruteForceSummary struct {
	// The number of anti-brute force IP blocking policies.
	//
	// example:
	//
	// 13
	AllStrategyCount *int32 `json:"AllStrategyCount,omitempty" xml:"AllStrategyCount,omitempty"`
	// The number of defense policies.
	//
	// example:
	//
	// 2
	AntiBruteForceRuleCount *string `json:"AntiBruteForceRuleCount,omitempty" xml:"AntiBruteForceRuleCount,omitempty"`
	// The number of custom blocking rules that are in effect.
	//
	// example:
	//
	// 3
	CustomEffectiveCount *string `json:"CustomEffectiveCount,omitempty" xml:"CustomEffectiveCount,omitempty"`
	// The number of custom blocking rules.
	//
	// example:
	//
	// 19730
	CustomRecordCount *string `json:"CustomRecordCount,omitempty" xml:"CustomRecordCount,omitempty"`
	// The number of anti-brute force IP blocking policies enabled.
	//
	// example:
	//
	// 2
	EffectiveCount *int32 `json:"EffectiveCount,omitempty" xml:"EffectiveCount,omitempty"`
	// The number of system blocking rules that are in effect.
	//
	// example:
	//
	// 1
	SystemEffectiveCount *string `json:"SystemEffectiveCount,omitempty" xml:"SystemEffectiveCount,omitempty"`
	// The number of system blocking rules.
	//
	// example:
	//
	// 2
	SystemRecordCount *string `json:"SystemRecordCount,omitempty" xml:"SystemRecordCount,omitempty"`
}

func (s DescribeBruteForceSummaryResponseBodyBruteForceSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceSummaryResponseBodyBruteForceSummary) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetAllStrategyCount() *int32 {
	return s.AllStrategyCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetAntiBruteForceRuleCount() *string {
	return s.AntiBruteForceRuleCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetCustomEffectiveCount() *string {
	return s.CustomEffectiveCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetCustomRecordCount() *string {
	return s.CustomRecordCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetEffectiveCount() *int32 {
	return s.EffectiveCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetSystemEffectiveCount() *string {
	return s.SystemEffectiveCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) GetSystemRecordCount() *string {
	return s.SystemRecordCount
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetAllStrategyCount(v int32) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.AllStrategyCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetAntiBruteForceRuleCount(v string) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.AntiBruteForceRuleCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetCustomEffectiveCount(v string) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.CustomEffectiveCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetCustomRecordCount(v string) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.CustomRecordCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetEffectiveCount(v int32) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.EffectiveCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetSystemEffectiveCount(v string) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.SystemEffectiveCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetSystemRecordCount(v string) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.SystemRecordCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) Validate() error {
	return dara.Validate(s)
}
