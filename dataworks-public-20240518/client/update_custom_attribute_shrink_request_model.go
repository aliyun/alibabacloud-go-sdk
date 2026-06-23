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
	// The new description for the custom attribute. It must be 256 characters or less.
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Whether to display the custom attribute in the UI.
	//
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// The new display name for the custom attribute. It must be 128 characters or less.
	//
	// example:
	//
	// 业务负责人
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The applicable entity types. This parameter supports specific types and wildcard formats, such as `*-table` and `*-column`. For example:
	//
	// - `dataworks-project`: A DataWorks workspace
	//
	// - `dataworks-dataset`: A DataWorks dataset
	//
	// - `maxcompute-table`: A MaxCompute table
	//
	// - `*-column`: All column types
	EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
	// The custom attribute ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom-attribute:biz_owner
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether the custom attribute can be used as a filter condition.
	//
	// example:
	//
	// false
	SearchFilterEnabled *bool `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	// The enumerated values. This applies only to custom attributes of the `enum` type. You can only append new values during an update.
	ValueEnumsShrink *string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty"`
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
