// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAITasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAITasksRequest
	GetNextToken() *string
	SetTaskId(v string) *ListAITasksRequest
	GetTaskId() *string
	SetTaskType(v string) *ListAITasksRequest
	GetTaskType() *string
}

type ListAITasksRequest struct {
	// The maximum number of data entries to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// U12WEI6Ro2ol3wA54rBNS3Cltv2VJyA+7hP4GqbIOhmWU5mWU9ZE3cXLgDaH4KSMRfIYcIVrvtHaAzCoyfo7VQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the AI task. You can filter AI tasks by task ID.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the AI task. You can filter AI tasks by task type.
	//
	// 	- GenerateTemplate: The AI task is used to generate a template.
	//
	// 	- FixTemplate: The AI task is used to fix a template.
	//
	// If you leave this parameter empty, all task types are queried.
	//
	// example:
	//
	// GenerateTemplate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListAITasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAITasksRequest) GoString() string {
	return s.String()
}

func (s *ListAITasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAITasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAITasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAITasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAITasksRequest) SetMaxResults(v int32) *ListAITasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAITasksRequest) SetNextToken(v string) *ListAITasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListAITasksRequest) SetTaskId(v string) *ListAITasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListAITasksRequest) SetTaskType(v string) *ListAITasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListAITasksRequest) Validate() error {
	return dara.Validate(s)
}
