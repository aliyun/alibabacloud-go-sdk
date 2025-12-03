// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *SubmitSnapshotJobRequest
	GetCount() *int64
	SetHeight(v string) *SubmitSnapshotJobRequest
	GetHeight() *string
	SetInterval(v int64) *SubmitSnapshotJobRequest
	GetInterval() *int64
	SetReferenceId(v string) *SubmitSnapshotJobRequest
	GetReferenceId() *string
	SetSnapshotTemplateId(v string) *SubmitSnapshotJobRequest
	GetSnapshotTemplateId() *string
	SetSpecifiedOffsetTime(v int64) *SubmitSnapshotJobRequest
	GetSpecifiedOffsetTime() *int64
	SetSpecifiedOffsetTimes(v []*int64) *SubmitSnapshotJobRequest
	GetSpecifiedOffsetTimes() []*int64
	SetSpriteSnapshotConfig(v string) *SubmitSnapshotJobRequest
	GetSpriteSnapshotConfig() *string
	SetUserData(v string) *SubmitSnapshotJobRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitSnapshotJobRequest
	GetVideoId() *string
	SetWidth(v string) *SubmitSnapshotJobRequest
	GetWidth() *string
}

type SubmitSnapshotJobRequest struct {
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
	SpecifiedOffsetTimes []*int64 `json:"SpecifiedOffsetTimes,omitempty" xml:"SpecifiedOffsetTimes,omitempty" type:"Repeated"`
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

func (s SubmitSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequest) GetCount() *int64 {
	return s.Count
}

func (s *SubmitSnapshotJobRequest) GetHeight() *string {
	return s.Height
}

func (s *SubmitSnapshotJobRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *SubmitSnapshotJobRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *SubmitSnapshotJobRequest) GetSnapshotTemplateId() *string {
	return s.SnapshotTemplateId
}

func (s *SubmitSnapshotJobRequest) GetSpecifiedOffsetTime() *int64 {
	return s.SpecifiedOffsetTime
}

func (s *SubmitSnapshotJobRequest) GetSpecifiedOffsetTimes() []*int64 {
	return s.SpecifiedOffsetTimes
}

func (s *SubmitSnapshotJobRequest) GetSpriteSnapshotConfig() *string {
	return s.SpriteSnapshotConfig
}

func (s *SubmitSnapshotJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitSnapshotJobRequest) GetWidth() *string {
	return s.Width
}

func (s *SubmitSnapshotJobRequest) SetCount(v int64) *SubmitSnapshotJobRequest {
	s.Count = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetHeight(v string) *SubmitSnapshotJobRequest {
	s.Height = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetInterval(v int64) *SubmitSnapshotJobRequest {
	s.Interval = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetReferenceId(v string) *SubmitSnapshotJobRequest {
	s.ReferenceId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSnapshotTemplateId(v string) *SubmitSnapshotJobRequest {
	s.SnapshotTemplateId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSpecifiedOffsetTime(v int64) *SubmitSnapshotJobRequest {
	s.SpecifiedOffsetTime = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSpecifiedOffsetTimes(v []*int64) *SubmitSnapshotJobRequest {
	s.SpecifiedOffsetTimes = v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSpriteSnapshotConfig(v string) *SubmitSnapshotJobRequest {
	s.SpriteSnapshotConfig = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetUserData(v string) *SubmitSnapshotJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetVideoId(v string) *SubmitSnapshotJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetWidth(v string) *SubmitSnapshotJobRequest {
	s.Width = &v
	return s
}

func (s *SubmitSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
