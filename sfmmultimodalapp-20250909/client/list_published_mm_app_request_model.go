// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListPublishedMmAppRequest
	GetAppId() *string
	SetPageNumber(v int32) *ListPublishedMmAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedMmAppRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListPublishedMmAppRequest
	GetWorkspaceId() *string
}

type ListPublishedMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPublishedMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedMmAppRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedMmAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedMmAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedMmAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPublishedMmAppRequest) SetAppId(v string) *ListPublishedMmAppRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishedMmAppRequest) SetPageNumber(v int32) *ListPublishedMmAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedMmAppRequest) SetPageSize(v int32) *ListPublishedMmAppRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedMmAppRequest) SetWorkspaceId(v string) *ListPublishedMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPublishedMmAppRequest) Validate() error {
	return dara.Validate(s)
}
