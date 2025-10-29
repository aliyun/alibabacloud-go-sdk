// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmonMonitorDataResponseBody
	GetCode() *string
	SetMessage(v string) *GetEmonMonitorDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEmonMonitorDataResponseBody
	GetRequestId() *string
	SetResult(v []*GetEmonMonitorDataResponseBodyResult) *GetEmonMonitorDataResponseBody
	GetResult() []*GetEmonMonitorDataResponseBodyResult
	SetSuccess(v bool) *GetEmonMonitorDataResponseBody
	GetSuccess() *bool
}

type GetEmonMonitorDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2D184B55-FA51-43F7-A1EF-E68A0545****
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetEmonMonitorDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmonMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEmonMonitorDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEmonMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmonMonitorDataResponseBody) GetResult() []*GetEmonMonitorDataResponseBodyResult {
	return s.Result
}

func (s *GetEmonMonitorDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEmonMonitorDataResponseBody) SetCode(v string) *GetEmonMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetMessage(v string) *GetEmonMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetRequestId(v string) *GetEmonMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetResult(v []*GetEmonMonitorDataResponseBodyResult) *GetEmonMonitorDataResponseBody {
	s.Result = v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetSuccess(v bool) *GetEmonMonitorDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEmonMonitorDataResponseBodyResult struct {
	// example:
	//
	// { "1586249280": 465.1980465119913, "1586249300": 213.45243650423305 }
	Dps map[string]interface{} `json:"dps,omitempty" xml:"dps,omitempty"`
	// example:
	//
	// 1.0
	Integrity *float32 `json:"integrity,omitempty" xml:"integrity,omitempty"`
	// example:
	//
	// 1522127381471
	MessageWatermark *int64 `json:"messageWatermark,omitempty" xml:"messageWatermark,omitempty"`
	// example:
	//
	// elasticbuild.elasticsearch.source.total_doc_count
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// example:
	//
	// 10
	Summary *float32 `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// {"taskName":"et-xxx","userId":"123456"}
	Tags map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetEmonMonitorDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetEmonMonitorDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponseBodyResult) GetDps() map[string]interface{} {
	return s.Dps
}

func (s *GetEmonMonitorDataResponseBodyResult) GetIntegrity() *float32 {
	return s.Integrity
}

func (s *GetEmonMonitorDataResponseBodyResult) GetMessageWatermark() *int64 {
	return s.MessageWatermark
}

func (s *GetEmonMonitorDataResponseBodyResult) GetMetric() *string {
	return s.Metric
}

func (s *GetEmonMonitorDataResponseBodyResult) GetSummary() *float32 {
	return s.Summary
}

func (s *GetEmonMonitorDataResponseBodyResult) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetEmonMonitorDataResponseBodyResult) SetDps(v map[string]interface{}) *GetEmonMonitorDataResponseBodyResult {
	s.Dps = v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetIntegrity(v float32) *GetEmonMonitorDataResponseBodyResult {
	s.Integrity = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetMessageWatermark(v int64) *GetEmonMonitorDataResponseBodyResult {
	s.MessageWatermark = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetMetric(v string) *GetEmonMonitorDataResponseBodyResult {
	s.Metric = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetSummary(v float32) *GetEmonMonitorDataResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetTags(v map[string]interface{}) *GetEmonMonitorDataResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
