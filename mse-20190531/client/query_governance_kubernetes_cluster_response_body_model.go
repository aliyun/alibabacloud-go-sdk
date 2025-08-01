// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGovernanceKubernetesClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryGovernanceKubernetesClusterResponseBodyData) *QueryGovernanceKubernetesClusterResponseBody
	GetData() *QueryGovernanceKubernetesClusterResponseBodyData
	SetMessage(v string) *QueryGovernanceKubernetesClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGovernanceKubernetesClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGovernanceKubernetesClusterResponseBody
	GetSuccess() *bool
}

type QueryGovernanceKubernetesClusterResponseBody struct {
	// The data returned.
	Data *QueryGovernanceKubernetesClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE34D413-2B79-5E77-9696-36D875E822AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGovernanceKubernetesClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGovernanceKubernetesClusterResponseBody) GetData() *QueryGovernanceKubernetesClusterResponseBodyData {
	return s.Data
}

func (s *QueryGovernanceKubernetesClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGovernanceKubernetesClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGovernanceKubernetesClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGovernanceKubernetesClusterResponseBody) SetData(v *QueryGovernanceKubernetesClusterResponseBodyData) *QueryGovernanceKubernetesClusterResponseBody {
	s.Data = v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBody) SetMessage(v string) *QueryGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *QueryGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBody) SetSuccess(v bool) *QueryGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryGovernanceKubernetesClusterResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of the data.
	Result []*QueryGovernanceKubernetesClusterResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of clusters.
	//
	// example:
	//
	// 3
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QueryGovernanceKubernetesClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryGovernanceKubernetesClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) GetResult() []*QueryGovernanceKubernetesClusterResponseBodyDataResult {
	return s.Result
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) SetPageNumber(v int32) *QueryGovernanceKubernetesClusterResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) SetPageSize(v int32) *QueryGovernanceKubernetesClusterResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) SetResult(v []*QueryGovernanceKubernetesClusterResponseBodyDataResult) *QueryGovernanceKubernetesClusterResponseBodyData {
	s.Result = v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) SetTotalSize(v int32) *QueryGovernanceKubernetesClusterResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryGovernanceKubernetesClusterResponseBodyDataResult struct {
	// The ID of the cluster.
	//
	// example:
	//
	// abcdef123456789
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// example-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.22
	K8sVersion *string `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	// The information about the namespace.
	//
	// example:
	//
	// [{"Name":"ack-onepilot","Tags":{"name":"ack-onepilot"}}]
	NamespaceInfos *string `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	// The time when the pilot component was started.
	//
	// example:
	//
	// 2022-05-17T05:39:43.000+0000
	PilotStartTime *string `json:"PilotStartTime,omitempty" xml:"PilotStartTime,omitempty"`
	PilotVersion   *string `json:"PilotVersion,omitempty" xml:"PilotVersion,omitempty"`
	// The region where the cluster resides.
	//
	// example:
	//
	// cn-shanghai
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	VersionLifeCycle *string `json:"VersionLifeCycle,omitempty" xml:"VersionLifeCycle,omitempty"`
}

func (s QueryGovernanceKubernetesClusterResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s QueryGovernanceKubernetesClusterResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetK8sVersion() *string {
	return s.K8sVersion
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetNamespaceInfos() *string {
	return s.NamespaceInfos
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetPilotStartTime() *string {
	return s.PilotStartTime
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetPilotVersion() *string {
	return s.PilotVersion
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetRegion() *string {
	return s.Region
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) GetVersionLifeCycle() *string {
	return s.VersionLifeCycle
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetClusterId(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.ClusterId = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetClusterName(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.ClusterName = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetK8sVersion(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.K8sVersion = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetNamespaceInfos(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.NamespaceInfos = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetPilotStartTime(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.PilotStartTime = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetPilotVersion(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.PilotVersion = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetRegion(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.Region = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) SetVersionLifeCycle(v string) *QueryGovernanceKubernetesClusterResponseBodyDataResult {
	s.VersionLifeCycle = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
