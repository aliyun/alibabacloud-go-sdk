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
	// The number of query results.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of user groups.
	Groups []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
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

type DescribeGroupsResponseBodyGroups struct {
	// The logon policy associated with the user group.
	AttachedLoginPolicy *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy `json:"AttachedLoginPolicy,omitempty" xml:"AttachedLoginPolicy,omitempty" type:"Struct"`
	// The list of allocated resources.
	AuthedResources map[string]*string `json:"AuthedResources,omitempty" xml:"AuthedResources,omitempty"`
	// The creation time in the ISO 8601 standard in the UTC+08:00 time zone.
	//
	// example:
	//
	// 2025-08-07T13:40:40+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user group.
	//
	// example:
	//
	// A test group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether approval is required for downloads.
	DownloadNeedApproval *bool `json:"DownloadNeedApproval,omitempty" xml:"DownloadNeedApproval,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// ug-2412ojkwtybd****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// TestGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether file approval is enabled.
	//
	// example:
	//
	// false
	TransferFileNeedApproval *bool `json:"TransferFileNeedApproval,omitempty" xml:"TransferFileNeedApproval,omitempty"`
	// Indicates whether approval is required for uploads.
	UploadNeedApproval *bool `json:"UploadNeedApproval,omitempty" xml:"UploadNeedApproval,omitempty"`
	// The number of members in the user group.
	//
	// example:
	//
	// 20
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) GetAttachedLoginPolicy() *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy {
	return s.AttachedLoginPolicy
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

func (s *DescribeGroupsResponseBodyGroups) GetDownloadNeedApproval() *bool {
	return s.DownloadNeedApproval
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

func (s *DescribeGroupsResponseBodyGroups) GetUploadNeedApproval() *bool {
	return s.UploadNeedApproval
}

func (s *DescribeGroupsResponseBodyGroups) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeGroupsResponseBodyGroups) SetAttachedLoginPolicy(v *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) *DescribeGroupsResponseBodyGroups {
	s.AttachedLoginPolicy = v
	return s
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

func (s *DescribeGroupsResponseBodyGroups) SetDownloadNeedApproval(v bool) *DescribeGroupsResponseBodyGroups {
	s.DownloadNeedApproval = &v
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

func (s *DescribeGroupsResponseBodyGroups) SetUploadNeedApproval(v bool) *DescribeGroupsResponseBodyGroups {
	s.UploadNeedApproval = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetUserCount(v int32) *DescribeGroupsResponseBodyGroups {
	s.UserCount = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) Validate() error {
	if s.AttachedLoginPolicy != nil {
		if err := s.AttachedLoginPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupsResponseBodyGroupsAttachedLoginPolicy struct {
	// The Policy Name of the logon policy.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the logon policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) GetName() *string {
	return s.Name
}

func (s *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) SetName(v string) *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy {
	s.Name = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) SetPolicyId(v string) *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy {
	s.PolicyId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsAttachedLoginPolicy) Validate() error {
	return dara.Validate(s)
}
