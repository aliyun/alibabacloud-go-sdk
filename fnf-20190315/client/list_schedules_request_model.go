// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *ListSchedulesRequest
	GetFlowName() *string
	SetLimit(v int32) *ListSchedulesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListSchedulesRequest
	GetNextToken() *string
}

type ListSchedulesRequest struct {
	// The name of the flow that is associated with the time-based schedules. The name is unique within the region and cannot be modified after the flow is created. The name must meet the following conventions:
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
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of schedules that you want to query. Valid values: 1 to 1000.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// For the first query, you do not need to specify this parameter. The system uses the value of the **FlowName*	- parameter as the value of the **NextToken*	- parameter. When the query ends, no value is returned for this parameter.
	//
	// example:
	//
	// testNextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListSchedulesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListSchedulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchedulesRequest) SetFlowName(v string) *ListSchedulesRequest {
	s.FlowName = &v
	return s
}

func (s *ListSchedulesRequest) SetLimit(v int32) *ListSchedulesRequest {
	s.Limit = &v
	return s
}

func (s *ListSchedulesRequest) SetNextToken(v string) *ListSchedulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
