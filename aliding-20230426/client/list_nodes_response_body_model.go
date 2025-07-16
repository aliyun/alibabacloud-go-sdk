// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody
	GetNodes() []*ListNodesResponseBodyNodes
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
}

type ListNodesResponseBody struct {
	// example:
	//
	// next_token
	NextToken *string                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Nodes     []*ListNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodesResponseBody) GetNodes() []*ListNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) SetNextToken(v string) *ListNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodesResponseBody) SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyNodes struct {
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
	// true
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
	Size            *int64                                     `json:"Size,omitempty" xml:"Size,omitempty"`
	StatisticalInfo *ListNodesResponseBodyNodesStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
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

func (s ListNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) GetCategory() *string {
	return s.Category
}

func (s *ListNodesResponseBodyNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNodesResponseBodyNodes) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListNodesResponseBodyNodes) GetExtension() *string {
	return s.Extension
}

func (s *ListNodesResponseBodyNodes) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *ListNodesResponseBodyNodes) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListNodesResponseBodyNodes) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListNodesResponseBodyNodes) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNodesResponseBodyNodes) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *ListNodesResponseBodyNodes) GetSize() *int64 {
	return s.Size
}

func (s *ListNodesResponseBodyNodes) GetStatisticalInfo() *ListNodesResponseBodyNodesStatisticalInfo {
	return s.StatisticalInfo
}

func (s *ListNodesResponseBodyNodes) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyNodes) GetUrl() *string {
	return s.Url
}

func (s *ListNodesResponseBodyNodes) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNodesResponseBodyNodes) SetCategory(v string) *ListNodesResponseBodyNodes {
	s.Category = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreateTime(v string) *ListNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreatorId(v string) *ListNodesResponseBodyNodes {
	s.CreatorId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExtension(v string) *ListNodesResponseBodyNodes {
	s.Extension = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHasChildren(v bool) *ListNodesResponseBodyNodes {
	s.HasChildren = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetModifiedTime(v string) *ListNodesResponseBodyNodes {
	s.ModifiedTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetModifierId(v string) *ListNodesResponseBodyNodes {
	s.ModifierId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetName(v string) *ListNodesResponseBodyNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetNodeId(v string) *ListNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetPermissionRole(v string) *ListNodesResponseBodyNodes {
	s.PermissionRole = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetSize(v int64) *ListNodesResponseBodyNodes {
	s.Size = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStatisticalInfo(v *ListNodesResponseBodyNodesStatisticalInfo) *ListNodesResponseBodyNodes {
	s.StatisticalInfo = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetType(v string) *ListNodesResponseBodyNodes {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetUrl(v string) *ListNodesResponseBodyNodes {
	s.Url = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetWorkspaceId(v string) *ListNodesResponseBodyNodes {
	s.WorkspaceId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyNodesStatisticalInfo struct {
	// example:
	//
	// 123
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s ListNodesResponseBodyNodesStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyNodesStatisticalInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *ListNodesResponseBodyNodesStatisticalInfo) SetWordCount(v int64) *ListNodesResponseBodyNodesStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *ListNodesResponseBodyNodesStatisticalInfo) Validate() error {
	return dara.Validate(s)
}
