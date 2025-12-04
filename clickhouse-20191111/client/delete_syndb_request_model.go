// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyndbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *DeleteSyndbRequest
	GetDbClusterId() *string
	SetOwnerAccount(v string) *DeleteSyndbRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSyndbRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSyndbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSyndbRequest
	GetResourceOwnerId() *int64
	SetSynDb(v string) *DeleteSyndbRequest
	GetSynDb() *string
}

type DeleteSyndbRequest struct {
	// The cluster ID.
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
	// The name of the database in the ApsaraDB RDS for MySQL instance. The database is used for data synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
}

func (s DeleteSyndbRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyndbRequest) GoString() string {
	return s.String()
}

func (s *DeleteSyndbRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DeleteSyndbRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSyndbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSyndbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSyndbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSyndbRequest) GetSynDb() *string {
	return s.SynDb
}

func (s *DeleteSyndbRequest) SetDbClusterId(v string) *DeleteSyndbRequest {
	s.DbClusterId = &v
	return s
}

func (s *DeleteSyndbRequest) SetOwnerAccount(v string) *DeleteSyndbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSyndbRequest) SetOwnerId(v int64) *DeleteSyndbRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSyndbRequest) SetResourceOwnerAccount(v string) *DeleteSyndbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSyndbRequest) SetResourceOwnerId(v int64) *DeleteSyndbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSyndbRequest) SetSynDb(v string) *DeleteSyndbRequest {
	s.SynDb = &v
	return s
}

func (s *DeleteSyndbRequest) Validate() error {
	return dara.Validate(s)
}
