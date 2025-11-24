// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSwimLaneGroupListResponseBody
	GetRequestId() *string
	SetSwimLaneGroupList(v []*GetSwimLaneGroupListResponseBodySwimLaneGroupList) *GetSwimLaneGroupListResponseBody
	GetSwimLaneGroupList() []*GetSwimLaneGroupListResponseBodySwimLaneGroupList
}

type GetSwimLaneGroupListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the lane group.
	SwimLaneGroupList []*GetSwimLaneGroupListResponseBodySwimLaneGroupList `json:"SwimLaneGroupList,omitempty" xml:"SwimLaneGroupList,omitempty" type:"Repeated"`
}

func (s GetSwimLaneGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSwimLaneGroupListResponseBody) GetSwimLaneGroupList() []*GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	return s.SwimLaneGroupList
}

func (s *GetSwimLaneGroupListResponseBody) SetRequestId(v string) *GetSwimLaneGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBody) SetSwimLaneGroupList(v []*GetSwimLaneGroupListResponseBodySwimLaneGroupList) *GetSwimLaneGroupListResponseBody {
	s.SwimLaneGroupList = v
	return s
}

func (s *GetSwimLaneGroupListResponseBody) Validate() error {
	if s.SwimLaneGroupList != nil {
		for _, item := range s.SwimLaneGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSwimLaneGroupListResponseBodySwimLaneGroupList struct {
	// The name of the baseline lane of the lane group in permissive mode. This parameter is valid only for a lane group in permissive mode.
	//
	// example:
	//
	// s1
	FallbackTarget *string `json:"FallbackTarget,omitempty" xml:"FallbackTarget,omitempty"`
	// The name of a lane group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the ASM ingress gateway.
	//
	// example:
	//
	// ingressgateway
	IngressGatewayName *string `json:"IngressGatewayName,omitempty" xml:"IngressGatewayName,omitempty"`
	// The policy used to route requests among multiple lanes in a lane group. Valid values:
	//
	// 	- weighted: a weight-based request routing policy. Requests are matched based on uniform rules and then routed to different lanes in a lane group at a specified ratio.
	//
	// 	- rule-based: a rule-based request routing policy. Each lane is configured with request routing rules. Only requests that match the request routing rules of a specific lane are routed to the lane.
	//
	// example:
	//
	// weighted
	IngressRoutingStrategy *string `json:"IngressRoutingStrategy,omitempty" xml:"IngressRoutingStrategy,omitempty"`
	// The type of gateways in which you can configure request routing rules. You can configure request routing rules only in ASM gateways.
	//
	// example:
	//
	// ASM
	IngressType *string `json:"IngressType,omitempty" xml:"IngressType,omitempty"`
	// Indicates whether the lane group is in permissive mode.
	//
	// example:
	//
	// false
	IsPermissive *bool `json:"IsPermissive,omitempty" xml:"IsPermissive,omitempty"`
	// The request routing header of the lane group. It is valid only for a lane group in permissive mode.
	//
	// example:
	//
	// x-asm-prefer-tag
	RouteHeader                *string `json:"RouteHeader,omitempty" xml:"RouteHeader,omitempty"`
	ServiceLevelFallbackTarget *string `json:"ServiceLevelFallbackTarget,omitempty" xml:"ServiceLevelFallbackTarget,omitempty"`
	// The Services associated with the lane group.
	//
	// example:
	//
	// ["sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc"]
	ServiceList *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
	// A serialized JSON string. The keys of the JSON object are the keys of the labels owned by all services in the lane group, and the values of the JSON object are arrays of all possible label values of the services in the lane group.
	//
	// example:
	//
	// {"ASM_TRAFFIC_TAG":["v2","v3","v1"],"version":["v3","v1","v2"]}
	SwimLaneLabels *string `json:"SwimLaneLabels,omitempty" xml:"SwimLaneLabels,omitempty"`
	// The end-to-end (E2E) pass-through request header of the lane group. It is valid only for a lane group in permissive mode.
	//
	// example:
	//
	// my-request-id
	TraceHeader *string `json:"TraceHeader,omitempty" xml:"TraceHeader,omitempty"`
	// The weight-based request routing rules for a lane group. This parameter is returned only when the IngressRoutingStrategy parameter is set to weighted.
	//
	// example:
	//
	// {"Domains":["*"],"MatchRequests":[{"URI":{"MatchingMode":"exact","MatchingContent":"/mock"},"Headers":[{"Name":"test","MatchingMode":"exact","MatchingContent":"yes"}]}]}
	WeightedIngressRule *string `json:"WeightedIngressRule,omitempty" xml:"WeightedIngressRule,omitempty"`
}

func (s GetSwimLaneGroupListResponseBodySwimLaneGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneGroupListResponseBodySwimLaneGroupList) GoString() string {
	return s.String()
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetFallbackTarget() *string {
	return s.FallbackTarget
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetIngressGatewayName() *string {
	return s.IngressGatewayName
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetIngressRoutingStrategy() *string {
	return s.IngressRoutingStrategy
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetIngressType() *string {
	return s.IngressType
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetIsPermissive() *bool {
	return s.IsPermissive
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetRouteHeader() *string {
	return s.RouteHeader
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetServiceLevelFallbackTarget() *string {
	return s.ServiceLevelFallbackTarget
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetServiceList() *string {
	return s.ServiceList
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetSwimLaneLabels() *string {
	return s.SwimLaneLabels
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetTraceHeader() *string {
	return s.TraceHeader
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) GetWeightedIngressRule() *string {
	return s.WeightedIngressRule
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetFallbackTarget(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.FallbackTarget = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetGroupName(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIngressGatewayName(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IngressGatewayName = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIngressRoutingStrategy(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IngressRoutingStrategy = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIngressType(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IngressType = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetIsPermissive(v bool) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.IsPermissive = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetRouteHeader(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.RouteHeader = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetServiceLevelFallbackTarget(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.ServiceLevelFallbackTarget = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetServiceList(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.ServiceList = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetSwimLaneLabels(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.SwimLaneLabels = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetTraceHeader(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.TraceHeader = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) SetWeightedIngressRule(v string) *GetSwimLaneGroupListResponseBodySwimLaneGroupList {
	s.WeightedIngressRule = &v
	return s
}

func (s *GetSwimLaneGroupListResponseBodySwimLaneGroupList) Validate() error {
	return dara.Validate(s)
}
