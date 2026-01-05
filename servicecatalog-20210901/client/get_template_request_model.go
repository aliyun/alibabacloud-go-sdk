// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductId(v string) *GetTemplateRequest
	GetProductId() *string
	SetProductVersionId(v string) *GetTemplateRequest
	GetProductVersionId() *string
}

type GetTemplateRequest struct {
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetTemplateRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetTemplateRequest) SetProductId(v string) *GetTemplateRequest {
	s.ProductId = &v
	return s
}

func (s *GetTemplateRequest) SetProductVersionId(v string) *GetTemplateRequest {
	s.ProductVersionId = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
