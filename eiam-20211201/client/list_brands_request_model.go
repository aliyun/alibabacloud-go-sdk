// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBrandsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListBrandsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListBrandsRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListBrandsRequest
	GetPreviousToken() *string
}

type ListBrandsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询上一页凭证（Token），取值为上一次API调用返回的previousToken参数值。
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListBrandsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBrandsRequest) GoString() string {
	return s.String()
}

func (s *ListBrandsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBrandsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListBrandsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBrandsRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListBrandsRequest) SetInstanceId(v string) *ListBrandsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBrandsRequest) SetMaxResults(v int64) *ListBrandsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBrandsRequest) SetNextToken(v string) *ListBrandsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBrandsRequest) SetPreviousToken(v string) *ListBrandsRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListBrandsRequest) Validate() error {
	return dara.Validate(s)
}
