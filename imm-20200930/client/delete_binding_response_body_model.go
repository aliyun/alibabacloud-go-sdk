// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBindingResponseBody
	GetRequestId() *string
}

type DeleteBindingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ACDFE467-C817-4B36-951A-6EB5A592****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBindingResponseBody) SetRequestId(v string) *DeleteBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
