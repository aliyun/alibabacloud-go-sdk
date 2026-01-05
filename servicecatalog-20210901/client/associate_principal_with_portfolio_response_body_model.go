// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePrincipalWithPortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociatePrincipalWithPortfolioResponseBody
	GetRequestId() *string
}

type AssociatePrincipalWithPortfolioResponseBody struct {
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociatePrincipalWithPortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociatePrincipalWithPortfolioResponseBody) SetRequestId(v string) *AssociatePrincipalWithPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
