// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceGetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *AISearchResourceGetListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *AISearchResourceGetListRequest
	GetPageSize() *int32
	SetResourceIds(v []*string) *AISearchResourceGetListRequest
	GetResourceIds() []*string
	SetType(v string) *AISearchResourceGetListRequest
	GetType() *string
}

type AISearchResourceGetListRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 20
	PageSize    *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceGetListRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceGetListRequest) GoString() string {
	return s.String()
}

func (s *AISearchResourceGetListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *AISearchResourceGetListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AISearchResourceGetListRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *AISearchResourceGetListRequest) GetType() *string {
	return s.Type
}

func (s *AISearchResourceGetListRequest) SetCurrentPage(v int32) *AISearchResourceGetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *AISearchResourceGetListRequest) SetPageSize(v int32) *AISearchResourceGetListRequest {
	s.PageSize = &v
	return s
}

func (s *AISearchResourceGetListRequest) SetResourceIds(v []*string) *AISearchResourceGetListRequest {
	s.ResourceIds = v
	return s
}

func (s *AISearchResourceGetListRequest) SetType(v string) *AISearchResourceGetListRequest {
	s.Type = &v
	return s
}

func (s *AISearchResourceGetListRequest) Validate() error {
	return dara.Validate(s)
}
