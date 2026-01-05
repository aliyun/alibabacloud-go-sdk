// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConstraintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConstraintId(v string) *CreateConstraintResponseBody
	GetConstraintId() *string
	SetRequestId(v string) *CreateConstraintResponseBody
	GetRequestId() *string
}

type CreateConstraintResponseBody struct {
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

func (s CreateConstraintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConstraintResponseBody) GetConstraintId() *string {
	return s.ConstraintId
}

func (s *CreateConstraintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConstraintResponseBody) SetConstraintId(v string) *CreateConstraintResponseBody {
	s.ConstraintId = &v
	return s
}

func (s *CreateConstraintResponseBody) SetRequestId(v string) *CreateConstraintResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConstraintResponseBody) Validate() error {
	return dara.Validate(s)
}
