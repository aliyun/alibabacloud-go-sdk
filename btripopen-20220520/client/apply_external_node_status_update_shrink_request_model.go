// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyExternalNodeStatusUpdateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *ApplyExternalNodeStatusUpdateShrinkRequest
	GetNodeId() *string
	SetOperationRecordsShrink(v string) *ApplyExternalNodeStatusUpdateShrinkRequest
	GetOperationRecordsShrink() *string
	SetProcessActionResult(v string) *ApplyExternalNodeStatusUpdateShrinkRequest
	GetProcessActionResult() *string
}

type ApplyExternalNodeStatusUpdateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	NodeId                 *string `json:"node_id,omitempty" xml:"node_id,omitempty"`
	OperationRecordsShrink *string `json:"operation_records,omitempty" xml:"operation_records,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_AGREE
	ProcessActionResult *string `json:"process_action_result,omitempty" xml:"process_action_result,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) GetOperationRecordsShrink() *string {
	return s.OperationRecordsShrink
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) GetProcessActionResult() *string {
	return s.ProcessActionResult
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) SetNodeId(v string) *ApplyExternalNodeStatusUpdateShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) SetOperationRecordsShrink(v string) *ApplyExternalNodeStatusUpdateShrinkRequest {
	s.OperationRecordsShrink = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) SetProcessActionResult(v string) *ApplyExternalNodeStatusUpdateShrinkRequest {
	s.ProcessActionResult = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
