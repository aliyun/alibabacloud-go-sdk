// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoPreviewPlayInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *VideoPreviewPlayInfo
	GetCategory() *string
	SetLiveTranscodingSubtitleTaskList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo
	GetLiveTranscodingSubtitleTaskList() []*VideoPreviewSubtitleInfo
	SetLiveTranscodingTaskList(v []*VideoPreviewPlayInfoLiveTranscodingTaskList) *VideoPreviewPlayInfo
	GetLiveTranscodingTaskList() []*VideoPreviewPlayInfoLiveTranscodingTaskList
	SetMasterUrl(v string) *VideoPreviewPlayInfo
	GetMasterUrl() *string
	SetMeta(v *VideoPreviewPlayInfoMeta) *VideoPreviewPlayInfo
	GetMeta() *VideoPreviewPlayInfoMeta
	SetOfflineVideoTranscodingList(v []*VideoPreviewPlayInfoOfflineVideoTranscodingList) *VideoPreviewPlayInfo
	GetOfflineVideoTranscodingList() []*VideoPreviewPlayInfoOfflineVideoTranscodingList
	SetOfflineVideoTranscodingSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo
	GetOfflineVideoTranscodingSubtitleList() []*VideoPreviewSubtitleInfo
	SetQuickVideoList(v []*VideoPreviewPlayInfoQuickVideoList) *VideoPreviewPlayInfo
	GetQuickVideoList() []*VideoPreviewPlayInfoQuickVideoList
	SetQuickVideoSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo
	GetQuickVideoSubtitleList() []*VideoPreviewSubtitleInfo
}

