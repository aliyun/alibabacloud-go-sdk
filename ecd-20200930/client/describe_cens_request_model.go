// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCensRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCensRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCensRequest
	GetRegionId() *string
}

type DescribeCensRequest struct {
	// The page number.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCensRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCensRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCensRequest) SetPageNumber(v int32) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int32) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetRegionId(v string) *DescribeCensRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCensRequest) Validate() error {
	return dara.Validate(s)
}
