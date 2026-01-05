// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConstraintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConstraintId(v string) *GetConstraintRequest
	GetConstraintId() *string
}

type GetConstraintRequest struct {
	// The ID of the constraint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cons-bp1yx7x42v****
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
}

func (s GetConstraintRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConstraintRequest) GoString() string {
	return s.String()
}

func (s *GetConstraintRequest) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *GetConstraintRequest) SetConstraintId(v string) *GetConstraintRequest {
	s.ConstraintId = &v
	return s
}

func (s *GetConstraintRequest) Validate() error {
	return dara.Validate(s)
}
