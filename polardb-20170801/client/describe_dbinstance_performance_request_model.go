// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBInstancePerformanceRequest
	GetKey() *string
	SetOwnerAccount(v string) *DescribeDBInstancePerformanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancePerformanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstancePerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBInstancePerformanceRequest
	GetStartTime() *string
}

type DescribeDBInstancePerformanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pi-*************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-15T17:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// PolarDBCPUForPCU,PolarDBPCU,PolarDBMemoryForPCU,PolarDBQPSTPS,PolarDBConnections
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancePerformanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancePerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
