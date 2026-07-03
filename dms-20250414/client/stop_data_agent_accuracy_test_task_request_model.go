// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataAgentAccuracyTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestTaskId(v string) *StopDataAgentAccuracyTestTaskRequest
	GetAccuracyTestTaskId() *string
	SetRegionId(v string) *StopDataAgentAccuracyTestTaskRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *StopDataAgentAccuracyTestTaskRequest
	GetWorkspaceId() *string
}

type StopDataAgentAccuracyTestTaskRequest struct {
	// The ID of the accuracy test task.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StopDataAgentAccuracyTestTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataAgentAccuracyTestTaskRequest) GoString() string {
	return s.String()
}

func (s *StopDataAgentAccuracyTestTaskRequest) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *StopDataAgentAccuracyTestTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDataAgentAccuracyTestTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *StopDataAgentAccuracyTestTaskRequest) SetAccuracyTestTaskId(v string) *StopDataAgentAccuracyTestTaskRequest {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskRequest) SetRegionId(v string) *StopDataAgentAccuracyTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskRequest) SetWorkspaceId(v string) *StopDataAgentAccuracyTestTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskRequest) Validate() error {
	return dara.Validate(s)
}
