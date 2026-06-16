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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// You can use this parameter to specify the query token.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token that is used to retrieve the previous page of results.
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
