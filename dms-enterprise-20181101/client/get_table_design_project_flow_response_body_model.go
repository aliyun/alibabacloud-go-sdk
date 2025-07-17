// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDesignProjectFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTableDesignProjectFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableDesignProjectFlowResponseBody
	GetErrorMessage() *string
	SetProjectFlow(v *GetTableDesignProjectFlowResponseBodyProjectFlow) *GetTableDesignProjectFlowResponseBody
	GetProjectFlow() *GetTableDesignProjectFlowResponseBodyProjectFlow
	SetRequestId(v string) *GetTableDesignProjectFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableDesignProjectFlowResponseBody
	GetSuccess() *bool
}

type GetTableDesignProjectFlowResponseBody struct {
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
	// The information about the schema design process.
	ProjectFlow *GetTableDesignProjectFlowResponseBodyProjectFlow `json:"ProjectFlow,omitempty" xml:"ProjectFlow,omitempty" type:"Struct"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// B5FD0BC8-2D90-4478-B8EC-A0E92E0B1773
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableDesignProjectFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableDesignProjectFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableDesignProjectFlowResponseBody) GetProjectFlow() *GetTableDesignProjectFlowResponseBodyProjectFlow {
	return s.ProjectFlow
}

func (s *GetTableDesignProjectFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableDesignProjectFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableDesignProjectFlowResponseBody) SetErrorCode(v string) *GetTableDesignProjectFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBody) SetErrorMessage(v string) *GetTableDesignProjectFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBody) SetProjectFlow(v *GetTableDesignProjectFlowResponseBodyProjectFlow) *GetTableDesignProjectFlowResponseBody {
	s.ProjectFlow = v
	return s
}

func (s *GetTableDesignProjectFlowResponseBody) SetRequestId(v string) *GetTableDesignProjectFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBody) SetSuccess(v bool) *GetTableDesignProjectFlowResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableDesignProjectFlowResponseBodyProjectFlow struct {
	// The position of the current node in the process.
	//
	// example:
	//
	// 2
	CurrentPosition *int32 `json:"CurrentPosition,omitempty" xml:"CurrentPosition,omitempty"`
	// The nodes in the process.
	FlowNodeArray []*GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray `json:"FlowNodeArray,omitempty" xml:"FlowNodeArray,omitempty" type:"Repeated"`
	// The description of the security rule set that is applied to the schema design ticket.
	//
	// example:
	//
	// mysq_test
	RuleComment *string `json:"RuleComment,omitempty" xml:"RuleComment,omitempty"`
	// The name of the security rule set that is applied to the schema design ticket.
	//
	// example:
	//
	// mysql default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetTableDesignProjectFlowResponseBodyProjectFlow) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectFlowResponseBodyProjectFlow) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) GetCurrentPosition() *int32 {
	return s.CurrentPosition
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) GetFlowNodeArray() []*GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	return s.FlowNodeArray
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) GetRuleComment() *string {
	return s.RuleComment
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) GetRuleName() *string {
	return s.RuleName
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) SetCurrentPosition(v int32) *GetTableDesignProjectFlowResponseBodyProjectFlow {
	s.CurrentPosition = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) SetFlowNodeArray(v []*GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) *GetTableDesignProjectFlowResponseBodyProjectFlow {
	s.FlowNodeArray = v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) SetRuleComment(v string) *GetTableDesignProjectFlowResponseBodyProjectFlow {
	s.RuleComment = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) SetRuleName(v string) *GetTableDesignProjectFlowResponseBodyProjectFlow {
	s.RuleName = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlow) Validate() error {
	return dara.Validate(s)
}

type GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray struct {
	// Indicates whether the ticket can be returned to the schema design node. Valid values:
	//
	// 	- **1**: The ticket can be returned to the schema design node.
	//
	// 	- **0**: The ticket cannot be returned to the schema design node. This value can be returned only for the PUBLISH node.
	//
	// example:
	//
	// 1
	BackToDesign *bool `json:"BackToDesign,omitempty" xml:"BackToDesign,omitempty"`
	// Indicates whether the current node can be skipped. Valid values:
	//
	// 	- **1**: The current node can be skipped.
	//
	// 	- **0**: The current node cannot be skipped. This value can be returned only for the PUBLISH node.
	//
	// example:
	//
	// 0
	CanSkip *bool `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	// The role of the node in the process.
	//
	// 	- START: The ticket was created.
	//
	// 	- DESIGN: The schema is being created.
	//
	// 	- PUBLISH: The schema is published.
	//
	// 	- END: The ticket is complete.
	//
	// example:
	//
	// DESIGN
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// The title of the node.
	NodeTitle *string `json:"NodeTitle,omitempty" xml:"NodeTitle,omitempty"`
	// The position of the node in the process. The value starts from 1.
	//
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Indicates whether the node is the anchor node. A schema design process has only one anchor node, on which the schema is published. After the schema is published on the anchor node, a post-publish image is generated and the DDL metadata lock is released.
	//
	// example:
	//
	// false
	PublishAnchor *bool `json:"PublishAnchor,omitempty" xml:"PublishAnchor,omitempty"`
	// The available publishing strategies.
	PublishStrategies []*string `json:"PublishStrategies,omitempty" xml:"PublishStrategies,omitempty" type:"Repeated"`
}

func (s GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) String() string {
	return dara.Prettify(s)
}

func (s GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GoString() string {
	return s.String()
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetBackToDesign() *bool {
	return s.BackToDesign
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetCanSkip() *bool {
	return s.CanSkip
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetNodeRole() *string {
	return s.NodeRole
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetNodeTitle() *string {
	return s.NodeTitle
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetPosition() *int32 {
	return s.Position
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetPublishAnchor() *bool {
	return s.PublishAnchor
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) GetPublishStrategies() []*string {
	return s.PublishStrategies
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetBackToDesign(v bool) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.BackToDesign = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetCanSkip(v bool) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.CanSkip = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetNodeRole(v string) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.NodeRole = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetNodeTitle(v string) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.NodeTitle = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetPosition(v int32) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.Position = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetPublishAnchor(v bool) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.PublishAnchor = &v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) SetPublishStrategies(v []*string) *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray {
	s.PublishStrategies = v
	return s
}

func (s *GetTableDesignProjectFlowResponseBodyProjectFlowFlowNodeArray) Validate() error {
	return dara.Validate(s)
}
