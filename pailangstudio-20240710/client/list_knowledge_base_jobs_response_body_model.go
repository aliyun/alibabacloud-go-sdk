// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseJobs(v []*KnowledgeBaseJob) *ListKnowledgeBaseJobsResponseBody
	GetKnowledgeBaseJobs() []*KnowledgeBaseJob
	SetMaxResults(v int32) *ListKnowledgeBaseJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListKnowledgeBaseJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKnowledgeBaseJobsResponseBody
	GetTotalCount() *int32
}

type ListKnowledgeBaseJobsResponseBody struct {
	KnowledgeBaseJobs []*KnowledgeBaseJob `json:"KnowledgeBaseJobs,omitempty" xml:"KnowledgeBaseJobs,omitempty" type:"Repeated"`
	MaxResults        *int32              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 963BD7F9-0C02-5594-9550-BCC6DD43E3C0
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKnowledgeBaseJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseJobsResponseBody) GetKnowledgeBaseJobs() []*KnowledgeBaseJob {
	return s.KnowledgeBaseJobs
}

func (s *ListKnowledgeBaseJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKnowledgeBaseJobsResponseBody) SetKnowledgeBaseJobs(v []*KnowledgeBaseJob) *ListKnowledgeBaseJobsResponseBody {
	s.KnowledgeBaseJobs = v
	return s
}

func (s *ListKnowledgeBaseJobsResponseBody) SetMaxResults(v int32) *ListKnowledgeBaseJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseJobsResponseBody) SetNextToken(v string) *ListKnowledgeBaseJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseJobsResponseBody) SetRequestId(v string) *ListKnowledgeBaseJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseJobsResponseBody) SetTotalCount(v int32) *ListKnowledgeBaseJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBaseJobsResponseBody) Validate() error {
	if s.KnowledgeBaseJobs != nil {
		for _, item := range s.KnowledgeBaseJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
