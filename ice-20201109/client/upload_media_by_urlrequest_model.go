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
	// The entity ID. You can call the CreateEntity operation to create an entity and specify a dynamic metadata structure.
	//
	// example:
	//
	// d67281da3c8743b8823ad12976187***
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metadata of the media file that you want to upload. The value must be a JSON string.
	//
	// 	- This parameter takes effect only if its value matches a URL that is specified in UploadURLs.
	//
	// 	- You must convert the JSON-formatted data, such as [UploadMetadata, UploadMetadata,…], into a JSON string.
	//
	// 	- For more information, see the "UploadMetadata" section of this topic.
	//
	// example:
	//
	// [{"SourceURL":"https://example.aliyundoc.com/video01.mp4","Title":"urlUploadTest"}]
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// The postprocessing configurations. You can specify this parameter if Type is set to video or audio.
	//
	// Set ProcessType to Workflow.
	//
	// example:
	//
	// {"ProcessType": "Workflow","ProcessID":"b72a06c6beeb4dcdb898feef067b1***"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage address.
	//
	// Set StorageType to oss.
	//
	// Set StorageLocation to an address in ApsaraVideo VOD. You cannot set this field to an OSS URL.
	//
	// example:
	//
	// {"StorageType":"oss","StorageLocation":"outin-***.oss-cn-shanghai.aliyuncs.com"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// The URL of the source file.
	//
	// 	- The URL must contain a file name extension, such as mp4 in `https://****.mp4`.
	//
	//     	- If the URL does not contain a file name extension, you can specify one by setting `FileExtension` in `UploadMetadata`.
	//
	//     	- If the URL contains a file name extension and `FileExtension` is also specified, the value of `FileExtension` prevails.
	//
	// 	- URL encoding is required. Separate multiple URLs with commas (,). You can specify a maximum of 20 URLs.
	//
	// 	- Special characters may cause upload failures. Therefore, you must encode URLs before you separate them with commas (,).
	//
	// example:
	//
	// https://diffurl.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	// The user data. The value must be a JSON string. You can configure settings such as message callbacks.
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
