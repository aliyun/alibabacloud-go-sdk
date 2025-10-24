// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMedia(v *AddMediaResponseBodyMedia) *AddMediaResponseBody
	GetMedia() *AddMediaResponseBodyMedia
	SetRequestId(v string) *AddMediaResponseBody
	GetRequestId() *string
}

type AddMediaResponseBody struct {
	// The detailed information about the media file.
	Media *AddMediaResponseBodyMedia `json:"Media,omitempty" xml:"Media,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 05F8B913-E9F3-4A6F-9922-48CADA0FFAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaResponseBody) GetMedia() *AddMediaResponseBodyMedia {
	return s.Media
}

func (s *AddMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaResponseBody) SetMedia(v *AddMediaResponseBodyMedia) *AddMediaResponseBody {
	s.Media = v
	return s
}

func (s *AddMediaResponseBody) SetRequestId(v string) *AddMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaResponseBody) Validate() error {
	if s.Media != nil {
		if err := s.Media.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMediaResponseBodyMedia struct {
	// The bitrate.
	//
	// example:
	//
	// 1148.77
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The ID of the category to which the media file belongs.
	//
	// example:
	//
	// 1
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The review status of the media file. Valid values:
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
	// http://bucket.oss-cn-hangzhou.aliyuncs.com/example/1.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media file was created.
	//
	// example:
	//
	// 2016-09-20T03:02:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the media file. The description can be up to 1,024 bytes in length.
	//
	// example:
	//
	// A test video
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 2.645333
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The information about the input file.
	File *AddMediaResponseBodyMediaFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The format of the media file. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mp4
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
	// 1280
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// 3e6149d5a8c944c09b1a8d2dc3e4****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The publishing status of the media file. Valid values:
	//
	// 	- **Initiated**: The media file is in the initial state.
	//
	// 	- **UnPublish**: The media file has not been published, and the playback permission on the OSS object is Private.
	//
	// 	- **Published**: The media file has been published, and the playback permission on the OSS object is Default.
	//
	// example:
	//
	// Published
	PublishState *string `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// The IDs of the media workflow execution instances.
	RunIdList *AddMediaResponseBodyMediaRunIdList `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Struct"`
	// The size of the media file.
	//
	// example:
	//
	// 379860
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The tags of the media file.
	Tags *AddMediaResponseBodyMediaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The title of the media file. The title can be up to 128 bytes in length.
	//
	// example:
	//
	// mytest.mp4
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The width of the media file.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddMediaResponseBodyMedia) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponseBodyMedia) GoString() string {
	return s.String()
}

func (s *AddMediaResponseBodyMedia) GetBitrate() *string {
	return s.Bitrate
}

func (s *AddMediaResponseBodyMedia) GetCateId() *int64 {
	return s.CateId
}

func (s *AddMediaResponseBodyMedia) GetCensorState() *string {
	return s.CensorState
}

func (s *AddMediaResponseBodyMedia) GetCoverURL() *string {
	return s.CoverURL
}

func (s *AddMediaResponseBodyMedia) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AddMediaResponseBodyMedia) GetDescription() *string {
	return s.Description
}

func (s *AddMediaResponseBodyMedia) GetDuration() *string {
	return s.Duration
}

func (s *AddMediaResponseBodyMedia) GetFile() *AddMediaResponseBodyMediaFile {
	return s.File
}

func (s *AddMediaResponseBodyMedia) GetFormat() *string {
	return s.Format
}

func (s *AddMediaResponseBodyMedia) GetFps() *string {
	return s.Fps
}

func (s *AddMediaResponseBodyMedia) GetHeight() *string {
	return s.Height
}

func (s *AddMediaResponseBodyMedia) GetMediaId() *string {
	return s.MediaId
}

func (s *AddMediaResponseBodyMedia) GetPublishState() *string {
	return s.PublishState
}

func (s *AddMediaResponseBodyMedia) GetRunIdList() *AddMediaResponseBodyMediaRunIdList {
	return s.RunIdList
}

func (s *AddMediaResponseBodyMedia) GetSize() *string {
	return s.Size
}

func (s *AddMediaResponseBodyMedia) GetTags() *AddMediaResponseBodyMediaTags {
	return s.Tags
}

func (s *AddMediaResponseBodyMedia) GetTitle() *string {
	return s.Title
}

func (s *AddMediaResponseBodyMedia) GetWidth() *string {
	return s.Width
}

