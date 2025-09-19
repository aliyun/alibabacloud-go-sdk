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
	SetUserVpc(v *UserVpc) *ResourceGroup
	GetUserVpc() *UserVpc
	SetVersion(v string) *ResourceGroup
	GetVersion() *string
	SetWorkspaceID(v string) *ResourceGroup
	GetWorkspaceID() *string
}

type ResourceGroup struct {
	CreatorID       *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	GmtCreatedTime  *string `json:"GmtCreatedTime,omitempty" xml:"GmtCreatedTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// RG1
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeCount *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string  `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	UserVpc         *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	Version         *string  `json:"Version,omitempty" xml:"Version,omitempty"`
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
	return dara.Validate(s)
}
