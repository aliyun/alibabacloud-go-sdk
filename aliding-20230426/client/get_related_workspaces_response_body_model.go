// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRelatedWorkspacesResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetRelatedWorkspacesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetRelatedWorkspacesResponseBody
	GetVendorType() *string
	SetWorkspaces(v []*GetRelatedWorkspacesResponseBodyWorkspaces) *GetRelatedWorkspacesResponseBody
	GetWorkspaces() []*GetRelatedWorkspacesResponseBodyWorkspaces
}

type GetRelatedWorkspacesResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string                                       `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	Workspaces []*GetRelatedWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s GetRelatedWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRelatedWorkspacesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetRelatedWorkspacesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetRelatedWorkspacesResponseBody) GetWorkspaces() []*GetRelatedWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *GetRelatedWorkspacesResponseBody) SetRequestId(v string) *GetRelatedWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBody) SetVendorRequestId(v string) *GetRelatedWorkspacesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBody) SetVendorType(v string) *GetRelatedWorkspacesResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBody) SetWorkspaces(v []*GetRelatedWorkspacesResponseBodyWorkspaces) *GetRelatedWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *GetRelatedWorkspacesResponseBody) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRelatedWorkspacesResponseBodyWorkspaces struct {
	// example:
	//
	// 1638256965936
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// example:
	//
	// 知识库
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Q2xwPOKiSLxxxx
	Owner      *string                                                 `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RecentList []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList `json:"RecentList,omitempty" xml:"RecentList,omitempty" type:"Repeated"`
	// example:
	//
	// OWNER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// https://alidocs.xxxx/nb9XJKdxxxxmyAp/docs/nb9XxxxxxxmyAp
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// nb9XJKdxxxxmyAp
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetRelatedWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetDeleted() *bool {
	return s.Deleted
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetName() *string {
	return s.Name
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetOwner() *string {
	return s.Owner
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetRecentList() []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	return s.RecentList
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetRole() *string {
	return s.Role
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetUrl() *string {
	return s.Url
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetCreateTime(v int64) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetDeleted(v bool) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Deleted = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetName(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetOwner(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Owner = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetRecentList(v []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.RecentList = v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetRole(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Role = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetUrl(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Url = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) Validate() error {
	if s.RecentList != nil {
		for _, item := range s.RecentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRelatedWorkspacesResponseBodyWorkspacesRecentList struct {
	// example:
	//
	// 1638256965936
	LastEditTime *int64 `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	// example:
	//
	// 知识库
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// nb9XxxxxxxmyAp
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// https://alidocs.xxxx/nb9XJKdxxxxmyAp
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRelatedWorkspacesResponseBodyWorkspacesRecentList) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GetLastEditTime() *int64 {
	return s.LastEditTime
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GetName() *string {
	return s.Name
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GetUrl() *string {
	return s.Url
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetLastEditTime(v int64) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.LastEditTime = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetName(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.Name = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetNodeId(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.NodeId = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetUrl(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.Url = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) Validate() error {
	return dara.Validate(s)
}
