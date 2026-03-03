// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceGroup interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorID(v string) *ResourceGroup
	GetCreatorID() *string
	SetGmtCreatedTime(v string) *ResourceGroup
	GetGmtCreatedTime() *string
	SetGmtModifiedTime(v string) *ResourceGroup
	GetGmtModifiedTime() *string
	SetName(v string) *ResourceGroup
	GetName() *string
	SetNodeCount(v int32) *ResourceGroup
	GetNodeCount() *int32
	SetResourceGroupID(v string) *ResourceGroup
	GetResourceGroupID() *string
	SetResourceType(v string) *ResourceGroup
	GetResourceType() *string
	SetStatus(v string) *ResourceGroup
	GetStatus() *string
	SetUserVpc(v *UserVpc) *ResourceGroup
	GetUserVpc() *UserVpc
	SetVersion(v string) *ResourceGroup
	GetVersion() *string
	SetWorkspaceID(v string) *ResourceGroup
	GetWorkspaceID() *string
}

type ResourceGroup struct {
	// CreatorID
	//
	// example:
	//
	// 1612285282502324
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	// pai resource created time
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtCreatedTime *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	// GmtModified
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// pai resource group name
	//
	// example:
	//
	// RG1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// NodeCount
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// pai resource group id
	//
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// UserVpc
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	Version *string  `json:"Version,omitempty" xml:"Version,omitempty"`
	// pworkspace id
	//
	// example:
	//
	// 23000
	WorkspaceID *string `json:"WorkspaceID,omitempty" xml:"WorkspaceID,omitempty"`
}

func (s ResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s ResourceGroup) GoString() string {
	return s.String()
}

func (s *ResourceGroup) GetCreatorID() *string {
	return s.CreatorID
}

func (s *ResourceGroup) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *ResourceGroup) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ResourceGroup) GetName() *string {
	return s.Name
}

func (s *ResourceGroup) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ResourceGroup) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ResourceGroup) GetResourceType() *string {
	return s.ResourceType
}

func (s *ResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *ResourceGroup) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *ResourceGroup) GetVersion() *string {
	return s.Version
}

func (s *ResourceGroup) GetWorkspaceID() *string {
	return s.WorkspaceID
}

func (s *ResourceGroup) SetCreatorID(v string) *ResourceGroup {
	s.CreatorID = &v
	return s
}

func (s *ResourceGroup) SetGmtCreatedTime(v string) *ResourceGroup {
	s.GmtCreatedTime = &v
	return s
}

func (s *ResourceGroup) SetGmtModifiedTime(v string) *ResourceGroup {
	s.GmtModifiedTime = &v
	return s
}

func (s *ResourceGroup) SetName(v string) *ResourceGroup {
	s.Name = &v
	return s
}

func (s *ResourceGroup) SetNodeCount(v int32) *ResourceGroup {
	s.NodeCount = &v
	return s
}

func (s *ResourceGroup) SetResourceGroupID(v string) *ResourceGroup {
	s.ResourceGroupID = &v
	return s
}

func (s *ResourceGroup) SetResourceType(v string) *ResourceGroup {
	s.ResourceType = &v
	return s
}

func (s *ResourceGroup) SetStatus(v string) *ResourceGroup {
	s.Status = &v
	return s
}

func (s *ResourceGroup) SetUserVpc(v *UserVpc) *ResourceGroup {
	s.UserVpc = v
	return s
}

func (s *ResourceGroup) SetVersion(v string) *ResourceGroup {
	s.Version = &v
	return s
}

func (s *ResourceGroup) SetWorkspaceID(v string) *ResourceGroup {
	s.WorkspaceID = &v
	return s
}

func (s *ResourceGroup) Validate() error {
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}
