// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *AddCategoryRequest
	GetCategoryName() *string
	SetCategoryType(v string) *AddCategoryRequest
	GetCategoryType() *string
	SetConnectorId(v string) *AddCategoryRequest
	GetConnectorId() *string
	SetParentCategoryId(v string) *AddCategoryRequest
	GetParentCategoryId() *string
}

type AddCategoryRequest struct {
	// This parameter is required.
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	ConnectorId  *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s AddCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddCategoryRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *AddCategoryRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddCategoryRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AddCategoryRequest) GetParentCategoryId() *string {
	return s.ParentCategoryId
}

func (s *AddCategoryRequest) SetCategoryName(v string) *AddCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *AddCategoryRequest) SetCategoryType(v string) *AddCategoryRequest {
	s.CategoryType = &v
	return s
}

func (s *AddCategoryRequest) SetConnectorId(v string) *AddCategoryRequest {
	s.ConnectorId = &v
	return s
}

func (s *AddCategoryRequest) SetParentCategoryId(v string) *AddCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *AddCategoryRequest) Validate() error {
	return dara.Validate(s)
}
