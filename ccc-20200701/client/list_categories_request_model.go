// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListCategoriesRequest
	GetCategoryId() *string
	SetInstanceId(v string) *ListCategoriesRequest
	GetInstanceId() *string
	SetType(v string) *ListCategoriesRequest
	GetType() *string
}

type ListCategoriesRequest struct {
	// example:
	//
	// 43c2671b-***-***-86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Ticket
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListCategoriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCategoriesRequest) GetType() *string {
	return s.Type
}

func (s *ListCategoriesRequest) SetCategoryId(v string) *ListCategoriesRequest {
	s.CategoryId = &v
	return s
}

func (s *ListCategoriesRequest) SetInstanceId(v string) *ListCategoriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCategoriesRequest) SetType(v string) *ListCategoriesRequest {
	s.Type = &v
	return s
}

func (s *ListCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
