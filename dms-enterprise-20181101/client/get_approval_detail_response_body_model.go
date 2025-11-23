// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalDetail(v *GetApprovalDetailResponseBodyApprovalDetail) *GetApprovalDetailResponseBody
	GetApprovalDetail() *GetApprovalDetailResponseBodyApprovalDetail
	SetErrorCode(v string) *GetApprovalDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetApprovalDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetApprovalDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApprovalDetailResponseBody
	GetSuccess() *bool
}

type GetApprovalDetailResponseBody struct {
	// The approval details of the ticket.
	ApprovalDetail *GetApprovalDetailResponseBodyApprovalDetail `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// MissingWorkflowInstanceId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// WorkflowInstanceId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66DE630B-ECA1-52A3-9198-602066F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApprovalDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBody) GetApprovalDetail() *GetApprovalDetailResponseBodyApprovalDetail {
	return s.ApprovalDetail
}

func (s *GetApprovalDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetApprovalDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApprovalDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApprovalDetailResponseBody) SetApprovalDetail(v *GetApprovalDetailResponseBodyApprovalDetail) *GetApprovalDetailResponseBody {
	s.ApprovalDetail = v
	return s
}

func (s *GetApprovalDetailResponseBody) SetErrorCode(v string) *GetApprovalDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetErrorMessage(v string) *GetApprovalDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetRequestId(v string) *GetApprovalDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetSuccess(v bool) *GetApprovalDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetApprovalDetailResponseBody) Validate() error {
	if s.ApprovalDetail != nil {
		if err := s.ApprovalDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalDetailResponseBodyApprovalDetail struct {
	// The ID of the approval process.
	//
	// example:
	//
	// 184****
	AuditId *int64 `json:"AuditId,omitempty" xml:"AuditId,omitempty"`
	// The time when the approval process was created.
	//
	// example:
	//
	// 2021-10-29 14:17:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the approver.
	CurrentHandlers *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers `json:"CurrentHandlers,omitempty" xml:"CurrentHandlers,omitempty" type:"Struct"`
	// The description of the approval process.
	//
	// example:
	//
	// [Instance permissions] Application\\<br/>Permission type: ⌈Logon⌋\\<br/>Application period: 30.0 days\\<br/>Background description: [Instance permissions] logon test\\<br/>\\<br/>[Application list]\\<br/>\\<span style=\\"color:red\\">product\\</span> rm-bp144d5ky4l4rli0417\\*\\*\\*\\*.mysql.rds.aliyuncs.com:3306 - PRODUCT\\<br/>
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ticket.
	//
	// example:
	//
	// 384****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The type of the ticket. Valid values:
	//
	// 	- **NDDL**: a schema design ticket
	//
	// 	- **DATA_TRACK**: a data tracking ticket
	//
	// 	- **TABLE_SYNC**: a table synchronization ticket
	//
	// 	- **PERM_APPLY**: a permission application ticket
	//
	// 	- **DATA_EXPORT**: a data export ticket
	//
	// 	- **DATA_CORRECT**: a data change ticket
	//
	// 	- **OWNER_APPLY**: an owner role application ticket
	//
	// 	- **SENSITIVITY**: a column sensitivity level change ticket
	//
	// example:
	//
	// PERM_APPLY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The reasons for the approval.
	ReasonList *GetApprovalDetailResponseBodyApprovalDetailReasonList `json:"ReasonList,omitempty" xml:"ReasonList,omitempty" type:"Struct"`
	// The ID of the workflow template.
	//
	// example:
	//
	// 1234
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Third-party approval flow remarks.
	//
	// example:
	//
	// test
	ThirdpartyWorkflowComment *string `json:"ThirdpartyWorkflowComment,omitempty" xml:"ThirdpartyWorkflowComment,omitempty"`
	// The third-party approval flow link.
	//
	// example:
	//
	// https://xxx
	ThirdpartyWorkflowUrl *string `json:"ThirdpartyWorkflowUrl,omitempty" xml:"ThirdpartyWorkflowUrl,omitempty"`
	// The title of the approval process.
	//
	// example:
	//
	// Permission application ticket - 384\\*\\*\\*\\*
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The approval status of the ticket. Valid values:
	//
	// 	- **AUDITING**: The ticket is being reviewed.
	//
	// 	- **REJECT**: The ticket was rejected.
	//
	// 	- **CANCEL**: The ticket was revoked.
	//
	// 	- **APPROVED**: The ticket was approved.
	//
	// > An approval process contains multiple approval nodes, and this parameter is returned for each approval node.
	//
	// example:
	//
	// APPROVED
	WorkflowInsCode *string `json:"WorkflowInsCode,omitempty" xml:"WorkflowInsCode,omitempty"`
	// The details of approval nodes.
	WorkflowNodes *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
}

func (s GetApprovalDetailResponseBodyApprovalDetail) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetail) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetAuditId() *int64 {
	return s.AuditId
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetCurrentHandlers() *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers {
	return s.CurrentHandlers
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetDescription() *string {
	return s.Description
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetOrderType() *string {
	return s.OrderType
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetReasonList() *GetApprovalDetailResponseBodyApprovalDetailReasonList {
	return s.ReasonList
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetThirdpartyWorkflowComment() *string {
	return s.ThirdpartyWorkflowComment
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetThirdpartyWorkflowUrl() *string {
	return s.ThirdpartyWorkflowUrl
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetTitle() *string {
	return s.Title
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetWorkflowInsCode() *string {
	return s.WorkflowInsCode
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) GetWorkflowNodes() *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes {
	return s.WorkflowNodes
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetAuditId(v int64) *GetApprovalDetailResponseBodyApprovalDetail {
	s.AuditId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetCreateTime(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.CreateTime = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetCurrentHandlers(v *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) *GetApprovalDetailResponseBodyApprovalDetail {
	s.CurrentHandlers = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetDescription(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.Description = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetOrderId(v int64) *GetApprovalDetailResponseBodyApprovalDetail {
	s.OrderId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetOrderType(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.OrderType = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetReasonList(v *GetApprovalDetailResponseBodyApprovalDetailReasonList) *GetApprovalDetailResponseBodyApprovalDetail {
	s.ReasonList = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetTemplateId(v int64) *GetApprovalDetailResponseBodyApprovalDetail {
	s.TemplateId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetThirdpartyWorkflowComment(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.ThirdpartyWorkflowComment = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetThirdpartyWorkflowUrl(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.ThirdpartyWorkflowUrl = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetTitle(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.Title = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetWorkflowInsCode(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.WorkflowInsCode = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetWorkflowNodes(v *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) *GetApprovalDetailResponseBodyApprovalDetail {
	s.WorkflowNodes = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) Validate() error {
	if s.CurrentHandlers != nil {
		if err := s.CurrentHandlers.Validate(); err != nil {
			return err
		}
	}
	if s.ReasonList != nil {
		if err := s.ReasonList.Validate(); err != nil {
			return err
		}
	}
	if s.WorkflowNodes != nil {
		if err := s.WorkflowNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers struct {
	CurrentHandler []*GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler `json:"CurrentHandler,omitempty" xml:"CurrentHandler,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) GetCurrentHandler() []*GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler {
	return s.CurrentHandler
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) SetCurrentHandler(v []*GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers {
	s.CurrentHandler = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) Validate() error {
	if s.CurrentHandler != nil {
		for _, item := range s.CurrentHandler {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler struct {
	// The ID of the user.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// nickName
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) GetId() *int64 {
	return s.Id
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) GetNickName() *string {
	return s.NickName
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) SetId(v int64) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler {
	s.Id = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) SetNickName(v string) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler {
	s.NickName = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) Validate() error {
	return dara.Validate(s)
}

type GetApprovalDetailResponseBodyApprovalDetailReasonList struct {
	Reasons []*string `json:"Reasons,omitempty" xml:"Reasons,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailReasonList) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailReasonList) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailReasonList) GetReasons() []*string {
	return s.Reasons
}

func (s *GetApprovalDetailResponseBodyApprovalDetailReasonList) SetReasons(v []*string) *GetApprovalDetailResponseBodyApprovalDetailReasonList {
	s.Reasons = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailReasonList) Validate() error {
	return dara.Validate(s)
}

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes struct {
	WorkflowNode []*GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) GetWorkflowNode() []*GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	return s.WorkflowNode
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) SetWorkflowNode(v []*GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes {
	s.WorkflowNode = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) Validate() error {
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

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode struct {
	// The IDs of the approvers.
	AuditUserIdList *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList `json:"AuditUserIdList,omitempty" xml:"AuditUserIdList,omitempty" type:"Struct"`
	// The name of the approval node.
	//
	// example:
	//
	// DBA
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The remarks of the approval.
	//
	// example:
	//
	// Reason: Approved
	OperateComment *string `json:"OperateComment,omitempty" xml:"OperateComment,omitempty"`
	// The time when the ticket was submitted.
	//
	// example:
	//
	// 2019-10-15 13:47:54
	OperateTime *string `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The ID of the user who submitted the ticket.
	//
	// example:
	//
	// 1****
	OperatorId *int64 `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The approval status of the ticket. Valid values:
	//
	// 	- **START**: The ticket was submitted.
	//
	// 	- **ERROR**: An error occurred.
	//
	// 	- **AUDITING**: The ticket is being reviewed.
	//
	// 	- **REJECT**: The ticket was rejected.
	//
	// 	- **CANCEL**: The ticket was revoked.
	//
	// 	- **APPROVED**: The ticket was approved.
	//
	// example:
	//
	// APPROVED
	WorkflowInsCode *string `json:"WorkflowInsCode,omitempty" xml:"WorkflowInsCode,omitempty"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetAuditUserIdList() *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList {
	return s.AuditUserIdList
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetOperateComment() *string {
	return s.OperateComment
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetOperateTime() *string {
	return s.OperateTime
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetOperatorId() *int64 {
	return s.OperatorId
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GetWorkflowInsCode() *string {
	return s.WorkflowInsCode
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetAuditUserIdList(v *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.AuditUserIdList = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetNodeName(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperateComment(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperateComment = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperateTime(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperateTime = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperatorId(v int64) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperatorId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetWorkflowInsCode(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.WorkflowInsCode = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) Validate() error {
	if s.AuditUserIdList != nil {
		if err := s.AuditUserIdList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList struct {
	AuditUserIds []*string `json:"AuditUserIds,omitempty" xml:"AuditUserIds,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) GetAuditUserIds() []*string {
	return s.AuditUserIds
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) SetAuditUserIds(v []*string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList {
	s.AuditUserIds = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) Validate() error {
	return dara.Validate(s)
}
