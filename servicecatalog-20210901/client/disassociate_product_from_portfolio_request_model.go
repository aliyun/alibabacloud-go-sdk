// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateProductFromPortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *DisassociateProductFromPortfolioRequest
	GetPortfolioId() *string
	SetProductId(v string) *DisassociateProductFromPortfolioRequest
	GetProductId() *string
}

type DisassociateProductFromPortfolioRequest struct {
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

func (s DisassociateProductFromPortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateProductFromPortfolioRequest) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *DisassociateProductFromPortfolioRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DisassociateProductFromPortfolioRequest) SetPortfolioId(v string) *DisassociateProductFromPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *DisassociateProductFromPortfolioRequest) SetProductId(v string) *DisassociateProductFromPortfolioRequest {
	s.ProductId = &v
	return s
}

func (s *DisassociateProductFromPortfolioRequest) Validate() error {
	return dara.Validate(s)
}
