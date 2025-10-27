// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionNamePrefix(v string) *ListExecutionsRequest
	GetExecutionNamePrefix() *string
	SetFlowName(v string) *ListExecutionsRequest
	GetFlowName() *string
	SetLimit(v int32) *ListExecutionsRequest
	GetLimit() *int32
	SetMapRunName(v string) *ListExecutionsRequest
	GetMapRunName() *string
	SetMetadataOnly(v bool) *ListExecutionsRequest
	GetMetadataOnly() *bool
	SetNextToken(v string) *ListExecutionsRequest
	GetNextToken() *string
	SetQualifier(v string) *ListExecutionsRequest
	GetQualifier() *string
	SetStartedTimeBegin(v string) *ListExecutionsRequest
	GetStartedTimeBegin() *string
	SetStartedTimeEnd(v string) *ListExecutionsRequest
	GetStartedTimeEnd() *string
	SetStatus(v string) *ListExecutionsRequest
	GetStatus() *string
}

type ListExecutionsRequest struct {
	// The name prefix of the execution.
	//
	// example:
	//
	// run
	ExecutionNamePrefix *string `json:"ExecutionNamePrefix,omitempty" xml:"ExecutionNamePrefix,omitempty"`
	// The name of the flow. The name must be unique within the region and cannot be modified after the flow is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of executions that you want to query. Valid values: 1-99. Default value: 60.
	//
	// example:
	//
	// 1
	Limit        *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MapRunName   *string `json:"MapRunName,omitempty" xml:"MapRunName,omitempty"`
	MetadataOnly *bool   `json:"MetadataOnly,omitempty" xml:"MetadataOnly,omitempty"`
	// The name of the execution to start the query. You can obtain the value from the response data. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// flow_xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	Qualifier *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	// The beginning of the time range to query executions. Specify the value in the UTC RFC3339 format.
	//
	// example:
	//
	// 2020-12-02T02:39:20.402Z
	StartedTimeBegin *string `json:"StartedTimeBegin,omitempty" xml:"StartedTimeBegin,omitempty"`
	// The end of the time range to query executions. Specify the value in the UTC RFC3339 format.
	//
	// example:
	//
	// 2020-12-02T02:23:54.817Z
	StartedTimeEnd *string `json:"StartedTimeEnd,omitempty" xml:"StartedTimeEnd,omitempty"`
	// The status of the execution that you want to filter. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) GetExecutionNamePrefix() *string {
	return s.ExecutionNamePrefix
}

func (s *ListExecutionsRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListExecutionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListExecutionsRequest) GetMapRunName() *string {
	return s.MapRunName
}

func (s *ListExecutionsRequest) GetMetadataOnly() *bool {
	return s.MetadataOnly
}

func (s *ListExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionsRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListExecutionsRequest) GetStartedTimeBegin() *string {
	return s.StartedTimeBegin
}

func (s *ListExecutionsRequest) GetStartedTimeEnd() *string {
	return s.StartedTimeEnd
}

func (s *ListExecutionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExecutionsRequest) SetExecutionNamePrefix(v string) *ListExecutionsRequest {
	s.ExecutionNamePrefix = &v
	return s
}

func (s *ListExecutionsRequest) SetFlowName(v string) *ListExecutionsRequest {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsRequest) SetLimit(v int32) *ListExecutionsRequest {
	s.Limit = &v
	return s
}

func (s *ListExecutionsRequest) SetMapRunName(v string) *ListExecutionsRequest {
	s.MapRunName = &v
	return s
}

func (s *ListExecutionsRequest) SetMetadataOnly(v bool) *ListExecutionsRequest {
	s.MetadataOnly = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsRequest) SetQualifier(v string) *ListExecutionsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListExecutionsRequest) SetStartedTimeBegin(v string) *ListExecutionsRequest {
	s.StartedTimeBegin = &v
	return s
}

func (s *ListExecutionsRequest) SetStartedTimeEnd(v string) *ListExecutionsRequest {
	s.StartedTimeEnd = &v
	return s
}

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
