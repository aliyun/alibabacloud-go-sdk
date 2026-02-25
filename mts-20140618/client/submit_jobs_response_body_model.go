// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobResultList(v *SubmitJobsResponseBodyJobResultList) *SubmitJobsResponseBody
	GetJobResultList() *SubmitJobsResponseBodyJobResultList
	SetRequestId(v string) *SubmitJobsResponseBody
	GetRequestId() *string
}

type SubmitJobsResponseBody struct {
	JobResultList *SubmitJobsResponseBodyJobResultList `json:"JobResultList,omitempty" xml:"JobResultList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A45S71F6-D73936451234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBody) GetJobResultList() *SubmitJobsResponseBodyJobResultList {
	return s.JobResultList
}

func (s *SubmitJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitJobsResponseBody) SetJobResultList(v *SubmitJobsResponseBodyJobResultList) *SubmitJobsResponseBody {
	s.JobResultList = v
	return s
}

func (s *SubmitJobsResponseBody) SetRequestId(v string) *SubmitJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitJobsResponseBody) Validate() error {
	if s.JobResultList != nil {
		if err := s.JobResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultList struct {
	JobResult []*SubmitJobsResponseBodyJobResultListJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultList) GetJobResult() []*SubmitJobsResponseBodyJobResultListJobResult {
	return s.JobResult
}

func (s *SubmitJobsResponseBodyJobResultList) SetJobResult(v []*SubmitJobsResponseBodyJobResultListJobResult) *SubmitJobsResponseBodyJobResultList {
	s.JobResult = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultList) Validate() error {
	if s.JobResult != nil {
		for _, item := range s.JobResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResult struct {
	Code    *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Job     *SubmitJobsResponseBodyJobResultListJobResultJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	Message *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResult) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) GetCode() *string {
	return s.Code
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) GetJob() *SubmitJobsResponseBodyJobResultListJobResultJob {
	return s.Job
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) GetMessage() *string {
	return s.Message
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) SetCode(v string) *SubmitJobsResponseBodyJobResultListJobResult {
	s.Code = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) SetJob(v *SubmitJobsResponseBodyJobResultListJobResultJob) *SubmitJobsResponseBodyJobResultListJobResult {
	s.Job = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) SetMessage(v string) *SubmitJobsResponseBodyJobResultListJobResult {
	s.Message = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) SetSuccess(v bool) *SubmitJobsResponseBodyJobResultListJobResult {
	s.Success = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResult) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJob struct {
	Code             *string                                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime     *string                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FinishTime       *string                                                          `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input            *SubmitJobsResponseBodyJobResultListJobResultJobInput            `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId            *string                                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MNSMessageResult *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Output           *SubmitJobsResponseBodyJobResultListJobResultJobOutput           `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Percent          *int64                                                           `json:"Percent,omitempty" xml:"Percent,omitempty"`
	PipelineId       *string                                                          `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	State            *string                                                          `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJob) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetCode() *string {
	return s.Code
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetInput() *SubmitJobsResponseBodyJobResultListJobResultJobInput {
	return s.Input
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetMNSMessageResult() *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetOutput() *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	return s.Output
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetPercent() *int64 {
	return s.Percent
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) GetState() *string {
	return s.State
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetCode(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.Code = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetCreationTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetFinishTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.FinishTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetInput(v *SubmitJobsResponseBodyJobResultListJobResultJobInput) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.Input = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetJobId(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.JobId = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetMNSMessageResult(v *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.MNSMessageResult = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetMessage(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.Message = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetOutput(v *SubmitJobsResponseBodyJobResultListJobResultJobOutput) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.Output = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetPercent(v int64) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.Percent = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetPipelineId(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.PipelineId = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) SetState(v string) *SubmitJobsResponseBodyJobResultListJobResultJob {
	s.State = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJob) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobInput struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobInput) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobInput {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobInput {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobInput {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobInput) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) SetErrorCode(v string) *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) SetErrorMessage(v string) *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) SetMessageId(v string) *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutput struct {
	AmixList               *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList               `json:"AmixList,omitempty" xml:"AmixList,omitempty" type:"Struct"`
	Audio                  *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio                  `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	AudioStreamMap         *string                                                                      `json:"AudioStreamMap,omitempty" xml:"AudioStreamMap,omitempty"`
	Clip                   *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip                   `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	Container              *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer              `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	DeWatermark            *string                                                                      `json:"DeWatermark,omitempty" xml:"DeWatermark,omitempty"`
	DigiWaterMark          *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark          `json:"DigiWaterMark,omitempty" xml:"DigiWaterMark,omitempty" type:"Struct"`
	Encryption             *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption             `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	M3U8NonStandardSupport *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport `json:"M3U8NonStandardSupport,omitempty" xml:"M3U8NonStandardSupport,omitempty" type:"Struct"`
	MergeConfigUrl         *string                                                                      `json:"MergeConfigUrl,omitempty" xml:"MergeConfigUrl,omitempty"`
	MergeList              *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList              `json:"MergeList,omitempty" xml:"MergeList,omitempty" type:"Struct"`
	MuxConfig              *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig              `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	OpeningList            *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList            `json:"OpeningList,omitempty" xml:"OpeningList,omitempty" type:"Struct"`
	OutSubtitleList        *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList        `json:"OutSubtitleList,omitempty" xml:"OutSubtitleList,omitempty" type:"Struct"`
	OutputFile             *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile             `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	Priority               *string                                                                      `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Properties             *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties             `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Rotate                 *string                                                                      `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SubtitleConfig         *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig         `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Struct"`
	SuperReso              *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso              `json:"SuperReso,omitempty" xml:"SuperReso,omitempty" type:"Struct"`
	TailSlateList          *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList          `json:"TailSlateList,omitempty" xml:"TailSlateList,omitempty" type:"Struct"`
	TemplateId             *string                                                                      `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransConfig            *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig            `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	UserData               *string                                                                      `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Video                  *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo                  `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
	VideoStreamMap         *string                                                                      `json:"VideoStreamMap,omitempty" xml:"VideoStreamMap,omitempty"`
	WaterMarkConfigUrl     *string                                                                      `json:"WaterMarkConfigUrl,omitempty" xml:"WaterMarkConfigUrl,omitempty"`
	WaterMarkList          *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList          `json:"WaterMarkList,omitempty" xml:"WaterMarkList,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutput) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetAmixList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList {
	return s.AmixList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetAudio() *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	return s.Audio
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetAudioStreamMap() *string {
	return s.AudioStreamMap
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetClip() *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip {
	return s.Clip
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetContainer() *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer {
	return s.Container
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetDeWatermark() *string {
	return s.DeWatermark
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetDigiWaterMark() *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark {
	return s.DigiWaterMark
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetEncryption() *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	return s.Encryption
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetM3U8NonStandardSupport() *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport {
	return s.M3U8NonStandardSupport
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetMergeConfigUrl() *string {
	return s.MergeConfigUrl
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetMergeList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList {
	return s.MergeList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetMuxConfig() *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig {
	return s.MuxConfig
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetOpeningList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList {
	return s.OpeningList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetOutSubtitleList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList {
	return s.OutSubtitleList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetOutputFile() *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile {
	return s.OutputFile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetPriority() *string {
	return s.Priority
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetProperties() *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	return s.Properties
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetRotate() *string {
	return s.Rotate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetSubtitleConfig() *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig {
	return s.SubtitleConfig
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetSuperReso() *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso {
	return s.SuperReso
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetTailSlateList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList {
	return s.TailSlateList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetTransConfig() *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	return s.TransConfig
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetUserData() *string {
	return s.UserData
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetVideo() *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	return s.Video
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetVideoStreamMap() *string {
	return s.VideoStreamMap
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetWaterMarkConfigUrl() *string {
	return s.WaterMarkConfigUrl
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) GetWaterMarkList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList {
	return s.WaterMarkList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetAmixList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.AmixList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetAudio(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Audio = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetAudioStreamMap(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.AudioStreamMap = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetClip(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Clip = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetContainer(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Container = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetDeWatermark(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.DeWatermark = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetDigiWaterMark(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.DigiWaterMark = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetEncryption(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Encryption = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetM3U8NonStandardSupport(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.M3U8NonStandardSupport = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetMergeConfigUrl(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.MergeConfigUrl = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetMergeList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.MergeList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetMuxConfig(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.MuxConfig = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetOpeningList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.OpeningList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetOutSubtitleList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.OutSubtitleList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetOutputFile(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.OutputFile = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetPriority(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Priority = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetProperties(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Properties = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetRotate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Rotate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetSubtitleConfig(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.SubtitleConfig = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetSuperReso(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.SuperReso = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetTailSlateList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.TailSlateList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetTemplateId(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.TemplateId = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetTransConfig(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.TransConfig = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetUserData(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.UserData = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetVideo(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.Video = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetVideoStreamMap(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.VideoStreamMap = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetWaterMarkConfigUrl(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.WaterMarkConfigUrl = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) SetWaterMarkList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) *SubmitJobsResponseBodyJobResultListJobResultJobOutput {
	s.WaterMarkList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutput) Validate() error {
	if s.AmixList != nil {
		if err := s.AmixList.Validate(); err != nil {
			return err
		}
	}
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
	if s.DigiWaterMark != nil {
		if err := s.DigiWaterMark.Validate(); err != nil {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList struct {
	Amix []*SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix `json:"Amix,omitempty" xml:"Amix,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) GetAmix() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	return s.Amix
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) SetAmix(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList {
	s.Amix = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixList) Validate() error {
	if s.Amix != nil {
		for _, item := range s.Amix {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix struct {
	AmixURL    *string `json:"AmixURL,omitempty" xml:"AmixURL,omitempty"`
	Duration   *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Map        *string `json:"Map,omitempty" xml:"Map,omitempty"`
	MixDurMode *string `json:"MixDurMode,omitempty" xml:"MixDurMode,omitempty"`
	Start      *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GetAmixURL() *string {
	return s.AmixURL
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GetMap() *string {
	return s.Map
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GetMixDurMode() *string {
	return s.MixDurMode
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) GetStart() *string {
	return s.Start
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) SetAmixURL(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	s.AmixURL = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) SetMap(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	s.Map = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) SetMixDurMode(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	s.MixDurMode = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) SetStart(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix {
	s.Start = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAmixListAmix) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio struct {
	Bitrate    *string                                                           `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *string                                                           `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string                                                           `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string                                                           `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale     *string                                                           `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Samplerate *string                                                           `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	Volume     *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetChannels() *string {
	return s.Channels
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetCodec() *string {
	return s.Codec
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetProfile() *string {
	return s.Profile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetQscale() *string {
	return s.Qscale
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) GetVolume() *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume {
	return s.Volume
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetChannels(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Channels = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetCodec(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Codec = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetProfile(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Profile = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetQscale(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Qscale = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetSamplerate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Samplerate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) SetVolume(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio {
	s.Volume = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume struct {
	Level  *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) SetLevel(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume {
	s.Level = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) SetMethod(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume {
	s.Method = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputAudioVolume) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputClip struct {
	TimeSpan *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) GetTimeSpan() *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan {
	return s.TimeSpan
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) SetTimeSpan(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip {
	s.TimeSpan = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClip) Validate() error {
	if s.TimeSpan != nil {
		if err := s.TimeSpan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Seek     *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) GetSeek() *string {
	return s.Seek
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) SetSeek(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan {
	s.Seek = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputClipTimeSpan) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) GetFormat() *string {
	return s.Format
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) SetFormat(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer {
	s.Format = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputContainer) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark struct {
	Alpha     *string                                                                      `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	InputFile *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	Type      *string                                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) GetAlpha() *string {
	return s.Alpha
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) GetInputFile() *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile {
	return s.InputFile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) GetType() *string {
	return s.Type
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) SetAlpha(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark {
	s.Alpha = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) SetInputFile(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark {
	s.InputFile = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) SetType(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark {
	s.Type = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMark) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputDigiWaterMarkInputFile) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	KeyUri  *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	SkipCnt *string `json:"SkipCnt,omitempty" xml:"SkipCnt,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetId() *string {
	return s.Id
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetKey() *string {
	return s.Key
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetKeyType() *string {
	return s.KeyType
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetKeyUri() *string {
	return s.KeyUri
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetSkipCnt() *string {
	return s.SkipCnt
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) GetType() *string {
	return s.Type
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetId(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.Id = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetKey(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.Key = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetKeyType(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.KeyType = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetKeyUri(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.KeyUri = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetSkipCnt(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.SkipCnt = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) SetType(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption {
	s.Type = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputEncryption) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport struct {
	TS *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS `json:"TS,omitempty" xml:"TS,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) GetTS() *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS {
	return s.TS
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) SetTS(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport {
	s.TS = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupport) Validate() error {
	if s.TS != nil {
		if err := s.TS.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS struct {
	Md5Support  *bool `json:"Md5Support,omitempty" xml:"Md5Support,omitempty"`
	SizeSupport *bool `json:"SizeSupport,omitempty" xml:"SizeSupport,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) GetMd5Support() *bool {
	return s.Md5Support
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) GetSizeSupport() *bool {
	return s.SizeSupport
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) SetMd5Support(v bool) *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS {
	s.Md5Support = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) SetSizeSupport(v bool) *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS {
	s.SizeSupport = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputM3U8NonStandardSupportTS) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList struct {
	Merge []*SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge `json:"Merge,omitempty" xml:"Merge,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) GetMerge() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge {
	return s.Merge
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) SetMerge(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList {
	s.Merge = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MergeURL *string `json:"MergeURL,omitempty" xml:"MergeURL,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) GetMergeURL() *string {
	return s.MergeURL
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) GetStart() *string {
	return s.Start
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) SetMergeURL(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge {
	s.MergeURL = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) SetRoleArn(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge {
	s.RoleArn = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) SetStart(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge {
	s.Start = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMergeListMerge) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig struct {
	Gif     *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif     `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	Segment *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	Webp    *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp    `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) GetGif() *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif {
	return s.Gif
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) GetSegment() *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment {
	return s.Segment
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) GetWebp() *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp {
	return s.Webp
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) SetGif(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig {
	s.Gif = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) SetSegment(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig {
	s.Segment = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) SetWebp(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig {
	s.Webp = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfig) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif struct {
	DitherMode      *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	FinalDelay      *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	Loop            *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) SetDitherMode(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) SetFinalDelay(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) SetIsCustomPalette(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) SetLoop(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp struct {
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) SetLoop(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList struct {
	Opening []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening `json:"Opening,omitempty" xml:"Opening,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) GetOpening() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening {
	return s.Opening
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) SetOpening(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList {
	s.Opening = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening struct {
	Height  *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Start   *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Width   *string `json:"Width,omitempty" xml:"Width,omitempty"`
	OpenUrl *string `json:"openUrl,omitempty" xml:"openUrl,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) GetStart() *string {
	return s.Start
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) GetOpenUrl() *string {
	return s.OpenUrl
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) SetStart(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening {
	s.Start = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) SetOpenUrl(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening {
	s.OpenUrl = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOpeningListOpening) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList struct {
	OutSubtitle []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle `json:"OutSubtitle,omitempty" xml:"OutSubtitle,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) GetOutSubtitle() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle {
	return s.OutSubtitle
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) SetOutSubtitle(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList {
	s.OutSubtitle = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle struct {
	Map             *string                                                                                         `json:"Map,omitempty" xml:"Map,omitempty"`
	Message         *string                                                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	OutSubtitleFile *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile `json:"OutSubtitleFile,omitempty" xml:"OutSubtitleFile,omitempty" type:"Struct"`
	Success         *bool                                                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) GetMap() *string {
	return s.Map
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) GetMessage() *string {
	return s.Message
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) GetOutSubtitleFile() *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	return s.OutSubtitleFile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) SetMap(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle {
	s.Map = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) SetMessage(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle {
	s.Message = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) SetOutSubtitleFile(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle {
	s.OutSubtitleFile = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) SetSuccess(v bool) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle {
	s.Success = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitle) Validate() error {
	if s.OutSubtitleFile != nil {
		if err := s.OutSubtitleFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetRoleArn(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.RoleArn = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) SetRoleArn(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile {
	s.RoleArn = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputOutputFile) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties struct {
	Bitrate    *string                                                                 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration   *string                                                                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileFormat *string                                                                 `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	FileSize   *string                                                                 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Format     *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat  `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	Fps        *string                                                                 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height     *string                                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
	Streams    *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	Width      *string                                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetFileFormat() *string {
	return s.FileFormat
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetFileSize() *string {
	return s.FileSize
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetFormat() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	return s.Format
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetFps() *string {
	return s.Fps
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetStreams() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams {
	return s.Streams
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetFileFormat(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.FileFormat = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetFileSize(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.FileSize = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetFormat(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Format = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetFps(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Fps = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetStreams(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Streams = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputProperties) Validate() error {
	if s.Format != nil {
		if err := s.Format.Validate(); err != nil {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat struct {
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	NumPrograms    *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	NumStreams     *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	Size           *string `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetSize() *string {
	return s.Size
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetFormatLongName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.FormatLongName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetFormatName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.FormatName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetNumPrograms(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.NumPrograms = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetNumStreams(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.NumStreams = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetSize(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.Size = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) SetStartTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat {
	s.StartTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams struct {
	AudioStreamList    *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList    `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	SubtitleStreamList *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	VideoStreamList    *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList    `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) GetAudioStreamList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) GetSubtitleStreamList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) GetVideoStreamList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) SetAudioStreamList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams {
	s.AudioStreamList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) SetSubtitleStreamList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) SetVideoStreamList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams {
	s.VideoStreamList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreams) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList struct {
	AudioStream []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) GetAudioStream() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) SetAudioStream(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream struct {
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout  *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels       *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Index          *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NumFrames      *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	SampleFmt      *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	Samplerate     *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannels(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTag(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetIndex(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetLang(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetNumFrames(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSamplerate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetStartTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) SetTimebase(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList struct {
	SubtitleStream []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) GetSubtitleStream() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) SetSubtitleStream(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream struct {
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList struct {
	VideoStream []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) GetVideoStream() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) SetVideoStream(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream struct {
	AvgFPS         *string                                                                                                      `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	Bitrate        *string                                                                                                      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName  *string                                                                                                      `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string                                                                                                      `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string                                                                                                      `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string                                                                                                      `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string                                                                                                      `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Dar            *string                                                                                                      `json:"Dar,omitempty" xml:"Dar,omitempty"`
	Duration       *string                                                                                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps            *string                                                                                                      `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HasBFrames     *string                                                                                                      `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height         *string                                                                                                      `json:"Height,omitempty" xml:"Height,omitempty"`
	Index          *string                                                                                                      `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string                                                                                                      `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Level          *string                                                                                                      `json:"Level,omitempty" xml:"Level,omitempty"`
	NetworkCost    *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	NumFrames      *string                                                                                                      `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	PixFmt         *string                                                                                                      `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Profile        *string                                                                                                      `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Sar            *string                                                                                                      `json:"Sar,omitempty" xml:"Sar,omitempty"`
	StartTime      *string                                                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string                                                                                                      `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	Width          *string                                                                                                      `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNetworkCost() *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTag(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDar(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetFps(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetIndex(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLang(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLevel(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNetworkCost(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNumFrames(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetPixFmt(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetProfile(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetSar(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetStartTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetTimebase(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	AvgBitrate    *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	PreloadTime   *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig struct {
	ExtSubtitleList *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList `json:"ExtSubtitleList,omitempty" xml:"ExtSubtitleList,omitempty" type:"Struct"`
	SubtitleList    *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList    `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) GetExtSubtitleList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList {
	return s.ExtSubtitleList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) GetSubtitleList() *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList {
	return s.SubtitleList
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) SetExtSubtitleList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig {
	s.ExtSubtitleList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) SetSubtitleList(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig {
	s.SubtitleList = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfig) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList struct {
	ExtSubtitle []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle `json:"ExtSubtitle,omitempty" xml:"ExtSubtitle,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) GetExtSubtitle() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	return s.ExtSubtitle
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) SetExtSubtitle(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList {
	s.ExtSubtitle = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle struct {
	CharEnc  *string                                                                                             `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	FontName *string                                                                                             `json:"FontName,omitempty" xml:"FontName,omitempty"`
	Input    *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetCharEnc() *string {
	return s.CharEnc
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetFontName() *string {
	return s.FontName
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetInput() *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	return s.Input
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetCharEnc(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.CharEnc = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetFontName(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.FontName = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetInput(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.Input = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitle) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList struct {
	Subtitle []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) GetSubtitle() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle {
	return s.Subtitle
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) SetSubtitle(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList {
	s.Subtitle = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle struct {
	Map *string `json:"Map,omitempty" xml:"Map,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) GetMap() *string {
	return s.Map
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) SetMap(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle {
	s.Map = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSubtitleConfigSubtitleListSubtitle) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso struct {
	IsHalfSample *string `json:"IsHalfSample,omitempty" xml:"IsHalfSample,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) GetIsHalfSample() *string {
	return s.IsHalfSample
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) SetIsHalfSample(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso {
	s.IsHalfSample = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputSuperReso) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList struct {
	TailSlate []*SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate `json:"TailSlate,omitempty" xml:"TailSlate,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) GetTailSlate() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	return s.TailSlate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) SetTailSlate(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList {
	s.TailSlate = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate struct {
	BgColor       *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	BlendDuration *string `json:"BlendDuration,omitempty" xml:"BlendDuration,omitempty"`
	Height        *string `json:"Height,omitempty" xml:"Height,omitempty"`
	IsMergeAudio  *bool   `json:"IsMergeAudio,omitempty" xml:"IsMergeAudio,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	TailUrl       *string `json:"TailUrl,omitempty" xml:"TailUrl,omitempty"`
	Width         *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetBgColor() *string {
	return s.BgColor
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetBlendDuration() *string {
	return s.BlendDuration
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetIsMergeAudio() *bool {
	return s.IsMergeAudio
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetStart() *string {
	return s.Start
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetTailUrl() *string {
	return s.TailUrl
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetBgColor(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.BgColor = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetBlendDuration(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.BlendDuration = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetIsMergeAudio(v bool) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.IsMergeAudio = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetStart(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.Start = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetTailUrl(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.TailUrl = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTailSlateListTailSlate) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig struct {
	AdjDarMethod            *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	IsCheckAudioBitrate     *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetAdjDarMethod(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckAudioBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckAudioBitrateFail(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckReso(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckResoFail(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckVideoBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetIsCheckVideoBitrateFail(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) SetTransMode(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig {
	s.TransMode = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputTransConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo struct {
	Bitrate      *string                                                               `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateBnd   *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	Bufsize      *string                                                               `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec        *string                                                               `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf          *string                                                               `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop         *string                                                               `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Degrain      *string                                                               `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	Fps          *string                                                               `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop          *string                                                               `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height       *string                                                               `json:"Height,omitempty" xml:"Height,omitempty"`
	MaxFps       *string                                                               `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	Maxrate      *string                                                               `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	Pad          *string                                                               `json:"Pad,omitempty" xml:"Pad,omitempty"`
	PixFmt       *string                                                               `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Preset       *string                                                               `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Profile      *string                                                               `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale       *string                                                               `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	ResoPriority *string                                                               `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	ScanMode     *string                                                               `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width        *string                                                               `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetBitrateBnd() *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetCodec() *string {
	return s.Codec
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetCrf() *string {
	return s.Crf
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetCrop() *string {
	return s.Crop
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetFps() *string {
	return s.Fps
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetGop() *string {
	return s.Gop
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetPad() *string {
	return s.Pad
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetPreset() *string {
	return s.Preset
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetProfile() *string {
	return s.Profile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetQscale() *string {
	return s.Qscale
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetBitrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Bitrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetBitrateBnd(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.BitrateBnd = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetBufsize(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Bufsize = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetCodec(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Codec = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetCrf(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Crf = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetCrop(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Crop = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetDegrain(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Degrain = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetFps(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Fps = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetGop(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Gop = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetMaxFps(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.MaxFps = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetMaxrate(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Maxrate = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetPad(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Pad = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetPixFmt(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.PixFmt = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetPreset(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Preset = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetProfile(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Profile = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetQscale(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Qscale = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetResoPriority(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.ResoPriority = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetScanMode(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.ScanMode = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd struct {
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) SetMax(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) SetMin(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList struct {
	WaterMark []*SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark `json:"WaterMark,omitempty" xml:"WaterMark,omitempty" type:"Repeated"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) GetWaterMark() []*SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	return s.WaterMark
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) SetWaterMark(v []*SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList {
	s.WaterMark = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkList) Validate() error {
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

type SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark struct {
	Dx                  *string                                                                               `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy                  *string                                                                               `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height              *string                                                                               `json:"Height,omitempty" xml:"Height,omitempty"`
	InputFile           *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	ReferPos            *string                                                                               `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	Type                *string                                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	WaterMarkTemplateId *string                                                                               `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
	Width               *string                                                                               `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetDx() *string {
	return s.Dx
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetDy() *string {
	return s.Dy
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetHeight() *string {
	return s.Height
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetInputFile() *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile {
	return s.InputFile
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetReferPos() *string {
	return s.ReferPos
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetType() *string {
	return s.Type
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) GetWidth() *string {
	return s.Width
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetDx(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.Dx = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetDy(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.Dy = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetHeight(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.Height = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetInputFile(v *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.InputFile = v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetReferPos(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.ReferPos = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetType(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.Type = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetWaterMarkTemplateId(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) SetWidth(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark {
	s.Width = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMark) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) SetBucket(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) SetLocation(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile {
	s.Location = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) SetObject(v string) *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile {
	s.Object = &v
	return s
}

func (s *SubmitJobsResponseBodyJobResultListJobResultJobOutputWaterMarkListWaterMarkInputFile) Validate() error {
	return dara.Validate(s)
}
