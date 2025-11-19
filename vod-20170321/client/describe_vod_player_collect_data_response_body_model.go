// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVodPlayerCollectDataResponseBodyDataList) *DescribeVodPlayerCollectDataResponseBody
	GetDataList() []*DescribeVodPlayerCollectDataResponseBodyDataList
	SetRequestId(v string) *DescribeVodPlayerCollectDataResponseBody
	GetRequestId() *string
}

type DescribeVodPlayerCollectDataResponseBody struct {
	DataList []*DescribeVodPlayerCollectDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodPlayerCollectDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataResponseBody) GetDataList() []*DescribeVodPlayerCollectDataResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVodPlayerCollectDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerCollectDataResponseBody) SetDataList(v []*DescribeVodPlayerCollectDataResponseBodyDataList) *DescribeVodPlayerCollectDataResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBody) SetRequestId(v string) *DescribeVodPlayerCollectDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodPlayerCollectDataResponseBodyDataList struct {
	// example:
	//
	// Vv
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// 100.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 0.5
	ValueRatio *float64 `json:"ValueRatio,omitempty" xml:"ValueRatio,omitempty"`
	// example:
	//
	// 200.0
	ValueRefer *float64 `json:"ValueRefer,omitempty" xml:"ValueRefer,omitempty"`
}

func (s DescribeVodPlayerCollectDataResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) GetMetric() *string {
	return s.Metric
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) GetValue() *float64 {
	return s.Value
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) GetValueRatio() *float64 {
	return s.ValueRatio
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) GetValueRefer() *float64 {
	return s.ValueRefer
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) SetMetric(v string) *DescribeVodPlayerCollectDataResponseBodyDataList {
	s.Metric = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) SetValue(v float64) *DescribeVodPlayerCollectDataResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) SetValueRatio(v float64) *DescribeVodPlayerCollectDataResponseBodyDataList {
	s.ValueRatio = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) SetValueRefer(v float64) *DescribeVodPlayerCollectDataResponseBodyDataList {
	s.ValueRefer = &v
	return s
}

func (s *DescribeVodPlayerCollectDataResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
