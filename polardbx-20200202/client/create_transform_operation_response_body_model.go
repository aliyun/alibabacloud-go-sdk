// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransformOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransformOperationResponseBody
	GetRequestId() *string
}

type CreateTransformOperationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransformOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransformOperationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransformOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransformOperationResponseBody) SetRequestId(v string) *CreateTransformOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransformOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
