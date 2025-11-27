// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDatabasesRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeDatabasesRequest
	GetDBName() *string
	SetDBStatus(v string) *DescribeDatabasesRequest
	GetDBStatus() *string
	SetOwnerAccount(v string) *DescribeDatabasesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDatabasesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabasesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeDatabasesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDatabasesRequest
	GetResourceOwnerId() *int64
}

type DescribeDatabasesRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testDB01
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The status of the database. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Creating
	DBStatus     *string `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDatabasesRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabasesRequest) GetDBStatus() *string {
	return s.DBStatus
}

func (s *DescribeDatabasesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDatabasesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabasesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDatabasesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDatabasesRequest) SetDBInstanceId(v string) *DescribeDatabasesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBName(v string) *DescribeDatabasesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBStatus(v string) *DescribeDatabasesRequest {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesRequest) SetOwnerAccount(v string) *DescribeDatabasesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetOwnerId(v int64) *DescribeDatabasesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageNumber(v int32) *DescribeDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageSize(v int32) *DescribeDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerAccount(v string) *DescribeDatabasesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerId(v int64) *DescribeDatabasesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
