// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateModelRequest
	GetDescription() *string
	SetGroupId(v string) *CreateModelRequest
	GetGroupId() *string
	SetModelName(v string) *CreateModelRequest
	GetModelName() *string
	SetSchema(v string) *CreateModelRequest
	GetSchema() *string
	SetTag(v []*CreateModelRequestTag) *CreateModelRequest
	GetTag() []*CreateModelRequestTag
}

type CreateModelRequest struct {
	// The description of the model definition.
	//
	// example:
	//
	// Model Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group to which the model belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30e792398d6c4569b04c0e53a3494381
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model. The name must be unique within the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The definition of the model in JSON Schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"type":"object","properties":{"id":{"format":"int64","maximum":100,"exclusiveMaximum":true,"type":"integer"},"name":{"maxLength":10,"type":"string"}}}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The object tags that match the lifecycle rule. You can specify multiple tags.
	Tag []*CreateModelRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelRequest) GetSchema() *string {
	return s.Schema
}

func (s *CreateModelRequest) GetTag() []*CreateModelRequestTag {
	return s.Tag
}

func (s *CreateModelRequest) SetDescription(v string) *CreateModelRequest {
	s.Description = &v
	return s
}

func (s *CreateModelRequest) SetGroupId(v string) *CreateModelRequest {
	s.GroupId = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetSchema(v string) *CreateModelRequest {
	s.Schema = &v
	return s
}

func (s *CreateModelRequest) SetTag(v []*CreateModelRequestTag) *CreateModelRequest {
	s.Tag = v
	return s
}

func (s *CreateModelRequest) Validate() error {
	return dara.Validate(s)
}

type CreateModelRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateModelRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestTag) GoString() string {
	return s.String()
}

func (s *CreateModelRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateModelRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateModelRequestTag) SetKey(v string) *CreateModelRequestTag {
	s.Key = &v
	return s
}

func (s *CreateModelRequestTag) SetValue(v string) *CreateModelRequestTag {
	s.Value = &v
	return s
}

func (s *CreateModelRequestTag) Validate() error {
	return dara.Validate(s)
}
