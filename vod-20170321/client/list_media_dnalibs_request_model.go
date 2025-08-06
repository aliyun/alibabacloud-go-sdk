// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNALibsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibRegion(v string) *ListMediaDNALibsRequest
	GetLibRegion() *string
	SetOwnerAccount(v string) *ListMediaDNALibsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListMediaDNALibsRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListMediaDNALibsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListMediaDNALibsRequest
	GetResourceOwnerId() *string
}

type ListMediaDNALibsRequest struct {
	// This parameter is required.
	LibRegion            *string `json:"LibRegion,omitempty" xml:"LibRegion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMediaDNALibsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNALibsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaDNALibsRequest) GetLibRegion() *string {
	return s.LibRegion
}

func (s *ListMediaDNALibsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListMediaDNALibsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListMediaDNALibsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMediaDNALibsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListMediaDNALibsRequest) SetLibRegion(v string) *ListMediaDNALibsRequest {
	s.LibRegion = &v
	return s
}

func (s *ListMediaDNALibsRequest) SetOwnerAccount(v string) *ListMediaDNALibsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListMediaDNALibsRequest) SetOwnerId(v string) *ListMediaDNALibsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaDNALibsRequest) SetResourceOwnerAccount(v string) *ListMediaDNALibsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaDNALibsRequest) SetResourceOwnerId(v string) *ListMediaDNALibsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMediaDNALibsRequest) Validate() error {
	return dara.Validate(s)
}
