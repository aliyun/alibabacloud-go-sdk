// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinlogFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeBinlogFilesRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeBinlogFilesRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeBinlogFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBinlogFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBinlogFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBinlogFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeBinlogFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBinlogFilesRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBinlogFilesRequest
	GetStartTime() *string
}

type DescribeBinlogFilesRequest struct {
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

func (s DescribeBinlogFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBinlogFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBinlogFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBinlogFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBinlogFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBinlogFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBinlogFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBinlogFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBinlogFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBinlogFilesRequest) SetDBInstanceId(v string) *DescribeBinlogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetEndTime(v string) *DescribeBinlogFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetOwnerAccount(v string) *DescribeBinlogFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetOwnerId(v int64) *DescribeBinlogFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetPageNumber(v int32) *DescribeBinlogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetPageSize(v int32) *DescribeBinlogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetResourceOwnerAccount(v string) *DescribeBinlogFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetResourceOwnerId(v int64) *DescribeBinlogFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetStartTime(v string) *DescribeBinlogFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBinlogFilesRequest) Validate() error {
	return dara.Validate(s)
}
