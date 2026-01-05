// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *CreatePortfolioResponseBody
	GetPortfolioId() *string
	SetRequestId(v string) *CreatePortfolioResponseBody
	GetRequestId() *string
}

type CreatePortfolioResponseBody struct {
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortfolioResponseBody) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *CreatePortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePortfolioResponseBody) SetPortfolioId(v string) *CreatePortfolioResponseBody {
	s.PortfolioId = &v
	return s
}

func (s *CreatePortfolioResponseBody) SetRequestId(v string) *CreatePortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
