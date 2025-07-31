// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGranularity(v string) *DescribeDBInstanceMonitorResponseBody
	GetGranularity() *string
	SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceMonitorResponseBody struct {
	// The collection frequency of monitoring data for the instance. Valid value: **5**. Unit: seconds.
	//
	// example:
	//
	// 5
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFD65226-08CC-4C4D-B6A4-CB3C382F67B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponseBody) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeDBInstanceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceMonitorResponseBody) SetGranularity(v string) *DescribeDBInstanceMonitorResponseBody {
	s.Granularity = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
