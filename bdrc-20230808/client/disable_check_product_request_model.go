// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *DisableCheckProductRequest
	GetProductType() *string
}

type DisableCheckProductRequest struct {
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DisableCheckProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckProductRequest) GoString() string {
	return s.String()
}

func (s *DisableCheckProductRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DisableCheckProductRequest) SetProductType(v string) *DisableCheckProductRequest {
	s.ProductType = &v
	return s
}

func (s *DisableCheckProductRequest) Validate() error {
	return dara.Validate(s)
}
