// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *QueryProductInfoRequest
	GetProductId() *string
}

type QueryProductInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s QueryProductInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryProductInfoRequest) GetProductId() *string {
	return s.ProductId
}

func (s *QueryProductInfoRequest) SetProductId(v string) *QueryProductInfoRequest {
	s.ProductId = &v
	return s
}

func (s *QueryProductInfoRequest) Validate() error {
	return dara.Validate(s)
}
