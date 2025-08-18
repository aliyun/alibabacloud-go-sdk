// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKvsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListKvsRequest
	GetNamespace() *string
	SetPageNumber(v int32) *ListKvsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKvsRequest
	GetPageSize() *int32
	SetPrefix(v string) *ListKvsRequest
	GetPrefix() *string
}

type ListKvsRequest struct {
	// The name of the namespace that you specify when you call the [CreatevNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. The product of PageNumber and PageSize cannot exceed 50,000.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 50. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The prefix to query.
	//
	// example:
	//
	// prefix-
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListKvsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKvsRequest) GoString() string {
	return s.String()
}

func (s *ListKvsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListKvsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKvsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKvsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListKvsRequest) SetNamespace(v string) *ListKvsRequest {
	s.Namespace = &v
	return s
}

func (s *ListKvsRequest) SetPageNumber(v int32) *ListKvsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKvsRequest) SetPageSize(v int32) *ListKvsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKvsRequest) SetPrefix(v string) *ListKvsRequest {
	s.Prefix = &v
	return s
}

func (s *ListKvsRequest) Validate() error {
	return dara.Validate(s)
}
