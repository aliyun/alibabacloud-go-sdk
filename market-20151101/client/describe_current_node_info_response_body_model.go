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
	// The request ID.
	//
	// example:
	//
	// 00eb4de1-6cff-4f56-833e-7b1e070e398d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The process node information returned upon success.
	Result *DescribeCurrentNodeInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCurrentNodeInfoResponseBodyResult struct {
	// Indicates whether rollback to the previous process node is allowed.
	//
	// example:
	//
	// false
	AllowRollbackNode *bool `json:"AllowRollbackNode,omitempty" xml:"AllowRollbackNode,omitempty"`
	// Indicates whether the node is automatically completed.
	//
	// example:
	//
	// false
	AutoFinishNode *bool `json:"AutoFinishNode,omitempty" xml:"AutoFinishNode,omitempty"`
	// The sequence number of the last node.
	//
	// example:
	//
	// 4
	FinalStepNo *int32 `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	// The timeout of the process node. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1588920725000
	GmtExpired *int64 `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The completion time of the process node. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1588920725000
	GmtFinished *int64 `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	// The start time of the process node. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1588834325000
	GmtStart *int64 `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// Indicates whether uploading an attachment is mandatory.
	//
	// example:
	//
	// false
	NeedAttachment *bool `json:"NeedAttachment,omitempty" xml:"NeedAttachment,omitempty"`
	// The ID of the next process node.
	//
	// example:
	//
	// 8473
	NextNodeId *int64 `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	// The ID of the process node.
	//
	// example:
	//
	// 8472
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Waiting for service activation
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The execution status of the node. Valid values:
	//
	// - NotStart: not started.
	//
	// - Starting: in progress.
	//
	// - Cancel: canceled.
	//
	// - Finish: completed.
	//
	// example:
	//
	// Starting
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The role of the operator for the process node. Valid values:
	//
	// - system: system.
	//
	// - custom: user.
	//
	// example:
	//
	// Provider
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	// The ID of the parent process node.
	//
	// example:
	//
	// 0
	ParentNodeId *int64 `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	// The ID of the previous process node.
	//
	// example:
	//
	// 8471
	PreviousNodeId *int64 `json:"PreviousNodeId,omitempty" xml:"PreviousNodeId,omitempty"`
	// The sequence number of the process node.
	//
	// example:
	//
	// 3
	StepNo *int32 `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
	// The form content of the process node in JSON format.
	//
	// Sample code:
	//
	// `[ { "fieldKey": "TFD0", //form key "isRequired": true, //whether the field is required "maxLength": 50, //maximum length "placeHolder": "Name", //form value description "stepNum": 0, //sort value (ascending) "title": "Consignee name", //form title "type": "text" //form type }, { "fieldKey": "TFD1", "isRequired": true, "maxLength": 50, "placeHolder": "Phone number", "stepNum": 1, "title": "Consignee phone number", "type": "text" }, { "fieldKey": "TFD2", "isRequired": true, "maxLength": 50, "placeHolder": "Province, city, district, street, and house number", "stepNum": 2, "title": "Detailed shipping address", "type": "textarea" } ]`.
	//
	// example:
	//
	// [ {"fieldKey": "TFD0","isRequired": true,"maxLength": 50, "placeHolder": "Name","stepNum": 0,"title": "Consignee name","type": "text" },{ "fieldKey": "TFD1","isRequired": true,"maxLength": 50,"placeHolder": "Phone number","stepNum": 1,"title": "Consignee phone number","type": "text"},{"fieldKey": "TFD2","isRequired": true,"maxLength": 50,"placeHolder": "Province, city, district, street, and house number","stepNum": 2,"title": "Detailed shipping address","type": "textarea"}]
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
