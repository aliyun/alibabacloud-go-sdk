// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *ModifyAccessRuleRequest
	GetAccessGroupId() *string
	SetAccessRuleId(v string) *ModifyAccessRuleRequest
	GetAccessRuleId() *string
	SetDescription(v string) *ModifyAccessRuleRequest
	GetDescription() *string
	SetInputRegionId(v string) *ModifyAccessRuleRequest
	GetInputRegionId() *string
	SetPriority(v int32) *ModifyAccessRuleRequest
	GetPriority() *int32
	SetRWAccessType(v string) *ModifyAccessRuleRequest
	GetRWAccessType() *string
}

type ModifyAccessRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acr-c38028f0-f313-4385-9456-3501b1f5****
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
}

func (s ModifyAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ModifyAccessRuleRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *ModifyAccessRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAccessRuleRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ModifyAccessRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyAccessRuleRequest) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *ModifyAccessRuleRequest) SetAccessGroupId(v string) *ModifyAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetAccessRuleId(v string) *ModifyAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetDescription(v string) *ModifyAccessRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetInputRegionId(v string) *ModifyAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
