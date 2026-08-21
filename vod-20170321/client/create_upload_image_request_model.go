// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateUploadImageRequest
	GetAppId() *string
	SetCateId(v int64) *CreateUploadImageRequest
	GetCateId() *int64
	SetDescription(v string) *CreateUploadImageRequest
	GetDescription() *string
	SetImageExt(v string) *CreateUploadImageRequest
	GetImageExt() *string
	SetImageType(v string) *CreateUploadImageRequest
	GetImageType() *string
	SetOriginalFileName(v string) *CreateUploadImageRequest
	GetOriginalFileName() *string
	SetStorageLocation(v string) *CreateUploadImageRequest
	GetStorageLocation() *string
	SetTags(v string) *CreateUploadImageRequest
	GetTags() *string
	SetTitle(v string) *CreateUploadImageRequest
	GetTitle() *string
	SetUserData(v string) *CreateUploadImageRequest
	GetUserData() *string
}

type CreateUploadImageRequest struct {
	// The application ID. Default value: **app-1000000**. If you have activated the multi-application service, specify the application ID to upload the image to the specified application. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID. You can obtain the category ID by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management*	- > **Categories*	- to view the category ID.
	//
	// - Obtain the value of CateId from the response when you call the [AddCategory](~~AddCategory~~) operation to create a category.
	//
	// - Obtain the value of CateId from the response when you call the [GetCategories](~~GetCategories~~) operation to query categories.
	//
	// example:
	//
	// 100036****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The description of the image.
	//
	// - The description can be up to 1024 characters in length.
	//
	// - The description must be encoded in UTF-8.
	//
	// example:
	//
	// Image upload test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file name extension of the image source file to upload. Valid values:
	//
	// - **png*	- (default)
	//
	// - **jpg**
	//
	// - **jpeg**
	//
	// - **gif**
	//
	// - **heic**
	//
	// - **webp**
	//
	// example:
	//
	// png
	ImageExt *string `json:"ImageExt,omitempty" xml:"ImageExt,omitempty"`
	// The type of the image. Valid values:
	//
	// - **default*	- (default): a common image.
	//
	// - **cover**: a video thumbnail.
	//
	// > The ApsaraVideo VOD console supports viewing and managing only images of the **default*	- type.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The address of the image source file to upload.
	//
	// > The file name extension is optional. If a file name extension is included here and is different from the value specified in `ImageExt`, the value of `ImageExt` takes precedence.
	//
	// example:
	//
	// D:\\picture_01
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The storage address. You can obtain the storage address by using the following method:
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management*	- > **Storage*	- to view the storage address.
	//
	// > If you do not specify this parameter, the image is uploaded to the default storage address. If you specify this parameter, the image is uploaded to the specified storage address.
	//
	// example:
	//
	// outin-****..oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the image. Rules:
	//
	// - Each tag can be up to 32 characters in length.
	//
	// - You can specify up to 16 tags.
	//
	// - Separate multiple tags with commas (,).
	//
	// - The tags must be encoded in UTF-8.
	//
	// example:
	//
	// Test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the image. Rules:
	//
	// - The title can be up to 128 characters in length.
	//
	// - The title must be encoded in UTF-8.
	//
	// example:
	//
	// mytitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings in a JSON string. The settings support message callbacks, upload acceleration, and other configurations. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
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

func (s CreateUploadImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadImageRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadImageRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateUploadImageRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *CreateUploadImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUploadImageRequest) GetImageExt() *string {
	return s.ImageExt
}

func (s *CreateUploadImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *CreateUploadImageRequest) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *CreateUploadImageRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CreateUploadImageRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateUploadImageRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateUploadImageRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadImageRequest) SetAppId(v string) *CreateUploadImageRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadImageRequest) SetCateId(v int64) *CreateUploadImageRequest {
	s.CateId = &v
	return s
}

func (s *CreateUploadImageRequest) SetDescription(v string) *CreateUploadImageRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadImageRequest) SetImageExt(v string) *CreateUploadImageRequest {
	s.ImageExt = &v
	return s
}

func (s *CreateUploadImageRequest) SetImageType(v string) *CreateUploadImageRequest {
	s.ImageType = &v
	return s
}

func (s *CreateUploadImageRequest) SetOriginalFileName(v string) *CreateUploadImageRequest {
	s.OriginalFileName = &v
	return s
}

func (s *CreateUploadImageRequest) SetStorageLocation(v string) *CreateUploadImageRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadImageRequest) SetTags(v string) *CreateUploadImageRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadImageRequest) SetTitle(v string) *CreateUploadImageRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadImageRequest) SetUserData(v string) *CreateUploadImageRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadImageRequest) Validate() error {
	return dara.Validate(s)
}
