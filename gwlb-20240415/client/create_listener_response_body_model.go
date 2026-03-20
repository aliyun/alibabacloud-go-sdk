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
	// The listener ID.
	//
	// example:
	//
	// lsn-wi3c0v30ivysrg****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A045E652-D298-5E70-A978-7247135336FB
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
