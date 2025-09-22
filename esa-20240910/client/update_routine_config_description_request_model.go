// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineConfigDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRoutineConfigDescriptionRequest
	GetDescription() *string
	SetName(v string) *UpdateRoutineConfigDescriptionRequest
	GetName() *string
}

type UpdateRoutineConfigDescriptionRequest struct {
	// example:
	//
	// description of this routine
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateRoutineConfigDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineConfigDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoutineConfigDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRoutineConfigDescriptionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRoutineConfigDescriptionRequest) SetDescription(v string) *UpdateRoutineConfigDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateRoutineConfigDescriptionRequest) SetName(v string) *UpdateRoutineConfigDescriptionRequest {
	s.Name = &v
	return s
}

func (s *UpdateRoutineConfigDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
