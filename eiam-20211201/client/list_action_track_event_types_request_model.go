// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionTrackEventTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListActionTrackEventTypesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListActionTrackEventTypesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListActionTrackEventTypesRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListActionTrackEventTypesRequest
	GetPreviousToken() *string
}

type ListActionTrackEventTypesRequest struct {
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

func (s ListActionTrackEventTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionTrackEventTypesRequest) GoString() string {
	return s.String()
}

func (s *ListActionTrackEventTypesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListActionTrackEventTypesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListActionTrackEventTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionTrackEventTypesRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListActionTrackEventTypesRequest) SetInstanceId(v string) *ListActionTrackEventTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListActionTrackEventTypesRequest) SetMaxResults(v int64) *ListActionTrackEventTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionTrackEventTypesRequest) SetNextToken(v string) *ListActionTrackEventTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionTrackEventTypesRequest) SetPreviousToken(v string) *ListActionTrackEventTypesRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListActionTrackEventTypesRequest) Validate() error {
	return dara.Validate(s)
}
