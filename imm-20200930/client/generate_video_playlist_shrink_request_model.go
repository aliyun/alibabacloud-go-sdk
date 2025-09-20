// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoPlaylistShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *GenerateVideoPlaylistShrinkRequest
	GetCredentialConfigShrink() *string
	SetMasterURI(v string) *GenerateVideoPlaylistShrinkRequest
	GetMasterURI() *string
	SetNotificationShrink(v string) *GenerateVideoPlaylistShrinkRequest
	GetNotificationShrink() *string
	SetOverwritePolicy(v string) *GenerateVideoPlaylistShrinkRequest
	GetOverwritePolicy() *string
	SetProjectName(v string) *GenerateVideoPlaylistShrinkRequest
	GetProjectName() *string
	SetSourceDuration(v float32) *GenerateVideoPlaylistShrinkRequest
	GetSourceDuration() *float32
	SetSourceStartTime(v float32) *GenerateVideoPlaylistShrinkRequest
	GetSourceStartTime() *float32
	SetSourceSubtitlesShrink(v string) *GenerateVideoPlaylistShrinkRequest
	GetSourceSubtitlesShrink() *string
	SetSourceURI(v string) *GenerateVideoPlaylistShrinkRequest
	GetSourceURI() *string
	SetTagsShrink(v string) *GenerateVideoPlaylistShrinkRequest
	GetTagsShrink() *string
	SetTargetsShrink(v string) *GenerateVideoPlaylistShrinkRequest
	GetTargetsShrink() *string
	SetUserData(v string) *GenerateVideoPlaylistShrinkRequest
	GetUserData() *string
}

type GenerateVideoPlaylistShrinkRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The OSS path of the master playlist.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the full path of the file that is suffixed with .m3u8.
	//
	// >  If a playlist contains subtitles or multiple outputs, the MasterURI parameter is required and the URI of subtitle files or outputs must be in the directory specified by the MasterURI parameter or its subdirectory.
	//
	// example:
	//
	// oss://bucket/object/master.m3u8
	MasterURI *string `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	// The notification settings. To view details, click Notification. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The overwrite policy when the media playlist exists. Valid values:
	//
	// 	- overwrite (default): overwrites an existing media playlist.
	//
	// 	- skip-existing: skips generation and retains the existing media playlist.
	//
	// example:
	//
	// overwrite
	OverwritePolicy *string `json:"OverwritePolicy,omitempty" xml:"OverwritePolicy,omitempty"`
	// The project name.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The period of time during which the playlist is generated. Unit: seconds.
	//
	// 	- If you set this parameter to 0 (default) or leave this parameter empty, a playlist is generated until the end time of the source video.
	//
	// 	- If you set this parameter to a value greater than 0, a playlist is generated for the specified period of time from the start time that you specify.
	//
	// >  If you set this parameter to a value that exceeds the end time of a source video, use the default value.
	//
	// example:
	//
	// 0
	SourceDuration *float32 `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	// The time when the playlist starts to generate. Unit: seconds.
	//
	// 	- If you set this parameter to 0 (default) or leave this parameter empty, the start time of the source video is used as the time when a playlist starts to generate.
	//
	// 	- If you set this parameter to a value greater than 0, the time when a playlist starts to generate is the specified point in time.
	//
	// >  If you use this parameter together with the **SourceDuration*	- parameter, a playlist can be generated based on the partial content of a source video.
	//
	// example:
	//
	// 0
	SourceStartTime *float32 `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	// The subtitle files. By default, this parameter is left empty. Up to two subtitle files are supported.
	SourceSubtitlesShrink *string `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty"`
	// The OSS path of the video file.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the full path of the file that contains the file name extension.
	//
	// >  Only OSS buckets of the Standard storage class are supported. OSS buckets for which hotlink protection whitelists are configured are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-test/testcases/video.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The [tags](https://help.aliyun.com/document_detail/106678.html) that you want to add to a TS file in OSS. You can use tags to manage the lifecycles of TS files in OSS.
	//
	// example:
	//
	// {"key1": "value1", "key2": "value2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The live transcoding playlists. Up to 6 playlists are supported. Each output corresponds to at most one video media playlist and one or more subtitle media playlists.
	//
	// >  If more than one output is configured, the **MasterURI*	- parameter is required.
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The custom user information, which is returned in asynchronous notifications to help you handle the notifications in the system. The maximum length of a notification is 2048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GenerateVideoPlaylistShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *GenerateVideoPlaylistShrinkRequest) GetMasterURI() *string {
	return s.MasterURI
}

func (s *GenerateVideoPlaylistShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *GenerateVideoPlaylistShrinkRequest) GetOverwritePolicy() *string {
	return s.OverwritePolicy
}

func (s *GenerateVideoPlaylistShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GenerateVideoPlaylistShrinkRequest) GetSourceDuration() *float32 {
	return s.SourceDuration
}

func (s *GenerateVideoPlaylistShrinkRequest) GetSourceStartTime() *float32 {
	return s.SourceStartTime
}

func (s *GenerateVideoPlaylistShrinkRequest) GetSourceSubtitlesShrink() *string {
	return s.SourceSubtitlesShrink
}

func (s *GenerateVideoPlaylistShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *GenerateVideoPlaylistShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *GenerateVideoPlaylistShrinkRequest) GetTargetsShrink() *string {
	return s.TargetsShrink
}

func (s *GenerateVideoPlaylistShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateVideoPlaylistShrinkRequest) SetCredentialConfigShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetMasterURI(v string) *GenerateVideoPlaylistShrinkRequest {
	s.MasterURI = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetNotificationShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetOverwritePolicy(v string) *GenerateVideoPlaylistShrinkRequest {
	s.OverwritePolicy = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetProjectName(v string) *GenerateVideoPlaylistShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceDuration(v float32) *GenerateVideoPlaylistShrinkRequest {
	s.SourceDuration = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceStartTime(v float32) *GenerateVideoPlaylistShrinkRequest {
	s.SourceStartTime = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceSubtitlesShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.SourceSubtitlesShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceURI(v string) *GenerateVideoPlaylistShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetTagsShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetTargetsShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.TargetsShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetUserData(v string) *GenerateVideoPlaylistShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) Validate() error {
	return dara.Validate(s)
}
