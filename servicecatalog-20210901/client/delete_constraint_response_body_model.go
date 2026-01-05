// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConstraintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConstraintResponseBody
	GetRequestId() *string
}

type DeleteConstraintResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConstraintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConstraintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConstraintResponseBody) SetRequestId(v string) *DeleteConstraintResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConstraintResponseBody) Validate() error {
	return dara.Validate(s)
}
