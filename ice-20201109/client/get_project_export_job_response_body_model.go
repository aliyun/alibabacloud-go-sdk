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
	// The project export task.
	ProjectExportJob *GetProjectExportJobResponseBodyProjectExportJob `json:"ProjectExportJob,omitempty" xml:"ProjectExportJob,omitempty" type:"Struct"`
	// The request ID.
	//
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
	if s.ProjectExportJob != nil {
		if err := s.ProjectExportJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectExportJobResponseBodyProjectExportJob struct {
	// The error code of the project export task.
	//
	// 	Notice: Check this field when the task fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The export result.
	ExportResult *GetProjectExportJobResponseBodyProjectExportJobExportResult `json:"ExportResult,omitempty" xml:"ExportResult,omitempty" type:"Struct"`
	// The type of the project export. Valid values:
	//
	// - **BaseTimeline**: timeline.
	//
	// - **AdobePremierePro**: Adobe Premiere Pro project.
	//
	// example:
	//
	// BaseTimeline
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The ID of the project export task.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message of the project export task.
	//
	// 	Notice: Check this field when the task fails.
	//
	// example:
	//
	// The specified parameter is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the project export task. Valid values:
	//
	// - **Init**: initial state.
	//
	// - **Processing**: processing.
	//
	// - **Success**: succeeded.
	//
	// - **Failed**: failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The custom settings in JSON format.
	//
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
	if s.ExportResult != nil {
		if err := s.ExportResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectExportJobResponseBodyProjectExportJobExportResult struct {
	AudioUrl *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	// The file URL of the exported project, which is typically an authenticated OSS URL. This field is returned when the export type is AdobePremierePro.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example_prefix/exported_project_1e8c39a502c3436c84f88290cd713bf3.zip?Expires=1750331685&....
	ProjectUrl *string                                                               `json:"ProjectUrl,omitempty" xml:"ProjectUrl,omitempty"`
	SrtList    []*GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList `json:"SrtList,omitempty" xml:"SrtList,omitempty" type:"Repeated"`
	// The online editing timeline. This field is returned when the export type is BaseTimeline. For more information about the structure, see [Timeline configuration](https://help.aliyun.com/document_detail/198823.html).
	//
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

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetAudioUrl() *string {
	return s.AudioUrl
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetProjectUrl() *string {
	return s.ProjectUrl
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetSrtList() []*GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	return s.SrtList
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) GetTimeline() *string {
	return s.Timeline
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetAudioUrl(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.AudioUrl = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetProjectUrl(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.ProjectUrl = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetSrtList(v []*GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.SrtList = v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) SetTimeline(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResult {
	s.Timeline = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResult) Validate() error {
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

type GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList struct {
	SrtUrl *string `json:"SrtUrl,omitempty" xml:"SrtUrl,omitempty"`
	Tag    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GetSrtUrl() *string {
	return s.SrtUrl
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) GetTag() *string {
	return s.Tag
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) SetSrtUrl(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	s.SrtUrl = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) SetTag(v string) *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList {
	s.Tag = &v
	return s
}

func (s *GetProjectExportJobResponseBodyProjectExportJobExportResultSrtList) Validate() error {
	return dara.Validate(s)
}
