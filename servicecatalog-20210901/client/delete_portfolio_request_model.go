// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *DeletePortfolioRequest
	GetPortfolioId() *string
}

type DeletePortfolioRequest struct {
	// The ID of the product portfolio.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s DeletePortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePortfolioRequest) GoString() string {
	return s.String()
}

func (s *DeletePortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *DeletePortfolioRequest) SetPortfolioId(v string) *DeletePortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *DeletePortfolioRequest) Validate() error {
	return dara.Validate(s)
}
