// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeProjectExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectExportJob(v *GetYikeProjectExportJobResponseBodyProjectExportJob) *GetYikeProjectExportJobResponseBody
	GetProjectExportJob() *GetYikeProjectExportJobResponseBodyProjectExportJob
	SetRequestId(v string) *GetYikeProjectExportJobResponseBody
	GetRequestId() *string
}

type GetYikeProjectExportJobResponseBody struct {
	// The export task object.
	ProjectExportJob *GetYikeProjectExportJobResponseBodyProjectExportJob `json:"ProjectExportJob,omitempty" xml:"ProjectExportJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeProjectExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobResponseBody) GetProjectExportJob() *GetYikeProjectExportJobResponseBodyProjectExportJob {
	return s.ProjectExportJob
}

func (s *GetYikeProjectExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeProjectExportJobResponseBody) SetProjectExportJob(v *GetYikeProjectExportJobResponseBodyProjectExportJob) *GetYikeProjectExportJobResponseBody {
	s.ProjectExportJob = v
	return s
}

func (s *GetYikeProjectExportJobResponseBody) SetRequestId(v string) *GetYikeProjectExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBody) Validate() error {
	if s.ProjectExportJob != nil {
		if err := s.ProjectExportJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeProjectExportJobResponseBodyProjectExportJob struct {
	// The task failure code.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The export result.
	ExportResult *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult `json:"ExportResult,omitempty" xml:"ExportResult,omitempty" type:"Struct"`
	// The project export type.
	//
	// example:
	//
	// PureAudio
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task failure message.
	//
	// example:
	//
	// The specified parameter is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The online editing project ID.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The custom data.
	//
	// example:
	//
	// {\\"testKey\\":\\"testValue\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJob) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJob) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetCode() *string {
	return s.Code
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetExportResult() *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult {
	return s.ExportResult
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetExportType() *string {
	return s.ExportType
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetMessage() *string {
	return s.Message
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetStatus() *string {
	return s.Status
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) GetUserData() *string {
	return s.UserData
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetCode(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.Code = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetExportResult(v *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.ExportResult = v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetExportType(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.ExportType = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetJobId(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.JobId = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetMessage(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.Message = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetProjectId(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.ProjectId = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetStatus(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.Status = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) SetUserData(v string) *GetYikeProjectExportJobResponseBodyProjectExportJob {
	s.UserData = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJob) Validate() error {
	if s.ExportResult != nil {
		if err := s.ExportResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeProjectExportJobResponseBodyProjectExportJobExportResult struct {
	// The download URL of the audio file.
	//
	// example:
	//
	// http://....MP3?Expires=
	AudioUrl *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	// The download URL of the PR project file (not supported).
	//
	// example:
	//
	// ....
	ProjectUrl *string `json:"ProjectUrl,omitempty" xml:"ProjectUrl,omitempty"`
	// The subtitle list.
	SrtList []*GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList `json:"SrtList,omitempty" xml:"SrtList,omitempty" type:"Repeated"`
	// The editing timeline (not supported).
	//
	// example:
	//
	// ....
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) GetAudioUrl() *string {
	return s.AudioUrl
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) GetProjectUrl() *string {
	return s.ProjectUrl
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) GetSrtList() []*GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	return s.SrtList
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) GetTimeline() *string {
	return s.Timeline
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) SetAudioUrl(v string) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult {
	s.AudioUrl = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) SetProjectUrl(v string) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult {
	s.ProjectUrl = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) SetSrtList(v []*GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult {
	s.SrtList = v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) SetTimeline(v string) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult {
	s.Timeline = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResult) Validate() error {
	if s.SrtList != nil {
		for _, item := range s.SrtList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList struct {
	// The download URL of the SRT file.
	//
	// example:
	//
	// http://xxx?Expires=
	SrtUrl *string `json:"SrtUrl,omitempty" xml:"SrtUrl,omitempty"`
	// The type enumeration. Currently, only VoiceOver is supported.
	//
	// example:
	//
	// VoiceOver
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GetSrtUrl() *string {
	return s.SrtUrl
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GetTag() *string {
	return s.Tag
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) SetSrtUrl(v string) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	s.SrtUrl = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) SetTag(v string) *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	s.Tag = &v
	return s
}

func (s *GetYikeProjectExportJobResponseBodyProjectExportJobExportResultSrtList) Validate() error {
	return dara.Validate(s)
}
