// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallAegisMachinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUninstallAegisMachinesRequest
	GetCurrentPage() *int32
	SetOs(v string) *ListUninstallAegisMachinesRequest
	GetOs() *string
	SetPageSize(v int32) *ListUninstallAegisMachinesRequest
	GetPageSize() *int32
	SetRegionIdStr(v string) *ListUninstallAegisMachinesRequest
	GetRegionIdStr() *string
	SetRegionNo(v string) *ListUninstallAegisMachinesRequest
	GetRegionNo() *string
	SetRemark(v string) *ListUninstallAegisMachinesRequest
	GetRemark() *string
	SetSourceIp(v string) *ListUninstallAegisMachinesRequest
	GetSourceIp() *string
	SetVendor(v int32) *ListUninstallAegisMachinesRequest
	GetVendor() *int32
}

type ListUninstallAegisMachinesRequest struct {
	// The page number of the first page to display in the query results. Default value: **1**, which indicates that the query results are displayed from page 1.
	//
	// example:
	//
	// 4
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The operating system.
	//
	// > You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to obtain supported operating systems from the **Values*	- of the item whose **Name*	- is **osType**.
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The number of entries per page in a paged query. Default value: **5**, which indicates that 5 entries are displayed per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the server resides.
	//
	// > You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to obtain supported regions from the **Values*	- of the item whose **Name*	- is **regionId**.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdStr *string `json:"RegionIdStr,omitempty" xml:"RegionIdStr,omitempty"`
	// The region where the server resides.
	//
	// > You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to obtain supported regions from the **Values*	- of the item whose **Name*	- is **regionId**.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The asset information to query. You can set this parameter to the asset name or public IP address.
	//
	// example:
	//
	// 172.20.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 180.113.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: non-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: third-party cloud asset
	//
	// - **8**: lightweight asset.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListUninstallAegisMachinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallAegisMachinesRequest) GoString() string {
	return s.String()
}

func (s *ListUninstallAegisMachinesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUninstallAegisMachinesRequest) GetOs() *string {
	return s.Os
}

func (s *ListUninstallAegisMachinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUninstallAegisMachinesRequest) GetRegionIdStr() *string {
	return s.RegionIdStr
}

func (s *ListUninstallAegisMachinesRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ListUninstallAegisMachinesRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListUninstallAegisMachinesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListUninstallAegisMachinesRequest) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListUninstallAegisMachinesRequest) SetCurrentPage(v int32) *ListUninstallAegisMachinesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetOs(v string) *ListUninstallAegisMachinesRequest {
	s.Os = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetPageSize(v int32) *ListUninstallAegisMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetRegionIdStr(v string) *ListUninstallAegisMachinesRequest {
	s.RegionIdStr = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetRegionNo(v string) *ListUninstallAegisMachinesRequest {
	s.RegionNo = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetRemark(v string) *ListUninstallAegisMachinesRequest {
	s.Remark = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetSourceIp(v string) *ListUninstallAegisMachinesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) SetVendor(v int32) *ListUninstallAegisMachinesRequest {
	s.Vendor = &v
	return s
}

func (s *ListUninstallAegisMachinesRequest) Validate() error {
	return dara.Validate(s)
}
