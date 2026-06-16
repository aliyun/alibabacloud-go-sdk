// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePcaAndExternalCACertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListRequest
	GetCurrentPage() *int32
	SetIdentifiers(v string) *DescribePcaAndExternalCACertificateListRequest
	GetIdentifiers() *string
	SetKeyWord(v string) *DescribePcaAndExternalCACertificateListRequest
	GetKeyWord() *string
	SetShowSize(v int32) *DescribePcaAndExternalCACertificateListRequest
	GetShowSize() *int32
}

type DescribePcaAndExternalCACertificateListRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// One or more certificate identifiers, separated by commas.
	//
	// example:
	//
	// aaa,bbb
	Identifiers *string `json:"Identifiers,omitempty" xml:"Identifiers,omitempty"`
	// The keyword for a fuzzy search on the name, domain name, and SAN fields.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries to return per page. The default value is 50.
	//
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribePcaAndExternalCACertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePcaAndExternalCACertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetIdentifiers() *string {
	return s.Identifiers
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribePcaAndExternalCACertificateListRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetIdentifiers(v string) *DescribePcaAndExternalCACertificateListRequest {
	s.Identifiers = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetKeyWord(v string) *DescribePcaAndExternalCACertificateListRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) SetShowSize(v int32) *DescribePcaAndExternalCACertificateListRequest {
	s.ShowSize = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListRequest) Validate() error {
	return dara.Validate(s)
}
