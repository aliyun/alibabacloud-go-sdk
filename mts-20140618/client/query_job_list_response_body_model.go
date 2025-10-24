// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v *QueryJobListResponseBodyJobList) *QueryJobListResponseBody
	GetJobList() *QueryJobListResponseBodyJobList
	SetNonExistJobIds(v *QueryJobListResponseBodyNonExistJobIds) *QueryJobListResponseBody
	GetNonExistJobIds() *QueryJobListResponseBodyNonExistJobIds
	SetRequestId(v string) *QueryJobListResponseBody
	GetRequestId() *string
}

type QueryJobListResponseBody struct {
	// The transcoding jobs.
	JobList *QueryJobListResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
	// The list of nonexistent job IDs. If all queried job IDs exist, the response does not contain this parameter.
	NonExistJobIds *QueryJobListResponseBodyNonExistJobIds `json:"NonExistJobIds,omitempty" xml:"NonExistJobIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 197ADF44-104C-514C-9F92-D8924CB34E2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBody) GetJobList() *QueryJobListResponseBodyJobList {
	return s.JobList
}

func (s *QueryJobListResponseBody) GetNonExistJobIds() *QueryJobListResponseBodyNonExistJobIds {
	return s.NonExistJobIds
}

func (s *QueryJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJobListResponseBody) SetJobList(v *QueryJobListResponseBodyJobList) *QueryJobListResponseBody {
	s.JobList = v
	return s
}

func (s *QueryJobListResponseBody) SetNonExistJobIds(v *QueryJobListResponseBodyNonExistJobIds) *QueryJobListResponseBody {
	s.NonExistJobIds = v
	return s
}

func (s *QueryJobListResponseBody) SetRequestId(v string) *QueryJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryJobListResponseBody) Validate() error {
	if s.JobList != nil {
		if err := s.JobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistJobIds != nil {
		if err := s.NonExistJobIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobList struct {
	Job []*QueryJobListResponseBodyJobListJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobList) GetJob() []*QueryJobListResponseBodyJobListJob {
	return s.Job
}

func (s *QueryJobListResponseBodyJobList) SetJob(v []*QueryJobListResponseBodyJobListJob) *QueryJobListResponseBodyJobList {
	s.Job = v
	return s
}

func (s *QueryJobListResponseBodyJobList) Validate() error {
	if s.Job != nil {
		for _, item := range s.Job {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJob struct {
	// The error code returned if the job failed. If the job was successful, this parameter is not returned.
	//
	// example:
	//
	// InvalidParameter.NullValue
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2014-01-10T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2014-01-10T12:20:25Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The information about the job input.
	Input *QueryJobListResponseBodyJobListJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// 31fa3c9ca8134fb4b0b0f7878301****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message sent by Message Service (MNS) to notify users of the job result.
	MNSMessageResult *QueryJobListResponseBodyJobListJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the job failed. If the job was successful, this parameter is not returned.
	//
	// example:
	//
	// The specified parameter "%s" cannot be null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The job output.
	Output *QueryJobListResponseBodyJobListJobOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The transcoding progress.
	//
	// example:
	//
	// 100
	Percent *int64 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the MPS queue that is used to run the job.
	//
	// example:
	//
	// 88c6ca184c0e47b665e2a1267971****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job state. Valid values:
	//
	// 	- **Submitted**: The job was submitted.
	//
	// 	- **Transcoding**: Transcoding is in process.
	//
	// 	- **TranscodeSuccess**: The job was successful.
	//
	// 	- **TranscodeFail**: The job failed.
	//
	// 	- **TranscodeCancelled**: The job was canceled.
	//
	// example:
	//
	// TranscodeSuccess
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 2021-03-04T06:44:43Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s QueryJobListResponseBodyJobListJob) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJob) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJob) GetCode() *string {
	return s.Code
}

func (s *QueryJobListResponseBodyJobListJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryJobListResponseBodyJobListJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryJobListResponseBodyJobListJob) GetInput() *QueryJobListResponseBodyJobListJobInput {
	return s.Input
}

func (s *QueryJobListResponseBodyJobListJob) GetJobId() *string {
	return s.JobId
}

func (s *QueryJobListResponseBodyJobListJob) GetMNSMessageResult() *QueryJobListResponseBodyJobListJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *QueryJobListResponseBodyJobListJob) GetMessage() *string {
	return s.Message
}

func (s *QueryJobListResponseBodyJobListJob) GetOutput() *QueryJobListResponseBodyJobListJobOutput {
	return s.Output
}

func (s *QueryJobListResponseBodyJobListJob) GetPercent() *int64 {
	return s.Percent
}

func (s *QueryJobListResponseBodyJobListJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryJobListResponseBodyJobListJob) GetState() *string {
	return s.State
}

func (s *QueryJobListResponseBodyJobListJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *QueryJobListResponseBodyJobListJob) SetCode(v string) *QueryJobListResponseBodyJobListJob {
	s.Code = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetCreationTime(v string) *QueryJobListResponseBodyJobListJob {
	s.CreationTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetFinishTime(v string) *QueryJobListResponseBodyJobListJob {
	s.FinishTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetInput(v *QueryJobListResponseBodyJobListJobInput) *QueryJobListResponseBodyJobListJob {
	s.Input = v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetJobId(v string) *QueryJobListResponseBodyJobListJob {
	s.JobId = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetMNSMessageResult(v *QueryJobListResponseBodyJobListJobMNSMessageResult) *QueryJobListResponseBodyJobListJob {
	s.MNSMessageResult = v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetMessage(v string) *QueryJobListResponseBodyJobListJob {
	s.Message = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetOutput(v *QueryJobListResponseBodyJobListJobOutput) *QueryJobListResponseBodyJobListJob {
	s.Output = v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetPercent(v int64) *QueryJobListResponseBodyJobListJob {
	s.Percent = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetPipelineId(v string) *QueryJobListResponseBodyJobListJob {
	s.PipelineId = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetState(v string) *QueryJobListResponseBodyJobListJob {
	s.State = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) SetSubmitTime(v string) *QueryJobListResponseBodyJobListJob {
	s.SubmitTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJob) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobInput struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// exampleBucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input file resides.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input file.
	//
	// example:
	//
	// video_01.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobInput) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobInput) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *QueryJobListResponseBodyJobListJobInput) GetLocation() *string {
	return s.Location
}

func (s *QueryJobListResponseBodyJobListJobInput) GetObject() *string {
	return s.Object
}

func (s *QueryJobListResponseBodyJobListJobInput) SetBucket(v string) *QueryJobListResponseBodyJobListJobInput {
	s.Bucket = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobInput) SetLocation(v string) *QueryJobListResponseBodyJobListJobInput {
	s.Location = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobInput) SetObject(v string) *QueryJobListResponseBodyJobListJobInput {
	s.Object = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobInput) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobMNSMessageResult struct {
	// The error code returned if the job failed. This parameter is not returned if the job was successful.
	//
	// example:
	//
	// InvalidParameter.ResourceNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the job failed. This parameter is not returned if the job was successful.
	//
	// example:
	//
	// The resource operated “%s” cannot be found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message returned if the job was successful.
	//
	// example:
	//
	// 123
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) SetErrorCode(v string) *QueryJobListResponseBodyJobListJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) SetErrorMessage(v string) *QueryJobListResponseBodyJobListJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) SetMessageId(v string) *QueryJobListResponseBodyJobListJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutput struct {
	// The audio configurations.
	//
	// >  If this parameter is specified in the request, the corresponding parameters in the specified transcoding template are overwritten.
	Audio *QueryJobListResponseBodyJobListJobOutputAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The sequence number of the audio stream.
	//
	// 	- Format: 0:a:{Sequence number}. Example: 0:a:0.
	//
	// 	- The sequence number is the index of the audio stream in the list and starts from 0.
	//
	// 	- If no sequence number is specified, the default audio stream is used.
	//
	// example:
	//
	// 0:a:0
	AudioStreamMap *string `json:"AudioStreamMap,omitempty" xml:"AudioStreamMap,omitempty"`
	// The information about clips.
	Clip *QueryJobListResponseBodyJobListJobOutputClip `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	// The container format configurations.
	Container *QueryJobListResponseBodyJobListJobOutputContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The configurations of watermark blurring. The value is a JSON object. For more information, see the DeWatermark section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// example:
	//
	// {"0":[{"l":10,"t":10,"w":10,"h":10},{"l":100,"t":0.1,"w":10,"h":10}],"128000":[],"250000":[{"l":0.2,"t":0.1,"w":0.01,"h":0.05}]}
	DeWatermark *string `json:"DeWatermark,omitempty" xml:"DeWatermark,omitempty"`
	// The encryption configurations. The encrypted video file is generated in the M3U8 format.
	Encryption *QueryJobListResponseBodyJobListJobOutputEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The custom fields.
	//
	// example:
	//
	// testid-002
	ExtendData *string `json:"ExtendData,omitempty" xml:"ExtendData,omitempty"`
	// The non-standard support configurations for M3U8. The value is a JSON object. For more information, see the M3U8NonStandardSupport section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	M3U8NonStandardSupport *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport `json:"M3U8NonStandardSupport,omitempty" xml:"M3U8NonStandardSupport,omitempty" type:"Struct"`
	// The URL of the merging configuration file. Only one of MergeList and MergeConfigUrl takes effect.
	//
	// 	- The configuration file specified by MergeConfigUrl can contain up to 50 clips.
	//
	// 	- MergeConfigUrl indicates the URL of the configuration file for merging clips. Make sure that the configuration file is stored as an object in OSS and that MPS can access the OSS object. For information about the file content, see the details about merging parameters.
	//
	// 	- Example of the content of the merging configuration file: `{"MergeList":[{"MergeURL":"http://exampleBucket****.oss-cn-hangzhou.aliyuncs.com/video_01.mp4"}]}`.
	//
	// example:
	//
	// https://ceshi-***.oss-cn-shanghai.aliyuncs.com/ccc/p0903q9wkkb.m3u8
	MergeConfigUrl *string `json:"MergeConfigUrl,omitempty" xml:"MergeConfigUrl,omitempty"`
	// The configurations of clip merging. Up to four clips can be merged.
	MergeList *QueryJobListResponseBodyJobListJobOutputMergeList `json:"MergeList,omitempty" xml:"MergeList,omitempty" type:"Struct"`
	// The information about the high-speed transcoding job. This information is available only for jobs that are submitted by using an MPS queue for high-speed transcoding. This does not support MPS queues for high-speed transcoding of an earlier version.
	MultiSpeedInfo *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo `json:"MultiSpeedInfo,omitempty" xml:"MultiSpeedInfo,omitempty" type:"Struct"`
	// The transmuxing configurations. The transmuxing configurations. If this parameter is specified in the request, the corresponding parameters in the specified transcoding template are overwritten.
	MuxConfig *QueryJobListResponseBodyJobListJobOutputMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The opening parts. The value is a JSON object.
	OpeningList *QueryJobListResponseBodyJobListJobOutputOpeningList `json:"OpeningList,omitempty" xml:"OpeningList,omitempty" type:"Struct"`
	// The output captions.
	OutSubtitleList *QueryJobListResponseBodyJobListJobOutputOutSubtitleList `json:"OutSubtitleList,omitempty" xml:"OutSubtitleList,omitempty" type:"Struct"`
	// The details of the output file.
	OutputFile *QueryJobListResponseBodyJobListJobOutputOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// The priority of the job in the ApsaraVideo Media Processing (MPS) queue to which the job is added.
	//
	// 	- A value of 10 indicates the highest priority.
	//
	// 	- Default value: **6**.
	//
	// example:
	//
	// 6
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The media properties.
	Properties *QueryJobListResponseBodyJobListJobOutputProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The rotation angle of the video.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The caption configurations.
	SubtitleConfig *QueryJobListResponseBodyJobListJobOutputSubtitleConfig `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Struct"`
	// The configurations for using the resolution of the source video.
	SuperReso *QueryJobListResponseBodyJobListJobOutputSuperReso `json:"SuperReso,omitempty" xml:"SuperReso,omitempty" type:"Struct"`
	// The ending parts.
	TailSlateList *QueryJobListResponseBodyJobListJobOutputTailSlateList `json:"TailSlateList,omitempty" xml:"TailSlateList,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// S00000001-200010
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The general transcoding configurations.
	//
	// >  If this parameter is specified in the request, the corresponding parameters in the specified transcoding template are overwritten.
	TransConfig *QueryJobListResponseBodyJobListJobOutputTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video configurations.
	Video *QueryJobListResponseBodyJobListJobOutputVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
	// The sequence number of the video stream. The sequence number is the index of the video stream in the list and starts from 0. If no sequence number is specified, the default video stream is used.
	//
	// example:
	//
	// 0
	VideoStreamMap *string `json:"VideoStreamMap,omitempty" xml:"VideoStreamMap,omitempty"`
	// The URL of the watermark configuration file.
	//
	// example:
	//
	// http://example.com/configure
	WaterMarkConfigUrl *string `json:"WaterMarkConfigUrl,omitempty" xml:"WaterMarkConfigUrl,omitempty"`
	// The watermarks.
	WaterMarkList *QueryJobListResponseBodyJobListJobOutputWaterMarkList `json:"WaterMarkList,omitempty" xml:"WaterMarkList,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutput) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetAudio() *QueryJobListResponseBodyJobListJobOutputAudio {
	return s.Audio
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetAudioStreamMap() *string {
	return s.AudioStreamMap
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetClip() *QueryJobListResponseBodyJobListJobOutputClip {
	return s.Clip
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetContainer() *QueryJobListResponseBodyJobListJobOutputContainer {
	return s.Container
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetDeWatermark() *string {
	return s.DeWatermark
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetEncryption() *QueryJobListResponseBodyJobListJobOutputEncryption {
	return s.Encryption
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetExtendData() *string {
	return s.ExtendData
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetM3U8NonStandardSupport() *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport {
	return s.M3U8NonStandardSupport
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetMergeConfigUrl() *string {
	return s.MergeConfigUrl
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetMergeList() *QueryJobListResponseBodyJobListJobOutputMergeList {
	return s.MergeList
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetMultiSpeedInfo() *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	return s.MultiSpeedInfo
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetMuxConfig() *QueryJobListResponseBodyJobListJobOutputMuxConfig {
	return s.MuxConfig
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetOpeningList() *QueryJobListResponseBodyJobListJobOutputOpeningList {
	return s.OpeningList
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetOutSubtitleList() *QueryJobListResponseBodyJobListJobOutputOutSubtitleList {
	return s.OutSubtitleList
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetOutputFile() *QueryJobListResponseBodyJobListJobOutputOutputFile {
	return s.OutputFile
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetPriority() *string {
	return s.Priority
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetProperties() *QueryJobListResponseBodyJobListJobOutputProperties {
	return s.Properties
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetRotate() *string {
	return s.Rotate
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetSubtitleConfig() *QueryJobListResponseBodyJobListJobOutputSubtitleConfig {
	return s.SubtitleConfig
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetSuperReso() *QueryJobListResponseBodyJobListJobOutputSuperReso {
	return s.SuperReso
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetTailSlateList() *QueryJobListResponseBodyJobListJobOutputTailSlateList {
	return s.TailSlateList
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetTransConfig() *QueryJobListResponseBodyJobListJobOutputTransConfig {
	return s.TransConfig
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetUserData() *string {
	return s.UserData
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetVideo() *QueryJobListResponseBodyJobListJobOutputVideo {
	return s.Video
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetVideoStreamMap() *string {
	return s.VideoStreamMap
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetWaterMarkConfigUrl() *string {
	return s.WaterMarkConfigUrl
}

func (s *QueryJobListResponseBodyJobListJobOutput) GetWaterMarkList() *QueryJobListResponseBodyJobListJobOutputWaterMarkList {
	return s.WaterMarkList
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetAudio(v *QueryJobListResponseBodyJobListJobOutputAudio) *QueryJobListResponseBodyJobListJobOutput {
	s.Audio = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetAudioStreamMap(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.AudioStreamMap = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetClip(v *QueryJobListResponseBodyJobListJobOutputClip) *QueryJobListResponseBodyJobListJobOutput {
	s.Clip = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetContainer(v *QueryJobListResponseBodyJobListJobOutputContainer) *QueryJobListResponseBodyJobListJobOutput {
	s.Container = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetDeWatermark(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.DeWatermark = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetEncryption(v *QueryJobListResponseBodyJobListJobOutputEncryption) *QueryJobListResponseBodyJobListJobOutput {
	s.Encryption = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetExtendData(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.ExtendData = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetM3U8NonStandardSupport(v *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) *QueryJobListResponseBodyJobListJobOutput {
	s.M3U8NonStandardSupport = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetMergeConfigUrl(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.MergeConfigUrl = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetMergeList(v *QueryJobListResponseBodyJobListJobOutputMergeList) *QueryJobListResponseBodyJobListJobOutput {
	s.MergeList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetMultiSpeedInfo(v *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) *QueryJobListResponseBodyJobListJobOutput {
	s.MultiSpeedInfo = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetMuxConfig(v *QueryJobListResponseBodyJobListJobOutputMuxConfig) *QueryJobListResponseBodyJobListJobOutput {
	s.MuxConfig = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetOpeningList(v *QueryJobListResponseBodyJobListJobOutputOpeningList) *QueryJobListResponseBodyJobListJobOutput {
	s.OpeningList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetOutSubtitleList(v *QueryJobListResponseBodyJobListJobOutputOutSubtitleList) *QueryJobListResponseBodyJobListJobOutput {
	s.OutSubtitleList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetOutputFile(v *QueryJobListResponseBodyJobListJobOutputOutputFile) *QueryJobListResponseBodyJobListJobOutput {
	s.OutputFile = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetPriority(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.Priority = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetProperties(v *QueryJobListResponseBodyJobListJobOutputProperties) *QueryJobListResponseBodyJobListJobOutput {
	s.Properties = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetRotate(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.Rotate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetSubtitleConfig(v *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) *QueryJobListResponseBodyJobListJobOutput {
	s.SubtitleConfig = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetSuperReso(v *QueryJobListResponseBodyJobListJobOutputSuperReso) *QueryJobListResponseBodyJobListJobOutput {
	s.SuperReso = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetTailSlateList(v *QueryJobListResponseBodyJobListJobOutputTailSlateList) *QueryJobListResponseBodyJobListJobOutput {
	s.TailSlateList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetTemplateId(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.TemplateId = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetTransConfig(v *QueryJobListResponseBodyJobListJobOutputTransConfig) *QueryJobListResponseBodyJobListJobOutput {
	s.TransConfig = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetUserData(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.UserData = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetVideo(v *QueryJobListResponseBodyJobListJobOutputVideo) *QueryJobListResponseBodyJobListJobOutput {
	s.Video = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetVideoStreamMap(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.VideoStreamMap = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetWaterMarkConfigUrl(v string) *QueryJobListResponseBodyJobListJobOutput {
	s.WaterMarkConfigUrl = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) SetWaterMarkList(v *QueryJobListResponseBodyJobListJobOutputWaterMarkList) *QueryJobListResponseBodyJobListJobOutput {
	s.WaterMarkList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutput) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Clip != nil {
		if err := s.Clip.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.M3U8NonStandardSupport != nil {
		if err := s.M3U8NonStandardSupport.Validate(); err != nil {
			return err
		}
	}
	if s.MergeList != nil {
		if err := s.MergeList.Validate(); err != nil {
			return err
		}
	}
	if s.MultiSpeedInfo != nil {
		if err := s.MultiSpeedInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OpeningList != nil {
		if err := s.OpeningList.Validate(); err != nil {
			return err
		}
	}
	if s.OutSubtitleList != nil {
		if err := s.OutSubtitleList.Validate(); err != nil {
			return err
		}
	}
	if s.OutputFile != nil {
		if err := s.OutputFile.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleConfig != nil {
		if err := s.SubtitleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SuperReso != nil {
		if err := s.SuperReso.Validate(); err != nil {
			return err
		}
	}
	if s.TailSlateList != nil {
		if err := s.TailSlateList.Validate(); err != nil {
			return err
		}
	}
	if s.TransConfig != nil {
		if err := s.TransConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	if s.WaterMarkList != nil {
		if err := s.WaterMarkList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: **128**.
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels.
	//
	// 	- Valid values: 1, 2, 3, 4, 5, 6, 7, and 8.
	//
	// 	- Default value: **2**.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec.
	//
	// 	- Valid values: aac, mp3, vorbis, and flac.
	//
	// 	- Default value: **aac**.
	//
	// example:
	//
	// aac
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The codec profile of the audio. Valid values when the value of Codec is aac: aaclow, aache, aachev2, aacld, and aaceld.
	//
	// example:
	//
	// aaclow
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of quality control on the audio.
	//
	// example:
	//
	// 15
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The sampling rate.
	//
	// 	- Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// 	- Unit: Hz.
	//
	// 	- Default value: 44100.
	//
	// >  If the video container format is FLV and the audio codec is MP3, the value of this parameter cannot be 32000, 48000, or 96000. If the audio codec is MP3, the value of this parameter cannot be 96000.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume configurations.
	Volume *QueryJobListResponseBodyJobListJobOutputAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputAudio) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputAudio) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetChannels() *string {
	return s.Channels
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetCodec() *string {
	return s.Codec
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetProfile() *string {
	return s.Profile
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetQscale() *string {
	return s.Qscale
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) GetVolume() *QueryJobListResponseBodyJobListJobOutputAudioVolume {
	return s.Volume
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetChannels(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Channels = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetCodec(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Codec = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetProfile(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Profile = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetQscale(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Qscale = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetSamplerate(v string) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Samplerate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) SetVolume(v *QueryJobListResponseBodyJobListJobOutputAudioVolume) *QueryJobListResponseBodyJobListJobOutputAudio {
	s.Volume = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputAudioVolume struct {
	// The volume adjustment range. Default value: -20. Unit: dB.
	//
	// example:
	//
	// -20
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The method that is used to adjust the volume. Valid values:
	//
	// 	- **auto**
	//
	// 	- **dynamic**
	//
	// 	- **linear**
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputAudioVolume) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *QueryJobListResponseBodyJobListJobOutputAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *QueryJobListResponseBodyJobListJobOutputAudioVolume) SetLevel(v string) *QueryJobListResponseBodyJobListJobOutputAudioVolume {
	s.Level = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudioVolume) SetMethod(v string) *QueryJobListResponseBodyJobListJobOutputAudioVolume {
	s.Method = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputAudioVolume) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputClip struct {
	// The time span of the clip.
	TimeSpan *QueryJobListResponseBodyJobListJobOutputClipTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputClip) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputClip) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputClip) GetTimeSpan() *QueryJobListResponseBodyJobListJobOutputClipTimeSpan {
	return s.TimeSpan
}

func (s *QueryJobListResponseBodyJobListJobOutputClip) SetTimeSpan(v *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) *QueryJobListResponseBodyJobListJobOutputClip {
	s.TimeSpan = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputClip) Validate() error {
	if s.TimeSpan != nil {
		if err := s.TimeSpan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputClipTimeSpan struct {
	// The duration of the clip.
	//
	// 	- Format: `hh:mm:ss[.SSS]`.
	//
	// 	- Example: 01:00:59.999.
	//
	// Or
	//
	// 	- Format: `sssss[.SSS]`.
	//
	// 	- Example: 32000.23.
	//
	// example:
	//
	// 01:00:59.999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The point in time when the clip starts.
	//
	// 	- Format: `hh:mm:ss[.SSS]`.
	//
	// 	- Example: 01:59:59.999.
	//
	// Or
	//
	// 	- Format: `sssss[.SSS]`.
	//
	// 	- Example: 32000.23.
	//
	// example:
	//
	// 01:59:59.999
	Seek *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputClipTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputClipTimeSpan) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) GetSeek() *string {
	return s.Seek
}

func (s *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputClipTimeSpan {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) SetSeek(v string) *QueryJobListResponseBodyJobListJobOutputClipTimeSpan {
	s.Seek = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputClipTimeSpan) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputContainer struct {
	// The container format.
	//
	// 	- Default value: mp4.
	//
	// 	- Video formats include FLV, MP4, HLS (M3U8 + TS), and MPEG-DASH (MPD + fMP4).
	//
	// 	- Audio formats include MP3, MP4, Ogg, FLAC, and M4A.
	//
	// 	- Image formats include GIF and WebP. If the container format is GIF, the video codec must be GIF.
	//
	// 	- If the container format is WebP, the video codec must be WebP.
	//
	// 	- If the container format is FLV, the video codec cannot be H.265.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputContainer) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputContainer) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputContainer) GetFormat() *string {
	return s.Format
}

func (s *QueryJobListResponseBodyJobListJobOutputContainer) SetFormat(v string) *QueryJobListResponseBodyJobListJobOutputContainer {
	s.Format = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputContainer) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputEncryption struct {
	// The encryption ID.
	//
	// example:
	//
	// 31fa3c9ca8134f9cec2b4b0b0f78****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The key that is used to encrypt the video.
	//
	// example:
	//
	// encryptionkey128
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The key encryption method. Valid values: Base64 and KMS.
	//
	// >  For example, if the key is `encryptionkey128`, the key can be encrypted as `Base64("encryptionkey128")` or `KMS(Base64("encryptionkey128")` depending on the encryption method used.
	//
	// example:
	//
	// Base64
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The URL that is used to request the key. The URL is Base64-encoded.
	//
	// example:
	//
	// https://1161758785*****.cn-shanghai.fc.aliyuncs.com/2016-08-15/proxy/HLS-decyptServer/decyptServer/
	KeyUri *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	// The number of unencrypted frames at the beginning of the video. Leaving these frames unencrypted enables video playback to quickly start.
	//
	// example:
	//
	// 3
	SkipCnt *string `json:"SkipCnt,omitempty" xml:"SkipCnt,omitempty"`
	// The encryption type. Only **hls-aes-128*	- may be returned.
	//
	// example:
	//
	// hls-aes-128
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputEncryption) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputEncryption) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetId() *string {
	return s.Id
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetKey() *string {
	return s.Key
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetKeyType() *string {
	return s.KeyType
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetKeyUri() *string {
	return s.KeyUri
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetSkipCnt() *string {
	return s.SkipCnt
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) GetType() *string {
	return s.Type
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetId(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.Id = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetKey(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.Key = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetKeyType(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.KeyType = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetKeyUri(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.KeyUri = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetSkipCnt(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.SkipCnt = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) SetType(v string) *QueryJobListResponseBodyJobListJobOutputEncryption {
	s.Type = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputEncryption) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport struct {
	// The non-standard support configurations for TS files. The value is a JSON object. For more information, see the TS section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	TS *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS `json:"TS,omitempty" xml:"TS,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) GetTS() *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	return s.TS
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) SetTS(v *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport {
	s.TS = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport) Validate() error {
	if s.TS != nil {
		if err := s.TS.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS struct {
	// Indicates whether the output of the MD5 value of the TS file is supported in the M3U8 file. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Md5Support *bool `json:"Md5Support,omitempty" xml:"Md5Support,omitempty"`
	// Indicates whether the output of the size of the TS file is supported in the M3U8 file. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SizeSupport *bool `json:"SizeSupport,omitempty" xml:"SizeSupport,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GetMd5Support() *bool {
	return s.Md5Support
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GetSizeSupport() *bool {
	return s.SizeSupport
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) SetMd5Support(v bool) *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	s.Md5Support = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) SetSizeSupport(v bool) *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	s.SizeSupport = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupportTS) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputMergeList struct {
	Merge []*QueryJobListResponseBodyJobListJobOutputMergeListMerge `json:"Merge,omitempty" xml:"Merge,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputMergeList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMergeList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeList) GetMerge() []*QueryJobListResponseBodyJobListJobOutputMergeListMerge {
	return s.Merge
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeList) SetMerge(v []*QueryJobListResponseBodyJobListJobOutputMergeListMerge) *QueryJobListResponseBodyJobListJobOutputMergeList {
	s.Merge = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeList) Validate() error {
	if s.Merge != nil {
		for _, item := range s.Merge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputMergeListMerge struct {
	// The duration of the clip.
	//
	// 	- Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// 	- Examples: 01:59:59.999 and 32000.23.
	//
	// example:
	//
	// 01:59:59.999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The OSS URL of the clip.
	//
	// 	- Example: `http://example-bucket-.oss-cn-hangzhou.aliyuncs.com/example-object.flv`.
	//
	// 	- The object must be URL-encoded by using the UTF-8 standard. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// example:
	//
	// http://example-bucket.oss-cn-hangzhou.aliyuncs.com/example-object.flv
	MergeURL *string `json:"MergeURL,omitempty" xml:"MergeURL,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the Resource Access Management (RAM) role used for delegated authorization.
	//
	// example:
	//
	// acs:ram::<your uid>:role/<your role name>
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The start point in time of the clip.
	//
	// 	- Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// 	- Examples: 01:59:59.999 and 32000.23.
	//
	// example:
	//
	// 01:59:59.999
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputMergeListMerge) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMergeListMerge) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) GetMergeURL() *string {
	return s.MergeURL
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) GetStart() *string {
	return s.Start
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputMergeListMerge {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) SetMergeURL(v string) *QueryJobListResponseBodyJobListJobOutputMergeListMerge {
	s.MergeURL = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) SetRoleArn(v string) *QueryJobListResponseBodyJobListJobOutputMergeListMerge {
	s.RoleArn = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) SetStart(v string) *QueryJobListResponseBodyJobListJobOutputMergeListMerge {
	s.Start = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMergeListMerge) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo struct {
	// The error code returned if high-speed transcoding is not enabled.
	//
	// example:
	//
	// Boost.NotNeedSpeed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The downgrade policy if high-speed transcoding is not supported.
	//
	// example:
	//
	// NormalSpeed
	DowngradePolicy *string `json:"DowngradePolicy,omitempty" xml:"DowngradePolicy,omitempty"`
	// The duration of the output video.
	//
	// example:
	//
	// 21.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether high-speed transcoding is enabled.
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The error message returned if high-speed transcoding is not enabled.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The actual transcoding speed.
	//
	// example:
	//
	// 6.576886940181647
	RealSpeed *float64 `json:"RealSpeed,omitempty" xml:"RealSpeed,omitempty"`
	// The speed setting.
	//
	// example:
	//
	// 30
	SettingSpeed *int32 `json:"SettingSpeed,omitempty" xml:"SettingSpeed,omitempty"`
	// The amount of time consumed.
	//
	// example:
	//
	// 3.193
	TimeCost *float64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetCode() *string {
	return s.Code
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetDowngradePolicy() *string {
	return s.DowngradePolicy
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetDuration() *float64 {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetEnable() *string {
	return s.Enable
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetMessage() *string {
	return s.Message
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetRealSpeed() *float64 {
	return s.RealSpeed
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetSettingSpeed() *int32 {
	return s.SettingSpeed
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) GetTimeCost() *float64 {
	return s.TimeCost
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetCode(v string) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.Code = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetDowngradePolicy(v string) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.DowngradePolicy = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetDuration(v float64) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetEnable(v string) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.Enable = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetMessage(v string) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.Message = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetRealSpeed(v float64) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.RealSpeed = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetSettingSpeed(v int32) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.SettingSpeed = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) SetTimeCost(v float64) *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo {
	s.TimeCost = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputMuxConfig struct {
	// The transmuxing configurations for GIF.
	Gif *QueryJobListResponseBodyJobListJobOutputMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment configurations. The value is a JSON object.
	Segment *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The transmuxing configurations for WebP.
	Webp *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfig) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) GetGif() *QueryJobListResponseBodyJobListJobOutputMuxConfigGif {
	return s.Gif
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) GetSegment() *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment {
	return s.Segment
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) GetWebp() *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp {
	return s.Webp
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) SetGif(v *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) *QueryJobListResponseBodyJobListJobOutputMuxConfig {
	s.Gif = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) SetSegment(v *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) *QueryJobListResponseBodyJobListJobOutputMuxConfig {
	s.Segment = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) SetWebp(v *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) *QueryJobListResponseBodyJobListJobOutputMuxConfig {
	s.Webp = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfig) Validate() error {
	if s.Gif != nil {
		if err := s.Gif.Validate(); err != nil {
			return err
		}
	}
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	if s.Webp != nil {
		if err := s.Webp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputMuxConfigGif struct {
	// The color dithering algorithm of the palette. Valid values: sierra and bayer.
	//
	// example:
	//
	// bayer
	DitherMode *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	// The duration for which the final frame is paused. Unit: centisecond.
	//
	// example:
	//
	// 0
	FinalDelay *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	// Indicates whether a custom palette is used. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	// The loop count.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigGif) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) SetDitherMode(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) SetFinalDelay(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) SetIsCustomPalette(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) SetLoop(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputMuxConfigSegment struct {
	// The segment length. Unit: seconds.
	//
	// example:
	//
	// 20
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputMuxConfigWebp struct {
	// The loop count.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) SetLoop(v string) *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputOpeningList struct {
	Opening []*QueryJobListResponseBodyJobListJobOutputOpeningListOpening `json:"Opening,omitempty" xml:"Opening,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputOpeningList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOpeningList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningList) GetOpening() []*QueryJobListResponseBodyJobListJobOutputOpeningListOpening {
	return s.Opening
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningList) SetOpening(v []*QueryJobListResponseBodyJobListJobOutputOpeningListOpening) *QueryJobListResponseBodyJobListJobOutputOpeningList {
	s.Opening = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningList) Validate() error {
	if s.Opening != nil {
		for _, item := range s.Opening {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputOpeningListOpening struct {
	// The height of the opening part.
	//
	// 	- Valid values: values in the range of (0,4096), -1, and full.
	//
	// 	- A value of -1 indicates that the original height of the opening part is retained.
	//
	// 	- A value of full indicates that the height of the opening part equals the height of the main part.
	//
	// 	- Default value: **-1**.
	//
	// example:
	//
	// -1
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The amount of time after which the opening part is played.
	//
	// 	- The value starts from 0.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **0**.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The width of the opening part.
	//
	// 	- Valid values: values in the range of (0,4096), -1, and full.
	//
	// 	- A value of -1 indicates that the original width of the opening part is retained.
	//
	// 	- A value of full indicates that the width of the opening part equals the width of the main part.
	//
	// 	- Default value: **-1**.
	//
	// example:
	//
	// -1
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The OSS URL of the opening part.
	//
	// example:
	//
	// http://example.oss-cn-shanghai.aliyuncs.com/t5.mp4
	OpenUrl *string `json:"openUrl,omitempty" xml:"openUrl,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputOpeningListOpening) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOpeningListOpening) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) GetStart() *string {
	return s.Start
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) GetOpenUrl() *string {
	return s.OpenUrl
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputOpeningListOpening {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) SetStart(v string) *QueryJobListResponseBodyJobListJobOutputOpeningListOpening {
	s.Start = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputOpeningListOpening {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) SetOpenUrl(v string) *QueryJobListResponseBodyJobListJobOutputOpeningListOpening {
	s.OpenUrl = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOpeningListOpening) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputOutSubtitleList struct {
	OutSubtitle []*QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle `json:"OutSubtitle,omitempty" xml:"OutSubtitle,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleList) GetOutSubtitle() []*QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	return s.OutSubtitle
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleList) SetOutSubtitle(v []*QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) *QueryJobListResponseBodyJobListJobOutputOutSubtitleList {
	s.OutSubtitle = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleList) Validate() error {
	if s.OutSubtitle != nil {
		for _, item := range s.OutSubtitle {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle struct {
	// The video track. Format: `0:{Stream}:{Stream sequence number}`, which is `0:v:{video_index}`. The value of Stream is v, which indicates a video stream. The sequence number is the index of the video stream in the list and starts from 0.
	//
	// example:
	//
	// 0:v:0
	Map *string `json:"Map,omitempty" xml:"Map,omitempty"`
	// The error message returned if the job failed to be created. This parameter is not returned if the job was created.
	//
	// example:
	//
	// The specified parameter “%s” cannot be null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the output caption.
	OutSubtitleFile *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile `json:"OutSubtitleFile,omitempty" xml:"OutSubtitleFile,omitempty" type:"Struct"`
	// Indicates whether the job was successful. Valid values:
	//
	// 	- **true**: The job was successful.
	//
	// 	- **false**: The job failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetMap() *string {
	return s.Map
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetMessage() *string {
	return s.Message
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetOutSubtitleFile() *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	return s.OutSubtitleFile
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetSuccess() *bool {
	return s.Success
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetMap(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Map = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetMessage(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Message = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetOutSubtitleFile(v *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.OutSubtitleFile = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetSuccess(v bool) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Success = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) Validate() error {
	if s.OutSubtitleFile != nil {
		if err := s.OutSubtitleFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile struct {
	// The name of the OSS bucket in which the output caption is stored.
	//
	// example:
	//
	// exampleBucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the output caption resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the output caption.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The ARN of the RAM role used for delegated authorization.
	//
	// example:
	//
	// acs:ram::<your uid>:role/<your role name>
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetLocation() *string {
	return s.Location
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetObject() *string {
	return s.Object
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetBucket(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Bucket = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetLocation(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Location = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetObject(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Object = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetRoleArn(v string) *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.RoleArn = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputOutputFile struct {
	// The name of the OSS bucket in which the output file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the output file resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the output file.
	//
	// example:
	//
	// example-output.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The ARN of the RAM role used for delegated authorization.
	//
	// example:
	//
	// acs:ram::<your uid>:role/<your role name>
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputOutputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputOutputFile) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) GetObject() *string {
	return s.Object
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) SetBucket(v string) *QueryJobListResponseBodyJobListJobOutputOutputFile {
	s.Bucket = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) SetLocation(v string) *QueryJobListResponseBodyJobListJobOutputOutputFile {
	s.Location = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) SetObject(v string) *QueryJobListResponseBodyJobListJobOutputOutputFile {
	s.Object = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) SetRoleArn(v string) *QueryJobListResponseBodyJobListJobOutputOutputFile {
	s.RoleArn = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputOutputFile) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputProperties struct {
	// The video bitrate.
	//
	// example:
	//
	// 490
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video duration.
	//
	// example:
	//
	// 17
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The video format.
	//
	// example:
	//
	// mp4
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The size of the media file.
	//
	// example:
	//
	// 1057273
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The format information.
	Format *QueryJobListResponseBodyJobListJobOutputPropertiesFormat `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	// The frame rate of the video.
	//
	// example:
	//
	// 30
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The video height.
	//
	// example:
	//
	// 1280
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The non-engine layer keywords.
	SourceLogos *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos `json:"SourceLogos,omitempty" xml:"SourceLogos,omitempty" type:"Struct"`
	// The stream information.
	Streams *QueryJobListResponseBodyJobListJobOutputPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	// The video width.
	//
	// example:
	//
	// 720
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputProperties) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputProperties) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetFileFormat() *string {
	return s.FileFormat
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetFileSize() *string {
	return s.FileSize
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetFormat() *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	return s.Format
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetFps() *string {
	return s.Fps
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetSourceLogos() *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos {
	return s.SourceLogos
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetStreams() *QueryJobListResponseBodyJobListJobOutputPropertiesStreams {
	return s.Streams
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetFileFormat(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.FileFormat = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetFileSize(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.FileSize = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetFormat(v *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Format = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetFps(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Fps = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetSourceLogos(v *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.SourceLogos = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetStreams(v *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Streams = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputProperties {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputProperties) Validate() error {
	if s.Format != nil {
		if err := s.Format.Validate(); err != nil {
			return err
		}
	}
	if s.SourceLogos != nil {
		if err := s.SourceLogos.Validate(); err != nil {
			return err
		}
	}
	if s.Streams != nil {
		if err := s.Streams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesFormat struct {
	// The total bitrate.
	//
	// example:
	//
	// 490.784
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 17.234000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The full name of the container format.
	//
	// example:
	//
	// QuickTime / MOV
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	// The short name of the container format. Valid values: mov, mp4, m4a, 3gp, 3g2, and mj2.
	//
	// example:
	//
	// mov
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The total number of program streams.
	//
	// example:
	//
	// 0
	NumPrograms *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	// The total number of media streams.
	//
	// example:
	//
	// 2
	NumStreams *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	// The size of the media file.
	//
	// example:
	//
	// 1057273
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time.
	//
	// example:
	//
	// -0.064000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesFormat) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetSize() *string {
	return s.Size
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetFormatLongName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.FormatLongName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetFormatName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.FormatName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetNumPrograms(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.NumPrograms = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetNumStreams(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.NumStreams = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetSize(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.Size = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) SetStartTime(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesFormat {
	s.StartTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos struct {
	SourceLogo []*QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo `json:"SourceLogo,omitempty" xml:"SourceLogo,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) GetSourceLogo() []*QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo {
	return s.SourceLogo
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) SetSourceLogo(v []*QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos {
	s.SourceLogo = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos) Validate() error {
	if s.SourceLogo != nil {
		for _, item := range s.SourceLogo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo struct {
	// The keyword.
	//
	// example:
	//
	// example
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) GetSource() *string {
	return s.Source
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) SetSource(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo {
	s.Source = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogosSourceLogo) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreams struct {
	// The audio streams.
	AudioStreamList *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	// The caption streams.
	SubtitleStreamList *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	// The video streams.
	VideoStreamList *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreams) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreams) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) GetAudioStreamList() *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) GetSubtitleStreamList() *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) GetVideoStreamList() *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) SetAudioStreamList(v *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) *QueryJobListResponseBodyJobListJobOutputPropertiesStreams {
	s.AudioStreamList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) SetSubtitleStreamList(v *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) *QueryJobListResponseBodyJobListJobOutputPropertiesStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) SetVideoStreamList(v *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) *QueryJobListResponseBodyJobListJobOutputPropertiesStreams {
	s.VideoStreamList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreams) Validate() error {
	if s.AudioStreamList != nil {
		if err := s.AudioStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleStreamList != nil {
		if err := s.SubtitleStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.VideoStreamList != nil {
		if err := s.VideoStreamList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList struct {
	AudioStream []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) GetAudioStream() []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) SetAudioStream(v []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) Validate() error {
	if s.AudioStream != nil {
		for _, item := range s.AudioStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream struct {
	// The bitrate of the audio stream.
	//
	// example:
	//
	// 64.136
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of the sound channels.
	//
	// example:
	//
	// mono
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 1
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the codec.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the codec.
	//
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the codec.
	//
	// example:
	//
	// mp4
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/32000
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration of the audio stream.
	//
	// example:
	//
	// 17.223562
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the audio stream. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata) and [ISO 639](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
	//
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 50
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 32000
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.064000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base of the audio stream.
	//
	// example:
	//
	// 1/32000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannels(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTag(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetIndex(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetLang(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetNumFrames(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSamplerate(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetStartTime(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetTimebase(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList struct {
	SubtitleStream []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) GetSubtitleStream() []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) SetSubtitleStream(v []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) Validate() error {
	if s.SubtitleStream != nil {
		for _, item := range s.SubtitleStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream struct {
	// The sequence number of the caption stream. The value indicates the position of the caption stream in all caption streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the caption stream. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata) and [ISO 639](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
	//
	// example:
	//
	// eng
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList struct {
	VideoStream []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) GetVideoStream() []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) SetVideoStream(v []*QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) Validate() error {
	if s.VideoStream != nil {
		for _, item := range s.VideoStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream struct {
	// The average frame rate of the video stream.
	//
	// example:
	//
	// 30.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The video bitrate.
	//
	// example:
	//
	// 421.117
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the codec.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the codec.
	//
	// example:
	//
	// 0x31637661
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the codec.
	//
	// example:
	//
	// avc1
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/60
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR) of the video stream.
	//
	// example:
	//
	// 9:16
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration of the video stream.
	//
	// example:
	//
	// 17.233333
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate of the video stream.
	//
	// example:
	//
	// 30.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether the video stream contains bidirectional frames (B-frames).
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height of the video stream in pixels.
	//
	// example:
	//
	// 1280
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value indicates the position of the video stream in all video streams.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the video stream. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html#Metadata) and [ISO 639](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
	//
	// example:
	//
	// eng
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
	//
	// example:
	//
	// 31
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network bandwidth that was consumed.
	NetworkCost *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	// The total number of frames.
	//
	// example:
	//
	// 30
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format of the video stream.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The codec profile.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The sample aspect ratio (SAR) of the video stream.
	//
	// example:
	//
	// 1:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base of the video stream.
	//
	// example:
	//
	// 1/15360
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width of the video stream in pixels.
	//
	// example:
	//
	// 720
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The number of binary bits used by each sample or pixel.
	//
	// example:
	//
	// 8
	BitsPerRawSample *string `json:"bitsPerRawSample,omitempty" xml:"bitsPerRawSample,omitempty"`
	// The primary colors.
	//
	// example:
	//
	// bt709
	ColorPrimaries *string `json:"colorPrimaries,omitempty" xml:"colorPrimaries,omitempty"`
	// The color transfer configuration.
	//
	// example:
	//
	// bt709
	ColorTransfer *string `json:"colorTransfer,omitempty" xml:"colorTransfer,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNetworkCost() *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetBitsPerRawSample() *string {
	return s.BitsPerRawSample
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetColorPrimaries() *string {
	return s.ColorPrimaries
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetColorTransfer() *string {
	return s.ColorTransfer
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecName(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTag(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDar(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDuration(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetFps(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetIndex(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLang(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLevel(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNetworkCost(v *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNumFrames(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetPixFmt(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetProfile(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetSar(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetStartTime(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetTimebase(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetBitsPerRawSample(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.BitsPerRawSample = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetColorPrimaries(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.ColorPrimaries = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetColorTransfer(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.ColorTransfer = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	// The average bitrate.
	//
	// example:
	//
	// 300
	AvgBitrate *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	// The maximum bandwidth that was consumed.
	//
	// example:
	//
	// 10
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	// The amount of time consumed to preload the video stream.
	//
	// example:
	//
	// 8
	PreloadTime *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfig struct {
	// The external captions.
	ExtSubtitleList *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList `json:"ExtSubtitleList,omitempty" xml:"ExtSubtitleList,omitempty" type:"Struct"`
	// The captions.
	SubtitleList *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfig) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) GetExtSubtitleList() *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList {
	return s.ExtSubtitleList
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) GetSubtitleList() *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList {
	return s.SubtitleList
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) SetExtSubtitleList(v *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) *QueryJobListResponseBodyJobListJobOutputSubtitleConfig {
	s.ExtSubtitleList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) SetSubtitleList(v *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) *QueryJobListResponseBodyJobListJobOutputSubtitleConfig {
	s.SubtitleList = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfig) Validate() error {
	if s.ExtSubtitleList != nil {
		if err := s.ExtSubtitleList.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleList != nil {
		if err := s.SubtitleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList struct {
	ExtSubtitle []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle `json:"ExtSubtitle,omitempty" xml:"ExtSubtitle,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) GetExtSubtitle() []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	return s.ExtSubtitle
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) SetExtSubtitle(v []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList {
	s.ExtSubtitle = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) Validate() error {
	if s.ExtSubtitle != nil {
		for _, item := range s.ExtSubtitle {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle struct {
	// The character set used by the external caption.
	//
	// 	- Valid values: UTF-8, GBK, BIG5, and auto.
	//
	// 	- Default value: **auto**.
	//
	// >  If the value of CharEnc is auto, the detected character set may not be the actual character set. We recommend that you set this parameter to another value.
	//
	// example:
	//
	// auto
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The font of the hardcoded captions converted from external captions. Default value: SimSum. For more information, see [Fonts](https://help.aliyun.com/document_detail/59950.html).
	//
	// example:
	//
	// "WenQuanYi Zen Hei", "Yuanti SC Regular", "SimSun"
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The input caption file.
	//
	// 	- SRT and ASS files are supported. For more information, see the Input section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// 	- Example: `{"Bucket":"example-bucket","Location":"oss-cn-hangzhou","Object":"example.srt"}`.
	Input *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetCharEnc() *string {
	return s.CharEnc
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetFontName() *string {
	return s.FontName
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetInput() *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	return s.Input
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetCharEnc(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.CharEnc = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetFontName(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.FontName = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetInput(v *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.Input = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput struct {
	// The name of the OSS bucket in which the input caption file is stored.
	//
	// example:
	//
	// example-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input caption file resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input caption file.
	//
	// example:
	//
	// example-output.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetBucket() *string {
	return s.Bucket
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetLocation() *string {
	return s.Location
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetObject() *string {
	return s.Object
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetBucket(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Bucket = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetLocation(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Location = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetObject(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Object = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList struct {
	Subtitle []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) GetSubtitle() []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle {
	return s.Subtitle
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) SetSubtitle(v []*QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList {
	s.Subtitle = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList) Validate() error {
	if s.Subtitle != nil {
		for _, item := range s.Subtitle {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle struct {
	// The audio track. Format: `0:{Stream}:{Stream sequence number}`, which is `0:a:{audio_index}`. The value of Stream is a, which indicates an audio stream. The sequence number is the index of the audio stream in the list and starts from 0.
	//
	// example:
	//
	// 0:a:0
	Map *string `json:"Map,omitempty" xml:"Map,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) GetMap() *string {
	return s.Map
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) SetMap(v string) *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle {
	s.Map = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputSuperReso struct {
	// Indicates whether parameters related to the sampling rate are obtained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsHalfSample *string `json:"IsHalfSample,omitempty" xml:"IsHalfSample,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputSuperReso) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputSuperReso) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputSuperReso) GetIsHalfSample() *string {
	return s.IsHalfSample
}

func (s *QueryJobListResponseBodyJobListJobOutputSuperReso) SetIsHalfSample(v string) *QueryJobListResponseBodyJobListJobOutputSuperReso {
	s.IsHalfSample = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputSuperReso) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputTailSlateList struct {
	TailSlate []*QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate `json:"TailSlate,omitempty" xml:"TailSlate,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputTailSlateList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputTailSlateList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateList) GetTailSlate() []*QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	return s.TailSlate
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateList) SetTailSlate(v []*QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) *QueryJobListResponseBodyJobListJobOutputTailSlateList {
	s.TailSlate = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateList) Validate() error {
	if s.TailSlate != nil {
		for _, item := range s.TailSlate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate struct {
	// The color of the bars that are added to the ending part if the size of the ending part is smaller than that of the main part. Default value: White. For more information, see [Parameter details](https://help.aliyun.com/document_detail/29253.html).
	//
	// example:
	//
	// White
	BgColor *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	// The duration of the transition between the main part and the ending part. A fade transition is used: The last frame of the main part fades out, and the first frame of the ending part fades in. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 0
	BlendDuration *string `json:"BlendDuration,omitempty" xml:"BlendDuration,omitempty"`
	// The height of the ending part.
	//
	// 	- Valid values: values in the range of (0,4096), -1, and full.
	//
	// 	- A value of -1 indicates that the original height of the ending part is retained.
	//
	// 	- A value of full indicates that the height of the ending part equals the height of the main part.
	//
	// 	- Default value: -1.
	//
	// example:
	//
	// -1
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Indicates whether the audio content of the ending part is merged. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsMergeAudio *bool `json:"IsMergeAudio,omitempty" xml:"IsMergeAudio,omitempty"`
	// The time when the ending part is played.
	//
	// example:
	//
	// 00000.00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The OSS URL of the ending part.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com/opening_01.flv
	TailUrl *string `json:"TailUrl,omitempty" xml:"TailUrl,omitempty"`
	// The width of the ending part. Valid values: values in the range of (0,4096), -1, and full.
	//
	// 	- A value of -1 indicates that the original width of the ending part is retained.
	//
	// 	- A value of full indicates that the width of the ending part equals the width of the main part.
	//
	// 	- Default value: -1.
	//
	// example:
	//
	// -1
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetBgColor() *string {
	return s.BgColor
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetBlendDuration() *string {
	return s.BlendDuration
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetIsMergeAudio() *bool {
	return s.IsMergeAudio
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetStart() *string {
	return s.Start
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetTailUrl() *string {
	return s.TailUrl
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetBgColor(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.BgColor = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetBlendDuration(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.BlendDuration = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetIsMergeAudio(v bool) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.IsMergeAudio = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetStart(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Start = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetTailUrl(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.TailUrl = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTailSlateListTailSlate) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputTransConfig struct {
	// The method of resolution adjustment. Default value: **none**. Valid values: rescale, crop, pad, and none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Indicates whether the audio bitrate is checked. If the bitrate of the output audio is higher than that of the input audio, the input bitrate is retained and the specified audio bitrate does not take effect. This parameter has a lower priority than IsCheckAudioBitrateFail. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value:
	//
	//     	- If this parameter is empty and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	//     	- If this parameter is empty and the codec of the output audio is the same as that of the input audio, the default value is true.
	//
	// example:
	//
	// false
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Indicates whether the audio bitrate is checked. If the bitrate of the output audio is higher than that of the input audio, the input audio is not transcoded and a transcoding failure is returned. This parameter has a higher priority than IsCheckAudioBitrate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Indicates whether the resolution is checked. If the output resolution is higher than the input resolution based on the width or height, the input resolution is retained. Valid values:
	//
	// 	- **true**:
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Indicates whether the resolution is checked. If the output resolution is higher than the input resolution based on the width or height, a transcoding failure is returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Indicates whether the video bitrate is checked. If the bitrate of the output video is higher than that of the input video, the input bitrate is retained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Indicates whether the video bitrate is checked. If the bitrate of the output video is higher than that of the input video, the input video is not transcoded and a transcoding failure is returned. This parameter has a higher priority than IsCheckVideoBitrate. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// false
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The transcoding mode.
	//
	// 	- Valid values: onepass, twopass, and CBR.
	//
	// 	- Default value: **onepass**.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputTransConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputTransConfig) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetAdjDarMethod(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckAudioBitrate(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckAudioBitrateFail(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckReso(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckResoFail(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckVideoBitrate(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetIsCheckVideoBitrateFail(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) SetTransMode(v string) *QueryJobListResponseBodyJobListJobOutputTransConfig {
	s.TransMode = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputTransConfig) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputVideo struct {
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 500
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The average bitrate range of the video.
	BitrateBnd *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	// The buffer size.
	//
	// 	- Unit: KB.
	//
	// 	- Default value: **6000**.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The video codec.
	//
	// 	- Valid values: H.264 and H.265.
	//
	// 	- Default value: H.264.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Default value when the value of Codec is H.264: **23**, default value when the value of Codec is H.265: **26**.
	//
	// 	- If the value of this parameter is returned, the value of Bitrate becomes invalid.
	//
	// example:
	//
	// 26
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The method of video cropping. Valid values:
	//
	// 	- **border**: automatically detects and removes borders.
	//
	// 	- A value in the width:height:left:top format: The video image is cropped based on custom settings.
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The strength of the independent noise reduction algorithm.
	//
	// example:
	//
	// 5
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	// The frame rate of the video.
	//
	// 	- Unit: frames per second.
	//
	// 	- The value is 60 if the frame rate of the input file exceeds 60.
	//
	// 	- Default value: the frame rate of the input video.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum interval between keyframes or the maximum number of frames in a frame group. Unit: seconds.
	//
	// 	- Default value: **250**.
	//
	// 	- If the maximum number of frames is returned, the value does not contain a unit.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: the height of the input video.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The maximum frame rate.
	//
	// example:
	//
	// 60
	MaxFps *string `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	// The maximum bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 3000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The black bars that are added to the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Format: width:height:left:top.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The pixel format of the video. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The preset video algorithm. Default value: **medium**. Valid values:
	//
	// 	- **veryfast**
	//
	// 	- **fast**
	//
	// 	- **medium**
	//
	// 	- **slow**
	//
	// 	- **slower**
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The codec profile of the video. Valid values: baseline, main, and high.
	//
	// >  If multiple definitions are involved, we recommend that you use baseline for the lowest definition to ensure normal playback on low-end devices, and use main or high for other definitions.
	//
	// 	- **baseline**: applicable to mobile devices.
	//
	// 	- **main**: applicable to standard-definition devices.
	//
	// 	- **high**: applicable to high-definition devices.
	//
	// 	- Default value: **high**.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of quality control on the video.
	//
	// example:
	//
	// 15
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The resource priority.
	//
	// example:
	//
	// 1
	ResoPriority *string `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- If this parameter is left **empty**, the scan mode of the input video is used.
	//
	// 	- **auto**: automatic deinterlacing.
	//
	// 	- **progressive**: progressive scan.
	//
	// 	- **interlaced**: interlaced scan.
	//
	// 	- **By default**, this parameter is left empty.
	//
	// **Best practice**: The interlaced scan mode saves data traffic than the progressive scan mode but provides poor image quality. Therefore, the progressive scan mode is commonly used in mainstream video production.
	//
	// 	- If **progressive*	- or **interlaced*	- is used when the scan mode of the input video is neither of them, the transcoding job fails.
	//
	// 	- We recommend that you use **the scan mode of the input video*	- or **automatic deinterlacing*	- for higher compatibility.
	//
	// example:
	//
	// interlaced
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: the width of the input video.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputVideo) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputVideo) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetBitrateBnd() *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetCodec() *string {
	return s.Codec
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetCrf() *string {
	return s.Crf
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetCrop() *string {
	return s.Crop
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetFps() *string {
	return s.Fps
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetGop() *string {
	return s.Gop
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetPad() *string {
	return s.Pad
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetPreset() *string {
	return s.Preset
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetProfile() *string {
	return s.Profile
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetQscale() *string {
	return s.Qscale
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetBitrate(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Bitrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetBitrateBnd(v *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.BitrateBnd = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetBufsize(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Bufsize = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetCodec(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Codec = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetCrf(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Crf = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetCrop(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Crop = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetDegrain(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Degrain = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetFps(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Fps = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetGop(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Gop = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetMaxFps(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.MaxFps = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetMaxrate(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Maxrate = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetPad(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Pad = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetPixFmt(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.PixFmt = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetPreset(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Preset = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetProfile(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Profile = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetQscale(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Qscale = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetResoPriority(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.ResoPriority = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetScanMode(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.ScanMode = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputVideo {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd struct {
	// The maximum bitrate.
	//
	// example:
	//
	// 1000
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum bitrate.
	//
	// example:
	//
	// 300
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) SetMax(v string) *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) SetMin(v string) *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyJobListJobOutputWaterMarkList struct {
	WaterMark []*QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark `json:"WaterMark,omitempty" xml:"WaterMark,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkList) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkList) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkList) GetWaterMark() []*QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	return s.WaterMark
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkList) SetWaterMark(v []*QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) *QueryJobListResponseBodyJobListJobOutputWaterMarkList {
	s.WaterMark = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkList) Validate() error {
	if s.WaterMark != nil {
		for _, item := range s.WaterMark {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark struct {
	// The horizontal offset of the watermark image relative to the output video. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. Default value: 0. The value can be an integer or a decimal number.
	//
	// 	- An integer indicates the pixel value of the horizontal offset.
	//
	//     	- Valid values: **[8,4096]**.
	//
	//     	- Unit: pixel.
	//
	// 	- A decimal number indicates the ratio of the horizontal offset to the width in the output video resolution.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 100
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the watermark image relative to the output video. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. The value can be an integer or a decimal number.
	//
	// 	- An integer indicates the pixel value of the vertical offset.
	//
	//     	- Valid values: **[8,4096]**.
	//
	//     	- Unit: pixel.
	//
	// 	- A decimal number indicates the ratio of the vertical offset to the height in the output video resolution.
	//
	//     	- Valid values: **(0,1)**.
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 100
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the watermark image. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. The value can be an integer or a decimal number.
	//
	// 	- An integer indicates the pixel value of the watermark height.
	//
	//     	- Valid values: **[8,4096]**.
	//
	//     	- Unit: pixel.
	//
	// 	- A decimal number indicates the ratio of the watermark height to the height in the output video resolution.
	//
	//     	- Valid values: **(0,1)**.
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 50
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The watermark input file. PNG images and MOV files are supported.
	InputFile *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The position of the watermark. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. Valid values:
	//
	// 	- TopRight
	//
	// 	- TopLeft
	//
	// 	- BottomRight
	//
	// 	- BottomLeft
	//
	// example:
	//
	// TopRight
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The type of the watermark. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. For more information, see [Parameter details](https://help.aliyun.com/document_detail/29253.html). Valid values:
	//
	// 	- Image
	//
	// 	- Text
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	WaterMarkTemplateId *string `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
	// The width of the watermark image. If this parameter is specified in the request, the corresponding parameter in the watermark template is overwritten. The value can be an integer or a decimal number.
	//
	// 	- An integer indicates the pixel value of the watermark width.
	//
	//     	- Valid values: **[8,4096]**.
	//
	//     	- Unit: pixel.
	//
	// 	- A decimal number indicates the ratio of the watermark width to the width in the output video resolution.
	//
	//     	- Valid values: **(0,1)**.
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 50
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetDx() *string {
	return s.Dx
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetDy() *string {
	return s.Dy
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetHeight() *string {
	return s.Height
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetInputFile() *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	return s.InputFile
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetReferPos() *string {
	return s.ReferPos
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetType() *string {
	return s.Type
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) GetWidth() *string {
	return s.Width
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetDx(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Dx = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetDy(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Dy = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetHeight(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Height = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetInputFile(v *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.InputFile = v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetReferPos(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.ReferPos = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetType(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Type = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetWaterMarkTemplateId(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) SetWidth(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Width = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMark) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input file resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the Object Storage Service (OSS) object that is used as the input file.
	//
	// example:
	//
	// example-logo-****.png
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetObject() *string {
	return s.Object
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetBucket(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Bucket = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetLocation(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Location = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetObject(v string) *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Object = &v
	return s
}

func (s *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) Validate() error {
	return dara.Validate(s)
}

type QueryJobListResponseBodyNonExistJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryJobListResponseBodyNonExistJobIds) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponseBodyNonExistJobIds) GoString() string {
	return s.String()
}

func (s *QueryJobListResponseBodyNonExistJobIds) GetString_() []*string {
	return s.String_
}

func (s *QueryJobListResponseBodyNonExistJobIds) SetString_(v []*string) *QueryJobListResponseBodyNonExistJobIds {
	s.String_ = v
	return s
}

func (s *QueryJobListResponseBodyNonExistJobIds) Validate() error {
	return dara.Validate(s)
}
