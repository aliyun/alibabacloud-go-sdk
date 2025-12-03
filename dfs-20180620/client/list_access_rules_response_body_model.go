// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRules(v []*ListAccessRulesResponseBodyAccessRules) *ListAccessRulesResponseBody
	GetAccessRules() []*ListAccessRulesResponseBodyAccessRules
	SetNextToken(v string) *ListAccessRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccessRulesResponseBody
	GetTotalCount() *int32
}

type ListAccessRulesResponseBody struct {
	AccessRules []*ListAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Repeated"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponseBody) GetAccessRules() []*ListAccessRulesResponseBodyAccessRules {
	return s.AccessRules
}

func (s *ListAccessRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccessRulesResponseBody) SetAccessRules(v []*ListAccessRulesResponseBodyAccessRules) *ListAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

func (s *ListAccessRulesResponseBody) SetNextToken(v string) *ListAccessRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessRulesResponseBody) SetRequestId(v string) *ListAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessRulesResponseBody) SetTotalCount(v int32) *ListAccessRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccessRulesResponseBody) Validate() error {
	if s.AccessRules != nil {
		for _, item := range s.AccessRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessRulesResponseBodyAccessRules struct {
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

func (s ListAccessRulesResponseBodyAccessRules) String() string {
	return dara.Prettify(s)
}

func (s ListAccessRulesResponseBodyAccessRules) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponseBodyAccessRules) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ListAccessRulesResponseBodyAccessRules) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *ListAccessRulesResponseBodyAccessRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessRulesResponseBodyAccessRules) GetDescription() *string {
	return s.Description
}

func (s *ListAccessRulesResponseBodyAccessRules) GetNetworkSegment() *string {
	return s.NetworkSegment
}

func (s *ListAccessRulesResponseBodyAccessRules) GetPriority() *int32 {
	return s.Priority
}

func (s *ListAccessRulesResponseBodyAccessRules) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *ListAccessRulesResponseBodyAccessRules) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccessRulesResponseBodyAccessRules) SetAccessGroupId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetAccessRuleId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.AccessRuleId = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetCreateTime(v string) *ListAccessRulesResponseBodyAccessRules {
	s.CreateTime = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetDescription(v string) *ListAccessRulesResponseBodyAccessRules {
	s.Description = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetNetworkSegment(v string) *ListAccessRulesResponseBodyAccessRules {
	s.NetworkSegment = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetPriority(v int32) *ListAccessRulesResponseBodyAccessRules {
	s.Priority = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetRWAccessType(v string) *ListAccessRulesResponseBodyAccessRules {
	s.RWAccessType = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetRegionId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.RegionId = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) Validate() error {
	return dara.Validate(s)
}
