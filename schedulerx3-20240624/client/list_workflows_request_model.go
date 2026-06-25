// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListWorkflowsRequest
	GetAppName() *string
	SetClusterId(v string) *ListWorkflowsRequest
	GetClusterId() *string
	SetDescription(v string) *ListWorkflowsRequest
	GetDescription() *string
	SetMaxResults(v int32) *ListWorkflowsRequest
	GetMaxResults() *int32
	SetName(v string) *ListWorkflowsRequest
	GetName() *string
	SetNextToken(v string) *ListWorkflowsRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListWorkflowsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListWorkflowsRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListWorkflowsRequest
	GetStatus() *int32
	SetWorkflowId(v int64) *ListWorkflowsRequest
	GetWorkflowId() *int64
}

type ListWorkflowsRequest struct {
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries to return. The default value is 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// myWorkflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If the return value is empty, no more data is returned.
	//
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the workflow. If this parameter is not specified, workflows in the `disabled` state are returned by default. Valid values:
	//
	// - `0`: disabled
	//
	// - `1`: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListWorkflowsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListWorkflowsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowsRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListWorkflowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkflowsRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowsRequest) SetAppName(v string) *ListWorkflowsRequest {
	s.AppName = &v
	return s
}

func (s *ListWorkflowsRequest) SetClusterId(v string) *ListWorkflowsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListWorkflowsRequest) SetDescription(v string) *ListWorkflowsRequest {
	s.Description = &v
	return s
}

func (s *ListWorkflowsRequest) SetMaxResults(v int32) *ListWorkflowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowsRequest) SetName(v string) *ListWorkflowsRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowsRequest) SetNextToken(v string) *ListWorkflowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsRequest) SetPageNum(v int32) *ListWorkflowsRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkflowsRequest) SetPageSize(v int32) *ListWorkflowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowsRequest) SetStatus(v int32) *ListWorkflowsRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowsRequest) SetWorkflowId(v int64) *ListWorkflowsRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
