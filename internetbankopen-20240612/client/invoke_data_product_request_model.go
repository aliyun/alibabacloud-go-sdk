// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDataProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InvokeDataProductRequest
	GetClientToken() *string
	SetProductCode(v string) *InvokeDataProductRequest
	GetProductCode() *string
	SetServiceParameters(v string) *InvokeDataProductRequest
	GetServiceParameters() *string
}

type InvokeDataProductRequest struct {
	ClientToken       *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ProductCode       *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ServiceParameters *string `json:"serviceParameters,omitempty" xml:"serviceParameters,omitempty"`
}

func (s InvokeDataProductRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeDataProductRequest) GoString() string {
	return s.String()
}

func (s *InvokeDataProductRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InvokeDataProductRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InvokeDataProductRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *InvokeDataProductRequest) SetClientToken(v string) *InvokeDataProductRequest {
	s.ClientToken = &v
	return s
}

func (s *InvokeDataProductRequest) SetProductCode(v string) *InvokeDataProductRequest {
	s.ProductCode = &v
	return s
}

func (s *InvokeDataProductRequest) SetServiceParameters(v string) *InvokeDataProductRequest {
	s.ServiceParameters = &v
	return s
}

func (s *InvokeDataProductRequest) Validate() error {
	return dara.Validate(s)
}
