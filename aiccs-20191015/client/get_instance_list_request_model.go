// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetInstanceListRequest
	GetName() *string
	SetPageNumber(v int32) *GetInstanceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetInstanceListRequest
	GetPageSize() *int32
}

type GetInstanceListRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequest) GetName() *string {
	return s.Name
}

func (s *GetInstanceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetInstanceListRequest) SetName(v string) *GetInstanceListRequest {
	s.Name = &v
	return s
}

func (s *GetInstanceListRequest) SetPageNumber(v int32) *GetInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceListRequest) SetPageSize(v int32) *GetInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
