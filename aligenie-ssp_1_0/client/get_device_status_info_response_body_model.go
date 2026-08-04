// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceStatusInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceStatusInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceStatusInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetDeviceStatusInfoResponseBodyResult) *GetDeviceStatusInfoResponseBody
	GetResult() *GetDeviceStatusInfoResponseBodyResult
}

type GetDeviceStatusInfoResponseBody struct {
	// The error code returned. A value of 200 indicates that the call succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return result of invoking this API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Detailed information returned.
	Result *GetDeviceStatusInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceStatusInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceStatusInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceStatusInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceStatusInfoResponseBody) GetResult() *GetDeviceStatusInfoResponseBodyResult {
	return s.Result
}

func (s *GetDeviceStatusInfoResponseBody) SetCode(v int32) *GetDeviceStatusInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetMessage(v string) *GetDeviceStatusInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetRequestId(v string) *GetDeviceStatusInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) SetResult(v *GetDeviceStatusInfoResponseBodyResult) *GetDeviceStatusInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceStatusInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceStatusInfoResponseBodyResult struct {
	// Indicates whether the device is online. A value of 0 means offline, and 1 means online.
	//
	// example:
	//
	// 1
	Online *int32 `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s GetDeviceStatusInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoResponseBodyResult) GetOnline() *int32 {
	return s.Online
}

func (s *GetDeviceStatusInfoResponseBodyResult) SetOnline(v int32) *GetDeviceStatusInfoResponseBodyResult {
	s.Online = &v
	return s
}

func (s *GetDeviceStatusInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
