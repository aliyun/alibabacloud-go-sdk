// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateUploadMediaRequest
	GetAppId() *string
	SetEntityId(v string) *CreateUploadMediaRequest
	GetEntityId() *string
	SetFileInfo(v string) *CreateUploadMediaRequest
	GetFileInfo() *string
	SetMediaMetaData(v string) *CreateUploadMediaRequest
	GetMediaMetaData() *string
	SetPostProcessConfig(v string) *CreateUploadMediaRequest
	GetPostProcessConfig() *string
	SetUploadTargetConfig(v string) *CreateUploadMediaRequest
	GetUploadTargetConfig() *string
	SetUserData(v string) *CreateUploadMediaRequest
	GetUserData() *string
}

type CreateUploadMediaRequest struct {
	// The application ID. Default value: app-1000000.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The entity ID. You can call the CreateEntity operation to create an entity and customize the dynamic metadata structure.
	//
	// example:
	//
	// 9e177cac2fb44f8b8c67b199fcc7bffd
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The file information in JSON format. This parameter contains the following fields:
	//
	// - Type (required): the file type. Valid values: video, image, audio, text, and other.
	//
	// - Name (required): the file name without the file name extension.
	//
	// - Size (optional): the file size.
	//
	// - Ext (required): the file name extension.
	//
	// example:
	//
	// {\\"Type\\":\\"video\\",\\"Name\\":\\"test\\",\\"Size\\":108078336,\\"Ext\\":\\"mp4\\"}
	FileInfo *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	// The metadata of the media asset to upload, in JSON format.
	//
	// Title (required):
	//
	// - The maximum length is 128 characters.
	//
	// - UTF-8 encoded.
	//
	// Description (optional):
	//
	// - The maximum length is 1024 characters.
	//
	// - UTF-8 encoded.
	//
	// CateId (optional): the category ID.
	//
	// Tags (optional): the tags.
	//
	// BusinessType (required): the business type. Valid values:
	//
	// - When Type = video:
	//
	// opening: opening credits. ending: ending credits.
	//
	// - When Type = image:
	//
	//   default: default.
	//
	//   cover: cover image.
	//
	// - When Type = text:
	//
	//   subtitles: subtitles.
	//
	//   font: font.
	//
	// - When Type = material:
	//
	//   watermark: watermark.
	//
	// - general: general-purpose.
	//
	// CoverURL (optional): the cover URL.
	//
	// DynamicMetaData: the dynamic metadata. The value is a string.
	//
	// example:
	//
	// {\\"Title\\": \\"UploadTest\\", \\"Description\\": \\"UploadImageTest\\", \\"Tags\\": \\"tag1,tag2\\",\\"BusinessType\\":\\"cover\\"}
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// Specifies the post-upload processing action when Type = video or audio.
	//
	// ProcessType: set to Workflow.
	//
	// >
	//
	// > - This parameter triggers an [asynchronous task](https://help.aliyun.com/document_detail/3027141.html). After submission, the task is not immediately completed and enters a background queue for asynchronous execution.
	//
	// example:
	//
	// {\\"ProcessType\\":\\"Workflow\\",\\"ProcessID\\":\\"74ba870f1a4873a3ba238e0bf6fa9***\\"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage address.
	//
	// - StorageType: only oss is supported.
	//
	// - StorageLocation: only VOD storage is supported. User-owned OSS storage is not supported.
	//
	// example:
	//
	// {\\"StorageType\\":\\"oss\\",\\"StorageLocation\\":\\"outin-***.oss-cn-shanghai.aliyuncs.com\\"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateUploadMediaRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateUploadMediaRequest) GetFileInfo() *string {
	return s.FileInfo
}

func (s *CreateUploadMediaRequest) GetMediaMetaData() *string {
	return s.MediaMetaData
}

func (s *CreateUploadMediaRequest) GetPostProcessConfig() *string {
	return s.PostProcessConfig
}

func (s *CreateUploadMediaRequest) GetUploadTargetConfig() *string {
	return s.UploadTargetConfig
}

func (s *CreateUploadMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadMediaRequest) SetAppId(v string) *CreateUploadMediaRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadMediaRequest) SetEntityId(v string) *CreateUploadMediaRequest {
	s.EntityId = &v
	return s
}

func (s *CreateUploadMediaRequest) SetFileInfo(v string) *CreateUploadMediaRequest {
	s.FileInfo = &v
	return s
}

func (s *CreateUploadMediaRequest) SetMediaMetaData(v string) *CreateUploadMediaRequest {
	s.MediaMetaData = &v
	return s
}

func (s *CreateUploadMediaRequest) SetPostProcessConfig(v string) *CreateUploadMediaRequest {
	s.PostProcessConfig = &v
	return s
}

func (s *CreateUploadMediaRequest) SetUploadTargetConfig(v string) *CreateUploadMediaRequest {
	s.UploadTargetConfig = &v
	return s
}

func (s *CreateUploadMediaRequest) SetUserData(v string) *CreateUploadMediaRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadMediaRequest) Validate() error {
	return dara.Validate(s)
}
