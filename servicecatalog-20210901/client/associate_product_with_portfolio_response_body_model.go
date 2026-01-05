// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProductWithPortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateProductWithPortfolioResponseBody
	GetRequestId() *string
}

type AssociateProductWithPortfolioResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateProductWithPortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateProductWithPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateProductWithPortfolioResponseBody) SetRequestId(v string) *AssociateProductWithPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateProductWithPortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
