// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateListShrinkRequest
	GetDescription() *string
	SetItemsShrink(v string) *CreateListShrinkRequest
	GetItemsShrink() *string
	SetKind(v string) *CreateListShrinkRequest
	GetKind() *string
	SetName(v string) *CreateListShrinkRequest
	GetName() *string
}

type CreateListShrinkRequest struct {
	// The description of the custom list. This parameter provides detailed information about the custom list.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list items. This parameter provides the specific item data for the list.
	//
	// example:
	//
	// a custom list
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// The kind of the custom list. This parameter specifies the type of the custom list.
	//
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the custom list.
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

func (s CreateListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateListShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateListShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *CreateListShrinkRequest) GetKind() *string {
	return s.Kind
}

func (s *CreateListShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateListShrinkRequest) SetDescription(v string) *CreateListShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateListShrinkRequest) SetItemsShrink(v string) *CreateListShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *CreateListShrinkRequest) SetKind(v string) *CreateListShrinkRequest {
	s.Kind = &v
	return s
}

func (s *CreateListShrinkRequest) SetName(v string) *CreateListShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
