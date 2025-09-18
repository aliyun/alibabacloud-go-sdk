// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListMmAppRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListMmAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMmAppRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListMmAppRequest
	GetStatus() *int32
	SetWorkspaceId(v string) *ListMmAppRequest
	GetWorkspaceId() *string
}

type ListMmAppRequest struct {
	// example:
	//
	// 多模态
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppRequest) GoString() string {
	return s.String()
}

func (s *ListMmAppRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMmAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMmAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmAppRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMmAppRequest) SetKeyword(v string) *ListMmAppRequest {
	s.Keyword = &v
	return s
}

func (s *ListMmAppRequest) SetPageNumber(v int32) *ListMmAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMmAppRequest) SetPageSize(v int32) *ListMmAppRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmAppRequest) SetStatus(v int32) *ListMmAppRequest {
	s.Status = &v
	return s
}

func (s *ListMmAppRequest) SetWorkspaceId(v string) *ListMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMmAppRequest) Validate() error {
	return dara.Validate(s)
}
