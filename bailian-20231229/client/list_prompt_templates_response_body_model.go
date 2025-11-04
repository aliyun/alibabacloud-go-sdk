// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPromptTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPromptTemplatesResponseBody
	GetNextToken() *string
	SetPromptTemplates(v []*ListPromptTemplatesResponseBodyPromptTemplates) *ListPromptTemplatesResponseBody
	GetPromptTemplates() []*ListPromptTemplatesResponseBodyPromptTemplates
	SetRequestId(v string) *ListPromptTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPromptTemplatesResponseBody
	GetTotalCount() *int32
	SetWorkspaceId(v string) *ListPromptTemplatesResponseBody
	GetWorkspaceId() *string
}

type ListPromptTemplatesResponseBody struct {
	// The maximum number of returned entries.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that determines the start position of the next query.
	//
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The templates.
	PromptTemplates []*ListPromptTemplatesResponseBodyPromptTemplates `json:"promptTemplates,omitempty" xml:"promptTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FE9B6CBF-47E6-5D76-9C5D-B814DD5AB071
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// llm-us9hjmt32nysdxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListPromptTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromptTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromptTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromptTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromptTemplatesResponseBody) GetPromptTemplates() []*ListPromptTemplatesResponseBodyPromptTemplates {
	return s.PromptTemplates
}

func (s *ListPromptTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromptTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromptTemplatesResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPromptTemplatesResponseBody) SetMaxResults(v int32) *ListPromptTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPromptTemplatesResponseBody) SetNextToken(v string) *ListPromptTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPromptTemplatesResponseBody) SetPromptTemplates(v []*ListPromptTemplatesResponseBodyPromptTemplates) *ListPromptTemplatesResponseBody {
	s.PromptTemplates = v
	return s
}

func (s *ListPromptTemplatesResponseBody) SetRequestId(v string) *ListPromptTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromptTemplatesResponseBody) SetTotalCount(v int32) *ListPromptTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPromptTemplatesResponseBody) SetWorkspaceId(v string) *ListPromptTemplatesResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *ListPromptTemplatesResponseBody) Validate() error {
	if s.PromptTemplates != nil {
		for _, item := range s.PromptTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPromptTemplatesResponseBodyPromptTemplates struct {
	// The template content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The template name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// d6935b7efbe34d11b13df9307151cf8c
	PromptTemplateId *string `json:"promptTemplateId,omitempty" xml:"promptTemplateId,omitempty"`
	// The template type.
	//
	// example:
	//
	// "System"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The variables of the template.
	Variables []*string `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListPromptTemplatesResponseBodyPromptTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListPromptTemplatesResponseBodyPromptTemplates) GoString() string {
	return s.String()
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) GetContent() *string {
	return s.Content
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) GetName() *string {
	return s.Name
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) GetPromptTemplateId() *string {
	return s.PromptTemplateId
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) GetType() *string {
	return s.Type
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) GetVariables() []*string {
	return s.Variables
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) SetContent(v string) *ListPromptTemplatesResponseBodyPromptTemplates {
	s.Content = &v
	return s
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) SetName(v string) *ListPromptTemplatesResponseBodyPromptTemplates {
	s.Name = &v
	return s
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) SetPromptTemplateId(v string) *ListPromptTemplatesResponseBodyPromptTemplates {
	s.PromptTemplateId = &v
	return s
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) SetType(v string) *ListPromptTemplatesResponseBodyPromptTemplates {
	s.Type = &v
	return s
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) SetVariables(v []*string) *ListPromptTemplatesResponseBodyPromptTemplates {
	s.Variables = v
	return s
}

func (s *ListPromptTemplatesResponseBodyPromptTemplates) Validate() error {
	return dara.Validate(s)
}
