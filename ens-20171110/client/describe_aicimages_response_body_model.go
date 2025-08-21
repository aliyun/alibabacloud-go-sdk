// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*DescribeAICImagesResponseBodyImages) *DescribeAICImagesResponseBody
	GetImages() []*DescribeAICImagesResponseBodyImages
	SetPageNumber(v int32) *DescribeAICImagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAICImagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAICImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAICImagesResponseBody
	GetTotalCount() *int32
}

type DescribeAICImagesResponseBody struct {
	// The information about the images.
	Images []*DescribeAICImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 72DC6C0A-D9A8-5345-A2BE-FE354CC728A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAICImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAICImagesResponseBody) GetImages() []*DescribeAICImagesResponseBodyImages {
	return s.Images
}

func (s *DescribeAICImagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAICImagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAICImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAICImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAICImagesResponseBody) SetImages(v []*DescribeAICImagesResponseBodyImages) *DescribeAICImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeAICImagesResponseBody) SetPageNumber(v int32) *DescribeAICImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAICImagesResponseBody) SetPageSize(v int32) *DescribeAICImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAICImagesResponseBody) SetRequestId(v string) *DescribeAICImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAICImagesResponseBody) SetTotalCount(v int32) *DescribeAICImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAICImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAICImagesResponseBodyImages struct {
	AndroidVersion *string `json:"AndroidVersion,omitempty" xml:"AndroidVersion,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2022-09-22 10:54:34
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// Test operation of console
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID of the AIC instance.
	//
	// example:
	//
	// m-ad0ddaddc2d54adeaa09b7c0f1e****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The URL of the AIC image repository.
	//
	// example:
	//
	// ****.alibaba-inc.com/aic/socimage:test
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The status of the image. **Available*	- is returned for this parameter. Available indicates that the image is available.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The username of the image repository.
	//
	// example:
	//
	// user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAICImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeAICImagesResponseBodyImages) GetAndroidVersion() *string {
	return s.AndroidVersion
}

func (s *DescribeAICImagesResponseBodyImages) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAICImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *DescribeAICImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAICImagesResponseBodyImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeAICImagesResponseBodyImages) GetStatus() *string {
	return s.Status
}

func (s *DescribeAICImagesResponseBodyImages) GetUser() *string {
	return s.User
}

func (s *DescribeAICImagesResponseBodyImages) SetAndroidVersion(v string) *DescribeAICImagesResponseBodyImages {
	s.AndroidVersion = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetCreationTime(v string) *DescribeAICImagesResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetDescription(v string) *DescribeAICImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetImageId(v string) *DescribeAICImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetImageUrl(v string) *DescribeAICImagesResponseBodyImages {
	s.ImageUrl = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetStatus(v string) *DescribeAICImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) SetUser(v string) *DescribeAICImagesResponseBodyImages {
	s.User = &v
	return s
}

func (s *DescribeAICImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
