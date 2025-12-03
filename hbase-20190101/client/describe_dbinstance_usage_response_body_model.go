// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBInstanceUsageResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeDBInstanceUsageResponseBody
	GetResult() *string
}

type DescribeDBInstanceUsageResponseBody struct {
	// example:
	//
	// A2D841CE-D066-53E8-B9AC-3731DCC85397
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"disk_usage_record\\":{\\"disk_used\\":\\"0.9GB\\",\\"disk_total\\":\\"1156.1GB\\",\\"usage_rate\\":\\"1%\\"}}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeDBInstanceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceUsageResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeDBInstanceUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceUsageResponseBody) SetResult(v string) *DescribeDBInstanceUsageResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeDBInstanceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
