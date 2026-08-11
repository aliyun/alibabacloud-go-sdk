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
	// The entity ID. You can call the CreateEntity operation to create an entity and define a custom dynamic metadata structure.
	//
	// example:
	//
	// d67281da3c8743b8823ad12976187***
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metadata of the media file to be uploaded. The value is a JSON string.
	//
	// - The metadata takes effect only when it matches a URL in UploadURLs.
	//
	// - JSON format: [UploadMetadata, UploadMetadata, ...]. The value must be converted to a JSON string.
	//
	// - For more information, see the UploadMetadata table below.
	//
	// example:
	//
	// [{"SourceURL":"https://example.aliyundoc.com/video01.mp4","Title":"urlUploadTest"}]
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// The post-upload processing action when Type is set to video or audio.
	//
	// Valid values of ProcessType: Workflow.
	//
	// example:
	//
	// {"ProcessType": "Workflow","ProcessID":"b72a06c6beeb4dcdb898feef067b1***"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage address.
	//
	// - StorageType: only oss is supported.
	//
	// - StorageLocation: only VOD storage is supported. User-owned OSS storage is not supported.
	//
	// example:
	//
	// {"StorageType":"oss","StorageLocation":"outin-***.oss-cn-shanghai.aliyuncs.com"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// The URL of the media source file.
	//
	// - The URL must include a file name extension. For example, mp4 is the file name extension in `https://****.mp4`.
	//
	//     - If the URL does not include a file name extension, you can specify the FileExtension parameter in `UploadMetadatas`.
	//
	//     - If the URL includes a file name extension and the FileExtension parameter is also specified, the value of FileExtension takes precedence.
	//
	// - URL-encode the URLs. Separate multiple URLs with commas (,). A maximum of 20 URLs are supported.
	//
	// - To prevent upload failures caused by special characters, URL-encode each URL before concatenating them with commas.
	//
	// example:
	//
	// https://diffurl.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks.
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
