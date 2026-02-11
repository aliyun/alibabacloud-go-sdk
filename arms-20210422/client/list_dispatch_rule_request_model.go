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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	System   *bool   `json:"System,omitempty" xml:"System,omitempty"`
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
