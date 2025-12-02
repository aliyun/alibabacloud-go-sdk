// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmStreamChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChoices(v []*LlmStreamChatResponseBodyChoices) *LlmStreamChatResponseBody
	GetChoices() []*LlmStreamChatResponseBodyChoices
	SetCreated(v int64) *LlmStreamChatResponseBody
	GetCreated() *int64
	SetError(v *LlmStreamChatResponseBodyError) *LlmStreamChatResponseBody
	GetError() *LlmStreamChatResponseBodyError
	SetId(v string) *LlmStreamChatResponseBody
	GetId() *string
	SetModel(v string) *LlmStreamChatResponseBody
	GetModel() *string
	SetObject(v string) *LlmStreamChatResponseBody
	GetObject() *string
	SetRequestId(v string) *LlmStreamChatResponseBody
	GetRequestId() *string
	SetSystemFingerprint(v string) *LlmStreamChatResponseBody
	GetSystemFingerprint() *string
	SetUsage(v string) *LlmStreamChatResponseBody
	GetUsage() *string
}

type LlmStreamChatResponseBody struct {
	// List of model generation results
	Choices []*LlmStreamChatResponseBodyChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
	// Timestamp of session creation
	//
	// example:
	//
	// 1750990728
	Created *int64 `json:"Created,omitempty" xml:"Created,omitempty"`
	// Streaming response error information content
	Error *LlmStreamChatResponseBodyError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// Unique ID for this session
	//
	// example:
	//
	// chatcmpl-777bce52-93d3-9f8c-89c3-e99884f4f57f
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Model identifier
	//
	// example:
	//
	// deepseek-v3
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Response type
	//
	// example:
	//
	// chat.completion.chunk
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Unique request ID
	//
	// example:
	//
	// 21d296d6-594e-97de-812f-925ec6e05673
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// System fingerprint
	//
	// example:
	//
	// null
	SystemFingerprint *string `json:"SystemFingerprint,omitempty" xml:"SystemFingerprint,omitempty"`
	// Token usage
	//
	// example:
	//
	// null
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s LlmStreamChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatResponseBody) GoString() string {
	return s.String()
}

func (s *LlmStreamChatResponseBody) GetChoices() []*LlmStreamChatResponseBodyChoices {
	return s.Choices
}

func (s *LlmStreamChatResponseBody) GetCreated() *int64 {
	return s.Created
}

func (s *LlmStreamChatResponseBody) GetError() *LlmStreamChatResponseBodyError {
	return s.Error
}

func (s *LlmStreamChatResponseBody) GetId() *string {
	return s.Id
}

func (s *LlmStreamChatResponseBody) GetModel() *string {
	return s.Model
}

func (s *LlmStreamChatResponseBody) GetObject() *string {
	return s.Object
}

func (s *LlmStreamChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmStreamChatResponseBody) GetSystemFingerprint() *string {
	return s.SystemFingerprint
}

func (s *LlmStreamChatResponseBody) GetUsage() *string {
	return s.Usage
}

func (s *LlmStreamChatResponseBody) SetChoices(v []*LlmStreamChatResponseBodyChoices) *LlmStreamChatResponseBody {
	s.Choices = v
	return s
}

