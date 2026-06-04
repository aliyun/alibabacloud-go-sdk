// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityDefShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateMetaEntityDefShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateMetaEntityDefShrinkRequest
	GetDisplayName() *string
	SetEntityType(v string) *UpdateMetaEntityDefShrinkRequest
	GetEntityType() *string
	SetNewAttributeDefsShrink(v string) *UpdateMetaEntityDefShrinkRequest
	GetNewAttributeDefsShrink() *string
	SetUpdateAttributeDefsShrink(v string) *UpdateMetaEntityDefShrinkRequest
	GetUpdateAttributeDefsShrink() *string
}

type UpdateMetaEntityDefShrinkRequest struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// []
	NewAttributeDefsShrink *string `json:"NewAttributeDefs,omitempty" xml:"NewAttributeDefs,omitempty"`
	// example:
	//
	// []
	UpdateAttributeDefsShrink *string `json:"UpdateAttributeDefs,omitempty" xml:"UpdateAttributeDefs,omitempty"`
}

func (s UpdateMetaEntityDefShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityDefShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityDefShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMetaEntityDefShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateMetaEntityDefShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateMetaEntityDefShrinkRequest) GetNewAttributeDefsShrink() *string {
	return s.NewAttributeDefsShrink
}

func (s *UpdateMetaEntityDefShrinkRequest) GetUpdateAttributeDefsShrink() *string {
	return s.UpdateAttributeDefsShrink
}

func (s *UpdateMetaEntityDefShrinkRequest) SetDescription(v string) *UpdateMetaEntityDefShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetaEntityDefShrinkRequest) SetDisplayName(v string) *UpdateMetaEntityDefShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateMetaEntityDefShrinkRequest) SetEntityType(v string) *UpdateMetaEntityDefShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateMetaEntityDefShrinkRequest) SetNewAttributeDefsShrink(v string) *UpdateMetaEntityDefShrinkRequest {
	s.NewAttributeDefsShrink = &v
	return s
}

func (s *UpdateMetaEntityDefShrinkRequest) SetUpdateAttributeDefsShrink(v string) *UpdateMetaEntityDefShrinkRequest {
	s.UpdateAttributeDefsShrink = &v
	return s
}

func (s *UpdateMetaEntityDefShrinkRequest) Validate() error {
	return dara.Validate(s)
}
