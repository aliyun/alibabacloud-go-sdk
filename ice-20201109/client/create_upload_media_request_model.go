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
	// The entity ID. You can call the CreateEntity operation to create an entity and specify a dynamic metadata structure.
	//
	// example:
	//
	// 9e177cac2fb44f8b8c67b199fcc7bffd
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The file information, which is in the JSON format and contains the following fields:
	//
	// 	- Type: required. The file type. Valid values: video, image, audio, text, and other.
	//
	// 	- Name: required. The file name without the extension.
	//
	// 	- Size: optional. The file size.
	//
	// 	- Ext: required. The file name extension.
	//
	// example:
	//
	// {\\"Type\\":\\"video\\",\\"Name\\":\\"test.mp4\\",\\"Size\\":108078336,\\"Ext\\":\\"mp4\\"}
	FileInfo *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	// The metadata of the media asset, which is a JSON string that contains the following fields:
	//
	// Title: required.
	//
	// 	- The value can be up to 128 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// Description: optional.
	//
	// 	- The value can be up to 1,024 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// CateId: optional.
	//
	// Tags: optional.
	//
	// BusinessType: required. Valid values:
	//
	// 	- opening or ending if Type is set to video
	//
	// 	- default or cover if Type is set to image
	//
	// 	- subtitles or font if Type is set to text
	//
	// 	- watermark if Type is set to material
	//
	// 	- general CoverURL: optional.
	//
	// DynamicMetaData: The value is a string.
	//
	// example:
	//
	// {\\"Title\\": \\"UploadTest\\", \\"Description\\": \\"UploadImageTest\\", \\"Tags\\": \\"tag1,tag2\\",\\"BusinessType\\":\\"cover\\"}
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// The postprocessing configurations. You can specify this parameter if Type is set to video or audio.
	//
	// Set ProcessType to Workflow.
	//
	// example:
	//
	// {\\"ProcessType\\":\\"Workflow\\",\\"ProcessID\\":\\"74ba870f1a4873a3ba238e0bf6fa9***\\"}
	PostProcessConfig *string `json:"PostProcessConfig,omitempty" xml:"PostProcessConfig,omitempty"`
	// The destination storage address.
	//
	// Set StorageType to oss.
	//
	// Set StorageLocation to an address in ApsaraVideo VOD. You cannot set this field to an OSS URL.
	//
	// example:
	//
	// {\\"StorageType\\":\\"oss\\",\\"StorageLocation\\":\\"outin-***.oss-cn-shanghai.aliyuncs.com\\"}
	UploadTargetConfig *string `json:"UploadTargetConfig,omitempty" xml:"UploadTargetConfig,omitempty"`
	// The user data. The value must be a JSON string. You can configure settings such as message callbacks.
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
