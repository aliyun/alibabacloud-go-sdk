// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduProblemSolveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEduProblemSolveResponseBody
	GetCode() *string
	SetFinished(v bool) *CreateEduProblemSolveResponseBody
	GetFinished() *bool
	SetItems(v []*CreateEduProblemSolveResponseBodyItems) *CreateEduProblemSolveResponseBody
	GetItems() []*CreateEduProblemSolveResponseBodyItems
	SetMessage(v string) *CreateEduProblemSolveResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEduProblemSolveResponseBody
	GetRequestId() *string
	SetResultType(v string) *CreateEduProblemSolveResponseBody
	GetResultType() *string
	SetRetryable(v bool) *CreateEduProblemSolveResponseBody
	GetRetryable() *bool
	SetUsage(v *CreateEduProblemSolveResponseBodyUsage) *CreateEduProblemSolveResponseBody
	GetUsage() *CreateEduProblemSolveResponseBodyUsage
}

type CreateEduProblemSolveResponseBody struct {
	// The API data identifier.
	//
	// example:
	//
	// change_batch
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Specifies whether this is the final full snapshot. The value is fixed to true in the complete event.
	Finished *bool `json:"finished,omitempty" xml:"finished,omitempty"`
	// The current list of complete question results.
	Items []*CreateEduProblemSolveResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The actual response message of the request.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Echo request ID echoed from echo request.
	//
	// example:
	//
	// 098E23F7-A74C-5584-87A1-39734D5B83CE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The type of the question resolution result. Valid values:
	//
	// 	- single: Single question.
	//
	// 	- multiple: Multiple questions.
	//
	// example:
	//
	// single
	ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
	// Specifies whether the error is retryable. Default value: true.
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// The token usage. Only the complete event carries the final value. The message event does not return this field.
	Usage *CreateEduProblemSolveResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s CreateEduProblemSolveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEduProblemSolveResponseBody) GetFinished() *bool {
	return s.Finished
}

func (s *CreateEduProblemSolveResponseBody) GetItems() []*CreateEduProblemSolveResponseBodyItems {
	return s.Items
}

func (s *CreateEduProblemSolveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEduProblemSolveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduProblemSolveResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *CreateEduProblemSolveResponseBody) GetRetryable() *bool {
	return s.Retryable
}

func (s *CreateEduProblemSolveResponseBody) GetUsage() *CreateEduProblemSolveResponseBodyUsage {
	return s.Usage
}

func (s *CreateEduProblemSolveResponseBody) SetCode(v string) *CreateEduProblemSolveResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetFinished(v bool) *CreateEduProblemSolveResponseBody {
	s.Finished = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetItems(v []*CreateEduProblemSolveResponseBodyItems) *CreateEduProblemSolveResponseBody {
	s.Items = v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetMessage(v string) *CreateEduProblemSolveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetRequestId(v string) *CreateEduProblemSolveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetResultType(v string) *CreateEduProblemSolveResponseBody {
	s.ResultType = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetRetryable(v bool) *CreateEduProblemSolveResponseBody {
	s.Retryable = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetUsage(v *CreateEduProblemSolveResponseBodyUsage) *CreateEduProblemSolveResponseBody {
	s.Usage = v
	return s
}

func (s *CreateEduProblemSolveResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveResponseBodyItems struct {
	// The complete Markdown content of the question resolution, including the deduction process and the final result.
	//
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The URL of the unique image for the current question. The URL may be time-sensitive. If the question is not associated with an image resource and no request URL is available for fallback, an empty string is returned.
	//
	// example:
	//
	// https://xxx.jpg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// The question identifier. The value is the question title, such as Question 1. The server does not return a separate title field. If the downstream service does not provide a question title, which is always the case for single-question scenarios, titles such as Question 1 and Question 2 are generated in the order of the questions.
	//
	// example:
	//
	// 50a10e43a33e36179e355556d319ebd5
	QuestionId *string `json:"questionId,omitempty" xml:"questionId,omitempty"`
	// The current question status. Valid values:
	//
	// 	- processing: Processing.
	//
	// 	- complete: Complete.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The thinking process. This field is returned only when result_type is single, and is not returned for multiple questions.
	Thinking *CreateEduProblemSolveResponseBodyItemsThinking `json:"thinking,omitempty" xml:"thinking,omitempty" type:"Struct"`
}

func (s CreateEduProblemSolveResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyItems) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *CreateEduProblemSolveResponseBodyItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateEduProblemSolveResponseBodyItems) GetQuestionId() *string {
	return s.QuestionId
}

func (s *CreateEduProblemSolveResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *CreateEduProblemSolveResponseBodyItems) GetThinking() *CreateEduProblemSolveResponseBodyItemsThinking {
	return s.Thinking
}

func (s *CreateEduProblemSolveResponseBodyItems) SetContent(v string) *CreateEduProblemSolveResponseBodyItems {
	s.Content = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItems) SetImageUrl(v string) *CreateEduProblemSolveResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItems) SetQuestionId(v string) *CreateEduProblemSolveResponseBodyItems {
	s.QuestionId = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItems) SetStatus(v string) *CreateEduProblemSolveResponseBodyItems {
	s.Status = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItems) SetThinking(v *CreateEduProblemSolveResponseBodyItemsThinking) *CreateEduProblemSolveResponseBodyItems {
	s.Thinking = v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItems) Validate() error {
	if s.Thinking != nil {
		if err := s.Thinking.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveResponseBodyItemsThinking struct {
	// The current thinking content.
	//
	// example:
	//
	// 8ci7nsj9WPYXvQb9MJfA
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The thinking status. Valid values:
	//
	// 	- processing: Processing.
	//
	// 	- complete: Complete.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateEduProblemSolveResponseBodyItemsThinking) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyItemsThinking) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyItemsThinking) GetContent() *string {
	return s.Content
}

func (s *CreateEduProblemSolveResponseBodyItemsThinking) GetStatus() *string {
	return s.Status
}

func (s *CreateEduProblemSolveResponseBodyItemsThinking) SetContent(v string) *CreateEduProblemSolveResponseBodyItemsThinking {
	s.Content = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItemsThinking) SetStatus(v string) *CreateEduProblemSolveResponseBodyItemsThinking {
	s.Status = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyItemsThinking) Validate() error {
	return dara.Validate(s)
}

type CreateEduProblemSolveResponseBodyUsage struct {
	// The number of output tokens.
	//
	// example:
	//
	// 1024
	CompletionTokens *int32 `json:"completionTokens,omitempty" xml:"completionTokens,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 1024
	PromptTokens *int32 `json:"promptTokens,omitempty" xml:"promptTokens,omitempty"`
}

func (s CreateEduProblemSolveResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyUsage) GetCompletionTokens() *int32 {
	return s.CompletionTokens
}

func (s *CreateEduProblemSolveResponseBodyUsage) GetPromptTokens() *int32 {
	return s.PromptTokens
}

func (s *CreateEduProblemSolveResponseBodyUsage) SetCompletionTokens(v int32) *CreateEduProblemSolveResponseBodyUsage {
	s.CompletionTokens = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyUsage) SetPromptTokens(v int32) *CreateEduProblemSolveResponseBodyUsage {
	s.PromptTokens = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
