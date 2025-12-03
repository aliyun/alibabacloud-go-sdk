// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendName(v string) *DescribeBackendListRequest
	GetBackendName() *string
	SetBackendType(v string) *DescribeBackendListRequest
	GetBackendType() *string
	SetPageNumber(v int32) *DescribeBackendListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackendListRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeBackendListRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeBackendListRequestTag) *DescribeBackendListRequest
	GetTag() []*DescribeBackendListRequestTag
}

type DescribeBackendListRequest struct {
	// The name of the backend service. You can use \\	- to perform fuzzy queries.
	//
	// example:
	//
	// test
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The type of the backend service.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
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
	// The list of the tag.
	Tag []*DescribeBackendListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeBackendListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackendListRequest) GetBackendName() *string {
	return s.BackendName
}

func (s *DescribeBackendListRequest) GetBackendType() *string {
	return s.BackendType
}

func (s *DescribeBackendListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackendListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackendListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeBackendListRequest) GetTag() []*DescribeBackendListRequestTag {
	return s.Tag
}

func (s *DescribeBackendListRequest) SetBackendName(v string) *DescribeBackendListRequest {
	s.BackendName = &v
	return s
}

func (s *DescribeBackendListRequest) SetBackendType(v string) *DescribeBackendListRequest {
	s.BackendType = &v
	return s
}

func (s *DescribeBackendListRequest) SetPageNumber(v int32) *DescribeBackendListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackendListRequest) SetPageSize(v int32) *DescribeBackendListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackendListRequest) SetSecurityToken(v string) *DescribeBackendListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackendListRequest) SetTag(v []*DescribeBackendListRequestTag) *DescribeBackendListRequest {
	s.Tag = v
	return s
}

func (s *DescribeBackendListRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackendListRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// test1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackendListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeBackendListRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeBackendListRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeBackendListRequestTag) SetKey(v string) *DescribeBackendListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeBackendListRequestTag) SetValue(v string) *DescribeBackendListRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeBackendListRequestTag) Validate() error {
	return dara.Validate(s)
}
