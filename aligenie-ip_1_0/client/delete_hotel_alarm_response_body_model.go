// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *DeleteHotelAlarmResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *DeleteHotelAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHotelAlarmResponseBody
	GetRequestId() *string
	SetResult(v int32) *DeleteHotelAlarmResponseBody
	GetResult() *int32
	SetStatusCode(v int32) *DeleteHotelAlarmResponseBody
	GetStatusCode() *int32
}

type DeleteHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteHotelAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *DeleteHotelAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHotelAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHotelAlarmResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *DeleteHotelAlarmResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *DeleteHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetMessage(v string) *DeleteHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetRequestId(v string) *DeleteHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetResult(v int32) *DeleteHotelAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) SetStatusCode(v int32) *DeleteHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
