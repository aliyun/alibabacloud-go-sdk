// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaDNALibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibRegion(v string) *CreateMediaDNALibRequest
	GetLibRegion() *string
	SetModelType(v string) *CreateMediaDNALibRequest
	GetModelType() *string
	SetOwnerAccount(v string) *CreateMediaDNALibRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CreateMediaDNALibRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *CreateMediaDNALibRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CreateMediaDNALibRequest
	GetResourceOwnerId() *string
}

type CreateMediaDNALibRequest struct {
	// This parameter is required.
	LibRegion            *string `json:"LibRegion,omitempty" xml:"LibRegion,omitempty"`
	ModelType            *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMediaDNALibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaDNALibRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaDNALibRequest) GetLibRegion() *string {
	return s.LibRegion
}

func (s *CreateMediaDNALibRequest) GetModelType() *string {
	return s.ModelType
}

func (s *CreateMediaDNALibRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateMediaDNALibRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateMediaDNALibRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMediaDNALibRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CreateMediaDNALibRequest) SetLibRegion(v string) *CreateMediaDNALibRequest {
	s.LibRegion = &v
	return s
}

func (s *CreateMediaDNALibRequest) SetModelType(v string) *CreateMediaDNALibRequest {
	s.ModelType = &v
	return s
}

func (s *CreateMediaDNALibRequest) SetOwnerAccount(v string) *CreateMediaDNALibRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMediaDNALibRequest) SetOwnerId(v string) *CreateMediaDNALibRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMediaDNALibRequest) SetResourceOwnerAccount(v string) *CreateMediaDNALibRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMediaDNALibRequest) SetResourceOwnerId(v string) *CreateMediaDNALibRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMediaDNALibRequest) Validate() error {
	return dara.Validate(s)
}
