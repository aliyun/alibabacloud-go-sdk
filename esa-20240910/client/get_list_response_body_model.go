// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetListResponseBody
	GetDescription() *string
	SetId(v int64) *GetListResponseBody
	GetId() *int64
	SetItems(v []*string) *GetListResponseBody
	GetItems() []*string
	SetKind(v string) *GetListResponseBody
	GetKind() *string
	SetName(v string) *GetListResponseBody
	GetName() *string
	SetRequestId(v string) *GetListResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *GetListResponseBody
	GetUpdateTime() *string
}

type GetListResponseBody struct {
	// The description of the custom list.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom list.[](~~2850217~~)
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The items in the custom list, which are displayed as an array.
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The type of the custom list.
	//
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the custom list.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the custom list was last modified.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListResponseBody) GoString() string {
	return s.String()
}

func (s *GetListResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetListResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetListResponseBody) GetItems() []*string {
	return s.Items
}

func (s *GetListResponseBody) GetKind() *string {
	return s.Kind
}

func (s *GetListResponseBody) GetName() *string {
	return s.Name
}

func (s *GetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetListResponseBody) SetDescription(v string) *GetListResponseBody {
	s.Description = &v
	return s
}

func (s *GetListResponseBody) SetId(v int64) *GetListResponseBody {
	s.Id = &v
	return s
}

func (s *GetListResponseBody) SetItems(v []*string) *GetListResponseBody {
	s.Items = v
	return s
}

func (s *GetListResponseBody) SetKind(v string) *GetListResponseBody {
	s.Kind = &v
	return s
}

func (s *GetListResponseBody) SetName(v string) *GetListResponseBody {
	s.Name = &v
	return s
}

func (s *GetListResponseBody) SetRequestId(v string) *GetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListResponseBody) SetUpdateTime(v string) *GetListResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetListResponseBody) Validate() error {
	return dara.Validate(s)
}
