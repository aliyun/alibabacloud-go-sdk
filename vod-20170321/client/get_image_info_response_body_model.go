// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageInfo(v *GetImageInfoResponseBodyImageInfo) *GetImageInfoResponseBody
	GetImageInfo() *GetImageInfoResponseBodyImageInfo
	SetRequestId(v string) *GetImageInfoResponseBody
	GetRequestId() *string
}

type GetImageInfoResponseBody struct {
	// The information about the image.
	ImageInfo *GetImageInfoResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AB99D4DF-FAFA-49DC-9C548C1E261E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBody) GetImageInfo() *GetImageInfoResponseBodyImageInfo {
	return s.ImageInfo
}

func (s *GetImageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageInfoResponseBody) SetImageInfo(v *GetImageInfoResponseBodyImageInfo) *GetImageInfoResponseBody {
	s.ImageInfo = v
	return s
}

func (s *GetImageInfoResponseBody) SetRequestId(v string) *GetImageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageInfoResponseBodyImageInfo struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category.
	//
	// example:
	//
	// 254766071
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// test name
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The time when the image was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-21T02:37:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// bbc65bba53f9*****ed90de118a7849
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the image. Valid values:
	//
	// 	- **CoverSnapshot**: thumbnail snapshot.
	//
	// 	- **NormalSnapshot**: normal snapshot.
	//
	// 	- **SpriteSnapshot**: sprite snapshot.
	//
	// 	- **SpriteOriginSnapshot**: sprite source snapshot.
	//
	// 	- **All**: images of all the preceding types. Multiple types other than All can return for this parameter. Multiple types are separated by commas (,).
	//
	// example:
	//
	// NormalSnapshot
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The source information about the image.
	Mezzanine *GetImageInfoResponseBodyImageInfoMezzanine `json:"Mezzanine,omitempty" xml:"Mezzanine,omitempty" type:"Struct"`
	// The status of the image. Valid values:
	//
	// 	- **Uploading**: The image is being uploaded. This is the initial status.
	//
	// 	- **Normal**: The image is uploaded.
	//
	// 	- **UploadFail**: The image fails to be uploaded.
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The bucket in which the image is stored.
	//
	// example:
	//
	// outin-****..oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the image. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// tag1,tag2,tag3
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the image.
	//
	// example:
	//
	// this is a sample
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The image URL. If a domain name for CDN is specified, a CDN URL is returned. Otherwise, an OSS URL is returned.
	//
	// example:
	//
	// http://example.aliyundoc.com/image/default/****.gif?auth_key=****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetImageInfoResponseBodyImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfoResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBodyImageInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetImageInfoResponseBodyImageInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetImageInfoResponseBodyImageInfo) GetCateName() *string {
	return s.CateName
}

func (s *GetImageInfoResponseBodyImageInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetImageInfoResponseBodyImageInfo) GetDescription() *string {
	return s.Description
}

func (s *GetImageInfoResponseBodyImageInfo) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageInfoResponseBodyImageInfo) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageInfoResponseBodyImageInfo) GetMezzanine() *GetImageInfoResponseBodyImageInfoMezzanine {
	return s.Mezzanine
}

func (s *GetImageInfoResponseBodyImageInfo) GetStatus() *string {
	return s.Status
}

func (s *GetImageInfoResponseBodyImageInfo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetImageInfoResponseBodyImageInfo) GetTags() *string {
	return s.Tags
}

func (s *GetImageInfoResponseBodyImageInfo) GetTitle() *string {
	return s.Title
}

func (s *GetImageInfoResponseBodyImageInfo) GetURL() *string {
	return s.URL
}

func (s *GetImageInfoResponseBodyImageInfo) SetAppId(v string) *GetImageInfoResponseBodyImageInfo {
	s.AppId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCateId(v int64) *GetImageInfoResponseBodyImageInfo {
	s.CateId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCateName(v string) *GetImageInfoResponseBodyImageInfo {
	s.CateName = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCreationTime(v string) *GetImageInfoResponseBodyImageInfo {
	s.CreationTime = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetDescription(v string) *GetImageInfoResponseBodyImageInfo {
	s.Description = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetImageId(v string) *GetImageInfoResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetImageType(v string) *GetImageInfoResponseBodyImageInfo {
	s.ImageType = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetMezzanine(v *GetImageInfoResponseBodyImageInfoMezzanine) *GetImageInfoResponseBodyImageInfo {
	s.Mezzanine = v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetStatus(v string) *GetImageInfoResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetStorageLocation(v string) *GetImageInfoResponseBodyImageInfo {
	s.StorageLocation = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetTags(v string) *GetImageInfoResponseBodyImageInfo {
	s.Tags = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetTitle(v string) *GetImageInfoResponseBodyImageInfo {
	s.Title = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetURL(v string) *GetImageInfoResponseBodyImageInfo {
	s.URL = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) Validate() error {
	return dara.Validate(s)
}

type GetImageInfoResponseBodyImageInfoMezzanine struct {
	// The size of the image. Unit: bytes.
	//
	// example:
	//
	// 8932
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The OSS URL of the image file.
	//
	// example:
	//
	// https://outin-bfefbb*****163e1c7426.oss-cn-XXXXXXXX.aliyuncs.com/image/default/5E84CD536*****D4DAD.png?Expires=1590982353&OSSAccessKeyId=*****&Signature=ALPET74o*****c%3D
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The height of the image. Unit: pixels.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The URL of the source file.
	//
	// example:
	//
	// ****.gif
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The width of the image. Unit: pixels.
	//
	// example:
	//
	// 200
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageInfoResponseBodyImageInfoMezzanine) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfoResponseBodyImageInfoMezzanine) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) GetFileSize() *string {
	return s.FileSize
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) GetFileURL() *string {
	return s.FileURL
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) GetHeight() *int32 {
	return s.Height
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) GetWidth() *int32 {
	return s.Width
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetFileSize(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.FileSize = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetFileURL(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.FileURL = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetHeight(v int32) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.Height = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetOriginalFileName(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.OriginalFileName = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetWidth(v int32) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.Width = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) Validate() error {
	return dara.Validate(s)
}
