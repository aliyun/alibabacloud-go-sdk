// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricListResponseBody
	GetCode() *string
	SetDatapoints(v string) *DescribeMetricListResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeMetricListResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeMetricListResponseBody
	GetNextToken() *string
	SetPeriod(v string) *DescribeMetricListResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeMetricListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricListResponseBody
	GetSuccess() *bool
}

type DescribeMetricListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data.
	//
	// example:
	//
	// [{"timestamp":1548777660000,"userId":"120886317861****","instanceId":"i-abc","Minimum":9.92,"Average":9.92,"Maximum":9.92}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paging token.
	//
	// example:
	//
	// 15761441850009dd70bb64cff1f0fff6d0b08ffff073be5fb1e785e2b020f7fed9b5e137bd810a6d6cff5ae****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried. Unit: seconds. Valid values: 60, 300, and 900.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3121AE7D-4AFF-4C25-8F1D-C8226EBB1F42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricListResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeMetricListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricListResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricListResponseBody) SetCode(v string) *DescribeMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetDatapoints(v string) *DescribeMetricListResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetMessage(v string) *DescribeMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetNextToken(v string) *DescribeMetricListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetPeriod(v string) *DescribeMetricListResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetRequestId(v string) *DescribeMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetSuccess(v bool) *DescribeMetricListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricListResponseBody) Validate() error {
	return dara.Validate(s)
}
