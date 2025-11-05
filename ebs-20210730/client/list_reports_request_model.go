// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListReportsRequest
	GetAppId() *string
	SetMaxResults(v int32) *ListReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListReportsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListReportsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListReportsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListReportsRequest
	GetRegionId() *string
}

type ListReportsRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// app-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Maximum number of items for Token-based pagination.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token (Token), the value is the NextToken parameter value returned from the previous API call.
	//
	// example:
	//
	// a6792e832ff0XXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Page number for paginated queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of rows per page when performing paginated queries.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) to query the list of regions supported by Block Storage Data Insights.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReportsRequest) GoString() string {
	return s.String()
}

func (s *ListReportsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReportsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListReportsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReportsRequest) SetAppId(v string) *ListReportsRequest {
	s.AppId = &v
	return s
}

func (s *ListReportsRequest) SetMaxResults(v int32) *ListReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReportsRequest) SetNextToken(v string) *ListReportsRequest {
	s.NextToken = &v
	return s
}

func (s *ListReportsRequest) SetPageNumber(v int32) *ListReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListReportsRequest) SetPageSize(v int32) *ListReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListReportsRequest) SetRegionId(v string) *ListReportsRequest {
	s.RegionId = &v
	return s
}

func (s *ListReportsRequest) Validate() error {
	return dara.Validate(s)
}
