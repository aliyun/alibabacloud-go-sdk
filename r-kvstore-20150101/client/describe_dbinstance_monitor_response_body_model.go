// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterval(v string) *DescribeDBInstanceMonitorResponseBody
	GetInterval() *string
	SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceMonitorResponseBody struct {
	// The collection frequency of the monitoring data. Unit: second.
	//
	// example:
	//
	// 5
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 276F582D-C6B2-519C-A5ED-2A92BB15EC07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponseBody) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDBInstanceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceMonitorResponseBody) SetInterval(v string) *DescribeDBInstanceMonitorResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
