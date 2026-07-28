// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResourceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindResourceControlResponseBody
	GetRequestId() *string
}

type UnbindResourceControlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C0ACF0-DD29-4B67-9190-B7A48C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResourceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindResourceControlResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResourceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindResourceControlResponseBody) SetRequestId(v string) *UnbindResourceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindResourceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
