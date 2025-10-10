// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPolicyRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSecurityPolicyRelationsResponseBody
	GetRequestId() *string
	SetSecrityPolicyRelations(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) *ListSecurityPolicyRelationsResponseBody
	GetSecrityPolicyRelations() []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations
}

type ListSecurityPolicyRelationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security policies and the listeners that are associated with the security policies.
	SecrityPolicyRelations []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations `json:"SecrityPolicyRelations,omitempty" xml:"SecrityPolicyRelations,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityPolicyRelationsResponseBody) GetSecrityPolicyRelations() []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	return s.SecrityPolicyRelations
}

func (s *ListSecurityPolicyRelationsResponseBody) SetRequestId(v string) *ListSecurityPolicyRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBody) SetSecrityPolicyRelations(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) *ListSecurityPolicyRelationsResponseBody {
	s.SecrityPolicyRelations = v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations struct {
	// The listeners that are associated with the security policy.
	RelatedListeners []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	// The security policy ID.
	//
	// example:
	//
	// scp-bp1bpn0kn9****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) GetRelatedListeners() []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	return s.RelatedListeners
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetRelatedListeners(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.RelatedListeners = v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetSecurityPolicyId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) Validate() error {
	return dara.Validate(s)
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners struct {
	// The listener ID.
	//
	// example:
	//
	// lsn-0bfuc****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol.
	//
	// example:
	//
	// HTTPS
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The Server Load Balancer (SLB) instance ID.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GetListenerPort() *int64 {
	return s.ListenerPort
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerPort(v int64) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerProtocol(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetLoadBalancerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) Validate() error {
	return dara.Validate(s)
}
