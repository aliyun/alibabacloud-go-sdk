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
	// This parameter is required.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
