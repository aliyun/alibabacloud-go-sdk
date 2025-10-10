// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyProductRequest
	GetDescription() *string
	SetName(v string) *ModifyProductRequest
	GetName() *string
	SetProductId(v string) *ModifyProductRequest
	GetProductId() *string
}

type ModifyProductRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ModifyProductRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProductRequest) GoString() string {
	return s.String()
}

func (s *ModifyProductRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProductRequest) GetName() *string {
	return s.Name
}

func (s *ModifyProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ModifyProductRequest) SetDescription(v string) *ModifyProductRequest {
	s.Description = &v
	return s
}

func (s *ModifyProductRequest) SetName(v string) *ModifyProductRequest {
	s.Name = &v
	return s
}

func (s *ModifyProductRequest) SetProductId(v string) *ModifyProductRequest {
	s.ProductId = &v
	return s
}

func (s *ModifyProductRequest) Validate() error {
	return dara.Validate(s)
}
