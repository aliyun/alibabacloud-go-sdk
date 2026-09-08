// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateAgentDataSemanticsProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGenerateAgentDataSemanticsProgressResponseBody
	GetCode() *string
	SetData(v *GetGenerateAgentDataSemanticsProgressResponseBodyData) *GetGenerateAgentDataSemanticsProgressResponseBody
	GetData() *GetGenerateAgentDataSemanticsProgressResponseBodyData
	SetMessage(v string) *GetGenerateAgentDataSemanticsProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGenerateAgentDataSemanticsProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGenerateAgentDataSemanticsProgressResponseBody
	GetSuccess() *bool
}

type GetGenerateAgentDataSemanticsProgressResponseBody struct {
	// The response code of the operation.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The generation progress details. When the initial generation is complete, a full snapshot of the current generation round is returned. When regeneration is complete, the current Metrics, Joins, Examples, and new Text are returned. To discard a regeneration, first call Get to retrieve the current official version, and then call Save with the four types of content unchanged to idempotently clean up temporary results.
	Data *GetGenerateAgentDataSemanticsProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. If the request fails, an error message is returned.
	//
	// example:
	//
	// Data semantics generation task not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 5DAF96FB-A4DF-548C-B8A1-F2A8D2F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGenerateAgentDataSemanticsProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateAgentDataSemanticsProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) GetData() *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	return s.Data
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) SetCode(v string) *GetGenerateAgentDataSemanticsProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) SetData(v *GetGenerateAgentDataSemanticsProgressResponseBodyData) *GetGenerateAgentDataSemanticsProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) SetMessage(v string) *GetGenerateAgentDataSemanticsProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) SetRequestId(v string) *GetGenerateAgentDataSemanticsProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) SetSuccess(v bool) *GetGenerateAgentDataSemanticsProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGenerateAgentDataSemanticsProgressResponseBodyData struct {
	// The error code returned when the generation task fails.
	//
	// example:
	//
	// DataSemanticsGenerateFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the generation task fails.
	//
	// example:
	//
	// Failed to generate data semantics. Please retry later.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The SQL example knowledge returned when the generation is complete. A maximum of 50 items can be returned.
	Examples []*AgentDataSemanticsExample `json:"Examples,omitempty" xml:"Examples,omitempty" type:"Repeated"`
	// The data association knowledge returned when the generation is complete. A maximum of 100 items can be returned.
	Joins []*AgentDataSemanticsJoin `json:"Joins,omitempty" xml:"Joins,omitempty" type:"Repeated"`
	// The SQL expression knowledge returned when the generation is complete. A maximum of 100 items can be returned.
	Metrics []*AgentDataSemanticsMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The four-phase stage progress. This parameter may not be returned when the overall generation is complete.
	Progress []*AgentDataSemanticsStageProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Repeated"`
	// The current overall stage.
	//
	// example:
	//
	// GENERATE
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The Markdown text knowledge returned when the generation is complete.
	Text *AgentDataSemanticsText `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetGenerateAgentDataSemanticsProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateAgentDataSemanticsProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetExamples() []*AgentDataSemanticsExample {
	return s.Examples
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetJoins() []*AgentDataSemanticsJoin {
	return s.Joins
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetMetrics() []*AgentDataSemanticsMetric {
	return s.Metrics
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetProgress() []*AgentDataSemanticsStageProgress {
	return s.Progress
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetStage() *string {
	return s.Stage
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) GetText() *AgentDataSemanticsText {
	return s.Text
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetErrorCode(v string) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetErrorMessage(v string) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetExamples(v []*AgentDataSemanticsExample) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Examples = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetJoins(v []*AgentDataSemanticsJoin) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Joins = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetMetrics(v []*AgentDataSemanticsMetric) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetProgress(v []*AgentDataSemanticsStageProgress) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Progress = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetStage(v string) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Stage = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) SetText(v *AgentDataSemanticsText) *GetGenerateAgentDataSemanticsProgressResponseBodyData {
	s.Text = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponseBodyData) Validate() error {
	if s.Examples != nil {
		for _, item := range s.Examples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Joins != nil {
		for _, item := range s.Joins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Progress != nil {
		for _, item := range s.Progress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Text != nil {
		if err := s.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}
