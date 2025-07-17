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
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
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
	return dara.Validate(s)
}

type PreviewWorkflowResponseBodyWorkflowDetail struct {
	Comment          *string                                                    `json:"Comment,omitempty" xml:"Comment,omitempty"`
	WfCateName       *string                                                    `json:"WfCateName,omitempty" xml:"WfCateName,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNode struct {
	AuditUserList *PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserList `json:"AuditUserList,omitempty" xml:"AuditUserList,omitempty" type:"Struct"`
	Comment       *string                                                                             `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// Owner
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type PreviewWorkflowResponseBodyWorkflowDetailWorkflowNodeListWorkflowNodeAuditUserListAuditUser struct {
	// example:
	//
	// Owner
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// db_test
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
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
