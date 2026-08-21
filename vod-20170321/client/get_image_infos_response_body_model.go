// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageInfo(v []*GetImageInfosResponseBodyImageInfo) *GetImageInfosResponseBody
	GetImageInfo() []*GetImageInfosResponseBodyImageInfo
	SetNonExistImageIds(v []*string) *GetImageInfosResponseBody
	GetNonExistImageIds() []*string
	SetRequestId(v string) *GetImageInfosResponseBody
	GetRequestId() *string
}

type GetImageInfosResponseBody struct {
	// The image information list.
	ImageInfo []*GetImageInfosResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
	// The list of image IDs that do not exist.
	NonExistImageIds []*string `json:"NonExistImageIds,omitempty" xml:"NonExistImageIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageInfosResponseBody) GetImageInfo() []*GetImageInfosResponseBodyImageInfo {
	return s.ImageInfo
}

func (s *GetImageInfosResponseBody) GetNonExistImageIds() []*string {
	return s.NonExistImageIds
}

func (s *GetImageInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageInfosResponseBody) SetImageInfo(v []*GetImageInfosResponseBodyImageInfo) *GetImageInfosResponseBody {
	s.ImageInfo = v
	return s
}

func (s *GetImageInfosResponseBody) SetNonExistImageIds(v []*string) *GetImageInfosResponseBody {
	s.NonExistImageIds = v
	return s
}

func (s *GetImageInfosResponseBody) SetRequestId(v string) *GetImageInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageInfosResponseBody) Validate() error {
	if s.ImageInfo != nil {
		for _, item := range s.ImageInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageInfosResponseBodyImageInfo struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 254766071
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// Test
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The time when the image was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-11-21T02:37:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The image description.
	//
	// example:
	//
	// Test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// bbc65bba53f9*****ed90de118a7849
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image type. Valid values:
	//
	// - **default**: regular image.
	//
	// - **cover**: video thumbnail.
	//
	// example:
	//
	// default
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The mezzanine information of the image.
	Mezzanine *GetImageInfosResponseBodyImageInfoMezzanine `json:"Mezzanine,omitempty" xml:"Mezzanine,omitempty" type:"Struct"`
	// The image status. Valid values:
	//
	// - **Uploading**: The image is being uploaded. This is the initial status.
	//
	// - **Normal**: The image is uploaded.
	//
	// - **UploadFail**: The image failed to be uploaded.
	//
	// example:
	//
	// Uploading
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address of the image file.
	//
	// example:
	//
	// outin-****..oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The image tags. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// tag1,tag2,tag3
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The image title.
	//
	// example:
	//
	// this is a sample
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The image access URL. If a CDN domain name is configured, the CDN URL is returned. Otherwise, the OSS URL is returned.
	//
	// example:
	//
	// http://example.aliyundoc.com/image/default/****.gif?auth_key=****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetImageInfosResponseBodyImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfosResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *GetImageInfosResponseBodyImageInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetImageInfosResponseBodyImageInfo) GetCateId() *int64 {
	return s.CateId
}

func (s *GetImageInfosResponseBodyImageInfo) GetCateName() *string {
	return s.CateName
}

func (s *GetImageInfosResponseBodyImageInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetImageInfosResponseBodyImageInfo) GetDescription() *string {
	return s.Description
}

func (s *GetImageInfosResponseBodyImageInfo) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageInfosResponseBodyImageInfo) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageInfosResponseBodyImageInfo) GetMezzanine() *GetImageInfosResponseBodyImageInfoMezzanine {
	return s.Mezzanine
}

func (s *GetImageInfosResponseBodyImageInfo) GetStatus() *string {
	return s.Status
}

func (s *GetImageInfosResponseBodyImageInfo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetImageInfosResponseBodyImageInfo) GetTags() *string {
	return s.Tags
}

func (s *GetImageInfosResponseBodyImageInfo) GetTitle() *string {
	return s.Title
}

func (s *GetImageInfosResponseBodyImageInfo) GetURL() *string {
	return s.URL
}

func (s *GetImageInfosResponseBodyImageInfo) SetAppId(v string) *GetImageInfosResponseBodyImageInfo {
	s.AppId = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetCateId(v int64) *GetImageInfosResponseBodyImageInfo {
	s.CateId = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetCateName(v string) *GetImageInfosResponseBodyImageInfo {
	s.CateName = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetCreationTime(v string) *GetImageInfosResponseBodyImageInfo {
	s.CreationTime = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetDescription(v string) *GetImageInfosResponseBodyImageInfo {
	s.Description = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetImageId(v string) *GetImageInfosResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetImageType(v string) *GetImageInfosResponseBodyImageInfo {
	s.ImageType = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetMezzanine(v *GetImageInfosResponseBodyImageInfoMezzanine) *GetImageInfosResponseBodyImageInfo {
	s.Mezzanine = v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetStatus(v string) *GetImageInfosResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetStorageLocation(v string) *GetImageInfosResponseBodyImageInfo {
	s.StorageLocation = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetTags(v string) *GetImageInfosResponseBodyImageInfo {
	s.Tags = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetTitle(v string) *GetImageInfosResponseBodyImageInfo {
	s.Title = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) SetURL(v string) *GetImageInfosResponseBodyImageInfo {
	s.URL = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfo) Validate() error {
	if s.Mezzanine != nil {
		if err := s.Mezzanine.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageInfosResponseBodyImageInfoMezzanine struct {
	// The file size. Unit: byte.
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
	// The image height. Unit: pixel.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The address of the uploaded source image file.
	//
	// example:
	//
	// ****.gif
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The image width. Unit: pixel.
	//
	// example:
	//
	// 200
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageInfosResponseBodyImageInfoMezzanine) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfosResponseBodyImageInfoMezzanine) GoString() string {
	return s.String()
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) GetFileSize() *string {
	return s.FileSize
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) GetFileURL() *string {
	return s.FileURL
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) GetHeight() *int32 {
	return s.Height
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) GetWidth() *int32 {
	return s.Width
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) SetFileSize(v string) *GetImageInfosResponseBodyImageInfoMezzanine {
	s.FileSize = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) SetFileURL(v string) *GetImageInfosResponseBodyImageInfoMezzanine {
	s.FileURL = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) SetHeight(v int32) *GetImageInfosResponseBodyImageInfoMezzanine {
	s.Height = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) SetOriginalFileName(v string) *GetImageInfosResponseBodyImageInfoMezzanine {
	s.OriginalFileName = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) SetWidth(v int32) *GetImageInfosResponseBodyImageInfoMezzanine {
	s.Width = &v
	return s
}

func (s *GetImageInfosResponseBodyImageInfoMezzanine) Validate() error {
	return dara.Validate(s)
}
