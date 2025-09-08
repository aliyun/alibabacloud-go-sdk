// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *DescribeCloudSiemAssetsRequest
	GetAssetName() *string
	SetAssetType(v string) *DescribeCloudSiemAssetsRequest
	GetAssetType() *string
	SetAssetUuid(v string) *DescribeCloudSiemAssetsRequest
	GetAssetUuid() *string
	SetCurrentPage(v int32) *DescribeCloudSiemAssetsRequest
	GetCurrentPage() *int32
	SetIncidentUuid(v string) *DescribeCloudSiemAssetsRequest
	GetIncidentUuid() *string
	SetPageSize(v int32) *DescribeCloudSiemAssetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudSiemAssetsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCloudSiemAssetsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCloudSiemAssetsRequest
	GetRoleType() *int32
}

type DescribeCloudSiemAssetsRequest struct {
	// example:
	//
	// test123
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// ip
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 123456-2222-3333-5555-3435345****
	AssetUuid *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeCloudSiemAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeCloudSiemAssetsRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeCloudSiemAssetsRequest) GetAssetUuid() *string {
	return s.AssetUuid
}

func (s *DescribeCloudSiemAssetsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudSiemAssetsRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudSiemAssetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudSiemAssetsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCloudSiemAssetsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetName(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetType(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetUuid(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetCurrentPage(v int32) *DescribeCloudSiemAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetPageSize(v int32) *DescribeCloudSiemAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRegionId(v string) *DescribeCloudSiemAssetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRoleFor(v int64) *DescribeCloudSiemAssetsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRoleType(v int32) *DescribeCloudSiemAssetsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) Validate() error {
	return dara.Validate(s)
}
