// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDebugExecutionRequest
	GetClientToken() *string
	SetProperties(v string) *StartDebugExecutionRequest
	GetProperties() *string
	SetRegionId(v string) *StartDebugExecutionRequest
	GetRegionId() *string
	SetTaskType(v string) *StartDebugExecutionRequest
	GetTaskType() *string
}

type StartDebugExecutionRequest struct {
	// example:
	//
	// abcde3OARpx77No54nv6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {"Service": "ecs", "API": "DescribeRegions"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ACS::ExecuteAPI
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s StartDebugExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDebugExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartDebugExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDebugExecutionRequest) GetProperties() *string {
	return s.Properties
}

func (s *StartDebugExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDebugExecutionRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *StartDebugExecutionRequest) SetClientToken(v string) *StartDebugExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDebugExecutionRequest) SetProperties(v string) *StartDebugExecutionRequest {
	s.Properties = &v
	return s
}

func (s *StartDebugExecutionRequest) SetRegionId(v string) *StartDebugExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *StartDebugExecutionRequest) SetTaskType(v string) *StartDebugExecutionRequest {
	s.TaskType = &v
	return s
}

func (s *StartDebugExecutionRequest) Validate() error {
	return dara.Validate(s)
}
