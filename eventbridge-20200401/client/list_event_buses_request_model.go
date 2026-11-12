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
	SetResourceGroupId(v string) *ListEventBusesRequest
	GetResourceGroupId() *string
}

type ListEventBusesRequest struct {
	// The maximum number of entries to return per page. You can use this parameter together with NextToken to implement pagination.
	//
	// 	Notice: The maximum number of entries returned per request is 100.</notice>
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name prefix of the event buses to retrieve.
	//
	// example:
	//
	// My
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// The pagination token. If the specified Limit is reached and more results are available, this token is returned. You can use this token together with Limit to implement pagination.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group, used to filter event buses in the specified resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ListEventBusesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *ListEventBusesRequest) SetResourceGroupId(v string) *ListEventBusesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEventBusesRequest) Validate() error {
	return dara.Validate(s)
}
