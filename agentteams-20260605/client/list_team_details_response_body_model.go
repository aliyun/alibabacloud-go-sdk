// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTeamDetailsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTeamDetailsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListTeamDetailsResponseBodyItems) *ListTeamDetailsResponseBody
	GetItems() []*ListTeamDetailsResponseBodyItems
	SetMaxResults(v int32) *ListTeamDetailsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTeamDetailsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTeamDetailsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTeamDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTeamDetailsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTeamDetailsResponseBody
	GetTotalCount() *int64
}

type ListTeamDetailsResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListTeamDetailsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTeamDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTeamDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTeamDetailsResponseBody) GetItems() []*ListTeamDetailsResponseBodyItems {
	return s.Items
}

func (s *ListTeamDetailsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTeamDetailsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTeamDetailsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTeamDetailsResponseBody) SetCode(v string) *ListTeamDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetHttpStatusCode(v int32) *ListTeamDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetItems(v []*ListTeamDetailsResponseBodyItems) *ListTeamDetailsResponseBody {
	s.Items = v
	return s
}

func (s *ListTeamDetailsResponseBody) SetMaxResults(v int32) *ListTeamDetailsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetMessage(v string) *ListTeamDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetNextToken(v string) *ListTeamDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetRequestId(v string) *ListTeamDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetSuccess(v bool) *ListTeamDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTeamDetailsResponseBody) SetTotalCount(v int64) *ListTeamDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTeamDetailsResponseBody) Validate() error {
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

type ListTeamDetailsResponseBodyItems struct {
	AvatarUrl   *string  `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Status      *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessRate *float64 `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	TaskCount   *int32   `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TeamName    *string  `json:"TeamName,omitempty" xml:"TeamName,omitempty"`
	TokenUsage  *int64   `json:"TokenUsage,omitempty" xml:"TokenUsage,omitempty"`
	WorkerCount *int32   `json:"WorkerCount,omitempty" xml:"WorkerCount,omitempty"`
}

func (s ListTeamDetailsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTeamDetailsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTeamDetailsResponseBodyItems) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListTeamDetailsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListTeamDetailsResponseBodyItems) GetSuccessRate() *float64 {
	return s.SuccessRate
}

func (s *ListTeamDetailsResponseBodyItems) GetTaskCount() *int32 {
	return s.TaskCount
}

func (s *ListTeamDetailsResponseBodyItems) GetTeamName() *string {
	return s.TeamName
}

func (s *ListTeamDetailsResponseBodyItems) GetTokenUsage() *int64 {
	return s.TokenUsage
}

func (s *ListTeamDetailsResponseBodyItems) GetWorkerCount() *int32 {
	return s.WorkerCount
}

func (s *ListTeamDetailsResponseBodyItems) SetAvatarUrl(v string) *ListTeamDetailsResponseBodyItems {
	s.AvatarUrl = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetStatus(v string) *ListTeamDetailsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetSuccessRate(v float64) *ListTeamDetailsResponseBodyItems {
	s.SuccessRate = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetTaskCount(v int32) *ListTeamDetailsResponseBodyItems {
	s.TaskCount = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetTeamName(v string) *ListTeamDetailsResponseBodyItems {
	s.TeamName = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetTokenUsage(v int64) *ListTeamDetailsResponseBodyItems {
	s.TokenUsage = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) SetWorkerCount(v int32) *ListTeamDetailsResponseBodyItems {
	s.WorkerCount = &v
	return s
}

func (s *ListTeamDetailsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
