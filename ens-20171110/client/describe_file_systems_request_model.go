// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeFileSystemsRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *DescribeFileSystemsRequest
	GetFileSystemId() *string
	SetFileSystemName(v string) *DescribeFileSystemsRequest
	GetFileSystemName() *string
	SetPageNumber(v int32) *DescribeFileSystemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemsRequest
	GetPageSize() *int32
}

type DescribeFileSystemsRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// c50f8*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the file system.
	//
	// example:
	//
	// FileSystem1
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeFileSystemsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsRequest) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *DescribeFileSystemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemsRequest) SetEnsRegionId(v string) *DescribeFileSystemsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemName(v string) *DescribeFileSystemsRequest {
	s.FileSystemName = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageNumber(v int32) *DescribeFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageSize(v int32) *DescribeFileSystemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
