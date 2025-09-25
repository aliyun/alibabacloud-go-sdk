// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnualDocSummaryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnaYears(v []*int32) *CreateAnnualDocSummaryTaskRequest
	GetAnaYears() []*int32
	SetDocInfos(v []*CreateAnnualDocSummaryTaskRequestDocInfos) *CreateAnnualDocSummaryTaskRequest
	GetDocInfos() []*CreateAnnualDocSummaryTaskRequestDocInfos
	SetEnableTable(v bool) *CreateAnnualDocSummaryTaskRequest
	GetEnableTable() *bool
	SetInstruction(v string) *CreateAnnualDocSummaryTaskRequest
	GetInstruction() *string
	SetModelId(v string) *CreateAnnualDocSummaryTaskRequest
	GetModelId() *string
}

type CreateAnnualDocSummaryTaskRequest struct {
	// This parameter is required.
	AnaYears []*int32 `json:"anaYears,omitempty" xml:"anaYears,omitempty" type:"Repeated"`
	// This parameter is required.
	DocInfos []*CreateAnnualDocSummaryTaskRequestDocInfos `json:"docInfos,omitempty" xml:"docInfos,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableTable *bool   `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s CreateAnnualDocSummaryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskRequest) GetAnaYears() []*int32 {
	return s.AnaYears
}

func (s *CreateAnnualDocSummaryTaskRequest) GetDocInfos() []*CreateAnnualDocSummaryTaskRequestDocInfos {
	return s.DocInfos
}

func (s *CreateAnnualDocSummaryTaskRequest) GetEnableTable() *bool {
	return s.EnableTable
}

func (s *CreateAnnualDocSummaryTaskRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateAnnualDocSummaryTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateAnnualDocSummaryTaskRequest) SetAnaYears(v []*int32) *CreateAnnualDocSummaryTaskRequest {
	s.AnaYears = v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetDocInfos(v []*CreateAnnualDocSummaryTaskRequestDocInfos) *CreateAnnualDocSummaryTaskRequest {
	s.DocInfos = v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetEnableTable(v bool) *CreateAnnualDocSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetInstruction(v string) *CreateAnnualDocSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetModelId(v string) *CreateAnnualDocSummaryTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAnnualDocSummaryTaskRequestDocInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 198386463432
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023
	DocYear *int32 `json:"docYear,omitempty" xml:"docYear,omitempty"`
	// example:
	//
	// 2
	EndPage *int32 `json:"endPage,omitempty" xml:"endPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rdxrmo6amk
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
}

func (s CreateAnnualDocSummaryTaskRequestDocInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskRequestDocInfos) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) GetDocId() *string {
	return s.DocId
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) GetDocYear() *int32 {
	return s.DocYear
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) GetEndPage() *int32 {
	return s.EndPage
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) GetStartPage() *int32 {
	return s.StartPage
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetDocId(v string) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.DocId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetDocYear(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.DocYear = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetEndPage(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.EndPage = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetLibraryId(v string) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.LibraryId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetStartPage(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.StartPage = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) Validate() error {
	return dara.Validate(s)
}
