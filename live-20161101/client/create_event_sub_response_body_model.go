// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEventSubResponseBody
	GetRequestId() *string
	SetSubscribeId(v string) *CreateEventSubResponseBody
	GetSubscribeId() *string
}

type CreateEventSubResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 760bad53276431c499e30dc36f6b****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s CreateEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventSubResponseBody) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *CreateEventSubResponseBody) SetRequestId(v string) *CreateEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventSubResponseBody) SetSubscribeId(v string) *CreateEventSubResponseBody {
	s.SubscribeId = &v
	return s
}

func (s *CreateEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}
