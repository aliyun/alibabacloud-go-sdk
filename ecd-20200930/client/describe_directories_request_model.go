// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v []*string) *DescribeDirectoriesRequest
	GetDirectoryId() []*string
	SetDirectoryStatus(v string) *DescribeDirectoriesRequest
	GetDirectoryStatus() *string
	SetDirectoryType(v string) *DescribeDirectoriesRequest
	GetDirectoryType() *string
	SetMaxResults(v int32) *DescribeDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDirectoriesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDirectoriesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeDirectoriesRequest
	GetStatus() *string
}

type DescribeDirectoriesRequest struct {
	// The directory IDs. You can specify one or more directory IDs.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// The directory status. This parameter is the same as Status.
	//
	// example:
	//
	// REGISTERED
	DirectoryStatus *string `json:"DirectoryStatus,omitempty" xml:"DirectoryStatus,omitempty"`
	// The directory type.
	//
	// example:
	//
	// SIMPLE
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The number of entries per page in a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. An empty value indicates that no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The directory status.
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) GetDirectoryId() []*string {
	return s.DirectoryId
}

func (s *DescribeDirectoriesRequest) GetDirectoryStatus() *string {
	return s.DirectoryStatus
}

func (s *DescribeDirectoriesRequest) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDirectoriesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryStatus(v string) *DescribeDirectoriesRequest {
	s.DirectoryStatus = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int32) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetStatus(v string) *DescribeDirectoriesRequest {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
