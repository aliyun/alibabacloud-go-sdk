// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketId(v string) *DescribeFileSystemsResponseBody
	GetBucketId() *string
	SetFileCount(v int32) *DescribeFileSystemsResponseBody
	GetFileCount() *int32
	SetFileSystemId(v string) *DescribeFileSystemsResponseBody
	GetFileSystemId() *string
	SetFileSystems(v []*DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody
	GetFileSystems() []*DescribeFileSystemsResponseBodyFileSystems
	SetIsTruncated(v bool) *DescribeFileSystemsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *DescribeFileSystemsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFileSystemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFileSystemsResponseBody
	GetRequestId() *string
}

type DescribeFileSystemsResponseBody struct {
	// The ID of the bucket.
	//
	// example:
	//
	// bucket-456
	BucketId *string `json:"BucketId,omitempty" xml:"BucketId,omitempty"`
	// The total number of files that meet the specified conditions.
	//
	// example:
	//
	// 100
	FileCount *int32 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// fs-123
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The list of files.
	FileSystems []*DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	// Indicates whether the results are truncated. If the value is true, you can use NextToken to query subsequent results.
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of records returned for this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// nextToken456
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) GetBucketId() *string {
	return s.BucketId
}

func (s *DescribeFileSystemsResponseBody) GetFileCount() *int32 {
	return s.FileCount
}

func (s *DescribeFileSystemsResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsResponseBody) GetFileSystems() []*DescribeFileSystemsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *DescribeFileSystemsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *DescribeFileSystemsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFileSystemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileSystemsResponseBody) SetBucketId(v string) *DescribeFileSystemsResponseBody {
	s.BucketId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetFileCount(v int32) *DescribeFileSystemsResponseBody {
	s.FileCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetFileSystemId(v string) *DescribeFileSystemsResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v []*DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetIsTruncated(v bool) *DescribeFileSystemsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetMaxResults(v int32) *DescribeFileSystemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetNextToken(v string) *DescribeFileSystemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) Validate() error {
	if s.FileSystems != nil {
		for _, item := range s.FileSystems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystems struct {
	// The capacity of the file.
	//
	// example:
	//
	// 100GB
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// test.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the file was last modified. The time is in the RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-03-25T09:37:10Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The storage type of the file.
	//
	// example:
	//
	// SSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetFileName() *string {
	return s.FileName
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetLastModified() *string {
	return s.LastModified
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetCapacity(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileName(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetLastModified(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.LastModified = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) Validate() error {
	return dara.Validate(s)
}
