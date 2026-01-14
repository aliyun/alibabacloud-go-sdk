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
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
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
