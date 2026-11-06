// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProductCategoryRequest
	GetDescription() *string
	SetIcon(v *HiMarketIcon) *CreateProductCategoryRequest
	GetIcon() *HiMarketIcon
	SetName(v string) *CreateProductCategoryRequest
	GetName() *string
}

type CreateProductCategoryRequest struct {
	// The description of the product category. Maximum length: 500 characters.
	//
	// example:
	//
	// Product category for AI-related capabilities
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The icon of the product category, which consists of the icon type and icon information fields.
	Icon *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	// The name of the product category. Maximum length: 100 characters.
	//
	// example:
	//
	// ai-capabilities
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProductCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateProductCategoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProductCategoryRequest) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *CreateProductCategoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateProductCategoryRequest) SetDescription(v string) *CreateProductCategoryRequest {
	s.Description = &v
	return s
}

func (s *CreateProductCategoryRequest) SetIcon(v *HiMarketIcon) *CreateProductCategoryRequest {
	s.Icon = v
	return s
}

func (s *CreateProductCategoryRequest) SetName(v string) *CreateProductCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateProductCategoryRequest) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}
