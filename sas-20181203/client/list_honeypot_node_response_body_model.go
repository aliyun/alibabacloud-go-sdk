// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotNodeResponseBody
	GetCode() *string
	SetHoneypotNodeList(v []*ListHoneypotNodeResponseBodyHoneypotNodeList) *ListHoneypotNodeResponseBody
	GetHoneypotNodeList() []*ListHoneypotNodeResponseBodyHoneypotNodeList
	SetHttpStatusCode(v int32) *ListHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHoneypotNodeResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotNodeResponseBodyPageInfo) *ListHoneypotNodeResponseBody
	GetPageInfo() *ListHoneypotNodeResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotNodeResponseBody
	GetSuccess() *bool
}

type ListHoneypotNodeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array that consists of the information about the management nodes.
	HoneypotNodeList []*ListHoneypotNodeResponseBodyHoneypotNodeList `json:"HoneypotNodeList,omitempty" xml:"HoneypotNodeList,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotNodeResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 337BEA70-B03D-5370-8420-436F3FCD9924
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotNodeResponseBody) GetHoneypotNodeList() []*ListHoneypotNodeResponseBodyHoneypotNodeList {
	return s.HoneypotNodeList
}

func (s *ListHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotNodeResponseBody) GetPageInfo() *ListHoneypotNodeResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotNodeResponseBody) SetCode(v string) *ListHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetHoneypotNodeList(v []*ListHoneypotNodeResponseBodyHoneypotNodeList) *ListHoneypotNodeResponseBody {
	s.HoneypotNodeList = v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *ListHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetMessage(v string) *ListHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetPageInfo(v *ListHoneypotNodeResponseBodyPageInfo) *ListHoneypotNodeResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetRequestId(v string) *ListHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotNodeResponseBody) SetSuccess(v bool) *ListHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotNodeResponseBodyHoneypotNodeList struct {
	// Indicates whether a honeypot is allowed to access the Internet. Valid values:
	//
	// 	- **true**: The honeypot is allowed to access the Internet.
	//
	// 	- **false**: The honeypot is not allowed to access the Internet.
	//
	// example:
	//
	// true
	AllowHoneypotAccessInternet *bool `json:"AllowHoneypotAccessInternet,omitempty" xml:"AllowHoneypotAccessInternet,omitempty"`
	// The time when the management node was created.
	//
	// example:
	//
	// 2022-08-04 15:52:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the management node. Default value: **false**. Valid values:
	//
	// 	- **false**: non-default type
	//
	// 	- **true**: default type
	//
	// example:
	//
	// false
	DefaultNode *bool `json:"DefaultNode,omitempty" xml:"DefaultNode,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1fs3qsc1msa3512k****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The maximum number of honeypots that can be deployed to the management node.
	//
	// example:
	//
	// 10
	HoneypotTotalCount *int32 `json:"HoneypotTotalCount,omitempty" xml:"HoneypotTotalCount,omitempty"`
	// The number of honeypots that are deployed to the management node.
	//
	// example:
	//
	// 2
	HoneypotUsedCount *int32 `json:"HoneypotUsedCount,omitempty" xml:"HoneypotUsedCount,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// 7d110ca6-05ee-4149-8042-13ad1a41fd****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The IP address of the management node.
	//
	// example:
	//
	// 119.180.XX.XX
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// cyct_cnymu
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The maximum number of probes that can be deployed for the management node.
	//
	// example:
	//
	// 5
	ProbeTotalCount *int32 `json:"ProbeTotalCount,omitempty" xml:"ProbeTotalCount,omitempty"`
	// The number of probes that are deployed for the management node.
	//
	// example:
	//
	// 2
	ProbeUsedCount *int32 `json:"ProbeUsedCount,omitempty" xml:"ProbeUsedCount,omitempty"`
	// An array consisting of the CIDR blocks that are allowed to access the management node.
	SecurityGroupProbeIpList []*string `json:"SecurityGroupProbeIpList,omitempty" xml:"SecurityGroupProbeIpList,omitempty" type:"Repeated"`
	// The status of the management node. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: abnormal
	//
	// example:
	//
	// 1
	TotalStatus *int32 `json:"TotalStatus,omitempty" xml:"TotalStatus,omitempty"`
	// Indicates whether the management node can be upgraded. Valid values:
	//
	// 	- **false**: no
	//
	// 	- **true**: yes
	//
	// example:
	//
	// true
	UpgradeAvailable *bool `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
}

func (s ListHoneypotNodeResponseBodyHoneypotNodeList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotNodeResponseBodyHoneypotNodeList) GoString() string {
	return s.String()
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetAllowHoneypotAccessInternet() *bool {
	return s.AllowHoneypotAccessInternet
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetDefaultNode() *bool {
	return s.DefaultNode
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetHoneypotTotalCount() *int32 {
	return s.HoneypotTotalCount
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetHoneypotUsedCount() *int32 {
	return s.HoneypotUsedCount
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetNodeIp() *string {
	return s.NodeIp
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetProbeTotalCount() *int32 {
	return s.ProbeTotalCount
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetProbeUsedCount() *int32 {
	return s.ProbeUsedCount
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetSecurityGroupProbeIpList() []*string {
	return s.SecurityGroupProbeIpList
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetTotalStatus() *int32 {
	return s.TotalStatus
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) GetUpgradeAvailable() *bool {
	return s.UpgradeAvailable
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetAllowHoneypotAccessInternet(v bool) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.AllowHoneypotAccessInternet = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetCreateTime(v string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.CreateTime = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetDefaultNode(v bool) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.DefaultNode = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetEcsInstanceId(v string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.EcsInstanceId = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetHoneypotTotalCount(v int32) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.HoneypotTotalCount = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetHoneypotUsedCount(v int32) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.HoneypotUsedCount = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetNodeId(v string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetNodeIp(v string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.NodeIp = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetNodeName(v string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.NodeName = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetProbeTotalCount(v int32) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.ProbeTotalCount = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetProbeUsedCount(v int32) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.ProbeUsedCount = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetSecurityGroupProbeIpList(v []*string) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.SecurityGroupProbeIpList = v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetTotalStatus(v int32) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.TotalStatus = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) SetUpgradeAvailable(v bool) *ListHoneypotNodeResponseBodyHoneypotNodeList {
	s.UpgradeAvailable = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyHoneypotNodeList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotNodeResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
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

func (s ListHoneypotNodeResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotNodeResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotNodeResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotNodeResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotNodeResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotNodeResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotNodeResponseBodyPageInfo) SetCount(v int32) *ListHoneypotNodeResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotNodeResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotNodeResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotNodeResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotNodeResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
