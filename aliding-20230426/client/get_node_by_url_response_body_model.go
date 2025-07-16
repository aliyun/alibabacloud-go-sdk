// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *GetNodeByUrlResponseBodyNode) *GetNodeByUrlResponseBody
	GetNode() *GetNodeByUrlResponseBodyNode
	SetRequestId(v string) *GetNodeByUrlResponseBody
	GetRequestId() *string
}

type GetNodeByUrlResponseBody struct {
	Node *GetNodeByUrlResponseBodyNode `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetNodeByUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBody) GetNode() *GetNodeByUrlResponseBodyNode {
	return s.Node
}

func (s *GetNodeByUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeByUrlResponseBody) SetNode(v *GetNodeByUrlResponseBodyNode) *GetNodeByUrlResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeByUrlResponseBody) SetRequestId(v string) *GetNodeByUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeByUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeByUrlResponseBodyNode struct {
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
	// EpGBa2Lm8aRmzkkNhplMx1prWgN7R35y
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// READER
	PermissionRole *string `json:"PermissionRole,omitempty" xml:"PermissionRole,omitempty"`
	// example:
	//
	// 512
	Size            *int64                                       `json:"Size,omitempty" xml:"Size,omitempty"`
	StatisticalInfo *GetNodeByUrlResponseBodyNodeStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
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
	// By8jQS1ZYjGn5b0M
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetNodeByUrlResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBodyNode) GetCategory() *string {
	return s.Category
}

func (s *GetNodeByUrlResponseBodyNode) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNodeByUrlResponseBodyNode) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetNodeByUrlResponseBodyNode) GetExtension() *string {
	return s.Extension
}

func (s *GetNodeByUrlResponseBodyNode) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *GetNodeByUrlResponseBodyNode) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetNodeByUrlResponseBodyNode) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetNodeByUrlResponseBodyNode) GetName() *string {
	return s.Name
}

func (s *GetNodeByUrlResponseBodyNode) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeByUrlResponseBodyNode) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetNodeByUrlResponseBodyNode) GetSize() *int64 {
	return s.Size
}

func (s *GetNodeByUrlResponseBodyNode) GetStatisticalInfo() *GetNodeByUrlResponseBodyNodeStatisticalInfo {
	return s.StatisticalInfo
}

func (s *GetNodeByUrlResponseBodyNode) GetType() *string {
	return s.Type
}

func (s *GetNodeByUrlResponseBodyNode) GetUrl() *string {
	return s.Url
}

func (s *GetNodeByUrlResponseBodyNode) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetNodeByUrlResponseBodyNode) SetCategory(v string) *GetNodeByUrlResponseBodyNode {
	s.Category = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetCreateTime(v string) *GetNodeByUrlResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetCreatorId(v string) *GetNodeByUrlResponseBodyNode {
	s.CreatorId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetExtension(v string) *GetNodeByUrlResponseBodyNode {
	s.Extension = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetHasChildren(v bool) *GetNodeByUrlResponseBodyNode {
	s.HasChildren = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetModifiedTime(v string) *GetNodeByUrlResponseBodyNode {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetModifierId(v string) *GetNodeByUrlResponseBodyNode {
	s.ModifierId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetName(v string) *GetNodeByUrlResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetNodeId(v string) *GetNodeByUrlResponseBodyNode {
	s.NodeId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetPermissionRole(v string) *GetNodeByUrlResponseBodyNode {
	s.PermissionRole = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetSize(v int64) *GetNodeByUrlResponseBodyNode {
	s.Size = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetStatisticalInfo(v *GetNodeByUrlResponseBodyNodeStatisticalInfo) *GetNodeByUrlResponseBodyNode {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetType(v string) *GetNodeByUrlResponseBodyNode {
	s.Type = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetUrl(v string) *GetNodeByUrlResponseBodyNode {
	s.Url = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetWorkspaceId(v string) *GetNodeByUrlResponseBodyNode {
	s.WorkspaceId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) Validate() error {
	return dara.Validate(s)
}

type GetNodeByUrlResponseBodyNodeStatisticalInfo struct {
	// example:
	//
	// 123
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetNodeByUrlResponseBodyNodeStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlResponseBodyNodeStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBodyNodeStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *GetNodeByUrlResponseBodyNodeStatisticalInfo) SetWordCount(v int64) *GetNodeByUrlResponseBodyNodeStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNodeStatisticalInfo) Validate() error {
	return dara.Validate(s)
}
