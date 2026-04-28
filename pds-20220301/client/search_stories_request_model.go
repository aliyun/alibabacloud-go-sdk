// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchStoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverImageThumbnailProcess(v string) *SearchStoriesRequest
	GetCoverImageThumbnailProcess() *string
	SetCoverVideoThumbnailProcess(v string) *SearchStoriesRequest
	GetCoverVideoThumbnailProcess() *string
	SetCreateTimeRange(v *SearchStoriesRequestCreateTimeRange) *SearchStoriesRequest
	GetCreateTimeRange() *SearchStoriesRequestCreateTimeRange
	SetCustomLabels(v string) *SearchStoriesRequest
	GetCustomLabels() *string
	SetDriveId(v string) *SearchStoriesRequest
	GetDriveId() *string
	SetFaceGroupIds(v []*string) *SearchStoriesRequest
	GetFaceGroupIds() []*string
	SetLimit(v int64) *SearchStoriesRequest
	GetLimit() *int64
	SetMarker(v string) *SearchStoriesRequest
	GetMarker() *string
	SetOrder(v string) *SearchStoriesRequest
	GetOrder() *string
	SetSort(v string) *SearchStoriesRequest
	GetSort() *string
	SetStoryEndTimeRange(v *SearchStoriesRequestStoryEndTimeRange) *SearchStoriesRequest
	GetStoryEndTimeRange() *SearchStoriesRequestStoryEndTimeRange
	SetStoryId(v string) *SearchStoriesRequest
	GetStoryId() *string
	SetStoryName(v string) *SearchStoriesRequest
	GetStoryName() *string
	SetStoryStartTimeRange(v *SearchStoriesRequestStoryStartTimeRange) *SearchStoriesRequest
	GetStoryStartTimeRange() *SearchStoriesRequestStoryStartTimeRange
	SetStoryType(v string) *SearchStoriesRequest
	GetStoryType() *string
	SetUrlExpireSec(v int64) *SearchStoriesRequest
	GetUrlExpireSec() *int64
	SetWithEmptyStories(v bool) *SearchStoriesRequest
	GetWithEmptyStories() *bool
}

