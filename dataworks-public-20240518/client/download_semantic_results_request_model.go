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
	// The node name. Use the Data.Name value from the CreateSemanticJob response, the Name value from a ListSemanticJobs list item, or the JobName value from a ListSemanticJobRuns record.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The optional run ID. If you specify the Data.JobRunId value from the RunSemanticJob response or the JobRunId value from a ListSemanticJobRuns record, only the artifacts of that specific run are returned. If you do not specify this parameter, the artifacts of the latest run of the node are returned.
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
