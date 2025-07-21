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
	// The maximum number of entries to return. Valid values: **1 to 100**
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
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
