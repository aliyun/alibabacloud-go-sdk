// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociatePrincipalFromPortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociatePrincipalFromPortfolioResponseBody
	GetRequestId() *string
}

type DisassociatePrincipalFromPortfolioResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociatePrincipalFromPortfolioResponseBody) SetRequestId(v string) *DisassociatePrincipalFromPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
