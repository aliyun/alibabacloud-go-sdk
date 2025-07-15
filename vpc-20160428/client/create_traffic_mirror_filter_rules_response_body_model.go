// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressRules(v []*CreateTrafficMirrorFilterRulesResponseBodyEgressRules) *CreateTrafficMirrorFilterRulesResponseBody
	GetEgressRules() []*CreateTrafficMirrorFilterRulesResponseBodyEgressRules
	SetIngressRules(v []*CreateTrafficMirrorFilterRulesResponseBodyIngressRules) *CreateTrafficMirrorFilterRulesResponseBody
	GetIngressRules() []*CreateTrafficMirrorFilterRulesResponseBodyIngressRules
	SetRequestId(v string) *CreateTrafficMirrorFilterRulesResponseBody
	GetRequestId() *string
}

type CreateTrafficMirrorFilterRulesResponseBody struct {
	// The list of outbound rules.
	EgressRules []*CreateTrafficMirrorFilterRulesResponseBodyEgressRules `json:"EgressRules,omitempty" xml:"EgressRules,omitempty" type:"Repeated"`
	// The list of inbound rules.
	IngressRules []*CreateTrafficMirrorFilterRulesResponseBodyIngressRules `json:"IngressRules,omitempty" xml:"IngressRules,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 07F272E2-6AD5-433A-8207-A607C76F1676
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) GetEgressRules() []*CreateTrafficMirrorFilterRulesResponseBodyEgressRules {
	return s.EgressRules
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) GetIngressRules() []*CreateTrafficMirrorFilterRulesResponseBodyIngressRules {
	return s.IngressRules
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) SetEgressRules(v []*CreateTrafficMirrorFilterRulesResponseBodyEgressRules) *CreateTrafficMirrorFilterRulesResponseBody {
	s.EgressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) SetIngressRules(v []*CreateTrafficMirrorFilterRulesResponseBodyIngressRules) *CreateTrafficMirrorFilterRulesResponseBody {
	s.IngressRules = v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) SetRequestId(v string) *CreateTrafficMirrorFilterRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRulesResponseBodyEgressRules struct {
	// The ID of the outbound rule.
	//
	// example:
	//
	// tmr-j6cok23ugp53eeib5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesResponseBodyEgressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesResponseBodyEgressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyEgressRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyEgressRules) SetInstanceId(v string) *CreateTrafficMirrorFilterRulesResponseBodyEgressRules {
	s.InstanceId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyEgressRules) Validate() error {
	return dara.Validate(s)
}

type CreateTrafficMirrorFilterRulesResponseBodyIngressRules struct {
	// The ID of the inbound rule.
	//
	// example:
	//
	// tmr-j6c6rtallo51ouzv3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateTrafficMirrorFilterRulesResponseBodyIngressRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterRulesResponseBodyIngressRules) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyIngressRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyIngressRules) SetInstanceId(v string) *CreateTrafficMirrorFilterRulesResponseBodyIngressRules {
	s.InstanceId = &v
	return s
}

func (s *CreateTrafficMirrorFilterRulesResponseBodyIngressRules) Validate() error {
	return dara.Validate(s)
}
