// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopLivyComputeResponseBody
	GetCode() *string
	SetMessage(v string) *StopLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopLivyComputeResponseBody
	GetRequestId() *string
}

type StopLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *StopLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLivyComputeResponseBody) SetCode(v string) *StopLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *StopLivyComputeResponseBody) SetMessage(v string) *StopLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *StopLivyComputeResponseBody) SetRequestId(v string) *StopLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLivyComputeResponseBody) Validate() error {
	return dara.Validate(s)
}
