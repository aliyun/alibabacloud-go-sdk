// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *UpdatePortfolioResponseBody
	GetPortfolioId() *string
	SetRequestId(v string) *UpdatePortfolioResponseBody
	GetRequestId() *string
}

type UpdatePortfolioResponseBody struct {
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioResponseBody) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *UpdatePortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePortfolioResponseBody) SetPortfolioId(v string) *UpdatePortfolioResponseBody {
	s.PortfolioId = &v
	return s
}

func (s *UpdatePortfolioResponseBody) SetRequestId(v string) *UpdatePortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
