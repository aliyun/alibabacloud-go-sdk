// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSUGIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHDFSUGIRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHDFSUGIRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorHDFSUGIRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHDFSUGIRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHDFSUGIRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHDFSUGIRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHDFSUGIRequest
	GetRegionId() *string
	SetType(v string) *ListDoctorHDFSUGIRequest
	GetType() *string
}

type ListDoctorHDFSUGIRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The basis on which you want to sort the query results. Valid values:
	//
	// 	- totalFileCount: the total number of files
	//
	// 	- totalDataSize: the total data size
	//
	// 	- totalDirCount: the total number of directories
	//
	// example:
	//
	// totalFileCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- ASC: in ascending order
	//
	// 	- DESC: in descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter condition. Valid values:
	//
	// 	- user
	//
	// 	- group
	//
	// This parameter is required.
	//
	// example:
	//
	// group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDoctorHDFSUGIRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHDFSUGIRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHDFSUGIRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHDFSUGIRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHDFSUGIRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHDFSUGIRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHDFSUGIRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHDFSUGIRequest) GetType() *string {
	return s.Type
}

func (s *ListDoctorHDFSUGIRequest) SetClusterId(v string) *ListDoctorHDFSUGIRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetDateTime(v string) *ListDoctorHDFSUGIRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetMaxResults(v int32) *ListDoctorHDFSUGIRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetNextToken(v string) *ListDoctorHDFSUGIRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetOrderBy(v string) *ListDoctorHDFSUGIRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetOrderType(v string) *ListDoctorHDFSUGIRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetRegionId(v string) *ListDoctorHDFSUGIRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) SetType(v string) *ListDoctorHDFSUGIRequest {
	s.Type = &v
	return s
}

func (s *ListDoctorHDFSUGIRequest) Validate() error {
	return dara.Validate(s)
}
