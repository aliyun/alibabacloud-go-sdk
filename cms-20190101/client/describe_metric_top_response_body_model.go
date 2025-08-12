// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricTopResponseBody
	GetCode() *string
	SetDatapoints(v string) *DescribeMetricTopResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeMetricTopResponseBody
	GetMessage() *string
	SetPeriod(v string) *DescribeMetricTopResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeMetricTopResponseBody
	GetRequestId() *string
}

type DescribeMetricTopResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data.
	//
	// example:
	//
	// [{\\"order\\":1,\\"timestamp\\":1620287520000,\\"userId\\":\\"120886317861****\\",\\"instanceId\\":\\"i-j6ccf7d5fn335qpo****\\",\\"Average\\":99.92,\\"Minimum\\":99.5,\\"Maximum\\":100.0,\\"_count\\":1.0},{\\"order\\":2,\\"timestamp\\":1620287520000,\\"userId\\":\\"120886317861****\\",\\"instanceId\\":\\"i-0xii2bvf42iqvxbp****\\",\\"Average\\":99.91,\\"Minimum\\":99.0,\\"Maximum\\":100.0,\\"_count\\":1.0}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The statistical period of the monitoring data. Unit: seconds. Valid values: 15, 60, 900, and 3600.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3121AE7D-4AFF-4C25-8F1D-C8226EBB1F42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricTopResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeMetricTopResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricTopResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricTopResponseBody) SetCode(v string) *DescribeMetricTopResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetDatapoints(v string) *DescribeMetricTopResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetMessage(v string) *DescribeMetricTopResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetPeriod(v string) *DescribeMetricTopResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetRequestId(v string) *DescribeMetricTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricTopResponseBody) Validate() error {
	return dara.Validate(s)
}
