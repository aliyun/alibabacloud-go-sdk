// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetRecordingRuleResponseBody
	GetCode() *int32
	SetData(v string) *GetRecordingRuleResponseBody
	GetData() *string
	SetMessage(v string) *GetRecordingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRecordingRuleResponseBody
	GetRequestId() *string
}

type GetRecordingRuleResponseBody struct {
	// The status code. The status code 200 indicates a successful request, whereas others indicate a failed request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The recording rule.
	//
	// example:
	//
	// --- groups: - name: "recording_demo"   rules:   - expr: "sum(jvm_memory_max_bytes)"     record: "rate_coredns_demo"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetRecordingRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *GetRecordingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRecordingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordingRuleResponseBody) SetCode(v int32) *GetRecordingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordingRuleResponseBody) SetData(v string) *GetRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *GetRecordingRuleResponseBody) SetMessage(v string) *GetRecordingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordingRuleResponseBody) SetRequestId(v string) *GetRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
