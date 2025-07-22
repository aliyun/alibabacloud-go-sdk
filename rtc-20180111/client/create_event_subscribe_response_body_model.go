// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEventSubscribeResponseBody
	GetRequestId() *string
	SetSubscribeId(v string) *CreateEventSubscribeResponseBody
	GetSubscribeId() *string
}

type CreateEventSubscribeResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s CreateEventSubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventSubscribeResponseBody) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *CreateEventSubscribeResponseBody) SetRequestId(v string) *CreateEventSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventSubscribeResponseBody) SetSubscribeId(v string) *CreateEventSubscribeResponseBody {
	s.SubscribeId = &v
	return s
}

func (s *CreateEventSubscribeResponseBody) Validate() error {
	return dara.Validate(s)
}
