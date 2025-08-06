// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v string) *GetMessageCallbackEventListResponseBody
	GetEventList() *string
	SetRequestId(v string) *GetMessageCallbackEventListResponseBody
	GetRequestId() *string
}

type GetMessageCallbackEventListResponseBody struct {
	EventList *string `json:"EventList,omitempty" xml:"EventList,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageCallbackEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackEventListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackEventListResponseBody) GetEventList() *string {
	return s.EventList
}

func (s *GetMessageCallbackEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageCallbackEventListResponseBody) SetEventList(v string) *GetMessageCallbackEventListResponseBody {
	s.EventList = &v
	return s
}

func (s *GetMessageCallbackEventListResponseBody) SetRequestId(v string) *GetMessageCallbackEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageCallbackEventListResponseBody) Validate() error {
	return dara.Validate(s)
}
