// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosMcpServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListNacosMcpServersRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListNacosMcpServersRequest
	GetInstanceId() *string
	SetName(v string) *ListNacosMcpServersRequest
	GetName() *string
	SetNamespaceId(v string) *ListNacosMcpServersRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListNacosMcpServersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListNacosMcpServersRequest
	GetPageSize() *int32
	SetSearch(v string) *ListNacosMcpServersRequest
	GetSearch() *string
}

type ListNacosMcpServersRequest struct {
	// The language of the response. Valid values: zh (default): Chinese, en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The MCP server name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 5e3ee449-b5c0-4aee-b857-32c0acbebf26
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query pattern.
	//
	// example:
	//
	// blur
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListNacosMcpServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersRequest) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListNacosMcpServersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNacosMcpServersRequest) GetName() *string {
	return s.Name
}

func (s *ListNacosMcpServersRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNacosMcpServersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListNacosMcpServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNacosMcpServersRequest) GetSearch() *string {
	return s.Search
}

func (s *ListNacosMcpServersRequest) SetAcceptLanguage(v string) *ListNacosMcpServersRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetInstanceId(v string) *ListNacosMcpServersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetName(v string) *ListNacosMcpServersRequest {
	s.Name = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetNamespaceId(v string) *ListNacosMcpServersRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetPageNum(v int32) *ListNacosMcpServersRequest {
	s.PageNum = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetPageSize(v int32) *ListNacosMcpServersRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacosMcpServersRequest) SetSearch(v string) *ListNacosMcpServersRequest {
	s.Search = &v
	return s
}

func (s *ListNacosMcpServersRequest) Validate() error {
	return dara.Validate(s)
}
