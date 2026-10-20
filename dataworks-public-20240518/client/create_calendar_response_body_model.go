// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateCalendarResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateCalendarResponseBody
	GetRequestId() *string
}

type CreateCalendarResponseBody struct {
	// The ID of the scheduling calendar.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCalendarResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCalendarResponseBody) SetId(v int64) *CreateCalendarResponseBody {
	s.Id = &v
	return s
}

func (s *CreateCalendarResponseBody) SetRequestId(v string) *CreateCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
