// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeWorkTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperator(v string) *DisposeWorkTaskShrinkRequest
	GetOperator() *string
	SetOptRemark(v string) *DisposeWorkTaskShrinkRequest
	GetOptRemark() *string
	SetStatus(v int32) *DisposeWorkTaskShrinkRequest
	GetStatus() *int32
	SetTaskIds(v string) *DisposeWorkTaskShrinkRequest
	GetTaskIds() *string
	SetWorkTaskAnalysisResultsShrink(v string) *DisposeWorkTaskShrinkRequest
	GetWorkTaskAnalysisResultsShrink() *string
}

type DisposeWorkTaskShrinkRequest struct {
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// WB01089929
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Operation remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 处理完成
	OptRemark *string `json:"OptRemark,omitempty" xml:"OptRemark,omitempty"`
	// Work order status.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Work order ID, multiple IDs separated by commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10310
	TaskIds                       *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	WorkTaskAnalysisResultsShrink *string `json:"WorkTaskAnalysisResults,omitempty" xml:"WorkTaskAnalysisResults,omitempty"`
}

func (s DisposeWorkTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DisposeWorkTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskShrinkRequest) GetOperator() *string {
	return s.Operator
}

func (s *DisposeWorkTaskShrinkRequest) GetOptRemark() *string {
	return s.OptRemark
}

func (s *DisposeWorkTaskShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DisposeWorkTaskShrinkRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DisposeWorkTaskShrinkRequest) GetWorkTaskAnalysisResultsShrink() *string {
	return s.WorkTaskAnalysisResultsShrink
}

func (s *DisposeWorkTaskShrinkRequest) SetOperator(v string) *DisposeWorkTaskShrinkRequest {
	s.Operator = &v
	return s
}

func (s *DisposeWorkTaskShrinkRequest) SetOptRemark(v string) *DisposeWorkTaskShrinkRequest {
	s.OptRemark = &v
	return s
}

func (s *DisposeWorkTaskShrinkRequest) SetStatus(v int32) *DisposeWorkTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *DisposeWorkTaskShrinkRequest) SetTaskIds(v string) *DisposeWorkTaskShrinkRequest {
	s.TaskIds = &v
	return s
}

func (s *DisposeWorkTaskShrinkRequest) SetWorkTaskAnalysisResultsShrink(v string) *DisposeWorkTaskShrinkRequest {
	s.WorkTaskAnalysisResultsShrink = &v
	return s
}

func (s *DisposeWorkTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
