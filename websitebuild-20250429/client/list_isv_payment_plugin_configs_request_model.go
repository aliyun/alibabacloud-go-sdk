// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsvPaymentPluginConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListIsvPaymentPluginConfigsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListIsvPaymentPluginConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIsvPaymentPluginConfigsRequest
	GetNextToken() *string
}

type ListIsvPaymentPluginConfigsRequest struct {
	// example:
	//
	// WD20250821161210000001
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

func (s ListIsvPaymentPluginConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIsvPaymentPluginConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListIsvPaymentPluginConfigsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListIsvPaymentPluginConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIsvPaymentPluginConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIsvPaymentPluginConfigsRequest) SetBizId(v string) *ListIsvPaymentPluginConfigsRequest {
	s.BizId = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsRequest) SetMaxResults(v int32) *ListIsvPaymentPluginConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsRequest) SetNextToken(v string) *ListIsvPaymentPluginConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsRequest) Validate() error {
	return dara.Validate(s)
}
