// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSandboxesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSandboxesRequest
	GetNextToken() *string
	SetSandboxId(v string) *ListSandboxesRequest
	GetSandboxId() *string
	SetStatus(v string) *ListSandboxesRequest
	GetStatus() *string
	SetTemplateName(v string) *ListSandboxesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ListSandboxesRequest
	GetTemplateType() *string
}

type ListSandboxesRequest struct {
	// The maximum number of results to return.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next set of results. A non-empty value indicates that more data is available.
	//
	// example:
	//
	// dnLkmeaJc9vHgbzREh2l0J4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 01KMB33KCB3YRYE9C2AJCW5DQK
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
	// Filters the results by status.
	//
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters the results by template name.
	//
	// example:
	//
	// templateName
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Filters the results by template type.
	//
	// example:
	//
	// TASK
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s ListSandboxesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesRequest) GoString() string {
	return s.String()
}

func (s *ListSandboxesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxesRequest) GetSandboxId() *string {
	return s.SandboxId
}

func (s *ListSandboxesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSandboxesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListSandboxesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListSandboxesRequest) SetMaxResults(v int32) *ListSandboxesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxesRequest) SetNextToken(v string) *ListSandboxesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSandboxesRequest) SetSandboxId(v string) *ListSandboxesRequest {
	s.SandboxId = &v
	return s
}

func (s *ListSandboxesRequest) SetStatus(v string) *ListSandboxesRequest {
	s.Status = &v
	return s
}

func (s *ListSandboxesRequest) SetTemplateName(v string) *ListSandboxesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListSandboxesRequest) SetTemplateType(v string) *ListSandboxesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListSandboxesRequest) Validate() error {
	return dara.Validate(s)
}
