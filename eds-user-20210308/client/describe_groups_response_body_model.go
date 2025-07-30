// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeGroupsResponseBody
	GetCount() *int32
	SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody
	GetGroups() []*DescribeGroupsResponseBodyGroups
	SetRequestId(v string) *DescribeGroupsResponseBody
	GetRequestId() *string
}

type DescribeGroupsResponseBody struct {
	Count  *int32                              `json:"Count,omitempty" xml:"Count,omitempty"`
	Groups []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGroupsResponseBody) GetGroups() []*DescribeGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupsResponseBody) SetCount(v int32) *DescribeGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupsResponseBody) SetRequestId(v string) *DescribeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupsResponseBodyGroups struct {
	AuthedResources map[string]*string `json:"AuthedResources,omitempty" xml:"AuthedResources,omitempty"`
	CreateTime      *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ug-2412ojkwtybd****
	GroupId                  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName                *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	TransferFileNeedApproval *bool   `json:"TransferFileNeedApproval,omitempty" xml:"TransferFileNeedApproval,omitempty"`
	UserCount                *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) GetAuthedResources() map[string]*string {
	return s.AuthedResources
}

func (s *DescribeGroupsResponseBodyGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribeGroupsResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGroupsResponseBodyGroups) GetTransferFileNeedApproval() *bool {
	return s.TransferFileNeedApproval
}

func (s *DescribeGroupsResponseBodyGroups) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeGroupsResponseBodyGroups) SetAuthedResources(v map[string]*string) *DescribeGroupsResponseBodyGroups {
	s.AuthedResources = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetCreateTime(v string) *DescribeGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetDescription(v string) *DescribeGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGroupId(v string) *DescribeGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGroupName(v string) *DescribeGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetTransferFileNeedApproval(v bool) *DescribeGroupsResponseBodyGroups {
	s.TransferFileNeedApproval = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetUserCount(v int32) *DescribeGroupsResponseBodyGroups {
	s.UserCount = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
