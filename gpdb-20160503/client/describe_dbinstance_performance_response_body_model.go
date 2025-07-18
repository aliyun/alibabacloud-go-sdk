// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetEndTime() *string
	SetEngine(v string) *DescribeDBInstancePerformanceResponseBody
	GetEngine() *string
	SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponseBody
	GetPerformanceKeys() []*string
	SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBInstancePerformanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2022-07-09T03:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The queried performance metrics.
	PerformanceKeys []*string `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5E85244A-AB47-46A3-A3AD-5F307DCB407E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2022-07-08T03:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancePerformanceResponseBody) GetPerformanceKeys() []*string {
	return s.PerformanceKeys
}

func (s *DescribeDBInstancePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}
