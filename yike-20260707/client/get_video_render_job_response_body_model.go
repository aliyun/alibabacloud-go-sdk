// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRenderJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetVideoRenderJobResponseBodyJob) *GetVideoRenderJobResponseBody
	GetJob() *GetVideoRenderJobResponseBodyJob
	SetRequestId(v string) *GetVideoRenderJobResponseBody
	GetRequestId() *string
}

type GetVideoRenderJobResponseBody struct {
	Job *GetVideoRenderJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoRenderJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRenderJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoRenderJobResponseBody) GetJob() *GetVideoRenderJobResponseBodyJob {
	return s.Job
}

func (s *GetVideoRenderJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoRenderJobResponseBody) SetJob(v *GetVideoRenderJobResponseBodyJob) *GetVideoRenderJobResponseBody {
	s.Job = v
	return s
}

func (s *GetVideoRenderJobResponseBody) SetRequestId(v string) *GetVideoRenderJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoRenderJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoRenderJobResponseBodyJob struct {
	// example:
	//
	// 8239345231244512***
	EditingProjectId *string `json:"EditingProjectId,omitempty" xml:"EditingProjectId,omitempty"`
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
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// https://xxxx.mp4
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

func (s GetVideoRenderJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRenderJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetVideoRenderJobResponseBodyJob) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *GetVideoRenderJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoRenderJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoRenderJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoRenderJobResponseBodyJob) GetLanguage() *string {
	return s.Language
}

func (s *GetVideoRenderJobResponseBodyJob) GetResult() *string {
	return s.Result
}

func (s *GetVideoRenderJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetVideoRenderJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoRenderJobResponseBodyJob) SetEditingProjectId(v string) *GetVideoRenderJobResponseBodyJob {
	s.EditingProjectId = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetErrorCode(v string) *GetVideoRenderJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetErrorMessage(v string) *GetVideoRenderJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetJobId(v string) *GetVideoRenderJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetLanguage(v string) *GetVideoRenderJobResponseBodyJob {
	s.Language = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetResult(v string) *GetVideoRenderJobResponseBodyJob {
	s.Result = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetStatus(v string) *GetVideoRenderJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) SetUserData(v string) *GetVideoRenderJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *GetVideoRenderJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
