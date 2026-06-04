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
	// example:
	//
	// []
	AttributeDefsShrink *string `json:"AttributeDefs,omitempty" xml:"AttributeDefs,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// NONE
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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
