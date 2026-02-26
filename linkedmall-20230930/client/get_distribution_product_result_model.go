// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDistributionProductResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDistributionProductResult
	GetCode() *string
	SetMessage(v string) *GetDistributionProductResult
	GetMessage() *string
	SetProducts(v []*DistributionProduct) *GetDistributionProductResult
	GetProducts() []*DistributionProduct
	SetRequestId(v string) *GetDistributionProductResult
	GetRequestId() *string
}

type GetDistributionProductResult struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	Products  []*DistributionProduct `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDistributionProductResult) String() string {
	return dara.Prettify(s)
}

func (s GetDistributionProductResult) GoString() string {
	return s.String()
}

func (s *GetDistributionProductResult) GetCode() *string {
	return s.Code
}

func (s *GetDistributionProductResult) GetMessage() *string {
	return s.Message
}

func (s *GetDistributionProductResult) GetProducts() []*DistributionProduct {
	return s.Products
}

func (s *GetDistributionProductResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDistributionProductResult) SetCode(v string) *GetDistributionProductResult {
	s.Code = &v
	return s
}

func (s *GetDistributionProductResult) SetMessage(v string) *GetDistributionProductResult {
	s.Message = &v
	return s
}

func (s *GetDistributionProductResult) SetProducts(v []*DistributionProduct) *GetDistributionProductResult {
	s.Products = v
	return s
}

func (s *GetDistributionProductResult) SetRequestId(v string) *GetDistributionProductResult {
	s.RequestId = &v
	return s
}

func (s *GetDistributionProductResult) Validate() error {
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
