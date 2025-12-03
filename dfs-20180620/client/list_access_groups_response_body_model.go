// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroups(v []*ListAccessGroupsResponseBodyAccessGroups) *ListAccessGroupsResponseBody
	GetAccessGroups() []*ListAccessGroupsResponseBodyAccessGroups
	SetNextToken(v string) *ListAccessGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccessGroupsResponseBody
	GetTotalCount() *int32
}

type ListAccessGroupsResponseBody struct {
	AccessGroups []*ListAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Repeated"`
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponseBody) GetAccessGroups() []*ListAccessGroupsResponseBodyAccessGroups {
	return s.AccessGroups
}

func (s *ListAccessGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccessGroupsResponseBody) SetAccessGroups(v []*ListAccessGroupsResponseBodyAccessGroups) *ListAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

func (s *ListAccessGroupsResponseBody) SetNextToken(v string) *ListAccessGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessGroupsResponseBody) SetRequestId(v string) *ListAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessGroupsResponseBody) SetTotalCount(v int32) *ListAccessGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccessGroupsResponseBody) Validate() error {
	if s.AccessGroups != nil {
		for _, item := range s.AccessGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessGroupsResponseBodyAccessGroups struct {
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// example:
	//
	// test-cluster-policy
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 1
	MountPointCount *int32 `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s ListAccessGroupsResponseBodyAccessGroups) String() string {
	return dara.Prettify(s)
}

func (s ListAccessGroupsResponseBodyAccessGroups) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetDescription() *string {
	return s.Description
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetMountPointCount() *int32 {
	return s.MountPointCount
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccessGroupsResponseBodyAccessGroups) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetAccessGroupId(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetAccessGroupName(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.AccessGroupName = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetCreateTime(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.CreateTime = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetDescription(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.Description = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetIsDefault(v bool) *ListAccessGroupsResponseBodyAccessGroups {
	s.IsDefault = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetMountPointCount(v int32) *ListAccessGroupsResponseBodyAccessGroups {
	s.MountPointCount = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetNetworkType(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.NetworkType = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetRegionId(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.RegionId = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetRuleCount(v int32) *ListAccessGroupsResponseBodyAccessGroups {
	s.RuleCount = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) Validate() error {
	return dara.Validate(s)
}
