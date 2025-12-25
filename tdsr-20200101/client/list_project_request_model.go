// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListProjectRequest
	GetName() *string
	SetPageNum(v int64) *ListProjectRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListProjectRequest
	GetPageSize() *int64
}

type ListProjectRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) GetName() *string {
	return s.Name
}

func (s *ListProjectRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListProjectRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectRequest) SetName(v string) *ListProjectRequest {
	s.Name = &v
	return s
}

func (s *ListProjectRequest) SetPageNum(v int64) *ListProjectRequest {
	s.PageNum = &v
	return s
}

func (s *ListProjectRequest) SetPageSize(v int64) *ListProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRequest) Validate() error {
	return dara.Validate(s)
}
