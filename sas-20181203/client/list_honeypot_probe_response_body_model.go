// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotProbeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListHoneypotProbeResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListHoneypotProbeResponseBodyList) *ListHoneypotProbeResponseBody
	GetList() []*ListHoneypotProbeResponseBodyList
	SetMessage(v string) *ListHoneypotProbeResponseBody
	GetMessage() *string
	SetPageInfo(v *ListHoneypotProbeResponseBodyPageInfo) *ListHoneypotProbeResponseBody
	GetPageInfo() *ListHoneypotProbeResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotProbeResponseBody
	GetSuccess() *bool
}

type ListHoneypotProbeResponseBody struct {
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// An array that consists of the details about the probe.
	List []*ListHoneypotProbeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListHoneypotProbeResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4BC9E610-21BE-537F-82EF-144A60D5A970
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

func (s ListHoneypotProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotProbeResponseBody) GetList() []*ListHoneypotProbeResponseBodyList {
	return s.List
}

func (s *ListHoneypotProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotProbeResponseBody) GetPageInfo() *ListHoneypotProbeResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotProbeResponseBody) SetCode(v string) *ListHoneypotProbeResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetHttpStatusCode(v int32) *ListHoneypotProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetList(v []*ListHoneypotProbeResponseBodyList) *ListHoneypotProbeResponseBody {
	s.List = v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetMessage(v string) *ListHoneypotProbeResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetPageInfo(v *ListHoneypotProbeResponseBodyPageInfo) *ListHoneypotProbeResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetRequestId(v string) *ListHoneypotProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotProbeResponseBody) SetSuccess(v bool) *ListHoneypotProbeResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotProbeResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type ListHoneypotProbeResponseBodyList struct {
	// The information about the management node.
	ControlNode *ListHoneypotProbeResponseBodyListControlNode `json:"ControlNode,omitempty" xml:"ControlNode,omitempty" type:"Struct"`
	// The time when the probe was deployed.
	//
	// example:
	//
	// 1669363825000
	DeployTime *int64 `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// prod-pinpoint-hd1b
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The IP address of the server on which the probe is installed.
	//
	// example:
	//
	// 33.53.XX.XX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// The operating system of the server on which the probe is deployed. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// example:
	//
	// windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The ID of the probe.
	//
	// example:
	//
	// 4d167bb3-dd09-4a6a-a179-d5d6a5b0****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The type of the probe. Valid values:
	//
	// 	- **host_probe**: host probe
	//
	// 	- **vpc_black_hole_probe**: VPC probe
	//
	// example:
	//
	// host_probe
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The version of the probe.
	//
	// example:
	//
	// 18060096
	ProbeVersion *string `json:"ProbeVersion,omitempty" xml:"ProbeVersion,omitempty"`
	// The status of the probe. Valid values:
	//
	// 	- **installed**: installed
	//
	// 	- **install_failed**: installation failed
	//
	// 	- **online**: online
	//
	// 	- **offline**: offline
	//
	// 	- **unnormal**: abnormal
	//
	// 	- **unprobe**: unauthorized
	//
	// 	- **uninstalling**: being uninstalled
	//
	// 	- **uninstalled**: uninstalled
	//
	// 	- **uninstall_failed**: uninstallation failed
	//
	// 	- **not_exist**: not installed
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server to which the host probe is deployed.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the VPC in which the VPC probe is deployed.
	//
	// example:
	//
	// vpc-5gu8iu68w9b472jbb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListHoneypotProbeResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeResponseBodyList) GetControlNode() *ListHoneypotProbeResponseBodyListControlNode {
	return s.ControlNode
}

func (s *ListHoneypotProbeResponseBodyList) GetDeployTime() *int64 {
	return s.DeployTime
}

func (s *ListHoneypotProbeResponseBodyList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListHoneypotProbeResponseBodyList) GetHostIp() *string {
	return s.HostIp
}

func (s *ListHoneypotProbeResponseBodyList) GetOsType() *string {
	return s.OsType
}

func (s *ListHoneypotProbeResponseBodyList) GetProbeId() *string {
	return s.ProbeId
}

func (s *ListHoneypotProbeResponseBodyList) GetProbeType() *string {
	return s.ProbeType
}

func (s *ListHoneypotProbeResponseBodyList) GetProbeVersion() *string {
	return s.ProbeVersion
}

func (s *ListHoneypotProbeResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListHoneypotProbeResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListHoneypotProbeResponseBodyList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListHoneypotProbeResponseBodyList) SetControlNode(v *ListHoneypotProbeResponseBodyListControlNode) *ListHoneypotProbeResponseBodyList {
	s.ControlNode = v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetDeployTime(v int64) *ListHoneypotProbeResponseBodyList {
	s.DeployTime = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetDisplayName(v string) *ListHoneypotProbeResponseBodyList {
	s.DisplayName = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetHostIp(v string) *ListHoneypotProbeResponseBodyList {
	s.HostIp = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetOsType(v string) *ListHoneypotProbeResponseBodyList {
	s.OsType = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetProbeId(v string) *ListHoneypotProbeResponseBodyList {
	s.ProbeId = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetProbeType(v string) *ListHoneypotProbeResponseBodyList {
	s.ProbeType = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetProbeVersion(v string) *ListHoneypotProbeResponseBodyList {
	s.ProbeVersion = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetStatus(v string) *ListHoneypotProbeResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetUuid(v string) *ListHoneypotProbeResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) SetVpcId(v string) *ListHoneypotProbeResponseBodyList {
	s.VpcId = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyList) Validate() error {
	if s.ControlNode != nil {
		if err := s.ControlNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHoneypotProbeResponseBodyListControlNode struct {
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// i-uf6eq0rlvu1mkh0p****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 8ec9da17-c0e7-4642-aad6-defc9722****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// HoneypotNode1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s ListHoneypotProbeResponseBodyListControlNode) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeResponseBodyListControlNode) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeResponseBodyListControlNode) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListHoneypotProbeResponseBodyListControlNode) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotProbeResponseBodyListControlNode) GetNodeName() *string {
	return s.NodeName
}

func (s *ListHoneypotProbeResponseBodyListControlNode) SetEcsInstanceId(v string) *ListHoneypotProbeResponseBodyListControlNode {
	s.EcsInstanceId = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyListControlNode) SetNodeId(v string) *ListHoneypotProbeResponseBodyListControlNode {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyListControlNode) SetNodeName(v string) *ListHoneypotProbeResponseBodyListControlNode {
	s.NodeName = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyListControlNode) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotProbeResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotProbeResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotProbeResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotProbeResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotProbeResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotProbeResponseBodyPageInfo) SetCount(v int32) *ListHoneypotProbeResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotProbeResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotProbeResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotProbeResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotProbeResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
