// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListTemplatesRequest
	GetStatus() *string
	SetTemplateName(v string) *ListTemplatesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ListTemplatesRequest
	GetTemplateType() *string
	SetWorkspaceId(v string) *ListTemplatesRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListTemplatesRequest
	GetWorkspaceIds() *string
}

type ListTemplatesRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The status of the template. Use this parameter to filter templates.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The name of the template. Use this parameter to filter templates.
	//
	// example:
	//
	// temp-abc
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The type of the template. Use this parameter to filter templates.
	//
	// example:
	//
	// Browser
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// The ID of the workspace to which the template belongs.
	//
	// example:
	//
	// aaa
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The IDs of the workspaces. You can use this parameter to query templates from multiple workspaces.
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplatesRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetStatus(v string) *ListTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) SetWorkspaceId(v string) *ListTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplatesRequest) SetWorkspaceIds(v string) *ListTemplatesRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
