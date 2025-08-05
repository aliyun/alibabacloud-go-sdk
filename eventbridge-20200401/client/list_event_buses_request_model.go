// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListEventBusesRequest
	GetLimit() *int32
	SetNamePrefix(v string) *ListEventBusesRequest
	GetNamePrefix() *string
	SetNextToken(v string) *ListEventBusesRequest
	GetNextToken() *string
}

type ListEventBusesRequest struct {
	// The maximum number of entries to return in a request. You can use this parameter and NextToken to implement paging.
	//
	// >  A maximum of 100 entries can be returned in a request.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The prefix of the names of the event buses that you want to query.
	//
	// example:
	//
	// My
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned. You can use this parameter and Limit to implement paging.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListEventBusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventBusesRequest) GoString() string {
	return s.String()
}

func (s *ListEventBusesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListEventBusesRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *ListEventBusesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventBusesRequest) SetLimit(v int32) *ListEventBusesRequest {
	s.Limit = &v
	return s
}

func (s *ListEventBusesRequest) SetNamePrefix(v string) *ListEventBusesRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListEventBusesRequest) SetNextToken(v string) *ListEventBusesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventBusesRequest) Validate() error {
	return dara.Validate(s)
}
