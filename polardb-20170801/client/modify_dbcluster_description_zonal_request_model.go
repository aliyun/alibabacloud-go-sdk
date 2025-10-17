// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDescriptionZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBClusterDescriptionZonalRequest
	GetClientToken() *string
	SetDBClusterDescription(v string) *ModifyDBClusterDescriptionZonalRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *ModifyDBClusterDescriptionZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterDescriptionZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterDescriptionZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionZonalRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterDescriptionZonalRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterDescriptionZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDescriptionZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterDescriptionZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetClientToken(v string) *ModifyDBClusterDescriptionZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionZonalRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionZonalRequest) Validate() error {
	return dara.Validate(s)
}
