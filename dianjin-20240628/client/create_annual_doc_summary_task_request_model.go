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
	// The list of analysis years.
	//
	// This parameter is required.
	AnaYears []*int32 `json:"anaYears,omitempty" xml:"anaYears,omitempty" type:"Repeated"`
	// The list of document information.
	//
	// This parameter is required.
	DocInfos []*CreateAnnualDocSummaryTaskRequestDocInfos `json:"docInfos,omitempty" xml:"docInfos,omitempty" type:"Repeated"`
	// Specifies whether to enable tables. Default value: true.
	//
	// example:
	//
	// true
	EnableTable *bool `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	// The instruction.
	//
	// example:
	//
	// You are a senior securities researcher conducting performance analysis on listed companies for the year XX. Based on the reference information, provide a detailed analysis covering the following aspects:
	//
	// 1. Overall performance changes, including detailed metrics such as revenue and profit.
	//
	// 2. Specific reasons for performance changes, including changes in each business segment.
	//
	// Strictly output only the information for the year XX
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The model ID.
	//
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
	if s.DocInfos != nil {
		for _, item := range s.DocInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAnnualDocSummaryTaskRequestDocInfos struct {
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 198386463432
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// The document year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023
	DocYear *int32 `json:"docYear,omitempty" xml:"docYear,omitempty"`
	// The end page.
	//
	// example:
	//
	// 2
	EndPage *int32 `json:"endPage,omitempty" xml:"endPage,omitempty"`
	// The document library ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rdxrmo6amk
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// The start page.
	//
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
