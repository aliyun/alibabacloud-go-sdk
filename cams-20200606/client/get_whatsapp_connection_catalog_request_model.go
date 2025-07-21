// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConnectionCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetWhatsappConnectionCatalogRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetWhatsappConnectionCatalogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetWhatsappConnectionCatalogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWhatsappConnectionCatalogRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *GetWhatsappConnectionCatalogRequest
	GetWabaId() *string
}

type GetWhatsappConnectionCatalogRequest struct {
	// The space ID of the user within the ISV account.
	//
	// example:
	//
	// C2020939922929292
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The WABA ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 292939399393
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s GetWhatsappConnectionCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConnectionCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetWhatsappConnectionCatalogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWhatsappConnectionCatalogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWhatsappConnectionCatalogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWhatsappConnectionCatalogRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *GetWhatsappConnectionCatalogRequest) SetCustSpaceId(v string) *GetWhatsappConnectionCatalogRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetOwnerId(v int64) *GetWhatsappConnectionCatalogRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetResourceOwnerAccount(v string) *GetWhatsappConnectionCatalogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetResourceOwnerId(v int64) *GetWhatsappConnectionCatalogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) SetWabaId(v string) *GetWhatsappConnectionCatalogRequest {
	s.WabaId = &v
	return s
}

func (s *GetWhatsappConnectionCatalogRequest) Validate() error {
	return dara.Validate(s)
}
