// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConstraintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConstraintId(v string) *UpdateConstraintResponseBody
	GetConstraintId() *string
	SetRequestId(v string) *UpdateConstraintResponseBody
	GetRequestId() *string
}

type UpdateConstraintResponseBody struct {
	// The ID of the constraint.
	//
	// example:
	//
	// cons-bp1yx7x42v****
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConstraintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConstraintResponseBody) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *UpdateConstraintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConstraintResponseBody) SetConstraintId(v string) *UpdateConstraintResponseBody {
	s.ConstraintId = &v
	return s
}

func (s *UpdateConstraintResponseBody) SetRequestId(v string) *UpdateConstraintResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConstraintResponseBody) Validate() error {
	return dara.Validate(s)
}
