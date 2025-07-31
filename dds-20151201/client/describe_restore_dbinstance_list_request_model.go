// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreDBInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTimeAfter(v string) *DescribeRestoreDBInstanceListRequest
	GetCreationTimeAfter() *string
	SetDBInstanceId(v string) *DescribeRestoreDBInstanceListRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeRestoreDBInstanceListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRestoreDBInstanceListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRestoreDBInstanceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreDBInstanceListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeRestoreDBInstanceListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRestoreDBInstanceListRequest
	GetResourceOwnerId() *int64
}

type DescribeRestoreDBInstanceListRequest struct {
	// Find instances created after the specified time, formatted as <i>yyyy-MM-dd</i>T<i>HH:00:00</i>Z (UTC time).
	//
	// >
	//
	// > - The time must be on the hour.
	//
	// > - The time cannot be earlier than 7 days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-07-24T14:00:00Z
	CreationTimeAfter *string `json:"CreationTimeAfter,omitempty" xml:"CreationTimeAfter,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp114f14849d****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRestoreDBInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreDBInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreDBInstanceListRequest) GetCreationTimeAfter() *string {
	return s.CreationTimeAfter
}

func (s *DescribeRestoreDBInstanceListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRestoreDBInstanceListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRestoreDBInstanceListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRestoreDBInstanceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreDBInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreDBInstanceListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRestoreDBInstanceListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRestoreDBInstanceListRequest) SetCreationTimeAfter(v string) *DescribeRestoreDBInstanceListRequest {
	s.CreationTimeAfter = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetDBInstanceId(v string) *DescribeRestoreDBInstanceListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetOwnerAccount(v string) *DescribeRestoreDBInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetOwnerId(v int64) *DescribeRestoreDBInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetPageNumber(v int32) *DescribeRestoreDBInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetPageSize(v int32) *DescribeRestoreDBInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetResourceOwnerAccount(v string) *DescribeRestoreDBInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) SetResourceOwnerId(v int64) *DescribeRestoreDBInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRestoreDBInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
