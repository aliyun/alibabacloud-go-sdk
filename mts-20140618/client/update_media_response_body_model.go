// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMedia(v *UpdateMediaResponseBodyMedia) *UpdateMediaResponseBody
	GetMedia() *UpdateMediaResponseBodyMedia
	SetRequestId(v string) *UpdateMediaResponseBody
	GetRequestId() *string
}

type UpdateMediaResponseBody struct {
	// The information about the media file.
	Media *UpdateMediaResponseBodyMedia `json:"Media,omitempty" xml:"Media,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6A88246F-C91F-42BD-BABE-DB0DF993F960
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBody) GetMedia() *UpdateMediaResponseBodyMedia {
	return s.Media
}

func (s *UpdateMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaResponseBody) SetMedia(v *UpdateMediaResponseBodyMedia) *UpdateMediaResponseBody {
	s.Media = v
	return s
}

func (s *UpdateMediaResponseBody) SetRequestId(v string) *UpdateMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaResponseBody) Validate() error {
	if s.Media != nil {
		if err := s.Media.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaResponseBodyMedia struct {
	// The bitrate of the media file.
	//
	// example:
	//
	// 2659.326
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The ID of the category to which the media file belongs.
	//
	// example:
	//
	// 1
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The review state of the media file. Valid values:
	//
	// 	- **Initiated**: The media file is uploaded but not reviewed.
	//
	// 	- **Pass**: The media file is uploaded and passes the review.
	//
	// example:
	//
	// Initiated
	CensorState *string `json:"CensorState,omitempty" xml:"CensorState,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-shanghai.aliyuncs.com/example-****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media file was created.
	//
	// example:
	//
	// 2016-09-14T08:30:33Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the media file.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 7.965000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The information about the input file.
	File *UpdateMediaResponseBodyMediaFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The format of the media file. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mov
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the media file.
	//
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the media file.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The publishing state of the media file. Valid values:
	//
	// 	- **Initiated**: The media file is in the initial state.
	//
	// 	- **UnPublish**: The media file has not been published, and the playback permission on the OSS object is Private.
	//
	// 	- **Published**: The media file has been published, and the playback permission on the OSS object is Default.
	//
	// 	- **Deleted**: The media file is deleted.
	//
	// example:
	//
	// Published
	PublishState *string `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// The IDs of the media workflow execution instances.
	RunIdList *UpdateMediaResponseBodyMediaRunIdList `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Struct"`
	// The size of the media file.
	//
	// example:
	//
	// 2647692
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The information about the tags.
	Tags *UpdateMediaResponseBodyMediaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The title of the media file.
	//
	// example:
	//
	// hello
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The width of the media file.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateMediaResponseBodyMedia) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBodyMedia) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBodyMedia) GetBitrate() *string {
	return s.Bitrate
}

func (s *UpdateMediaResponseBodyMedia) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateMediaResponseBodyMedia) GetCensorState() *string {
	return s.CensorState
}

func (s *UpdateMediaResponseBodyMedia) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateMediaResponseBodyMedia) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateMediaResponseBodyMedia) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaResponseBodyMedia) GetDuration() *string {
	return s.Duration
}

func (s *UpdateMediaResponseBodyMedia) GetFile() *UpdateMediaResponseBodyMediaFile {
	return s.File
}

func (s *UpdateMediaResponseBodyMedia) GetFormat() *string {
	return s.Format
}

func (s *UpdateMediaResponseBodyMedia) GetFps() *string {
	return s.Fps
}

func (s *UpdateMediaResponseBodyMedia) GetHeight() *string {
	return s.Height
}

func (s *UpdateMediaResponseBodyMedia) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaResponseBodyMedia) GetPublishState() *string {
	return s.PublishState
}

func (s *UpdateMediaResponseBodyMedia) GetRunIdList() *UpdateMediaResponseBodyMediaRunIdList {
	return s.RunIdList
}

func (s *UpdateMediaResponseBodyMedia) GetSize() *string {
	return s.Size
}

func (s *UpdateMediaResponseBodyMedia) GetTags() *UpdateMediaResponseBodyMediaTags {
	return s.Tags
}

