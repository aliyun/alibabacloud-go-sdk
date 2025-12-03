// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v *DescribeIpWhitelistResponseBodyGroups) *DescribeIpWhitelistResponseBody
	GetGroups() *DescribeIpWhitelistResponseBodyGroups
	SetRequestId(v string) *DescribeIpWhitelistResponseBody
	GetRequestId() *string
}

type DescribeIpWhitelistResponseBody struct {
	Groups *DescribeIpWhitelistResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// example:
	//
	// AFAA617B-3268-5883-982B-DB8EC8CC1F1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBody) GetGroups() *DescribeIpWhitelistResponseBodyGroups {
	return s.Groups
}

func (s *DescribeIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpWhitelistResponseBody) SetGroups(v *DescribeIpWhitelistResponseBodyGroups) *DescribeIpWhitelistResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeIpWhitelistResponseBody) SetRequestId(v string) *DescribeIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpWhitelistResponseBody) Validate() error {
	if s.Groups != nil {
		if err := s.Groups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpWhitelistResponseBodyGroups struct {
	Group []*DescribeIpWhitelistResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroups) GetGroup() []*DescribeIpWhitelistResponseBodyGroupsGroup {
	return s.Group
}

func (s *DescribeIpWhitelistResponseBodyGroups) SetGroup(v []*DescribeIpWhitelistResponseBodyGroupsGroup) *DescribeIpWhitelistResponseBodyGroups {
	s.Group = v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroups) Validate() error {
	if s.Group != nil {
		for _, item := range s.Group {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpWhitelistResponseBodyGroupsGroup struct {
	// example:
	//
	// default
	GroupName *string                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *DescribeIpWhitelistResponseBodyGroupsGroupIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Struct"`
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) GetIpList() *DescribeIpWhitelistResponseBodyGroupsGroupIpList {
	return s.IpList
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetGroupName(v string) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpList(v *DescribeIpWhitelistResponseBodyGroupsGroupIpList) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpList = v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpVersion(v int32) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) Validate() error {
	if s.IpList != nil {
		if err := s.IpList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpWhitelistResponseBodyGroupsGroupIpList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroupIpList) GetIp() []*string {
	return s.Ip
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroupIpList) SetIp(v []*string) *DescribeIpWhitelistResponseBodyGroupsGroupIpList {
	s.Ip = v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroupIpList) Validate() error {
	return dara.Validate(s)
}
