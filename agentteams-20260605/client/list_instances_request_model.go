// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListInstancesRequest
	GetInstanceName() *string
	SetMaxResults(v int32) *ListInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesRequest
	GetNextToken() *string
	SetSkip(v int32) *ListInstancesRequest
	GetSkip() *int32
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// example:
	//
	// AgentTeams
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 20
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Skip      *int32  `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetSkip(v int32) *ListInstancesRequest {
	s.Skip = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
