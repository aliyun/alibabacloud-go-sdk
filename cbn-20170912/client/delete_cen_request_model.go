// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DeleteCenRequest
	GetCenId() *string
	SetOwnerAccount(v string) *DeleteCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenRequest
	GetResourceOwnerId() *int64
}

type DeleteCenRequest struct {
	// The CEN instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenRequest) SetCenId(v string) *DeleteCenRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenRequest) SetOwnerAccount(v string) *DeleteCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenRequest) SetOwnerId(v int64) *DeleteCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenRequest) SetResourceOwnerAccount(v string) *DeleteCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenRequest) SetResourceOwnerId(v int64) *DeleteCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenRequest) Validate() error {
	return dara.Validate(s)
}
