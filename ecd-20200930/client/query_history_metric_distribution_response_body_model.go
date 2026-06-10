// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryMetricDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistributionList(v []*QueryHistoryMetricDistributionResponseBodyDistributionList) *QueryHistoryMetricDistributionResponseBody
	GetDistributionList() []*QueryHistoryMetricDistributionResponseBodyDistributionList
	SetRequestId(v string) *QueryHistoryMetricDistributionResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryHistoryMetricDistributionResponseBody
	GetTotalCount() *int64
}

type QueryHistoryMetricDistributionResponseBody struct {
	// An array of objects, each representing the distribution data for a specific value range.
	DistributionList []*QueryHistoryMetricDistributionResponseBodyDistributionList `json:"DistributionList,omitempty" xml:"DistributionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2F2BF549-CBD9-1FED-9ABB-086B62D7B293
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data points that match the query.
	//
	// example:
	//
	// 94
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHistoryMetricDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryMetricDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoryMetricDistributionResponseBody) GetDistributionList() []*QueryHistoryMetricDistributionResponseBodyDistributionList {
	return s.DistributionList
}

func (s *QueryHistoryMetricDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoryMetricDistributionResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryHistoryMetricDistributionResponseBody) SetDistributionList(v []*QueryHistoryMetricDistributionResponseBodyDistributionList) *QueryHistoryMetricDistributionResponseBody {
	s.DistributionList = v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBody) SetRequestId(v string) *QueryHistoryMetricDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBody) SetTotalCount(v int64) *QueryHistoryMetricDistributionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBody) Validate() error {
	if s.DistributionList != nil {
		for _, item := range s.DistributionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoryMetricDistributionResponseBodyDistributionList struct {
	// The number of data points in this value range.
	//
	// example:
	//
	// 40
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The range label that you defined in the request.
	//
	// example:
	//
	// 2F2BF549-CBD9-1FED-9ABB-086B62D7B293
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum value of the value range.
	//
	// example:
	//
	// 20
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum value of the value range.
	//
	// example:
	//
	// 0
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s QueryHistoryMetricDistributionResponseBodyDistributionList) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryMetricDistributionResponseBodyDistributionList) GoString() string {
	return s.String()
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) GetCount() *int32 {
	return s.Count
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) GetLabel() *string {
	return s.Label
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) GetMax() *float32 {
	return s.Max
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) GetMin() *float32 {
	return s.Min
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) SetCount(v int32) *QueryHistoryMetricDistributionResponseBodyDistributionList {
	s.Count = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) SetLabel(v string) *QueryHistoryMetricDistributionResponseBodyDistributionList {
	s.Label = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) SetMax(v float32) *QueryHistoryMetricDistributionResponseBodyDistributionList {
	s.Max = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) SetMin(v float32) *QueryHistoryMetricDistributionResponseBodyDistributionList {
	s.Min = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponseBodyDistributionList) Validate() error {
	return dara.Validate(s)
}
