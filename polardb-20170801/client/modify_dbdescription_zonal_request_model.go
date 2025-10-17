// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBDescriptionZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyDBDescriptionZonalRequest
	GetDBClusterId() *string
	SetDBDescription(v string) *ModifyDBDescriptionZonalRequest
	GetDBDescription() *string
	SetDBName(v string) *ModifyDBDescriptionZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *ModifyDBDescriptionZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBDescriptionZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBDescriptionZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBDescriptionZonalRequest
	GetResourceOwnerId() *int64
}

type ModifyDBDescriptionZonalRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdesc
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBDescriptionZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBDescriptionZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBDescriptionZonalRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *ModifyDBDescriptionZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyDBDescriptionZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBDescriptionZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBDescriptionZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBDescriptionZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBDescriptionZonalRequest) SetClientToken(v string) *ModifyDBDescriptionZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetDBClusterId(v string) *ModifyDBDescriptionZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetDBDescription(v string) *ModifyDBDescriptionZonalRequest {
	s.DBDescription = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetDBName(v string) *ModifyDBDescriptionZonalRequest {
	s.DBName = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetOwnerAccount(v string) *ModifyDBDescriptionZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetOwnerId(v int64) *ModifyDBDescriptionZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetResourceOwnerAccount(v string) *ModifyDBDescriptionZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) SetResourceOwnerId(v int64) *ModifyDBDescriptionZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBDescriptionZonalRequest) Validate() error {
	return dara.Validate(s)
}
