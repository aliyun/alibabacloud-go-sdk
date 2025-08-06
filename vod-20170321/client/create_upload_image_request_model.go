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
	// The ID of the application. Default value: **app-1000000**. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID of the image. You can use one of the following methods to obtain the category ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Categories**. On the Categories page, you can view the category ID of the image.
	//
	// 	- Obtain the value of CateId from the response to the [AddCategory](https://help.aliyun.com/document_detail/56401.html) operation.
	//
	// 	- Obtain the value of CateId from the response to the [GetCategories](https://help.aliyun.com/document_detail/56406.html) operation.
	//
	// example:
	//
	// 100036****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The description of the image.
	//
	// 	- The description can be up to 1,024 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// The description of the image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file name extension of the image. Valid values:
	//
	// 	- **png*	- (default)
	//
	// 	- **jpg**
	//
	// 	- **jpeg**
	//
	// 	- **gif**
	//
	// example:
	//
	// png
	ImageExt *string `json:"ImageExt,omitempty" xml:"ImageExt,omitempty"`
	// The type of the image. Valid values:
	//
	// 	- **default**: the default image type.
	//
	// 	- **cover**: the thumbnail.
	//
	// > You can manage only images of the **default*	- type in the ApsaraVideo VOD console.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The name of the source file.
	//
	// > The name must contain a file name extension. The file name extension is not case-sensitive.
	//
	// example:
	//
	// D:\\picture_01.png
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The storage address. Perform the following operations to obtain the storage address: Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Storage**. On the Storage page, view the storage address.
	//
	// > If you specify a storage address, media files are uploaded to the specified address.
	//
	// example:
	//
	// outin-****..oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the image. The following rules apply:
	//
	// 	- Each tag can be up to 32 characters in length.
	//
	// 	- You can specify a maximum of 16 tags for an image.
	//
	// 	- Separate multiple tags with commas (,).
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the image. The following rules apply:
	//
	// 	- The title can be up to 128 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// mytitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom configurations, including callback configurations and upload acceleration configurations. The value must be a JSON string. For more information, see the "UserData: specifies the custom configurations for media upload" section of the [Request parameters](https://help.aliyun.com/document_detail/86952.html) topic.
	//
	// > 	- The callback configurations take effect only after you specify the HTTP callback URL and select specific callback events in the ApsaraVideo VOD console. For more information about how to configure HTTP callback settings in the ApsaraVideo VOD console, see [Configure callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > 	- If you want to enable the upload acceleration feature, submit a ticket. For more information, see [Overview](https://help.aliyun.com/document_detail/55396.html). For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
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
