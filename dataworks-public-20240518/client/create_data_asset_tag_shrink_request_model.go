// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAssetTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDataAssetTagShrinkRequest
	GetDescription() *string
	SetKey(v string) *CreateDataAssetTagShrinkRequest
	GetKey() *string
	SetManagersShrink(v string) *CreateDataAssetTagShrinkRequest
	GetManagersShrink() *string
	SetValueType(v string) *CreateDataAssetTagShrinkRequest
	GetValueType() *string
	SetValuesShrink(v string) *CreateDataAssetTagShrinkRequest
	GetValuesShrink() *string
}

type CreateDataAssetTagShrinkRequest struct {
	// The description of the tag.
	//
	// example:
	//
	// This is a description
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
	// The type of the tag value. Valid values:
	//
	// 	- Boolean
	//
	// 	- Int
	//
	// 	- String
	//
	// 	- Double
	//
	// example:
	//
	// String
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The tag values.
	ValuesShrink *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s CreateDataAssetTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAssetTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAssetTagShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAssetTagShrinkRequest) GetKey() *string {
	return s.Key
}

func (s *CreateDataAssetTagShrinkRequest) GetManagersShrink() *string {
	return s.ManagersShrink
}

func (s *CreateDataAssetTagShrinkRequest) GetValueType() *string {
	return s.ValueType
}

func (s *CreateDataAssetTagShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *CreateDataAssetTagShrinkRequest) SetDescription(v string) *CreateDataAssetTagShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDataAssetTagShrinkRequest) SetKey(v string) *CreateDataAssetTagShrinkRequest {
	s.Key = &v
	return s
}

func (s *CreateDataAssetTagShrinkRequest) SetManagersShrink(v string) *CreateDataAssetTagShrinkRequest {
	s.ManagersShrink = &v
	return s
}

func (s *CreateDataAssetTagShrinkRequest) SetValueType(v string) *CreateDataAssetTagShrinkRequest {
	s.ValueType = &v
	return s
}

func (s *CreateDataAssetTagShrinkRequest) SetValuesShrink(v string) *CreateDataAssetTagShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *CreateDataAssetTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
