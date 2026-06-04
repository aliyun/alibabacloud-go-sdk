// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateCustomAttributeRequest
	GetComment() *string
	SetDisplayEnabled(v bool) *UpdateCustomAttributeRequest
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *UpdateCustomAttributeRequest
	GetDisplayName() *string
	SetEntityTypes(v []*string) *UpdateCustomAttributeRequest
	GetEntityTypes() []*string
	SetId(v string) *UpdateCustomAttributeRequest
	GetId() *string
	SetSearchFilterEnabled(v bool) *UpdateCustomAttributeRequest
	GetSearchFilterEnabled() *bool
	SetValueEnums(v []*string) *UpdateCustomAttributeRequest
	GetValueEnums() []*string
}

type UpdateCustomAttributeRequest struct {
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// example:
	//
	// 业务负责人
	DisplayName *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	SearchFilterEnabled *bool     `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	ValueEnums          []*string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty" type:"Repeated"`
}

func (s UpdateCustomAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAttributeRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateCustomAttributeRequest) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *UpdateCustomAttributeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateCustomAttributeRequest) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *UpdateCustomAttributeRequest) GetId() *string {
	return s.Id
}

func (s *UpdateCustomAttributeRequest) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *UpdateCustomAttributeRequest) GetValueEnums() []*string {
	return s.ValueEnums
}

func (s *UpdateCustomAttributeRequest) SetComment(v string) *UpdateCustomAttributeRequest {
	s.Comment = &v
	return s
}

func (s *UpdateCustomAttributeRequest) SetDisplayEnabled(v bool) *UpdateCustomAttributeRequest {
	s.DisplayEnabled = &v
	return s
}

func (s *UpdateCustomAttributeRequest) SetDisplayName(v string) *UpdateCustomAttributeRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateCustomAttributeRequest) SetEntityTypes(v []*string) *UpdateCustomAttributeRequest {
	s.EntityTypes = v
	return s
}

func (s *UpdateCustomAttributeRequest) SetId(v string) *UpdateCustomAttributeRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomAttributeRequest) SetSearchFilterEnabled(v bool) *UpdateCustomAttributeRequest {
	s.SearchFilterEnabled = &v
	return s
}

func (s *UpdateCustomAttributeRequest) SetValueEnums(v []*string) *UpdateCustomAttributeRequest {
	s.ValueEnums = v
	return s
}

func (s *UpdateCustomAttributeRequest) Validate() error {
	return dara.Validate(s)
}
