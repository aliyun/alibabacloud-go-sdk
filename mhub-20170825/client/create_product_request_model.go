// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProductRequest
	GetDescription() *string
	SetName(v string) *CreateProductRequest
	GetName() *string
}

type CreateProductRequest struct {
	// example:
	//
	// AAA
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProductRequest) GetName() *string {
	return s.Name
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetName(v string) *CreateProductRequest {
	s.Name = &v
	return s
}

func (s *CreateProductRequest) Validate() error {
	return dara.Validate(s)
}
