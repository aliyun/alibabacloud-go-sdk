// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListWorkFlowNodesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListWorkFlowNodesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListWorkFlowNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkFlowNodesResponseBody
	GetSuccess() *bool
	SetWorkflowNodes(v *ListWorkFlowNodesResponseBodyWorkflowNodes) *ListWorkFlowNodesResponseBody
	GetWorkflowNodes() *ListWorkFlowNodesResponseBodyWorkflowNodes
}

type ListWorkFlowNodesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CFD8FE00-36D9-4C1B-940D-65A7B73D9066
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of approval nodes.
	WorkflowNodes *ListWorkFlowNodesResponseBodyWorkflowNodes `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
}

func (s ListWorkFlowNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkFlowNodesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListWorkFlowNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkFlowNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkFlowNodesResponseBody) GetWorkflowNodes() *ListWorkFlowNodesResponseBodyWorkflowNodes {
	return s.WorkflowNodes
}

func (s *ListWorkFlowNodesResponseBody) SetErrorCode(v string) *ListWorkFlowNodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetErrorMessage(v string) *ListWorkFlowNodesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetRequestId(v string) *ListWorkFlowNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetSuccess(v bool) *ListWorkFlowNodesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetWorkflowNodes(v *ListWorkFlowNodesResponseBodyWorkflowNodes) *ListWorkFlowNodesResponseBody {
	s.WorkflowNodes = v
	return s
}

func (s *ListWorkFlowNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWorkFlowNodesResponseBodyWorkflowNodes struct {
	WorkflowNode []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodes) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodes) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodes) GetWorkflowNode() []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	return s.WorkflowNode
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodes) SetWorkflowNode(v []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) *ListWorkFlowNodesResponseBodyWorkflowNodes {
	s.WorkflowNode = v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodes) Validate() error {
	return dara.Validate(s)
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode struct {
	// The details about approvers.
	AuditUsers *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers `json:"AuditUsers,omitempty" xml:"AuditUsers,omitempty" type:"Struct"`
	// The description of the approval template.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the creator. This ID is different from the ID of the Alibaba Cloud account of the creator.
	//
	// example:
	//
	// 123
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The name of the user who creates the approval node.
	//
	// example:
	//
	// test
	CreateUserNickName *string `json:"CreateUserNickName,omitempty" xml:"CreateUserNickName,omitempty"`
	// The ID of the approval node.
	//
	// example:
	//
	// 123
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the approval node.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the approval node. Valid values:
	//
	// 	- SYS: The approval node is predefined by the system.
	//
	// 	- USER_LIST: The approval node is created by a user.
	//
	// example:
	//
	// SYS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetAuditUsers() *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers {
	return s.AuditUsers
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetComment() *string {
	return s.Comment
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetCreateUserNickName() *string {
	return s.CreateUserNickName
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetNodeName() *string {
	return s.NodeName
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GetNodeType() *string {
	return s.NodeType
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetAuditUsers(v *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.AuditUsers = v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetComment(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetCreateUserId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.CreateUserId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetCreateUserNickName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.CreateUserNickName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeType(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeType = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) Validate() error {
	return dara.Validate(s)
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers struct {
	AuditUser []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser `json:"AuditUser,omitempty" xml:"AuditUser,omitempty" type:"Repeated"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) GetAuditUser() []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	return s.AuditUser
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) SetAuditUser(v []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers {
	s.AuditUser = v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) Validate() error {
	return dara.Validate(s)
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser struct {
	// The nickname of the approver.
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The real name of the approver.
	//
	// example:
	//
	// test
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// The ID of the approver. The ID is different from the ID of the Alibaba Cloud account of the approver.
	//
	// example:
	//
	// 123
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) GetNickName() *string {
	return s.NickName
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) GetRealName() *string {
	return s.RealName
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) GetUserId() *int64 {
	return s.UserId
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetNickName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.NickName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetRealName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.RealName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetUserId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.UserId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) Validate() error {
	return dara.Validate(s)
}