type SearchStoriesRequest struct {
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	CoverImageThumbnailProcess *string `json:"cover_image_thumbnail_process,omitempty" xml:"cover_image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	CoverVideoThumbnailProcess *string `json:"cover_video_thumbnail_process,omitempty" xml:"cover_video_thumbnail_process,omitempty"`
	// if can be null:
	// true
	CreateTimeRange *SearchStoriesRequestCreateTimeRange `json:"create_time_range,omitempty" xml:"create_time_range,omitempty" type:"Struct"`
	// Deprecated
	//
	// example:
	//
	// key1=value1,key2!=value2
	CustomLabels *string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId      *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FaceGroupIds []*string `json:"face_group_ids,omitempty" xml:"face_group_ids,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// NWQ1Yjk4YmI1ZDODBhNDQ2Nzhl***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// CreateTime
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// if can be null:
	// true
	StoryEndTimeRange *SearchStoriesRequestStoryEndTimeRange `json:"story_end_time_range,omitempty" xml:"story_end_time_range,omitempty" type:"Struct"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId   *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// if can be null:
	// true
	StoryStartTimeRange *SearchStoriesRequestStoryStartTimeRange `json:"story_start_time_range,omitempty" xml:"story_start_time_range,omitempty" type:"Struct"`
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	WithEmptyStories *bool `json:"with_empty_stories,omitempty" xml:"with_empty_stories,omitempty"`
}

func (s SearchStoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequest) GetCoverImageThumbnailProcess() *string {
	return s.CoverImageThumbnailProcess
}

func (s *SearchStoriesRequest) GetCoverVideoThumbnailProcess() *string {
	return s.CoverVideoThumbnailProcess
}

func (s *SearchStoriesRequest) GetCreateTimeRange() *SearchStoriesRequestCreateTimeRange {
	return s.CreateTimeRange
}

func (s *SearchStoriesRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *SearchStoriesRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *SearchStoriesRequest) GetFaceGroupIds() []*string {
	return s.FaceGroupIds
}

func (s *SearchStoriesRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *SearchStoriesRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchStoriesRequest) GetOrder() *string {
	return s.Order
}

func (s *SearchStoriesRequest) GetSort() *string {
	return s.Sort
}

func (s *SearchStoriesRequest) GetStoryEndTimeRange() *SearchStoriesRequestStoryEndTimeRange {
	return s.StoryEndTimeRange
}

func (s *SearchStoriesRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *SearchStoriesRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *SearchStoriesRequest) GetStoryStartTimeRange() *SearchStoriesRequestStoryStartTimeRange {
	return s.StoryStartTimeRange
}

func (s *SearchStoriesRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *SearchStoriesRequest) GetUrlExpireSec() *int64 {
	return s.UrlExpireSec
}

func (s *SearchStoriesRequest) GetWithEmptyStories() *bool {
	return s.WithEmptyStories
}

func (s *SearchStoriesRequest) SetCoverImageThumbnailProcess(v string) *SearchStoriesRequest {
	s.CoverImageThumbnailProcess = &v
	return s
}

func (s *SearchStoriesRequest) SetCoverVideoThumbnailProcess(v string) *SearchStoriesRequest {
	s.CoverVideoThumbnailProcess = &v
	return s
}

func (s *SearchStoriesRequest) SetCreateTimeRange(v *SearchStoriesRequestCreateTimeRange) *SearchStoriesRequest {
	s.CreateTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetCustomLabels(v string) *SearchStoriesRequest {
	s.CustomLabels = &v
	return s
}

func (s *SearchStoriesRequest) SetDriveId(v string) *SearchStoriesRequest {
	s.DriveId = &v
	return s
}

func (s *SearchStoriesRequest) SetFaceGroupIds(v []*string) *SearchStoriesRequest {
	s.FaceGroupIds = v
	return s
}

func (s *SearchStoriesRequest) SetLimit(v int64) *SearchStoriesRequest {
	s.Limit = &v
	return s
}

func (s *SearchStoriesRequest) SetMarker(v string) *SearchStoriesRequest {
	s.Marker = &v
	return s
}

func (s *SearchStoriesRequest) SetOrder(v string) *SearchStoriesRequest {
	s.Order = &v
	return s
}

func (s *SearchStoriesRequest) SetSort(v string) *SearchStoriesRequest {
	s.Sort = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryEndTimeRange(v *SearchStoriesRequestStoryEndTimeRange) *SearchStoriesRequest {
	s.StoryEndTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetStoryId(v string) *SearchStoriesRequest {
	s.StoryId = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryName(v string) *SearchStoriesRequest {
	s.StoryName = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryStartTimeRange(v *SearchStoriesRequestStoryStartTimeRange) *SearchStoriesRequest {
	s.StoryStartTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetStoryType(v string) *SearchStoriesRequest {
	s.StoryType = &v
	return s
}

func (s *SearchStoriesRequest) SetUrlExpireSec(v int64) *SearchStoriesRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *SearchStoriesRequest) SetWithEmptyStories(v bool) *SearchStoriesRequest {
	s.WithEmptyStories = &v
	return s
}

func (s *SearchStoriesRequest) Validate() error {
	if s.CreateTimeRange != nil {
		if err := s.CreateTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.StoryEndTimeRange != nil {
		if err := s.StoryEndTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.StoryStartTimeRange != nil {
		if err := s.StoryStartTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchStoriesRequestCreateTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestCreateTimeRange) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesRequestCreateTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestCreateTimeRange) GetEnd() *string {
	return s.End
}

func (s *SearchStoriesRequestCreateTimeRange) GetStart() *string {
	return s.Start
}

func (s *SearchStoriesRequestCreateTimeRange) SetEnd(v string) *SearchStoriesRequestCreateTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestCreateTimeRange) SetStart(v string) *SearchStoriesRequestCreateTimeRange {
	s.Start = &v
	return s
}

func (s *SearchStoriesRequestCreateTimeRange) Validate() error {
	return dara.Validate(s)
}

type SearchStoriesRequestStoryEndTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestStoryEndTimeRange) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesRequestStoryEndTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestStoryEndTimeRange) GetEnd() *string {
	return s.End
}

func (s *SearchStoriesRequestStoryEndTimeRange) GetStart() *string {
	return s.Start
}

func (s *SearchStoriesRequestStoryEndTimeRange) SetEnd(v string) *SearchStoriesRequestStoryEndTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestStoryEndTimeRange) SetStart(v string) *SearchStoriesRequestStoryEndTimeRange {
	s.Start = &v
	return s
}

func (s *SearchStoriesRequestStoryEndTimeRange) Validate() error {
	return dara.Validate(s)
}

type SearchStoriesRequestStoryStartTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestStoryStartTimeRange) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesRequestStoryStartTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestStoryStartTimeRange) GetEnd() *string {
	return s.End
}

func (s *SearchStoriesRequestStoryStartTimeRange) GetStart() *string {
	return s.Start
}

func (s *SearchStoriesRequestStoryStartTimeRange) SetEnd(v string) *SearchStoriesRequestStoryStartTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestStoryStartTimeRange) SetStart(v string) *SearchStoriesRequestStoryStartTimeRange {
	s.Start = &v
	return s
}

func (s *SearchStoriesRequestStoryStartTimeRange) Validate() error {
	return dara.Validate(s)
}
