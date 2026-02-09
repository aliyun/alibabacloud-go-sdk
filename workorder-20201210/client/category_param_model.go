// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryParam interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *CategoryParam
	GetCategoryId() *int64
	SetCategoryName(v string) *CategoryParam
	GetCategoryName() *string
	SetIsSelectFromDialog(v bool) *CategoryParam
	GetIsSelectFromDialog() *bool
	SetProductId(v int64) *CategoryParam
	GetProductId() *int64
	SetProductName(v string) *CategoryParam
	GetProductName() *string
}

type CategoryParam struct {
	CategoryId         *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName       *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	IsSelectFromDialog *bool   `json:"IsSelectFromDialog,omitempty" xml:"IsSelectFromDialog,omitempty"`
	ProductId          *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName        *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s CategoryParam) String() string {
	return dara.Prettify(s)
}

func (s CategoryParam) GoString() string {
	return s.String()
}

func (s *CategoryParam) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CategoryParam) GetCategoryName() *string {
	return s.CategoryName
}

func (s *CategoryParam) GetIsSelectFromDialog() *bool {
	return s.IsSelectFromDialog
}

func (s *CategoryParam) GetProductId() *int64 {
	return s.ProductId
}

func (s *CategoryParam) GetProductName() *string {
	return s.ProductName
}

func (s *CategoryParam) SetCategoryId(v int64) *CategoryParam {
	s.CategoryId = &v
	return s
}

func (s *CategoryParam) SetCategoryName(v string) *CategoryParam {
	s.CategoryName = &v
	return s
}

func (s *CategoryParam) SetIsSelectFromDialog(v bool) *CategoryParam {
	s.IsSelectFromDialog = &v
	return s
}

func (s *CategoryParam) SetProductId(v int64) *CategoryParam {
	s.ProductId = &v
	return s
}

func (s *CategoryParam) SetProductName(v string) *CategoryParam {
	s.ProductName = &v
	return s
}

func (s *CategoryParam) Validate() error {
	return dara.Validate(s)
}
