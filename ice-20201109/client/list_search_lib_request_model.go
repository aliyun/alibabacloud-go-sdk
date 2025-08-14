// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSearchLibRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSearchLibRequest
	GetPageSize() *int32
}

type ListSearchLibRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLibRequest) GoString() string {
	return s.String()
}

func (s *ListSearchLibRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSearchLibRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchLibRequest) SetPageNo(v int32) *ListSearchLibRequest {
	s.PageNo = &v
	return s
}

func (s *ListSearchLibRequest) SetPageSize(v int32) *ListSearchLibRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchLibRequest) Validate() error {
	return dara.Validate(s)
}
