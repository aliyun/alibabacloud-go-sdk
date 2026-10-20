// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCalendarResponseBody
	GetSuccess() *bool
}

type DeleteCalendarResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCalendarResponseBody) SetRequestId(v string) *DeleteCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCalendarResponseBody) SetSuccess(v bool) *DeleteCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
