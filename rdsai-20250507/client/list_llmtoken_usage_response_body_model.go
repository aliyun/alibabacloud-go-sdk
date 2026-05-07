// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMTokenUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompletionTokens(v int64) *ListLLMTokenUsageResponseBody
	GetCompletionTokens() *int64
	SetExplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBody
	GetExplicitCachedTokens() *int64
	SetImplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBody
	GetImplicitCachedTokens() *int64
	SetPromptTokens(v int64) *ListLLMTokenUsageResponseBody
	GetPromptTokens() *int64
	SetRequestId(v string) *ListLLMTokenUsageResponseBody
	GetRequestId() *string
	SetTokenUsages(v []*ListLLMTokenUsageResponseBodyTokenUsages) *ListLLMTokenUsageResponseBody
	GetTokenUsages() []*ListLLMTokenUsageResponseBodyTokenUsages
	SetTotalTokens(v int64) *ListLLMTokenUsageResponseBody
	GetTotalTokens() *int64
}

type ListLLMTokenUsageResponseBody struct {
	// example:
	//
	// 11345
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 1500
	ExplicitCachedTokens *int64 `json:"ExplicitCachedTokens,omitempty" xml:"ExplicitCachedTokens,omitempty"`
	// example:
	//
	// 1500
	ImplicitCachedTokens *int64 `json:"ImplicitCachedTokens,omitempty" xml:"ImplicitCachedTokens,omitempty"`
	// example:
	//
	// 1000
	PromptTokens *int64 `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenUsages []*ListLLMTokenUsageResponseBodyTokenUsages `json:"TokenUsages,omitempty" xml:"TokenUsages,omitempty" type:"Repeated"`
	// example:
	//
	// 15345
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s ListLLMTokenUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLLMTokenUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListLLMTokenUsageResponseBody) GetCompletionTokens() *int64 {
	return s.CompletionTokens
}

func (s *ListLLMTokenUsageResponseBody) GetExplicitCachedTokens() *int64 {
	return s.ExplicitCachedTokens
}

func (s *ListLLMTokenUsageResponseBody) GetImplicitCachedTokens() *int64 {
	return s.ImplicitCachedTokens
}

func (s *ListLLMTokenUsageResponseBody) GetPromptTokens() *int64 {
	return s.PromptTokens
}

func (s *ListLLMTokenUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLLMTokenUsageResponseBody) GetTokenUsages() []*ListLLMTokenUsageResponseBodyTokenUsages {
	return s.TokenUsages
}

func (s *ListLLMTokenUsageResponseBody) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ListLLMTokenUsageResponseBody) SetCompletionTokens(v int64) *ListLLMTokenUsageResponseBody {
	s.CompletionTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetExplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBody {
	s.ExplicitCachedTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetImplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBody {
	s.ImplicitCachedTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetPromptTokens(v int64) *ListLLMTokenUsageResponseBody {
	s.PromptTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetRequestId(v string) *ListLLMTokenUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetTokenUsages(v []*ListLLMTokenUsageResponseBodyTokenUsages) *ListLLMTokenUsageResponseBody {
	s.TokenUsages = v
	return s
}

func (s *ListLLMTokenUsageResponseBody) SetTotalTokens(v int64) *ListLLMTokenUsageResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBody) Validate() error {
	if s.TokenUsages != nil {
		for _, item := range s.TokenUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLLMTokenUsageResponseBodyTokenUsages struct {
	// example:
	//
	// 1024
	CompletionReasoningTokens *int64 `json:"CompletionReasoningTokens,omitempty" xml:"CompletionReasoningTokens,omitempty"`
	// example:
	//
	// 10000
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 2025-12-01T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1024
	ExplicitCachedTokens *int64 `json:"ExplicitCachedTokens,omitempty" xml:"ExplicitCachedTokens,omitempty"`
	// example:
	//
	// 1408
	ImplicitCachedTokens *int64 `json:"ImplicitCachedTokens,omitempty" xml:"ImplicitCachedTokens,omitempty"`
	// example:
	//
	// glm-5
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 1234
	PromptTokens *int64 `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// example:
	//
	// 2025-03-27
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 12345
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s ListLLMTokenUsageResponseBodyTokenUsages) String() string {
	return dara.Prettify(s)
}

func (s ListLLMTokenUsageResponseBodyTokenUsages) GoString() string {
	return s.String()
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetCompletionReasoningTokens() *int64 {
	return s.CompletionReasoningTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetCompletionTokens() *int64 {
	return s.CompletionTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetExplicitCachedTokens() *int64 {
	return s.ExplicitCachedTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetImplicitCachedTokens() *int64 {
	return s.ImplicitCachedTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetModel() *string {
	return s.Model
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetPromptTokens() *int64 {
	return s.PromptTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetCompletionReasoningTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.CompletionReasoningTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetCompletionTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.CompletionTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetEndTime(v string) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.EndTime = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetExplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.ExplicitCachedTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetImplicitCachedTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.ImplicitCachedTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetModel(v string) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.Model = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetPromptTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.PromptTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetStartTime(v string) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.StartTime = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) SetTotalTokens(v int64) *ListLLMTokenUsageResponseBodyTokenUsages {
	s.TotalTokens = &v
	return s
}

func (s *ListLLMTokenUsageResponseBodyTokenUsages) Validate() error {
	return dara.Validate(s)
}
