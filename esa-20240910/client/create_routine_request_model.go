// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRoutineRequest
	GetDescription() *string
	SetName(v string) *CreateRoutineRequest
	GetName() *string
}

type CreateRoutineRequest struct {
	// The routine description.
	//
	// example:
	//
	// the description of this routine
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The routine name, which must be unique in the same account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRoutineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoutineRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineRequest) SetDescription(v string) *CreateRoutineRequest {
	s.Description = &v
	return s
}

func (s *CreateRoutineRequest) SetName(v string) *CreateRoutineRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineRequest) Validate() error {
	return dara.Validate(s)
}
