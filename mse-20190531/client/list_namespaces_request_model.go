// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNamespacesRequest
	GetAcceptLanguage() *string
	SetName(v string) *ListNamespacesRequest
	GetName() *string
	SetPageNumber(v int32) *ListNamespacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNamespacesRequest
	GetPageSize() *int32
	SetRegion(v string) *ListNamespacesRequest
	GetRegion() *string
	SetTag(v []*ListNamespacesRequestTag) *ListNamespacesRequest
	GetTag() []*ListNamespacesRequestTag
}

type ListNamespacesRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// myNamespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string                     `json:"Region,omitempty" xml:"Region,omitempty"`
	Tag    []*ListNamespacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNamespacesRequest) GetName() *string {
	return s.Name
}

func (s *ListNamespacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNamespacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListNamespacesRequest) GetTag() []*ListNamespacesRequestTag {
	return s.Tag
}

func (s *ListNamespacesRequest) SetAcceptLanguage(v string) *ListNamespacesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNamespacesRequest) SetName(v string) *ListNamespacesRequest {
	s.Name = &v
	return s
}

func (s *ListNamespacesRequest) SetPageNumber(v int32) *ListNamespacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNamespacesRequest) SetPageSize(v int32) *ListNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesRequest) SetRegion(v string) *ListNamespacesRequest {
	s.Region = &v
	return s
}

func (s *ListNamespacesRequest) SetTag(v []*ListNamespacesRequestTag) *ListNamespacesRequest {
	s.Tag = v
	return s
}

func (s *ListNamespacesRequest) Validate() error {
	return dara.Validate(s)
}

type ListNamespacesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNamespacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListNamespacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListNamespacesRequestTag) SetKey(v string) *ListNamespacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListNamespacesRequestTag) SetValue(v string) *ListNamespacesRequestTag {
	s.Value = &v
	return s
}

func (s *ListNamespacesRequestTag) Validate() error {
	return dara.Validate(s)
}
