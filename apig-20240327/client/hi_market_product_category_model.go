// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketProductCategory interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *HiMarketProductCategory
	GetCategoryId() *string
	SetDescription(v string) *HiMarketProductCategory
	GetDescription() *string
	SetIcon(v *HiMarketIcon) *HiMarketProductCategory
	GetIcon() *HiMarketIcon
	SetName(v string) *HiMarketProductCategory
	GetName() *string
}

type HiMarketProductCategory struct {
	CategoryId  *string       `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	Description *string       `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *HiMarketIcon `json:"icon,omitempty" xml:"icon,omitempty"`
	Name        *string       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HiMarketProductCategory) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductCategory) GoString() string {
	return s.String()
}

func (s *HiMarketProductCategory) GetCategoryId() *string {
	return s.CategoryId
}

func (s *HiMarketProductCategory) GetDescription() *string {
	return s.Description
}

func (s *HiMarketProductCategory) GetIcon() *HiMarketIcon {
	return s.Icon
}

func (s *HiMarketProductCategory) GetName() *string {
	return s.Name
}

func (s *HiMarketProductCategory) SetCategoryId(v string) *HiMarketProductCategory {
	s.CategoryId = &v
	return s
}

func (s *HiMarketProductCategory) SetDescription(v string) *HiMarketProductCategory {
	s.Description = &v
	return s
}

func (s *HiMarketProductCategory) SetIcon(v *HiMarketIcon) *HiMarketProductCategory {
	s.Icon = v
	return s
}

func (s *HiMarketProductCategory) SetName(v string) *HiMarketProductCategory {
	s.Name = &v
	return s
}

func (s *HiMarketProductCategory) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}
