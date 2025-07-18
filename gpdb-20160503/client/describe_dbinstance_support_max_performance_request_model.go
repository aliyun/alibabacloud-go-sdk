// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSupportMaxPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSupportMaxPerformanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBInstanceSupportMaxPerformanceRequest
	GetOwnerId() *int64
}

type DescribeDBInstanceSupportMaxPerformanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDBInstanceSupportMaxPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSupportMaxPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSupportMaxPerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSupportMaxPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceSupportMaxPerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstanceSupportMaxPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceRequest) SetOwnerId(v int64) *DescribeDBInstanceSupportMaxPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
