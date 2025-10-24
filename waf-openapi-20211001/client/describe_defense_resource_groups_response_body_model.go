// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*DescribeDefenseResourceGroupsResponseBodyGroups) *DescribeDefenseResourceGroupsResponseBody
	GetGroups() []*DescribeDefenseResourceGroupsResponseBodyGroups
	SetRequestId(v string) *DescribeDefenseResourceGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDefenseResourceGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseResourceGroupsResponseBody struct {
	// The list of protected object groups.
	Groups []*DescribeDefenseResourceGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BFEC5C77-049B-5E88-A5B6-CB0C****B66E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponseBody) GetGroups() []*DescribeDefenseResourceGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeDefenseResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseResourceGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetGroups(v []*DescribeDefenseResourceGroupsResponseBodyGroups) *DescribeDefenseResourceGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDefenseResourceGroupsResponseBodyGroups struct {
	// The description of the protected object group.
	//
	// example:
	//
	// This is test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protected object group was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624343180000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The most recent time when the protected object group was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1701656305000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the protected object group.
	//
	// example:
	//
	// apptest
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The names of the protected objects that are added to the protected object group. Separate multiple protected objects with commas (,).
	//
	// example:
	//
	// example02.aliyun-waf,example01.aliyun-waf
	ResourceList *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) GetResourceList() *string {
	return s.ResourceList
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetDescription(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGmtCreate(v int64) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGmtModified(v int64) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGroupName(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetResourceList(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.ResourceList = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
