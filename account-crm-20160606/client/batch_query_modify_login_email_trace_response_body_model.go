// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryModifyLoginEmailTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchQueryModifyLoginEmailTraceResponseBody
	GetCode() *string
	SetMessage(v string) *BatchQueryModifyLoginEmailTraceResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchQueryModifyLoginEmailTraceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *BatchQueryModifyLoginEmailTraceResponseBody
	GetSuccess() *string
	SetTraces(v []*BatchQueryModifyLoginEmailTraceResponseBodyTraces) *BatchQueryModifyLoginEmailTraceResponseBody
	GetTraces() []*BatchQueryModifyLoginEmailTraceResponseBodyTraces
}

type BatchQueryModifyLoginEmailTraceResponseBody struct {
	Code      *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Traces    []*BatchQueryModifyLoginEmailTraceResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s BatchQueryModifyLoginEmailTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryModifyLoginEmailTraceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) GetTraces() []*BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	return s.Traces
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) SetCode(v string) *BatchQueryModifyLoginEmailTraceResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) SetMessage(v string) *BatchQueryModifyLoginEmailTraceResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) SetRequestId(v string) *BatchQueryModifyLoginEmailTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) SetSuccess(v string) *BatchQueryModifyLoginEmailTraceResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) SetTraces(v []*BatchQueryModifyLoginEmailTraceResponseBodyTraces) *BatchQueryModifyLoginEmailTraceResponseBody {
	s.Traces = v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBody) Validate() error {
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

type BatchQueryModifyLoginEmailTraceResponseBodyTraces struct {
	NewLoginEmail *string `json:"NewLoginEmail,omitempty" xml:"NewLoginEmail,omitempty"`
	NowLoginEmail *string `json:"NowLoginEmail,omitempty" xml:"NowLoginEmail,omitempty"`
	OldLoginEmail *string `json:"OldLoginEmail,omitempty" xml:"OldLoginEmail,omitempty"`
	Pk            *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceNo       *string `json:"TraceNo,omitempty" xml:"TraceNo,omitempty"`
}

func (s BatchQueryModifyLoginEmailTraceResponseBodyTraces) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryModifyLoginEmailTraceResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetNewLoginEmail() *string {
	return s.NewLoginEmail
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetNowLoginEmail() *string {
	return s.NowLoginEmail
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetOldLoginEmail() *string {
	return s.OldLoginEmail
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetPk() *string {
	return s.Pk
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetStatus() *string {
	return s.Status
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) GetTraceNo() *string {
	return s.TraceNo
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetNewLoginEmail(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.NewLoginEmail = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetNowLoginEmail(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.NowLoginEmail = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetOldLoginEmail(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.OldLoginEmail = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetPk(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.Pk = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetStatus(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.Status = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) SetTraceNo(v string) *BatchQueryModifyLoginEmailTraceResponseBodyTraces {
	s.TraceNo = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponseBodyTraces) Validate() error {
	return dara.Validate(s)
}
