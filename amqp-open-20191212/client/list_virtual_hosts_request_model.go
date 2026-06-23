// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVirtualHostsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListVirtualHostsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVirtualHostsRequest
	GetNextToken() *string
}

type ListVirtualHostsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return. Valid values:
	//
	// **1 to 100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If the number of results exceeds the value of MaxResults, NextToken is returned. You can include this parameter in the next call to retrieve the next page of results. Leave this parameter empty for the first call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListVirtualHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualHostsRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVirtualHostsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVirtualHostsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualHostsRequest) SetInstanceId(v string) *ListVirtualHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVirtualHostsRequest) SetMaxResults(v int32) *ListVirtualHostsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsRequest) SetNextToken(v string) *ListVirtualHostsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVirtualHostsRequest) Validate() error {
	return dara.Validate(s)
}
