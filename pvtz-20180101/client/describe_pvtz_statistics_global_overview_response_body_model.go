// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsGlobalOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePvtzStatisticsGlobalOverviewResponseBodyData) *DescribePvtzStatisticsGlobalOverviewResponseBody
	GetData() *DescribePvtzStatisticsGlobalOverviewResponseBodyData
	SetRequestId(v string) *DescribePvtzStatisticsGlobalOverviewResponseBody
	GetRequestId() *string
}

type DescribePvtzStatisticsGlobalOverviewResponseBody struct {
	Data *DescribePvtzStatisticsGlobalOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePvtzStatisticsGlobalOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsGlobalOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBody) GetData() *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	return s.Data
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBody) SetData(v *DescribePvtzStatisticsGlobalOverviewResponseBodyData) *DescribePvtzStatisticsGlobalOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBody) SetRequestId(v string) *DescribePvtzStatisticsGlobalOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePvtzStatisticsGlobalOverviewResponseBodyData struct {
	// example:
	//
	// 15
	AvgResolveLatency *int64 `json:"AvgResolveLatency,omitempty" xml:"AvgResolveLatency,omitempty"`
	// example:
	//
	// -2
	AvgResolveLatencyTrend *int64 `json:"AvgResolveLatencyTrend,omitempty" xml:"AvgResolveLatencyTrend,omitempty"`
	// example:
	//
	// 98
	AvgSuccessRatio *int64 `json:"AvgSuccessRatio,omitempty" xml:"AvgSuccessRatio,omitempty"`
	// example:
	//
	// 1
	AvgSuccessRatioTrend *int64 `json:"AvgSuccessRatioTrend,omitempty" xml:"AvgSuccessRatioTrend,omitempty"`
	// example:
	//
	// 100000
	TotalResolveCount *int64 `json:"TotalResolveCount,omitempty" xml:"TotalResolveCount,omitempty"`
	// example:
	//
	// 5
	TotalResolveCountTrend *int64 `json:"TotalResolveCountTrend,omitempty" xml:"TotalResolveCountTrend,omitempty"`
}

func (s DescribePvtzStatisticsGlobalOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsGlobalOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetAvgResolveLatency() *int64 {
	return s.AvgResolveLatency
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetAvgResolveLatencyTrend() *int64 {
	return s.AvgResolveLatencyTrend
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetAvgSuccessRatio() *int64 {
	return s.AvgSuccessRatio
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetAvgSuccessRatioTrend() *int64 {
	return s.AvgSuccessRatioTrend
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetTotalResolveCount() *int64 {
	return s.TotalResolveCount
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) GetTotalResolveCountTrend() *int64 {
	return s.TotalResolveCountTrend
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetAvgResolveLatency(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.AvgResolveLatency = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetAvgResolveLatencyTrend(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.AvgResolveLatencyTrend = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetAvgSuccessRatio(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.AvgSuccessRatio = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetAvgSuccessRatioTrend(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.AvgSuccessRatioTrend = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetTotalResolveCount(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.TotalResolveCount = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) SetTotalResolveCountTrend(v int64) *DescribePvtzStatisticsGlobalOverviewResponseBodyData {
	s.TotalResolveCountTrend = &v
	return s
}

func (s *DescribePvtzStatisticsGlobalOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
