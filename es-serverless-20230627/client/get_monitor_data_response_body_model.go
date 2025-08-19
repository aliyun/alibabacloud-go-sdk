// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMonitorDataResponseBody
	GetCode() *string
	SetMessage(v string) *GetMonitorDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMonitorDataResponseBody
	GetRequestId() *string
	SetResult(v []*GetMonitorDataResponseBodyResult) *GetMonitorDataResponseBody
	GetResult() []*GetMonitorDataResponseBodyResult
	SetSuccess(v bool) *GetMonitorDataResponseBody
	GetSuccess() *bool
}

type GetMonitorDataResponseBody struct {
	// example:
	//
	// InternalServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// internal server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetMonitorDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMonitorDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMonitorDataResponseBody) GetResult() []*GetMonitorDataResponseBodyResult {
	return s.Result
}

func (s *GetMonitorDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMonitorDataResponseBody) SetCode(v string) *GetMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetMessage(v string) *GetMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetRequestId(v string) *GetMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetResult(v []*GetMonitorDataResponseBodyResult) *GetMonitorDataResponseBody {
	s.Result = v
	return s
}

func (s *GetMonitorDataResponseBody) SetSuccess(v bool) *GetMonitorDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetMonitorDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMonitorDataResponseBodyResult struct {
	// example:
	//
	// {
	//
	//     "1689480600":28676235.104761902
	//
	// }
	Dps map[string]interface{} `json:"dps,omitempty" xml:"dps,omitempty"`
	// example:
	//
	// 1
	Integrity *float32 `json:"integrity,omitempty" xml:"integrity,omitempty"`
	// example:
	//
	// 1689566839447
	MessageWatermark *int64 `json:"messageWatermark,omitempty" xml:"messageWatermark,omitempty"`
	// example:
	//
	// elasticsearch-server.logic_cpu.cpu
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// example:
	//
	// 172455913.18935508
	Summary *float32 `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// {
	//
	//                 "indexName":"test"
	//
	//             }
	Tags map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetMonitorDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponseBodyResult) GetDps() map[string]interface{} {
	return s.Dps
}

func (s *GetMonitorDataResponseBodyResult) GetIntegrity() *float32 {
	return s.Integrity
}

func (s *GetMonitorDataResponseBodyResult) GetMessageWatermark() *int64 {
	return s.MessageWatermark
}

func (s *GetMonitorDataResponseBodyResult) GetMetric() *string {
	return s.Metric
}

func (s *GetMonitorDataResponseBodyResult) GetSummary() *float32 {
	return s.Summary
}

func (s *GetMonitorDataResponseBodyResult) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetMonitorDataResponseBodyResult) SetDps(v map[string]interface{}) *GetMonitorDataResponseBodyResult {
	s.Dps = v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetIntegrity(v float32) *GetMonitorDataResponseBodyResult {
	s.Integrity = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetMessageWatermark(v int64) *GetMonitorDataResponseBodyResult {
	s.MessageWatermark = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetMetric(v string) *GetMonitorDataResponseBodyResult {
	s.Metric = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetSummary(v float32) *GetMonitorDataResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetTags(v map[string]interface{}) *GetMonitorDataResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetMonitorDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
