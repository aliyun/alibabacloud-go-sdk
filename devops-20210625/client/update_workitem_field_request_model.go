// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateWorkitemPropertyRequest(v []*UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) *UpdateWorkitemFieldRequest
	GetUpdateWorkitemPropertyRequest() []*UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest
	SetWorkitemIdentifier(v string) *UpdateWorkitemFieldRequest
	GetWorkitemIdentifier() *string
}

type UpdateWorkitemFieldRequest struct {
	// This parameter is required.
	UpdateWorkitemPropertyRequest []*UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest `json:"updateWorkitemPropertyRequest,omitempty" xml:"updateWorkitemPropertyRequest,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s UpdateWorkitemFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemFieldRequest) GetUpdateWorkitemPropertyRequest() []*UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest {
	return s.UpdateWorkitemPropertyRequest
}

func (s *UpdateWorkitemFieldRequest) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *UpdateWorkitemFieldRequest) SetUpdateWorkitemPropertyRequest(v []*UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) *UpdateWorkitemFieldRequest {
	s.UpdateWorkitemPropertyRequest = v
	return s
}

func (s *UpdateWorkitemFieldRequest) SetWorkitemIdentifier(v string) *UpdateWorkitemFieldRequest {
	s.WorkitemIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldRequest) Validate() error {
	if s.UpdateWorkitemPropertyRequest != nil {
		for _, item := range s.UpdateWorkitemPropertyRequest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tag
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// This parameter is required.
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) SetFieldIdentifier(v string) *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest {
	s.FieldIdentifier = &v
	return s
}

func (s *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) SetFieldValue(v string) *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest {
	s.FieldValue = &v
	return s
}

func (s *UpdateWorkitemFieldRequestUpdateWorkitemPropertyRequest) Validate() error {
	return dara.Validate(s)
}
