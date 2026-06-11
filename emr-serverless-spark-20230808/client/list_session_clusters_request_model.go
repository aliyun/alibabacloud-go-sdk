// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKind(v string) *ListSessionClustersRequest
	GetKind() *string
	SetMaxResults(v int32) *ListSessionClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSessionClustersRequest
	GetNextToken() *string
	SetQueueName(v string) *ListSessionClustersRequest
	GetQueueName() *string
	SetRegionId(v string) *ListSessionClustersRequest
	GetRegionId() *string
	SetSessionClusterId(v string) *ListSessionClustersRequest
	GetSessionClusterId() *string
}

type ListSessionClustersRequest struct {
	// The session type.
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The maximum number of records to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that marks the start of the next page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// emr-spark-demo-job
	SessionClusterId *string `json:"sessionClusterId,omitempty" xml:"sessionClusterId,omitempty"`
}

func (s ListSessionClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersRequest) GoString() string {
	return s.String()
}

func (s *ListSessionClustersRequest) GetKind() *string {
	return s.Kind
}

func (s *ListSessionClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionClustersRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *ListSessionClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSessionClustersRequest) GetSessionClusterId() *string {
	return s.SessionClusterId
}

func (s *ListSessionClustersRequest) SetKind(v string) *ListSessionClustersRequest {
	s.Kind = &v
	return s
}

func (s *ListSessionClustersRequest) SetMaxResults(v int32) *ListSessionClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSessionClustersRequest) SetNextToken(v string) *ListSessionClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionClustersRequest) SetQueueName(v string) *ListSessionClustersRequest {
	s.QueueName = &v
	return s
}

func (s *ListSessionClustersRequest) SetRegionId(v string) *ListSessionClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListSessionClustersRequest) SetSessionClusterId(v string) *ListSessionClustersRequest {
	s.SessionClusterId = &v
	return s
}

func (s *ListSessionClustersRequest) Validate() error {
	return dara.Validate(s)
}
