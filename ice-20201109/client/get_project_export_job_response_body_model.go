// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectExportJob(v *GetProjectExportJobResponseBodyProjectExportJob) *GetProjectExportJobResponseBody
	GetProjectExportJob() *GetProjectExportJobResponseBodyProjectExportJob
	SetRequestId(v string) *GetProjectExportJobResponseBody
	GetRequestId() *string
}

type GetProjectExportJobResponseBody struct {
	ProjectExportJob *GetProjectExportJobResponseBodyProjectExportJob `json:"ProjectExportJob,omitempty" xml:"ProjectExportJob,omitempty" type:"Struct"`
	// example:
	//
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobResponseBody) GetProjectExportJob() *GetProjectExportJobResponseBodyProjectExportJob {
	return s.ProjectExportJob
}

func (s *GetProjectExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectExportJobResponseBody) SetProjectExportJob(v *GetProjectExportJobResponseBodyProjectExportJob) *GetProjectExportJobResponseBody {
	s.ProjectExportJob = v
	return s
}

func (s *GetProjectExportJobResponseBody) SetRequestId(v string) *GetProjectExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProjectExportJobResponseBodyProjectExportJob struct {
	// example:
	//
	// InvalidParameter
	Code         *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	ExportResult *GetProjectExportJobResponseBodyProjectExportJobExportResult `json:"ExportResult,omitempty" xml:"ExportResult,omitempty" type:"Struct"`
	// example:
	//
	// BaseTimeline
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// The specified parameter is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx","Key":"Valuexxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetProjectExportJobResponseBodyProjectExportJob) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobResponseBodyProjectExportJob) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetCode() *string {
	return s.Code
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetExportResult() *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	return s.ExportResult
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetExportType() *string {
	return s.ExportType
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetJobId() *string {
	return s.JobId
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetMessage() *string {
	return s.Message
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetStatus() *string {
	return s.Status
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) GetUserData() *string {
	return s.UserData
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetCode(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.Code = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetExportResult(v *GetProjectExportJobResponseBodyProjectExportJobExportResult) *GetProjectExportJobResponseBodyProjectExportJob {
	s.ExportResult = v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetExportType(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.ExportType = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetJobId(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.JobId = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetMessage(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.Message = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetProjectId(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.ProjectId = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetStatus(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.Status = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) SetUserData(v string) *GetProjectExportJobResponseBodyProjectExportJob {
	s.UserData = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJob) Validate() error {
	return dara.Validate(s)
}

type GetProjectExportJobResponseBodyProjectExportJobExportResult struct {
	ProjectUrl *string `json:"ProjectUrl,omitempty" xml:"ProjectUrl,omitempty"`
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"Type":"Video","MediaId":"****4d7cf14dc7b83b0e801c****","MediaURL":"https://test-bucket.oss-cn-shanghai.aliyuncs.com/test.mp4","TimelineIn":0.0,"TimelineOut":5.0,"In":0.0,"Out":5.0,"Speed":1.0,"Duration":5.0,"VirginDuration":13.334,"Height":1.0,"Width":1.0,"X":0.0,"Y":0.0}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
}

func (s GetProjectExportJobResponseBodyProjectExportJobExportResult) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobResponseBodyProjectExportJobExportResult) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetProjectUrl() *string {
	return s.ProjectUrl
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetTimeline() *string {
	return s.Timeline
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetProjectUrl(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.ProjectUrl = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetTimeline(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.Timeline = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) Validate() error {
	return dara.Validate(s)
}
