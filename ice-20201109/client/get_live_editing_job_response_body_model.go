// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveEditingJob(v *GetLiveEditingJobResponseBodyLiveEditingJob) *GetLiveEditingJobResponseBody
	GetLiveEditingJob() *GetLiveEditingJobResponseBodyLiveEditingJob
	SetRequestId(v string) *GetLiveEditingJobResponseBody
	GetRequestId() *string
}

type GetLiveEditingJobResponseBody struct {
	// The information about the live editing job.
	LiveEditingJob *GetLiveEditingJobResponseBodyLiveEditingJob `json:"LiveEditingJob,omitempty" xml:"LiveEditingJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLiveEditingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponseBody) GetLiveEditingJob() *GetLiveEditingJobResponseBodyLiveEditingJob {
	return s.LiveEditingJob
}

func (s *GetLiveEditingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveEditingJobResponseBody) SetLiveEditingJob(v *GetLiveEditingJobResponseBodyLiveEditingJob) *GetLiveEditingJobResponseBody {
	s.LiveEditingJob = v
	return s
}

func (s *GetLiveEditingJobResponseBody) SetRequestId(v string) *GetLiveEditingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveEditingJobResponseBody) Validate() error {
	if s.LiveEditingJob != nil {
		if err := s.LiveEditingJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveEditingJobResponseBodyLiveEditingJob struct {
	// The clips.
	//
	// example:
	//
	// [{\\"StartTime\\": \\" 2021-06-21T08:01:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:03:00Z\\" }]
	Clips *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	// The response code. Note: Pay attention to this parameter if the job failed.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the live editing job was completed. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:52Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the live editing job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the live editing job.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The live editing configurations.
	LiveStreamConfig *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig `json:"LiveStreamConfig,omitempty" xml:"LiveStreamConfig,omitempty" type:"Struct"`
	// The media asset ID of the output file.
	//
	// example:
	//
	// ****0cc6ba49eab379332c5b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The production configurations.
	MediaProduceConfig *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig `json:"MediaProduceConfig,omitempty" xml:"MediaProduceConfig,omitempty" type:"Struct"`
	// The URL of the output file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example2.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// The returned message. Note: Pay attention to this parameter if the job failed.
	//
	// example:
	//
	// The specific parameter LiveStreamConfig is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the live editing job was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:49Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The storage configurations of the output file.
	OutputMediaConfig *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty" type:"Struct"`
	// The ID of the live editing project.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The state of the live editing job. Valid values: Init, Queuing, Processing, Success, and Failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"key": "value\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetLiveEditingJobResponseBodyLiveEditingJob) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponseBodyLiveEditingJob) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetClips() *string {
	return s.Clips
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetCode() *string {
	return s.Code
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetLiveStreamConfig() *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig {
	return s.LiveStreamConfig
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetMediaId() *string {
	return s.MediaId
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetMediaProduceConfig() *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig {
	return s.MediaProduceConfig
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetMediaURL() *string {
	return s.MediaURL
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetMessage() *string {
	return s.Message
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetOutputMediaConfig() *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	return s.OutputMediaConfig
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetStatus() *string {
	return s.Status
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) GetUserData() *string {
	return s.UserData
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetClips(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.Clips = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetCode(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.Code = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetCompleteTime(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.CompleteTime = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetCreationTime(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.CreationTime = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetJobId(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.JobId = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetLiveStreamConfig(v *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.LiveStreamConfig = v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetMediaId(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.MediaId = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetMediaProduceConfig(v *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.MediaProduceConfig = v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetMediaURL(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.MediaURL = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetMessage(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.Message = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetModifiedTime(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetOutputMediaConfig(v *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.OutputMediaConfig = v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetProjectId(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.ProjectId = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetStatus(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.Status = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) SetUserData(v string) *GetLiveEditingJobResponseBodyLiveEditingJob {
	s.UserData = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJob) Validate() error {
	if s.LiveStreamConfig != nil {
		if err := s.LiveStreamConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MediaProduceConfig != nil {
		if err := s.MediaProduceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OutputMediaConfig != nil {
		if err := s.OutputMediaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name of the live stream.
	//
	// example:
	//
	// domain.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// streamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) GetAppName() *string {
	return s.AppName
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) SetAppName(v string) *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig {
	s.AppName = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) SetDomainName(v string) *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig {
	s.DomainName = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) SetStreamName(v string) *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig {
	s.StreamName = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobLiveStreamConfig) Validate() error {
	return dara.Validate(s)
}

type GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig struct {
	// The editing mode. Default value: Accurate.
	//
	// example:
	//
	// Accurate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) GetMode() *string {
	return s.Mode
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) SetMode(v string) *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig {
	s.Mode = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobMediaProduceConfig) Validate() error {
	return dara.Validate(s)
}

type GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig struct {
	// The bitrate of the output file. Unit: Kbit/s. You can leave this parameter empty. The default value is the maximum bitrate of the input materials.
	//
	// example:
	//
	// 1000
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// If OutputMediaTarget is set to vod-media, this parameter indicates the file name of the output file. The value contains the file name extension but not the path.
	//
	// example:
	//
	// test.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The height of the output file. You can leave this parameter empty. The default value is the maximum height of the input materials.
	//
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The URL of the output file.
	//
	// example:
	//
	// https://testice-testbucket.oss-cn-shanghai.aliyuncs.com/test.mp4
	MediaURL *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	// If OutputMediaTarget is set to vod-media, this parameter indicates the storage location of the media asset in ApsaraVideo VOD. The storage location is the path of the file in ApsaraVideo VOD, excluding the prefix http://. Example: outin-xxxxxx.oss-cn-shanghai.aliyuncs.com.
	//
	// example:
	//
	// outin-xxxxxx.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The ID of the VOD transcoding template group. If VOD transcoding is not required, set the value to VOD_NO_TRANSCODE.
	//
	// example:
	//
	// VOD_NO_TRANSCODE
	VodTemplateGroupId *string `json:"VodTemplateGroupId,omitempty" xml:"VodTemplateGroupId,omitempty"`
	// The width of the output file. You can leave this parameter empty. The default value is the maximum width of the input materials.
	//
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetFileName() *string {
	return s.FileName
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetHeight() *int32 {
	return s.Height
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetMediaURL() *string {
	return s.MediaURL
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetVodTemplateGroupId() *string {
	return s.VodTemplateGroupId
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) GetWidth() *int32 {
	return s.Width
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetBitrate(v int64) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.Bitrate = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetFileName(v string) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.FileName = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetHeight(v int32) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.Height = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetMediaURL(v string) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.MediaURL = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetStorageLocation(v string) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.StorageLocation = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetVodTemplateGroupId(v string) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.VodTemplateGroupId = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) SetWidth(v int32) *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig {
	s.Width = &v
	return s
}

func (s *GetLiveEditingJobResponseBodyLiveEditingJobOutputMediaConfig) Validate() error {
	return dara.Validate(s)
}
