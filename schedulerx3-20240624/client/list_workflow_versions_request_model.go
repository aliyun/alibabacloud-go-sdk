// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListWorkflowVersionsRequest
	GetAppName() *string
	SetClusterId(v string) *ListWorkflowVersionsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListWorkflowVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowVersionsRequest
	GetNextToken() *string
	SetWorkflowId(v int64) *ListWorkflowVersionsRequest
	GetWorkflowId() *int64
}

type ListWorkflowVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowVersionsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListWorkflowVersionsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListWorkflowVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowVersionsRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowVersionsRequest) SetAppName(v string) *ListWorkflowVersionsRequest {
	s.AppName = &v
	return s
}

func (s *ListWorkflowVersionsRequest) SetClusterId(v string) *ListWorkflowVersionsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListWorkflowVersionsRequest) SetMaxResults(v int32) *ListWorkflowVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowVersionsRequest) SetNextToken(v string) *ListWorkflowVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowVersionsRequest) SetWorkflowId(v int64) *ListWorkflowVersionsRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowVersionsRequest) Validate() error {
	return dara.Validate(s)
}
