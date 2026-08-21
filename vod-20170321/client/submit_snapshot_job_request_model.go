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
	// The snapshot height. Valid values: `[8,4096]`. Default value: the source video height. Unit: px.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The snapshot interval. The value must be **greater than or equal to 0**.
	//
	// - Unit: seconds.
	//
	// - Default value: **1**.
	//
	// - If Interval is set to **0**, snapshots are evenly captured based on the value of Count and the video duration.
	//
	// example:
	//
	// 1
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens, and underscores are supported. Length: 6 to 64 characters. The value must be unique at the user level.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The snapshot template ID.
	//
	// - We recommend that you create a snapshot template first and then pass the snapshot template ID. For more information about how to create a snapshot template, see [Add a snapshot template](https://help.aliyun.com/document_detail/99406.html).
	//
	// - If you pass the snapshot template ID, all request parameters except Action and VideoId are ignored.
	//
	// example:
	//
	// f5b228fe693bf55bd87b789****
	SnapshotTemplateId *string `json:"SnapshotTemplateId,omitempty" xml:"SnapshotTemplateId,omitempty"`
	// The start time for the snapshot.
	//
	// - Unit: milliseconds.
	//
	// - Default value: **0**.
	//
	// example:
	//
	// 0
	SpecifiedOffsetTime *int64 `json:"SpecifiedOffsetTime,omitempty" xml:"SpecifiedOffsetTime,omitempty"`
	// The points in time at which snapshots are captured. Unit: milliseconds. You can specify up to 30 points in time at a time.
	SpecifiedOffsetTimes []*int64 `json:"SpecifiedOffsetTimes,omitempty" xml:"SpecifiedOffsetTimes,omitempty" type:"Repeated"`
	// The sprite configuration. If this parameter is not empty, a sprite is generated. For more information about the parameter structure, see [SpriteSnapshotConfig](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {\\"CellWidth\\": 120, \\"CellHeight\\": 68, \\"Columns\\": 3,\\"Lines\\": 10, \\"Padding\\": 20, \\"Margin\\": 50}
	SpriteSnapshotConfig *string `json:"SpriteSnapshotConfig,omitempty" xml:"SpriteSnapshotConfig,omitempty"`
	// The custom settings. Only JSON strings are supported. You can use this parameter to pass through custom data and specify callback URL settings. For more information about the parameter structure, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// > To use the message callback in this parameter, configure the HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://.example.aliyundoc.com"},"Extend":{"localId":"xxx","example":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - Obtain the video ID from the value of the VideoId response parameter when you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential.
	//
	// - After the video is uploaded, call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// d3e680e618708efbf2cae7cc9312****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// The snapshot width. Valid values: `[8,4096]`. Default value: the source video width. Unit: px.
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
