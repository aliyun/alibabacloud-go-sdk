// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProductWithPortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *AssociateProductWithPortfolioRequest
	GetPortfolioId() *string
	SetProductId(v string) *AssociateProductWithPortfolioRequest
	GetProductId() *string
}

type AssociateProductWithPortfolioRequest struct {
	// The ID of the product portfolio.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s AssociateProductWithPortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateProductWithPortfolioRequest) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *AssociateProductWithPortfolioRequest) GetProductId() *string {
	return s.ProductId
}

func (s *AssociateProductWithPortfolioRequest) SetPortfolioId(v string) *AssociateProductWithPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *AssociateProductWithPortfolioRequest) SetProductId(v string) *AssociateProductWithPortfolioRequest {
	s.ProductId = &v
	return s
}

func (s *AssociateProductWithPortfolioRequest) Validate() error {
	return dara.Validate(s)
}
