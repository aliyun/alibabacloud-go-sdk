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
	// The ID of the application. Default value: **app-1000000**. If you have activated the multi-application service, specify the ID of the application to add the watermark template in the specified application. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the auxiliary media asset. Valid values:
	//
	// 	- **watermark**
	//
	// 	- **subtitle**
	//
	// 	- **material**
	//
	// This parameter is required.
	//
	// example:
	//
	// watermark
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The ID of the category. Separate multiple IDs with commas (,). You can specify up to five IDs. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Categories*	- to view the category ID of the media file.
	//
	// 	- Obtain the category ID from the response to the [AddCategory](~~AddCategory~~) operation that you call to create a category.
	//
	// 	- Obtain the category ID from the response to the [GetCategories](~~GetCategories~~) operation that you call to query categories.
	//
	// example:
	//
	// 1298****,0813****
	CateIds *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	// The description of the auxiliary media asset. Take note of the following items:
	//
	// 	- The description can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// uploadTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source file URL of the auxiliary media asset.
	//
	// >  The file name extension is optional. If the file name extension that you specified for this parameter is different from the value of MediaExt, the value of MediaExt takes effect.
	//
	// example:
	//
	// D:\\test.png
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the auxiliary media asset. Unit: byte.
	//
	// example:
	//
	// 123
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file name extension of the auxiliary media asset.
	//
	// 	- Valid values for watermarks: **png, gif, apng, and mov**
	//
	// 	- Valid values for subtitles: **srt, ass, stl, ttml, and vtt**
	//
	// 	- Valid values for materials: **jpg, gif, png, mp4, mat, zip, and apk**
	//
	// example:
	//
	// png
	MediaExt *string `json:"MediaExt,omitempty" xml:"MediaExt,omitempty"`
	// The storage address. Perform the following operations to obtain the storage address:
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Storage**. On the Storage page, view the storage address.
	//
	// >  If you leave this parameter empty, the auxiliary media asset is uploaded to the default storage address. If you specify this parameter, the auxiliary media asset is uploaded to the specified storage address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The one or more tags of the auxiliary media asset. Take note of the following items:
	//
	// 	- You can specify a maximum of 16 tags.
	//
	// 	- If you need to specify multiple tags, separate the tags with commas (,).
	//
	// 	- Each tag can be up to 32 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the auxiliary media asset. The following rules apply:
	//
	// 	- The title cannot exceed 128 bytes.
	//
	// 	- The title must be encoded in UTF-8.
	//
	// example:
	//
	// testTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom configurations. For example, you can specify callback configurations and upload acceleration configurations. The value must be a JSON string. For more information, see [Request parameters](~~86952#section-6fg-qll-v3w~~).
	//
	// > 	- The callback configurations take effect only after you specify the HTTP callback URL and select the specific callback events in the ApsaraVideo VOD console. For more information about how to configure HTTP callback settings in the ApsaraVideo VOD console, see [Configure callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > 	- If you want to enable the upload acceleration feature, submit a ticket. For more information, see [Overview](https://help.aliyun.com/document_detail/55396.html). For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
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
