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
	// The application ID. The default value is `app-1000000`.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the entity. You can call the `CreateEntity` operation to create an entity and define a custom schema for dynamic metadata.
	//
	// example:
	//
	// 9e177cac2fb44f8b8c67b199fcc7bffd
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The file information, provided as a JSON string containing the following fields:
	//
	// - `Type` (Required): The file type. Valid values: `video`, `image`, `audio`, `text`, and `other`.
	//
	// - `Name` (Required): The filename without the extension.
	//
	// - `Size` (Optional): The file size.
	//
	// - `Ext` (Required): The file extension.
	//
	// example:
	//
	// {\\"Type\\":\\"video\\",\\"Name\\":\\"test\\",\\"Size\\":108078336,\\"Ext\\":\\"mp4\\"}
	FileInfo *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	// The media asset metadata, provided as a JSON string.
	//
	// `Title` (Required):
	//
	// - The title can be up to 128 characters in length.
	//
	// - The title must be UTF-8 encoded.
	//
	// `Description` (Optional):
	//
	// - The description can be up to 1,024 characters in length.
	//
	// - The description must be UTF-8 encoded.
	//
	// `CateId` (Optional): The category ID.
	//
	// `Tags` (Optional): The tags of the media asset, separated by commas.
	//
	// `BusinessType` (Required): The business type. Valid values depend on the `Type` specified in `FileInfo`.
	//
	// - If `Type` is `video`: `opening` or `ending`.
	//
	// - If `Type` is `image`: `default`, `cover`, or `watermark`.
	//
	// - If `Type` is `text`: `subtitles` or `font`.
	//
	// -
	//
	// - If `Type` is `other`: `general`.
	//
	// `CoverURL` (Optional): The URL of the cover image.<br>`DynamicMetaData` (Optional): A string for custom dynamic metadata.<br>
	//
	// example:
	//
	// {\\"Title\\": \\"UploadTest\\", \\"Description\\": \\"UploadImageTest\\", \\"Tags\\": \\"tag1,tag2\\",\\"BusinessType\\":\\"cover\\"}
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// The post-processing configuration for `video` or `audio` uploads.
	//
	// Set `ProcessType` to `Workflow`.
	//
	// > - This parameter specifies an [asynchronous task](https://help.aliyun.com/document_detail/3027141.html), which is queued and runs in the background after you submit the request.
	//
	// example:
	//
	// {\\"ProcessType\\":\\"Workflow\\",\\"ProcessID\\":\\"74ba870f1a4873a3ba238e0bf6fa9***\\"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage configuration, provided as a JSON string.
	//
	// - `StorageType`: Only `oss` is supported.
	//
	// - `StorageLocation`: Only VOD storage is supported. You cannot upload to your own OSS buckets.
	//
	// example:
	//
	// {\\"StorageType\\":\\"oss\\",\\"StorageLocation\\":\\"outin-***.oss-cn-shanghai.aliyuncs.com\\"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// A JSON string for custom settings, such as configuring a message callback.
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
