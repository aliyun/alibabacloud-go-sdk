// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeFlowMetricResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeFlowMetricResponseBody
	GetRequestId() *string
}

type DescribeFlowMetricResponseBody struct {
	// example:
	//
	// {\\"instanceId\\": \\"np-4wrye3ishxi47****\\", \\"requestId\\": \\"4F0CD5B6-70D6-5115-A2F7-7EAC3981****\\", \\"dataPoints\\": [{\\"timeStamp\\": 1636510320000, \\"Average\\": 293752.0}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 6857EDCB-631F-5405-BE95-45CBB4C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeFlowMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowMetricResponseBody) SetData(v string) *DescribeFlowMetricResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetRequestId(v string) *DescribeFlowMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
