// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCalendarResponseBody
	GetSuccess() *bool
}

type DeleteCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CF99C381-C8F6-5A8D-8C24-57F46B706D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCalendarResponseBody) SetCode(v int32) *DeleteCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCalendarResponseBody) SetMessage(v string) *DeleteCalendarResponseBody {
	s.Message = &v
	return s
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
