// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListenerId(v string) *CreateListenerResponseBody
	GetListenerId() *string
	SetRequestId(v string) *CreateListenerResponseBody
	GetRequestId() *string
}

type CreateListenerResponseBody struct {
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
