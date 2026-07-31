// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSemanticResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DownloadSemanticResultsResponseBodyData) *DownloadSemanticResultsResponseBody
	GetData() *DownloadSemanticResultsResponseBodyData
	SetRequestId(v string) *DownloadSemanticResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadSemanticResultsResponseBody
	GetSuccess() *bool
}

type DownloadSemanticResultsResponseBody struct {
	// The collection of result files for the specified node run. Multiple items are returned if a single run generates multiple files.
	Data *DownloadSemanticResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadSemanticResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSemanticResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSemanticResultsResponseBody) GetData() *DownloadSemanticResultsResponseBodyData {
	return s.Data
}

func (s *DownloadSemanticResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSemanticResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadSemanticResultsResponseBody) SetData(v *DownloadSemanticResultsResponseBodyData) *DownloadSemanticResultsResponseBody {
	s.Data = v
	return s
}

func (s *DownloadSemanticResultsResponseBody) SetRequestId(v string) *DownloadSemanticResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSemanticResultsResponseBody) SetSuccess(v bool) *DownloadSemanticResultsResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadSemanticResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadSemanticResultsResponseBodyData struct {
	// The list of result files. Each item contains the associated node name, the associated run ID, and a short-lived download URL.
	Results []*DownloadSemanticResultsResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DownloadSemanticResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DownloadSemanticResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DownloadSemanticResultsResponseBodyData) GetResults() []*DownloadSemanticResultsResponseBodyDataResults {
	return s.Results
}

func (s *DownloadSemanticResultsResponseBodyData) SetResults(v []*DownloadSemanticResultsResponseBodyDataResults) *DownloadSemanticResultsResponseBodyData {
	s.Results = v
	return s
}

func (s *DownloadSemanticResultsResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DownloadSemanticResultsResponseBodyDataResults struct {
	// The temporary pre-signed download URL of the result file. Download the file by using an HTTP GET request as soon as possible. Do not log, share, or treat the full URL as a long-term address.
	//
	// example:
	//
	// https://example.com/temporary-download-url
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The node name to which the artifact belongs. This value is the same as the JobName value in the request.
	//
	// example:
	//
	// semantic-job-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The run ID to which the artifact belongs. You can compare this value with the Data.JobRunId value from the RunSemanticJob response or the JobRunId value from ListSemanticJobRuns.
	//
	// example:
	//
	// 01H00000000000000000000000
	JobRunId *string `json:"JobRunId,omitempty" xml:"JobRunId,omitempty"`
}

func (s DownloadSemanticResultsResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s DownloadSemanticResultsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *DownloadSemanticResultsResponseBodyDataResults) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadSemanticResultsResponseBodyDataResults) GetJobName() *string {
	return s.JobName
}

func (s *DownloadSemanticResultsResponseBodyDataResults) GetJobRunId() *string {
	return s.JobRunId
}

func (s *DownloadSemanticResultsResponseBodyDataResults) SetDownloadUrl(v string) *DownloadSemanticResultsResponseBodyDataResults {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadSemanticResultsResponseBodyDataResults) SetJobName(v string) *DownloadSemanticResultsResponseBodyDataResults {
	s.JobName = &v
	return s
}

func (s *DownloadSemanticResultsResponseBodyDataResults) SetJobRunId(v string) *DownloadSemanticResultsResponseBodyDataResults {
	s.JobRunId = &v
	return s
}

func (s *DownloadSemanticResultsResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
