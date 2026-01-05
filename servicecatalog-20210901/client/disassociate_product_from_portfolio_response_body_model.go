// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateProductFromPortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateProductFromPortfolioResponseBody
	GetRequestId() *string
}

type DisassociateProductFromPortfolioResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateProductFromPortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateProductFromPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateProductFromPortfolioResponseBody) SetRequestId(v string) *DisassociateProductFromPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateProductFromPortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
