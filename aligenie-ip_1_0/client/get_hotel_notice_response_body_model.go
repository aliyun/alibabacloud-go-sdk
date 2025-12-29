// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelNoticeResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelNoticeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelNoticeResponseBody
	GetRequestId() *string
	SetResult(v string) *GetHotelNoticeResponseBody
	GetResult() *string
}

type GetHotelNoticeResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test notice...
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetHotelNoticeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelNoticeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelNoticeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelNoticeResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetHotelNoticeResponseBody) SetCode(v int32) *GetHotelNoticeResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetMessage(v string) *GetHotelNoticeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetRequestId(v string) *GetHotelNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelNoticeResponseBody) SetResult(v string) *GetHotelNoticeResponseBody {
	s.Result = &v
	return s
}

func (s *GetHotelNoticeResponseBody) Validate() error {
	return dara.Validate(s)
}
