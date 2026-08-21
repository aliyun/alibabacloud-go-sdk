// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadAttachedMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateUploadAttachedMediaRequest
	GetAppId() *string
	SetBusinessType(v string) *CreateUploadAttachedMediaRequest
	GetBusinessType() *string
	SetCateIds(v string) *CreateUploadAttachedMediaRequest
	GetCateIds() *string
	SetDescription(v string) *CreateUploadAttachedMediaRequest
	GetDescription() *string
	SetFileName(v string) *CreateUploadAttachedMediaRequest
	GetFileName() *string
	SetFileSize(v string) *CreateUploadAttachedMediaRequest
	GetFileSize() *string
	SetMediaExt(v string) *CreateUploadAttachedMediaRequest
	GetMediaExt() *string
	SetStorageLocation(v string) *CreateUploadAttachedMediaRequest
	GetStorageLocation() *string
	SetTags(v string) *CreateUploadAttachedMediaRequest
	GetTags() *string
	SetTitle(v string) *CreateUploadAttachedMediaRequest
	GetTitle() *string
	SetUserData(v string) *CreateUploadAttachedMediaRequest
	GetUserData() *string
}

type CreateUploadAttachedMediaRequest struct {
	// The application ID. Default value: **app-1000000**. If you have activated the multi-application service, specify the application ID to upload the auxiliary media asset to the specified application. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the auxiliary media asset. Valid values:
	//
	// - **watermark**: watermark.
	//
	// - **subtitle**: subtitle.
	//
	// - **material**: material.
	//
	// This parameter is required.
	//
	// example:
	//
	// watermark
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category IDs. Separate multiple IDs with commas (,). A maximum of 5 IDs are supported. You can obtain category IDs by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management Configuration*	- > **Category Management*	- to view category IDs.
	//
	// - The category ID is returned when you call the [AddCategory](~~AddCategory~~) operation to create a category.
	//
	// - The category ID is returned when you call the [GetCategories](~~GetCategories~~) operation to query categories.
	//
	// example:
	//
	// 1298****,0813****
	CateIds *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	// The description of the media asset. Rules:
	//
	// - The description can be up to 1024 bytes in length.
	//
	// - The description must be encoded in UTF-8.
	//
	// example:
	//
	// uploadTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source file address of the auxiliary media asset to be uploaded.
	//
	// >The file name extension is optional. If a file name extension is specified here and is different from the extension specified in MediaExt, the value of MediaExt takes precedence.
	//
	// example:
	//
	// D:\\test.png
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 123
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file name extension of the auxiliary media asset source file to be uploaded. Valid values:
	//
	// - Watermark: **png, gif, apng, mov**.
	//
	// - Subtitle: **srt, ass, stl, ttml, vtt**.
	//
	// - Material: **jpg, gif, png, mp4, mat, zip, apk**.
	//
	// example:
	//
	// png
	MediaExt *string `json:"MediaExt,omitempty" xml:"MediaExt,omitempty"`
	// The storage address. You can obtain the storage address by using the following method:
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management Configuration*	- > **Storage Management*	- to view the storage address.
	//
	// > If you do not specify this parameter, the auxiliary media asset is uploaded to the default storage address. If you specify this parameter, the auxiliary media asset is uploaded to the specified storage address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags. Rules:
	//
	// - A maximum of 16 tags are supported.
	//
	// - Separate multiple tags with commas (,).
	//
	// - Each tag can be up to 32 characters or Chinese characters in length.
	//
	// - The tags must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the auxiliary media asset. Rules:
	//
	// - The title can be up to 128 bytes in length.
	//
	// - The title must be encoded in UTF-8.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings, which is a JSON string. The settings support message callbacks, upload acceleration, and other configurations. For more information, see [UserData](~~86952#section-6fg-qll-v3w~~).
	//
	// > - To use message callbacks in this parameter, you must configure an HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. For information about how to configure HTTP callbacks in the console, see [Callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > - To use the upload acceleration feature, submit a ticket to activate it. For more information, see [Upload instructions](https://help.aliyun.com/document_detail/55396.html). For information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadAttachedMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadAttachedMediaRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateUploadAttachedMediaRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateUploadAttachedMediaRequest) GetCateIds() *string {
	return s.CateIds
}

func (s *CreateUploadAttachedMediaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUploadAttachedMediaRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadAttachedMediaRequest) GetFileSize() *string {
	return s.FileSize
}

func (s *CreateUploadAttachedMediaRequest) GetMediaExt() *string {
	return s.MediaExt
}

func (s *CreateUploadAttachedMediaRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CreateUploadAttachedMediaRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateUploadAttachedMediaRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateUploadAttachedMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadAttachedMediaRequest) SetAppId(v string) *CreateUploadAttachedMediaRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetBusinessType(v string) *CreateUploadAttachedMediaRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetCateIds(v string) *CreateUploadAttachedMediaRequest {
	s.CateIds = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetDescription(v string) *CreateUploadAttachedMediaRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetFileName(v string) *CreateUploadAttachedMediaRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetFileSize(v string) *CreateUploadAttachedMediaRequest {
	s.FileSize = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetMediaExt(v string) *CreateUploadAttachedMediaRequest {
	s.MediaExt = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetStorageLocation(v string) *CreateUploadAttachedMediaRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetTags(v string) *CreateUploadAttachedMediaRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetTitle(v string) *CreateUploadAttachedMediaRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetUserData(v string) *CreateUploadAttachedMediaRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) Validate() error {
	return dara.Validate(s)
}
