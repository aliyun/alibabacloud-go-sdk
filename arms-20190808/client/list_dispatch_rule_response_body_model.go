// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDispatchRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRules(v []*ListDispatchRuleResponseBodyDispatchRules) *ListDispatchRuleResponseBody
	GetDispatchRules() []*ListDispatchRuleResponseBodyDispatchRules
	SetRequestId(v string) *ListDispatchRuleResponseBody
	GetRequestId() *string
}

type ListDispatchRuleResponseBody struct {
	// The returned struct.
	DispatchRules []*ListDispatchRuleResponseBodyDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDispatchRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBody) GetDispatchRules() []*ListDispatchRuleResponseBodyDispatchRules {
	return s.DispatchRules
}

func (s *ListDispatchRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDispatchRuleResponseBody) SetDispatchRules(v []*ListDispatchRuleResponseBodyDispatchRules) *ListDispatchRuleResponseBody {
	s.DispatchRules = v
	return s
}

func (s *ListDispatchRuleResponseBody) SetRequestId(v string) *ListDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDispatchRuleResponseBody) Validate() error {
	if s.DispatchRules != nil {
		for _, item := range s.DispatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDispatchRuleResponseBodyDispatchRules struct {
	// The name of the notification policy.
	//
	// example:
	//
	// Prod
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 10282
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Indicates whether the notification policy is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDispatchRuleResponseBodyDispatchRules) String() string {
	return dara.Prettify(s)
}

func (s ListDispatchRuleResponseBodyDispatchRules) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBodyDispatchRules) GetName() *string {
	return s.Name
}

func (s *ListDispatchRuleResponseBodyDispatchRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListDispatchRuleResponseBodyDispatchRules) GetState() *string {
	return s.State
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetName(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetRuleId(v int64) *ListDispatchRuleResponseBodyDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetState(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.State = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) Validate() error {
	return dara.Validate(s)
}
