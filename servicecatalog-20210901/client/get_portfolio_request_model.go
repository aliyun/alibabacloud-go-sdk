// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *GetPortfolioRequest
	GetPortfolioId() *string
}

type GetPortfolioRequest struct {
	// The ID of the product portfolio.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s GetPortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPortfolioRequest) GoString() string {
	return s.String()
}

func (s *GetPortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetPortfolioRequest) SetPortfolioId(v string) *GetPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *GetPortfolioRequest) Validate() error {
	return dara.Validate(s)
}
