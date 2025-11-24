// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIngressRule(v string) *GetSwimLaneDetailResponseBody
	GetIngressRule() *string
	SetIngressService(v string) *GetSwimLaneDetailResponseBody
	GetIngressService() *string
	SetLabelSelectorKey(v string) *GetSwimLaneDetailResponseBody
	GetLabelSelectorKey() *string
	SetLabelSelectorValue(v string) *GetSwimLaneDetailResponseBody
	GetLabelSelectorValue() *string
	SetRequestId(v string) *GetSwimLaneDetailResponseBody
	GetRequestId() *string
	SetServicesList(v string) *GetSwimLaneDetailResponseBody
	GetServicesList() *string
	SetValidationMessage(v string) *GetSwimLaneDetailResponseBody
	GetValidationMessage() *string
	SetWeightedIngressDestination(v string) *GetSwimLaneDetailResponseBody
	GetWeightedIngressDestination() *string
}

type GetSwimLaneDetailResponseBody struct {
	// The traffic routing rule that routes traffic to the lane by using the ingress gateway. The traffic routing rule contains one or more custom routes.
	//
	// example:
	//
	// [{"Domains":["*"],"RouteName":"r1","MatchRequest":{"Headers":[{"Name":"x-asm-prefer-tag","MatchingMode":"exact","MatchingContent":"s1"}],"URI":{"MatchingMode":"exact","MatchingContent":"/mock"}},"RouteDestinations":[{"Destination":{"Host":"mocka.default.svc.cluster.local","Subset":"s1"}}]},{"Domains":["*"],"RouteName":"hello","MatchRequest":{"Headers":[{"Name":"x-asm-prefer-tag","MatchingMode":"exact","MatchingContent":"s1"}],"URI":{"MatchingMode":"exact","MatchingContent":"/mocktest"}},"RouteDestinations":[{"Destination":{"Host":"mocka.default.svc.cluster.local","Subset":"s1"}}]}]
	IngressRule *string `json:"IngressRule,omitempty" xml:"IngressRule,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// mocka.default.svc.cluster.local
	IngressService *string `json:"IngressService,omitempty" xml:"IngressService,omitempty"`
	// The label key of the associated service workload. The value is fixed as **ASM_TRAFFIC_TAG**.
	//
	// example:
	//
	// ASM_TRAFFIC_TAG
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The value of ASM_TRAFFIC_TAG.
	//
	// example:
	//
	// v1
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of services associated with the lane.
	//
	// example:
	//
	// ["sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc"]
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The status of the swimlane. If the configuration is successful and takes effect, a `null` is returned. Otherwise, an error message is returned.
	//
	// example:
	//
	// [{"code":"CODE_UNEFFECTED_SWIMLANE_LABEL","level":"warning","message":"The label selector defined in the ASMSwimLane has no effect on any workload instance."}]
	ValidationMessage *string `json:"ValidationMessage,omitempty" xml:"ValidationMessage,omitempty"`
	// This parameter is returned only when the IngressRoutingStrategy parameter is set to weighted. This parameter indicates the domain name of Services in each lane and the request routing weight. The value of this parameter is a serialized JSON string.
	//
	// example:
	//
	// {"RouteDestination":{"Host":"mocka.default.svc.cluster.local","Subset":"s1"},"Weight":40}
	WeightedIngressDestination *string `json:"WeightedIngressDestination,omitempty" xml:"WeightedIngressDestination,omitempty"`
}

func (s GetSwimLaneDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneDetailResponseBody) GetIngressRule() *string {
	return s.IngressRule
}

func (s *GetSwimLaneDetailResponseBody) GetIngressService() *string {
	return s.IngressService
}

func (s *GetSwimLaneDetailResponseBody) GetLabelSelectorKey() *string {
	return s.LabelSelectorKey
}

func (s *GetSwimLaneDetailResponseBody) GetLabelSelectorValue() *string {
	return s.LabelSelectorValue
}

func (s *GetSwimLaneDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSwimLaneDetailResponseBody) GetServicesList() *string {
	return s.ServicesList
}

func (s *GetSwimLaneDetailResponseBody) GetValidationMessage() *string {
	return s.ValidationMessage
}

func (s *GetSwimLaneDetailResponseBody) GetWeightedIngressDestination() *string {
	return s.WeightedIngressDestination
}

func (s *GetSwimLaneDetailResponseBody) SetIngressRule(v string) *GetSwimLaneDetailResponseBody {
	s.IngressRule = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetIngressService(v string) *GetSwimLaneDetailResponseBody {
	s.IngressService = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetLabelSelectorKey(v string) *GetSwimLaneDetailResponseBody {
	s.LabelSelectorKey = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetLabelSelectorValue(v string) *GetSwimLaneDetailResponseBody {
	s.LabelSelectorValue = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetRequestId(v string) *GetSwimLaneDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetServicesList(v string) *GetSwimLaneDetailResponseBody {
	s.ServicesList = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetValidationMessage(v string) *GetSwimLaneDetailResponseBody {
	s.ValidationMessage = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) SetWeightedIngressDestination(v string) *GetSwimLaneDetailResponseBody {
	s.WeightedIngressDestination = &v
	return s
}

func (s *GetSwimLaneDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
