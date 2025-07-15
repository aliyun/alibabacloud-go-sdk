// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncEditTimelineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoClips(v bool) *AsyncEditTimelineShrinkRequest
	GetAutoClips() *bool
	SetTaskId(v string) *AsyncEditTimelineShrinkRequest
	GetTaskId() *string
	SetTimelinesShrink(v string) *AsyncEditTimelineShrinkRequest
	GetTimelinesShrink() *string
	SetWorkspaceId(v string) *AsyncEditTimelineShrinkRequest
	GetWorkspaceId() *string
}

type AsyncEditTimelineShrinkRequest struct {
	AutoClips *bool `json:"AutoClips,omitempty" xml:"AutoClips,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0dbf1055f8a2475d99904c3b76a0ffba
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	TimelinesShrink *string `json:"Timelines,omitempty" xml:"Timelines,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncEditTimelineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineShrinkRequest) GetAutoClips() *bool {
	return s.AutoClips
}

func (s *AsyncEditTimelineShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncEditTimelineShrinkRequest) GetTimelinesShrink() *string {
	return s.TimelinesShrink
}

func (s *AsyncEditTimelineShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncEditTimelineShrinkRequest) SetAutoClips(v bool) *AsyncEditTimelineShrinkRequest {
	s.AutoClips = &v
	return s
}

func (s *AsyncEditTimelineShrinkRequest) SetTaskId(v string) *AsyncEditTimelineShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncEditTimelineShrinkRequest) SetTimelinesShrink(v string) *AsyncEditTimelineShrinkRequest {
	s.TimelinesShrink = &v
	return s
}

func (s *AsyncEditTimelineShrinkRequest) SetWorkspaceId(v string) *AsyncEditTimelineShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncEditTimelineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
