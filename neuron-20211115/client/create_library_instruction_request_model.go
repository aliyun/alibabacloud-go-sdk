// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *CreateLibraryInstructionRequest
	GetCompanyId() *int64
	SetDocument(v string) *CreateLibraryInstructionRequest
	GetDocument() *string
	SetId(v int32) *CreateLibraryInstructionRequest
	GetId() *int32
	SetLibraryId(v string) *CreateLibraryInstructionRequest
	GetLibraryId() *string
	SetMarketId(v int64) *CreateLibraryInstructionRequest
	GetMarketId() *int64
}

type CreateLibraryInstructionRequest struct {
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Document  *string `json:"document,omitempty" xml:"document,omitempty"`
	Id        *int32  `json:"id,omitempty" xml:"id,omitempty"`
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	MarketId  *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s CreateLibraryInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryInstructionRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryInstructionRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CreateLibraryInstructionRequest) GetDocument() *string {
	return s.Document
}

func (s *CreateLibraryInstructionRequest) GetId() *int32 {
	return s.Id
}

func (s *CreateLibraryInstructionRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreateLibraryInstructionRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *CreateLibraryInstructionRequest) SetCompanyId(v int64) *CreateLibraryInstructionRequest {
	s.CompanyId = &v
	return s
}

func (s *CreateLibraryInstructionRequest) SetDocument(v string) *CreateLibraryInstructionRequest {
	s.Document = &v
	return s
}

func (s *CreateLibraryInstructionRequest) SetId(v int32) *CreateLibraryInstructionRequest {
	s.Id = &v
	return s
}

func (s *CreateLibraryInstructionRequest) SetLibraryId(v string) *CreateLibraryInstructionRequest {
	s.LibraryId = &v
	return s
}

func (s *CreateLibraryInstructionRequest) SetMarketId(v int64) *CreateLibraryInstructionRequest {
	s.MarketId = &v
	return s
}

func (s *CreateLibraryInstructionRequest) Validate() error {
	return dara.Validate(s)
}
