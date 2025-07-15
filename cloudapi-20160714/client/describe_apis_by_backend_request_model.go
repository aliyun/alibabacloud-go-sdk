// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByBackendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *DescribeApisByBackendRequest
	GetBackendId() *string
	SetPageNumber(v int32) *DescribeApisByBackendRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByBackendRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisByBackendRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApisByBackendRequest
	GetStageName() *string
}

type DescribeApisByBackendRequest struct {
	// The ID of the backend service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4ac69b7a17524781b275ed4c5eb25c54
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment to which the API is published. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE**
	//
	// 	- **TEST**
	//
	// If you do not specify this parameter, APIs in the draft state are returned.
	//
	// example:
	//
	// PRE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApisByBackendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByBackendRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByBackendRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeApisByBackendRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByBackendRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByBackendRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisByBackendRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisByBackendRequest) SetBackendId(v string) *DescribeApisByBackendRequest {
	s.BackendId = &v
	return s
}

func (s *DescribeApisByBackendRequest) SetPageNumber(v int32) *DescribeApisByBackendRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByBackendRequest) SetPageSize(v int32) *DescribeApisByBackendRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByBackendRequest) SetSecurityToken(v string) *DescribeApisByBackendRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByBackendRequest) SetStageName(v string) *DescribeApisByBackendRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApisByBackendRequest) Validate() error {
	return dara.Validate(s)
}
