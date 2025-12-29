// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *UpdateHotelAlarmResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *UpdateHotelAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHotelAlarmResponseBody
	GetRequestId() *string
	SetResult(v int32) *UpdateHotelAlarmResponseBody
	GetResult() *int32
	SetStatusCode(v int32) *UpdateHotelAlarmResponseBody
	GetStatusCode() *int32
}

type UpdateHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
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

func (s UpdateHotelAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *UpdateHotelAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHotelAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotelAlarmResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *UpdateHotelAlarmResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *UpdateHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetMessage(v string) *UpdateHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetRequestId(v string) *UpdateHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetResult(v int32) *UpdateHotelAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) SetStatusCode(v int32) *UpdateHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
