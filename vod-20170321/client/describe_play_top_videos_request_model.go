// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayTopVideosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v string) *DescribePlayTopVideosRequest
	GetBizDate() *string
	SetOwnerId(v int64) *DescribePlayTopVideosRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *DescribePlayTopVideosRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayTopVideosRequest
	GetPageSize() *int64
}

type DescribePlayTopVideosRequest struct {
	// The date to query. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T13:00:00Z
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **100**. Maximum value: **1000**.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePlayTopVideosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayTopVideosRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosRequest) GetBizDate() *string {
	return s.BizDate
}

func (s *DescribePlayTopVideosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePlayTopVideosRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayTopVideosRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayTopVideosRequest) SetBizDate(v string) *DescribePlayTopVideosRequest {
	s.BizDate = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetOwnerId(v int64) *DescribePlayTopVideosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetPageNo(v int64) *DescribePlayTopVideosRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetPageSize(v int64) *DescribePlayTopVideosRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayTopVideosRequest) Validate() error {
	return dara.Validate(s)
}
