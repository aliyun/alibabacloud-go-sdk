// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrincipalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *ListPrincipalsRequest
	GetPortfolioId() *string
}

type ListPrincipalsRequest struct {
	// The ID of the product portfolio.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s ListPrincipalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrincipalsRequest) GoString() string {
	return s.String()
}

func (s *ListPrincipalsRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListPrincipalsRequest) SetPortfolioId(v string) *ListPrincipalsRequest {
	s.PortfolioId = &v
	return s
}

func (s *ListPrincipalsRequest) Validate() error {
	return dara.Validate(s)
}
