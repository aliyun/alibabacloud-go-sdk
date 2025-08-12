// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricDataResponseBody
	GetCode() *string
	SetDatapoints(v string) *DescribeMetricDataResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeMetricDataResponseBody
	GetMessage() *string
	SetPeriod(v string) *DescribeMetricDataResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeMetricDataResponseBody
	GetRequestId() *string
}

type DescribeMetricDataResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data. The value includes the following fields:
	//
	// 	- `timestamp`: the time when the alert was triggered.
	//
	// 	- `userId`: the ID of the user for which the alert was triggered.
	//
	// 	- `instanceId`: the ID of the instance for which the alert was triggered.
	//
	// 	- `Minimum`, `Average`, and `Maximum`: the aggregation methods.
	//
	// example:
	//
	// [{\\"timestamp\\":1618368900000,\\"Average\\":95.8291666666667,\\"Minimum\\":65.48,\\"Maximum\\":100.0},{\\"timestamp\\":1618368960000,\\"Average\\":95.8683333333333,\\"Minimum\\":67.84,\\"Maximum\\":100.0}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Valid values: 15, 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A5F022D-AC7C-460E-94AE-B9E75083D027
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricDataResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeMetricDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricDataResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricDataResponseBody) SetCode(v string) *DescribeMetricDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetDatapoints(v string) *DescribeMetricDataResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetMessage(v string) *DescribeMetricDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetPeriod(v string) *DescribeMetricDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetRequestId(v string) *DescribeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}
