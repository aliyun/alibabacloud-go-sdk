// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdHocTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopAdHocTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopAdHocTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopAdHocTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopAdHocTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopAdHocTaskResponseBody
	GetSuccess() *bool
}

type StopAdHocTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopAdHocTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAdHocTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopAdHocTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopAdHocTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopAdHocTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopAdHocTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAdHocTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopAdHocTaskResponseBody) SetCode(v string) *StopAdHocTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopAdHocTaskResponseBody) SetHttpStatusCode(v int32) *StopAdHocTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopAdHocTaskResponseBody) SetMessage(v string) *StopAdHocTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopAdHocTaskResponseBody) SetRequestId(v string) *StopAdHocTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAdHocTaskResponseBody) SetSuccess(v bool) *StopAdHocTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopAdHocTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
