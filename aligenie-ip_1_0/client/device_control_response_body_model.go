// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeviceControlResponseBody
	GetCode() *int32
	SetMessage(v string) *DeviceControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeviceControlResponseBody
	GetRequestId() *string
	SetResult(v *DeviceControlResponseBodyResult) *DeviceControlResponseBody
	GetResult() *DeviceControlResponseBodyResult
}

type DeviceControlResponseBody struct {
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
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeviceControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeviceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeviceControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeviceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeviceControlResponseBody) GetResult() *DeviceControlResponseBodyResult {
	return s.Result
}

func (s *DeviceControlResponseBody) SetCode(v int32) *DeviceControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceControlResponseBody) SetMessage(v string) *DeviceControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceControlResponseBody) SetRequestId(v string) *DeviceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceControlResponseBody) SetResult(v *DeviceControlResponseBodyResult) *DeviceControlResponseBody {
	s.Result = v
	return s
}

func (s *DeviceControlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeviceControlResponseBodyResult struct {
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeviceControlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DeviceControlResponseBodyResult) SetStatus(v string) *DeviceControlResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeviceControlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
