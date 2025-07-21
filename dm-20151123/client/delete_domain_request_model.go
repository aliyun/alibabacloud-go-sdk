// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v int32) *DeleteDomainRequest
	GetDomainId() *int32
	SetOwnerId(v int64) *DeleteDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDomainRequest
	GetResourceOwnerId() *int64
}

type DeleteDomainRequest struct {
	// Domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 326***
	DomainId             *int32  `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetDomainId() *int32 {
	return s.DomainId
}

func (s *DeleteDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDomainRequest) SetDomainId(v int32) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) SetOwnerId(v int64) *DeleteDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDomainRequest) SetResourceOwnerAccount(v string) *DeleteDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDomainRequest) SetResourceOwnerId(v int64) *DeleteDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
