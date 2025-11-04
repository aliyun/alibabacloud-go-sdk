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
	// The status of the task. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Fail**
	//
	// 	- **Processing**
	//
	// 	- **Submitted**
	//
	// example:
	//
	// Success
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of analysis result objects.
	Results *QueryVideoCognitionJobResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The user-defined data.
	//
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
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVideoCognitionJobResponseBodyResultsResult struct {
	// A JSON string containing the detailed analysis data. The structure of this data depends on the Type field. For details, see the Result parameters section below.
	//
	// example:
	//
	// {"title":"example-title-****"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of analysis result. Valid values:
	//
	// 1.  TextLabel: Tags from text content.
	//
	// 2.  VideoLabel: Tags from video content.
	//
	// 3.  ASR: Raw speech recognition results. Not returned by default.
	//
	// 4.  OCR: Raw text recognition results. Not returned by default.
	//
	// 5.  NLP: Natural Language Processing results. Not returned by default.
	//
	// 6.  Process: URL to the raw algorithm output. Not returned by default.
	//
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
