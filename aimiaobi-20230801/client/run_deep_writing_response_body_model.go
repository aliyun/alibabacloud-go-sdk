// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDeepWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunDeepWritingResponseBody
	GetCode() *string
	SetHeader(v *RunDeepWritingResponseBodyHeader) *RunDeepWritingResponseBody
	GetHeader() *RunDeepWritingResponseBodyHeader
	SetHttpStatusCode(v string) *RunDeepWritingResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunDeepWritingResponseBody
	GetMessage() *string
	SetPayload(v *RunDeepWritingResponseBodyPayload) *RunDeepWritingResponseBody
	GetPayload() *RunDeepWritingResponseBodyPayload
	SetRequestId(v string) *RunDeepWritingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunDeepWritingResponseBody
	GetSuccess() *bool
}

type RunDeepWritingResponseBody struct {
	// example:
	//
	// successful
	Code   *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Header *RunDeepWritingResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Payload *RunDeepWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 31AC01F1-88FB-5C4D-B6F5-E8BB136CD5A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunDeepWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunDeepWritingResponseBody) GetHeader() *RunDeepWritingResponseBodyHeader {
	return s.Header
}

func (s *RunDeepWritingResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunDeepWritingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunDeepWritingResponseBody) GetPayload() *RunDeepWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunDeepWritingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDeepWritingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunDeepWritingResponseBody) SetCode(v string) *RunDeepWritingResponseBody {
	s.Code = &v
	return s
}

func (s *RunDeepWritingResponseBody) SetHeader(v *RunDeepWritingResponseBodyHeader) *RunDeepWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunDeepWritingResponseBody) SetHttpStatusCode(v string) *RunDeepWritingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunDeepWritingResponseBody) SetMessage(v string) *RunDeepWritingResponseBody {
	s.Message = &v
	return s
}

func (s *RunDeepWritingResponseBody) SetPayload(v *RunDeepWritingResponseBodyPayload) *RunDeepWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunDeepWritingResponseBody) SetRequestId(v string) *RunDeepWritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDeepWritingResponseBody) SetSuccess(v bool) *RunDeepWritingResponseBody {
	s.Success = &v
	return s
}

func (s *RunDeepWritingResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDeepWritingResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// response.output_item.done
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// c2e2e991-f96a-4fcc-9ff7-d0df46c6d232
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// b84d31a5-44b2-4a35-9c6d-878d459c93d0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// FAB10D42-F081-557B-8DCB-D6FB7AAF100B
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDeepWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDeepWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDeepWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDeepWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDeepWritingResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDeepWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDeepWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDeepWritingResponseBodyHeader) SetErrorCode(v string) *RunDeepWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetErrorMessage(v string) *RunDeepWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetEvent(v string) *RunDeepWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetSessionId(v string) *RunDeepWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetStatusCode(v int32) *RunDeepWritingResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetTaskId(v string) *RunDeepWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) SetTraceId(v string) *RunDeepWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDeepWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDeepWritingResponseBodyPayload struct {
	Output *RunDeepWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
}

