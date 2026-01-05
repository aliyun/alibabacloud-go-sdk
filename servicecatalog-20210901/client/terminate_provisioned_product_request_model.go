// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateProvisionedProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionedProductId(v string) *TerminateProvisionedProductRequest
	GetProvisionedProductId() *string
}

type TerminateProvisionedProductRequest struct {
	// The ID of the product instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
}

func (s TerminateProvisionedProductRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateProvisionedProductRequest) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductRequest) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *TerminateProvisionedProductRequest) SetProvisionedProductId(v string) *TerminateProvisionedProductRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *TerminateProvisionedProductRequest) Validate() error {
	return dara.Validate(s)
}
