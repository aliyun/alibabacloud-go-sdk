// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageList(v []*ListDynamicImageResponseBodyDynamicImageList) *ListDynamicImageResponseBody
	GetDynamicImageList() []*ListDynamicImageResponseBodyDynamicImageList
	SetRequestId(v string) *ListDynamicImageResponseBody
	GetRequestId() *string
}

type ListDynamicImageResponseBody struct {
	// The list of animated stickers.
	DynamicImageList []*ListDynamicImageResponseBodyDynamicImageList `json:"DynamicImageList,omitempty" xml:"DynamicImageList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 570189B6-572E-4953-13B4278EE0D8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDynamicImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponseBody) GetDynamicImageList() []*ListDynamicImageResponseBodyDynamicImageList {
	return s.DynamicImageList
}

func (s *ListDynamicImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicImageResponseBody) SetDynamicImageList(v []*ListDynamicImageResponseBodyDynamicImageList) *ListDynamicImageResponseBody {
	s.DynamicImageList = v
	return s
}

func (s *ListDynamicImageResponseBody) SetRequestId(v string) *ListDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDynamicImageResponseBodyDynamicImageList struct {
	// The time when the animated sticker was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-28T02:01:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The duration of the animated sticker. Unit: seconds.
	//
	// example:
	//
	// 2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the animated sticker.
	//
	// example:
	//
	// 2b4e51df60323ef43d6e336ecf****
	DynamicImageId *string `json:"DynamicImageId,omitempty" xml:"DynamicImageId,omitempty"`
	// The size of the animated sticker file. Unit: byte.
	//
	// example:
	//
	// 119866
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The URL of the animated sticker file.
	//
	// example:
	//
	// https://example.aliyundoc.com/2e114f110059*****0c3193918fd449a/image/dynamic/2b4e51df60*****323ef43d6e336ecf.webp?auth_key=1597296785-0-0-4a48e85*****bd2bb358e0b3cade
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The format of the animated sticker. Valid values: gif and webp.
	//
	// example:
	//
	// webp
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the animated sticker. Unit: frames per second.
	//
	// example:
	//
	// 10
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the animated sticker. Unit: pixel.
	//
	// example:
	//
	// 360
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The job ID for creating the animated sticker.
	//
	// example:
	//
	// 2bf4390af9e5491c09cc720ad****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 2e114f1100590c3193918fd449a****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// The width of the animated sticker. Unit: pixel.
	//
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListDynamicImageResponseBodyDynamicImageList) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageResponseBodyDynamicImageList) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetDuration() *string {
	return s.Duration
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetDynamicImageId() *string {
	return s.DynamicImageId
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetFileSize() *string {
	return s.FileSize
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetFileURL() *string {
	return s.FileURL
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetFormat() *string {
	return s.Format
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetFps() *string {
	return s.Fps
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetHeight() *string {
	return s.Height
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetJobId() *string {
	return s.JobId
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetVideoId() *string {
	return s.VideoId
}

func (s *ListDynamicImageResponseBodyDynamicImageList) GetWidth() *string {
	return s.Width
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetCreationTime(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.CreationTime = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetDuration(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Duration = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetDynamicImageId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.DynamicImageId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFileSize(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.FileSize = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFileURL(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.FileURL = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFormat(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Format = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFps(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Fps = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetHeight(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Height = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetJobId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.JobId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetVideoId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.VideoId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetWidth(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Width = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) Validate() error {
	return dara.Validate(s)
}
