// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceBasicInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceBasicInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceBasicInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetDeviceBasicInfoResponseBodyResult) *GetDeviceBasicInfoResponseBody
	GetResult() *GetDeviceBasicInfoResponseBodyResult
}

type GetDeviceBasicInfoResponseBody struct {
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
	// 0EC7*726E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceBasicInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceBasicInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceBasicInfoResponseBody) GetResult() *GetDeviceBasicInfoResponseBodyResult {
	return s.Result
}

func (s *GetDeviceBasicInfoResponseBody) SetCode(v int32) *GetDeviceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetMessage(v string) *GetDeviceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetRequestId(v string) *GetDeviceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) SetResult(v *GetDeviceBasicInfoResponseBodyResult) *GetDeviceBasicInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceBasicInfoResponseBodyResult struct {
	// example:
	//
	// 2.0.3
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// b4:xx:xx:xx:65:2b
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 我的设备
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1200xxx048
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetDeviceBasicInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoResponseBodyResult) GetFirmwareVersion() *string {
	return s.FirmwareVersion
}

func (s *GetDeviceBasicInfoResponseBodyResult) GetMac() *string {
	return s.Mac
}

func (s *GetDeviceBasicInfoResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetDeviceBasicInfoResponseBodyResult) GetSn() *string {
	return s.Sn
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetFirmwareVersion(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetMac(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetName(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) SetSn(v string) *GetDeviceBasicInfoResponseBodyResult {
	s.Sn = &v
	return s
}

func (s *GetDeviceBasicInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
