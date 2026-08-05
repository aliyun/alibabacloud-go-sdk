// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListCredentialsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListCredentialsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCredentialsRequest
	GetPageSize() *int32
}

type ListCredentialsRequest struct {
	// The keyword used to search for credential values.
	//
	// example:
	//
	// OS-**
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCredentialsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCredentialsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCredentialsRequest) SetKeyword(v string) *ListCredentialsRequest {
	s.Keyword = &v
	return s
}

func (s *ListCredentialsRequest) SetPageNumber(v int32) *ListCredentialsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCredentialsRequest) SetPageSize(v int32) *ListCredentialsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
