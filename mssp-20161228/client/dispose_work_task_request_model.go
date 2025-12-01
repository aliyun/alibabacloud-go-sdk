// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeWorkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperator(v string) *DisposeWorkTaskRequest
	GetOperator() *string
	SetOptRemark(v string) *DisposeWorkTaskRequest
	GetOptRemark() *string
	SetStatus(v int32) *DisposeWorkTaskRequest
	GetStatus() *int32
	SetTaskIds(v string) *DisposeWorkTaskRequest
	GetTaskIds() *string
	SetWorkTaskAnalysisResults(v []*DisposeWorkTaskRequestWorkTaskAnalysisResults) *DisposeWorkTaskRequest
	GetWorkTaskAnalysisResults() []*DisposeWorkTaskRequestWorkTaskAnalysisResults
}

type DisposeWorkTaskRequest struct {
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
	TaskIds                 *string                                          `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	WorkTaskAnalysisResults []*DisposeWorkTaskRequestWorkTaskAnalysisResults `json:"WorkTaskAnalysisResults,omitempty" xml:"WorkTaskAnalysisResults,omitempty" type:"Repeated"`
}

func (s DisposeWorkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DisposeWorkTaskRequest) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskRequest) GetOperator() *string {
	return s.Operator
}

func (s *DisposeWorkTaskRequest) GetOptRemark() *string {
	return s.OptRemark
}

func (s *DisposeWorkTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DisposeWorkTaskRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DisposeWorkTaskRequest) GetWorkTaskAnalysisResults() []*DisposeWorkTaskRequestWorkTaskAnalysisResults {
	return s.WorkTaskAnalysisResults
}

func (s *DisposeWorkTaskRequest) SetOperator(v string) *DisposeWorkTaskRequest {
	s.Operator = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetOptRemark(v string) *DisposeWorkTaskRequest {
	s.OptRemark = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetStatus(v int32) *DisposeWorkTaskRequest {
	s.Status = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetTaskIds(v string) *DisposeWorkTaskRequest {
	s.TaskIds = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetWorkTaskAnalysisResults(v []*DisposeWorkTaskRequestWorkTaskAnalysisResults) *DisposeWorkTaskRequest {
	s.WorkTaskAnalysisResults = v
	return s
}

func (s *DisposeWorkTaskRequest) Validate() error {
	if s.WorkTaskAnalysisResults != nil {
		for _, item := range s.WorkTaskAnalysisResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DisposeWorkTaskRequestWorkTaskAnalysisResults struct {
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DisposeWorkTaskRequestWorkTaskAnalysisResults) String() string {
	return dara.Prettify(s)
}

func (s DisposeWorkTaskRequestWorkTaskAnalysisResults) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskRequestWorkTaskAnalysisResults) GetAnalysisResult() *string {
	return s.AnalysisResult
}

func (s *DisposeWorkTaskRequestWorkTaskAnalysisResults) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DisposeWorkTaskRequestWorkTaskAnalysisResults) SetAnalysisResult(v string) *DisposeWorkTaskRequestWorkTaskAnalysisResults {
	s.AnalysisResult = &v
	return s
}

func (s *DisposeWorkTaskRequestWorkTaskAnalysisResults) SetTaskId(v int64) *DisposeWorkTaskRequestWorkTaskAnalysisResults {
	s.TaskId = &v
	return s
}

func (s *DisposeWorkTaskRequestWorkTaskAnalysisResults) Validate() error {
	return dara.Validate(s)
}
