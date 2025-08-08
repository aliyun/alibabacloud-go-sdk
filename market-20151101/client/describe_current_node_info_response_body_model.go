// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentNodeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCurrentNodeInfoResponseBody
	GetRequestId() *string
	SetResult(v *DescribeCurrentNodeInfoResponseBodyResult) *DescribeCurrentNodeInfoResponseBody
	GetResult() *DescribeCurrentNodeInfoResponseBodyResult
	SetSuccess(v bool) *DescribeCurrentNodeInfoResponseBody
	GetSuccess() *bool
}

type DescribeCurrentNodeInfoResponseBody struct {
	// example:
	//
	// 00eb4de1-6cff-4f56-833e-7b1e070e398d
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeCurrentNodeInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCurrentNodeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCurrentNodeInfoResponseBody) GetResult() *DescribeCurrentNodeInfoResponseBodyResult {
	return s.Result
}

func (s *DescribeCurrentNodeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCurrentNodeInfoResponseBody) SetRequestId(v string) *DescribeCurrentNodeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBody) SetResult(v *DescribeCurrentNodeInfoResponseBodyResult) *DescribeCurrentNodeInfoResponseBody {
	s.Result = v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBody) SetSuccess(v bool) *DescribeCurrentNodeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCurrentNodeInfoResponseBodyResult struct {
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
	// 1588920725000
	GmtExpired *int64 `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// example:
	//
	// 1588920725000
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
	// 8473
	NextNodeId *int64 `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	// example:
	//
	// 8472
	NodeId   *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// Starting
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// example:
	//
	// Provider
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	// example:
	//
	// 0
	ParentNodeId *int64 `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	// example:
	//
	// 8471
	PreviousNodeId *int64 `json:"PreviousNodeId,omitempty" xml:"PreviousNodeId,omitempty"`
	// example:
	//
	// 3
	StepNo       *int32  `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
	TemplateForm *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s DescribeCurrentNodeInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetAllowRollbackNode() *bool {
	return s.AllowRollbackNode
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetAutoFinishNode() *bool {
	return s.AutoFinishNode
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetFinalStepNo() *int32 {
	return s.FinalStepNo
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetGmtExpired() *int64 {
	return s.GmtExpired
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetNeedAttachment() *bool {
	return s.NeedAttachment
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetNextNodeId() *int64 {
	return s.NextNodeId
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetNodeId() *int64 {
	return s.NodeId
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetParentNodeId() *int64 {
	return s.ParentNodeId
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetPreviousNodeId() *int64 {
	return s.PreviousNodeId
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetStepNo() *int32 {
	return s.StepNo
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) GetTemplateForm() *string {
	return s.TemplateForm
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetAllowRollbackNode(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.AllowRollbackNode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetAutoFinishNode(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.AutoFinishNode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetFinalStepNo(v int32) *DescribeCurrentNodeInfoResponseBodyResult {
	s.FinalStepNo = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtExpired(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtExpired = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtFinished(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtFinished = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtStart(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtStart = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNeedAttachment(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NeedAttachment = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNextNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NextNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeName(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeStatus(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeStatus = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetOperatorRole(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetParentNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.ParentNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetPreviousNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.PreviousNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetStepNo(v int32) *DescribeCurrentNodeInfoResponseBodyResult {
	s.StepNo = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetTemplateForm(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.TemplateForm = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
