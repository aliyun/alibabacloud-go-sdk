// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListCommandRequest
	GetAppId() *string
	SetDomainCode(v string) *ListCommandRequest
	GetDomainCode() *string
	SetPageNumber(v int32) *ListCommandRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommandRequest
	GetPageSize() *int32
	SetToolName(v string) *ListCommandRequest
	GetToolName() *string
	SetWorkspaceId(v string) *ListCommandRequest
	GetWorkspaceId() *string
}

type ListCommandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8453564564
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xl-sxx
	ToolName *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommandRequest) GoString() string {
	return s.String()
}

func (s *ListCommandRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCommandRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *ListCommandRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommandRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommandRequest) GetToolName() *string {
	return s.ToolName
}

func (s *ListCommandRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCommandRequest) SetAppId(v string) *ListCommandRequest {
	s.AppId = &v
	return s
}

func (s *ListCommandRequest) SetDomainCode(v string) *ListCommandRequest {
	s.DomainCode = &v
	return s
}

func (s *ListCommandRequest) SetPageNumber(v int32) *ListCommandRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommandRequest) SetPageSize(v int32) *ListCommandRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommandRequest) SetToolName(v string) *ListCommandRequest {
	s.ToolName = &v
	return s
}

func (s *ListCommandRequest) SetWorkspaceId(v string) *ListCommandRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCommandRequest) Validate() error {
	return dara.Validate(s)
}
