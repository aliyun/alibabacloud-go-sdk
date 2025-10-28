// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sIngressRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListK8sIngressRulesRequest
	GetClusterId() *string
	SetCondition(v string) *ListK8sIngressRulesRequest
	GetCondition() *string
	SetNamespace(v string) *ListK8sIngressRulesRequest
	GetNamespace() *string
	SetRegionId(v string) *ListK8sIngressRulesRequest
	GetRegionId() *string
}

type ListK8sIngressRulesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// 5b2b4ab4-efbc-4a81-9c45-xxxxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The filter conditions. Set the value to a JSON string in the format of {"field":"Name", "pattern":"my-"}, where:
	//
	// 	- field: the parameter to be matched. Valid values: Name and ClusterName.
	//
	// 	- pattern: the content to be matched.
	//
	// For example, a value of {"field":"Name", "pattern":"my-"} indicates that the specified filter conditions match the routing rules whose names start with my-.
	//
	// example:
	//
	// {"field":"Name", "pattern":"my-"}
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region where the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListK8sIngressRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesRequest) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sIngressRulesRequest) GetCondition() *string {
	return s.Condition
}

func (s *ListK8sIngressRulesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sIngressRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListK8sIngressRulesRequest) SetClusterId(v string) *ListK8sIngressRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListK8sIngressRulesRequest) SetCondition(v string) *ListK8sIngressRulesRequest {
	s.Condition = &v
	return s
}

func (s *ListK8sIngressRulesRequest) SetNamespace(v string) *ListK8sIngressRulesRequest {
	s.Namespace = &v
	return s
}

func (s *ListK8sIngressRulesRequest) SetRegionId(v string) *ListK8sIngressRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListK8sIngressRulesRequest) Validate() error {
	return dara.Validate(s)
}
