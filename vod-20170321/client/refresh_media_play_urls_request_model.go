// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshMediaPlayUrlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinitions(v string) *RefreshMediaPlayUrlsRequest
	GetDefinitions() *string
	SetFormats(v string) *RefreshMediaPlayUrlsRequest
	GetFormats() *string
	SetMediaIds(v string) *RefreshMediaPlayUrlsRequest
	GetMediaIds() *string
	SetResultType(v string) *RefreshMediaPlayUrlsRequest
	GetResultType() *string
	SetSliceCount(v int32) *RefreshMediaPlayUrlsRequest
	GetSliceCount() *int32
	SetSliceFlag(v bool) *RefreshMediaPlayUrlsRequest
	GetSliceFlag() *bool
	SetStreamType(v string) *RefreshMediaPlayUrlsRequest
	GetStreamType() *string
	SetTaskType(v string) *RefreshMediaPlayUrlsRequest
	GetTaskType() *string
	SetUserData(v string) *RefreshMediaPlayUrlsRequest
	GetUserData() *string
}

type RefreshMediaPlayUrlsRequest struct {
	// Specifies the definitions of the streams that you want to purge or prefetch. You can specify multiple definitions. Separate multiple definitions with commas (,). If you do not specify this parameter, **streams in all definitions are purged or prefetched by default**.
	//
	// > The value must be one of the values defined in **Definition*	- in [Metric description for media assets](https://help.aliyun.com/document_detail/124671.html).
	//
	// example:
	//
	// HD, SD
	Definitions *string `json:"Definitions,omitempty" xml:"Definitions,omitempty"`
	// The streaming formats that you want to refresh or prefetch. You can specify multiple formats. Separate multiple formats with commas (,). If you do not specify this parameter, **streams in all formats are refreshed or prefetched by default**. Valid values:
	//
	// - **mp4**
	//
	// - **m3u8**
	//
	// - **mp3**
	//
	// - **flv**
	//
	// - **webm**
	//
	// - **ts**
	//
	// example:
	//
	// mp4,m3u8
	Formats *string `json:"Formats,omitempty" xml:"Formats,omitempty"`
	// The IDs of the audio or video files that you want to refresh or prefetch. You can specify one or more IDs. Separate multiple IDs with commas (,). You can specify up to 20 IDs.
	//
	// You can obtain audio or video IDs by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****, a6e49sfgd23p5g9ja7095863****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The result type of the refresh or prefetch task. Valid values:
	//
	// - **Single*	- (default): Only the latest transcoded stream for each definition and format is refreshed or prefetched.
	//
	// - **Multiple**: All transcoded streams for each definition and format are refreshed or prefetched.
	//
	// example:
	//
	// Single
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The number of TS file playback URLs to refresh or prefetch for M3U8 streams. Only the first N TS file playback URLs of each M3U8 stream are refreshed or prefetched. Valid values: 1 to 20. **Default value: 5**.
	//
	// example:
	//
	// 5
	SliceCount *int32 `json:"SliceCount,omitempty" xml:"SliceCount,omitempty"`
	// Specifies whether to refresh or prefetch the playback URLs of TS files in M3U8 streams. Valid values:
	//
	// - **false*	- (default): No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// false
	SliceFlag *bool `json:"SliceFlag,omitempty" xml:"SliceFlag,omitempty"`
	// The types of the streams that you want to refresh or prefetch. You can specify multiple stream types. Separate multiple stream types with commas (,). If you do not specify this parameter, **all stream types are refreshed or prefetched by default**. Valid values:
	//
	// - **video**: video.
	//
	// - **audio**: audio.
	//
	// example:
	//
	// video
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The type of the task. Valid values:
	//
	// - **Refresh**: purge.
	//
	// - **Preload**: prefetch.
	//
	// This parameter is required.
	//
	// example:
	//
	// Preload
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks and upload acceleration. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// > - To use message callbacks in this parameter, configure an HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. For information about how to configure HTTP callbacks in the console, see [Callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > - To use the upload acceleration feature, submit a ticket to activate it. For more information, see [Upload instructions](https://help.aliyun.com/document_detail/55396.html). For information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"}, "Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s RefreshMediaPlayUrlsRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshMediaPlayUrlsRequest) GoString() string {
	return s.String()
}

func (s *RefreshMediaPlayUrlsRequest) GetDefinitions() *string {
	return s.Definitions
}

func (s *RefreshMediaPlayUrlsRequest) GetFormats() *string {
	return s.Formats
}

func (s *RefreshMediaPlayUrlsRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *RefreshMediaPlayUrlsRequest) GetResultType() *string {
	return s.ResultType
}

func (s *RefreshMediaPlayUrlsRequest) GetSliceCount() *int32 {
	return s.SliceCount
}

func (s *RefreshMediaPlayUrlsRequest) GetSliceFlag() *bool {
	return s.SliceFlag
}

func (s *RefreshMediaPlayUrlsRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *RefreshMediaPlayUrlsRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *RefreshMediaPlayUrlsRequest) GetUserData() *string {
	return s.UserData
}

func (s *RefreshMediaPlayUrlsRequest) SetDefinitions(v string) *RefreshMediaPlayUrlsRequest {
	s.Definitions = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetFormats(v string) *RefreshMediaPlayUrlsRequest {
	s.Formats = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetMediaIds(v string) *RefreshMediaPlayUrlsRequest {
	s.MediaIds = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetResultType(v string) *RefreshMediaPlayUrlsRequest {
	s.ResultType = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetSliceCount(v int32) *RefreshMediaPlayUrlsRequest {
	s.SliceCount = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetSliceFlag(v bool) *RefreshMediaPlayUrlsRequest {
	s.SliceFlag = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetStreamType(v string) *RefreshMediaPlayUrlsRequest {
	s.StreamType = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetTaskType(v string) *RefreshMediaPlayUrlsRequest {
	s.TaskType = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) SetUserData(v string) *RefreshMediaPlayUrlsRequest {
	s.UserData = &v
	return s
}

func (s *RefreshMediaPlayUrlsRequest) Validate() error {
	return dara.Validate(s)
}
