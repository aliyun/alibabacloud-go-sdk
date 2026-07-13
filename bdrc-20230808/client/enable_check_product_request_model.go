// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckProductRequest interface {
  dara.Model
  String() string
  GoString() string
  SetProductType(v string) *EnableCheckProductRequest
  GetProductType() *string 
}

type EnableCheckProductRequest struct {
  // The product type.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ecs
  ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s EnableCheckProductRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckProductRequest) GoString() string {
  return s.String()
}

func (s *EnableCheckProductRequest) GetProductType() *string  {
  return s.ProductType
}

func (s *EnableCheckProductRequest) SetProductType(v string) *EnableCheckProductRequest {
  s.ProductType = &v
  return s
}

func (s *EnableCheckProductRequest) Validate() error {
  return dara.Validate(s)
}

