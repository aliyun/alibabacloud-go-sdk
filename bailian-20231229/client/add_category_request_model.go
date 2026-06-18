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
	// The name of the category. The name must be 1 to 20 characters long. It can contain Unicode letters, such as English letters and Chinese characters, along with digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// 产品清单
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The type of the category. Valid value:
	//
	// - UNSTRUCTURED: A category.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The ID of the connector instance. You can obtain the ID from the Alibaba Cloud Model Studio console.
	//
	// example:
	//
	// conn_xxxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The ID of the parent category under which the new category is created. If you leave this parameter empty, a top-level category is created.
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