type VideoPreviewPlayInfo struct {
	// Category
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Live transcoding subtitle information.
	LiveTranscodingSubtitleTaskList []*VideoPreviewSubtitleInfo `json:"live_transcoding_subtitle_task_list,omitempty" xml:"live_transcoding_subtitle_task_list,omitempty" type:"Repeated"`
	// The information about video playback.
	LiveTranscodingTaskList []*VideoPreviewPlayInfoLiveTranscodingTaskList `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	// Playback URL of master m3u8.
	//
	// example:
	//
	// https://pds-xxx-valueadd.oss-xxx.aliyuncs.com/qv/xxx/master.m3u8
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// Video meta information.
	Meta *VideoPreviewPlayInfoMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// Offline transcoding playback information.
	OfflineVideoTranscodingList []*VideoPreviewPlayInfoOfflineVideoTranscodingList `json:"offline_video_transcoding_list,omitempty" xml:"offline_video_transcoding_list,omitempty" type:"Repeated"`
	// Offline transcoding subtitle information.
	OfflineVideoTranscodingSubtitleList []*VideoPreviewSubtitleInfo `json:"offline_video_transcoding_subtitle_list,omitempty" xml:"offline_video_transcoding_subtitle_list,omitempty" type:"Repeated"`
	// The information about video playback.
	QuickVideoList []*VideoPreviewPlayInfoQuickVideoList `json:"quick_video_list,omitempty" xml:"quick_video_list,omitempty" type:"Repeated"`
	// Quick transcoding subtitle information.
	QuickVideoSubtitleList []*VideoPreviewSubtitleInfo `json:"quick_video_subtitle_list,omitempty" xml:"quick_video_subtitle_list,omitempty" type:"Repeated"`
}

func (s VideoPreviewPlayInfo) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayInfo) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfo) GetCategory() *string {
	return s.Category
}

func (s *VideoPreviewPlayInfo) GetLiveTranscodingSubtitleTaskList() []*VideoPreviewSubtitleInfo {
	return s.LiveTranscodingSubtitleTaskList
}

func (s *VideoPreviewPlayInfo) GetLiveTranscodingTaskList() []*VideoPreviewPlayInfoLiveTranscodingTaskList {
	return s.LiveTranscodingTaskList
}

func (s *VideoPreviewPlayInfo) GetMasterUrl() *string {
	return s.MasterUrl
}

func (s *VideoPreviewPlayInfo) GetMeta() *VideoPreviewPlayInfoMeta {
	return s.Meta
}

func (s *VideoPreviewPlayInfo) GetOfflineVideoTranscodingList() []*VideoPreviewPlayInfoOfflineVideoTranscodingList {
	return s.OfflineVideoTranscodingList
}

func (s *VideoPreviewPlayInfo) GetOfflineVideoTranscodingSubtitleList() []*VideoPreviewSubtitleInfo {
	return s.OfflineVideoTranscodingSubtitleList
}

func (s *VideoPreviewPlayInfo) GetQuickVideoList() []*VideoPreviewPlayInfoQuickVideoList {
	return s.QuickVideoList
}

func (s *VideoPreviewPlayInfo) GetQuickVideoSubtitleList() []*VideoPreviewSubtitleInfo {
	return s.QuickVideoSubtitleList
}

func (s *VideoPreviewPlayInfo) SetCategory(v string) *VideoPreviewPlayInfo {
	s.Category = &v
	return s
}

func (s *VideoPreviewPlayInfo) SetLiveTranscodingSubtitleTaskList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.LiveTranscodingSubtitleTaskList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetLiveTranscodingTaskList(v []*VideoPreviewPlayInfoLiveTranscodingTaskList) *VideoPreviewPlayInfo {
	s.LiveTranscodingTaskList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetMasterUrl(v string) *VideoPreviewPlayInfo {
	s.MasterUrl = &v
	return s
}

func (s *VideoPreviewPlayInfo) SetMeta(v *VideoPreviewPlayInfoMeta) *VideoPreviewPlayInfo {
	s.Meta = v
	return s
}

func (s *VideoPreviewPlayInfo) SetOfflineVideoTranscodingList(v []*VideoPreviewPlayInfoOfflineVideoTranscodingList) *VideoPreviewPlayInfo {
	s.OfflineVideoTranscodingList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetOfflineVideoTranscodingSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.OfflineVideoTranscodingSubtitleList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetQuickVideoList(v []*VideoPreviewPlayInfoQuickVideoList) *VideoPreviewPlayInfo {
	s.QuickVideoList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetQuickVideoSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.QuickVideoSubtitleList = v
	return s
}

func (s *VideoPreviewPlayInfo) Validate() error {
	if s.LiveTranscodingSubtitleTaskList != nil {
		for _, item := range s.LiveTranscodingSubtitleTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
	if s.OfflineVideoTranscodingSubtitleList != nil {
		for _, item := range s.OfflineVideoTranscodingSubtitleList {
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
	if s.QuickVideoSubtitleList != nil {
		for _, item := range s.QuickVideoSubtitleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoPreviewPlayInfoLiveTranscodingTaskList struct {
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
	// Template ID.
	//
	// example:
	//
	// 264_480p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// Playback URL.
	//
	// example:
	//
	// https://example.data.aliyunpds.com/lt/xxx/media.m3u8
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoLiveTranscodingTaskList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayInfoLiveTranscodingTaskList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) GetKeepOriginalResolution() *bool {
	return s.KeepOriginalResolution
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) GetUrl() *string {
	return s.Url
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetStatus(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetTemplateId(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetUrl(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.Url = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayInfoMeta struct {
	// Video length.
	//
	// example:
	//
	// 4.2898
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Height of the video
	//
	// example:
	//
	// 1080
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// Width of the video.
	//
	// example:
	//
	// 1920
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoPreviewPlayInfoMeta) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayInfoMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoMeta) GetDuration() *float64 {
	return s.Duration
}

func (s *VideoPreviewPlayInfoMeta) GetHeight() *int64 {
	return s.Height
}

func (s *VideoPreviewPlayInfoMeta) GetWidth() *int64 {
	return s.Width
}

func (s *VideoPreviewPlayInfoMeta) SetDuration(v float64) *VideoPreviewPlayInfoMeta {
	s.Duration = &v
	return s
}

func (s *VideoPreviewPlayInfoMeta) SetHeight(v int64) *VideoPreviewPlayInfoMeta {
	s.Height = &v
	return s
}

func (s *VideoPreviewPlayInfoMeta) SetWidth(v int64) *VideoPreviewPlayInfoMeta {
	s.Width = &v
	return s
}

func (s *VideoPreviewPlayInfoMeta) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayInfoOfflineVideoTranscodingList struct {
	// Whether the transcoded video has the same resolution as the source video.
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
	// The task status.
	//
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 264_1080p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// Playback URL.
	//
	// example:
	//
	// https://example.data.aliyunpds.com/xxx/master.mp4
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoOfflineVideoTranscodingList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayInfoOfflineVideoTranscodingList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) GetKeepOriginalResolution() *bool {
	return s.KeepOriginalResolution
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) GetUrl() *string {
	return s.Url
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetStatus(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetTemplateId(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetUrl(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.Url = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) Validate() error {
	return dara.Validate(s)
}

type VideoPreviewPlayInfoQuickVideoList struct {
	// Valid values:
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
	// 264_480p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// Playback URL.
	//
	// example:
	//
	// https://example.data.aliyunpds.com/qv/xxx/media.m3u8
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoQuickVideoList) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewPlayInfoQuickVideoList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoQuickVideoList) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewPlayInfoQuickVideoList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *VideoPreviewPlayInfoQuickVideoList) GetUrl() *string {
	return s.Url
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetStatus(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetTemplateId(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetUrl(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.Url = &v
	return s
}

func (s *VideoPreviewPlayInfoQuickVideoList) Validate() error {
	return dara.Validate(s)
}
