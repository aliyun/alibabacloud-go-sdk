// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDistributionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopDistributionResult
	GetCode() *string
	SetMessage(v string) *StopDistributionResult
	GetMessage() *string
	SetProducts(v []*DistributionProduct) *StopDistributionResult
	GetProducts() []*DistributionProduct
	SetRequestId(v string) *StopDistributionResult
	GetRequestId() *string
}

type StopDistributionResult struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	Products  []*DistributionProduct `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopDistributionResult) String() string {
	return dara.Prettify(s)
}

func (s StopDistributionResult) GoString() string {
	return s.String()
}

func (s *StopDistributionResult) GetCode() *string {
	return s.Code
}

func (s *StopDistributionResult) GetMessage() *string {
	return s.Message
}

func (s *StopDistributionResult) GetProducts() []*DistributionProduct {
	return s.Products
}

func (s *StopDistributionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDistributionResult) SetCode(v string) *StopDistributionResult {
	s.Code = &v
	return s
}

func (s *StopDistributionResult) SetMessage(v string) *StopDistributionResult {
	s.Message = &v
	return s
}

func (s *StopDistributionResult) SetProducts(v []*DistributionProduct) *StopDistributionResult {
	s.Products = v
	return s
}

func (s *StopDistributionResult) SetRequestId(v string) *StopDistributionResult {
	s.RequestId = &v
	return s
}

func (s *StopDistributionResult) Validate() error {
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
