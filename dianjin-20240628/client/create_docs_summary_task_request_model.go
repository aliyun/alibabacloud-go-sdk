// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocsSummaryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocInfos(v []*CreateDocsSummaryTaskRequestDocInfos) *CreateDocsSummaryTaskRequest
	GetDocInfos() []*CreateDocsSummaryTaskRequestDocInfos
	SetEnableTable(v bool) *CreateDocsSummaryTaskRequest
	GetEnableTable() *bool
	SetInstruction(v string) *CreateDocsSummaryTaskRequest
	GetInstruction() *string
	SetModelId(v string) *CreateDocsSummaryTaskRequest
	GetModelId() *string
}

type CreateDocsSummaryTaskRequest struct {
	// This parameter is required.
	DocInfos []*CreateDocsSummaryTaskRequestDocInfos `json:"docInfos,omitempty" xml:"docInfos,omitempty" type:"Repeated"`
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

func (s CreateDocsSummaryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocsSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskRequest) GetDocInfos() []*CreateDocsSummaryTaskRequestDocInfos {
	return s.DocInfos
}

func (s *CreateDocsSummaryTaskRequest) GetEnableTable() *bool {
	return s.EnableTable
}

func (s *CreateDocsSummaryTaskRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateDocsSummaryTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateDocsSummaryTaskRequest) SetDocInfos(v []*CreateDocsSummaryTaskRequestDocInfos) *CreateDocsSummaryTaskRequest {
	s.DocInfos = v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetEnableTable(v bool) *CreateDocsSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetInstruction(v string) *CreateDocsSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetModelId(v string) *CreateDocsSummaryTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateDocsSummaryTaskRequest) Validate() error {
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

type CreateDocsSummaryTaskRequestDocInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 198386463432
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
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

func (s CreateDocsSummaryTaskRequestDocInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateDocsSummaryTaskRequestDocInfos) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskRequestDocInfos) GetDocId() *string {
	return s.DocId
}

func (s *CreateDocsSummaryTaskRequestDocInfos) GetEndPage() *int32 {
	return s.EndPage
}

func (s *CreateDocsSummaryTaskRequestDocInfos) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreateDocsSummaryTaskRequestDocInfos) GetStartPage() *int32 {
	return s.StartPage
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetDocId(v string) *CreateDocsSummaryTaskRequestDocInfos {
	s.DocId = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetEndPage(v int32) *CreateDocsSummaryTaskRequestDocInfos {
	s.EndPage = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetLibraryId(v string) *CreateDocsSummaryTaskRequestDocInfos {
	s.LibraryId = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetStartPage(v int32) *CreateDocsSummaryTaskRequestDocInfos {
	s.StartPage = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) Validate() error {
	return dara.Validate(s)
}
