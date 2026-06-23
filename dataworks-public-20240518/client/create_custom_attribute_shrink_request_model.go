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
	// The description of the custom attribute. The description must be less than 256 characters in length.
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether to display the attribute on the product page. The default value is true.
	//
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// The display name of the custom attribute. The name must be less than 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 业务负责人
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The list of applicable entity types. You can specify precise entity types or use wildcards such as `*-table` and `*-column`. Examples:
	//
	// - dataworks-project: a DataWorks workspace.
	//
	// - dataworks-dataset: a DataWorks dataset.
	//
	// - maxcompute-table: a MaxCompute table.
	//
	// - \\*-column: all field types.
	//
	// This parameter is required.
	EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
	// The ID of the custom attribute. The ID must match the regular expression `^custom-attribute:[A-Za-z][A-Za-z0-9_]{0,98}$`. The part after \\`custom-attribute:\\` must be less than 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom-attribute:biz_owner
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether the attribute can be used as a filter on the Data Map search page. The default value is false. Currently, you can set this parameter to true only for attributes of the ENUM type.
	//
	// example:
	//
	// false
	SearchFilterEnabled *bool `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	// The type of the custom attribute. Valid values are ENUM, TEXT, and HYPERLINK.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The enumeration values. This parameter is required when \\`Type\\` is set to \\`ENUM\\`. It is not supported for the TEXT and HYPERLINK types.
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
