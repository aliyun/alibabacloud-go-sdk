// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *SubmitSnapshotJobShrinkRequest
	GetCount() *int64
	SetHeight(v string) *SubmitSnapshotJobShrinkRequest
	GetHeight() *string
	SetInterval(v int64) *SubmitSnapshotJobShrinkRequest
	GetInterval() *int64
	SetReferenceId(v string) *SubmitSnapshotJobShrinkRequest
	GetReferenceId() *string
	SetSnapshotTemplateId(v string) *SubmitSnapshotJobShrinkRequest
	GetSnapshotTemplateId() *string
	SetSpecifiedOffsetTime(v int64) *SubmitSnapshotJobShrinkRequest
	GetSpecifiedOffsetTime() *int64
	SetSpecifiedOffsetTimesShrink(v string) *SubmitSnapshotJobShrinkRequest
	GetSpecifiedOffsetTimesShrink() *string
	SetSpriteSnapshotConfig(v string) *SubmitSnapshotJobShrinkRequest
	GetSpriteSnapshotConfig() *string
	SetUserData(v string) *SubmitSnapshotJobShrinkRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitSnapshotJobShrinkRequest
	GetVideoId() *string
	SetWidth(v string) *SubmitSnapshotJobShrinkRequest
	GetWidth() *string
}

type SubmitSnapshotJobShrinkRequest struct {
	// The maximum number of snapshots. Default value: **1**.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The height of each snapshot. Valid values: `[8,4096]`. By default, the height of the video source is used. Unit: pixels.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The snapshot interval. The value must be **greater than or equal to 0**.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **1**.
	//
	// 	- If you set this parameter to **0**, snapshots are captured at even intervals based on the video duration divided by the value of the Count parameter.
	//
	// example:
	//
	// 1
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The ID of the snapshot template.
	//
	// 	- We recommend that you create a snapshot template before you specify the template ID. For more information about how to create a snapshot template, see [AddVodTemplate](https://help.aliyun.com/document_detail/99406.html).
	//
	// 	- If you set the SnapshotTemplateId parameter, all the other request parameters except the Action and VideoId parameters are ignored.
	//
	// example:
	//
	// f5b228fe693bf55bd87b789****
	SnapshotTemplateId *string `json:"SnapshotTemplateId,omitempty" xml:"SnapshotTemplateId,omitempty"`
	// The point in time when the first snapshot is captured.
	//
	// 	- Unit: milliseconds.
	//
	// 	- Default value: **0**.
	//
	// example:
	//
	// 0
	SpecifiedOffsetTime *int64 `json:"SpecifiedOffsetTime,omitempty" xml:"SpecifiedOffsetTime,omitempty"`
	// The playback positions at which you want to capture snapshots. Unit: milliseconds. You can specify up to 30 playback positions in a request.
	SpecifiedOffsetTimesShrink *string `json:"SpecifiedOffsetTimes,omitempty" xml:"SpecifiedOffsetTimes,omitempty"`
	// The sprite snapshot configuration. If you set this parameter, sprite snapshots are generated. For more information, see [SpriteSnapshotConfig](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {\\"CellWidth\\": 120, \\"CellHeight\\": 68, \\"Columns\\": 3,\\"Lines\\": 10, \\"Padding\\": 20, \\"Margin\\": 50}
	SpriteSnapshotConfig *string `json:"SpriteSnapshotConfig,omitempty" xml:"SpriteSnapshotConfig,omitempty"`
	// The custom configurations including the configuration of transparent data transmission and callback configurations. The value must be a JSON string. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// >  To use the message callback feature, you must specify an HTTP callback URL and the callback events in the ApsaraVideo VOD console. Otherwise, the callback settings do not take effect.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://.example.aliyundoc.com"},"Extend":{"localId":"xxx","example":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- After you upload a video in the ApsaraVideo VOD console, you can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the ID of the video.
	//
	// 	- Obtain the video ID from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to obtain the upload URL and credential.
	//
	// 	- Obtain the video ID from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you called to query media information after the audio or video file is uploaded.
	//
	// example:
	//
	// d3e680e618708efbf2cae7cc9312****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// The width of each snapshot. Valid values: `[8,4096]`. By default, the width of the video source is used. Unit: pixels.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitSnapshotJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobShrinkRequest) GetCount() *int64 {
	return s.Count
}

func (s *SubmitSnapshotJobShrinkRequest) GetHeight() *string {
	return s.Height
}

func (s *SubmitSnapshotJobShrinkRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *SubmitSnapshotJobShrinkRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *SubmitSnapshotJobShrinkRequest) GetSnapshotTemplateId() *string {
	return s.SnapshotTemplateId
}

func (s *SubmitSnapshotJobShrinkRequest) GetSpecifiedOffsetTime() *int64 {
	return s.SpecifiedOffsetTime
}

func (s *SubmitSnapshotJobShrinkRequest) GetSpecifiedOffsetTimesShrink() *string {
	return s.SpecifiedOffsetTimesShrink
}

func (s *SubmitSnapshotJobShrinkRequest) GetSpriteSnapshotConfig() *string {
	return s.SpriteSnapshotConfig
}

func (s *SubmitSnapshotJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobShrinkRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitSnapshotJobShrinkRequest) GetWidth() *string {
	return s.Width
}

func (s *SubmitSnapshotJobShrinkRequest) SetCount(v int64) *SubmitSnapshotJobShrinkRequest {
	s.Count = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetHeight(v string) *SubmitSnapshotJobShrinkRequest {
	s.Height = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetInterval(v int64) *SubmitSnapshotJobShrinkRequest {
	s.Interval = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetReferenceId(v string) *SubmitSnapshotJobShrinkRequest {
	s.ReferenceId = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetSnapshotTemplateId(v string) *SubmitSnapshotJobShrinkRequest {
	s.SnapshotTemplateId = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetSpecifiedOffsetTime(v int64) *SubmitSnapshotJobShrinkRequest {
	s.SpecifiedOffsetTime = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetSpecifiedOffsetTimesShrink(v string) *SubmitSnapshotJobShrinkRequest {
	s.SpecifiedOffsetTimesShrink = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetSpriteSnapshotConfig(v string) *SubmitSnapshotJobShrinkRequest {
	s.SpriteSnapshotConfig = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetUserData(v string) *SubmitSnapshotJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetVideoId(v string) *SubmitSnapshotJobShrinkRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) SetWidth(v string) *SubmitSnapshotJobShrinkRequest {
	s.Width = &v
	return s
}

func (s *SubmitSnapshotJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
