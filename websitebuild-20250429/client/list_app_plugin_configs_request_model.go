// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppPluginConfigsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListAppPluginConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppPluginConfigsRequest
	GetNextToken() *string
}

type ListAppPluginConfigsRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAppPluginConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAppPluginConfigsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppPluginConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPluginConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPluginConfigsRequest) SetBizId(v string) *ListAppPluginConfigsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppPluginConfigsRequest) SetMaxResults(v int32) *ListAppPluginConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppPluginConfigsRequest) SetNextToken(v string) *ListAppPluginConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppPluginConfigsRequest) Validate() error {
	return dara.Validate(s)
}
