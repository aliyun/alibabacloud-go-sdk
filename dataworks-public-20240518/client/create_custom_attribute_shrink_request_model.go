// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateCustomAttributeShrinkRequest
	GetComment() *string
	SetDisplayEnabled(v bool) *CreateCustomAttributeShrinkRequest
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *CreateCustomAttributeShrinkRequest
	GetDisplayName() *string
	SetEntityTypesShrink(v string) *CreateCustomAttributeShrinkRequest
	GetEntityTypesShrink() *string
	SetId(v string) *CreateCustomAttributeShrinkRequest
	GetId() *string
	SetSearchFilterEnabled(v bool) *CreateCustomAttributeShrinkRequest
	GetSearchFilterEnabled() *bool
	SetType(v string) *CreateCustomAttributeShrinkRequest
	GetType() *string
	SetValueEnumsShrink(v string) *CreateCustomAttributeShrinkRequest
	GetValueEnumsShrink() *string
}

type CreateCustomAttributeShrinkRequest struct {
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
	EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
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
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueEnumsShrink *string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty"`
}

func (s CreateCustomAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAttributeShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateCustomAttributeShrinkRequest) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *CreateCustomAttributeShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCustomAttributeShrinkRequest) GetEntityTypesShrink() *string {
	return s.EntityTypesShrink
}

func (s *CreateCustomAttributeShrinkRequest) GetId() *string {
	return s.Id
}

func (s *CreateCustomAttributeShrinkRequest) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *CreateCustomAttributeShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateCustomAttributeShrinkRequest) GetValueEnumsShrink() *string {
	return s.ValueEnumsShrink
}

func (s *CreateCustomAttributeShrinkRequest) SetComment(v string) *CreateCustomAttributeShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetDisplayEnabled(v bool) *CreateCustomAttributeShrinkRequest {
	s.DisplayEnabled = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetDisplayName(v string) *CreateCustomAttributeShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetEntityTypesShrink(v string) *CreateCustomAttributeShrinkRequest {
	s.EntityTypesShrink = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetId(v string) *CreateCustomAttributeShrinkRequest {
	s.Id = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetSearchFilterEnabled(v bool) *CreateCustomAttributeShrinkRequest {
	s.SearchFilterEnabled = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetType(v string) *CreateCustomAttributeShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) SetValueEnumsShrink(v string) *CreateCustomAttributeShrinkRequest {
	s.ValueEnumsShrink = &v
	return s
}

func (s *CreateCustomAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
