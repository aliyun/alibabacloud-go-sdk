// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateListRequest
	GetDescription() *string
	SetItems(v []*string) *CreateListRequest
	GetItems() []*string
	SetKind(v string) *CreateListRequest
	GetKind() *string
	SetName(v string) *CreateListRequest
	GetName() *string
}

type CreateListRequest struct {
	// The description of the list that you want to create.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The items in the list that you want to create.
	//
	// example:
	//
	// a custom list
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the list that you want to create.
	//
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the list that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateListRequest) GoString() string {
	return s.String()
}

func (s *CreateListRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateListRequest) GetItems() []*string {
	return s.Items
}

func (s *CreateListRequest) GetKind() *string {
	return s.Kind
}

func (s *CreateListRequest) GetName() *string {
	return s.Name
}

func (s *CreateListRequest) SetDescription(v string) *CreateListRequest {
	s.Description = &v
	return s
}

func (s *CreateListRequest) SetItems(v []*string) *CreateListRequest {
	s.Items = v
	return s
}

func (s *CreateListRequest) SetKind(v string) *CreateListRequest {
	s.Kind = &v
	return s
}

func (s *CreateListRequest) SetName(v string) *CreateListRequest {
	s.Name = &v
	return s
}

func (s *CreateListRequest) Validate() error {
	return dara.Validate(s)
}
