// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDayuMigrateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryDayuMigrateStatusRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryDayuMigrateStatusRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryDayuMigrateStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryDayuMigrateStatusRequest
	GetResourceOwnerId() *int64
}

type QueryDayuMigrateStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryDayuMigrateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDayuMigrateStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDayuMigrateStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryDayuMigrateStatusRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryDayuMigrateStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryDayuMigrateStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryDayuMigrateStatusRequest) SetOwnerId(v int64) *QueryDayuMigrateStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDayuMigrateStatusRequest) SetProdCode(v string) *QueryDayuMigrateStatusRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryDayuMigrateStatusRequest) SetResourceOwnerAccount(v string) *QueryDayuMigrateStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDayuMigrateStatusRequest) SetResourceOwnerId(v int64) *QueryDayuMigrateStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDayuMigrateStatusRequest) Validate() error {
	return dara.Validate(s)
}