func (s *LlmStreamChatResponseBody) SetCreated(v int64) *LlmStreamChatResponseBody {
	s.Created = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetError(v *LlmStreamChatResponseBodyError) *LlmStreamChatResponseBody {
	s.Error = v
	return s
}

func (s *LlmStreamChatResponseBody) SetId(v string) *LlmStreamChatResponseBody {
	s.Id = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetModel(v string) *LlmStreamChatResponseBody {
	s.Model = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetObject(v string) *LlmStreamChatResponseBody {
	s.Object = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetRequestId(v string) *LlmStreamChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetSystemFingerprint(v string) *LlmStreamChatResponseBody {
	s.SystemFingerprint = &v
	return s
}

func (s *LlmStreamChatResponseBody) SetUsage(v string) *LlmStreamChatResponseBody {
	s.Usage = &v
	return s
}

func (s *LlmStreamChatResponseBody) Validate() error {
	if s.Choices != nil {
		for _, item := range s.Choices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LlmStreamChatResponseBodyChoices struct {
	// Incremental content object
	Delta *LlmStreamChatResponseBodyChoicesDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Struct"`
	// For streaming output, it is null while generating and becomes \\"stop\\" if the generation ends due to a stop token.
	//
	// example:
	//
	// stop
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// Stream sequence number
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Token probability information
	//
	// example:
	//
	// null
	Logprobs *string `json:"Logprobs,omitempty" xml:"Logprobs,omitempty"`
}

func (s LlmStreamChatResponseBodyChoices) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *LlmStreamChatResponseBodyChoices) GetDelta() *LlmStreamChatResponseBodyChoicesDelta {
	return s.Delta
}

func (s *LlmStreamChatResponseBodyChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *LlmStreamChatResponseBodyChoices) GetIndex() *int64 {
	return s.Index
}

func (s *LlmStreamChatResponseBodyChoices) GetLogprobs() *string {
	return s.Logprobs
}

func (s *LlmStreamChatResponseBodyChoices) SetDelta(v *LlmStreamChatResponseBodyChoicesDelta) *LlmStreamChatResponseBodyChoices {
	s.Delta = v
	return s
}

func (s *LlmStreamChatResponseBodyChoices) SetFinishReason(v string) *LlmStreamChatResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *LlmStreamChatResponseBodyChoices) SetIndex(v int64) *LlmStreamChatResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *LlmStreamChatResponseBodyChoices) SetLogprobs(v string) *LlmStreamChatResponseBodyChoices {
	s.Logprobs = &v
	return s
}

func (s *LlmStreamChatResponseBodyChoices) Validate() error {
	if s.Delta != nil {
		if err := s.Delta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LlmStreamChatResponseBodyChoicesDelta struct {
	// Real-time generated text content
	//
	// example:
	//
	// 我是Deepseek-V3，有什么可以帮你
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Role identifier
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s LlmStreamChatResponseBodyChoicesDelta) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatResponseBodyChoicesDelta) GoString() string {
	return s.String()
}

func (s *LlmStreamChatResponseBodyChoicesDelta) GetContent() *string {
	return s.Content
}

func (s *LlmStreamChatResponseBodyChoicesDelta) GetRole() *string {
	return s.Role
}

func (s *LlmStreamChatResponseBodyChoicesDelta) SetContent(v string) *LlmStreamChatResponseBodyChoicesDelta {
	s.Content = &v
	return s
}

func (s *LlmStreamChatResponseBodyChoicesDelta) SetRole(v string) *LlmStreamChatResponseBodyChoicesDelta {
	s.Role = &v
	return s
}

func (s *LlmStreamChatResponseBodyChoicesDelta) Validate() error {
	return dara.Validate(s)
}

type LlmStreamChatResponseBodyError struct {
	// Error code
	//
	// example:
	//
	// data_inspection_failed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message
	//
	// example:
	//
	// Input data may contain inappropriate content.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Parameter that caused the error
	//
	// example:
	//
	// null
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// Error type
	//
	// example:
	//
	// data_inspection_failed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LlmStreamChatResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatResponseBodyError) GoString() string {
	return s.String()
}

func (s *LlmStreamChatResponseBodyError) GetCode() *string {
	return s.Code
}

func (s *LlmStreamChatResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *LlmStreamChatResponseBodyError) GetParam() *string {
	return s.Param
}

func (s *LlmStreamChatResponseBodyError) GetType() *string {
	return s.Type
}

func (s *LlmStreamChatResponseBodyError) SetCode(v string) *LlmStreamChatResponseBodyError {
	s.Code = &v
	return s
}

func (s *LlmStreamChatResponseBodyError) SetMessage(v string) *LlmStreamChatResponseBodyError {
	s.Message = &v
	return s
}

func (s *LlmStreamChatResponseBodyError) SetParam(v string) *LlmStreamChatResponseBodyError {
	s.Param = &v
	return s
}

func (s *LlmStreamChatResponseBodyError) SetType(v string) *LlmStreamChatResponseBodyError {
	s.Type = &v
	return s
}

func (s *LlmStreamChatResponseBodyError) Validate() error {
	return dara.Validate(s)
}
