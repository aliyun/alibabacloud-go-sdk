// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSemanticResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobName(v string) *DownloadSemanticResultsRequest
	GetJobName() *string
	SetJobRunId(v string) *DownloadSemanticResultsRequest
	GetJobRunId() *string
}

type DownloadSemanticResultsRequest struct {
	// The job name. You can obtain this value from Data.Name in the CreateSemanticJob response, Name in the ListSemanticJobs response, or JobName in the ListSemanticJobRuns response.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The optional run ID. If you specify the JobRunId from the RunSemanticJob response (Data.JobRunId) or the ListSemanticJobRuns response, only the artifacts of the specified run are returned. If you do not specify this parameter, the artifacts of the most recent run of the job are returned.
	//
	// example:
	//
	// 01H00000000000000000000000
	JobRunId *string `json:"JobRunId,omitempty" xml:"JobRunId,omitempty"`
}

func (s DownloadSemanticResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSemanticResultsRequest) GoString() string {
	return s.String()
}

func (s *DownloadSemanticResultsRequest) GetJobName() *string {
	return s.JobName
}

func (s *DownloadSemanticResultsRequest) GetJobRunId() *string {
	return s.JobRunId
}

func (s *DownloadSemanticResultsRequest) SetJobName(v string) *DownloadSemanticResultsRequest {
	s.JobName = &v
	return s
}

func (s *DownloadSemanticResultsRequest) SetJobRunId(v string) *DownloadSemanticResultsRequest {
	s.JobRunId = &v
	return s
}

func (s *DownloadSemanticResultsRequest) Validate() error {
	return dara.Validate(s)
}
