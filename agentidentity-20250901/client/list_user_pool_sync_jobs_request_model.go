// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolSyncJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolSyncJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolSyncJobsRequest
	GetNextToken() *string
	SetUserPoolName(v string) *ListUserPoolSyncJobsRequest
	GetUserPoolName() *string
}

type ListUserPoolSyncJobsRequest struct {
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListUserPoolSyncJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolSyncJobsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPoolSyncJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolSyncJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolSyncJobsRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListUserPoolSyncJobsRequest) SetMaxResults(v int32) *ListUserPoolSyncJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolSyncJobsRequest) SetNextToken(v string) *ListUserPoolSyncJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolSyncJobsRequest) SetUserPoolName(v string) *ListUserPoolSyncJobsRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListUserPoolSyncJobsRequest) Validate() error {
	return dara.Validate(s)
}
