// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDataAssetTagShrinkRequest
	GetDescription() *string
	SetKey(v string) *UpdateDataAssetTagShrinkRequest
	GetKey() *string
	SetManagersShrink(v string) *UpdateDataAssetTagShrinkRequest
	GetManagersShrink() *string
	SetValuesShrink(v string) *UpdateDataAssetTagShrinkRequest
	GetValuesShrink() *string
}

type UpdateDataAssetTagShrinkRequest struct {
	// The description of the tag.
	//
	// example:
	//
	// This is a description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag administrators.
	ManagersShrink *string `json:"Managers,omitempty" xml:"Managers,omitempty"`
	// The tag values.
	ValuesShrink *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s UpdateDataAssetTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetTagShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataAssetTagShrinkRequest) GetKey() *string {
	return s.Key
}

func (s *UpdateDataAssetTagShrinkRequest) GetManagersShrink() *string {
	return s.ManagersShrink
}

func (s *UpdateDataAssetTagShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *UpdateDataAssetTagShrinkRequest) SetDescription(v string) *UpdateDataAssetTagShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataAssetTagShrinkRequest) SetKey(v string) *UpdateDataAssetTagShrinkRequest {
	s.Key = &v
	return s
}

func (s *UpdateDataAssetTagShrinkRequest) SetManagersShrink(v string) *UpdateDataAssetTagShrinkRequest {
	s.ManagersShrink = &v
	return s
}

func (s *UpdateDataAssetTagShrinkRequest) SetValuesShrink(v string) *UpdateDataAssetTagShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *UpdateDataAssetTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
