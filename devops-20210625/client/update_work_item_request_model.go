// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldType(v string) *UpdateWorkItemRequest
	GetFieldType() *string
	SetIdentifier(v string) *UpdateWorkItemRequest
	GetIdentifier() *string
	SetPropertyKey(v string) *UpdateWorkItemRequest
	GetPropertyKey() *string
	SetPropertyValue(v string) *UpdateWorkItemRequest
	GetPropertyValue() *string
}

type UpdateWorkItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// subject
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e8b2xxxxxx2abdxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// subject
	PropertyKey *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// newValue
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s UpdateWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemRequest) GetFieldType() *string {
	return s.FieldType
}

func (s *UpdateWorkItemRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateWorkItemRequest) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *UpdateWorkItemRequest) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdateWorkItemRequest) SetFieldType(v string) *UpdateWorkItemRequest {
	s.FieldType = &v
	return s
}

func (s *UpdateWorkItemRequest) SetIdentifier(v string) *UpdateWorkItemRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkItemRequest) SetPropertyKey(v string) *UpdateWorkItemRequest {
	s.PropertyKey = &v
	return s
}

func (s *UpdateWorkItemRequest) SetPropertyValue(v string) *UpdateWorkItemRequest {
	s.PropertyValue = &v
	return s
}

func (s *UpdateWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
