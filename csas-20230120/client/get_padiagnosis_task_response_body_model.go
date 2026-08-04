// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPADiagnosisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosisTask(v *GetPADiagnosisTaskResponseBodyDiagnosisTask) *GetPADiagnosisTaskResponseBody
	GetDiagnosisTask() *GetPADiagnosisTaskResponseBodyDiagnosisTask
	SetRequestId(v string) *GetPADiagnosisTaskResponseBody
	GetRequestId() *string
}

type GetPADiagnosisTaskResponseBody struct {
	// The diagnostic task.
	DiagnosisTask *GetPADiagnosisTaskResponseBodyDiagnosisTask `json:"DiagnosisTask,omitempty" xml:"DiagnosisTask,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPADiagnosisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBody) GetDiagnosisTask() *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	return s.DiagnosisTask
}

func (s *GetPADiagnosisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPADiagnosisTaskResponseBody) SetDiagnosisTask(v *GetPADiagnosisTaskResponseBodyDiagnosisTask) *GetPADiagnosisTaskResponseBody {
	s.DiagnosisTask = v
	return s
}

func (s *GetPADiagnosisTaskResponseBody) SetRequestId(v string) *GetPADiagnosisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBody) Validate() error {
	if s.DiagnosisTask != nil {
		if err := s.DiagnosisTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTask struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2023-08-17 09:49:03
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the device.
	//
	// example:
	//
	// 76C08B0F-CEFD-8F01-C1D3-0D5B493B5EAF
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The ID of the diagnostic task.
	//
	// example:
	//
	// diag-3e0d36d6c15a0502
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The diagnostic type. Valid values:
	//
	// - **FullLink**: full-link diagnostics
	//
	// - **Application**: application diagnostics
	//
	// example:
	//
	// FullLink
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The domain name to be diagnosed.
	//
	// example:
	//
	// mtools-admin.redotpay.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the POP.
	//
	// example:
	//
	// pop-2504cd067e094750
	PopId *string `json:"PopId,omitempty" xml:"PopId,omitempty"`
	// The point of presence (POP) selection mode:
	//
	// - **AutoSelect**: automatic selection
	//
	// - **ManualSelect**: manual selection
	//
	// example:
	//
	// AutoSelect
	PopMode *string `json:"PopMode,omitempty" xml:"PopMode,omitempty"`
	// The port.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private access application protocol. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The diagnostic result.
	Result *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the task. Valid values:
	//
	// - **Running**: The task is running.
	//
	// - **Finished**: The task is complete.
	//
	// - **Failed**: The task failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Additional UDP configurations.
	UdpExtraConfigs *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs `json:"UdpExtraConfigs,omitempty" xml:"UdpExtraConfigs,omitempty" type:"Struct"`
	// The user group.
	UserGroup *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTask) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTask) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetDevTag() *string {
	return s.DevTag
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetHost() *string {
	return s.Host
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetPopId() *string {
	return s.PopId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetPopMode() *string {
	return s.PopMode
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetPort() *string {
	return s.Port
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetProtocol() *string {
	return s.Protocol
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetResult() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	return s.Result
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetStatus() *string {
	return s.Status
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetUdpExtraConfigs() *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	return s.UdpExtraConfigs
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetUserGroup() *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	return s.UserGroup
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) GetUsername() *string {
	return s.Username
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetCreateTime(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.CreateTime = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetDevTag(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.DevTag = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetDiagnoseId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.DiagnoseId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetDiagnoseType(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.DiagnoseType = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetHost(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Host = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetPopId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.PopId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetPopMode(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.PopMode = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetPort(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Port = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetProtocol(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Protocol = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetResult(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Result = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetStatus(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Status = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetUdpExtraConfigs(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.UdpExtraConfigs = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetUserGroup(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.UserGroup = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) SetUsername(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTask {
	s.Username = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTask) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.UdpExtraConfigs != nil {
		if err := s.UdpExtraConfigs.Validate(); err != nil {
			return err
		}
	}
	if s.UserGroup != nil {
		if err := s.UserGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResult struct {
	// The error message.
	//
	// example:
	//
	// device offline
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request stream.
	//
	// example:
	//
	// flow-d918b12f9b974f6489fc
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The network connectivity information.
	NetworkLinkInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo `json:"NetworkLinkInfo,omitempty" xml:"NetworkLinkInfo,omitempty" type:"Struct"`
	// The policy information.
	PolicyInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo `json:"PolicyInfo,omitempty" xml:"PolicyInfo,omitempty" type:"Struct"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GetFlowId() *string {
	return s.FlowId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GetNetworkLinkInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo {
	return s.NetworkLinkInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GetPolicyInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	return s.PolicyInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) SetErrorMessage(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) SetFlowId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	s.FlowId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) SetNetworkLinkInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	s.NetworkLinkInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) SetPolicyInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	s.PolicyInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) SetSuccess(v bool) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult {
	s.Success = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResult) Validate() error {
	if s.NetworkLinkInfo != nil {
		if err := s.NetworkLinkInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyInfo != nil {
		if err := s.PolicyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo struct {
	// The Domain Name System (DNS) information.
	Dns *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Struct"`
	// The time to first byte.
	//
	// example:
	//
	// 300
	FBT *string `json:"FBT,omitempty" xml:"FBT,omitempty"`
	// The connections between nodes.
	Links []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	// The forwarding nodes.
	Nodes []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) GetDns() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	return s.Dns
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) GetFBT() *string {
	return s.FBT
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) GetLinks() []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	return s.Links
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) GetNodes() []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	return s.Nodes
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) SetDns(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo {
	s.Dns = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) SetFBT(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo {
	s.FBT = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) SetLinks(v []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo {
	s.Links = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) SetNodes(v []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo {
	s.Nodes = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfo) Validate() error {
	if s.Dns != nil {
		if err := s.Dns.Validate(); err != nil {
			return err
		}
	}
	if s.Links != nil {
		for _, item := range s.Links {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns struct {
	// The DNS server.
	//
	// example:
	//
	// 100.100.2.136,100.100.2.138
	DnsServer *string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty"`
	// The DNS type.
	//
	// example:
	//
	// private-zone
	DnsType *string `json:"DnsType,omitempty" xml:"DnsType,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The source node.
	//
	// example:
	//
	// 2
	FromNode *int64 `json:"FromNode,omitempty" xml:"FromNode,omitempty"`
	// The intermediate hops.
	Hops [][]*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The latency.
	//
	// example:
	//
	// 10
	Latency *string `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The DNS result.
	//
	// example:
	//
	// 10.0.0.1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The destination node.
	//
	// example:
	//
	// 3
	ToNode *int64 `json:"ToNode,omitempty" xml:"ToNode,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetDnsServer() *string {
	return s.DnsServer
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetDnsType() *string {
	return s.DnsType
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetError() *string {
	return s.Error
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetFromNode() *int64 {
	return s.FromNode
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetHops() [][]*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops {
	return s.Hops
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetLatency() *string {
	return s.Latency
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetResult() *string {
	return s.Result
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetSuccess() *bool {
	return s.Success
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) GetToNode() *int64 {
	return s.ToNode
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetDnsServer(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.DnsServer = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetDnsType(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.DnsType = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetError(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.Error = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetFromNode(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.FromNode = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetHops(v [][]*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.Hops = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetLatency(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.Latency = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetResult(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.Result = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetSuccess(v bool) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.Success = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) SetToNode(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns {
	s.ToNode = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDns) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops struct {
	// The private access application address. The address can be 1 to 128 characters long and can be an IPv4 address, a CIDR block, a domain name, or a wildcard domain name.
	//
	// example:
	//
	// *******************************************
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The TTL.
	//
	// example:
	//
	// 10
	TTL *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The latency.
	//
	// example:
	//
	// 10
	Latency *string `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The geographic location.
	GeoData *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData `json:"GeoData,omitempty" xml:"GeoData,omitempty" type:"Struct"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) GetAddress() *string {
	return s.Address
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) GetTTL() *string {
	return s.TTL
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) GetLatency() *string {
	return s.Latency
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) GetGeoData() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData {
	return s.GeoData
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) SetAddress(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops {
	s.Address = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) SetTTL(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops {
	s.TTL = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) SetLatency(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops {
	s.Latency = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) SetGeoData(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops {
	s.GeoData = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHops) Validate() error {
	if s.GeoData != nil {
		if err := s.GeoData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData struct {
	// The country.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The province.
	//
	// example:
	//
	// Zhejiang
	Prov *string `json:"Prov,omitempty" xml:"Prov,omitempty"`
	// The city.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The ISP.
	//
	// example:
	//
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) GetCountry() *string {
	return s.Country
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) GetProv() *string {
	return s.Prov
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) GetCity() *string {
	return s.City
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) GetIsp() *string {
	return s.Isp
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) SetCountry(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData {
	s.Country = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) SetProv(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData {
	s.Prov = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) SetCity(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData {
	s.City = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) SetIsp(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData {
	s.Isp = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoDnsHopsGeoData) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks struct {
	// The error message.
	//
	// example:
	//
	// 0
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The source node.
	//
	// example:
	//
	// 1
	FromNode *int64 `json:"FromNode,omitempty" xml:"FromNode,omitempty"`
	// The intermediate hops.
	Hops []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops `json:"Hops,omitempty" xml:"Hops,omitempty" type:"Repeated"`
	// The latency.
	//
	// example:
	//
	// 10
	Latency *string `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The destination node.
	//
	// example:
	//
	// 2
	ToNode *int64 `json:"ToNode,omitempty" xml:"ToNode,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetError() *string {
	return s.Error
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetFromNode() *int64 {
	return s.FromNode
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetHops() []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops {
	return s.Hops
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetLatency() *string {
	return s.Latency
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetSuccess() *bool {
	return s.Success
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) GetToNode() *int64 {
	return s.ToNode
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetError(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.Error = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetFromNode(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.FromNode = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetHops(v []*GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.Hops = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetLatency(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.Latency = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetSuccess(v bool) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.Success = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) SetToNode(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks {
	s.ToNode = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinks) Validate() error {
	if s.Hops != nil {
		for _, item := range s.Hops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops struct {
	// The address.
	//
	// example:
	//
	// **********************
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The geographic location.
	GeoData *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData `json:"GeoData,omitempty" xml:"GeoData,omitempty" type:"Struct"`
	// The latency.
	//
	// example:
	//
	// 10
	Latency *string `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The Time to Live (TTL).
	//
	// example:
	//
	// 10
	TTL *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) GetAddress() *string {
	return s.Address
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) GetGeoData() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData {
	return s.GeoData
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) GetLatency() *string {
	return s.Latency
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) GetTTL() *string {
	return s.TTL
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) SetAddress(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops {
	s.Address = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) SetGeoData(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops {
	s.GeoData = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) SetLatency(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops {
	s.Latency = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) SetTTL(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops {
	s.TTL = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHops) Validate() error {
	if s.GeoData != nil {
		if err := s.GeoData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData struct {
	// The city.
	//
	// example:
	//
	// Haikou City
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The country.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The ISP.
	//
	// example:
	//
	// ChinaMobile_L2
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The province.
	//
	// example:
	//
	// ZHejiang
	Prov *string `json:"Prov,omitempty" xml:"Prov,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) GetCity() *string {
	return s.City
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) GetCountry() *string {
	return s.Country
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) GetIsp() *string {
	return s.Isp
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) GetProv() *string {
	return s.Prov
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) SetCity(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData {
	s.City = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) SetCountry(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData {
	s.Country = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) SetIsp(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData {
	s.Isp = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) SetProv(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData {
	s.Prov = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoLinksHopsGeoData) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes struct {
	// The private access application address. The address can be 1 to 128 characters long and can be an IPv4 address, a CIDR block, a domain name, or a wildcard domain name.
	//
	// example:
	//
	// 172.27.228.132
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the cloud network instance.
	//
	// example:
	//
	// vpc-xxxxxx
	CloudNetId *string `json:"CloudNetId,omitempty" xml:"CloudNetId,omitempty"`
	// The error message.
	//
	// example:
	//
	// 1
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The geographic location information.
	GeoData *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData `json:"GeoData,omitempty" xml:"GeoData,omitempty" type:"Struct"`
	// The node ID.
	//
	// example:
	//
	// 1237
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
	// example:
	//
	// 全局加速
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The English name of the node.
	//
	// example:
	//
	// Japan Private POP
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The node type.
	//
	// example:
	//
	// stunnel
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// pop-xxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetAddress() *string {
	return s.Address
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetCloudNetId() *string {
	return s.CloudNetId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetError() *string {
	return s.Error
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetGeoData() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData {
	return s.GeoData
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetId() *int64 {
	return s.Id
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetName() *string {
	return s.Name
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetNameEn() *string {
	return s.NameEn
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) GetSuccess() *bool {
	return s.Success
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetAddress(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.Address = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetCloudNetId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.CloudNetId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetError(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.Error = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetGeoData(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.GeoData = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetId(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.Id = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetName(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.Name = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetNameEn(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.NameEn = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetNodeType(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.NodeType = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetResourceId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.ResourceId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) SetSuccess(v bool) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes {
	s.Success = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodes) Validate() error {
	if s.GeoData != nil {
		if err := s.GeoData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData struct {
	// The city.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The country.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The Internet Service Provider (ISP).
	//
	// example:
	//
	// ChinaTelecom_L2
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The province.
	//
	// example:
	//
	// Zhejiang
	Prov *string `json:"Prov,omitempty" xml:"Prov,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) GetCity() *string {
	return s.City
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) GetCountry() *string {
	return s.Country
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) GetIsp() *string {
	return s.Isp
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) GetProv() *string {
	return s.Prov
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) SetCity(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData {
	s.City = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) SetCountry(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData {
	s.Country = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) SetIsp(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData {
	s.Isp = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) SetProv(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData {
	s.Prov = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultNetworkLinkInfoNodesGeoData) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo struct {
	// The device information.
	DeviceAttributeInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo `json:"DeviceAttributeInfo,omitempty" xml:"DeviceAttributeInfo,omitempty" type:"Struct"`
	// The processing duration.
	//
	// example:
	//
	// 1000
	ProcessTime *int64 `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	// The name of the routing policy.
	RouteStrategyInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo `json:"RouteStrategyInfo,omitempty" xml:"RouteStrategyInfo,omitempty" type:"Struct"`
	// The user group information.
	UserGroupInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo `json:"UserGroupInfo,omitempty" xml:"UserGroupInfo,omitempty" type:"Struct"`
	// The zero-trust policy information.
	ZeroTrustPolicyInfo *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo `json:"ZeroTrustPolicyInfo,omitempty" xml:"ZeroTrustPolicyInfo,omitempty" type:"Struct"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GetDeviceAttributeInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	return s.DeviceAttributeInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GetProcessTime() *int64 {
	return s.ProcessTime
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GetRouteStrategyInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo {
	return s.RouteStrategyInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GetUserGroupInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	return s.UserGroupInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) GetZeroTrustPolicyInfo() *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo {
	return s.ZeroTrustPolicyInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) SetDeviceAttributeInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	s.DeviceAttributeInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) SetProcessTime(v int64) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	s.ProcessTime = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) SetRouteStrategyInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	s.RouteStrategyInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) SetUserGroupInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	s.UserGroupInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) SetZeroTrustPolicyInfo(v *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo {
	s.ZeroTrustPolicyInfo = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfo) Validate() error {
	if s.DeviceAttributeInfo != nil {
		if err := s.DeviceAttributeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RouteStrategyInfo != nil {
		if err := s.RouteStrategyInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserGroupInfo != nil {
		if err := s.UserGroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ZeroTrustPolicyInfo != nil {
		if err := s.ZeroTrustPolicyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo struct {
	// The ID of the device.
	//
	// example:
	//
	// E9EE1CE7-4AA0-521D-B8E1-E13E47F05E94
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The operating system of the device. Valid values:
	//
	// - **Windows**: Windows
	//
	// - **macOS**: macOS
	//
	// - **Linux**: Linux
	//
	// - **Android**: Android
	//
	// - **iOS**: iOS
	//
	// - **Windows_Wuying**: Cloud Desktop
	//
	// example:
	//
	// macos
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The name of the file.
	File []*string `json:"File,omitempty" xml:"File,omitempty" type:"Repeated"`
	// The firewall.
	//
	// example:
	//
	// [{\\"Platform\\":\\"windows\\",\\"Status\\":\\"disabled\\"},{\\"Platform\\":\\"macos\\",\\"Status\\":\\"disabled\\"},{\\"Platform\\":\\"linux\\",\\"Status\\":\\"disabled\\"}]
	Firewall *string `json:"Firewall,omitempty" xml:"Firewall,omitempty"`
	// The name of the device. The name can be 1 to 128 characters long and can include letters, numbers, and the following special characters: . , ; - _ / @ and spaces. To query for all devices with names containing 4-byte UTF-8 characters, enter only an underscore (_).
	//
	// example:
	//
	// DESKTOP-CVTB5KT.CXISHD01.CATHAY_INS.CHN
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The private IP address of the device.
	//
	// example:
	//
	// 10.5.208.122
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.98.146.136
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The MAC address of the device.
	//
	// example:
	//
	// `curl Rj0F9uvI.popscan.xaliyun.com`
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The matched security baseline.
	//
	// example:
	//
	// Test
	MatchedSecurityBaseline *string `json:"MatchedSecurityBaseline,omitempty" xml:"MatchedSecurityBaseline,omitempty"`
	// The list of security baseline processes.
	Process []*string `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
	// The SSID.
	//
	// example:
	//
	// abcd
	Ssid *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetDevTag() *string {
	return s.DevTag
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetFile() []*string {
	return s.File
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetFirewall() *string {
	return s.Firewall
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetHostname() *string {
	return s.Hostname
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetInnerIp() *string {
	return s.InnerIp
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetMac() *string {
	return s.Mac
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetMatchedSecurityBaseline() *string {
	return s.MatchedSecurityBaseline
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetProcess() []*string {
	return s.Process
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) GetSsid() *string {
	return s.Ssid
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetDevTag(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.DevTag = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetDeviceType(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.DeviceType = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetFile(v []*string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.File = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetFirewall(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.Firewall = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetHostname(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.Hostname = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetInnerIp(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.InnerIp = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetInternetIp(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.InternetIp = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetMac(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.Mac = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetMatchedSecurityBaseline(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.MatchedSecurityBaseline = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetProcess(v []*string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.Process = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) SetSsid(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo {
	s.Ssid = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoDeviceAttributeInfo) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo struct {
	// The policy type.
	//
	// example:
	//
	// connector
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// av-rtd-091c2d6e3f24aae4
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// 1
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) GetRouteType() *string {
	return s.RouteType
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) GetStrategyName() *string {
	return s.StrategyName
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) SetRouteType(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo {
	s.RouteType = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) SetStrategyId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo {
	s.StrategyId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) SetStrategyName(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo {
	s.StrategyName = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoRouteStrategyInfo) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo struct {
	// The email address.
	//
	// example:
	//
	// 1234@xxxx.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The list of department names.
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// The matched user group.
	//
	// example:
	//
	// IT
	MatchedUserGroups *string `json:"MatchedUserGroups,omitempty" xml:"MatchedUserGroups,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 123456789
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GetEmail() *string {
	return s.Email
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GetGroup() []*string {
	return s.Group
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GetMatchedUserGroups() *string {
	return s.MatchedUserGroups
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GetTelephone() *string {
	return s.Telephone
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) GetUsername() *string {
	return s.Username
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) SetEmail(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	s.Email = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) SetGroup(v []*string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	s.Group = v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) SetMatchedUserGroups(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	s.MatchedUserGroups = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) SetTelephone(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	s.Telephone = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) SetUsername(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo {
	s.Username = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoUserGroupInfo) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo struct {
	// The policy action:
	//
	// - **Allow**: allow
	//
	// - **Block**: block
	//
	// - **Observe**: monitor mode
	//
	// example:
	//
	// block
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The application name.
	//
	// example:
	//
	// MyApp2
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The block information.
	//
	// example:
	//
	// access denied
	BlockInfo *string `json:"BlockInfo,omitempty" xml:"BlockInfo,omitempty"`
	// The name of the zero-trust policy.
	//
	// example:
	//
	// 保密测试
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) GetAction() *string {
	return s.Action
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) GetAppName() *string {
	return s.AppName
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) GetBlockInfo() *string {
	return s.BlockInfo
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) SetAction(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo {
	s.Action = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) SetAppName(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo {
	s.AppName = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) SetBlockInfo(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo {
	s.BlockInfo = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) SetPolicyName(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo {
	s.PolicyName = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskResultPolicyInfoZeroTrustPolicyInfo) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs struct {
	// The expected response from the backend server.
	//
	// example:
	//
	// hello
	ExpectedResponse *string `json:"ExpectedResponse,omitempty" xml:"ExpectedResponse,omitempty"`
	// The content of the UDP request.
	//
	// example:
	//
	// hello
	RequestContent *string `json:"RequestContent,omitempty" xml:"RequestContent,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GetExpectedResponse() *string {
	return s.ExpectedResponse
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) GetRequestContent() *string {
	return s.RequestContent
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) SetExpectedResponse(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	s.ExpectedResponse = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) SetRequestContent(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs {
	s.RequestContent = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUdpExtraConfigs) Validate() error {
	return dara.Validate(s)
}

type GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup struct {
	// The ID of the user group.
	//
	// example:
	//
	// ug-xxxxx
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// IT
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) String() string {
	return dara.Prettify(s)
}

func (s GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GoString() string {
	return s.String()
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) SetUserGroupId(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	s.UserGroupId = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) SetUserGroupName(v string) *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup {
	s.UserGroupName = &v
	return s
}

func (s *GetPADiagnosisTaskResponseBodyDiagnosisTaskUserGroup) Validate() error {
	return dara.Validate(s)
}
