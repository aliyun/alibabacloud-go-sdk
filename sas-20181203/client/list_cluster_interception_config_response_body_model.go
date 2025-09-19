// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInterceptionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfigList(v []*ListClusterInterceptionConfigResponseBodyClusterConfigList) *ListClusterInterceptionConfigResponseBody
	GetClusterConfigList() []*ListClusterInterceptionConfigResponseBodyClusterConfigList
	SetPageInfo(v *ListClusterInterceptionConfigResponseBodyPageInfo) *ListClusterInterceptionConfigResponseBody
	GetPageInfo() *ListClusterInterceptionConfigResponseBodyPageInfo
	SetRequestId(v string) *ListClusterInterceptionConfigResponseBody
	GetRequestId() *string
}

type ListClusterInterceptionConfigResponseBody struct {
	// An array that consists of the configurations of the cluster.
	ClusterConfigList []*ListClusterInterceptionConfigResponseBodyClusterConfigList `json:"ClusterConfigList,omitempty" xml:"ClusterConfigList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListClusterInterceptionConfigResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 49FDE92F-A0B8-56CC-B7A8-23B17646****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterInterceptionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInterceptionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterInterceptionConfigResponseBody) GetClusterConfigList() []*ListClusterInterceptionConfigResponseBodyClusterConfigList {
	return s.ClusterConfigList
}

func (s *ListClusterInterceptionConfigResponseBody) GetPageInfo() *ListClusterInterceptionConfigResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListClusterInterceptionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterInterceptionConfigResponseBody) SetClusterConfigList(v []*ListClusterInterceptionConfigResponseBodyClusterConfigList) *ListClusterInterceptionConfigResponseBody {
	s.ClusterConfigList = v
	return s
}

func (s *ListClusterInterceptionConfigResponseBody) SetPageInfo(v *ListClusterInterceptionConfigResponseBodyPageInfo) *ListClusterInterceptionConfigResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListClusterInterceptionConfigResponseBody) SetRequestId(v string) *ListClusterInterceptionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterInterceptionConfigResponseBodyClusterConfigList struct {
	// The status of the container firewall feature. Valid values:
	//
	// 	- **-1**: unknown
	//
	// 	- **0**: abnormal
	//
	// 	- **1**: normal
	//
	// 	- **2**: normal to be confirmed
	//
	// example:
	//
	// 0
	ClusterCNNFStatus *int32 `json:"ClusterCNNFStatus,omitempty" xml:"ClusterCNNFStatus,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c9051d30d8a044b4d99e1cb5d25ac****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// container-opa-kill-02
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **ManagedKubernetes**: managed Kubernetes cluster
	//
	// 	- **NotManagedKubernetes**: non-managed Kubernetes cluster
	//
	// 	- **PrivateKubernetes**: private cluster
	//
	// 	- **kubernetes**: dedicated Kubernetes cluster
	//
	// 	- **ask**: dedicated serverless Kubernetes (ASK) cluster
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The status of the defense rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	InterceptionSwitch *int32 `json:"InterceptionSwitch,omitempty" xml:"InterceptionSwitch,omitempty"`
	// The number of defense rules that are in effect.
	//
	// example:
	//
	// 12
	OpenRuleCount *int64 `json:"OpenRuleCount,omitempty" xml:"OpenRuleCount,omitempty"`
	// Indicates whether the container firewall feature is supported.
	//
	// example:
	//
	// false
	SupportCNNF *bool `json:"SupportCNNF,omitempty" xml:"SupportCNNF,omitempty"`
	// The total number of defense rules.
	//
	// example:
	//
	// 123
	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" xml:"TotalRuleCount,omitempty"`
}

func (s ListClusterInterceptionConfigResponseBodyClusterConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInterceptionConfigResponseBodyClusterConfigList) GoString() string {
	return s.String()
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetClusterCNNFStatus() *int32 {
	return s.ClusterCNNFStatus
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetInterceptionSwitch() *int32 {
	return s.InterceptionSwitch
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetOpenRuleCount() *int64 {
	return s.OpenRuleCount
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetSupportCNNF() *bool {
	return s.SupportCNNF
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) GetTotalRuleCount() *int64 {
	return s.TotalRuleCount
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetClusterCNNFStatus(v int32) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.ClusterCNNFStatus = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetClusterId(v string) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.ClusterId = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetClusterName(v string) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.ClusterName = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetClusterType(v string) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetInterceptionSwitch(v int32) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.InterceptionSwitch = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetOpenRuleCount(v int64) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.OpenRuleCount = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetSupportCNNF(v bool) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.SupportCNNF = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) SetTotalRuleCount(v int64) *ListClusterInterceptionConfigResponseBodyClusterConfigList {
	s.TotalRuleCount = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyClusterConfigList) Validate() error {
	return dara.Validate(s)
}

type ListClusterInterceptionConfigResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 11
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrrentPage *int32 `json:"CurrrentPage,omitempty" xml:"CurrrentPage,omitempty"`
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
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterInterceptionConfigResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInterceptionConfigResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) GetCurrrentPage() *int32 {
	return s.CurrrentPage
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) SetCount(v int32) *ListClusterInterceptionConfigResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) SetCurrrentPage(v int32) *ListClusterInterceptionConfigResponseBodyPageInfo {
	s.CurrrentPage = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) SetPageSize(v int32) *ListClusterInterceptionConfigResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) SetTotalCount(v int32) *ListClusterInterceptionConfigResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListClusterInterceptionConfigResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
