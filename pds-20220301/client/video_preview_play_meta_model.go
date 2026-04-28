// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoPreviewPlayMeta interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *VideoPreviewPlayMeta
	GetCategory() *string
	SetLiveTranscodingTaskList(v []*VideoPreviewPlayMetaLiveTranscodingTaskList) *VideoPreviewPlayMeta
	GetLiveTranscodingTaskList() []*VideoPreviewPlayMetaLiveTranscodingTaskList
	SetMeta(v *VideoPreviewPlayMetaMeta) *VideoPreviewPlayMeta
	GetMeta() *VideoPreviewPlayMetaMeta
	SetOfflineVideoTranscodingList(v []*VideoPreviewPlayMetaOfflineVideoTranscodingList) *VideoPreviewPlayMeta
	GetOfflineVideoTranscodingList() []*VideoPreviewPlayMetaOfflineVideoTranscodingList
	SetQuickVideoList(v []*VideoPreviewPlayMetaQuickVideoList) *VideoPreviewPlayMeta
	GetQuickVideoList() []*VideoPreviewPlayMetaQuickVideoList
}

type VideoPreviewPlayMeta struct {
	// Category
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Status of the live transcoding job.
	LiveTranscodingTaskList []*VideoPreviewPlayMetaLiveTranscodingTaskList `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	// Video meta information.
	Meta *VideoPreviewPlayMetaMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The status of the offline transcoding job.
	OfflineVideoTranscodingList []*VideoPreviewPlayMetaOfflineVideoTranscodingList `json:"offline_video_transcoding_list,omitempty" xml:"offline_video_transcoding_list,omitempty" type:"Repeated"`
	// The state of the transcoding job.
	QuickVideoList []*VideoPreviewPlayMetaQuickVideoList `json:"quick_video_list,omitempty" xml:"quick_video_list,omitempty" type:"Repeated"`
}

func (s VideoPreviewPlayMeta) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMeta) GetCategory() *string {
	return s.Category
}

func (s *VideoPreviewPlayMeta) GetLiveTranscodingTaskList() []*VideoPreviewPlayMetaLiveTranscodingTaskList {
	return s.LiveTranscodingTaskList
}

func (s *VideoPreviewPlayMeta) GetMeta() *VideoPreviewPlayMetaMeta {
	return s.Meta
}

func (s *VideoPreviewPlayMeta) GetOfflineVideoTranscodingList() []*VideoPreviewPlayMetaOfflineVideoTranscodingList {
	return s.OfflineVideoTranscodingList
}

func (s *VideoPreviewPlayMeta) GetQuickVideoList() []*VideoPreviewPlayMetaQuickVideoList {
	return s.QuickVideoList
}

func (s *VideoPreviewPlayMeta) SetCategory(v string) *VideoPreviewPlayMeta {
	s.Category = &v
	return s
}

func (s *VideoPreviewPlayMeta) SetLiveTranscodingTaskList(v []*VideoPreviewPlayMetaLiveTranscodingTaskList) *VideoPreviewPlayMeta {
	s.LiveTranscodingTaskList = v
	return s
}

func (s *VideoPreviewPlayMeta) SetMeta(v *VideoPreviewPlayMetaMeta) *VideoPreviewPlayMeta {
	s.Meta = v
	return s
}

func (s *VideoPreviewPlayMeta) SetOfflineVideoTranscodingList(v []*VideoPreviewPlayMetaOfflineVideoTranscodingList) *VideoPreviewPlayMeta {
	s.OfflineVideoTranscodingList = v
	return s
}

func (s *VideoPreviewPlayMeta) SetQuickVideoList(v []*VideoPreviewPlayMetaQuickVideoList) *VideoPreviewPlayMeta {
	s.QuickVideoList = v
	return s
}

func (s *VideoPreviewPlayMeta) Validate() error {
	if s.LiveTranscodingTaskList != nil {
		for _, item := range s.LiveTranscodingTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	if s.OfflineVideoTranscodingList != nil {
		for _, item := range s.OfflineVideoTranscodingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuickVideoList != nil {
		for _, item := range s.QuickVideoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoPreviewPlayMetaLiveTranscodingTaskList struct {
	// Whether the original resolution is maintained.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	KeepOriginalResolution *bool `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	// The status. Valid values:
	//
	// 	- finished: The index is complete, and the url can be obtained.
	//
	// 	- running: Indexing in progress. Wait a moment and try again.
	//
	// 	- failed: Transcoding failed. Check the media file. If you have any questions, contact customer service.
	//
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Template ID
	//
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaLiveTranscodingTaskList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayMetaLiveTranscodingTaskList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) GetKeepOriginalResolution() *bool {
	return s.KeepOriginalResolution
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetStatus(v string) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetTemplateId(v string) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayMetaMeta struct {
	// Length of the video.
	//
	// example:
	//
	// 10
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Height of the video.
	//
	// example:
	//
	// 720
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// Width of the video.
	//
	// example:
	//
	// 1280
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoPreviewPlayMetaMeta) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayMetaMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaMeta) GetDuration() *float64 {
	return s.Duration
}

func (s *VideoPreviewPlayMetaMeta) GetHeight() *int64 {
	return s.Height
}

func (s *VideoPreviewPlayMetaMeta) GetWidth() *int64 {
	return s.Width
}

func (s *VideoPreviewPlayMetaMeta) SetDuration(v float64) *VideoPreviewPlayMetaMeta {
	s.Duration = &v
	return s
}

func (s *VideoPreviewPlayMetaMeta) SetHeight(v int64) *VideoPreviewPlayMetaMeta {
	s.Height = &v
	return s
}

func (s *VideoPreviewPlayMetaMeta) SetWidth(v int64) *VideoPreviewPlayMetaMeta {
	s.Width = &v
	return s
}

func (s *VideoPreviewPlayMetaMeta) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayMetaOfflineVideoTranscodingList struct {
	// Whether the original resolution is maintained.
	//
	// example:
	//
	// true
	KeepOriginalResolution *string `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	// The status. finished: The index is completed, and the url can be obtained. running: Indexing in progress. Wait a moment and try again. failed: Transcoding failed. Check the media file. If you have any questions, contact customer service.
	//
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Template ID
	//
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaOfflineVideoTranscodingList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayMetaOfflineVideoTranscodingList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) GetKeepOriginalResolution() *string {
	return s.KeepOriginalResolution
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetKeepOriginalResolution(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetStatus(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetTemplateId(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayMetaQuickVideoList struct {
	// The status. finished: The index is completed, and the url can be obtained. running: Indexing in progress. Wait a moment and try again. failed: Transcoding failed. Check the media file. If you have any questions, contact customer service.
	//
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Template ID
	//
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaQuickVideoList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayMetaQuickVideoList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaQuickVideoList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayMetaQuickVideoList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayMetaQuickVideoList) SetStatus(v string) *VideoPreviewPlayMetaQuickVideoList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaQuickVideoList) SetTemplateId(v string) *VideoPreviewPlayMetaQuickVideoList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayMetaQuickVideoList) Validate() error {
	return dara.Validate(s)
}
