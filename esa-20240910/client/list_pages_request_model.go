// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPagesRequest
	GetPageSize() *int32
	SetQueryArgs(v *ListPagesRequestQueryArgs) *ListPagesRequest
	GetQueryArgs() *ListPagesRequestQueryArgs
}

type ListPagesRequest struct {
	// The page number. Valid values: **1 to 100000**. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize  *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryArgs *ListPagesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
}

func (s ListPagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPagesRequest) GoString() string {
	return s.String()
}

func (s *ListPagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPagesRequest) GetQueryArgs() *ListPagesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListPagesRequest) SetPageNumber(v int32) *ListPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPagesRequest) SetPageSize(v int32) *ListPagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPagesRequest) SetQueryArgs(v *ListPagesRequestQueryArgs) *ListPagesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListPagesRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPagesRequestQueryArgs struct {
	NameDescriptionLike *string `json:"NameDescriptionLike,omitempty" xml:"NameDescriptionLike,omitempty"`
}

func (s ListPagesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListPagesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListPagesRequestQueryArgs) GetNameDescriptionLike() *string {
	return s.NameDescriptionLike
}

func (s *ListPagesRequestQueryArgs) SetNameDescriptionLike(v string) *ListPagesRequestQueryArgs {
	s.NameDescriptionLike = &v
	return s
}

func (s *ListPagesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
