// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributesShrink(v string) *UpdateMetaEntityShrinkRequest
	GetAttributesShrink() *string
	SetComment(v string) *UpdateMetaEntityShrinkRequest
	GetComment() *string
	SetCustomAttributesShrink(v string) *UpdateMetaEntityShrinkRequest
	GetCustomAttributesShrink() *string
	SetId(v string) *UpdateMetaEntityShrinkRequest
	GetId() *string
}

type UpdateMetaEntityShrinkRequest struct {
	// example:
	//
	// []
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// []
	CustomAttributesShrink *string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api:api_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateMetaEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityShrinkRequest) GetAttributesShrink() *string {
	return s.AttributesShrink
}

func (s *UpdateMetaEntityShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateMetaEntityShrinkRequest) GetCustomAttributesShrink() *string {
	return s.CustomAttributesShrink
}

func (s *UpdateMetaEntityShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMetaEntityShrinkRequest) SetAttributesShrink(v string) *UpdateMetaEntityShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *UpdateMetaEntityShrinkRequest) SetComment(v string) *UpdateMetaEntityShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateMetaEntityShrinkRequest) SetCustomAttributesShrink(v string) *UpdateMetaEntityShrinkRequest {
	s.CustomAttributesShrink = &v
	return s
}

func (s *UpdateMetaEntityShrinkRequest) SetId(v string) *UpdateMetaEntityShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMetaEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
