// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyClusterListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v []*DescribeHybridProxyClusterListResponseBodyClusterList) *DescribeHybridProxyClusterListResponseBody
	GetClusterList() []*DescribeHybridProxyClusterListResponseBodyClusterList
	SetPageInfo(v *DescribeHybridProxyClusterListResponseBodyPageInfo) *DescribeHybridProxyClusterListResponseBody
	GetPageInfo() *DescribeHybridProxyClusterListResponseBodyPageInfo
	SetRequestId(v string) *DescribeHybridProxyClusterListResponseBody
	GetRequestId() *string
}

type DescribeHybridProxyClusterListResponseBody struct {
	// The proxy clusters.
	ClusterList []*DescribeHybridProxyClusterListResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeHybridProxyClusterListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C1A36413-50B2-5B2F-843F-EB14C582713F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridProxyClusterListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyClusterListResponseBody) GetClusterList() []*DescribeHybridProxyClusterListResponseBodyClusterList {
	return s.ClusterList
}

func (s *DescribeHybridProxyClusterListResponseBody) GetPageInfo() *DescribeHybridProxyClusterListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeHybridProxyClusterListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridProxyClusterListResponseBody) SetClusterList(v []*DescribeHybridProxyClusterListResponseBodyClusterList) *DescribeHybridProxyClusterListResponseBody {
	s.ClusterList = v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBody) SetPageInfo(v *DescribeHybridProxyClusterListResponseBodyPageInfo) *DescribeHybridProxyClusterListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBody) SetRequestId(v string) *DescribeHybridProxyClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyClusterListResponseBodyClusterList struct {
	// The ID of the credential that is used for cluster authentication.
	//
	// example:
	//
	// test
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The key of the credential that is used for cluster authentication.
	//
	// example:
	//
	// test
	AuthKeySecret *string `json:"AuthKeySecret,omitempty" xml:"AuthKeySecret,omitempty"`
	// The number of servers that are connected to the proxy cluster.
	//
	// example:
	//
	// 10
	ClientCount *int32 `json:"ClientCount,omitempty" xml:"ClientCount,omitempty"`
	// The name of the proxy cluster.
	//
	// example:
	//
	// idc-sas-proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The installation command for the node of the proxy cluster.
	//
	// example:
	//
	// test
	InstallCommand *string `json:"InstallCommand,omitempty" xml:"InstallCommand,omitempty"`
	// The endpoint of the cluster. An IP address or a domain name is specified.
	//
	// example:
	//
	// 114.115.XXX.XXX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The timestamp when the cluster last sent a heartbeat message. Unit: milliseconds.
	//
	// example:
	//
	// 1608304654000
	LastHeartTime *int64 `json:"LastHeartTime,omitempty" xml:"LastHeartTime,omitempty"`
	// The number of proxy nodes.
	//
	// example:
	//
	// 3
	ProxyCount *int32 `json:"ProxyCount,omitempty" xml:"ProxyCount,omitempty"`
	// The description of the proxy cluster.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the cluster.
	//
	// example:
	//
	// offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHybridProxyClusterListResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyClusterListResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetAuthKeySecret() *string {
	return s.AuthKeySecret
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetClientCount() *int32 {
	return s.ClientCount
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetLastHeartTime() *int64 {
	return s.LastHeartTime
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetProxyCount() *int32 {
	return s.ProxyCount
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetRemark() *string {
	return s.Remark
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetAuthKey(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.AuthKey = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetAuthKeySecret(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.AuthKeySecret = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetClientCount(v int32) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.ClientCount = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetClusterName(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetInstallCommand(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.InstallCommand = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetIp(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.Ip = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetLastHeartTime(v int64) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.LastHeartTime = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetProxyCount(v int32) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.ProxyCount = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetRemark(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.Remark = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) SetStatus(v string) *DescribeHybridProxyClusterListResponseBodyClusterList {
	s.Status = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyClusterList) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridProxyClusterListResponseBodyPageInfo struct {
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
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridProxyClusterListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyClusterListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) SetCount(v int32) *DescribeHybridProxyClusterListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeHybridProxyClusterListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) SetPageSize(v int32) *DescribeHybridProxyClusterListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeHybridProxyClusterListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
