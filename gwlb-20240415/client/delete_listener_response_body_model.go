// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteListenerResponseBody
	GetRequestId() *string
}

type DeleteListenerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5C6E3548-086F-5FF6-A2B3-B1871B3AB488
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
