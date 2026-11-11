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
	// The updated description of the list.
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
	// The updated items in the list. The value is a JSON array string, such as `["1.1.1.1","2.2.2.2"]`.
	//
	// **Full overwrite semantics**: The specified `Items` completely overwrites the original list content instead of appending to it.
	//
	// > ⚠️ **If this parameter is not specified or is left empty, the original list content is cleared**. To retain the original items and append new ones, call `GetList` to retrieve the existing `Items`, concatenate them, and submit them at once.
	//
	// **Element format**: The format depends on the `Kind` specified when the list was created. UpdateList does not support modifying the kind.
	//
	// - Kind = `ip`: Each element must be a valid IP address or CIDR block. If an element is invalid, `WrongValueMatched` is returned.
	//
	// - Other kinds: The element format is subject to conventions, and the number of elements is limited by the tenant quota `NumberItemsPerList`. Only the `ip` kind is not subject to this limit.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom list
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The updated name of the custom list. If this parameter is not specified, the original name is retained.
	//
	// **Naming rules**: Only letters, digits, and underscores are supported. The regular expression is `^\\w{1,64}$`, and the length is 1 to 64 characters.
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
