// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyExternalNodeStatusUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *ApplyExternalNodeStatusUpdateRequest
	GetNodeId() *string
	SetOperationRecords(v []*ApplyExternalNodeStatusUpdateRequestOperationRecords) *ApplyExternalNodeStatusUpdateRequest
	GetOperationRecords() []*ApplyExternalNodeStatusUpdateRequestOperationRecords
	SetProcessActionResult(v string) *ApplyExternalNodeStatusUpdateRequest
	GetProcessActionResult() *string
}

type ApplyExternalNodeStatusUpdateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	NodeId           *string                                                 `json:"node_id,omitempty" xml:"node_id,omitempty"`
	OperationRecords []*ApplyExternalNodeStatusUpdateRequestOperationRecords `json:"operation_records,omitempty" xml:"operation_records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_AGREE
	ProcessActionResult *string `json:"process_action_result,omitempty" xml:"process_action_result,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ApplyExternalNodeStatusUpdateRequest) GetOperationRecords() []*ApplyExternalNodeStatusUpdateRequestOperationRecords {
	return s.OperationRecords
}

func (s *ApplyExternalNodeStatusUpdateRequest) GetProcessActionResult() *string {
	return s.ProcessActionResult
}

func (s *ApplyExternalNodeStatusUpdateRequest) SetNodeId(v string) *ApplyExternalNodeStatusUpdateRequest {
	s.NodeId = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequest) SetOperationRecords(v []*ApplyExternalNodeStatusUpdateRequestOperationRecords) *ApplyExternalNodeStatusUpdateRequest {
	s.OperationRecords = v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequest) SetProcessActionResult(v string) *ApplyExternalNodeStatusUpdateRequest {
	s.ProcessActionResult = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequest) Validate() error {
	return dara.Validate(s)
}

type ApplyExternalNodeStatusUpdateRequestOperationRecords struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// 2023-05-28 11:33:28
	OperateTime  *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	OperatorName *string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// example:
	//
	// AGREE
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// PROCESS_APPROVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateRequestOperationRecords) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateRequestOperationRecords) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) GetComment() *string {
	return s.Comment
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) GetOperateTime() *string {
	return s.OperateTime
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) GetResult() *string {
	return s.Result
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) GetType() *string {
	return s.Type
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) SetComment(v string) *ApplyExternalNodeStatusUpdateRequestOperationRecords {
	s.Comment = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) SetOperateTime(v string) *ApplyExternalNodeStatusUpdateRequestOperationRecords {
	s.OperateTime = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) SetOperatorName(v string) *ApplyExternalNodeStatusUpdateRequestOperationRecords {
	s.OperatorName = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) SetResult(v string) *ApplyExternalNodeStatusUpdateRequestOperationRecords {
	s.Result = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) SetType(v string) *ApplyExternalNodeStatusUpdateRequestOperationRecords {
	s.Type = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateRequestOperationRecords) Validate() error {
	return dara.Validate(s)
}
