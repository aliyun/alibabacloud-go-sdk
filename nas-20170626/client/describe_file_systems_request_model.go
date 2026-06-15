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
	SetFileSystemType(v string) *DescribeFileSystemsRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeFileSystemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeFileSystemsRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest
	GetTag() []*DescribeFileSystemsRequestTag
	SetVpcId(v string) *DescribeFileSystemsRequest
	GetVpcId() *string
}

type DescribeFileSystemsRequest struct {
	// The ID of the file system.
	//
	// - For General-purpose NAS file systems, the ID is a string of characters, such as `31a8e4****`.
	//
	// - For Extreme NAS file systems, the ID must start with `extreme-`, such as `extreme-0015****`.
	//
	// - For Cloud Parallel File System (CPFS) file systems, the ID must start with `cpfs-`, such as `cpfs-125487****`.
	//
	// - For Cloud Parallel File System SE (CPFS SE) file systems, the ID must start with `cpfsse-`, such as `cpfsse-022c71b134****`.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The file system type.
	//
	// Valid values:
	//
	// - `all` (default): all file system types.
	//
	// - `standard`: General-purpose NAS.
	//
	// - `extreme`: Extreme NAS.
	//
	// - `cpfs`: Cloud Parallel File System (CPFS).
	//
	// - `cpfsse`: Cloud Parallel File System SE (CPFS SE).
	//
	// > Separate multiple types with commas.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number of the file system list.
	//
	// The page number starts at 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of file systems to return on each page.
	//
	// Value range: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// You can view this ID in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?).
	//
	// example:
	//
	// rg-acfmwavnfef****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags used to filter file systems. You can specify 1 to 20 tags.
	Tag []*DescribeFileSystemsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// The file system and the ECS instance used for mounting must be in the same VPC.
	//
	// example:
	//
	// vpc-bp1sevsgtqvk5gxbl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeFileSystemsRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeFileSystemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFileSystemsRequest) GetTag() []*DescribeFileSystemsRequestTag {
	return s.Tag
}

func (s *DescribeFileSystemsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemType(v string) *DescribeFileSystemsRequest {
	s.FileSystemType = &v
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

func (s *DescribeFileSystemsRequest) SetResourceGroupId(v string) *DescribeFileSystemsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsRequest) SetVpcId(v string) *DescribeFileSystemsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// -
	//
	// - The tag key can be up to 128 characters long.
	//
	// - It cannot start with `aliyun` or `acs:`.
	//
	// - It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// -
	//
	// - The tag value can be up to 128 characters long.
	//
	// - It cannot start with `aliyun` or `acs:`.
	//
	// - It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFileSystemsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFileSystemsRequestTag) SetKey(v string) *DescribeFileSystemsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) SetValue(v string) *DescribeFileSystemsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) Validate() error {
	return dara.Validate(s)
}
