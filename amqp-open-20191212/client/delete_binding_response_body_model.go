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
	// 021788F6-E50C-4BD6-9F80-66B0A19A6***
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
