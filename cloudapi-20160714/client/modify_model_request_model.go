// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyModelRequest
	GetDescription() *string
	SetGroupId(v string) *ModifyModelRequest
	GetGroupId() *string
	SetModelName(v string) *ModifyModelRequest
	GetModelName() *string
	SetNewModelName(v string) *ModifyModelRequest
	GetNewModelName() *string
	SetSchema(v string) *ModifyModelRequest
	GetSchema() *string
}

type ModifyModelRequest struct {
	// The description of the new model definition.
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
	// The name of the model.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The new name of the model.
	//
	// example:
	//
	// NewTest
	NewModelName *string `json:"NewModelName,omitempty" xml:"NewModelName,omitempty"`
	// The new definition of the model.
	//
	// example:
	//
	// {\\"type\\":\\"object\\",\\"properties\\":{\\"id\\":{\\"format\\":\\"int64\\",\\"maximum\\":100,\\"exclusiveMaximum\\":true,\\"type\\":\\"integer\\"},\\"name\\":{\\"maxLength\\":10,\\"type\\":\\"string\\"}}}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ModifyModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyModelRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModifyModelRequest) GetNewModelName() *string {
	return s.NewModelName
}

func (s *ModifyModelRequest) GetSchema() *string {
	return s.Schema
}

func (s *ModifyModelRequest) SetDescription(v string) *ModifyModelRequest {
	s.Description = &v
	return s
}

func (s *ModifyModelRequest) SetGroupId(v string) *ModifyModelRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyModelRequest) SetModelName(v string) *ModifyModelRequest {
	s.ModelName = &v
	return s
}

func (s *ModifyModelRequest) SetNewModelName(v string) *ModifyModelRequest {
	s.NewModelName = &v
	return s
}

func (s *ModifyModelRequest) SetSchema(v string) *ModifyModelRequest {
	s.Schema = &v
	return s
}

func (s *ModifyModelRequest) Validate() error {
	return dara.Validate(s)
}
