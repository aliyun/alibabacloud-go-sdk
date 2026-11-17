// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduProblemSolveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateEduProblemSolveResponseBodyData) *CreateEduProblemSolveResponseBody
	GetData() *CreateEduProblemSolveResponseBodyData
	SetEvent(v string) *CreateEduProblemSolveResponseBody
	GetEvent() *string
}

type CreateEduProblemSolveResponseBody struct {
	Data *CreateEduProblemSolveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// message
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
}

func (s CreateEduProblemSolveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBody) GetData() *CreateEduProblemSolveResponseBodyData {
	return s.Data
}

func (s *CreateEduProblemSolveResponseBody) GetEvent() *string {
	return s.Event
}

func (s *CreateEduProblemSolveResponseBody) SetData(v *CreateEduProblemSolveResponseBodyData) *CreateEduProblemSolveResponseBody {
	s.Data = v
	return s
}

func (s *CreateEduProblemSolveResponseBody) SetEvent(v string) *CreateEduProblemSolveResponseBody {
	s.Event = &v
	return s
}

func (s *CreateEduProblemSolveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveResponseBodyData struct {
	// example:
	//
	// change_batch
	Code     *string                                       `json:"code,omitempty" xml:"code,omitempty"`
	Finished *bool                                         `json:"finished,omitempty" xml:"finished,omitempty"`
	Items    []*CreateEduProblemSolveResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 098E23F7-A74C-5584-87A1-39734D5B83CE
	RequestId  *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultType *string                                     `json:"resultType,omitempty" xml:"resultType,omitempty"`
	Retryable  *bool                                       `json:"retryable,omitempty" xml:"retryable,omitempty"`
	Usage      *CreateEduProblemSolveResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s CreateEduProblemSolveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *CreateEduProblemSolveResponseBodyData) GetFinished() *bool {
	return s.Finished
}

func (s *CreateEduProblemSolveResponseBodyData) GetItems() []*CreateEduProblemSolveResponseBodyDataItems {
	return s.Items
}

func (s *CreateEduProblemSolveResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateEduProblemSolveResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduProblemSolveResponseBodyData) GetResultType() *string {
	return s.ResultType
}

func (s *CreateEduProblemSolveResponseBodyData) GetRetryable() *bool {
	return s.Retryable
}

func (s *CreateEduProblemSolveResponseBodyData) GetUsage() *CreateEduProblemSolveResponseBodyDataUsage {
	return s.Usage
}

func (s *CreateEduProblemSolveResponseBodyData) SetCode(v string) *CreateEduProblemSolveResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetFinished(v bool) *CreateEduProblemSolveResponseBodyData {
	s.Finished = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetItems(v []*CreateEduProblemSolveResponseBodyDataItems) *CreateEduProblemSolveResponseBodyData {
	s.Items = v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetMessage(v string) *CreateEduProblemSolveResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetRequestId(v string) *CreateEduProblemSolveResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetResultType(v string) *CreateEduProblemSolveResponseBodyData {
	s.ResultType = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetRetryable(v bool) *CreateEduProblemSolveResponseBodyData {
	s.Retryable = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) SetUsage(v *CreateEduProblemSolveResponseBodyDataUsage) *CreateEduProblemSolveResponseBodyData {
	s.Usage = v
	return s
}

func (s *CreateEduProblemSolveResponseBodyData) Validate() error {
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

type CreateEduProblemSolveResponseBodyDataItems struct {
	// example:
	//
	// {}
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	ImageUrl   *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	QuestionId *string `json:"questionId,omitempty" xml:"questionId,omitempty"`
	// example:
	//
	// 200
	Status   *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	Thinking *CreateEduProblemSolveResponseBodyDataItemsThinking `json:"thinking,omitempty" xml:"thinking,omitempty" type:"Struct"`
}

func (s CreateEduProblemSolveResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyDataItems) GetContent() *string {
	return s.Content
}

func (s *CreateEduProblemSolveResponseBodyDataItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateEduProblemSolveResponseBodyDataItems) GetQuestionId() *string {
	return s.QuestionId
}

func (s *CreateEduProblemSolveResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *CreateEduProblemSolveResponseBodyDataItems) GetThinking() *CreateEduProblemSolveResponseBodyDataItemsThinking {
	return s.Thinking
}

func (s *CreateEduProblemSolveResponseBodyDataItems) SetContent(v string) *CreateEduProblemSolveResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItems) SetImageUrl(v string) *CreateEduProblemSolveResponseBodyDataItems {
	s.ImageUrl = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItems) SetQuestionId(v string) *CreateEduProblemSolveResponseBodyDataItems {
	s.QuestionId = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItems) SetStatus(v string) *CreateEduProblemSolveResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItems) SetThinking(v *CreateEduProblemSolveResponseBodyDataItemsThinking) *CreateEduProblemSolveResponseBodyDataItems {
	s.Thinking = v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItems) Validate() error {
	if s.Thinking != nil {
		if err := s.Thinking.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveResponseBodyDataItemsThinking struct {
	// example:
	//
	// 8ci7nsj9WPYXvQb9MJfA
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateEduProblemSolveResponseBodyDataItemsThinking) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyDataItemsThinking) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyDataItemsThinking) GetContent() *string {
	return s.Content
}

func (s *CreateEduProblemSolveResponseBodyDataItemsThinking) GetStatus() *string {
	return s.Status
}

func (s *CreateEduProblemSolveResponseBodyDataItemsThinking) SetContent(v string) *CreateEduProblemSolveResponseBodyDataItemsThinking {
	s.Content = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItemsThinking) SetStatus(v string) *CreateEduProblemSolveResponseBodyDataItemsThinking {
	s.Status = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataItemsThinking) Validate() error {
	return dara.Validate(s)
}

type CreateEduProblemSolveResponseBodyDataUsage struct {
	CompletionTokens *int32 `json:"completionTokens,omitempty" xml:"completionTokens,omitempty"`
	PromptTokens     *int32 `json:"promptTokens,omitempty" xml:"promptTokens,omitempty"`
}

func (s CreateEduProblemSolveResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponseBodyDataUsage) GetCompletionTokens() *int32 {
	return s.CompletionTokens
}

func (s *CreateEduProblemSolveResponseBodyDataUsage) GetPromptTokens() *int32 {
	return s.PromptTokens
}

func (s *CreateEduProblemSolveResponseBodyDataUsage) SetCompletionTokens(v int32) *CreateEduProblemSolveResponseBodyDataUsage {
	s.CompletionTokens = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataUsage) SetPromptTokens(v int32) *CreateEduProblemSolveResponseBodyDataUsage {
	s.PromptTokens = &v
	return s
}

func (s *CreateEduProblemSolveResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
