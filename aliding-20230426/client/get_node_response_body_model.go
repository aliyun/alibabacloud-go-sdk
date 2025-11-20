// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody
	GetNode() *GetNodeResponseBodyNode
	SetRequestId(v string) *GetNodeResponseBody
	GetRequestId() *string
}

type GetNodeResponseBody struct {
	Node *GetNodeResponseBodyNode `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) GetNode() *GetNodeResponseBodyNode {
	return s.Node
}

func (s *GetNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeResponseBody) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeResponseBodyNode struct {
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
	// 123123
	Size            *int64                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	StatisticalInfo *GetNodeResponseBodyNodeStatisticalInfo `json:"StatisticalInfo,omitempty" xml:"StatisticalInfo,omitempty" type:"Struct"`
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

func (s GetNodeResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) GetCategory() *string {
	return s.Category
}

func (s *GetNodeResponseBodyNode) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNodeResponseBodyNode) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetNodeResponseBodyNode) GetExtension() *string {
	return s.Extension
}

func (s *GetNodeResponseBodyNode) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *GetNodeResponseBodyNode) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetNodeResponseBodyNode) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetNodeResponseBodyNode) GetName() *string {
	return s.Name
}

func (s *GetNodeResponseBodyNode) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeResponseBodyNode) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetNodeResponseBodyNode) GetSize() *int64 {
	return s.Size
}

func (s *GetNodeResponseBodyNode) GetStatisticalInfo() *GetNodeResponseBodyNodeStatisticalInfo {
	return s.StatisticalInfo
}

func (s *GetNodeResponseBodyNode) GetType() *string {
	return s.Type
}

func (s *GetNodeResponseBodyNode) GetUrl() *string {
	return s.Url
}

func (s *GetNodeResponseBodyNode) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetNodeResponseBodyNode) SetCategory(v string) *GetNodeResponseBodyNode {
	s.Category = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v string) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetCreatorId(v string) *GetNodeResponseBodyNode {
	s.CreatorId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetExtension(v string) *GetNodeResponseBodyNode {
	s.Extension = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetHasChildren(v bool) *GetNodeResponseBodyNode {
	s.HasChildren = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifiedTime(v string) *GetNodeResponseBodyNode {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifierId(v string) *GetNodeResponseBodyNode {
	s.ModifierId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetNodeId(v string) *GetNodeResponseBodyNode {
	s.NodeId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetPermissionRole(v string) *GetNodeResponseBodyNode {
	s.PermissionRole = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSize(v int64) *GetNodeResponseBodyNode {
	s.Size = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetStatisticalInfo(v *GetNodeResponseBodyNodeStatisticalInfo) *GetNodeResponseBodyNode {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodeResponseBodyNode) SetType(v string) *GetNodeResponseBodyNode {
	s.Type = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetUrl(v string) *GetNodeResponseBodyNode {
	s.Url = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetWorkspaceId(v string) *GetNodeResponseBodyNode {
	s.WorkspaceId = &v
	return s
}

func (s *GetNodeResponseBodyNode) Validate() error {
	if s.StatisticalInfo != nil {
		if err := s.StatisticalInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeResponseBodyNodeStatisticalInfo struct {
	// example:
	//
	// 200
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetNodeResponseBodyNodeStatisticalInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBodyNodeStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNodeStatisticalInfo) GetWordCount() *int64 {
	return s.WordCount
}

func (s *GetNodeResponseBodyNodeStatisticalInfo) SetWordCount(v int64) *GetNodeResponseBodyNodeStatisticalInfo {
	s.WordCount = &v
	return s
}

func (s *GetNodeResponseBodyNodeStatisticalInfo) Validate() error {
	return dara.Validate(s)
}
