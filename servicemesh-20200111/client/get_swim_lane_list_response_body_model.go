// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwimLaneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSwimLaneListResponseBody
	GetRequestId() *string
	SetSwimLaneList(v []*GetSwimLaneListResponseBodySwimLaneList) *GetSwimLaneListResponseBody
	GetSwimLaneList() []*GetSwimLaneListResponseBodySwimLaneList
}

type GetSwimLaneListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The swimlanes.
	SwimLaneList []*GetSwimLaneListResponseBodySwimLaneList `json:"SwimLaneList,omitempty" xml:"SwimLaneList,omitempty" type:"Repeated"`
}

func (s GetSwimLaneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSwimLaneListResponseBody) GetSwimLaneList() []*GetSwimLaneListResponseBodySwimLaneList {
	return s.SwimLaneList
}

func (s *GetSwimLaneListResponseBody) SetRequestId(v string) *GetSwimLaneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwimLaneListResponseBody) SetSwimLaneList(v []*GetSwimLaneListResponseBodySwimLaneList) *GetSwimLaneListResponseBody {
	s.SwimLaneList = v
	return s
}

func (s *GetSwimLaneListResponseBody) Validate() error {
	if s.SwimLaneList != nil {
		for _, item := range s.SwimLaneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSwimLaneListResponseBodySwimLaneList struct {
	// The name of a lane group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request routing rules.
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
	// The label key of the associated service workload. The value is fixed as `ASM_TRAFFIC_TAG`.
	//
	// example:
	//
	// ASM_TRAFFIC_TAG
	LabelSelectorKey *string `json:"LabelSelectorKey,omitempty" xml:"LabelSelectorKey,omitempty"`
	// The label value of the associated service workload. The value is fixed as `ASM_TRAFFIC_TAG`.
	//
	// example:
	//
	// v1
	LabelSelectorValue *string `json:"LabelSelectorValue,omitempty" xml:"LabelSelectorValue,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// s1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Services associated with the lane.
	//
	// example:
	//
	// ["sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb","sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc"]
	ServiceList *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
	// The verification messages of the lane group. If the service does not exist in the lane group, the verification message is displayed in the verification messages of the lane group.
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
	WeightedIngressDestinatin *string `json:"WeightedIngressDestinatin,omitempty" xml:"WeightedIngressDestinatin,omitempty"`
}

func (s GetSwimLaneListResponseBodySwimLaneList) String() string {
	return dara.Prettify(s)
}

func (s GetSwimLaneListResponseBodySwimLaneList) GoString() string {
	return s.String()
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetIngressRule() *string {
	return s.IngressRule
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetIngressService() *string {
	return s.IngressService
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetLabelSelectorKey() *string {
	return s.LabelSelectorKey
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetLabelSelectorValue() *string {
	return s.LabelSelectorValue
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetName() *string {
	return s.Name
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetServiceList() *string {
	return s.ServiceList
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetValidationMessage() *string {
	return s.ValidationMessage
}

func (s *GetSwimLaneListResponseBodySwimLaneList) GetWeightedIngressDestinatin() *string {
	return s.WeightedIngressDestinatin
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetGroupName(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.GroupName = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetIngressRule(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.IngressRule = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetIngressService(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.IngressService = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetLabelSelectorKey(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.LabelSelectorKey = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetLabelSelectorValue(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.LabelSelectorValue = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetName(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.Name = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetServiceList(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.ServiceList = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetValidationMessage(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.ValidationMessage = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) SetWeightedIngressDestinatin(v string) *GetSwimLaneListResponseBodySwimLaneList {
	s.WeightedIngressDestinatin = &v
	return s
}

func (s *GetSwimLaneListResponseBodySwimLaneList) Validate() error {
	return dara.Validate(s)
}
