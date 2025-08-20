// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticKey(v string) *ListDiagnosticsRequest
	GetDiagnosticKey() *string
	SetDiagnosticProduct(v string) *ListDiagnosticsRequest
	GetDiagnosticProduct() *string
	SetMaxResults(v string) *ListDiagnosticsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListDiagnosticsRequest
	GetNextToken() *string
	SetStatus(v string) *ListDiagnosticsRequest
	GetStatus() *string
}

type ListDiagnosticsRequest struct {
	// The keyword in the diagnosis.
	//
	// example:
	//
	// 2829A772-B154-5A0A-B61B-DEE8A9EE8A5D
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The product that is diagnosed.
	//
	// example:
	//
	// ros
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The maximum number of results to be returned in a single call when NextToken is used for the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f01****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The diagnosis status.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsRequest) GetDiagnosticKey() *string {
	return s.DiagnosticKey
}

func (s *ListDiagnosticsRequest) GetDiagnosticProduct() *string {
	return s.DiagnosticProduct
}

func (s *ListDiagnosticsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListDiagnosticsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosticsRequest) SetDiagnosticKey(v string) *ListDiagnosticsRequest {
	s.DiagnosticKey = &v
	return s
}

func (s *ListDiagnosticsRequest) SetDiagnosticProduct(v string) *ListDiagnosticsRequest {
	s.DiagnosticProduct = &v
	return s
}

func (s *ListDiagnosticsRequest) SetMaxResults(v string) *ListDiagnosticsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticsRequest) SetNextToken(v string) *ListDiagnosticsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticsRequest) SetStatus(v string) *ListDiagnosticsRequest {
	s.Status = &v
	return s
}

func (s *ListDiagnosticsRequest) Validate() error {
	return dara.Validate(s)
}
