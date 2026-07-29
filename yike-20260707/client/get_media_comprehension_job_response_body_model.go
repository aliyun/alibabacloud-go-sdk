// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaComprehensionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaComprehensionJob(v *GetMediaComprehensionJobResponseBodyMediaComprehensionJob) *GetMediaComprehensionJobResponseBody
	GetMediaComprehensionJob() *GetMediaComprehensionJobResponseBodyMediaComprehensionJob
	SetRequestId(v string) *GetMediaComprehensionJobResponseBody
	GetRequestId() *string
}

type GetMediaComprehensionJobResponseBody struct {
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

func (s *GetMediaComprehensionJobResponseBody) GetMediaComprehensionJob() *GetMediaComprehensionJobResponseBodyMediaComprehensionJob {
	return s.MediaComprehensionJob
}

func (s *GetMediaComprehensionJobResponseBody) GetRequestId() *string {
	return s.RequestId
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
	if s.MediaComprehensionJob != nil {
		if err := s.MediaComprehensionJob.Validate(); err != nil {
			return err
		}
	}
	return nil
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
