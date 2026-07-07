// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateListRequest
	GetDescription() *string
	SetId(v int64) *UpdateListRequest
	GetId() *int64
	SetItems(v []*string) *UpdateListRequest
	GetItems() []*string
	SetName(v string) *UpdateListRequest
	GetName() *string
}

type UpdateListRequest struct {
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
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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

func (s UpdateListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateListRequest) GoString() string {
	return s.String()
}

func (s *UpdateListRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateListRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateListRequest) GetItems() []*string {
	return s.Items
}

func (s *UpdateListRequest) GetName() *string {
	return s.Name
}

func (s *UpdateListRequest) SetDescription(v string) *UpdateListRequest {
	s.Description = &v
	return s
}

func (s *UpdateListRequest) SetId(v int64) *UpdateListRequest {
	s.Id = &v
	return s
}

func (s *UpdateListRequest) SetItems(v []*string) *UpdateListRequest {
	s.Items = v
	return s
}

func (s *UpdateListRequest) SetName(v string) *UpdateListRequest {
	s.Name = &v
	return s
}

func (s *UpdateListRequest) Validate() error {
	return dara.Validate(s)
}
