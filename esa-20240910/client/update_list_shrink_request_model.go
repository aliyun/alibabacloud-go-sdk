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
	// The ID of the custom list. You can obtain the ID by calling the [ListLists](https://help.aliyun.com/document_detail/2850217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new list content. The value is a JSON array string, for example, `["1.1.1.1","2.2.2.2"]`.
	//
	// **Full overwrite semantics**: The specified `Items` value completely overwrites the existing list content instead of appending to it.
	//
	// > ⚠️ **If this parameter is not specified or is set to an empty value, the existing list content is cleared**. To retain existing items and append new ones, call `GetList` to retrieve the current `Items`, merge them, and then submit the combined list.
	//
	// **Element format**: The format depends on the `Kind` value specified when the list was created. UpdateList does not support modifying Kind.
	//
	// - Kind = `ip`: Each element must be a valid IP address or CIDR block. If an element is invalid, `WrongValueMatched` is returned.
	//
	// - Other Kind values: The element format is subject to the relevant specifications. The number of elements is limited by the tenant quota `NumberItemsPerList`. This limit does not apply to the `ip` Kind.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom list
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// The new name of the custom list. If this parameter is not specified, the original name is retained.
	//
	// **Naming rules**: Only letters, digits, and underscores are supported (`^\\w{1,64}$`). The name must be 1 to 64 characters in length.
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
