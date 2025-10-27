// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfflineMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOfflineMachinesResponseBody
	GetCurrentPage() *int32
	SetMachineList(v []*DescribeOfflineMachinesResponseBodyMachineList) *DescribeOfflineMachinesResponseBody
	GetMachineList() []*DescribeOfflineMachinesResponseBodyMachineList
	SetPageSize(v int32) *DescribeOfflineMachinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeOfflineMachinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOfflineMachinesResponseBody
	GetTotalCount() *int32
}

type DescribeOfflineMachinesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the information about servers.
	MachineList []*DescribeOfflineMachinesResponseBodyMachineList `json:"MachineList,omitempty" xml:"MachineList,omitempty" type:"Repeated"`
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
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 44
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOfflineMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfflineMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfflineMachinesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOfflineMachinesResponseBody) GetMachineList() []*DescribeOfflineMachinesResponseBodyMachineList {
	return s.MachineList
}

func (s *DescribeOfflineMachinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOfflineMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOfflineMachinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOfflineMachinesResponseBody) SetCurrentPage(v int32) *DescribeOfflineMachinesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBody) SetMachineList(v []*DescribeOfflineMachinesResponseBodyMachineList) *DescribeOfflineMachinesResponseBody {
	s.MachineList = v
	return s
}

func (s *DescribeOfflineMachinesResponseBody) SetPageSize(v int32) *DescribeOfflineMachinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBody) SetRequestId(v string) *DescribeOfflineMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBody) SetTotalCount(v int32) *DescribeOfflineMachinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBody) Validate() error {
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

type DescribeOfflineMachinesResponseBodyMachineList struct {
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

func (s DescribeOfflineMachinesResponseBodyMachineList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfflineMachinesResponseBodyMachineList) GoString() string {
	return s.String()
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetMachineRegion() *string {
	return s.MachineRegion
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetOs() *string {
	return s.Os
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) GetVendorName() *string {
	return s.VendorName
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetInstanceId(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.InstanceId = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetInstanceName(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.InstanceName = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetInternetIp(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.InternetIp = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetIntranetIp(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetMachineRegion(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.MachineRegion = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetOs(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.Os = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetRegionId(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.RegionId = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetUuid(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.Uuid = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetVendor(v int32) *DescribeOfflineMachinesResponseBodyMachineList {
	s.Vendor = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) SetVendorName(v string) *DescribeOfflineMachinesResponseBodyMachineList {
	s.VendorName = &v
	return s
}

func (s *DescribeOfflineMachinesResponseBodyMachineList) Validate() error {
	return dara.Validate(s)
}
