// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTeamTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTeamTasksResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListTeamTasksResponseBodyItems) *ListTeamTasksResponseBody
	GetItems() []*ListTeamTasksResponseBodyItems
	SetMaxResults(v int32) *ListTeamTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTeamTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTeamTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTeamTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTeamTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTeamTasksResponseBody
	GetTotalCount() *int64
}

type ListTeamTasksResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListTeamTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTeamTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTeamTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTeamTasksResponseBody) GetItems() []*ListTeamTasksResponseBodyItems {
	return s.Items
}

func (s *ListTeamTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTeamTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTeamTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTeamTasksResponseBody) SetCode(v string) *ListTeamTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetHttpStatusCode(v int32) *ListTeamTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetItems(v []*ListTeamTasksResponseBodyItems) *ListTeamTasksResponseBody {
	s.Items = v
	return s
}

func (s *ListTeamTasksResponseBody) SetMaxResults(v int32) *ListTeamTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetMessage(v string) *ListTeamTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetNextToken(v string) *ListTeamTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetRequestId(v string) *ListTeamTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetSuccess(v bool) *ListTeamTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListTeamTasksResponseBody) SetTotalCount(v int64) *ListTeamTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTeamTasksResponseBody) Validate() error {
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

type ListTeamTasksResponseBodyItems struct {
	AssignedTo *string `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	CreatedAt  *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskTitle  *string `json:"TaskTitle,omitempty" xml:"TaskTitle,omitempty"`
}

func (s ListTeamTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTeamTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTeamTasksResponseBodyItems) GetAssignedTo() *string {
	return s.AssignedTo
}

func (s *ListTeamTasksResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListTeamTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListTeamTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTeamTasksResponseBodyItems) GetTaskTitle() *string {
	return s.TaskTitle
}

func (s *ListTeamTasksResponseBodyItems) SetAssignedTo(v string) *ListTeamTasksResponseBodyItems {
	s.AssignedTo = &v
	return s
}

func (s *ListTeamTasksResponseBodyItems) SetCreatedAt(v string) *ListTeamTasksResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTeamTasksResponseBodyItems) SetStatus(v string) *ListTeamTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTeamTasksResponseBodyItems) SetTaskId(v string) *ListTeamTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListTeamTasksResponseBodyItems) SetTaskTitle(v string) *ListTeamTasksResponseBodyItems {
	s.TaskTitle = &v
	return s
}

func (s *ListTeamTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
