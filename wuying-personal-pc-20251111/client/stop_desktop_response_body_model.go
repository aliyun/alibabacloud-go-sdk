// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopDesktopResponseBody
	GetCode() *string
	SetRequestId(v string) *StopDesktopResponseBody
	GetRequestId() *string
	SetTraceId(v string) *StopDesktopResponseBody
	GetTraceId() *string
}

type StopDesktopResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StopDesktopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopDesktopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDesktopResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *StopDesktopResponseBody) SetCode(v string) *StopDesktopResponseBody {
	s.Code = &v
	return s
}

func (s *StopDesktopResponseBody) SetRequestId(v string) *StopDesktopResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDesktopResponseBody) SetTraceId(v string) *StopDesktopResponseBody {
	s.TraceId = &v
	return s
}

func (s *StopDesktopResponseBody) Validate() error {
	return dara.Validate(s)
}
