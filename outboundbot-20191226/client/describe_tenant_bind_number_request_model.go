// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTenantBindNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *DescribeTenantBindNumberRequest
	GetNumber() *string
}

type DescribeTenantBindNumberRequest struct {
	// Phone number to query (required)
	//
	// > Query the binding status of this number across all instances. Get the number from the response of ListAllTenantBindNumberBinding.
	//
	// example:
	//
	// 15005059355
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribeTenantBindNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *DescribeTenantBindNumberRequest) SetNumber(v string) *DescribeTenantBindNumberRequest {
	s.Number = &v
	return s
}

func (s *DescribeTenantBindNumberRequest) Validate() error {
	return dara.Validate(s)
}
