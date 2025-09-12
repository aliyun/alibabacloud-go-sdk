// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPcaCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPcaCaCertificateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPcaCaCertificateRequest
	GetNextToken() *string
}

type ListPcaCaCertificateRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPcaCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPcaCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListPcaCaCertificateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPcaCaCertificateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPcaCaCertificateRequest) SetMaxResults(v int32) *ListPcaCaCertificateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPcaCaCertificateRequest) SetNextToken(v string) *ListPcaCaCertificateRequest {
	s.NextToken = &v
	return s
}

func (s *ListPcaCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
