// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineItemListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineClassKey(v string) *DescribeImageBaselineItemListRequest
	GetBaselineClassKey() *string
	SetBaselineNameKey(v string) *DescribeImageBaselineItemListRequest
	GetBaselineNameKey() *string
	SetCurrentPage(v int32) *DescribeImageBaselineItemListRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageBaselineItemListRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageBaselineItemListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageBaselineItemListRequest
	GetPageSize() *int32
	SetScanRange(v []*string) *DescribeImageBaselineItemListRequest
	GetScanRange() []*string
	SetStatus(v string) *DescribeImageBaselineItemListRequest
	GetStatus() *string
	SetUuids(v []*string) *DescribeImageBaselineItemListRequest
	GetUuids() []*string
}

type DescribeImageBaselineItemListRequest struct {
	// The key of the baseline type.
	//
	// example:
	//
	// ak_leak
	BaselineClassKey *string `json:"BaselineClassKey,omitempty" xml:"BaselineClassKey,omitempty"`
	// The key of the baseline name.
	//
	// example:
	//
	// AccessKey pair leak
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b5c7193300da2070220038718ad****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The types of the assets that are scanned.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The status of the baseline risk item. Valid values:
	//
	// 	- **0**: unfixed
	//
	// 	- **1**: fixed
	//
	// 	- **2**: pending verification
	//
	// 	- **3**: fixing failed
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUIDs of images.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DescribeImageBaselineItemListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineItemListRequest) GetBaselineClassKey() *string {
	return s.BaselineClassKey
}

func (s *DescribeImageBaselineItemListRequest) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *DescribeImageBaselineItemListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineItemListRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageBaselineItemListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBaselineItemListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineItemListRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageBaselineItemListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageBaselineItemListRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *DescribeImageBaselineItemListRequest) SetBaselineClassKey(v string) *DescribeImageBaselineItemListRequest {
	s.BaselineClassKey = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetBaselineNameKey(v string) *DescribeImageBaselineItemListRequest {
	s.BaselineNameKey = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetCurrentPage(v int32) *DescribeImageBaselineItemListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetImageUuid(v string) *DescribeImageBaselineItemListRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetLang(v string) *DescribeImageBaselineItemListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetPageSize(v int32) *DescribeImageBaselineItemListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetScanRange(v []*string) *DescribeImageBaselineItemListRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetStatus(v string) *DescribeImageBaselineItemListRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageBaselineItemListRequest) SetUuids(v []*string) *DescribeImageBaselineItemListRequest {
	s.Uuids = v
	return s
}

func (s *DescribeImageBaselineItemListRequest) Validate() error {
	return dara.Validate(s)
}
