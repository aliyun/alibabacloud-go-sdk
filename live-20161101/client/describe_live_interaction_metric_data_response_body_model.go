// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveInteractionMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribeLiveInteractionMetricDataResponseBodyNodes) *DescribeLiveInteractionMetricDataResponseBody
	GetNodes() []*DescribeLiveInteractionMetricDataResponseBodyNodes
	SetRequestId(v string) *DescribeLiveInteractionMetricDataResponseBody
	GetRequestId() *string
	SetSummaryData(v string) *DescribeLiveInteractionMetricDataResponseBody
	GetSummaryData() *string
}

type DescribeLiveInteractionMetricDataResponseBody struct {
	// The node data.
	Nodes []*DescribeLiveInteractionMetricDataResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A01C98C5-25AE-124A-83FE-514DF5C5BE36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The summary data.
	//
	// example:
	//
	// 2000
	SummaryData *string `json:"SummaryData,omitempty" xml:"SummaryData,omitempty"`
}

func (s DescribeLiveInteractionMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveInteractionMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveInteractionMetricDataResponseBody) GetNodes() []*DescribeLiveInteractionMetricDataResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeLiveInteractionMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveInteractionMetricDataResponseBody) GetSummaryData() *string {
	return s.SummaryData
}

func (s *DescribeLiveInteractionMetricDataResponseBody) SetNodes(v []*DescribeLiveInteractionMetricDataResponseBodyNodes) *DescribeLiveInteractionMetricDataResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponseBody) SetRequestId(v string) *DescribeLiveInteractionMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponseBody) SetSummaryData(v string) *DescribeLiveInteractionMetricDataResponseBody {
	s.SummaryData = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveInteractionMetricDataResponseBodyNodes struct {
	// The time when the metric was queried. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1548670257000
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 66.670000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveInteractionMetricDataResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveInteractionMetricDataResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeLiveInteractionMetricDataResponseBodyNodes) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeLiveInteractionMetricDataResponseBodyNodes) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveInteractionMetricDataResponseBodyNodes) SetTimestamp(v string) *DescribeLiveInteractionMetricDataResponseBodyNodes {
	s.Timestamp = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponseBodyNodes) SetValue(v string) *DescribeLiveInteractionMetricDataResponseBodyNodes {
	s.Value = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}
