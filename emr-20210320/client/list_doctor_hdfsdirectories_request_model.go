// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHDFSDirectoriesRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHDFSDirectoriesRequest
	GetDateTime() *string
	SetDirPath(v string) *ListDoctorHDFSDirectoriesRequest
	GetDirPath() *string
	SetMaxResults(v int32) *ListDoctorHDFSDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHDFSDirectoriesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHDFSDirectoriesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHDFSDirectoriesRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHDFSDirectoriesRequest
	GetRegionId() *string
}

type ListDoctorHDFSDirectoriesRequest struct {
	// 集群ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// /tmp/test
	DirPath *string `json:"DirPath,omitempty" xml:"DirPath,omitempty"`
	// 一次获取的最大记录数。取值范围：1~100。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// smallFileCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// 区域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDoctorHDFSDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHDFSDirectoriesRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHDFSDirectoriesRequest) GetDirPath() *string {
	return s.DirPath
}

func (s *ListDoctorHDFSDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHDFSDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHDFSDirectoriesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHDFSDirectoriesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHDFSDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHDFSDirectoriesRequest) SetClusterId(v string) *ListDoctorHDFSDirectoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetDateTime(v string) *ListDoctorHDFSDirectoriesRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetDirPath(v string) *ListDoctorHDFSDirectoriesRequest {
	s.DirPath = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetMaxResults(v int32) *ListDoctorHDFSDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetNextToken(v string) *ListDoctorHDFSDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetOrderBy(v string) *ListDoctorHDFSDirectoriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetOrderType(v string) *ListDoctorHDFSDirectoriesRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetRegionId(v string) *ListDoctorHDFSDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
