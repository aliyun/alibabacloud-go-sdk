// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoCognitionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *QueryVideoCognitionJobResponseBodyInput) *QueryVideoCognitionJobResponseBody
	GetInput() *QueryVideoCognitionJobResponseBodyInput
	SetJobStatus(v string) *QueryVideoCognitionJobResponseBody
	GetJobStatus() *string
	SetParams(v string) *QueryVideoCognitionJobResponseBody
	GetParams() *string
	SetRequestId(v string) *QueryVideoCognitionJobResponseBody
	GetRequestId() *string
	SetResults(v *QueryVideoCognitionJobResponseBodyResults) *QueryVideoCognitionJobResponseBody
	GetResults() *QueryVideoCognitionJobResponseBodyResults
	SetTemplateId(v string) *QueryVideoCognitionJobResponseBody
	GetTemplateId() *string
	SetUserData(v string) *QueryVideoCognitionJobResponseBody
	GetUserData() *string
}

type QueryVideoCognitionJobResponseBody struct {
	// The input file.
	Input *QueryVideoCognitionJobResponseBodyInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job status. Valid values:
	//
	// - **Success**: The job succeeded.
	//
	// - **Fail**: The job failed.
	//
	// - **Processing**: The job is in progress.
	//
	// - **Submitted**: The job has been submitted and is awaiting processing.
	//
	// example:
	//
	// Success
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The request parameters.
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *QueryVideoCognitionJobResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The template ID.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The user data.
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

func (s *QueryVideoCognitionJobResponseBody) GetInput() *QueryVideoCognitionJobResponseBodyInput {
	return s.Input
}

func (s *QueryVideoCognitionJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *QueryVideoCognitionJobResponseBody) GetParams() *string {
	return s.Params
}

func (s *QueryVideoCognitionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVideoCognitionJobResponseBody) GetResults() *QueryVideoCognitionJobResponseBodyResults {
	return s.Results
}

func (s *QueryVideoCognitionJobResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryVideoCognitionJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *QueryVideoCognitionJobResponseBody) SetInput(v *QueryVideoCognitionJobResponseBodyInput) *QueryVideoCognitionJobResponseBody {
	s.Input = v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetJobStatus(v string) *QueryVideoCognitionJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetParams(v string) *QueryVideoCognitionJobResponseBody {
	s.Params = &v
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

func (s *QueryVideoCognitionJobResponseBody) SetTemplateId(v string) *QueryVideoCognitionJobResponseBody {
	s.TemplateId = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) SetUserData(v string) *QueryVideoCognitionJobResponseBody {
	s.UserData = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBody) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVideoCognitionJobResponseBodyInput struct {
	// The URL of the input file.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid value: OSS.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryVideoCognitionJobResponseBodyInput) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobResponseBodyInput) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobResponseBodyInput) GetMedia() *string {
	return s.Media
}

func (s *QueryVideoCognitionJobResponseBodyInput) GetType() *string {
	return s.Type
}

func (s *QueryVideoCognitionJobResponseBodyInput) SetMedia(v string) *QueryVideoCognitionJobResponseBodyInput {
	s.Media = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBodyInput) SetType(v string) *QueryVideoCognitionJobResponseBodyInput {
	s.Type = &v
	return s
}

func (s *QueryVideoCognitionJobResponseBodyInput) Validate() error {
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
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
