// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *DescribeScheduleRequest
	GetFlowName() *string
	SetScheduleName(v string) *DescribeScheduleRequest
	GetScheduleName() *string
}

type DescribeScheduleRequest struct {
	// The name of the flow that is associated with the time-based schedule. The name must be unique within the region and cannot be modified after the time-based schedule is created. The name must meet the following conventions:
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
	// The name of the time-based schedule. The name must meet the following conventions:
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
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DescribeScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribeScheduleRequest) GetScheduleName() *string {
	return s.ScheduleName
}

func (s *DescribeScheduleRequest) SetFlowName(v string) *DescribeScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeScheduleRequest) SetScheduleName(v string) *DescribeScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *DescribeScheduleRequest) Validate() error {
	return dara.Validate(s)
}
