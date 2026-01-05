// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortfolioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePortfolioResponseBody
	GetRequestId() *string
}

type DeletePortfolioResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePortfolioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortfolioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePortfolioResponseBody) SetRequestId(v string) *DeletePortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePortfolioResponseBody) Validate() error {
	return dara.Validate(s)
}
