// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFinReportSummaryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *CreateFinReportSummaryTaskRequest
	GetDocId() *string
	SetEnableTable(v bool) *CreateFinReportSummaryTaskRequest
	GetEnableTable() *bool
	SetEndPage(v int32) *CreateFinReportSummaryTaskRequest
	GetEndPage() *int32
	SetInstruction(v string) *CreateFinReportSummaryTaskRequest
	GetInstruction() *string
	SetLibraryId(v string) *CreateFinReportSummaryTaskRequest
	GetLibraryId() *string
	SetModelId(v string) *CreateFinReportSummaryTaskRequest
	GetModelId() *string
	SetStartPage(v int32) *CreateFinReportSummaryTaskRequest
	GetStartPage() *int32
	SetTaskType(v string) *CreateFinReportSummaryTaskRequest
	GetTaskType() *string
}

type CreateFinReportSummaryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableTable *bool `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	// example:
	//
	// 10
	EndPage     *int32  `json:"endPage,omitempty" xml:"endPage,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
	// example:
	//
	// custom
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateFinReportSummaryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFinReportSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskRequest) GetDocId() *string {
	return s.DocId
}

func (s *CreateFinReportSummaryTaskRequest) GetEnableTable() *bool {
	return s.EnableTable
}

func (s *CreateFinReportSummaryTaskRequest) GetEndPage() *int32 {
	return s.EndPage
}

func (s *CreateFinReportSummaryTaskRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateFinReportSummaryTaskRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreateFinReportSummaryTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateFinReportSummaryTaskRequest) GetStartPage() *int32 {
	return s.StartPage
}

func (s *CreateFinReportSummaryTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateFinReportSummaryTaskRequest) SetDocId(v string) *CreateFinReportSummaryTaskRequest {
	s.DocId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEnableTable(v bool) *CreateFinReportSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEndPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetInstruction(v string) *CreateFinReportSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetLibraryId(v string) *CreateFinReportSummaryTaskRequest {
	s.LibraryId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetModelId(v string) *CreateFinReportSummaryTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetStartPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetTaskType(v string) *CreateFinReportSummaryTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) Validate() error {
	return dara.Validate(s)
}
