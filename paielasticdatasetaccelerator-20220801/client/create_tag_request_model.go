// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *CreateTagRequest
	GetKey() *string
	SetResourceId(v string) *CreateTagRequest
	GetResourceId() *string
	SetResourceType(v string) *CreateTagRequest
	GetResourceType() *string
	SetValue(v string) *CreateTagRequest
	GetValue() *string
}

type CreateTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// inst-my1tk3jggooi5uwks
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetKey() *string {
	return s.Key
}

func (s *CreateTagRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateTagRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateTagRequest) GetValue() *string {
	return s.Value
}

func (s *CreateTagRequest) SetKey(v string) *CreateTagRequest {
	s.Key = &v
	return s
}

func (s *CreateTagRequest) SetResourceId(v string) *CreateTagRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateTagRequest) SetResourceType(v string) *CreateTagRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateTagRequest) SetValue(v string) *CreateTagRequest {
	s.Value = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
