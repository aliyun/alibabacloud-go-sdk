// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerIdeImagesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerIdeImagesResponseBody
	GetNextToken() *string
	SetPagingInfo(v *ListServerIdeImagesResponseBodyPagingInfo) *ListServerIdeImagesResponseBody
	GetPagingInfo() *ListServerIdeImagesResponseBodyPagingInfo
	SetRequestId(v string) *ListServerIdeImagesResponseBody
	GetRequestId() *string
}

type ListServerIdeImagesResponseBody struct {
	// The maximum number of records returned in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination information.
	PagingInfo *ListServerIdeImagesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServerIdeImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerIdeImagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeImagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeImagesResponseBody) GetPagingInfo() *ListServerIdeImagesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListServerIdeImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerIdeImagesResponseBody) SetMaxResults(v int32) *ListServerIdeImagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeImagesResponseBody) SetNextToken(v string) *ListServerIdeImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeImagesResponseBody) SetPagingInfo(v *ListServerIdeImagesResponseBodyPagingInfo) *ListServerIdeImagesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListServerIdeImagesResponseBody) SetRequestId(v string) *ListServerIdeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerIdeImagesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServerIdeImagesResponseBodyPagingInfo struct {
	// The list of images available for personal development environments.
	Images []*ListServerIdeImagesResponseBodyPagingInfoImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the filter conditions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerIdeImagesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeImagesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) GetImages() []*ListServerIdeImagesResponseBodyPagingInfoImages {
	return s.Images
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) SetImages(v []*ListServerIdeImagesResponseBodyPagingInfoImages) *ListServerIdeImagesResponseBodyPagingInfo {
	s.Images = v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) SetPageNumber(v int32) *ListServerIdeImagesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) SetPageSize(v int32) *ListServerIdeImagesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) SetTotalCount(v int32) *ListServerIdeImagesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfo) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerIdeImagesResponseBodyPagingInfoImages struct {
	// The image ID used by the instance.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// serveride_notebook
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The image status.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServerIdeImagesResponseBodyPagingInfoImages) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeImagesResponseBodyPagingInfoImages) GoString() string {
	return s.String()
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) GetImageName() *string {
	return s.ImageName
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) GetStatus() *string {
	return s.Status
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) SetImageId(v string) *ListServerIdeImagesResponseBodyPagingInfoImages {
	s.ImageId = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) SetImageName(v string) *ListServerIdeImagesResponseBodyPagingInfoImages {
	s.ImageName = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) SetImageUrl(v string) *ListServerIdeImagesResponseBodyPagingInfoImages {
	s.ImageUrl = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) SetStatus(v string) *ListServerIdeImagesResponseBodyPagingInfoImages {
	s.Status = &v
	return s
}

func (s *ListServerIdeImagesResponseBodyPagingInfoImages) Validate() error {
	return dara.Validate(s)
}
