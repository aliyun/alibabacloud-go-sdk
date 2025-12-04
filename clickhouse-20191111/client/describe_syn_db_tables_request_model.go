// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *DescribeSynDbTablesRequest
	GetDbClusterId() *string
	SetOwnerAccount(v string) *DescribeSynDbTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSynDbTablesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSynDbTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSynDbTablesRequest
	GetResourceOwnerId() *int64
	SetSynDb(v string) *DescribeSynDbTablesRequest
	GetSynDb() *string
}

type DescribeSynDbTablesRequest struct {
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp158i5wvj436****
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
}

func (s DescribeSynDbTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeSynDbTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSynDbTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSynDbTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSynDbTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSynDbTablesRequest) GetSynDb() *string {
	return s.SynDb
}

func (s *DescribeSynDbTablesRequest) SetDbClusterId(v string) *DescribeSynDbTablesRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetOwnerAccount(v string) *DescribeSynDbTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetOwnerId(v int64) *DescribeSynDbTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetResourceOwnerAccount(v string) *DescribeSynDbTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetResourceOwnerId(v int64) *DescribeSynDbTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetSynDb(v string) *DescribeSynDbTablesRequest {
	s.SynDb = &v
	return s
}

func (s *DescribeSynDbTablesRequest) Validate() error {
	return dara.Validate(s)
}
