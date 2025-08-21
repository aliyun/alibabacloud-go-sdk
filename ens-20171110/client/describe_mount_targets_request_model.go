// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeMountTargetsRequest
	GetEnsRegionId() *string
	SetFileSystemId(v string) *DescribeMountTargetsRequest
	GetFileSystemId() *string
	SetMountTargetName(v string) *DescribeMountTargetsRequest
	GetMountTargetName() *string
	SetPageNumber(v int32) *DescribeMountTargetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountTargetsRequest
	GetPageSize() *int32
}

type DescribeMountTargetsRequest struct {
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// c50f8*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the mount target.
	//
	// example:
	//
	// TestMountPath
	MountTargetName *string `json:"MountTargetName,omitempty" xml:"MountTargetName,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMountTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeMountTargetsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeMountTargetsRequest) GetMountTargetName() *string {
	return s.MountTargetName
}

func (s *DescribeMountTargetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountTargetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountTargetsRequest) SetEnsRegionId(v string) *DescribeMountTargetsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetFileSystemId(v string) *DescribeMountTargetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetMountTargetName(v string) *DescribeMountTargetsRequest {
	s.MountTargetName = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageNumber(v int32) *DescribeMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageSize(v int32) *DescribeMountTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsRequest) Validate() error {
	return dara.Validate(s)
}
