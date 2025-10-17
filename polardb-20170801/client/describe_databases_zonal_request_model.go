// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDatabasesZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeDatabasesZonalRequest
	GetDBName() *string
	SetMaxResults(v int32) *DescribeDatabasesZonalRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDatabasesZonalRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDatabasesZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDatabasesZonalRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDatabasesZonalRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabasesZonalRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeDatabasesZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDatabasesZonalRequest
	GetResourceOwnerId() *int64
}

type DescribeDatabasesZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDatabasesZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDatabasesZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDatabasesZonalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDatabasesZonalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDatabasesZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDatabasesZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDatabasesZonalRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabasesZonalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabasesZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDatabasesZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDatabasesZonalRequest) SetDBClusterId(v string) *DescribeDatabasesZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetDBName(v string) *DescribeDatabasesZonalRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetMaxResults(v int32) *DescribeDatabasesZonalRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetNextToken(v string) *DescribeDatabasesZonalRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetOwnerAccount(v string) *DescribeDatabasesZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetOwnerId(v int64) *DescribeDatabasesZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetPageNumber(v int32) *DescribeDatabasesZonalRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetPageSize(v int32) *DescribeDatabasesZonalRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetResourceOwnerAccount(v string) *DescribeDatabasesZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) SetResourceOwnerId(v int64) *DescribeDatabasesZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDatabasesZonalRequest) Validate() error {
	return dara.Validate(s)
}
