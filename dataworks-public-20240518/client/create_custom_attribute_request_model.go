// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateCustomAttributeRequest
	GetComment() *string
	SetDisplayEnabled(v bool) *CreateCustomAttributeRequest
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *CreateCustomAttributeRequest
	GetDisplayName() *string
	SetEntityTypes(v []*string) *CreateCustomAttributeRequest
	GetEntityTypes() []*string
	SetId(v string) *CreateCustomAttributeRequest
	GetId() *string
	SetSearchFilterEnabled(v bool) *CreateCustomAttributeRequest
	GetSearchFilterEnabled() *bool
	SetType(v string) *CreateCustomAttributeRequest
	GetType() *string
	SetValueEnums(v []*string) *CreateCustomAttributeRequest
	GetValueEnums() []*string
}

type CreateCustomAttributeRequest struct {
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 业务负责人
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	EntityTypes []*string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// custom-attribute:biz_owner
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	SearchFilterEnabled *bool `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueEnums []*string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty" type:"Repeated"`
}

func (s CreateCustomAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAttributeRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAttributeRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateCustomAttributeRequest) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *CreateCustomAttributeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCustomAttributeRequest) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *CreateCustomAttributeRequest) GetId() *string {
	return s.Id
}

func (s *CreateCustomAttributeRequest) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *CreateCustomAttributeRequest) GetType() *string {
	return s.Type
}

func (s *CreateCustomAttributeRequest) GetValueEnums() []*string {
	return s.ValueEnums
}

func (s *CreateCustomAttributeRequest) SetComment(v string) *CreateCustomAttributeRequest {
	s.Comment = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetDisplayEnabled(v bool) *CreateCustomAttributeRequest {
	s.DisplayEnabled = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetDisplayName(v string) *CreateCustomAttributeRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetEntityTypes(v []*string) *CreateCustomAttributeRequest {
	s.EntityTypes = v
	return s
}

func (s *CreateCustomAttributeRequest) SetId(v string) *CreateCustomAttributeRequest {
	s.Id = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetSearchFilterEnabled(v bool) *CreateCustomAttributeRequest {
	s.SearchFilterEnabled = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetType(v string) *CreateCustomAttributeRequest {
	s.Type = &v
	return s
}

func (s *CreateCustomAttributeRequest) SetValueEnums(v []*string) *CreateCustomAttributeRequest {
	s.ValueEnums = v
	return s
}

func (s *CreateCustomAttributeRequest) Validate() error {
	return dara.Validate(s)
}
