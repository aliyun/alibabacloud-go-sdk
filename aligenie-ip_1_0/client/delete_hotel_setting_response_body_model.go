// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteHotelSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHotelSettingResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteHotelSettingResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *DeleteHotelSettingResponseBody
	GetStatusCode() *int32
}

type DeleteHotelSettingResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67****BB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteHotelSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHotelSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHotelSettingResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteHotelSettingResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotelSettingResponseBody) SetMessage(v string) *DeleteHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetRequestId(v string) *DeleteHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetResult(v bool) *DeleteHotelSettingResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) SetStatusCode(v int32) *DeleteHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
