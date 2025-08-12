// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricLastResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricLastResponseBody
	GetCode() *string
	SetDatapoints(v string) *DescribeMetricLastResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeMetricLastResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeMetricLastResponseBody
	GetNextToken() *string
	SetPeriod(v string) *DescribeMetricLastResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeMetricLastResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricLastResponseBody
	GetSuccess() *bool
}

type DescribeMetricLastResponseBody struct {
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
	// [{"timestamp":1548777660000,"userId":"123456789876****","instanceId":"i-abcdefgh12****","Minimum":93.1,"Average":99.52,"Maximum":100}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// xxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The time interval.
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
	// 021472A6-25E3-4094-8D00-BA4B6A5486C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricLastResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricLastResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeMetricLastResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricLastResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricLastResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricLastResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricLastResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricLastResponseBody) SetCode(v string) *DescribeMetricLastResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetDatapoints(v string) *DescribeMetricLastResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetMessage(v string) *DescribeMetricLastResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetNextToken(v string) *DescribeMetricLastResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetPeriod(v string) *DescribeMetricLastResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetRequestId(v string) *DescribeMetricLastResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetSuccess(v bool) *DescribeMetricLastResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricLastResponseBody) Validate() error {
	return dara.Validate(s)
}
