// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutePolicy interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *RoutePolicy
	GetCondition() *string
	SetPolicyItems(v []*PolicyItem) *RoutePolicy
	GetPolicyItems() []*PolicyItem
}

type RoutePolicy struct {
	Condition   *string       `json:"condition,omitempty" xml:"condition,omitempty"`
	PolicyItems []*PolicyItem `json:"policyItems,omitempty" xml:"policyItems,omitempty" type:"Repeated"`
}

func (s RoutePolicy) String() string {
	return dara.Prettify(s)
}

func (s RoutePolicy) GoString() string {
	return s.String()
}

func (s *RoutePolicy) GetCondition() *string {
	return s.Condition
}

func (s *RoutePolicy) GetPolicyItems() []*PolicyItem {
	return s.PolicyItems
}

func (s *RoutePolicy) SetCondition(v string) *RoutePolicy {
	s.Condition = &v
	return s
}

func (s *RoutePolicy) SetPolicyItems(v []*PolicyItem) *RoutePolicy {
	s.PolicyItems = v
	return s
}

func (s *RoutePolicy) Validate() error {
	if s.PolicyItems != nil {
		for _, item := range s.PolicyItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
