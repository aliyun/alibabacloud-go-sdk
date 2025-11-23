// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *PreviewWorkflowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PreviewWorkflowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PreviewWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreviewWorkflowResponseBody
	GetSuccess() *bool
	SetWorkflowDetail(v *PreviewWorkflowResponseBodyWorkflowDetail) *PreviewWorkflowResponseBody
	GetWorkflowDetail() *PreviewWorkflowResponseBodyWorkflowDetail
}

type PreviewWorkflowResponseBody struct {
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
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the workflow.
	WorkflowDetail *PreviewWorkflowResponseBodyWorkflowDetail `json:"WorkflowDetail,omitempty" xml:"WorkflowDetail,omitempty" type:"Struct"`
}

func (s PreviewWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PreviewWorkflowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PreviewWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreviewWorkflowResponseBody) GetWorkflowDetail() *PreviewWorkflowResponseBodyWorkflowDetail {
	return s.WorkflowDetail
}

func (s *PreviewWorkflowResponseBody) SetErrorCode(v string) *PreviewWorkflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PreviewWorkflowResponseBody) SetErrorMessage(v string) *PreviewWorkflowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PreviewWorkflowResponseBody) SetRequestId(v string) *PreviewWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewWorkflowResponseBody) SetSuccess(v bool) *PreviewWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewWorkflowResponseBody) SetWorkflowDetail(v *PreviewWorkflowResponseBodyWorkflowDetail) *PreviewWorkflowResponseBody {
	s.WorkflowDetail = v
	return s
}

func (s *PreviewWorkflowResponseBody) Validate() error {
	if s.WorkflowDetail != nil {
		if err := s.WorkflowDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewWorkflowResponseBodyWorkflowDetail struct {
	// The remarks of the approval template.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the approval template.
	WfCateName *string `json:"WfCateName,omitempty" xml:"WfCateName,omitempty"`
	// The approval nodes.
	WorkflowNodeList *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList `json:"WorkflowNodeList,omitempty" xml:"WorkflowNodeList,omitempty" type:"Struct"`
}

func (s PreviewWorkflowResponseBodyWorkflowDetail) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBodyWorkflowDetail) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) GetComment() *string {
	return s.Comment
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) GetWfCateName() *string {
	return s.WfCateName
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) GetWorkflowNodeList() *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList {
	return s.WorkflowNodeList
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) SetComment(v string) *PreviewWorkflowResponseBodyWorkflowDetail {
	s.Comment = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) SetWfCateName(v string) *PreviewWorkflowResponseBodyWorkflowDetail {
	s.WfCateName = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) SetWorkflowNodeList(v *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) *PreviewWorkflowResponseBodyWorkflowDetail {
	s.WorkflowNodeList = v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetail) Validate() error {
	if s.WorkflowNodeList != nil {
		if err := s.WorkflowNodeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList struct {
	WorkflowNode []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) GetWorkflowNode() []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode {
	return s.WorkflowNode
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) SetWorkflowNode(v []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList {
	s.WorkflowNode = v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeList) Validate() error {
	if s.WorkflowNode != nil {
		for _, item := range s.WorkflowNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode struct {
	// The approvers.
	AuditUserList *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList `json:"AuditUserList,omitempty" xml:"AuditUserList,omitempty" type:"Struct"`
	// The remarks of the approval node.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the approval node.
	//
	// example:
	//
	// Owner
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the approval node.
	//
	// Valid values:
	//
	// 	- USER_LIST: The approval node is created by a user.
	//
	// 	- UNKNOWN: The source of the approval node is unknown.
	//
	// 	- SYS: The approval node is predefined by the system.
	//
	// example:
	//
	// SYS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) GetAuditUserList() *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList {
	return s.AuditUserList
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) GetComment() *string {
	return s.Comment
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) GetNodeName() *string {
	return s.NodeName
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) GetNodeType() *string {
	return s.NodeType
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) SetAuditUserList(v *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode {
	s.AuditUserList = v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) SetComment(v string) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode {
	s.Comment = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) SetNodeName(v string) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) SetNodeType(v string) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode {
	s.NodeType = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode) Validate() error {
	if s.AuditUserList != nil {
		if err := s.AuditUserList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList struct {
	AuditUser []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser `json:"AuditUser,omitempty" xml:"AuditUser,omitempty" type:"Repeated"`
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) GetAuditUser() []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser {
	return s.AuditUser
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) SetAuditUser(v []*PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList {
	s.AuditUser = v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList) Validate() error {
	if s.AuditUser != nil {
		for _, item := range s.AuditUser {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser struct {
	// The nickname of the approver.
	//
	// example:
	//
	// Owner
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The name of the approver.
	//
	// example:
	//
	// db_test
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// The ID of the approver.
	//
	// example:
	//
	// 16***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) String() string {
	return dara.Prettify(s)
}

func (s PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) GoString() string {
	return s.String()
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) GetNickName() *string {
	return s.NickName
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) GetRealName() *string {
	return s.RealName
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) GetUserId() *int64 {
	return s.UserId
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) SetNickName(v string) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser {
	s.NickName = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) SetRealName(v string) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser {
	s.RealName = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) SetUserId(v int64) *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser {
	s.UserId = &v
	return s
}

func (s *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser) Validate() error {
	return dara.Validate(s)
}
