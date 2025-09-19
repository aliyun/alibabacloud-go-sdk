// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterPluginInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListClusterPluginInfoResponseBodyData) *ListClusterPluginInfoResponseBody
	GetData() []*ListClusterPluginInfoResponseBodyData
	SetRequestId(v string) *ListClusterPluginInfoResponseBody
	GetRequestId() *string
}

type ListClusterPluginInfoResponseBody struct {
	// The information about the plug-in.
	Data []*ListClusterPluginInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterPluginInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterPluginInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterPluginInfoResponseBody) GetData() []*ListClusterPluginInfoResponseBodyData {
	return s.Data
}

func (s *ListClusterPluginInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterPluginInfoResponseBody) SetData(v []*ListClusterPluginInfoResponseBodyData) *ListClusterPluginInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterPluginInfoResponseBody) SetRequestId(v string) *ListClusterPluginInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterPluginInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterPluginInfoResponseBodyData struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// lmftest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- 1: normal
	//
	// 	- 2: abnormal
	//
	// 	- 3: offline
	//
	// example:
	//
	// ABNORMAL
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The plug-ins.
	NodePluginInfoList []*ListClusterPluginInfoResponseBodyDataNodePluginInfoList `json:"NodePluginInfoList,omitempty" xml:"NodePluginInfoList,omitempty" type:"Repeated"`
}

func (s ListClusterPluginInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterPluginInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterPluginInfoResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterPluginInfoResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterPluginInfoResponseBodyData) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *ListClusterPluginInfoResponseBodyData) GetNodePluginInfoList() []*ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	return s.NodePluginInfoList
}

func (s *ListClusterPluginInfoResponseBodyData) SetClusterId(v string) *ListClusterPluginInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyData) SetClusterName(v string) *ListClusterPluginInfoResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyData) SetClusterStatus(v string) *ListClusterPluginInfoResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyData) SetNodePluginInfoList(v []*ListClusterPluginInfoResponseBodyDataNodePluginInfoList) *ListClusterPluginInfoResponseBodyData {
	s.NodePluginInfoList = v
	return s
}

func (s *ListClusterPluginInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListClusterPluginInfoResponseBodyDataNodePluginInfoList struct {
	// The error code returned.
	//
	// example:
	//
	// kenerl not support
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// kenerl not support
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the plug-in is installed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Installed *bool `json:"Installed,omitempty" xml:"Installed,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 100.100.XXX.XX
	MachineInternetIp *string `json:"MachineInternetIp,omitempty" xml:"MachineInternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 10.XXX.XXX.XX
	MachineIntranetIp *string `json:"MachineIntranetIp,omitempty" xml:"MachineIntranetIp,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// npznas05
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The type of the instance. Valid values include:
	//
	// 	- **ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **slb**: Server Load Balancer (SLB) instance
	//
	// example:
	//
	// ECS
	MachineType *int64 `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Indicates whether the Security Center agent is online. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If the Security Center agent of the server is offline, Security Center does not protect the server.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// alihips
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// 1.3.1
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// bc563d2b-2a3d-411b-8bbe-d75b8d3c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// tpp-cn-2r42njq4y001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListClusterPluginInfoResponseBodyDataNodePluginInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GoString() string {
	return s.String()
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetInstalled() *bool {
	return s.Installed
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetMachineInternetIp() *string {
	return s.MachineInternetIp
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetMachineIntranetIp() *string {
	return s.MachineIntranetIp
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetMachineName() *string {
	return s.MachineName
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetMachineType() *int64 {
	return s.MachineType
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetOnline() *bool {
	return s.Online
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetPluginName() *string {
	return s.PluginName
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetUuid() *string {
	return s.Uuid
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetErrorCode(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetErrorMsg(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.ErrorMsg = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetInstalled(v bool) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.Installed = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetMachineInternetIp(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.MachineInternetIp = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetMachineIntranetIp(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.MachineIntranetIp = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetMachineName(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.MachineName = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetMachineType(v int64) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.MachineType = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetOnline(v bool) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.Online = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetPluginName(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.PluginName = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetPluginVersion(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.PluginVersion = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetUuid(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.Uuid = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) SetInstanceId(v string) *ListClusterPluginInfoResponseBodyDataNodePluginInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListClusterPluginInfoResponseBodyDataNodePluginInfoList) Validate() error {
	return dara.Validate(s)
}
