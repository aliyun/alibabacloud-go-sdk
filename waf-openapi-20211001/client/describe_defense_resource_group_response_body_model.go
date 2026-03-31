// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *DescribeDefenseResourceGroupResponseBodyGroup) *DescribeDefenseResourceGroupResponseBody
	GetGroup() *DescribeDefenseResourceGroupResponseBodyGroup
	SetRequestId(v string) *DescribeDefenseResourceGroupResponseBody
	GetRequestId() *string
}

type DescribeDefenseResourceGroupResponseBody struct {
	// The information about the protected object group.
	Group *DescribeDefenseResourceGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E67D21C6-5376-5F94-B745-70E08D03E3CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBody) GetGroup() *DescribeDefenseResourceGroupResponseBodyGroup {
	return s.Group
}

func (s *DescribeDefenseResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceGroupResponseBody) SetGroup(v *DescribeDefenseResourceGroupResponseBodyGroup) *DescribeDefenseResourceGroupResponseBody {
	s.Group = v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDefenseResourceGroupResponseBodyGroup struct {
	// The description of the protected object group.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protected object group was created.
	//
	// example:
	//
	// 23242312312
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The most recent time when the protected object group was modified.
	//
	// example:
	//
	// 23242312312
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the protected object group.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The protected objects in the protected object group. The protected objects are separated with commas (,).
	//
	// example:
	//
	// test1.aliyundoc.com,test2.aliyundoc.com
	ResourceList *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) GetResourceList() *string {
	return s.ResourceList
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetDescription(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtCreate(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtModified(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGroupName(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetResourceList(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.ResourceList = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}
