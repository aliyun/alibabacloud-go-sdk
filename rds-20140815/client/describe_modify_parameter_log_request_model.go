// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeModifyParameterLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeModifyParameterLogRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeModifyParameterLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeModifyParameterLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeModifyParameterLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeModifyParameterLogRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeModifyParameterLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeModifyParameterLogRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeModifyParameterLogRequest
	GetStartTime() *string
}

type DescribeModifyParameterLogRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModifyParameterLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModifyParameterLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModifyParameterLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeModifyParameterLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeModifyParameterLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModifyParameterLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModifyParameterLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeModifyParameterLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeModifyParameterLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModifyParameterLogRequest) SetDBInstanceId(v string) *DescribeModifyParameterLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetOwnerAccount(v string) *DescribeModifyParameterLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetOwnerId(v int64) *DescribeModifyParameterLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetPageNumber(v int32) *DescribeModifyParameterLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetPageSize(v int32) *DescribeModifyParameterLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetResourceOwnerAccount(v string) *DescribeModifyParameterLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetResourceOwnerId(v int64) *DescribeModifyParameterLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) Validate() error {
	return dara.Validate(s)
}
