// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *CreateHotelAlarmResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *CreateHotelAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHotelAlarmResponseBody
	GetRequestId() *string
	SetResult(v []*CreateHotelAlarmResponseBodyResult) *CreateHotelAlarmResponseBody
	GetResult() []*CreateHotelAlarmResponseBodyResult
	SetStatusCode(v int32) *CreateHotelAlarmResponseBody
	GetStatusCode() *int32
}

type CreateHotelAlarmResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	Message    *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43***86881
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*CreateHotelAlarmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateHotelAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *CreateHotelAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHotelAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHotelAlarmResponseBody) GetResult() []*CreateHotelAlarmResponseBodyResult {
	return s.Result
}

func (s *CreateHotelAlarmResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHotelAlarmResponseBody) SetExtentions(v map[string]interface{}) *CreateHotelAlarmResponseBody {
	s.Extentions = v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetMessage(v string) *CreateHotelAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetRequestId(v string) *CreateHotelAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetResult(v []*CreateHotelAlarmResponseBodyResult) *CreateHotelAlarmResponseBody {
	s.Result = v
	return s
}

func (s *CreateHotelAlarmResponseBody) SetStatusCode(v int32) *CreateHotelAlarmResponseBody {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelAlarmResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHotelAlarmResponseBodyResult struct {
	// example:
	//
	// 94
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// example:
	//
	// Pvk***TA==
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	FailMsg      *string `json:"FailMsg,omitempty" xml:"FailMsg,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// mg***Qd
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s CreateHotelAlarmResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponseBodyResult) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *CreateHotelAlarmResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *CreateHotelAlarmResponseBodyResult) GetFailMsg() *string {
	return s.FailMsg
}

func (s *CreateHotelAlarmResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *CreateHotelAlarmResponseBodyResult) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *CreateHotelAlarmResponseBodyResult) SetAlarmId(v int64) *CreateHotelAlarmResponseBodyResult {
	s.AlarmId = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetDeviceOpenId(v string) *CreateHotelAlarmResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetFailMsg(v string) *CreateHotelAlarmResponseBodyResult {
	s.FailMsg = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetRoomNo(v string) *CreateHotelAlarmResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) SetUserOpenId(v string) *CreateHotelAlarmResponseBodyResult {
	s.UserOpenId = &v
	return s
}

func (s *CreateHotelAlarmResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
