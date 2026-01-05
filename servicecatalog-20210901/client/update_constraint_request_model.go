// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConstraintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateConstraintRequest
	GetConfig() *string
	SetConstraintId(v string) *UpdateConstraintRequest
	GetConstraintId() *string
	SetDescription(v string) *UpdateConstraintRequest
	GetDescription() *string
}

type UpdateConstraintRequest struct {
	// The configurations of the constraint.
	//
	// Format: { "LocalRoleName": "\\<role_name>" }.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "LocalRoleName": "TestRole" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the constraint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cons-bp1yx7x42v****
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The description of the constraint.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Launch as local role TestRole
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateConstraintRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConstraintRequest) GoString() string {
	return s.String()
}

func (s *UpdateConstraintRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateConstraintRequest) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *UpdateConstraintRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConstraintRequest) SetConfig(v string) *UpdateConstraintRequest {
	s.Config = &v
	return s
}

func (s *UpdateConstraintRequest) SetConstraintId(v string) *UpdateConstraintRequest {
	s.ConstraintId = &v
	return s
}

func (s *UpdateConstraintRequest) SetDescription(v string) *UpdateConstraintRequest {
	s.Description = &v
	return s
}

func (s *UpdateConstraintRequest) Validate() error {
	return dara.Validate(s)
}
