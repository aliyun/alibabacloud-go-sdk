// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatapoints(v string) *DescribeMonitorDataResponseBody
	GetDatapoints() *string
	SetNextToken(v string) *DescribeMonitorDataResponseBody
	GetNextToken() *string
	SetPeriod(v string) *DescribeMonitorDataResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeMonitorDataResponseBody struct {
	// The monitoring data.
	//
	// example:
	//
	// []
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried. Valid values: 60, 300, and 900. Unit: seconds.
	//
	// > If you set the MetricName request parameter to FLOW_USED, the value of Period is 3600 (one hour).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeMonitorDataResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMonitorDataResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorDataResponseBody) SetDatapoints(v string) *DescribeMonitorDataResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetNextToken(v string) *DescribeMonitorDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetPeriod(v string) *DescribeMonitorDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetRequestId(v string) *DescribeMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) Validate() error {
	return dara.Validate(s)
}
