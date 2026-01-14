// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateListenerResponseBody
	GetRequestId() *string
}

type UpdateListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateListenerResponseBody) SetRequestId(v string) *UpdateListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
