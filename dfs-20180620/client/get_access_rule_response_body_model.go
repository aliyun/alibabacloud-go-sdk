// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRule(v *GetAccessRuleResponseBodyAccessRule) *GetAccessRuleResponseBody
	GetAccessRule() *GetAccessRuleResponseBodyAccessRule
	SetRequestId(v string) *GetAccessRuleResponseBody
	GetRequestId() *string
}

type GetAccessRuleResponseBody struct {
	AccessRule *GetAccessRuleResponseBodyAccessRule `json:"AccessRule,omitempty" xml:"AccessRule,omitempty" type:"Struct"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponseBody) GetAccessRule() *GetAccessRuleResponseBodyAccessRule {
	return s.AccessRule
}

func (s *GetAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessRuleResponseBody) SetAccessRule(v *GetAccessRuleResponseBodyAccessRule) *GetAccessRuleResponseBody {
	s.AccessRule = v
	return s
}

func (s *GetAccessRuleResponseBody) SetRequestId(v string) *GetAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessRuleResponseBody) Validate() error {
	if s.AccessRule != nil {
		if err := s.AccessRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessRuleResponseBodyAccessRule struct {
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// example:
	//
	// acr-c38028f0-f313-4385-9456-3501b1f5****
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 192.0.2.0/24
	NetworkSegment *string `json:"NetworkSegment,omitempty" xml:"NetworkSegment,omitempty"`
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAccessRuleResponseBodyAccessRule) String() string {
	return dara.Prettify(s)
}

func (s GetAccessRuleResponseBodyAccessRule) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponseBodyAccessRule) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *GetAccessRuleResponseBodyAccessRule) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *GetAccessRuleResponseBodyAccessRule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAccessRuleResponseBodyAccessRule) GetDescription() *string {
	return s.Description
}

func (s *GetAccessRuleResponseBodyAccessRule) GetNetworkSegment() *string {
	return s.NetworkSegment
}

func (s *GetAccessRuleResponseBodyAccessRule) GetPriority() *int32 {
	return s.Priority
}

func (s *GetAccessRuleResponseBodyAccessRule) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *GetAccessRuleResponseBodyAccessRule) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAccessRuleResponseBodyAccessRule) SetAccessGroupId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetAccessRuleId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.AccessRuleId = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetCreateTime(v string) *GetAccessRuleResponseBodyAccessRule {
	s.CreateTime = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetDescription(v string) *GetAccessRuleResponseBodyAccessRule {
	s.Description = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetNetworkSegment(v string) *GetAccessRuleResponseBodyAccessRule {
	s.NetworkSegment = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetPriority(v int32) *GetAccessRuleResponseBodyAccessRule {
	s.Priority = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetRWAccessType(v string) *GetAccessRuleResponseBodyAccessRule {
	s.RWAccessType = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetRegionId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.RegionId = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) Validate() error {
	return dara.Validate(s)
}
