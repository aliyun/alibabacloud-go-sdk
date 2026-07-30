// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemakeScriptJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetRemakeScriptJobResponseBodyJob) *GetRemakeScriptJobResponseBody
	GetJob() *GetRemakeScriptJobResponseBodyJob
	SetRequestId(v string) *GetRemakeScriptJobResponseBody
	GetRequestId() *string
}

type GetRemakeScriptJobResponseBody struct {
	// The generated creative script object.
	Job *GetRemakeScriptJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRemakeScriptJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRemakeScriptJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemakeScriptJobResponseBody) GetJob() *GetRemakeScriptJobResponseBodyJob {
	return s.Job
}

func (s *GetRemakeScriptJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRemakeScriptJobResponseBody) SetJob(v *GetRemakeScriptJobResponseBodyJob) *GetRemakeScriptJobResponseBody {
	s.Job = v
	return s
}

func (s *GetRemakeScriptJobResponseBody) SetRequestId(v string) *GetRemakeScriptJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRemakeScriptJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRemakeScriptJobResponseBodyJob struct {
	// The error code. Returned when the status is Failed.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. Returned when the status is Failed.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The file URL of the script imitation result. The file content is in JSON format of the creative script.
	//
	// example:
	//
	// http://xxxx.json
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The file status. Valid values:
	//
	// - Created: Created.
	//
	// - Executing: Executing.
	//
	// - Finished: Finished.
	//
	// - Failed: Failed.
	//
	// - Deleted: Deleted.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined parameter, in JSON format string.
	//
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetRemakeScriptJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetRemakeScriptJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetRemakeScriptJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRemakeScriptJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRemakeScriptJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetRemakeScriptJobResponseBodyJob) GetResult() *string {
	return s.Result
}

func (s *GetRemakeScriptJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetRemakeScriptJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *GetRemakeScriptJobResponseBodyJob) SetErrorCode(v string) *GetRemakeScriptJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) SetErrorMessage(v string) *GetRemakeScriptJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) SetJobId(v string) *GetRemakeScriptJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) SetResult(v string) *GetRemakeScriptJobResponseBodyJob {
	s.Result = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) SetStatus(v string) *GetRemakeScriptJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) SetUserData(v string) *GetRemakeScriptJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *GetRemakeScriptJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