func (s RunDeepWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyPayload) GetOutput() *RunDeepWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDeepWritingResponseBodyPayload) SetOutput(v *RunDeepWritingResponseBodyPayloadOutput) *RunDeepWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDeepWritingResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDeepWritingResponseBodyPayloadOutput struct {
	Item *RunDeepWritingResponseBodyPayloadOutputItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
	// example:
	//
	// 1
	OutputIndex *int32                                           `json:"OutputIndex,omitempty" xml:"OutputIndex,omitempty"`
	Response    *RunDeepWritingResponseBodyPayloadOutputResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SequenceNumber *string `json:"SequenceNumber,omitempty" xml:"SequenceNumber,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunDeepWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyPayloadOutput) GetItem() *RunDeepWritingResponseBodyPayloadOutputItem {
	return s.Item
}

func (s *RunDeepWritingResponseBodyPayloadOutput) GetOutputIndex() *int32 {
	return s.OutputIndex
}

func (s *RunDeepWritingResponseBodyPayloadOutput) GetResponse() *RunDeepWritingResponseBodyPayloadOutputResponse {
	return s.Response
}

func (s *RunDeepWritingResponseBodyPayloadOutput) GetSequenceNumber() *string {
	return s.SequenceNumber
}

func (s *RunDeepWritingResponseBodyPayloadOutput) GetType() *string {
	return s.Type
}

func (s *RunDeepWritingResponseBodyPayloadOutput) SetItem(v *RunDeepWritingResponseBodyPayloadOutputItem) *RunDeepWritingResponseBodyPayloadOutput {
	s.Item = v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutput) SetOutputIndex(v int32) *RunDeepWritingResponseBodyPayloadOutput {
	s.OutputIndex = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutput) SetResponse(v *RunDeepWritingResponseBodyPayloadOutputResponse) *RunDeepWritingResponseBodyPayloadOutput {
	s.Response = v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutput) SetSequenceNumber(v string) *RunDeepWritingResponseBodyPayloadOutput {
	s.SequenceNumber = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutput) SetType(v string) *RunDeepWritingResponseBodyPayloadOutput {
	s.Type = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutput) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDeepWritingResponseBodyPayloadOutputItem struct {
	// example:
	//
	// ProjectManager
	Agent     *string                                               `json:"Agent,omitempty" xml:"Agent,omitempty"`
	Arguments *string                                               `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Content   []*RunDeepWritingResponseBodyPayloadOutputItemContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 88f6ed9e85c4f9377378da23e6a370d1
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// function_call
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunDeepWritingResponseBodyPayloadOutputItem) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyPayloadOutputItem) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetAgent() *string {
	return s.Agent
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetArguments() *string {
	return s.Arguments
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetContent() []*RunDeepWritingResponseBodyPayloadOutputItemContent {
	return s.Content
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetId() *string {
	return s.Id
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetName() *string {
	return s.Name
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetResult() *string {
	return s.Result
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetStatus() *string {
	return s.Status
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) GetType() *string {
	return s.Type
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetAgent(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Agent = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetArguments(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Arguments = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetContent(v []*RunDeepWritingResponseBodyPayloadOutputItemContent) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Content = v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetId(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Id = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetName(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Name = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetResult(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Result = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetStatus(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Status = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) SetType(v string) *RunDeepWritingResponseBodyPayloadOutputItem {
	s.Type = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItem) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunDeepWritingResponseBodyPayloadOutputItemContent struct {
	// example:
	//
	// <TASK_DONE>
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// output_text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RunDeepWritingResponseBodyPayloadOutputItemContent) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyPayloadOutputItemContent) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyPayloadOutputItemContent) GetText() *string {
	return s.Text
}

func (s *RunDeepWritingResponseBodyPayloadOutputItemContent) GetType() *string {
	return s.Type
}

func (s *RunDeepWritingResponseBodyPayloadOutputItemContent) SetText(v string) *RunDeepWritingResponseBodyPayloadOutputItemContent {
	s.Text = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItemContent) SetType(v string) *RunDeepWritingResponseBodyPayloadOutputItemContent {
	s.Type = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputItemContent) Validate() error {
	return dara.Validate(s)
}

type RunDeepWritingResponseBodyPayloadOutputResponse struct {
	// example:
	//
	// b2dc224b38694e0b668020159a7c5732
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// in_progress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RunDeepWritingResponseBodyPayloadOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponseBodyPayloadOutputResponse) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponseBodyPayloadOutputResponse) GetId() *string {
	return s.Id
}

func (s *RunDeepWritingResponseBodyPayloadOutputResponse) GetStatus() *string {
	return s.Status
}

func (s *RunDeepWritingResponseBodyPayloadOutputResponse) SetId(v string) *RunDeepWritingResponseBodyPayloadOutputResponse {
	s.Id = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputResponse) SetStatus(v string) *RunDeepWritingResponseBodyPayloadOutputResponse {
	s.Status = &v
	return s
}

func (s *RunDeepWritingResponseBodyPayloadOutputResponse) Validate() error {
	return dara.Validate(s)
}
