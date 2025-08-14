// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoCognitionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobStatus(v string) *QueryVideoCognitionJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *QueryVideoCognitionJobResponseBody
	GetRequestId() *string
	SetResults(v *QueryVideoCognitionJobResponseBodyResults) *QueryVideoCognitionJobResponseBody
	GetResults() *QueryVideoCognitionJobResponseBodyResults
	SetUserData(v string) *QueryVideoCognitionJobResponseBody
	GetUserData() *string
}

type QueryVideoCognitionJobResponseBody struct {
	// example:
	//
	// Success
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *QueryVideoCognitionJobResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// example:
	//
	// {"userId":"123432412831"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryVideoCognitionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *QueryVideoCognitionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVideoCognitionJobResponseBody) GetResults() *QueryVideoCognitionJobResponseBodyResults {
	return s.Results
}

func (s *QueryVideoCognitionJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *QueryVideoCognitionJobResponseBody) SetJobStatus(v string) *QueryVideoCognitionJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetRequestId(v string) *QueryVideoCognitionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetResults(v *QueryVideoCognitionJobResponseBodyResults) *QueryVideoCognitionJobResponseBody {
	s.Results = v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetUserData(v string) *QueryVideoCognitionJobResponseBody {
	s.UserData = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryVideoCognitionJobResponseBodyResults struct {
	Result []*QueryVideoCognitionJobResponseBodyResultsResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryVideoCognitionJobResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobResponseBodyResults) GetResult() []*QueryVideoCognitionJobResponseBodyResultsResult {
	return s.Result
}

func (s *QueryVideoCognitionJobResponseBodyResults) SetResult(v []*QueryVideoCognitionJobResponseBodyResultsResult) *QueryVideoCognitionJobResponseBodyResults {
	s.Result = v
	return s
}

func (s *QueryVideoCognitionJobResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type QueryVideoCognitionJobResponseBodyResultsResult struct {
	// example:
	//
	// {"title":"example-title-****"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ASR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryVideoCognitionJobResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobResponseBodyResultsResult) GetData() *string {
	return s.Data
}

func (s *QueryVideoCognitionJobResponseBodyResultsResult) GetType() *string {
	return s.Type
}

func (s *QueryVideoCognitionJobResponseBodyResultsResult) SetData(v string) *QueryVideoCognitionJobResponseBodyResultsResult {
	s.Data = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBodyResultsResult) SetType(v string) *QueryVideoCognitionJobResponseBodyResultsResult {
	s.Type = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBodyResultsResult) Validate() error {
	return dara.Validate(s)
}
