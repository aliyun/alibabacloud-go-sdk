// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaComprehensionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetMediaComprehensionJobResponseBodyJob) *GetMediaComprehensionJobResponseBody
	GetJob() *GetMediaComprehensionJobResponseBodyJob
	SetMediaComprehensionJob(v *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) *GetMediaComprehensionJobResponseBody
	GetMediaComprehensionJob() *GetMediaComprehensionJobResponseBodyMediaComprehensionJob
	SetRequestId(v string) *GetMediaComprehensionJobResponseBody
	GetRequestId() *string
}

type GetMediaComprehensionJobResponseBody struct {
	Job                   *GetMediaComprehensionJobResponseBodyJob                   `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	MediaComprehensionJob *GetMediaComprehensionJobResponseBodyMediaComprehensionJob `json:"MediaComprehensionJob,omitempty" xml:"MediaComprehensionJob,omitempty" type:"Struct"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaComprehensionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaComprehensionJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaComprehensionJobResponseBody) GetJob() *GetMediaComprehensionJobResponseBodyJob {
	return s.Job
}

func (s *GetMediaComprehensionJobResponseBody) GetMediaComprehensionJob() *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	return s.MediaComprehensionJob
}

func (s *GetMediaComprehensionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaComprehensionJobResponseBody) SetJob(v *GetMediaComprehensionJobResponseBodyJob) *GetMediaComprehensionJobResponseBody {
	s.Job = v
	return s
}

func (s *GetMediaComprehensionJobResponseBody) SetMediaComprehensionJob(v *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) *GetMediaComprehensionJobResponseBody {
	s.MediaComprehensionJob = v
	return s
}

func (s *GetMediaComprehensionJobResponseBody) SetRequestId(v string) *GetMediaComprehensionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	if s.MediaComprehensionJob != nil {
		if err := s.MediaComprehensionJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaComprehensionJobResponseBodyJob struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MediaIds     []*string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty" type:"Repeated"`
	// example:
	//
	// "{\\"source_video_url\\":\\"http://xxx.mp4\\",\\"narrative_overview\\":{******}}"
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaComprehensionJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaComprehensionJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetMediaIds() []*string {
	return s.MediaIds
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetResult() *string {
	return s.Result
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetMediaComprehensionJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetErrorCode(v string) *GetMediaComprehensionJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetErrorMessage(v string) *GetMediaComprehensionJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetMediaIds(v []*string) *GetMediaComprehensionJobResponseBodyJob {
	s.MediaIds = v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetResult(v string) *GetMediaComprehensionJobResponseBodyJob {
	s.Result = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetStatus(v string) *GetMediaComprehensionJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) SetUserData(v string) *GetMediaComprehensionJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

type GetMediaComprehensionJobResponseBodyMediaComprehensionJob struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// ******afaa6f37457******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******307e9971f1******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// "{\\"source_video_url\\":\\"http://xxx.mp4\\",\\"narrative_overview\\":{******}}"
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// Executing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaComprehensionJobResponseBodyMediaComprehensionJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GoString() string {
	return s.String()
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetResult() *string {
	return s.Result
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetState() *string {
	return s.State
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetErrorCode(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.ErrorCode = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetErrorMessage(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetJobId(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.JobId = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetMediaId(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.MediaId = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetResult(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.Result = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetState(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.State = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) SetUserData(v string) *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	s.UserData = &v
	return s
}

func (s *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) Validate() error {
	return dara.Validate(s)
}
