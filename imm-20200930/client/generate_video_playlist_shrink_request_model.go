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
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The OSS URI of the Master Playlist.
	//
	// The OSS URI must be in the format of oss\\://${Bucket}/${Object}. ${Bucket} is the name of the OSS bucket that is in the same region as the current project. ${Object} is the full path of the file with the .m3u8 file name extension.
	//
	// > If the playlist has subtitle inputs or multiple target outputs, MasterURI is required. The subtitle URI or target URI must be in the same directory as or a subdirectory of the directory specified by MasterURI.
	//
	// example:
	//
	// oss://test-bucket/test-object/master.m3u8
	MasterURI *string `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	// The message notification configuration. For more information, click Notification. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The policy to overwrite an existing Media Playlist. Valid values:
	//
	// - overwrite (default): Overwrites the existing Media Playlist.
	//
	// - skip-existing: Skips the generation and retains the existing Media Playlist.
	//
	// example:
	//
	// overwrite
	OverwritePolicy *string `json:"OverwritePolicy,omitempty" xml:"OverwritePolicy,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The duration for which the playlist is generated. Unit: seconds (s). Valid values:
	//
	// - 0 (default) or empty: continues to the end of the source video.
	//
	// - Greater than 0: lasts for the specified duration from the start time.
	//
	// > If the specified duration extends beyond the end of the source video, the default value is used.
	//
	// example:
	//
	// 0
	SourceDuration *float32 `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	// The start time for generating the playlist. Unit: seconds (s). Valid values:
	//
	// - 0 (default) or empty: starts from the beginning of the source video.
	//
	// - Greater than 0: starts from the specified time point in the source video.
	//
	// > You can set this parameter together with the **SourceDuration*	- parameter to generate a playlist for a specific part of the source video.
	//
	// example:
	//
	// 0
	SourceStartTime *float32 `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	// The list of subtitles to add. The default value is empty. You can add up to two subtitles.
	SourceSubtitlesShrink *string `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty"`
	// The OSS URI of the video.
	//
	// The OSS URI must be in the format of oss\\://${Bucket}/${Object}. ${Bucket} is the name of the OSS bucket that is in the same region as the current project. ${Object} is the full path of the file, including the file name extension.
	//
	// > Only OSS Standard storage buckets are supported. Buckets with hotlink protection whitelists are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-source-object/video.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Adds OSS object [tags](https://help.aliyun.com/document_detail/106678.html) to the generated TS files. You can use tags to control the lifecycle of OSS files.
	//
	// example:
	//
	// {"key1": "value1", "key2": "value2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// An array of live transcoding playlists. The maximum array length is 6. Each target corresponds to a maximum of one video Media Playlist and one or more subtitle Media Playlists.
	//
	// > If you configure more than one target, the **MasterURI*	- parameter must not be empty.
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The custom information. This information is returned in the asynchronous notification message to help you associate the message with your services. The maximum length is 2,048 bytes.
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
