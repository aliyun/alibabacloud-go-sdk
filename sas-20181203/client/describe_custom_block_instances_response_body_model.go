// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*DescribeCustomBlockInstancesResponseBodyInstanceList) *DescribeCustomBlockInstancesResponseBody
	GetInstanceList() []*DescribeCustomBlockInstancesResponseBodyInstanceList
	SetPageInfo(v *DescribeCustomBlockInstancesResponseBodyPageInfo) *DescribeCustomBlockInstancesResponseBody
	GetPageInfo() *DescribeCustomBlockInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeCustomBlockInstancesResponseBody
	GetRequestId() *string
}

type DescribeCustomBlockInstancesResponseBody struct {
	// The server ID.
	InstanceList []*DescribeCustomBlockInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeCustomBlockInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D81DD78E-E006-5C65-A171-C8CB09XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomBlockInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockInstancesResponseBody) GetInstanceList() []*DescribeCustomBlockInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeCustomBlockInstancesResponseBody) GetPageInfo() *DescribeCustomBlockInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCustomBlockInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomBlockInstancesResponseBody) SetInstanceList(v []*DescribeCustomBlockInstancesResponseBodyInstanceList) *DescribeCustomBlockInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBody) SetPageInfo(v *DescribeCustomBlockInstancesResponseBodyPageInfo) *DescribeCustomBlockInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBody) SetRequestId(v string) *DescribeCustomBlockInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomBlockInstancesResponseBodyInstanceList struct {
	// The status of the host network extension. Valid values:
	//
	// 	- **true**: online
	//
	// 	- **false**: offline
	//
	// example:
	//
	// true
	AliNetOnline *bool `json:"AliNetOnline,omitempty" xml:"AliNetOnline,omitempty"`
	// The blocking type. Valid values:
	//
	// 	- **group**: security group
	//
	// 	- **alinet**: host network extension
	//
	// example:
	//
	// group
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// AliNetNotOnline
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// myInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 116.62.121.1xx
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.1.xx
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Indicates whether the rule is enabled for the server.
	//
	// 	- **2**: enabling failed
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information that is returned after brute-force attacks are blocked.
	//
	// example:
	//
	// {"aliUid":*******,"groupId":"sg-xxxx","groupName":"Sas_Malicious_Ip_Security_Group","groupType":"normal","instanceId":"i-xxxx","regionId":"cn-shenzhen","vpcId":"vpc-xxxxxxxx"}
	SuccessInfo *string `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// f2d6e901-1004-4ca8-9dae-53ec04a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCustomBlockInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetAliNetOnline() *bool {
	return s.AliNetOnline
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetBlockType() *string {
	return s.BlockType
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetSuccessInfo() *string {
	return s.SuccessInfo
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetAliNetOnline(v bool) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.AliNetOnline = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetBlockType(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.BlockType = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetErrorCode(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetInstanceName(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetInternetIp(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.InternetIp = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetIntranetIp(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetStatus(v int32) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetSuccessInfo(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.SuccessInfo = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) SetUuid(v string) *DescribeCustomBlockInstancesResponseBodyInstanceList {
	s.Uuid = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomBlockInstancesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
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
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of servers to which the defense rule is applied.
	//
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomBlockInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeCustomBlockInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCustomBlockInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeCustomBlockInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCustomBlockInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCustomBlockInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
