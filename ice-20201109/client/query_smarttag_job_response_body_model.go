// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobStatus(v string) *QuerySmarttagJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *QuerySmarttagJobResponseBody
	GetRequestId() *string
	SetResults(v *QuerySmarttagJobResponseBodyResults) *QuerySmarttagJobResponseBody
	GetResults() *QuerySmarttagJobResponseBodyResults
	SetUserData(v string) *QuerySmarttagJobResponseBody
	GetUserData() *string
}

type QuerySmarttagJobResponseBody struct {
	// The status of the job. Valid values:
	//
	// 	- **Success**: The job was successful.
	//
	// 	- **Fail**: The job failed.
	//
	// 	- **Processing**: The job is in progress.
	//
	// 	- **Submitted**: The job is submitted and waiting to be processed.
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
	// The analysis results of the smart tagging job. The value is an array.
	Results *QuerySmarttagJobResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The content of callback messages that are sent to Simple Message Queue (SMQ) when the information of the smart tagging job changes. For more information about the parameters contained in the callback message, see the "Callback parameters" section of this topic.
	//
	// example:
	//
	// {"userId":"123432412831"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QuerySmarttagJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *QuerySmarttagJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmarttagJobResponseBody) GetResults() *QuerySmarttagJobResponseBodyResults {
	return s.Results
}

func (s *QuerySmarttagJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *QuerySmarttagJobResponseBody) SetJobStatus(v string) *QuerySmarttagJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *QuerySmarttagJobResponseBody) SetRequestId(v string) *QuerySmarttagJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmarttagJobResponseBody) SetResults(v *QuerySmarttagJobResponseBodyResults) *QuerySmarttagJobResponseBody {
	s.Results = v
	return s
}

func (s *QuerySmarttagJobResponseBody) SetUserData(v string) *QuerySmarttagJobResponseBody {
	s.UserData = &v
	return s
}

func (s *QuerySmarttagJobResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySmarttagJobResponseBodyResults struct {
	Result []*QuerySmarttagJobResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QuerySmarttagJobResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobResponseBodyResults) GetResult() []*QuerySmarttagJobResponseBodyResultsResult {
	return s.Result
}

func (s *QuerySmarttagJobResponseBodyResults) SetResult(v []*QuerySmarttagJobResponseBodyResultsResult) *QuerySmarttagJobResponseBodyResults {
	s.Result = v
	return s
}

func (s *QuerySmarttagJobResponseBodyResults) Validate() error {
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

type QuerySmarttagJobResponseBodyResultsResult struct {
	// The details of the analysis result. The value is a JSON string. For more information about the parameters of different result types, see the "Parameters of different result types" section of this topic.
	//
	// example:
	//
	// {"title":"example-title-****"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of the analysis result.
	//
	// 	- The type of the analysis result based on Smart tagging V1.0. Valid values:
	//
	// 1.  TextLabel: the text tag.
	//
	// 2.  VideoLabel: the video tag.
	//
	// 3.  ASR: the original result of automatic speech recognition (ASR). By default, this type of result is not returned.
	//
	// 4.  OCR: the original result of optical character recognition (OCR). By default, this type of result is not returned.
	//
	// 5.  NLP: the natural language processing (NLP)-based result. By default, this type of result is not returned.
	//
	// 	- The type of the analysis result based on Smart tagging V2.0. Valid values:
	//
	// 1.  CPVLabel
	//
	// 2.  Meta: the information about the video file, such as the title of the video. By default, this type of information is not returned.
	//
	// 	- The type of the analysis result based on Smart tagging V2.0-custom. Valid values:
	//
	// 1.  CPVLabel
	//
	// 2.  Meta: the information about the video file, such as the title of the video. By default, this type of information is not returned.
	//
	// example:
	//
	// Meta
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySmarttagJobResponseBodyResultsResult) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobResponseBodyResultsResult) GetData() *string {
	return s.Data
}

func (s *QuerySmarttagJobResponseBodyResultsResult) GetType() *string {
	return s.Type
}

func (s *QuerySmarttagJobResponseBodyResultsResult) SetData(v string) *QuerySmarttagJobResponseBodyResultsResult {
	s.Data = &v
	return s
}

func (s *QuerySmarttagJobResponseBodyResultsResult) SetType(v string) *QuerySmarttagJobResponseBodyResultsResult {
	s.Type = &v
	return s
}

func (s *QuerySmarttagJobResponseBodyResultsResult) Validate() error {
	return dara.Validate(s)
}
