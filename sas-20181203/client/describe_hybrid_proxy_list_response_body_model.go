// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeHybridProxyListResponseBodyPageInfo) *DescribeHybridProxyListResponseBody
	GetPageInfo() *DescribeHybridProxyListResponseBodyPageInfo
	SetProxyList(v []*DescribeHybridProxyListResponseBodyProxyList) *DescribeHybridProxyListResponseBody
	GetProxyList() []*DescribeHybridProxyListResponseBodyProxyList
	SetRequestId(v string) *DescribeHybridProxyListResponseBody
	GetRequestId() *string
}

type DescribeHybridProxyListResponseBody struct {
	// The pagination information.
	PageInfo *DescribeHybridProxyListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The information about the proxy clusters.
	ProxyList []*DescribeHybridProxyListResponseBodyProxyList `json:"ProxyList,omitempty" xml:"ProxyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 75801E5D-E2EB-5C1D-B65D-2F7D2B00EF93
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridProxyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyListResponseBody) GetPageInfo() *DescribeHybridProxyListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeHybridProxyListResponseBody) GetProxyList() []*DescribeHybridProxyListResponseBodyProxyList {
	return s.ProxyList
}

func (s *DescribeHybridProxyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridProxyListResponseBody) SetPageInfo(v *DescribeHybridProxyListResponseBodyPageInfo) *DescribeHybridProxyListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeHybridProxyListResponseBody) SetProxyList(v []*DescribeHybridProxyListResponseBodyProxyList) *DescribeHybridProxyListResponseBody {
	s.ProxyList = v
	return s
}

func (s *DescribeHybridProxyListResponseBody) SetRequestId(v string) *DescribeHybridProxyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridProxyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyListResponseBodyPageInfo struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
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
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridProxyListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) SetCount(v int32) *DescribeHybridProxyListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeHybridProxyListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) SetPageSize(v int32) *DescribeHybridProxyListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeHybridProxyListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyListResponseBodyProxyList struct {
	// The number of servers that are connected to the proxy instance.
	//
	// example:
	//
	// 10
	ClientCount *int32 `json:"ClientCount,omitempty" xml:"ClientCount,omitempty"`
	// The version of the proxy instance.
	//
	// example:
	//
	// proxy_01_05
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-uf61q03boqhhmeai1XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 47.76.XXX.XXX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.23.XXX.XXX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The UUID of the proxy node.
	//
	// example:
	//
	// inet-proxy-3bb11fad-37d6-4aee-9c37-b0ad1612a18e
	ProxyUuid *string `json:"ProxyUuid,omitempty" xml:"ProxyUuid,omitempty"`
	// The status of the proxy server. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server that is connected to the proxy instance.
	//
	// example:
	//
	// 59a9d158-b2f0-4766-a893-ae67b943XXXX
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeHybridProxyListResponseBodyProxyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyListResponseBodyProxyList) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetClientCount() *int32 {
	return s.ClientCount
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetProxyUuid() *string {
	return s.ProxyUuid
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridProxyListResponseBodyProxyList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetClientCount(v int32) *DescribeHybridProxyListResponseBodyProxyList {
	s.ClientCount = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetCurrentVersion(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetInstanceId(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetInstanceName(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetInternetIp(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.InternetIp = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetIntranetIp(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetProxyUuid(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.ProxyUuid = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetStatus(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.Status = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) SetUuid(v string) *DescribeHybridProxyListResponseBodyProxyList {
	s.Uuid = &v
	return s
}

func (s *DescribeHybridProxyListResponseBodyProxyList) Validate() error {
	return dara.Validate(s)
}
