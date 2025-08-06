// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayFirstFrameDurationMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) *DescribePlayFirstFrameDurationMetricDataResponseBody
	GetNodes() []*DescribePlayFirstFrameDurationMetricDataResponseBodyNodes
	SetRequestId(v string) *DescribePlayFirstFrameDurationMetricDataResponseBody
	GetRequestId() *string
}

type DescribePlayFirstFrameDurationMetricDataResponseBody struct {
	Nodes     []*DescribePlayFirstFrameDurationMetricDataResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlayFirstFrameDurationMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayFirstFrameDurationMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBody) GetNodes() []*DescribePlayFirstFrameDurationMetricDataResponseBodyNodes {
	return s.Nodes
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBody) SetNodes(v []*DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) *DescribePlayFirstFrameDurationMetricDataResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBody) SetRequestId(v string) *DescribePlayFirstFrameDurationMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayFirstFrameDurationMetricDataResponseBodyNodes struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) GetX() *int64 {
	return s.X
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) GetY() *int64 {
	return s.Y
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) SetX(v int64) *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes {
	s.X = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) SetY(v int64) *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes {
	s.Y = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}
