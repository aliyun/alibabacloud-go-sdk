// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateCustomAttributeShrinkRequest
	GetComment() *string
	SetDisplayEnabled(v bool) *UpdateCustomAttributeShrinkRequest
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *UpdateCustomAttributeShrinkRequest
	GetDisplayName() *string
	SetEntityTypesShrink(v string) *UpdateCustomAttributeShrinkRequest
	GetEntityTypesShrink() *string
	SetId(v string) *UpdateCustomAttributeShrinkRequest
	GetId() *string
	SetSearchFilterEnabled(v bool) *UpdateCustomAttributeShrinkRequest
	GetSearchFilterEnabled() *bool
	SetValueEnumsShrink(v string) *UpdateCustomAttributeShrinkRequest
	GetValueEnumsShrink() *string
}

type UpdateCustomAttributeShrinkRequest struct {
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
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	SearchFilterEnabled *bool   `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	ValueEnumsShrink    *string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty"`
}

func (s UpdateCustomAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAttributeShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateCustomAttributeShrinkRequest) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *UpdateCustomAttributeShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateCustomAttributeShrinkRequest) GetEntityTypesShrink() *string {
	return s.EntityTypesShrink
}

func (s *UpdateCustomAttributeShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateCustomAttributeShrinkRequest) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *UpdateCustomAttributeShrinkRequest) GetValueEnumsShrink() *string {
	return s.ValueEnumsShrink
}

func (s *UpdateCustomAttributeShrinkRequest) SetComment(v string) *UpdateCustomAttributeShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetDisplayEnabled(v bool) *UpdateCustomAttributeShrinkRequest {
	s.DisplayEnabled = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetDisplayName(v string) *UpdateCustomAttributeShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetEntityTypesShrink(v string) *UpdateCustomAttributeShrinkRequest {
	s.EntityTypesShrink = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetId(v string) *UpdateCustomAttributeShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetSearchFilterEnabled(v bool) *UpdateCustomAttributeShrinkRequest {
	s.SearchFilterEnabled = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) SetValueEnumsShrink(v string) *UpdateCustomAttributeShrinkRequest {
	s.ValueEnumsShrink = &v
	return s
}

func (s *UpdateCustomAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
