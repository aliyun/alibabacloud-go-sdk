// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribePlayMetricDataResponseBodyNodes) *DescribePlayMetricDataResponseBody
	GetNodes() []*DescribePlayMetricDataResponseBodyNodes
	SetRequestId(v string) *DescribePlayMetricDataResponseBody
	GetRequestId() *string
	SetSummaryData(v string) *DescribePlayMetricDataResponseBody
	GetSummaryData() *string
}

type DescribePlayMetricDataResponseBody struct {
	Nodes       []*DescribePlayMetricDataResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SummaryData *string                                    `json:"SummaryData,omitempty" xml:"SummaryData,omitempty"`
}

func (s DescribePlayMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricDataResponseBody) GetNodes() []*DescribePlayMetricDataResponseBodyNodes {
	return s.Nodes
}

func (s *DescribePlayMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayMetricDataResponseBody) GetSummaryData() *string {
	return s.SummaryData
}

func (s *DescribePlayMetricDataResponseBody) SetNodes(v []*DescribePlayMetricDataResponseBodyNodes) *DescribePlayMetricDataResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribePlayMetricDataResponseBody) SetRequestId(v string) *DescribePlayMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayMetricDataResponseBody) SetSummaryData(v string) *DescribePlayMetricDataResponseBody {
	s.SummaryData = &v
	return s
}

func (s *DescribePlayMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayMetricDataResponseBodyNodes struct {
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribePlayMetricDataResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricDataResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricDataResponseBodyNodes) GetX() *string {
	return s.X
}

func (s *DescribePlayMetricDataResponseBodyNodes) GetY() *string {
	return s.Y
}

func (s *DescribePlayMetricDataResponseBodyNodes) SetX(v string) *DescribePlayMetricDataResponseBodyNodes {
	s.X = &v
	return s
}

func (s *DescribePlayMetricDataResponseBodyNodes) SetY(v string) *DescribePlayMetricDataResponseBodyNodes {
	s.Y = &v
	return s
}

func (s *DescribePlayMetricDataResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}
