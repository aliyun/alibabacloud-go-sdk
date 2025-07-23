// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListCsrRequest
	GetAlgorithm() *string
	SetCurrentPage(v int64) *ListCsrRequest
	GetCurrentPage() *int64
	SetKeyWord(v string) *ListCsrRequest
	GetKeyWord() *string
	SetShowSize(v int64) *ListCsrRequest
	GetShowSize() *int64
}

type ListCsrRequest struct {
	// The algorithm. Valid values: RSA, ECC, and SM2.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test_name
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCsrRequest) GoString() string {
	return s.String()
}

func (s *ListCsrRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCsrRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCsrRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListCsrRequest) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCsrRequest) SetAlgorithm(v string) *ListCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *ListCsrRequest) SetCurrentPage(v int64) *ListCsrRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCsrRequest) SetKeyWord(v string) *ListCsrRequest {
	s.KeyWord = &v
	return s
}

func (s *ListCsrRequest) SetShowSize(v int64) *ListCsrRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCsrRequest) Validate() error {
	return dara.Validate(s)
}
