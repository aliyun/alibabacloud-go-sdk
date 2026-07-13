// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListModelProvidersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListModelProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListModelProvidersRequest
	GetNextToken() *string
}

type ListModelProvidersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListModelProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListModelProvidersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListModelProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelProvidersRequest) SetInstanceId(v string) *ListModelProvidersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListModelProvidersRequest) SetMaxResults(v int32) *ListModelProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelProvidersRequest) SetNextToken(v string) *ListModelProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelProvidersRequest) Validate() error {
	return dara.Validate(s)
}
