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
	SetMessage(v string) *QuerySmarttagJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySmarttagJobResponseBody
	GetRequestId() *string
	SetResults(v *QuerySmarttagJobResponseBodyResults) *QuerySmarttagJobResponseBody
	GetResults() *QuerySmarttagJobResponseBodyResults
	SetUserData(v string) *QuerySmarttagJobResponseBody
	GetUserData() *string
}

type QuerySmarttagJobResponseBody struct {
	// example:
	//
	// Success
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *QuerySmarttagJobResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// example:
	//
	// example UserData ****
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

func (s *QuerySmarttagJobResponseBody) GetMessage() *string {
	return s.Message
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

func (s *QuerySmarttagJobResponseBody) SetMessage(v string) *QuerySmarttagJobResponseBody {
	s.Message = &v
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
	// example:
	//
	// {"title":"example-title-****"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
