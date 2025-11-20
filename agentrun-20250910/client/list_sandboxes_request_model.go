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
	SetStatus(v string) *ListSandboxesRequest
	GetStatus() *string
	SetTemplateName(v string) *ListSandboxesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ListSandboxesRequest
	GetTemplateType() *string
}

type ListSandboxesRequest struct {
	// 当前页码，从1开始计数
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// dnLkmeaJc9vHgbzREh2l0J4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 按模板名称过滤
	//
	// example:
	//
	// templateName
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
