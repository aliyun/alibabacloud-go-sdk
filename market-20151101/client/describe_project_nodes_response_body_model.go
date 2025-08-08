// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeProjectNodesResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeProjectNodesResponseBodyResult) *DescribeProjectNodesResponseBody
	GetResult() []*DescribeProjectNodesResponseBodyResult
	SetSuccess(v bool) *DescribeProjectNodesResponseBody
	GetSuccess() *bool
}

type DescribeProjectNodesResponseBody struct {
	// example:
	//
	// 937fee1f-26bb-4b6e-8def-977a6bdaa1e5
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectNodesResponseBody) GetResult() []*DescribeProjectNodesResponseBodyResult {
	return s.Result
}

func (s *DescribeProjectNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectNodesResponseBody) SetRequestId(v string) *DescribeProjectNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectNodesResponseBody) SetResult(v []*DescribeProjectNodesResponseBodyResult) *DescribeProjectNodesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectNodesResponseBody) SetSuccess(v bool) *DescribeProjectNodesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectNodesResponseBodyResult struct {
	// example:
	//
	// false
	AllowRollbackNode *bool `json:"AllowRollbackNode,omitempty" xml:"AllowRollbackNode,omitempty"`
	// example:
	//
	// false
	AutoFinishNode *bool `json:"AutoFinishNode,omitempty" xml:"AutoFinishNode,omitempty"`
	// example:
	//
	// 4
	FinalStepNo *int32 `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	// example:
	//
	// 1588834325000
	GmtExpired *int64 `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// example:
	//
	// 1588834325000
	GmtFinished *int64 `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	// example:
	//
	// 1588834325000
	GmtStart *int64 `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// example:
	//
	// false
	NeedAttachment *bool `json:"NeedAttachment,omitempty" xml:"NeedAttachment,omitempty"`
	// example:
	//
	// 8472
	NextNodeId *int64 `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	// example:
	//
	// 8471
	NodeId   *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// Finish
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// example:
	//
	// System
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	// example:
	//
	// 0
	ParentNodeId *int64 `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	// example:
	//
	// 8470
	PreviousNodeId *int64 `json:"PreviousNodeId,omitempty" xml:"PreviousNodeId,omitempty"`
	// example:
	//
	// 2
	StepNo       *int32  `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
	TemplateForm *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s DescribeProjectNodesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponseBodyResult) GetAllowRollbackNode() *bool {
	return s.AllowRollbackNode
}

func (s *DescribeProjectNodesResponseBodyResult) GetAutoFinishNode() *bool {
	return s.AutoFinishNode
}

func (s *DescribeProjectNodesResponseBodyResult) GetFinalStepNo() *int32 {
	return s.FinalStepNo
}

func (s *DescribeProjectNodesResponseBodyResult) GetGmtExpired() *int64 {
	return s.GmtExpired
}

func (s *DescribeProjectNodesResponseBodyResult) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *DescribeProjectNodesResponseBodyResult) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *DescribeProjectNodesResponseBodyResult) GetNeedAttachment() *bool {
	return s.NeedAttachment
}

func (s *DescribeProjectNodesResponseBodyResult) GetNextNodeId() *int64 {
	return s.NextNodeId
}

func (s *DescribeProjectNodesResponseBodyResult) GetNodeId() *int64 {
	return s.NodeId
}

func (s *DescribeProjectNodesResponseBodyResult) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeProjectNodesResponseBodyResult) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeProjectNodesResponseBodyResult) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *DescribeProjectNodesResponseBodyResult) GetParentNodeId() *int64 {
	return s.ParentNodeId
}

func (s *DescribeProjectNodesResponseBodyResult) GetPreviousNodeId() *int64 {
	return s.PreviousNodeId
}

func (s *DescribeProjectNodesResponseBodyResult) GetStepNo() *int32 {
	return s.StepNo
}

func (s *DescribeProjectNodesResponseBodyResult) GetTemplateForm() *string {
	return s.TemplateForm
}

func (s *DescribeProjectNodesResponseBodyResult) SetAllowRollbackNode(v bool) *DescribeProjectNodesResponseBodyResult {
	s.AllowRollbackNode = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetAutoFinishNode(v bool) *DescribeProjectNodesResponseBodyResult {
	s.AutoFinishNode = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetFinalStepNo(v int32) *DescribeProjectNodesResponseBodyResult {
	s.FinalStepNo = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtExpired(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtExpired = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtFinished(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtFinished = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtStart(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtStart = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNeedAttachment(v bool) *DescribeProjectNodesResponseBodyResult {
	s.NeedAttachment = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNextNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.NextNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeName(v string) *DescribeProjectNodesResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeStatus(v string) *DescribeProjectNodesResponseBodyResult {
	s.NodeStatus = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetOperatorRole(v string) *DescribeProjectNodesResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetParentNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.ParentNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetPreviousNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.PreviousNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetStepNo(v int32) *DescribeProjectNodesResponseBodyResult {
	s.StepNo = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetTemplateForm(v string) *DescribeProjectNodesResponseBodyResult {
	s.TemplateForm = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
