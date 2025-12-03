// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroup(v *GetAccessGroupResponseBodyAccessGroup) *GetAccessGroupResponseBody
	GetAccessGroup() *GetAccessGroupResponseBodyAccessGroup
	SetRequestId(v string) *GetAccessGroupResponseBody
	GetRequestId() *string
}

type GetAccessGroupResponseBody struct {
	AccessGroup *GetAccessGroupResponseBodyAccessGroup `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty" type:"Struct"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponseBody) GetAccessGroup() *GetAccessGroupResponseBodyAccessGroup {
	return s.AccessGroup
}

func (s *GetAccessGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessGroupResponseBody) SetAccessGroup(v *GetAccessGroupResponseBodyAccessGroup) *GetAccessGroupResponseBody {
	s.AccessGroup = v
	return s
}

func (s *GetAccessGroupResponseBody) SetRequestId(v string) *GetAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessGroupResponseBody) Validate() error {
	if s.AccessGroup != nil {
		if err := s.AccessGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessGroupResponseBodyAccessGroup struct {
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

func (s GetAccessGroupResponseBodyAccessGroup) String() string {
	return dara.Prettify(s)
}

func (s GetAccessGroupResponseBodyAccessGroup) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetDescription() *string {
	return s.Description
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetMountPointCount() *int32 {
	return s.MountPointCount
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAccessGroupResponseBodyAccessGroup) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetAccessGroupId(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetAccessGroupName(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.AccessGroupName = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetCreateTime(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.CreateTime = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetDescription(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.Description = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetIsDefault(v bool) *GetAccessGroupResponseBodyAccessGroup {
	s.IsDefault = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetMountPointCount(v int32) *GetAccessGroupResponseBodyAccessGroup {
	s.MountPointCount = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetNetworkType(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.NetworkType = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetRegionId(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.RegionId = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetRuleCount(v int32) *GetAccessGroupResponseBodyAccessGroup {
	s.RuleCount = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) Validate() error {
	return dara.Validate(s)
}
