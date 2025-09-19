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
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 4
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The operating system of the server.
	//
	// >  The value of this parameter is the value of the Values parameter that is returned by calling the [DescribeCriteria](~~DescribeCriteria~~) operation. If the value of the **Name*	- parameter in the response is **osType**, the value of the **Values*	- parameter indicates an operating system.
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The number of entries to return on each page. Default value: **5**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the server resides.
	//
	// >  The value of this parameter is the value of the Values parameter that is returned by calling the [DescribeCriteria](~~DescribeCriteria~~) operation. If the value of the **Name*	- parameter in the response is **regionId**, the value of the **Values*	- parameter indicates a region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdStr *string `json:"RegionIdStr,omitempty" xml:"RegionIdStr,omitempty"`
	// The region in which the server resides.
	//
	// >  The value of this parameter is the value of the Values parameter that is returned by calling the [DescribeCriteria](~~DescribeCriteria~~) operation. If the value of the **Name*	- parameter in the response is **regionId**, the value of the **Values*	- parameter indicates a region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The information about the server that you want to query. The value can be the name or the public IP address of the server.
	//
	// example:
	//
	// 172.20.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 180.113.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: a third-party cloud server
	//
	// 	- **2**: a server in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
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
