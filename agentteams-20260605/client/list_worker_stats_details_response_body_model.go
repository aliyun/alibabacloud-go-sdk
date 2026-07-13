// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerStatsDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkerStatsDetailsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListWorkerStatsDetailsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListWorkerStatsDetailsResponseBodyItems) *ListWorkerStatsDetailsResponseBody
	GetItems() []*ListWorkerStatsDetailsResponseBodyItems
	SetMaxResults(v int32) *ListWorkerStatsDetailsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWorkerStatsDetailsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkerStatsDetailsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkerStatsDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkerStatsDetailsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListWorkerStatsDetailsResponseBody
	GetTotalCount() *int64
}

type ListWorkerStatsDetailsResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListWorkerStatsDetailsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkerStatsDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerStatsDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkerStatsDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkerStatsDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWorkerStatsDetailsResponseBody) GetItems() []*ListWorkerStatsDetailsResponseBodyItems {
	return s.Items
}

func (s *ListWorkerStatsDetailsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkerStatsDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkerStatsDetailsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkerStatsDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkerStatsDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkerStatsDetailsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkerStatsDetailsResponseBody) SetCode(v string) *ListWorkerStatsDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetHttpStatusCode(v int32) *ListWorkerStatsDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetItems(v []*ListWorkerStatsDetailsResponseBodyItems) *ListWorkerStatsDetailsResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetMaxResults(v int32) *ListWorkerStatsDetailsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetMessage(v string) *ListWorkerStatsDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetNextToken(v string) *ListWorkerStatsDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetRequestId(v string) *ListWorkerStatsDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetSuccess(v bool) *ListWorkerStatsDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) SetTotalCount(v int64) *ListWorkerStatsDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkerStatsDetailsResponseBodyItems struct {
	LlmCallCount *int64  `json:"LlmCallCount,omitempty" xml:"LlmCallCount,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskCount    *int64  `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TokenUsage   *int64  `json:"TokenUsage,omitempty" xml:"TokenUsage,omitempty"`
}

func (s ListWorkerStatsDetailsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerStatsDetailsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetLlmCallCount() *int64 {
	return s.LlmCallCount
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetModel() *string {
	return s.Model
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *ListWorkerStatsDetailsResponseBodyItems) GetTokenUsage() *int64 {
	return s.TokenUsage
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetLlmCallCount(v int64) *ListWorkerStatsDetailsResponseBodyItems {
	s.LlmCallCount = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetModel(v string) *ListWorkerStatsDetailsResponseBodyItems {
	s.Model = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetName(v string) *ListWorkerStatsDetailsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetStatus(v string) *ListWorkerStatsDetailsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetTaskCount(v int64) *ListWorkerStatsDetailsResponseBodyItems {
	s.TaskCount = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) SetTokenUsage(v int64) *ListWorkerStatsDetailsResponseBodyItems {
	s.TokenUsage = &v
	return s
}

func (s *ListWorkerStatsDetailsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
