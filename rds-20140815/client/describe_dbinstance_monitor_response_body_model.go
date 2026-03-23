// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPeriod(v string) *DescribeDBInstanceMonitorResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceMonitorResponseBody struct {
	Period    *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeDBInstanceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceMonitorResponseBody) SetPeriod(v string) *DescribeDBInstanceMonitorResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
