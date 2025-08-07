// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDatabasesRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeDatabasesRequest
	GetDBName() *string
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
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// > You cannot specify multiple database names.
	//
	// example:
	//
	// test_db
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
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

func (s *DescribeDatabasesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDatabasesRequest) GetDBName() *string {
	return s.DBName
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

func (s *DescribeDatabasesRequest) SetDBClusterId(v string) *DescribeDatabasesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBName(v string) *DescribeDatabasesRequest {
	s.DBName = &v
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
