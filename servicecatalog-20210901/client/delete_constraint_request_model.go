// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConstraintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConstraintId(v string) *DeleteConstraintRequest
	GetConstraintId() *string
}

type DeleteConstraintRequest struct {
	// The ID of the constraint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cons-bp1yx7x42v****
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
}

func (s DeleteConstraintRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConstraintRequest) GoString() string {
	return s.String()
}

func (s *DeleteConstraintRequest) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *DeleteConstraintRequest) SetConstraintId(v string) *DeleteConstraintRequest {
	s.ConstraintId = &v
	return s
}

func (s *DeleteConstraintRequest) Validate() error {
	return dara.Validate(s)
}
