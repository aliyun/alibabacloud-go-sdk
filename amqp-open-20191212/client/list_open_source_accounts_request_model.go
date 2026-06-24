// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourceAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOpenSourceAccountsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListOpenSourceAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpenSourceAccountsRequest
	GetNextToken() *string
}

type ListOpenSourceAccountsRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-7pp2mwbc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current query. This token is passed as a parameter in the next call to continue pagination. Set this parameter to an empty string for the first call or when the last page is returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListOpenSourceAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourceAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListOpenSourceAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOpenSourceAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpenSourceAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpenSourceAccountsRequest) SetInstanceId(v string) *ListOpenSourceAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOpenSourceAccountsRequest) SetMaxResults(v int32) *ListOpenSourceAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpenSourceAccountsRequest) SetNextToken(v string) *ListOpenSourceAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpenSourceAccountsRequest) Validate() error {
	return dara.Validate(s)
}
