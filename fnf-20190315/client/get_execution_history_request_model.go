// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionName(v string) *GetExecutionHistoryRequest
	GetExecutionName() *string
	SetFlowName(v string) *GetExecutionHistoryRequest
	GetFlowName() *string
	SetLimit(v int32) *GetExecutionHistoryRequest
	GetLimit() *int32
	SetNextToken(v string) *GetExecutionHistoryRequest
	GetNextToken() *string
}

type GetExecutionHistoryRequest struct {
	// The name of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of workflows that you want to query. Valid values: 1-999. Default value: 60.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event to start the query. You can obtain the value from the response data.
	//
	// example:
	//
	// flow_xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetExecutionHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *GetExecutionHistoryRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *GetExecutionHistoryRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *GetExecutionHistoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetExecutionHistoryRequest) SetExecutionName(v string) *GetExecutionHistoryRequest {
	s.ExecutionName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetFlowName(v string) *GetExecutionHistoryRequest {
	s.FlowName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetLimit(v int32) *GetExecutionHistoryRequest {
	s.Limit = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetNextToken(v string) *GetExecutionHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *GetExecutionHistoryRequest) Validate() error {
	return dara.Validate(s)
}
