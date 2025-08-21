// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v *DescribeExportImageInfoResponseBodyImages) *DescribeExportImageInfoResponseBody
	GetImages() *DescribeExportImageInfoResponseBodyImages
	SetPageNumber(v int32) *DescribeExportImageInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExportImageInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeExportImageInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeExportImageInfoResponseBody
	GetTotalCount() *int32
}

type DescribeExportImageInfoResponseBody struct {
	// The returned result. For more information, see the Images parameter described in the JSON-formatted sample success response.
	Images *DescribeExportImageInfoResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExportImageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBody) GetImages() *DescribeExportImageInfoResponseBodyImages {
	return s.Images
}

func (s *DescribeExportImageInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExportImageInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExportImageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExportImageInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExportImageInfoResponseBody) SetImages(v *DescribeExportImageInfoResponseBodyImages) *DescribeExportImageInfoResponseBody {
	s.Images = v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetPageNumber(v int32) *DescribeExportImageInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetPageSize(v int32) *DescribeExportImageInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetRequestId(v string) *DescribeExportImageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) SetTotalCount(v int32) *DescribeExportImageInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExportImageInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExportImageInfoResponseBodyImages struct {
	Image []*DescribeExportImageInfoResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeExportImageInfoResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageInfoResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBodyImages) GetImage() []*DescribeExportImageInfoResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeExportImageInfoResponseBodyImages) SetImage(v []*DescribeExportImageInfoResponseBodyImagesImage) *DescribeExportImageInfoResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

type DescribeExportImageInfoResponseBodyImagesImage struct {
	// The architecture of the image. Example: **x86_64**.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2017-12-08T12:10:03Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The URL of the exported image.
	//
	// example:
	//
	// https://user_defined_bucket.oss-cn-beijing.aliyuncs.com/m-xxxxxxxxx1.qcow2
	ExportedImageURL *string `json:"ExportedImageURL,omitempty" xml:"ExportedImageURL,omitempty"`
	// The export status of the image. Valid values:
	//
	// 	- Exporting
	//
	// 	- Exported
	//
	// 	- ExportError
	//
	// 	- Unexported
	//
	// example:
	//
	// Exported
	ImageExportStatus *string `json:"ImageExportStatus,omitempty" xml:"ImageExportStatus,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-xxxxxxxxx1
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// m-xxxxxxxxx1
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The source of the image. Valid values:
	//
	// 	- system: public image
	//
	// 	- self: custom image
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The OS.
	//
	// example:
	//
	// centos
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DescribeExportImageInfoResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageInfoResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetExportedImageURL() *string {
	return s.ExportedImageURL
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetImageExportStatus() *string {
	return s.ImageExportStatus
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetArchitecture(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetCreationTime(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetExportedImageURL(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ExportedImageURL = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageExportStatus(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageExportStatus = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageId(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageName(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) SetPlatform(v string) *DescribeExportImageInfoResponseBodyImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeExportImageInfoResponseBodyImagesImage) Validate() error {
	return dara.Validate(s)
}
