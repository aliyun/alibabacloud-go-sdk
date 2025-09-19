// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyLinkedClientListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeHybridProxyLinkedClientListResponseBodyList) *DescribeHybridProxyLinkedClientListResponseBody
	GetList() []*DescribeHybridProxyLinkedClientListResponseBodyList
	SetPageInfo(v *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) *DescribeHybridProxyLinkedClientListResponseBody
	GetPageInfo() *DescribeHybridProxyLinkedClientListResponseBodyPageInfo
	SetRequestId(v string) *DescribeHybridProxyLinkedClientListResponseBody
	GetRequestId() *string
}

type DescribeHybridProxyLinkedClientListResponseBody struct {
	// The returned data.
	List []*DescribeHybridProxyLinkedClientListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeHybridProxyLinkedClientListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B01B804F-947C-5623-B050-1C8FDFA796CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridProxyLinkedClientListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyLinkedClientListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) GetList() []*DescribeHybridProxyLinkedClientListResponseBodyList {
	return s.List
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) GetPageInfo() *DescribeHybridProxyLinkedClientListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) SetList(v []*DescribeHybridProxyLinkedClientListResponseBodyList) *DescribeHybridProxyLinkedClientListResponseBody {
	s.List = v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) SetPageInfo(v *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) *DescribeHybridProxyLinkedClientListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) SetRequestId(v string) *DescribeHybridProxyLinkedClientListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyLinkedClientListResponseBodyList struct {
	// The name of the server group.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-bp1a69mvjujbakxu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-0****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the operating system.
	//
	// example:
	//
	// centos-xxx
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The name of the operating system for your asset.
	//
	// example:
	//
	// centos
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The ID of the region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region in which the server resides.
	//
	// example:
	//
	// cn-qingdao
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The status of the Security Center agent.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the tag added to the server.
	//
	// example:
	//
	// latest
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// ALIYUN
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s DescribeHybridProxyLinkedClientListResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyLinkedClientListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetOs() *string {
	return s.Os
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetOsName() *string {
	return s.OsName
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetTag() *string {
	return s.Tag
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) GetVendorName() *string {
	return s.VendorName
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetGroupName(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetInstanceId(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetInstanceName(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetInternetIp(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetIntranetIp(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetOs(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.Os = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetOsName(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.OsName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetRegionId(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetRegionName(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.RegionName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetStatus(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetTag(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.Tag = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetUuid(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) SetVendorName(v string) *DescribeHybridProxyLinkedClientListResponseBodyList {
	s.VendorName = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyLinkedClientListResponseBodyPageInfo struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridProxyLinkedClientListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyLinkedClientListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) SetCount(v int32) *DescribeHybridProxyLinkedClientListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeHybridProxyLinkedClientListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) SetPageSize(v int32) *DescribeHybridProxyLinkedClientListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeHybridProxyLinkedClientListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