func (s *UpdateMediaResponseBodyMedia) GetTitle() *string {
	return s.Title
}

func (s *UpdateMediaResponseBodyMedia) GetWidth() *string {
	return s.Width
}

func (s *UpdateMediaResponseBodyMedia) SetBitrate(v string) *UpdateMediaResponseBodyMedia {
	s.Bitrate = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetCateId(v int64) *UpdateMediaResponseBodyMedia {
	s.CateId = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetCensorState(v string) *UpdateMediaResponseBodyMedia {
	s.CensorState = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetCoverURL(v string) *UpdateMediaResponseBodyMedia {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetCreationTime(v string) *UpdateMediaResponseBodyMedia {
	s.CreationTime = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetDescription(v string) *UpdateMediaResponseBodyMedia {
	s.Description = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetDuration(v string) *UpdateMediaResponseBodyMedia {
	s.Duration = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetFile(v *UpdateMediaResponseBodyMediaFile) *UpdateMediaResponseBodyMedia {
	s.File = v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetFormat(v string) *UpdateMediaResponseBodyMedia {
	s.Format = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetFps(v string) *UpdateMediaResponseBodyMedia {
	s.Fps = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetHeight(v string) *UpdateMediaResponseBodyMedia {
	s.Height = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetMediaId(v string) *UpdateMediaResponseBodyMedia {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetPublishState(v string) *UpdateMediaResponseBodyMedia {
	s.PublishState = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetRunIdList(v *UpdateMediaResponseBodyMediaRunIdList) *UpdateMediaResponseBodyMedia {
	s.RunIdList = v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetSize(v string) *UpdateMediaResponseBodyMedia {
	s.Size = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetTags(v *UpdateMediaResponseBodyMediaTags) *UpdateMediaResponseBodyMedia {
	s.Tags = v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetTitle(v string) *UpdateMediaResponseBodyMedia {
	s.Title = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) SetWidth(v string) *UpdateMediaResponseBodyMedia {
	s.Width = &v
	return s
}

func (s *UpdateMediaResponseBodyMedia) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.RunIdList != nil {
		if err := s.RunIdList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMediaResponseBodyMediaFile struct {
	// The state of the input file. Valid values:
	//
	// 	- **Normal**: The input file is normal.
	//
	// 	- **Deleted**: The input file is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The name of the OSS bucket in which the input media file is stored.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s UpdateMediaResponseBodyMediaFile) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBodyMediaFile) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBodyMediaFile) GetState() *string {
	return s.State
}

func (s *UpdateMediaResponseBodyMediaFile) GetURL() *string {
	return s.URL
}

func (s *UpdateMediaResponseBodyMediaFile) SetState(v string) *UpdateMediaResponseBodyMediaFile {
	s.State = &v
	return s
}

func (s *UpdateMediaResponseBodyMediaFile) SetURL(v string) *UpdateMediaResponseBodyMediaFile {
	s.URL = &v
	return s
}

func (s *UpdateMediaResponseBodyMediaFile) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaResponseBodyMediaRunIdList struct {
	RunId []*string `json:"RunId,omitempty" xml:"RunId,omitempty" type:"Repeated"`
}

func (s UpdateMediaResponseBodyMediaRunIdList) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBodyMediaRunIdList) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBodyMediaRunIdList) GetRunId() []*string {
	return s.RunId
}

func (s *UpdateMediaResponseBodyMediaRunIdList) SetRunId(v []*string) *UpdateMediaResponseBodyMediaRunIdList {
	s.RunId = v
	return s
}

func (s *UpdateMediaResponseBodyMediaRunIdList) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaResponseBodyMediaTags struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateMediaResponseBodyMediaTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBodyMediaTags) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBodyMediaTags) GetTag() []*string {
	return s.Tag
}

func (s *UpdateMediaResponseBodyMediaTags) SetTag(v []*string) *UpdateMediaResponseBodyMediaTags {
	s.Tag = v
	return s
}

func (s *UpdateMediaResponseBodyMediaTags) Validate() error {
	return dara.Validate(s)
}
