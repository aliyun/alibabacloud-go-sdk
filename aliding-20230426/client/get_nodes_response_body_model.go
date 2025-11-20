// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*GetNodesResponseBodyNodes) *GetNodesResponseBody
	GetNodes() []*GetNodesResponseBodyNodes
	SetRequestId(v string) *GetNodesResponseBody
	GetRequestId() *string
}

type GetNodesResponseBody struct {
	Nodes []*GetNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBody) GetNodes() []*GetNodesResponseBodyNodes {
	return s.Nodes
}

func (s *GetNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodesResponseBody) SetNodes(v []*GetNodesResponseBodyNodes) *GetNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *GetNodesResponseBody) SetRequestId(v string) *GetNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodesResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNodesResponseBodyNodes struct {
	// example:
	//
	// ALIDOC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 01472825524039877041
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// adoc
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// false
	HasChildren *bool `json:"HasChildren,omitempty" xml:"HasChildren,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 01472825524039877041
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// node_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MNDoBb60VLBPraakI1Ywxyyn8lemrZQ3
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// READER
	PermissionRole *string `json:"PermissionRole,omitempty" xml:"PermissionRole,omitempty"`
	// example:
	//
	// 512
	Size            *int64                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	StatisticalInfo *GetNodesResponseBodyNodesStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// node_url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// MNDoBb60VLBPraakI1Ywxyyn8lemrZQ3
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s GetNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBodyNodes) GetCategory() *string {
	return s.Category
}

func (s *GetNodesResponseBodyNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNodesResponseBodyNodes) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetNodesResponseBodyNodes) GetExtension() *string {
	return s.Extension
}

func (s *GetNodesResponseBodyNodes) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *GetNodesResponseBodyNodes) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetNodesResponseBodyNodes) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetNodesResponseBodyNodes) GetName() *string {
	return s.Name
}

func (s *GetNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodesResponseBodyNodes) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetNodesResponseBodyNodes) GetSize() *int64 {
	return s.Size
}

func (s *GetNodesResponseBodyNodes) GetStatisticalInfo() *GetNodesResponseBodyNodesStatisticalInfo {
	return s.StatisticalInfo
}

func (s *GetNodesResponseBodyNodes) GetType() *string {
	return s.Type
}

func (s *GetNodesResponseBodyNodes) GetUrl() *string {
	return s.Url
}

func (s *GetNodesResponseBodyNodes) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetNodesResponseBodyNodes) SetCategory(v string) *GetNodesResponseBodyNodes {
	s.Category = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetCreateTime(v string) *GetNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetCreatorId(v string) *GetNodesResponseBodyNodes {
	s.CreatorId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetExtension(v string) *GetNodesResponseBodyNodes {
	s.Extension = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetHasChildren(v bool) *GetNodesResponseBodyNodes {
	s.HasChildren = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetModifiedTime(v string) *GetNodesResponseBodyNodes {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetModifierId(v string) *GetNodesResponseBodyNodes {
	s.ModifierId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetName(v string) *GetNodesResponseBodyNodes {
	s.Name = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetNodeId(v string) *GetNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetPermissionRole(v string) *GetNodesResponseBodyNodes {
	s.PermissionRole = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetSize(v int64) *GetNodesResponseBodyNodes {
	s.Size = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetStatisticalInfo(v *GetNodesResponseBodyNodesStatisticalInfo) *GetNodesResponseBodyNodes {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodesResponseBodyNodes) SetType(v string) *GetNodesResponseBodyNodes {
	s.Type = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetUrl(v string) *GetNodesResponseBodyNodes {
	s.Url = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetWorkspaceId(v string) *GetNodesResponseBodyNodes {
	s.WorkspaceId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) Validate() error {
	if s.StatisticalInfo != nil {
		if err := s.StatisticalInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodesResponseBodyNodesStatisticalInfo struct {
	// example:
	//
	// 123
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetNodesResponseBodyNodesStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNodesResponseBodyNodesStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBodyNodesStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *GetNodesResponseBodyNodesStatisticalInfo) SetWordCount(v int64) *GetNodesResponseBodyNodesStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *GetNodesResponseBodyNodesStatisticalInfo) Validate() error {
	return dara.Validate(s)
}
