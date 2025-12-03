// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPipelineGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *JoinPipelineGroupRequest
	GetGroupId() *int64
	SetPipelineIds(v string) *JoinPipelineGroupRequest
	GetPipelineIds() *string
}

type JoinPipelineGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 122,122
	PipelineIds *string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty"`
}

func (s JoinPipelineGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinPipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *JoinPipelineGroupRequest) GetPipelineIds() *string {
	return s.PipelineIds
}

func (s *JoinPipelineGroupRequest) SetGroupId(v int64) *JoinPipelineGroupRequest {
	s.GroupId = &v
	return s
}

func (s *JoinPipelineGroupRequest) SetPipelineIds(v string) *JoinPipelineGroupRequest {
	s.PipelineIds = &v
	return s
}

func (s *JoinPipelineGroupRequest) Validate() error {
	return dara.Validate(s)
}
