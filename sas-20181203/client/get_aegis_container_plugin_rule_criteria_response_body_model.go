// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) *GetAegisContainerPluginRuleCriteriaResponseBody
	GetCriteriaList() []*GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *GetAegisContainerPluginRuleCriteriaResponseBody
	GetRequestId() *string
}

type GetAegisContainerPluginRuleCriteriaResponseBody struct {
	// The information about the search condition.
	CriteriaList []*GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DA8133CC-CCA0-5CF2-BF64-FE7D52C44***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAegisContainerPluginRuleCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBody) GetCriteriaList() []*GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBody) SetCriteriaList(v []*GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) *GetAegisContainerPluginRuleCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBody) SetRequestId(v string) *GetAegisContainerPluginRuleCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition. Valid values:
	//
	// 	- **instanceId**: the ID of the container instance.
	//
	// 	- **clusterId**: the cluster ID.
	//
	// 	- **regionId**: the ID of the region in which the container resides.
	//
	// 	- **clusterName**: the name of the cluster.
	//
	// 	- **clusterType**: the type of the cluster.
	//
	// 	- **hostIp**: the public IP address.
	//
	// 	- **pod**: the pod.
	//
	// 	- **podIp**: the IP address of the pod.
	//
	// 	- **containerId**: the container ID.
	//
	// 	- **containerScope**: the type of the container.
	//
	// example:
	//
	// containerScope
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: The search condition needs to be specified.
	//
	// 	- **select**: The search condition is an option that can be selected from the drop-down list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values.
	//
	// example:
	//
	// NO,YES
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) SetName(v string) *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) SetType(v string) *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) SetValues(v string) *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
