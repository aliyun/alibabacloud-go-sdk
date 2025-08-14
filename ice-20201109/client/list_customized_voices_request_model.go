// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListCustomizedVoicesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListCustomizedVoicesRequest
	GetPageSize() *int32
	SetType(v string) *ListCustomizedVoicesRequest
	GetType() *string
}

type ListCustomizedVoicesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 	- The voice type. Valid values:
	//
	//     	- Basic
	//
	//     	- Standard
	//
	// 	- If you do not specify this parameter, the default value Basic is used.
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomizedVoicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoicesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoicesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListCustomizedVoicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomizedVoicesRequest) GetType() *string {
	return s.Type
}

func (s *ListCustomizedVoicesRequest) SetPageNo(v int32) *ListCustomizedVoicesRequest {
	s.PageNo = &v
	return s
}

func (s *ListCustomizedVoicesRequest) SetPageSize(v int32) *ListCustomizedVoicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomizedVoicesRequest) SetType(v string) *ListCustomizedVoicesRequest {
	s.Type = &v
	return s
}

func (s *ListCustomizedVoicesRequest) Validate() error {
	return dara.Validate(s)
}
