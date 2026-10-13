// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProductCategoryRequest
	GetDescription() *string
	SetIcon(v *HiMarketIcon) *UpdateProductCategoryRequest
	GetIcon() *HiMarketIcon
	SetName(v string) *UpdateProductCategoryRequest
	GetName() *string
}

type UpdateProductCategoryRequest struct {
	// The description of the product category. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Product category that includes AI-related capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The icon of the product category. The icon consists of two fields: icon type (type) and icon information (value).
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The name of the product category. The name can be up to 100 characters in length.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProductCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductCategoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProductCategoryRequest) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *UpdateProductCategoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProductCategoryRequest) SetDescription(v string) *UpdateProductCategoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductCategoryRequest) SetIcon(v *HiMarketIcon) *UpdateProductCategoryRequest {
	s.Icon = v
	return s
}

func (s *UpdateProductCategoryRequest) SetName(v string) *UpdateProductCategoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateProductCategoryRequest) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}
