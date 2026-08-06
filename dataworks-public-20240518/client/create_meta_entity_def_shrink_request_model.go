// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaEntityDefShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeDefsShrink(v string) *CreateMetaEntityDefShrinkRequest
	GetAttributeDefsShrink() *string
	SetDescription(v string) *CreateMetaEntityDefShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateMetaEntityDefShrinkRequest
	GetDisplayName() *string
	SetExtend(v string) *CreateMetaEntityDefShrinkRequest
	GetExtend() *string
	SetName(v string) *CreateMetaEntityDefShrinkRequest
	GetName() *string
}

type CreateMetaEntityDefShrinkRequest struct {
	// The attribute definition list for custom entities. This parameter cannot be specified when extend is set to TABLE.
	//
	// example:
	//
	// []
	AttributeDefsShrink *string `json:"AttributeDefs,omitempty" xml:"AttributeDefs,omitempty"`
	// The description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name, up to 32 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The extension mode. Valid values:
	//
	// - NONE: default value. Indicates a custom entity with freely defined attributes.
	//
	// - TABLE: indicates an extended table type. This type integrates in the same way as existing table types in DataWorks Data Map. You do not need to provide attribute definitions and can create corresponding Database/Table objects.
	//
	// example:
	//
	// TABLE
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The name of the type definition. Custom types must match `^[a-z0-9][a-z0-9_]*$`. Extended table types must match `^[a-z0-9][a-z0-9_]*-table$`.
	//
	// This parameter is required.
	//
	// example:
	//
	// biz_api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMetaEntityDefShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaEntityDefShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaEntityDefShrinkRequest) GetAttributeDefsShrink() *string {
	return s.AttributeDefsShrink
}

func (s *CreateMetaEntityDefShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetaEntityDefShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateMetaEntityDefShrinkRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateMetaEntityDefShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaEntityDefShrinkRequest) SetAttributeDefsShrink(v string) *CreateMetaEntityDefShrinkRequest {
	s.AttributeDefsShrink = &v
	return s
}

func (s *CreateMetaEntityDefShrinkRequest) SetDescription(v string) *CreateMetaEntityDefShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateMetaEntityDefShrinkRequest) SetDisplayName(v string) *CreateMetaEntityDefShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateMetaEntityDefShrinkRequest) SetExtend(v string) *CreateMetaEntityDefShrinkRequest {
	s.Extend = &v
	return s
}

func (s *CreateMetaEntityDefShrinkRequest) SetName(v string) *CreateMetaEntityDefShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaEntityDefShrinkRequest) Validate() error {
	return dara.Validate(s)
}
