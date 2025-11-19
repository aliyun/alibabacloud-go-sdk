// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartDesktopResponseBody
	GetCode() *string
	SetRequestId(v string) *StartDesktopResponseBody
	GetRequestId() *string
	SetTraceId(v string) *StartDesktopResponseBody
	GetTraceId() *string
}

type StartDesktopResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StartDesktopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartDesktopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDesktopResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *StartDesktopResponseBody) SetCode(v string) *StartDesktopResponseBody {
	s.Code = &v
	return s
}

func (s *StartDesktopResponseBody) SetRequestId(v string) *StartDesktopResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDesktopResponseBody) SetTraceId(v string) *StartDesktopResponseBody {
	s.TraceId = &v
	return s
}

func (s *StartDesktopResponseBody) Validate() error {
	return dara.Validate(s)
}
