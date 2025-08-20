// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMSimilarQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeLLMSimilarQuestionsResponseBodyItems) *DescribeLLMSimilarQuestionsResponseBody
	GetItems() []*DescribeLLMSimilarQuestionsResponseBodyItems
	SetRequestId(v string) *DescribeLLMSimilarQuestionsResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DescribeLLMSimilarQuestionsResponseBody
	GetSessionId() *string
}

type DescribeLLMSimilarQuestionsResponseBody struct {
	// The queried similar questions.
	Items []*DescribeLLMSimilarQuestionsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 96A55627-28E9-5E47-B8F6-D786BE551349
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 4847
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeLLMSimilarQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMSimilarQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLLMSimilarQuestionsResponseBody) GetItems() []*DescribeLLMSimilarQuestionsResponseBodyItems {
	return s.Items
}

func (s *DescribeLLMSimilarQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLLMSimilarQuestionsResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeLLMSimilarQuestionsResponseBody) SetItems(v []*DescribeLLMSimilarQuestionsResponseBodyItems) *DescribeLLMSimilarQuestionsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBody) SetRequestId(v string) *DescribeLLMSimilarQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBody) SetSessionId(v string) *DescribeLLMSimilarQuestionsResponseBody {
	s.SessionId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLLMSimilarQuestionsResponseBodyItems struct {
	// The answer to the similar question.
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The ID of the similar question.
	//
	// example:
	//
	// 2389899
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The similarity of the similar question.
	//
	// example:
	//
	// 0.58
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The source of the similar question.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The summary of the similar question.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The content of the similar question.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the answer to the similar question.
	//
	// example:
	//
	// www.aliyun.com/product
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeLLMSimilarQuestionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMSimilarQuestionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetAnswer() *string {
	return s.Answer
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetScore() *float64 {
	return s.Score
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetSource() *string {
	return s.Source
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetSummary() *string {
	return s.Summary
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetAnswer(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Answer = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetId(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetScore(v float64) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Score = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetSource(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Source = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetSummary(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Summary = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetTitle(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) SetUrl(v string) *DescribeLLMSimilarQuestionsResponseBodyItems {
	s.Url = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
