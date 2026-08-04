// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryCreateAccountTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchQueryCreateAccountTraceResponseBody
	GetCode() *string
	SetMessage(v string) *BatchQueryCreateAccountTraceResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchQueryCreateAccountTraceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchQueryCreateAccountTraceResponseBody
	GetSuccess() *bool
	SetTraces(v []*BatchQueryCreateAccountTraceResponseBodyTraces) *BatchQueryCreateAccountTraceResponseBody
	GetTraces() []*BatchQueryCreateAccountTraceResponseBodyTraces
}

type BatchQueryCreateAccountTraceResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Traces    []*BatchQueryCreateAccountTraceResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s BatchQueryCreateAccountTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryCreateAccountTraceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryCreateAccountTraceResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryCreateAccountTraceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchQueryCreateAccountTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryCreateAccountTraceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchQueryCreateAccountTraceResponseBody) GetTraces() []*BatchQueryCreateAccountTraceResponseBodyTraces {
	return s.Traces
}

func (s *BatchQueryCreateAccountTraceResponseBody) SetCode(v string) *BatchQueryCreateAccountTraceResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBody) SetMessage(v string) *BatchQueryCreateAccountTraceResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBody) SetRequestId(v string) *BatchQueryCreateAccountTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBody) SetSuccess(v bool) *BatchQueryCreateAccountTraceResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBody) SetTraces(v []*BatchQueryCreateAccountTraceResponseBodyTraces) *BatchQueryCreateAccountTraceResponseBody {
	s.Traces = v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBody) Validate() error {
	if s.Traces != nil {
		for _, item := range s.Traces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryCreateAccountTraceResponseBodyTraces struct {
	NowLoginEmail *string `json:"NowLoginEmail,omitempty" xml:"NowLoginEmail,omitempty"`
	Pk            *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceNo       *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s BatchQueryCreateAccountTraceResponseBodyTraces) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryCreateAccountTraceResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) GetNowLoginEmail() *string {
	return s.NowLoginEmail
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) GetPk() *string {
	return s.Pk
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) GetStatus() *string {
	return s.Status
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) GetTraceNo() *string {
	return s.TraceNo
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) SetNowLoginEmail(v string) *BatchQueryCreateAccountTraceResponseBodyTraces {
	s.NowLoginEmail = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) SetPk(v string) *BatchQueryCreateAccountTraceResponseBodyTraces {
	s.Pk = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) SetStatus(v string) *BatchQueryCreateAccountTraceResponseBodyTraces {
	s.Status = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) SetTraceNo(v string) *BatchQueryCreateAccountTraceResponseBodyTraces {
	s.TraceNo = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponseBodyTraces) Validate() error {
	return dara.Validate(s)
}
