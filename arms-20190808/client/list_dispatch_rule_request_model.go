// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListDispatchRuleRequest
	GetName() *string
	SetRegionId(v string) *ListDispatchRuleRequest
	GetRegionId() *string
	SetSystem(v bool) *ListDispatchRuleRequest
	GetSystem() *bool
}

type ListDispatchRuleRequest struct {
	// The name of the notification policy. Fuzzy match is supported.
	//
	// example:
	//
	// Prod
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 	- The type of notification policies to be queried. Valid values: `false` (default): notification policies created in Application Real-Time Monitoring Service (ARMS).
	//
	// 	- `true`: notification policies created in an external system.
	//
	// >  You cannot use the ARMS console to modify the dispatch rules of a notification policy that is created in an external system.
	//
	// example:
	//
	// true
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
}

func (s ListDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleRequest) GetName() *string {
	return s.Name
}

func (s *ListDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDispatchRuleRequest) GetSystem() *bool {
	return s.System
}

func (s *ListDispatchRuleRequest) SetName(v string) *ListDispatchRuleRequest {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleRequest) SetRegionId(v string) *ListDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListDispatchRuleRequest) SetSystem(v bool) *ListDispatchRuleRequest {
	s.System = &v
	return s
}

func (s *ListDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
