// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneralConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListGeneralConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGeneralConfigsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListGeneralConfigsRequest
	GetWorkspaceId() *string
}

type ListGeneralConfigsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGeneralConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGeneralConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListGeneralConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGeneralConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGeneralConfigsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGeneralConfigsRequest) SetPageNumber(v int32) *ListGeneralConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGeneralConfigsRequest) SetPageSize(v int32) *ListGeneralConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGeneralConfigsRequest) SetWorkspaceId(v string) *ListGeneralConfigsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGeneralConfigsRequest) Validate() error {
	return dara.Validate(s)
}
