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
	// Console test operation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Android in Container (AIC) image.
	//
	// example:
	//
	// m-ad0ddaddc2d54adeaa09b7c0f1e****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image type. Valid values:
	//
	// - **public**: public image.
	//
	// - **private**: custom image.
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
	// The end time of the query. Format: "2006-01-02". By default, no time limit is applied.
	//
	// example:
	//
	// 2022-02-02
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The start time of the query. Format: "2006-01-02". By default, no time limit is applied.
	//
	// example:
	//
	// 2022-01-02
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The page number. The value starts from **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging. Maximum value: **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The image status. Valid values:
	//
	// - **verifying**: pending verification.
	//
	// - **disable**: unavailable.
	//
	// - **available**: available.
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
