// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListWorkFlowTemplatesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListWorkFlowTemplatesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListWorkFlowTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkFlowTemplatesResponseBody
	GetSuccess() *bool
	SetWorkFlowTemplates(v *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) *ListWorkFlowTemplatesResponseBody
	GetWorkFlowTemplates() *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates
}

type ListWorkFlowTemplatesResponseBody struct {
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
	// 41067071-0243-4AAB-B3CF-4DE6D54F53B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success           *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	WorkFlowTemplates *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates `json:"WorkFlowTemplates,omitempty" xml:"WorkFlowTemplates,omitempty" type:"Struct"`
}

func (s ListWorkFlowTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkFlowTemplatesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListWorkFlowTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkFlowTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkFlowTemplatesResponseBody) GetWorkFlowTemplates() *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates {
	return s.WorkFlowTemplates
}

func (s *ListWorkFlowTemplatesResponseBody) SetErrorCode(v string) *ListWorkFlowTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetErrorMessage(v string) *ListWorkFlowTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetRequestId(v string) *ListWorkFlowTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetSuccess(v bool) *ListWorkFlowTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetWorkFlowTemplates(v *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) *ListWorkFlowTemplatesResponseBody {
	s.WorkFlowTemplates = v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) Validate() error {
	if s.WorkFlowTemplates != nil {
		if err := s.WorkFlowTemplates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplates struct {
	WorkFlowTemplate []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate `json:"WorkFlowTemplate,omitempty" xml:"WorkFlowTemplate,omitempty" type:"Repeated"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) GetWorkFlowTemplate() []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	return s.WorkFlowTemplate
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) SetWorkFlowTemplate(v []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates {
	s.WorkFlowTemplate = v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) Validate() error {
	if s.WorkFlowTemplate != nil {
		for _, item := range s.WorkFlowTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate struct {
	Comment       *string                                                                          `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateUserId  *int64                                                                           `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	Enabled       *string                                                                          `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	IsSystem      *int32                                                                           `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	TemplateId    *int64                                                                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName  *string                                                                          `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	WorkflowNodes *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetComment() *string {
	return s.Comment
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetEnabled() *string {
	return s.Enabled
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetIsSystem() *int32 {
	return s.IsSystem
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GetWorkflowNodes() *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes {
	return s.WorkflowNodes
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetComment(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetCreateUserId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.CreateUserId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetEnabled(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.Enabled = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetIsSystem(v int32) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.IsSystem = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetTemplateId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetTemplateName(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.TemplateName = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetWorkflowNodes(v *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.WorkflowNodes = v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) Validate() error {
	if s.WorkflowNodes != nil {
		if err := s.WorkflowNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes struct {
	WorkflowNode []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) GetWorkflowNode() []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	return s.WorkflowNode
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) SetWorkflowNode(v []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes {
	s.WorkflowNode = v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) Validate() error {
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

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode struct {
	Comment      *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateUserId *int64  `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	NodeId       *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Position     *int32  `json:"Position,omitempty" xml:"Position,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetComment() *string {
	return s.Comment
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetNodeName() *string {
	return s.NodeName
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetNodeType() *string {
	return s.NodeType
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetPosition() *int32 {
	return s.Position
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetComment(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetCreateUserId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.CreateUserId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeName(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeType(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeType = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetPosition(v int32) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.Position = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetTemplateId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.TemplateId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) Validate() error {
	return dara.Validate(s)
}
