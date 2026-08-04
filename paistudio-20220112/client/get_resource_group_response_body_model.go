// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterID(v string) *GetResourceGroupResponseBody
	GetClusterID() *string
	SetComputingResourceProvider(v string) *GetResourceGroupResponseBody
	GetComputingResourceProvider() *string
	SetCreatorID(v string) *GetResourceGroupResponseBody
	GetCreatorID() *string
	SetDescription(v string) *GetResourceGroupResponseBody
	GetDescription() *string
	SetGmtCreatedTime(v string) *GetResourceGroupResponseBody
	GetGmtCreatedTime() *string
	SetGmtModifiedTime(v string) *GetResourceGroupResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetResourceGroupResponseBody
	GetName() *string
	SetRequestId(v string) *GetResourceGroupResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetResourceGroupResponseBody
	GetResourceType() *string
	SetStatus(v string) *GetResourceGroupResponseBody
	GetStatus() *string
	SetSupportRDMA(v bool) *GetResourceGroupResponseBody
	GetSupportRDMA() *bool
	SetTags(v []*GetResourceGroupResponseBodyTags) *GetResourceGroupResponseBody
	GetTags() []*GetResourceGroupResponseBodyTags
	SetUserVpc(v *UserVpc) *GetResourceGroupResponseBody
	GetUserVpc() *UserVpc
	SetVersion(v string) *GetResourceGroupResponseBody
	GetVersion() *string
	SetWorkspaceID(v string) *GetResourceGroupResponseBody
	GetWorkspaceID() *string
}

type GetResourceGroupResponseBody struct {
	// The ID of the cluster to which the resource group belongs.
	//
	// example:
	//
	// cb2c7bde30b774e46a329c
	ClusterID *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	// The provider of computing resources for the resource group.
	//
	// example:
	//
	// ECS
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	// The ID of the user who created the resource group.
	//
	// example:
	//
	// 1612285282502324
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// The description of the resource group.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// The time when the resource group was last modified.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// TestResourceGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of computing resource for the resource group.
	//
	// example:
	//
	// Ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the resource group.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the resource group supports RDMA.
	//
	// example:
	//
	// true
	SupportRDMA *bool `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	// A list of tags.
	Tags []*GetResourceGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The details of the VPC connected to the resource group.
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	Version *string  `json:"Version,omitempty" xml:"Version,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 35201
	WorkspaceID *string `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
}

func (s GetResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) GetClusterID() *string {
	return s.ClusterID
}

func (s *GetResourceGroupResponseBody) GetComputingResourceProvider() *string {
	return s.ComputingResourceProvider
}

func (s *GetResourceGroupResponseBody) GetCreatorID() *string {
	return s.CreatorID
}

func (s *GetResourceGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetResourceGroupResponseBody) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *GetResourceGroupResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetResourceGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *GetResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetResourceGroupResponseBody) GetSupportRDMA() *bool {
	return s.SupportRDMA
}

func (s *GetResourceGroupResponseBody) GetTags() []*GetResourceGroupResponseBodyTags {
	return s.Tags
}

func (s *GetResourceGroupResponseBody) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *GetResourceGroupResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetResourceGroupResponseBody) GetWorkspaceID() *string {
	return s.WorkspaceID
}

func (s *GetResourceGroupResponseBody) SetClusterID(v string) *GetResourceGroupResponseBody {
	s.ClusterID = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetComputingResourceProvider(v string) *GetResourceGroupResponseBody {
	s.ComputingResourceProvider = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetCreatorID(v string) *GetResourceGroupResponseBody {
	s.CreatorID = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetDescription(v string) *GetResourceGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetGmtCreatedTime(v string) *GetResourceGroupResponseBody {
	s.GmtCreatedTime = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetGmtModifiedTime(v string) *GetResourceGroupResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetName(v string) *GetResourceGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceType(v string) *GetResourceGroupResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetStatus(v string) *GetResourceGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetSupportRDMA(v bool) *GetResourceGroupResponseBody {
	s.SupportRDMA = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetTags(v []*GetResourceGroupResponseBodyTags) *GetResourceGroupResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceGroupResponseBody) SetUserVpc(v *UserVpc) *GetResourceGroupResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetResourceGroupResponseBody) SetVersion(v string) *GetResourceGroupResponseBody {
	s.Version = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetWorkspaceID(v string) *GetResourceGroupResponseBody {
	s.WorkspaceID = &v
	return s
}

func (s *GetResourceGroupResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceGroupResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetResourceGroupResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetResourceGroupResponseBodyTags) SetTagKey(v string) *GetResourceGroupResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupResponseBodyTags) SetTagValue(v string) *GetResourceGroupResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetResourceGroupResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
