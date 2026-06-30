// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelGeneratedResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLabelGeneratedResultResponseBody
	GetCode() *string
	SetData(v *GetLabelGeneratedResultResponseBodyData) *GetLabelGeneratedResultResponseBody
	GetData() *GetLabelGeneratedResultResponseBodyData
	SetMessage(v string) *GetLabelGeneratedResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLabelGeneratedResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLabelGeneratedResultResponseBody
	GetSuccess() *bool
}

type GetLabelGeneratedResultResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetLabelGeneratedResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLabelGeneratedResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLabelGeneratedResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelGeneratedResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLabelGeneratedResultResponseBody) GetData() *GetLabelGeneratedResultResponseBodyData {
	return s.Data
}

func (s *GetLabelGeneratedResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLabelGeneratedResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLabelGeneratedResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLabelGeneratedResultResponseBody) SetCode(v string) *GetLabelGeneratedResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBody) SetData(v *GetLabelGeneratedResultResponseBodyData) *GetLabelGeneratedResultResponseBody {
	s.Data = v
	return s
}

func (s *GetLabelGeneratedResultResponseBody) SetMessage(v string) *GetLabelGeneratedResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBody) SetRequestId(v string) *GetLabelGeneratedResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBody) SetSuccess(v bool) *GetLabelGeneratedResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLabelGeneratedResultResponseBodyData struct {
	// The number of input tokens for the LLM.
	//
	// example:
	//
	// 7371
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The number of LLM calls.
	//
	// example:
	//
	// 4
	LlmCallNum *int32 `json:"LlmCallNum,omitempty" xml:"LlmCallNum,omitempty"`
	// The number of output tokens for the LLM.
	//
	// example:
	//
	// 355
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The pre-signed download URL of the result file.
	//
	// example:
	//
	// https://sca-eas-mining.oss-cn-beijing.aliyuncs.com/xxx.xlsx?Expires=...
	ResultFileUrl *string `json:"ResultFileUrl,omitempty" xml:"ResultFileUrl,omitempty"`
	// The ID of the generation task.
	//
	// example:
	//
	// 20260616-4955F615-A74E-171E-86ED-080F60C72EC9
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetLabelGeneratedResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLabelGeneratedResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLabelGeneratedResultResponseBodyData) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetLabelGeneratedResultResponseBodyData) GetLlmCallNum() *int32 {
	return s.LlmCallNum
}

func (s *GetLabelGeneratedResultResponseBodyData) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetLabelGeneratedResultResponseBodyData) GetResultFileUrl() *string {
	return s.ResultFileUrl
}

func (s *GetLabelGeneratedResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetLabelGeneratedResultResponseBodyData) SetInputTokens(v int32) *GetLabelGeneratedResultResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBodyData) SetLlmCallNum(v int32) *GetLabelGeneratedResultResponseBodyData {
	s.LlmCallNum = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBodyData) SetOutputTokens(v int32) *GetLabelGeneratedResultResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBodyData) SetResultFileUrl(v string) *GetLabelGeneratedResultResponseBodyData {
	s.ResultFileUrl = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBodyData) SetTaskId(v string) *GetLabelGeneratedResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetLabelGeneratedResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
