// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaByURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadMediaByURLRequest
	GetAppId() *string
	SetEntityId(v string) *UploadMediaByURLRequest
	GetEntityId() *string
	SetMediaMetaData(v string) *UploadMediaByURLRequest
	GetMediaMetaData() *string
	SetPostProcessConfig(v string) *UploadMediaByURLRequest
	GetPostProcessConfig() *string
	SetUploadTargetConfig(v string) *UploadMediaByURLRequest
	GetUploadTargetConfig() *string
	SetUploadURLs(v string) *UploadMediaByURLRequest
	GetUploadURLs() *string
	SetUserData(v string) *UploadMediaByURLRequest
	GetUserData() *string
}

type UploadMediaByURLRequest struct {
	// The application ID.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The entity ID. You can call the CreateEntity API to create an entity and define a custom dynamic metadata schema.
	//
	// example:
	//
	// d67281da3c8743b8823ad12976187***
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metadata of the media file, provided as a JSON string.
	//
	// - This metadata takes effect only when it matches a URL in `UploadURLs`.
	//
	// - The value must be a JSON array in the `[UploadMetadata, UploadMetadata, ...]` format, passed as a JSON string.
	//
	// - For more information, see the UploadMetadata table below.
	//
	// example:
	//
	// [{"SourceURL":"https://example.aliyundoc.com/video01.mp4","Title":"urlUploadTest"}]
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// Specifies post-upload processing actions for media files of type `video` or `audio`.
	//
	// The only supported value for `ProcessType` is `Workflow`.
	//
	// example:
	//
	// {"ProcessType": "Workflow","ProcessID":"b72a06c6beeb4dcdb898feef067b1***"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage location.
	//
	// - The only valid value for `StorageType` is `oss`.
	//
	// - `StorageLocation` supports VOD storage only and does not support your own OSS buckets.
	//
	// example:
	//
	// {"StorageType":"oss","StorageLocation":"outin-***.oss-cn-shanghai.aliyuncs.com"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// The source URL of the media file.
	//
	// - The URL must include a file extension. For example, in `https://****.mp4`, mp4 is the file extension.
	//
	//   - If the URL does not include a file extension, you can specify it by using the `FileExtension` parameter in `MediaMetaData`.
	//
	//   - If a file extension is present in both the URL and the `FileExtension` parameter, the value of `FileExtension` takes precedence.
	//
	// - The URLs must be URL-encoded. Separate multiple URLs with commas (,). You can specify up to 20 URLs.
	//
	// - To prevent upload failures due to special characters, URL-encode each URL before concatenating them with commas.
	//
	// example:
	//
	// https://diffurl.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	// Custom settings, provided as a JSON string. This parameter supports configurations such as message callbacks.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UploadMediaByURLRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadMediaByURLRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *UploadMediaByURLRequest) GetMediaMetaData() *string {
	return s.MediaMetaData
}

func (s *UploadMediaByURLRequest) GetPostProcessConfig() *string {
	return s.PostProcessConfig
}

func (s *UploadMediaByURLRequest) GetUploadTargetConfig() *string {
	return s.UploadTargetConfig
}

func (s *UploadMediaByURLRequest) GetUploadURLs() *string {
	return s.UploadURLs
}

func (s *UploadMediaByURLRequest) GetUserData() *string {
	return s.UserData
}

func (s *UploadMediaByURLRequest) SetAppId(v string) *UploadMediaByURLRequest {
	s.AppId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetEntityId(v string) *UploadMediaByURLRequest {
	s.EntityId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetMediaMetaData(v string) *UploadMediaByURLRequest {
	s.MediaMetaData = &v
	return s
}

func (s *UploadMediaByURLRequest) SetPostProcessConfig(v string) *UploadMediaByURLRequest {
	s.PostProcessConfig = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadTargetConfig(v string) *UploadMediaByURLRequest {
	s.UploadTargetConfig = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadURLs(v string) *UploadMediaByURLRequest {
	s.UploadURLs = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUserData(v string) *UploadMediaByURLRequest {
	s.UserData = &v
	return s
}

func (s *UploadMediaByURLRequest) Validate() error {
	return dara.Validate(s)
}
