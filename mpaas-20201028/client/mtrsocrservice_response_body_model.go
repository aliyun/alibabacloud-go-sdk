// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMTRSOCRServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MTRSOCRServiceResponseBody
	GetCode() *string
	SetMsg(v string) *MTRSOCRServiceResponseBody
	GetMsg() *string
	SetRequestId(v string) *MTRSOCRServiceResponseBody
	GetRequestId() *string
	SetResult(v string) *MTRSOCRServiceResponseBody
	GetResult() *string
	SetStatus(v bool) *MTRSOCRServiceResponseBody
	GetStatus() *bool
	SetTraceId(v string) *MTRSOCRServiceResponseBody
	GetTraceId() *string
}

type MTRSOCRServiceResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg  *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// D9B3C4E7-BEC7-1E2C-86A3-EA985B4FFD73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"aa":"ss"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true/false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 777799aa
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s MTRSOCRServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MTRSOCRServiceResponseBody) GoString() string {
	return s.String()
}

func (s *MTRSOCRServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MTRSOCRServiceResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *MTRSOCRServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MTRSOCRServiceResponseBody) GetResult() *string {
	return s.Result
}

func (s *MTRSOCRServiceResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *MTRSOCRServiceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MTRSOCRServiceResponseBody) SetCode(v string) *MTRSOCRServiceResponseBody {
	s.Code = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) SetMsg(v string) *MTRSOCRServiceResponseBody {
	s.Msg = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) SetRequestId(v string) *MTRSOCRServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) SetResult(v string) *MTRSOCRServiceResponseBody {
	s.Result = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) SetStatus(v bool) *MTRSOCRServiceResponseBody {
	s.Status = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) SetTraceId(v string) *MTRSOCRServiceResponseBody {
	s.TraceId = &v
	return s
}

func (s *MTRSOCRServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
