// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionedProductId(v string) *GetProvisionedProductRequest
	GetProvisionedProductId() *string
}

type GetProvisionedProductRequest struct {
	// The ID of the product instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
}

func (s GetProvisionedProductRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductRequest) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *GetProvisionedProductRequest) SetProvisionedProductId(v string) *GetProvisionedProductRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *GetProvisionedProductRequest) Validate() error {
	return dara.Validate(s)
}
