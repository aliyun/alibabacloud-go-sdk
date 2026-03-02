// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryInstruction interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *LibraryInstruction
	GetCompanyId() *int64
	SetDocument(v string) *LibraryInstruction
	GetDocument() *string
	SetId(v int64) *LibraryInstruction
	GetId() *int64
	SetLibraryId(v string) *LibraryInstruction
	GetLibraryId() *string
	SetMarketId(v int64) *LibraryInstruction
	GetMarketId() *int64
	SetRequestId(v string) *LibraryInstruction
	GetRequestId() *string
}

type LibraryInstruction struct {
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Document  *string `json:"document,omitempty" xml:"document,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	MarketId  *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LibraryInstruction) String() string {
	return dara.Prettify(s)
}

func (s LibraryInstruction) GoString() string {
	return s.String()
}

func (s *LibraryInstruction) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *LibraryInstruction) GetDocument() *string {
	return s.Document
}

func (s *LibraryInstruction) GetId() *int64 {
	return s.Id
}

func (s *LibraryInstruction) GetLibraryId() *string {
	return s.LibraryId
}

func (s *LibraryInstruction) GetMarketId() *int64 {
	return s.MarketId
}

func (s *LibraryInstruction) GetRequestId() *string {
	return s.RequestId
}

func (s *LibraryInstruction) SetCompanyId(v int64) *LibraryInstruction {
	s.CompanyId = &v
	return s
}

func (s *LibraryInstruction) SetDocument(v string) *LibraryInstruction {
	s.Document = &v
	return s
}

func (s *LibraryInstruction) SetId(v int64) *LibraryInstruction {
	s.Id = &v
	return s
}

func (s *LibraryInstruction) SetLibraryId(v string) *LibraryInstruction {
	s.LibraryId = &v
	return s
}

func (s *LibraryInstruction) SetMarketId(v int64) *LibraryInstruction {
	s.MarketId = &v
	return s
}

func (s *LibraryInstruction) SetRequestId(v string) *LibraryInstruction {
	s.RequestId = &v
	return s
}

func (s *LibraryInstruction) Validate() error {
	return dara.Validate(s)
}
