// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCmsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *GetCmsServiceRequest
	GetProduct() *string
	SetService(v string) *GetCmsServiceRequest
	GetService() *string
}

type GetCmsServiceRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// prometheus
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// prometheus
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
}

func (s GetCmsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCmsServiceRequest) GoString() string {
	return s.String()
}

func (s *GetCmsServiceRequest) GetProduct() *string {
	return s.Product
}

func (s *GetCmsServiceRequest) GetService() *string {
	return s.Service
}

func (s *GetCmsServiceRequest) SetProduct(v string) *GetCmsServiceRequest {
	s.Product = &v
	return s
}

func (s *GetCmsServiceRequest) SetService(v string) *GetCmsServiceRequest {
	s.Service = &v
	return s
}

func (s *GetCmsServiceRequest) Validate() error {
	return dara.Validate(s)
}
