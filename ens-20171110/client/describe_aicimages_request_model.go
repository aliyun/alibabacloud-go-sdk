// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeAICImagesRequest
	GetDescription() *string
	SetImageId(v string) *DescribeAICImagesRequest
	GetImageId() *string
	SetImageType(v string) *DescribeAICImagesRequest
	GetImageType() *string
	SetImageUrl(v string) *DescribeAICImagesRequest
	GetImageUrl() *string
	SetMaxDate(v string) *DescribeAICImagesRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeAICImagesRequest
	GetMinDate() *string
	SetPageNumber(v string) *DescribeAICImagesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAICImagesRequest
	GetPageSize() *string
	SetStatus(v string) *DescribeAICImagesRequest
	GetStatus() *string
}

type DescribeAICImagesRequest struct {
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
	// The type of the image. Valid values:
	//
	// 	- **public**: public image
	//
	// 	- **private**: custom image
	//
	// example:
	//
	// public
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The URL of the AIC image repository.
	//
	// example:
	//
	// ****.alibaba-inc.com/aic/socimage:test
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The end of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-02-02
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-01-02
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- **verifying**
	//
	// 	- **disable**
	//
	// 	- **available**
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAICImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAICImagesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeAICImagesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAICImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeAICImagesRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeAICImagesRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeAICImagesRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeAICImagesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAICImagesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAICImagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAICImagesRequest) SetDescription(v string) *DescribeAICImagesRequest {
	s.Description = &v
	return s
}

func (s *DescribeAICImagesRequest) SetImageId(v string) *DescribeAICImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeAICImagesRequest) SetImageType(v string) *DescribeAICImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeAICImagesRequest) SetImageUrl(v string) *DescribeAICImagesRequest {
	s.ImageUrl = &v
	return s
}

func (s *DescribeAICImagesRequest) SetMaxDate(v string) *DescribeAICImagesRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeAICImagesRequest) SetMinDate(v string) *DescribeAICImagesRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeAICImagesRequest) SetPageNumber(v string) *DescribeAICImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAICImagesRequest) SetPageSize(v string) *DescribeAICImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAICImagesRequest) SetStatus(v string) *DescribeAICImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAICImagesRequest) Validate() error {
	return dara.Validate(s)
}
