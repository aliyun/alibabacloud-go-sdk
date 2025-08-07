// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPeriod(v string) *DescribeDBClusterMonitorResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeDBClusterMonitorResponseBody
	GetRequestId() *string
}

type DescribeDBClusterMonitorResponseBody struct {
	// The interval at which monitoring data is collected. Unit: seconds.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593AE1C5-B70C-463F-9207-074639******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeDBClusterMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterMonitorResponseBody) SetPeriod(v string) *DescribeDBClusterMonitorResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeDBClusterMonitorResponseBody) SetRequestId(v string) *DescribeDBClusterMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
