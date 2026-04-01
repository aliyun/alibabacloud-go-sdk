// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeErrorLogsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeErrorLogsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeErrorLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeErrorLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeErrorLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeErrorLogsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeErrorLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeErrorLogsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeErrorLogsRequest
	GetStartTime() *string
}

type DescribeErrorLogsRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span between the start time and the end time must be less than 31 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-05-30T20:10Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-05-01T20:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeErrorLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeErrorLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeErrorLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeErrorLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeErrorLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeErrorLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeErrorLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeErrorLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeErrorLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeErrorLogsRequest) SetDBInstanceId(v string) *DescribeErrorLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetEndTime(v string) *DescribeErrorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetOwnerAccount(v string) *DescribeErrorLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetOwnerId(v int64) *DescribeErrorLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetPageNumber(v int32) *DescribeErrorLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetPageSize(v int32) *DescribeErrorLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetResourceOwnerAccount(v string) *DescribeErrorLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetResourceOwnerId(v int64) *DescribeErrorLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeErrorLogsRequest) SetStartTime(v string) *DescribeErrorLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeErrorLogsRequest) Validate() error {
	return dara.Validate(s)
}
