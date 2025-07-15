// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExecutionResponseBody
	GetRequestId() *string
}

type UpdateExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C8345E88-5334-469E-901D-F912C8CB9C55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExecutionResponseBody) SetRequestId(v string) *UpdateExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
