// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateListShrinkRequest
	GetDescription() *string
	SetId(v int64) *UpdateListShrinkRequest
	GetId() *int64
	SetItemsShrink(v string) *UpdateListShrinkRequest
	GetItemsShrink() *string
	SetName(v string) *UpdateListShrinkRequest
	GetName() *string
}

type UpdateListShrinkRequest struct {
	// The new description of the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom list, which can be obtained by calling the [ListLists](https://help.aliyun.com/document_detail/2850217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The items in the updated list. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom list
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// The new name of the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateListShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateListShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateListShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *UpdateListShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateListShrinkRequest) SetDescription(v string) *UpdateListShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateListShrinkRequest) SetId(v int64) *UpdateListShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateListShrinkRequest) SetItemsShrink(v string) *UpdateListShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *UpdateListShrinkRequest) SetName(v string) *UpdateListShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
