// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelRoomDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelRoomDeviceResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelRoomDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelRoomDeviceResponseBody
	GetRequestId() *string
	SetResult(v []*GetHotelRoomDeviceResponseBodyResult) *GetHotelRoomDeviceResponseBody
	GetResult() []*GetHotelRoomDeviceResponseBodyResult
}

type GetHotelRoomDeviceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// vrehvuifdsgrts
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelRoomDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelRoomDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelRoomDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelRoomDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelRoomDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelRoomDeviceResponseBody) GetResult() []*GetHotelRoomDeviceResponseBodyResult {
	return s.Result
}

func (s *GetHotelRoomDeviceResponseBody) SetCode(v int32) *GetHotelRoomDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetMessage(v string) *GetHotelRoomDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetRequestId(v string) *GetHotelRoomDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) SetResult(v []*GetHotelRoomDeviceResponseBodyResult) *GetHotelRoomDeviceResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelRoomDeviceResponseBody) Validate() error {
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

type GetHotelRoomDeviceResponseBodyResult struct {
	// example:
	//
	// 1.0.0-release
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// aa:aa:aa:aa:aa:aa
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// dsfdsfrgreg
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetHotelRoomDeviceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelRoomDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetFirmwareVersion() *string {
	return s.FirmwareVersion
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetMac() *string {
	return s.Mac
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetOnlineStatus() *int32 {
	return s.OnlineStatus
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *GetHotelRoomDeviceResponseBodyResult) GetSn() *string {
	return s.Sn
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetFirmwareVersion(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetHotelId(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetMac(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetOnlineStatus(v int32) *GetHotelRoomDeviceResponseBodyResult {
	s.OnlineStatus = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetRoomNo(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) SetSn(v string) *GetHotelRoomDeviceResponseBodyResult {
	s.Sn = &v
	return s
}

func (s *GetHotelRoomDeviceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
