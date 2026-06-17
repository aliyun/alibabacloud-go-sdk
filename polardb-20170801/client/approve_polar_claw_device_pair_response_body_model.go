// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePolarClawDevicePairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ApprovePolarClawDevicePairResponseBody
	GetApplicationId() *string
	SetCode(v int32) *ApprovePolarClawDevicePairResponseBody
	GetCode() *int32
	SetDevice(v *ApprovePolarClawDevicePairResponseBodyDevice) *ApprovePolarClawDevicePairResponseBody
	GetDevice() *ApprovePolarClawDevicePairResponseBodyDevice
	SetMessage(v string) *ApprovePolarClawDevicePairResponseBody
	GetMessage() *string
	SetPairRequestId(v string) *ApprovePolarClawDevicePairResponseBody
	GetPairRequestId() *string
	SetRequestId(v string) *ApprovePolarClawDevicePairResponseBody
	GetRequestId() *string
}

type ApprovePolarClawDevicePairResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-********************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paired device information.
	Device *ApprovePolarClawDevicePairResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pairing request ID.
	//
	// example:
	//
	// req-abc-123
	PairRequestId *string `json:"PairRequestId,omitempty" xml:"PairRequestId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25C70FF3-D49B-594D-BECE-0DE2BA1D8BBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApprovePolarClawDevicePairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApprovePolarClawDevicePairResponseBody) GoString() string {
	return s.String()
}

func (s *ApprovePolarClawDevicePairResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ApprovePolarClawDevicePairResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ApprovePolarClawDevicePairResponseBody) GetDevice() *ApprovePolarClawDevicePairResponseBodyDevice {
	return s.Device
}

func (s *ApprovePolarClawDevicePairResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApprovePolarClawDevicePairResponseBody) GetPairRequestId() *string {
	return s.PairRequestId
}

func (s *ApprovePolarClawDevicePairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApprovePolarClawDevicePairResponseBody) SetApplicationId(v string) *ApprovePolarClawDevicePairResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) SetCode(v int32) *ApprovePolarClawDevicePairResponseBody {
	s.Code = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) SetDevice(v *ApprovePolarClawDevicePairResponseBodyDevice) *ApprovePolarClawDevicePairResponseBody {
	s.Device = v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) SetMessage(v string) *ApprovePolarClawDevicePairResponseBody {
	s.Message = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) SetPairRequestId(v string) *ApprovePolarClawDevicePairResponseBody {
	s.PairRequestId = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) SetRequestId(v string) *ApprovePolarClawDevicePairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBody) Validate() error {
	if s.Device != nil {
		if err := s.Device.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApprovePolarClawDevicePairResponseBodyDevice struct {
	// The creation time of the device pairing, in Unix milliseconds.
	//
	// example:
	//
	// 1778662316663
	CreatedAtMs *int64 `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	// The device family.
	//
	// example:
	//
	// server
	DeviceFamily *string `json:"DeviceFamily,omitempty" xml:"DeviceFamily,omitempty"`
	// The unique ID of the device.
	//
	// example:
	//
	// device-l7rdl36iz6op66zf
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The display name of the device.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The operating system of the device.
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The role of the device.
	//
	// example:
	//
	// operator
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// A list of permission scopes.
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
}

func (s ApprovePolarClawDevicePairResponseBodyDevice) String() string {
	return dara.Prettify(s)
}

func (s ApprovePolarClawDevicePairResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetDeviceFamily() *string {
	return s.DeviceFamily
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetPlatform() *string {
	return s.Platform
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetRole() *string {
	return s.Role
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) GetScopes() []*string {
	return s.Scopes
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetCreatedAtMs(v int64) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.CreatedAtMs = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetDeviceFamily(v string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.DeviceFamily = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetDeviceId(v string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.DeviceId = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetDisplayName(v string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.DisplayName = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetPlatform(v string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.Platform = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetRole(v string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.Role = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) SetScopes(v []*string) *ApprovePolarClawDevicePairResponseBodyDevice {
	s.Scopes = v
	return s
}

func (s *ApprovePolarClawDevicePairResponseBodyDevice) Validate() error {
	return dara.Validate(s)
}
