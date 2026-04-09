// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *ListDataPermissionsRequest
	GetDataId() *string
	SetDataType(v string) *ListDataPermissionsRequest
	GetDataType() *string
	SetPageNumber(v int32) *ListDataPermissionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataPermissionsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListDataPermissionsRequest
	GetWorkspaceId() *string
}

type ListDataPermissionsRequest struct {
	// example:
	//
	// SystemSearch.QuarkCommonNews
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// dataset
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataPermissionsRequest) GetDataId() *string {
	return s.DataId
}

func (s *ListDataPermissionsRequest) GetDataType() *string {
	return s.DataType
}

func (s *ListDataPermissionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataPermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataPermissionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataPermissionsRequest) SetDataId(v string) *ListDataPermissionsRequest {
	s.DataId = &v
	return s
}

func (s *ListDataPermissionsRequest) SetDataType(v string) *ListDataPermissionsRequest {
	s.DataType = &v
	return s
}

func (s *ListDataPermissionsRequest) SetPageNumber(v int32) *ListDataPermissionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataPermissionsRequest) SetPageSize(v int32) *ListDataPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataPermissionsRequest) SetWorkspaceId(v string) *ListDataPermissionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
