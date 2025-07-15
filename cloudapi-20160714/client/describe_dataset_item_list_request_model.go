// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatasetItemListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *DescribeDatasetItemListRequest
	GetDatasetId() *string
	SetDatasetItemIds(v string) *DescribeDatasetItemListRequest
	GetDatasetItemIds() *string
	SetPageNumber(v int32) *DescribeDatasetItemListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatasetItemListRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeDatasetItemListRequest
	GetSecurityToken() *string
}

type DescribeDatasetItemListRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 602e1f6b3543200eaab0a89e********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the data entry. You can enter multiple IDs. Separate them with commas (,).
	//
	// example:
	//
	// 5045****
	DatasetItemIds *string `json:"DatasetItemIds,omitempty" xml:"DatasetItemIds,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDatasetItemListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatasetItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatasetItemListRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeDatasetItemListRequest) GetDatasetItemIds() *string {
	return s.DatasetItemIds
}

func (s *DescribeDatasetItemListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatasetItemListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatasetItemListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDatasetItemListRequest) SetDatasetId(v string) *DescribeDatasetItemListRequest {
	s.DatasetId = &v
	return s
}

func (s *DescribeDatasetItemListRequest) SetDatasetItemIds(v string) *DescribeDatasetItemListRequest {
	s.DatasetItemIds = &v
	return s
}

func (s *DescribeDatasetItemListRequest) SetPageNumber(v int32) *DescribeDatasetItemListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatasetItemListRequest) SetPageSize(v int32) *DescribeDatasetItemListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatasetItemListRequest) SetSecurityToken(v string) *DescribeDatasetItemListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDatasetItemListRequest) Validate() error {
	return dara.Validate(s)
}