func (s *AddMediaResponseBodyMedia) SetBitrate(v string) *AddMediaResponseBodyMedia {
	s.Bitrate = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetCateId(v int64) *AddMediaResponseBodyMedia {
	s.CateId = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetCensorState(v string) *AddMediaResponseBodyMedia {
	s.CensorState = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetCoverURL(v string) *AddMediaResponseBodyMedia {
	s.CoverURL = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetCreationTime(v string) *AddMediaResponseBodyMedia {
	s.CreationTime = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetDescription(v string) *AddMediaResponseBodyMedia {
	s.Description = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetDuration(v string) *AddMediaResponseBodyMedia {
	s.Duration = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetFile(v *AddMediaResponseBodyMediaFile) *AddMediaResponseBodyMedia {
	s.File = v
	return s
}

func (s *AddMediaResponseBodyMedia) SetFormat(v string) *AddMediaResponseBodyMedia {
	s.Format = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetFps(v string) *AddMediaResponseBodyMedia {
	s.Fps = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetHeight(v string) *AddMediaResponseBodyMedia {
	s.Height = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetMediaId(v string) *AddMediaResponseBodyMedia {
	s.MediaId = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetPublishState(v string) *AddMediaResponseBodyMedia {
	s.PublishState = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetRunIdList(v *AddMediaResponseBodyMediaRunIdList) *AddMediaResponseBodyMedia {
	s.RunIdList = v
	return s
}

func (s *AddMediaResponseBodyMedia) SetSize(v string) *AddMediaResponseBodyMedia {
	s.Size = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetTags(v *AddMediaResponseBodyMediaTags) *AddMediaResponseBodyMedia {
	s.Tags = v
	return s
}

func (s *AddMediaResponseBodyMedia) SetTitle(v string) *AddMediaResponseBodyMedia {
	s.Title = &v
	return s
}

func (s *AddMediaResponseBodyMedia) SetWidth(v string) *AddMediaResponseBodyMedia {
	s.Width = &v
	return s
}

func (s *AddMediaResponseBodyMedia) Validate() error {
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

type AddMediaResponseBodyMediaFile struct {
	// The status of the file. The default value is **Normal**.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The URL of the media file.
	//
	// example:
	//
	// http://bucket.oss-cn-hangzhou.aliyuncs.com/A/B/C/test.mp4
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s AddMediaResponseBodyMediaFile) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponseBodyMediaFile) GoString() string {
	return s.String()
}

func (s *AddMediaResponseBodyMediaFile) GetState() *string {
	return s.State
}

func (s *AddMediaResponseBodyMediaFile) GetURL() *string {
	return s.URL
}

func (s *AddMediaResponseBodyMediaFile) SetState(v string) *AddMediaResponseBodyMediaFile {
	s.State = &v
	return s
}

func (s *AddMediaResponseBodyMediaFile) SetURL(v string) *AddMediaResponseBodyMediaFile {
	s.URL = &v
	return s
}

func (s *AddMediaResponseBodyMediaFile) Validate() error {
	return dara.Validate(s)
}

type AddMediaResponseBodyMediaRunIdList struct {
	RunId []*string `json:"RunId,omitempty" xml:"RunId,omitempty" type:"Repeated"`
}

func (s AddMediaResponseBodyMediaRunIdList) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponseBodyMediaRunIdList) GoString() string {
	return s.String()
}

func (s *AddMediaResponseBodyMediaRunIdList) GetRunId() []*string {
	return s.RunId
}

func (s *AddMediaResponseBodyMediaRunIdList) SetRunId(v []*string) *AddMediaResponseBodyMediaRunIdList {
	s.RunId = v
	return s
}

func (s *AddMediaResponseBodyMediaRunIdList) Validate() error {
	return dara.Validate(s)
}

type AddMediaResponseBodyMediaTags struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddMediaResponseBodyMediaTags) String() string {
	return dara.Prettify(s)
}

func (s AddMediaResponseBodyMediaTags) GoString() string {
	return s.String()
}

func (s *AddMediaResponseBodyMediaTags) GetTag() []*string {
	return s.Tag
}

func (s *AddMediaResponseBodyMediaTags) SetTag(v []*string) *AddMediaResponseBodyMediaTags {
	s.Tag = v
	return s
}

func (s *AddMediaResponseBodyMediaTags) Validate() error {
	return dara.Validate(s)
}
