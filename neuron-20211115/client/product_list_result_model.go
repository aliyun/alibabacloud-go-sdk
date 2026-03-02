// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductListResult interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*Product) *ProductListResult
	GetProducts() []*Product
	SetRequestId(v string) *ProductListResult
	GetRequestId() *string
	SetTotal(v int32) *ProductListResult
	GetTotal() *int32
}

type ProductListResult struct {
	Products  []*Product `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total     *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ProductListResult) String() string {
	return dara.Prettify(s)
}

func (s ProductListResult) GoString() string {
	return s.String()
}

func (s *ProductListResult) GetProducts() []*Product {
	return s.Products
}

func (s *ProductListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductListResult) GetTotal() *int32 {
	return s.Total
}

func (s *ProductListResult) SetProducts(v []*Product) *ProductListResult {
	s.Products = v
	return s
}

func (s *ProductListResult) SetRequestId(v string) *ProductListResult {
	s.RequestId = &v
	return s
}

func (s *ProductListResult) SetTotal(v int32) *ProductListResult {
	s.Total = &v
	return s
}

func (s *ProductListResult) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
