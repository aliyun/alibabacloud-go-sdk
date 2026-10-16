// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolsRequest
	GetNextToken() *string
}

type ListUserPoolsRequest struct {
	// The maximum number of entries to return per page. Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page of data. You do not need to specify this parameter for the first call. If a NextToken is returned in the response of the previous call, specify it in the current call to obtain the next page.
	//
	// example:
	//
	// TGlzdFVzZXJQb29sczo6MTA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolsRequest) SetMaxResults(v int32) *ListUserPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolsRequest) SetNextToken(v string) *ListUserPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolsRequest) Validate() error {
	return dara.Validate(s)
}
