// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessAssetRequest
	GetCurrentPage() *int32
	SetDiskType(v string) *ListAgentlessAssetRequest
	GetDiskType() *string
	SetInstanceId(v string) *ListAgentlessAssetRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListAgentlessAssetRequest
	GetInstanceName() *string
	SetPageSize(v int32) *ListAgentlessAssetRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListAgentlessAssetRequest
	GetPlatform() *string
	SetScanRegionId(v string) *ListAgentlessAssetRequest
	GetScanRegionId() *string
	SetTargetType(v int32) *ListAgentlessAssetRequest
	GetTargetType() *int32
}

type ListAgentlessAssetRequest struct {
	// The page number in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the cloud disk. Values:
	//
	// - **system**: System disk
	//
	// - **data**: Data disk
	//
	// example:
	//
	// data
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the asset instance.
	//
	// example:
	//
	// s-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance.
	//
	// example:
	//
	// ca_cpm_******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum number of items to return per page in a paginated query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	ScanRegionId *string `json:"ScanRegionId,omitempty" xml:"ScanRegionId,omitempty"`
	// The type of the detection target. Values:
	//
	// - **3**: User snapshot
	//
	// - **4**: User-defined image
	//
	// example:
	//
	// 1
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAgentlessAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessAssetRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessAssetRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessAssetRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *ListAgentlessAssetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentlessAssetRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessAssetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessAssetRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListAgentlessAssetRequest) GetScanRegionId() *string {
	return s.ScanRegionId
}

func (s *ListAgentlessAssetRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessAssetRequest) SetCurrentPage(v int32) *ListAgentlessAssetRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetDiskType(v string) *ListAgentlessAssetRequest {
	s.DiskType = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetInstanceId(v string) *ListAgentlessAssetRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetInstanceName(v string) *ListAgentlessAssetRequest {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetPageSize(v int32) *ListAgentlessAssetRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetPlatform(v string) *ListAgentlessAssetRequest {
	s.Platform = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetScanRegionId(v string) *ListAgentlessAssetRequest {
	s.ScanRegionId = &v
	return s
}

func (s *ListAgentlessAssetRequest) SetTargetType(v int32) *ListAgentlessAssetRequest {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessAssetRequest) Validate() error {
	return dara.Validate(s)
}
