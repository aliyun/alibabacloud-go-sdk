// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *RunAgentTaskRequest
	GetBizRegionId() *string
	SetInstanceIds(v []*string) *RunAgentTaskRequest
	GetInstanceIds() []*string
	SetMaxSteps(v int32) *RunAgentTaskRequest
	GetMaxSteps() *int32
	SetTimeoutSeconds(v int32) *RunAgentTaskRequest
	GetTimeoutSeconds() *int32
	SetUserPrompt(v string) *RunAgentTaskRequest
	GetUserPrompt() *string
}

type RunAgentTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// This parameter is required.
	InstanceIds    []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	MaxSteps       *int32    `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	TimeoutSeconds *int32    `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// This parameter is required.
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s RunAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *RunAgentTaskRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *RunAgentTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunAgentTaskRequest) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *RunAgentTaskRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *RunAgentTaskRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *RunAgentTaskRequest) SetBizRegionId(v string) *RunAgentTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *RunAgentTaskRequest) SetInstanceIds(v []*string) *RunAgentTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *RunAgentTaskRequest) SetMaxSteps(v int32) *RunAgentTaskRequest {
	s.MaxSteps = &v
	return s
}

func (s *RunAgentTaskRequest) SetTimeoutSeconds(v int32) *RunAgentTaskRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *RunAgentTaskRequest) SetUserPrompt(v string) *RunAgentTaskRequest {
	s.UserPrompt = &v
	return s
}

func (s *RunAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
