// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindResponseBody
	GetRequestId() *string
}

type UnbindResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6EBD4046-2202-5FBD-8595-4B631F0C484B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindResponseBody) Validate() error {
	return dara.Validate(s)
}
