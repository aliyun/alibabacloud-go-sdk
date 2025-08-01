// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceKubernetesClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGovernanceKubernetesClusterResponseBodyData) *GetGovernanceKubernetesClusterResponseBody
	GetData() *GetGovernanceKubernetesClusterResponseBodyData
	SetMessage(v string) *GetGovernanceKubernetesClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGovernanceKubernetesClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGovernanceKubernetesClusterResponseBody
	GetSuccess() *bool
}

type GetGovernanceKubernetesClusterResponseBody struct {
	// The details of the data.
	Data *GetGovernanceKubernetesClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBody) GetData() *GetGovernanceKubernetesClusterResponseBodyData {
	return s.Data
}

func (s *GetGovernanceKubernetesClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGovernanceKubernetesClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGovernanceKubernetesClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetData(v *GetGovernanceKubernetesClusterResponseBodyData) *GetGovernanceKubernetesClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetMessage(v string) *GetGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *GetGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetSuccess(v bool) *GetGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGovernanceKubernetesClusterResponseBodyData struct {
	// The ID of the instance.
	//
	// example:
	//
	// cd23228b3c80c4d4f9ad7af1d87cc30d5
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// myCluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The version of Kubernetes.
	//
	// example:
	//
	// 1.20.11-aliyun.1
	K8sVersion *string `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	// The information of the namespace.
	//
	// example:
	//
	// [{\\"Name\\":\\"ack-onepilot\\",\\"Tags\\":null},{\\"Name\\":\\"default\\",\\"Tags\\":{\\"mse-enable\\":\\"enabled\\"}},{\\"Name\\":\\"kube-node-lease\\",\\"Tags\\":null},{\\"Name\\":\\"kube-public\\",\\"Tags\\":null},{\\"Name\\":\\"kube-system\\",\\"Tags\\":null}]
	NamespaceInfos *string `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	// The queried namespaces.
	Namespaces []*GetGovernanceKubernetesClusterResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The time when the pilot component was started.
	//
	// example:
	//
	// 2022-01-11T11:50:38.000+0000
	PilotStartTime *string `json:"PilotStartTime,omitempty" xml:"PilotStartTime,omitempty"`
	PilotVersion   *string `json:"PilotVersion,omitempty" xml:"PilotVersion,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by MSE.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time of the last modification.
	//
	// example:
	//
	// 2022-01-12T05:24:31.000+0000
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionLifeCycle *string `json:"VersionLifeCycle,omitempty" xml:"VersionLifeCycle,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetK8sVersion() *string {
	return s.K8sVersion
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetNamespaceInfos() *string {
	return s.NamespaceInfos
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetNamespaces() []*GetGovernanceKubernetesClusterResponseBodyDataNamespaces {
	return s.Namespaces
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetPilotStartTime() *string {
	return s.PilotStartTime
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetPilotVersion() *string {
	return s.PilotVersion
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) GetVersionLifeCycle() *string {
	return s.VersionLifeCycle
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetClusterId(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetClusterName(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetK8sVersion(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.K8sVersion = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetNamespaceInfos(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.NamespaceInfos = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetNamespaces(v []*GetGovernanceKubernetesClusterResponseBodyDataNamespaces) *GetGovernanceKubernetesClusterResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetPilotStartTime(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.PilotStartTime = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetPilotVersion(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.PilotVersion = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetRegion(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetUpdateTime(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetVersionLifeCycle(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.VersionLifeCycle = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetGovernanceKubernetesClusterResponseBodyDataNamespaces struct {
	// The name of the MSE namespace that you want to access.
	//
	// example:
	//
	// default
	MseNamespace *string `json:"MseNamespace,omitempty" xml:"MseNamespace,omitempty"`
	// The name of the namespace in the ACK cluster.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBodyDataNamespaces) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) GetMseNamespace() *string {
	return s.MseNamespace
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) GetName() *string {
	return s.Name
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) SetMseNamespace(v string) *GetGovernanceKubernetesClusterResponseBodyDataNamespaces {
	s.MseNamespace = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) SetName(v string) *GetGovernanceKubernetesClusterResponseBodyDataNamespaces {
	s.Name = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) Validate() error {
	return dara.Validate(s)
}
