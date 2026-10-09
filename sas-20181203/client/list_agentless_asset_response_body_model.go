// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*ListAgentlessAssetResponseBodyAssetList) *ListAgentlessAssetResponseBody
	GetAssetList() []*ListAgentlessAssetResponseBodyAssetList
	SetPageInfo(v *ListAgentlessAssetResponseBodyPageInfo) *ListAgentlessAssetResponseBody
	GetPageInfo() *ListAgentlessAssetResponseBodyPageInfo
	SetRequestId(v string) *ListAgentlessAssetResponseBody
	GetRequestId() *string
}

type ListAgentlessAssetResponseBody struct {
	// The returned list of assets.
	AssetList []*ListAgentlessAssetResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAgentlessAssetResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates this ID as a unique identifier for the request. You can use this ID to troubleshoot and locate issues.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessAssetResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessAssetResponseBody) GetAssetList() []*ListAgentlessAssetResponseBodyAssetList {
	return s.AssetList
}

func (s *ListAgentlessAssetResponseBody) GetPageInfo() *ListAgentlessAssetResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAgentlessAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessAssetResponseBody) SetAssetList(v []*ListAgentlessAssetResponseBodyAssetList) *ListAgentlessAssetResponseBody {
	s.AssetList = v
	return s
}

func (s *ListAgentlessAssetResponseBody) SetPageInfo(v *ListAgentlessAssetResponseBodyPageInfo) *ListAgentlessAssetResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAgentlessAssetResponseBody) SetRequestId(v string) *ListAgentlessAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessAssetResponseBody) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentlessAssetResponseBodyAssetList struct {
	// The type of the cloud disk. Valid values:
	//
	// - system: system cloud disk.
	//
	// - data: data cloud disk.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// s-rj9gda4wolo0zixi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// TestInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The asset type. Valid values:
	//
	// - **3**: user snapshot
	//
	// - **4**: user-defined image
	//
	// example:
	//
	// 3
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAgentlessAssetResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessAssetResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetDiskType() *string {
	return s.DiskType
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetPlatform() *string {
	return s.Platform
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentlessAssetResponseBodyAssetList) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetDiskType(v string) *ListAgentlessAssetResponseBodyAssetList {
	s.DiskType = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetInstanceId(v string) *ListAgentlessAssetResponseBodyAssetList {
	s.InstanceId = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetInstanceName(v string) *ListAgentlessAssetResponseBodyAssetList {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetPlatform(v string) *ListAgentlessAssetResponseBodyAssetList {
	s.Platform = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetRegionId(v string) *ListAgentlessAssetResponseBodyAssetList {
	s.RegionId = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) SetTargetType(v int32) *ListAgentlessAssetResponseBodyAssetList {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessAssetResponseBodyPageInfo struct {
	// The page number in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 90
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentlessAssetResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessAssetResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAgentlessAssetResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessAssetResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessAssetResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentlessAssetResponseBodyPageInfo) SetCurrentPage(v int32) *ListAgentlessAssetResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyPageInfo) SetPageSize(v int32) *ListAgentlessAssetResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyPageInfo) SetTotalCount(v int32) *ListAgentlessAssetResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAgentlessAssetResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
