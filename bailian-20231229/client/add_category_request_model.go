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
	// The category name. The name must be 1 to 20 characters in length and can contain characters in the Unicode letter category, including letters, Chinese characters, and digits. The name can also contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ProductList
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The category type. Valid values:
	//
	// - UNSTRUCTURED: category.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The instance ID of the data connector (Connector). Obtain this value from the Model Studio console.
	//
	// example:
	//
	// conn_xxxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The ID of the parent category. The new category is created under this parent category. If this parameter is left empty, a first-level category is created.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3xxxxxxxx
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
