// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeFileSystemsRequest
	GetFileSystemId() *string
	SetMaxResults(v int32) *DescribeFileSystemsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFileSystemsRequest
	GetNextToken() *string
	SetPolarFsInstanceId(v string) *DescribeFileSystemsRequest
	GetPolarFsInstanceId() *string
}

type DescribeFileSystemsRequest struct {
	// The absolute path of the directory to query. If not specified, the default path is used.
	//
	// example:
	//
	// /test
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of records to return per request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned from the previous query to retrieve the next page.
	//
	// example:
	//
	// token123
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the PolarFS instance to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DescribeFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFileSystemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFileSystemsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetMaxResults(v int32) *DescribeFileSystemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetNextToken(v string) *DescribeFileSystemsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPolarFsInstanceId(v string) *DescribeFileSystemsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
