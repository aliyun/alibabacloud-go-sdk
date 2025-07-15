// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncEditTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoClips(v bool) *AsyncEditTimelineRequest
	GetAutoClips() *bool
	SetTaskId(v string) *AsyncEditTimelineRequest
	GetTaskId() *string
	SetTimelines(v []*AsyncEditTimelineRequestTimelines) *AsyncEditTimelineRequest
	GetTimelines() []*AsyncEditTimelineRequestTimelines
	SetWorkspaceId(v string) *AsyncEditTimelineRequest
	GetWorkspaceId() *string
}

type AsyncEditTimelineRequest struct {
	AutoClips *bool `json:"AutoClips,omitempty" xml:"AutoClips,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0dbf1055f8a2475d99904c3b76a0ffba
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	Timelines []*AsyncEditTimelineRequestTimelines `json:"Timelines,omitempty" xml:"Timelines,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncEditTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineRequest) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineRequest) GetAutoClips() *bool {
	return s.AutoClips
}

func (s *AsyncEditTimelineRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncEditTimelineRequest) GetTimelines() []*AsyncEditTimelineRequestTimelines {
	return s.Timelines
}

func (s *AsyncEditTimelineRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncEditTimelineRequest) SetAutoClips(v bool) *AsyncEditTimelineRequest {
	s.AutoClips = &v
	return s
}

func (s *AsyncEditTimelineRequest) SetTaskId(v string) *AsyncEditTimelineRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncEditTimelineRequest) SetTimelines(v []*AsyncEditTimelineRequestTimelines) *AsyncEditTimelineRequest {
	s.Timelines = v
	return s
}

func (s *AsyncEditTimelineRequest) SetWorkspaceId(v string) *AsyncEditTimelineRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncEditTimelineRequest) Validate() error {
	return dara.Validate(s)
}

type AsyncEditTimelineRequestTimelines struct {
	// This parameter is required.
	Clips      []*AsyncEditTimelineRequestTimelinesClips `json:"Clips,omitempty" xml:"Clips,omitempty" type:"Repeated"`
	TimelineId *string                                   `json:"TimelineId,omitempty" xml:"TimelineId,omitempty"`
}

func (s AsyncEditTimelineRequestTimelines) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineRequestTimelines) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineRequestTimelines) GetClips() []*AsyncEditTimelineRequestTimelinesClips {
	return s.Clips
}

func (s *AsyncEditTimelineRequestTimelines) GetTimelineId() *string {
	return s.TimelineId
}

func (s *AsyncEditTimelineRequestTimelines) SetClips(v []*AsyncEditTimelineRequestTimelinesClips) *AsyncEditTimelineRequestTimelines {
	s.Clips = v
	return s
}

func (s *AsyncEditTimelineRequestTimelines) SetTimelineId(v string) *AsyncEditTimelineRequestTimelines {
	s.TimelineId = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelines) Validate() error {
	return dara.Validate(s)
}

type AsyncEditTimelineRequestTimelinesClips struct {
	ClipId       *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	ContentInner *string `json:"ContentInner,omitempty" xml:"ContentInner,omitempty"`
	In           *int32  `json:"In,omitempty" xml:"In,omitempty"`
	Out          *int32  `json:"Out,omitempty" xml:"Out,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	VideoName    *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
}

func (s AsyncEditTimelineRequestTimelinesClips) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineRequestTimelinesClips) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetClipId() *string {
	return s.ClipId
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetContentInner() *string {
	return s.ContentInner
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetIn() *int32 {
	return s.In
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetOut() *int32 {
	return s.Out
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetVideoId() *string {
	return s.VideoId
}

func (s *AsyncEditTimelineRequestTimelinesClips) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetClipId(v string) *AsyncEditTimelineRequestTimelinesClips {
	s.ClipId = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetContentInner(v string) *AsyncEditTimelineRequestTimelinesClips {
	s.ContentInner = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetIn(v int32) *AsyncEditTimelineRequestTimelinesClips {
	s.In = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetOut(v int32) *AsyncEditTimelineRequestTimelinesClips {
	s.Out = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetVideoId(v string) *AsyncEditTimelineRequestTimelinesClips {
	s.VideoId = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) SetVideoName(v string) *AsyncEditTimelineRequestTimelinesClips {
	s.VideoName = &v
	return s
}

func (s *AsyncEditTimelineRequestTimelinesClips) Validate() error {
	return dara.Validate(s)
}
