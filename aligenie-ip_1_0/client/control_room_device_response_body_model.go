// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlRoomDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ControlRoomDeviceResponseBody
	GetCode() *int32
	SetMessage(v string) *ControlRoomDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ControlRoomDeviceResponseBody
	GetRequestId() *string
	SetResult(v *ControlRoomDeviceResponseBodyResult) *ControlRoomDeviceResponseBody
	GetResult() *ControlRoomDeviceResponseBodyResult
}

type ControlRoomDeviceResponseBody struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ControlRoomDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ControlRoomDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ControlRoomDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ControlRoomDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ControlRoomDeviceResponseBody) GetResult() *ControlRoomDeviceResponseBodyResult {
	return s.Result
}

func (s *ControlRoomDeviceResponseBody) SetCode(v int32) *ControlRoomDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetMessage(v string) *ControlRoomDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetRequestId(v string) *ControlRoomDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ControlRoomDeviceResponseBody) SetResult(v *ControlRoomDeviceResponseBodyResult) *ControlRoomDeviceResponseBody {
	s.Result = v
	return s
}

func (s *ControlRoomDeviceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ControlRoomDeviceResponseBodyResult struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ControlRoomDeviceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *ControlRoomDeviceResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ControlRoomDeviceResponseBodyResult) SetMessage(v string) *ControlRoomDeviceResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ControlRoomDeviceResponseBodyResult) SetStatus(v int32) *ControlRoomDeviceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ControlRoomDeviceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
