// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestCaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateWorkitemPropertyRequest(v []*UpdateTestCaseRequestUpdateWorkitemPropertyRequest) *UpdateTestCaseRequest
	GetUpdateWorkitemPropertyRequest() []*UpdateTestCaseRequestUpdateWorkitemPropertyRequest
}

type UpdateTestCaseRequest struct {
	// This parameter is required.
	UpdateWorkitemPropertyRequest []*UpdateTestCaseRequestUpdateWorkitemPropertyRequest `json:"updateWorkitemPropertyRequest,omitempty" xml:"updateWorkitemPropertyRequest,omitempty" type:"Repeated"`
}

func (s UpdateTestCaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseRequest) GetUpdateWorkitemPropertyRequest() []*UpdateTestCaseRequestUpdateWorkitemPropertyRequest {
	return s.UpdateWorkitemPropertyRequest
}

func (s *UpdateTestCaseRequest) SetUpdateWorkitemPropertyRequest(v []*UpdateTestCaseRequestUpdateWorkitemPropertyRequest) *UpdateTestCaseRequest {
	s.UpdateWorkitemPropertyRequest = v
	return s
}

func (s *UpdateTestCaseRequest) Validate() error {
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

type UpdateTestCaseRequestUpdateWorkitemPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc.type
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0a032xx28107xxxx53e87a9
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s UpdateTestCaseRequestUpdateWorkitemPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseRequestUpdateWorkitemPropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseRequestUpdateWorkitemPropertyRequest) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *UpdateTestCaseRequestUpdateWorkitemPropertyRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateTestCaseRequestUpdateWorkitemPropertyRequest) SetFieldIdentifier(v string) *UpdateTestCaseRequestUpdateWorkitemPropertyRequest {
	s.FieldIdentifier = &v
	return s
}

func (s *UpdateTestCaseRequestUpdateWorkitemPropertyRequest) SetFieldValue(v string) *UpdateTestCaseRequestUpdateWorkitemPropertyRequest {
	s.FieldValue = &v
	return s
}

func (s *UpdateTestCaseRequestUpdateWorkitemPropertyRequest) Validate() error {
	return dara.Validate(s)
}
