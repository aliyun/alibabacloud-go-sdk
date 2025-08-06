// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataDemoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVodPlayerCollectDataDemoResponseBodyDataList) *DescribeVodPlayerCollectDataDemoResponseBody
	GetDataList() []*DescribeVodPlayerCollectDataDemoResponseBodyDataList
	SetRequestId(v string) *DescribeVodPlayerCollectDataDemoResponseBody
	GetRequestId() *string
}

type DescribeVodPlayerCollectDataDemoResponseBody struct {
	DataList []*DescribeVodPlayerCollectDataDemoResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodPlayerCollectDataDemoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataDemoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataDemoResponseBody) GetDataList() []*DescribeVodPlayerCollectDataDemoResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVodPlayerCollectDataDemoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerCollectDataDemoResponseBody) SetDataList(v []*DescribeVodPlayerCollectDataDemoResponseBodyDataList) *DescribeVodPlayerCollectDataDemoResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBody) SetRequestId(v string) *DescribeVodPlayerCollectDataDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodPlayerCollectDataDemoResponseBodyDataList struct {
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

func (s DescribeVodPlayerCollectDataDemoResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataDemoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) GetMetric() *string {
	return s.Metric
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) GetValue() *float64 {
	return s.Value
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) GetValueRatio() *float64 {
	return s.ValueRatio
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) GetValueRefer() *float64 {
	return s.ValueRefer
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) SetMetric(v string) *DescribeVodPlayerCollectDataDemoResponseBodyDataList {
	s.Metric = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) SetValue(v float64) *DescribeVodPlayerCollectDataDemoResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) SetValueRatio(v float64) *DescribeVodPlayerCollectDataDemoResponseBodyDataList {
	s.ValueRatio = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) SetValueRefer(v float64) *DescribeVodPlayerCollectDataDemoResponseBodyDataList {
	s.ValueRefer = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
