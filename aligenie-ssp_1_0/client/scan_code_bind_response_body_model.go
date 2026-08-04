// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanCodeBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ScanCodeBindResponseBody
	GetCode() *int32
	SetMessage(v string) *ScanCodeBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *ScanCodeBindResponseBody
	GetRequestId() *string
	SetResult(v *ScanCodeBindResponseBodyResult) *ScanCodeBindResponseBody
	GetResult() *ScanCodeBindResponseBodyResult
}

type ScanCodeBindResponseBody struct {
	// The returned error code. A value of 200 indicates that the invocation succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Result message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 686DF82F-45C4-7DF7-8B67-27B91CFD63A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed information returned.
	Result *ScanCodeBindResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ScanCodeBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindResponseBody) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScanCodeBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScanCodeBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanCodeBindResponseBody) GetResult() *ScanCodeBindResponseBodyResult {
	return s.Result
}

func (s *ScanCodeBindResponseBody) SetCode(v int32) *ScanCodeBindResponseBody {
	s.Code = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetMessage(v string) *ScanCodeBindResponseBody {
	s.Message = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetRequestId(v string) *ScanCodeBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanCodeBindResponseBody) SetResult(v *ScanCodeBindResponseBodyResult) *ScanCodeBindResponseBody {
	s.Result = v
	return s
}

func (s *ScanCodeBindResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScanCodeBindResponseBodyResult struct {
	// Product group
	//
	// example:
	//
	// X1
	BizGroup *string `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	// Product categorization
	//
	// example:
	//
	// AILABS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// A963*0158
	//
	// example:
	//
	// 设备OpenId
	DeviceOpenId *string `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	// DAFE****ce3ej=
	//
	// example:
	//
	// 用户OpenId
	UserOpenId *string `json:"UserOpenId,omitempty" xml:"UserOpenId,omitempty"`
}

func (s ScanCodeBindResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ScanCodeBindResponseBodyResult) GetBizGroup() *string {
	return s.BizGroup
}

func (s *ScanCodeBindResponseBodyResult) GetBizType() *string {
	return s.BizType
}

func (s *ScanCodeBindResponseBodyResult) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *ScanCodeBindResponseBodyResult) GetUserOpenId() *string {
	return s.UserOpenId
}

func (s *ScanCodeBindResponseBodyResult) SetBizGroup(v string) *ScanCodeBindResponseBodyResult {
	s.BizGroup = &v
	return s
}

func (s *ScanCodeBindResponseBodyResult) SetBizType(v string) *ScanCodeBindResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *ScanCodeBindResponseBodyResult) SetDeviceOpenId(v string) *ScanCodeBindResponseBodyResult {
	s.DeviceOpenId = &v
	return s
}

func (s *ScanCodeBindResponseBodyResult) SetUserOpenId(v string) *ScanCodeBindResponseBodyResult {
	s.UserOpenId = &v
	return s
}

func (s *ScanCodeBindResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
