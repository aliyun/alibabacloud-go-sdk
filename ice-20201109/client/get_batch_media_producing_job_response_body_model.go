// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchMediaProducingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEditingBatchJob(v *GetBatchMediaProducingJobResponseBodyEditingBatchJob) *GetBatchMediaProducingJobResponseBody
	GetEditingBatchJob() *GetBatchMediaProducingJobResponseBodyEditingBatchJob
	SetRequestId(v string) *GetBatchMediaProducingJobResponseBody
	GetRequestId() *string
}

type GetBatchMediaProducingJobResponseBody struct {
	// The information about the quick video production job.
	EditingBatchJob *GetBatchMediaProducingJobResponseBodyEditingBatchJob `json:"EditingBatchJob,omitempty" xml:"EditingBatchJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBatchMediaProducingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchMediaProducingJobResponseBody) GetEditingBatchJob() *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	return s.EditingBatchJob
}

func (s *GetBatchMediaProducingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchMediaProducingJobResponseBody) SetEditingBatchJob(v *GetBatchMediaProducingJobResponseBodyEditingBatchJob) *GetBatchMediaProducingJobResponseBody {
	s.EditingBatchJob = v
	return s
}

func (s *GetBatchMediaProducingJobResponseBody) SetRequestId(v string) *GetBatchMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBody) Validate() error {
	if s.EditingBatchJob != nil {
		if err := s.EditingBatchJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchMediaProducingJobResponseBodyEditingBatchJob struct {
	// The time when the job was complete.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-06-13T08:57:07Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-06-13T08:47:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The editing configurations. For more information, see [EditingConfig](~~2692547#1be9bba03b7qu~~).
	//
	// example:
	//
	// {
	//
	//   "MediaConfig": {
	//
	//       "Volume": 0
	//
	//   },
	//
	//   "SpeechConfig": {
	//
	//       "Volume": 1
	//
	//   },
	//
	//  "BackgroundMusicConfig": {
	//
	//       "Volume": 0.3
	//
	//   }
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The extended information. This parameter contains the following fields:
	//
	// ErrorCode: the error code of the main job.
	//
	// ErrorMessage: the error message of the main job.
	//
	// example:
	//
	// {
	//
	// 	"ErrorCode": "InvalidMaterial.NotFound",
	//
	// 	"ErrorMessage": "The specified clips id not found:[\\"****30d0b5e871eebb2ff7f6c75a****\\"]"
	//
	// }
	Extend       *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	ExtendInput  *string `json:"ExtendInput,omitempty" xml:"ExtendInput,omitempty"`
	ExtendOutput *string `json:"ExtendOutput,omitempty" xml:"ExtendOutput,omitempty"`
	// The input configurations. For more information, see [InputConfig](~~2692547#2faed1559549n~~).
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****b6b2750d4308892ac3330238****
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the job was last modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-06-13T08:57:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The output configurations. For more information, see [OutputConfig](~~2692547#447b928fcbuoa~~).
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://xxx.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 20,
	//
	//   "MaxDuration": 15,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920,
	//
	//   "Video": {"Crf": 27}
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The job state. Valid values:
	//
	// Init: The job is initialized.
	//
	// Processing: The job is in progress.
	//
	// Finished: The job is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The quick video production subjobs.
	SubJobList []*GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList `json:"SubJobList,omitempty" xml:"SubJobList,omitempty" type:"Repeated"`
	// The user-defined data, including the business and callback configurations. For more information, see [UserData](https://help.aliyun.com/document_detail/357745.html).
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetBatchMediaProducingJobResponseBodyEditingBatchJob) String() string {
	return dara.Prettify(s)
}

func (s GetBatchMediaProducingJobResponseBodyEditingBatchJob) GoString() string {
	return s.String()
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetExtend() *string {
	return s.Extend
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetExtendInput() *string {
	return s.ExtendInput
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetExtendOutput() *string {
	return s.ExtendOutput
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetInputConfig() *string {
	return s.InputConfig
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetJobId() *string {
	return s.JobId
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetJobType() *string {
	return s.JobType
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetStatus() *string {
	return s.Status
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetSubJobList() []*GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	return s.SubJobList
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) GetUserData() *string {
	return s.UserData
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetCompleteTime(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.CompleteTime = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetCreateTime(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.CreateTime = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetEditingConfig(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.EditingConfig = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetExtend(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.Extend = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetExtendInput(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.ExtendInput = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetExtendOutput(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.ExtendOutput = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetInputConfig(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.InputConfig = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetJobId(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.JobId = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetJobType(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.JobType = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetModifiedTime(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetOutputConfig(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.OutputConfig = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetStatus(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.Status = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetSubJobList(v []*GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.SubJobList = v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) SetUserData(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJob {
	s.UserData = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJob) Validate() error {
	if s.SubJobList != nil {
		for _, item := range s.SubJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList struct {
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The error code that is returned if the subjob failed. This parameter is not returned if the subjob is successful.
	//
	// example:
	//
	// InvalidMaterial.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the subjob failed. This parameter is not returned if the subjob is successful.
	//
	// example:
	//
	// The specified clips id not found:["****30d0b5e871eebb2ff7f6c75a****"]
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The subjob ID.
	//
	// example:
	//
	// ****8e81933d44e3ae69e2f81485****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the output media asset.
	//
	// example:
	//
	// ****1470b11171ee9d19e7e6c66a****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the output file.
	//
	// example:
	//
	// http:/xxx.oss-cn-shanghai.aliyuncs.com/xxx_0.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ****7cc47fe04eaa81bd853acb6a****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The subjob state. Valid values:
	//
	// Init: The subjob is initialized.
	//
	// Processing: The subjob is in progress.
	//
	// Success: The subjob is successful.
	//
	// Failed: The subjob failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GoString() string {
	return s.String()
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetDuration() *float32 {
	return s.Duration
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetJobId() *string {
	return s.JobId
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetMediaURL() *string {
	return s.MediaURL
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) GetStatus() *string {
	return s.Status
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetDuration(v float32) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.Duration = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetErrorCode(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.ErrorCode = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetErrorMessage(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetJobId(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.JobId = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetMediaId(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.MediaId = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetMediaURL(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.MediaURL = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetProjectId(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) SetStatus(v string) *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList {
	s.Status = &v
	return s
}

func (s *GetBatchMediaProducingJobResponseBodyEditingBatchJobSubJobList) Validate() error {
	return dara.Validate(s)
}
