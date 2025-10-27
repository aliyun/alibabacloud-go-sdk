// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallAegisMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUninstallAegisMachinesResponseBody
	GetCurrentPage() *int32
	SetMachineList(v []*ListUninstallAegisMachinesResponseBodyMachineList) *ListUninstallAegisMachinesResponseBody
	GetMachineList() []*ListUninstallAegisMachinesResponseBodyMachineList
	SetPageSize(v int32) *ListUninstallAegisMachinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUninstallAegisMachinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUninstallAegisMachinesResponseBody
	GetTotalCount() *int32
}

type ListUninstallAegisMachinesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the information about servers.
	MachineList []*ListUninstallAegisMachinesResponseBodyMachineList `json:"MachineList,omitempty" xml:"MachineList,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 151F6EB6-D5F3-417A-AF7B-4D84975DB586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 44
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUninstallAegisMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallAegisMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUninstallAegisMachinesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUninstallAegisMachinesResponseBody) GetMachineList() []*ListUninstallAegisMachinesResponseBodyMachineList {
	return s.MachineList
}

func (s *ListUninstallAegisMachinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUninstallAegisMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUninstallAegisMachinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUninstallAegisMachinesResponseBody) SetCurrentPage(v int32) *ListUninstallAegisMachinesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBody) SetMachineList(v []*ListUninstallAegisMachinesResponseBodyMachineList) *ListUninstallAegisMachinesResponseBody {
	s.MachineList = v
	return s
}

func (s *ListUninstallAegisMachinesResponseBody) SetPageSize(v int32) *ListUninstallAegisMachinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBody) SetRequestId(v string) *ListUninstallAegisMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBody) SetTotalCount(v int32) *ListUninstallAegisMachinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBody) Validate() error {
	if s.MachineList != nil {
		for _, item := range s.MachineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUninstallAegisMachinesResponseBodyMachineList struct {
	// The ID of the server.
	//
	// example:
	//
	// sas-bdrvxb4b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 120.79.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	MachineRegion *string `json:"MachineRegion,omitempty" xml:"MachineRegion,omitempty"`
	// The operating system of the server. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// 	- **windows-2003**
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The ID of the region in which the server resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 6d5b361f-958d-48a8-a9d2-d6e82c1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The source of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
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
	// The name of the service provider (SP) for the server.
	//
	// Valid values:
	//
	// 	- **ALIYUN**: Alibaba Cloud
	//
	// 	- **OUT**: a third-party service provider
	//
	// 	- **IDC**: a data center
	//
	// 	- **TENCENT**: Tencent Cloud
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud
	//
	// 	- **Microsoft**: Microsoft
	//
	// 	- **AWS**: Amazon Web Services (AWS)
	//
	// 	- **TRIPARTITE**: a lightweight server
	//
	// example:
	//
	// ALIYUN
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s ListUninstallAegisMachinesResponseBodyMachineList) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallAegisMachinesResponseBodyMachineList) GoString() string {
	return s.String()
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetMachineRegion() *string {
	return s.MachineRegion
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetOs() *string {
	return s.Os
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetUuid() *string {
	return s.Uuid
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) GetVendorName() *string {
	return s.VendorName
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetInstanceId(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.InstanceId = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetInstanceName(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.InstanceName = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetInternetIp(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.InternetIp = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetIntranetIp(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.IntranetIp = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetMachineRegion(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.MachineRegion = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetOs(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.Os = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetRegionId(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.RegionId = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetUuid(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.Uuid = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetVendor(v int32) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.Vendor = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) SetVendorName(v string) *ListUninstallAegisMachinesResponseBodyMachineList {
	s.VendorName = &v
	return s
}

func (s *ListUninstallAegisMachinesResponseBodyMachineList) Validate() error {
	return dara.Validate(s)
}
