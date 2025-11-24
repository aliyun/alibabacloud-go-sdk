// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFallbackTarget(v string) *UpdateSwimLaneGroupRequest
	GetFallbackTarget() *string
	SetGroupName(v string) *UpdateSwimLaneGroupRequest
	GetGroupName() *string
	SetIngressRoutingStrategy(v string) *UpdateSwimLaneGroupRequest
	GetIngressRoutingStrategy() *string
	SetServiceLevelFallbackTarget(v string) *UpdateSwimLaneGroupRequest
	GetServiceLevelFallbackTarget() *string
	SetServiceMeshId(v string) *UpdateSwimLaneGroupRequest
	GetServiceMeshId() *string
	SetServicesList(v string) *UpdateSwimLaneGroupRequest
	GetServicesList() *string
	SetWeightedIngressRule(v string) *UpdateSwimLaneGroupRequest
	GetWeightedIngressRule() *string
}

type UpdateSwimLaneGroupRequest struct {
	// The name of the baseline lane in the lane group if the lane group is in permissive mode. This parameter is valid only for a lane group in permissive mode.
	//
	// example:
	//
	// s1
	FallbackTarget *string `json:"FallbackTarget,omitempty" xml:"FallbackTarget,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The policy used to route requests among multiple lanes in a lane group. Valid values:
	//
	// 	- weighted: a weight-based request routing policy. Requests are matched based on uniform rules and then routed to different lanes in a lane group at a specified ratio.
	//
	// 	- rule-based: a rule-based request routing policy. Each lane is configured with request routing rules. Only requests that match the request routing rules of a specific lane are routed to the lane.
	//
	// example:
	//
	// weighted
	IngressRoutingStrategy     *string `json:"IngressRoutingStrategy,omitempty" xml:"IngressRoutingStrategy,omitempty"`
	ServiceLevelFallbackTarget *string `json:"ServiceLevelFallbackTarget,omitempty" xml:"ServiceLevelFallbackTarget,omitempty"`
	// The Service Mesh (ASM) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane group.
	//
	// example:
	//
	// ["sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc"]
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The weight-based request routing rules for a lane group. You can specify this parameter only when the IngressRoutingStrategy parameter is set to weighted.
	//
	// example:
	//
	// {"Domains":["*"],"MatchRequests":[{"URI":{"MatchingMode":"exact","MatchingContent":"/mock"},"Headers":[{"Name":"test","MatchingMode":"exact","MatchingContent":"yes"}]}]}
	WeightedIngressRule *string `json:"WeightedIngressRule,omitempty" xml:"WeightedIngressRule,omitempty"`
}

func (s UpdateSwimLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupRequest) GetFallbackTarget() *string {
	return s.FallbackTarget
}

func (s *UpdateSwimLaneGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateSwimLaneGroupRequest) GetIngressRoutingStrategy() *string {
	return s.IngressRoutingStrategy
}

func (s *UpdateSwimLaneGroupRequest) GetServiceLevelFallbackTarget() *string {
	return s.ServiceLevelFallbackTarget
}

func (s *UpdateSwimLaneGroupRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateSwimLaneGroupRequest) GetServicesList() *string {
	return s.ServicesList
}

func (s *UpdateSwimLaneGroupRequest) GetWeightedIngressRule() *string {
	return s.WeightedIngressRule
}

func (s *UpdateSwimLaneGroupRequest) SetFallbackTarget(v string) *UpdateSwimLaneGroupRequest {
	s.FallbackTarget = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetGroupName(v string) *UpdateSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetIngressRoutingStrategy(v string) *UpdateSwimLaneGroupRequest {
	s.IngressRoutingStrategy = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetServiceLevelFallbackTarget(v string) *UpdateSwimLaneGroupRequest {
	s.ServiceLevelFallbackTarget = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetServiceMeshId(v string) *UpdateSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetServicesList(v string) *UpdateSwimLaneGroupRequest {
	s.ServicesList = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) SetWeightedIngressRule(v string) *UpdateSwimLaneGroupRequest {
	s.WeightedIngressRule = &v
	return s
}

func (s *UpdateSwimLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
