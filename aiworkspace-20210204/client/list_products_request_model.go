// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCodes(v string) *ListProductsRequest
	GetProductCodes() *string
	SetServiceCodes(v string) *ListProductsRequest
	GetServiceCodes() *string
	SetVerbose(v bool) *ListProductsRequest
	GetVerbose() *bool
}

type ListProductsRequest struct {
	// example:
	//
	// PAI_isolate
	ProductCodes *string `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty"`
	// example:
	//
	// oss
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
	Verbose      *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetProductCodes() *string {
	return s.ProductCodes
}

func (s *ListProductsRequest) GetServiceCodes() *string {
	return s.ServiceCodes
}

func (s *ListProductsRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListProductsRequest) SetProductCodes(v string) *ListProductsRequest {
	s.ProductCodes = &v
	return s
}

func (s *ListProductsRequest) SetServiceCodes(v string) *ListProductsRequest {
	s.ServiceCodes = &v
	return s
}

func (s *ListProductsRequest) SetVerbose(v bool) *ListProductsRequest {
	s.Verbose = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
