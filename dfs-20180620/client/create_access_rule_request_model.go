// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *CreateAccessRuleRequest
	GetAccessGroupId() *string
	SetDescription(v string) *CreateAccessRuleRequest
	GetDescription() *string
	SetInputRegionId(v string) *CreateAccessRuleRequest
	GetInputRegionId() *string
	SetNetworkSegment(v string) *CreateAccessRuleRequest
	GetNetworkSegment() *string
	SetPriority(v int32) *CreateAccessRuleRequest
	GetPriority() *int32
	SetRWAccessType(v string) *CreateAccessRuleRequest
	GetRWAccessType() *string
}

type CreateAccessRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.0.2.0/24
	NetworkSegment *string `json:"NetworkSegment,omitempty" xml:"NetworkSegment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
}

func (s CreateAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *CreateAccessRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessRuleRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateAccessRuleRequest) GetNetworkSegment() *string {
	return s.NetworkSegment
}

func (s *CreateAccessRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateAccessRuleRequest) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *CreateAccessRuleRequest) SetAccessGroupId(v string) *CreateAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *CreateAccessRuleRequest) SetDescription(v string) *CreateAccessRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessRuleRequest) SetInputRegionId(v string) *CreateAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateAccessRuleRequest) SetNetworkSegment(v string) *CreateAccessRuleRequest {
	s.NetworkSegment = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
