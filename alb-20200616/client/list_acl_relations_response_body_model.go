// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclRelations(v []*ListAclRelationsResponseBodyAclRelations) *ListAclRelationsResponseBody
	GetAclRelations() []*ListAclRelationsResponseBodyAclRelations
	SetRequestId(v string) *ListAclRelationsResponseBody
	GetRequestId() *string
}

type ListAclRelationsResponseBody struct {
	// The relations between the specified ACL and the listeners.
	AclRelations []*ListAclRelationsResponseBodyAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAclRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAclRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBody) GetAclRelations() []*ListAclRelationsResponseBodyAclRelations {
	return s.AclRelations
}

func (s *ListAclRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAclRelationsResponseBody) SetAclRelations(v []*ListAclRelationsResponseBodyAclRelations) *ListAclRelationsResponseBody {
	s.AclRelations = v
	return s
}

func (s *ListAclRelationsResponseBody) SetRequestId(v string) *ListAclRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclRelationsResponseBody) Validate() error {
	if s.AclRelations != nil {
		for _, item := range s.AclRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclRelationsResponseBodyAclRelations struct {
	// ACL ID
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The listeners that are associated with the ACL.
	RelatedListeners []*ListAclRelationsResponseBodyAclRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
}

func (s ListAclRelationsResponseBodyAclRelations) String() string {
	return dara.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelations) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelations) GetAclId() *string {
	return s.AclId
}

func (s *ListAclRelationsResponseBodyAclRelations) GetRelatedListeners() []*ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	return s.RelatedListeners
}

func (s *ListAclRelationsResponseBodyAclRelations) SetAclId(v string) *ListAclRelationsResponseBodyAclRelations {
	s.AclId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelations) SetRelatedListeners(v []*ListAclRelationsResponseBodyAclRelationsRelatedListeners) *ListAclRelationsResponseBodyAclRelations {
	s.RelatedListeners = v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelations) Validate() error {
	if s.RelatedListeners != nil {
		for _, item := range s.RelatedListeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclRelationsResponseBodyAclRelationsRelatedListeners struct {
	// The listener ID.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol.
	//
	// example:
	//
	// HTTPS
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The association status between the ACL and the listener.
	//
	// 	- **Associating**
	//
	// 	- **Associated**
	//
	// 	- **Dissociating**
	//
	// example:
	//
	// Associated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) String() string {
	return dara.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) GetStatus() *string {
	return s.Status
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerPort(v int32) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerProtocol(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetLoadBalancerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetStatus(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.Status = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) Validate() error {
	return dara.Validate(s)
}
