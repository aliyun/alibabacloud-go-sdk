// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaProducingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaProducingJob(v *GetMediaProducingJobResponseBodyMediaProducingJob) *GetMediaProducingJobResponseBody
	GetMediaProducingJob() *GetMediaProducingJobResponseBodyMediaProducingJob
	SetRequestId(v string) *GetMediaProducingJobResponseBody
	GetRequestId() *string
}

type GetMediaProducingJobResponseBody struct {
	// The details of the media producing job.
	MediaProducingJob *GetMediaProducingJobResponseBodyMediaProducingJob `json:"MediaProducingJob,omitempty" xml:"MediaProducingJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****83B7-7F87-4792-BFE9-63CD2137****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaProducingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBody) GetMediaProducingJob() *GetMediaProducingJobResponseBodyMediaProducingJob {
	return s.MediaProducingJob
}

func (s *GetMediaProducingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaProducingJobResponseBody) SetMediaProducingJob(v *GetMediaProducingJobResponseBodyMediaProducingJob) *GetMediaProducingJobResponseBody {
	s.MediaProducingJob = v
	return s
}

func (s *GetMediaProducingJobResponseBody) SetRequestId(v string) *GetMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaProducingJobResponseBody) Validate() error {
	if s.MediaProducingJob != nil {
		if err := s.MediaProducingJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaProducingJobResponseBodyMediaProducingJob struct {
	// The clip parameters of the template used for the job.
	//
	// example:
	//
	// {"VideoArray":["****05512043f49f697f7425****","****05512043f49f697f7425****","****05512043f49f697f7425****"]}
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The error code of the media producing job. 	Notice: This parameter is returned only if the job fails.
	//
	// example:
	//
	// ExceededMaximumValue
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time in UTC when the media producing job was completed, formatted as `YYYY-MM-DD\\"T\\"hh:mm:ss\\"Z\\"`.
	//
	// example:
	//
	// 2020-12-23T13:33:52Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time in UTC when the media producing job was created, formatted as `YYYY-MM-DD\\"T\\"hh:mm:ss\\"Z\\"`.
	//
	// example:
	//
	// 2020-12-23T13:33:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration of the output media. 	Notice: This parameter is returned only if the job is successful and the output is an audio or video file.
	//
	// example:
	//
	// 30.500000
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the output media asset.
	//
	// example:
	//
	// ****0cc6ba49eab379332c5b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The URL of the output media file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example2.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The error message of the media producing job.
	//
	// 	Notice:
	//
	// This parameter is returned only if the job fails.
	//
	// example:
	//
	// The specified "Width_Height" has exceeded maximum value.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time in UTC when the media producing job was last modified, formatted as `YYYY-MM-DD\\"T\\"hh:mm:ss\\"Z\\"`.
	//
	// example:
	//
	// 2020-12-23T13:33:49Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The progress of the job. Valid values: 0 to 100.
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The project ID.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the media producing job. Valid values:
	//
	// - `Init` (Initialized)
	//
	// - `Queuing` (In queue)
	//
	// - `Processing` (In progress)
	//
	// - `Success` (Successful)
	//
	// - `Failed` (Failed)
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// If the job is a subtask of a batch smart one-click video creation task, this parameter contains the subtask\\"s materials, such as the title and speech script.
	//
	// example:
	//
	// {"Title": "标题", "SpeechText": "批量智能一键成片口播文案"}
	SubJobMaterials *string `json:"SubJobMaterials,omitempty" xml:"SubJobMaterials,omitempty"`
	// The ID of the template used for the job.
	//
	// example:
	//
	// ****6e76134d739cc3e85d3e****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The timeline of the media producing job.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The user data, in JSON format.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx","Key":"Valuexxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the media asset in VOD (Video on Demand). This ID is returned if the job\\"s output is stored in VOD.
	//
	// example:
	//
	// ****332c5b0cc6ba49eab379****
	VodMediaId *string `json:"VodMediaId,omitempty" xml:"VodMediaId,omitempty"`
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetCode() *string {
	return s.Code
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetDuration() *float32 {
	return s.Duration
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetMediaURL() *string {
	return s.MediaURL
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetMessage() *string {
	return s.Message
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetProgress() *int32 {
	return s.Progress
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetStatus() *string {
	return s.Status
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetSubJobMaterials() *string {
	return s.SubJobMaterials
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetTimeline() *string {
	return s.Timeline
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) GetVodMediaId() *string {
	return s.VodMediaId
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetClipsParam(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ClipsParam = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCode(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Code = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCompleteTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CompleteTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCreateTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CreateTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetDuration(v float32) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Duration = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetJobId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.JobId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaURL(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaURL = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMessage(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Message = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetModifiedTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetProgress(v int32) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Progress = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetProjectId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ProjectId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetStatus(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Status = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetSubJobMaterials(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.SubJobMaterials = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTemplateId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.TemplateId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTimeline(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Timeline = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetUserData(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.UserData = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetVodMediaId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.VodMediaId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) Validate() error {
	return dara.Validate(s)
}
