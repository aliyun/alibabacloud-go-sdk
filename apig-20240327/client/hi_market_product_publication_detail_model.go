// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketProductPublicationDetail interface {
	dara.Model
	String() string
	GoString() string
	SetPortalId(v string) *HiMarketProductPublicationDetail
	GetPortalId() *string
	SetPortalName(v string) *HiMarketProductPublicationDetail
	GetPortalName() *string
	SetProductId(v string) *HiMarketProductPublicationDetail
	GetProductId() *string
	SetProductName(v string) *HiMarketProductPublicationDetail
	GetProductName() *string
	SetProductType(v string) *HiMarketProductPublicationDetail
	GetProductType() *string
	SetPublicationId(v string) *HiMarketProductPublicationDetail
	GetPublicationId() *string
}

type HiMarketProductPublicationDetail struct {
	// The unique ID of the portal.
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The name of the portal.
	PortalName *string `json:"portalName,omitempty" xml:"portalName,omitempty"`
	// The unique ID of the product.
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The name of the product.
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// The type of the product.
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The unique ID of the product publication.
	PublicationId *string `json:"publicationId,omitempty" xml:"publicationId,omitempty"`
}

func (s HiMarketProductPublicationDetail) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductPublicationDetail) GoString() string {
	return s.String()
}

func (s *HiMarketProductPublicationDetail) GetPortalId() *string {
	return s.PortalId
}

func (s *HiMarketProductPublicationDetail) GetPortalName() *string {
	return s.PortalName
}

func (s *HiMarketProductPublicationDetail) GetProductId() *string {
	return s.ProductId
}

func (s *HiMarketProductPublicationDetail) GetProductName() *string {
	return s.ProductName
}

func (s *HiMarketProductPublicationDetail) GetProductType() *string {
	return s.ProductType
}

func (s *HiMarketProductPublicationDetail) GetPublicationId() *string {
	return s.PublicationId
}

func (s *HiMarketProductPublicationDetail) SetPortalId(v string) *HiMarketProductPublicationDetail {
	s.PortalId = &v
	return s
}

func (s *HiMarketProductPublicationDetail) SetPortalName(v string) *HiMarketProductPublicationDetail {
	s.PortalName = &v
	return s
}

func (s *HiMarketProductPublicationDetail) SetProductId(v string) *HiMarketProductPublicationDetail {
	s.ProductId = &v
	return s
}

func (s *HiMarketProductPublicationDetail) SetProductName(v string) *HiMarketProductPublicationDetail {
	s.ProductName = &v
	return s
}

func (s *HiMarketProductPublicationDetail) SetProductType(v string) *HiMarketProductPublicationDetail {
	s.ProductType = &v
	return s
}

func (s *HiMarketProductPublicationDetail) SetPublicationId(v string) *HiMarketProductPublicationDetail {
	s.PublicationId = &v
	return s
}

func (s *HiMarketProductPublicationDetail) Validate() error {
	return dara.Validate(s)
}
