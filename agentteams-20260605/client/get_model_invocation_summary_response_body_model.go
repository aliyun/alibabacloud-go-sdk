// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInvocationSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelInvocationSummaryResponseBody
	GetCode() *string
	SetData(v *GetModelInvocationSummaryResponseBodyData) *GetModelInvocationSummaryResponseBody
	GetData() *GetModelInvocationSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetModelInvocationSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetModelInvocationSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelInvocationSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelInvocationSummaryResponseBody
	GetSuccess() *bool
}

type GetModelInvocationSummaryResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetModelInvocationSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelInvocationSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelInvocationSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelInvocationSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelInvocationSummaryResponseBody) GetData() *GetModelInvocationSummaryResponseBodyData {
	return s.Data
}

func (s *GetModelInvocationSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelInvocationSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelInvocationSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelInvocationSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelInvocationSummaryResponseBody) SetCode(v string) *GetModelInvocationSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) SetData(v *GetModelInvocationSummaryResponseBodyData) *GetModelInvocationSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) SetHttpStatusCode(v int32) *GetModelInvocationSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) SetMessage(v string) *GetModelInvocationSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) SetRequestId(v string) *GetModelInvocationSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) SetSuccess(v bool) *GetModelInvocationSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelInvocationSummaryResponseBodyData struct {
	CallFrequency        *float64                                                         `json:"CallFrequency,omitempty" xml:"CallFrequency,omitempty"`
	ProviderDistribution []*GetModelInvocationSummaryResponseBodyDataProviderDistribution `json:"ProviderDistribution,omitempty" xml:"ProviderDistribution,omitempty" type:"Repeated"`
	TodayCallCount       *int32                                                           `json:"TodayCallCount,omitempty" xml:"TodayCallCount,omitempty"`
	TodayChangeRate      *float64                                                         `json:"TodayChangeRate,omitempty" xml:"TodayChangeRate,omitempty"`
	WeekCallCount        *int32                                                           `json:"WeekCallCount,omitempty" xml:"WeekCallCount,omitempty"`
	WeekChangeRate       *float64                                                         `json:"WeekChangeRate,omitempty" xml:"WeekChangeRate,omitempty"`
}

func (s GetModelInvocationSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelInvocationSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelInvocationSummaryResponseBodyData) GetCallFrequency() *float64 {
	return s.CallFrequency
}

func (s *GetModelInvocationSummaryResponseBodyData) GetProviderDistribution() []*GetModelInvocationSummaryResponseBodyDataProviderDistribution {
	return s.ProviderDistribution
}

func (s *GetModelInvocationSummaryResponseBodyData) GetTodayCallCount() *int32 {
	return s.TodayCallCount
}

func (s *GetModelInvocationSummaryResponseBodyData) GetTodayChangeRate() *float64 {
	return s.TodayChangeRate
}

func (s *GetModelInvocationSummaryResponseBodyData) GetWeekCallCount() *int32 {
	return s.WeekCallCount
}

func (s *GetModelInvocationSummaryResponseBodyData) GetWeekChangeRate() *float64 {
	return s.WeekChangeRate
}

func (s *GetModelInvocationSummaryResponseBodyData) SetCallFrequency(v float64) *GetModelInvocationSummaryResponseBodyData {
	s.CallFrequency = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) SetProviderDistribution(v []*GetModelInvocationSummaryResponseBodyDataProviderDistribution) *GetModelInvocationSummaryResponseBodyData {
	s.ProviderDistribution = v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) SetTodayCallCount(v int32) *GetModelInvocationSummaryResponseBodyData {
	s.TodayCallCount = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) SetTodayChangeRate(v float64) *GetModelInvocationSummaryResponseBodyData {
	s.TodayChangeRate = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) SetWeekCallCount(v int32) *GetModelInvocationSummaryResponseBodyData {
	s.WeekCallCount = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) SetWeekChangeRate(v float64) *GetModelInvocationSummaryResponseBodyData {
	s.WeekChangeRate = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyData) Validate() error {
	if s.ProviderDistribution != nil {
		for _, item := range s.ProviderDistribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelInvocationSummaryResponseBodyDataProviderDistribution struct {
	Count        *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	Percentage   *float64 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	ProviderName *string  `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetModelInvocationSummaryResponseBodyDataProviderDistribution) String() string {
	return dara.Prettify(s)
}

func (s GetModelInvocationSummaryResponseBodyDataProviderDistribution) GoString() string {
	return s.String()
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) GetCount() *int32 {
	return s.Count
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) GetPercentage() *float64 {
	return s.Percentage
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) SetCount(v int32) *GetModelInvocationSummaryResponseBodyDataProviderDistribution {
	s.Count = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) SetPercentage(v float64) *GetModelInvocationSummaryResponseBodyDataProviderDistribution {
	s.Percentage = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) SetProviderName(v string) *GetModelInvocationSummaryResponseBodyDataProviderDistribution {
	s.ProviderName = &v
	return s
}

func (s *GetModelInvocationSummaryResponseBodyDataProviderDistribution) Validate() error {
	return dara.Validate(s)
}
