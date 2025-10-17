// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckDBNameZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *CheckDBNameZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CheckDBNameZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckDBNameZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckDBNameZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckDBNameZonalRequest
	GetResourceOwnerId() *int64
}

type CheckDBNameZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckDBNameZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameZonalRequest) GoString() string {
	return s.String()
}

func (s *CheckDBNameZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckDBNameZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *CheckDBNameZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckDBNameZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckDBNameZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckDBNameZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckDBNameZonalRequest) SetDBClusterId(v string) *CheckDBNameZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckDBNameZonalRequest) SetDBName(v string) *CheckDBNameZonalRequest {
	s.DBName = &v
	return s
}

func (s *CheckDBNameZonalRequest) SetOwnerAccount(v string) *CheckDBNameZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckDBNameZonalRequest) SetOwnerId(v int64) *CheckDBNameZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDBNameZonalRequest) SetResourceOwnerAccount(v string) *CheckDBNameZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDBNameZonalRequest) SetResourceOwnerId(v int64) *CheckDBNameZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDBNameZonalRequest) Validate() error {
	return dara.Validate(s)
}
