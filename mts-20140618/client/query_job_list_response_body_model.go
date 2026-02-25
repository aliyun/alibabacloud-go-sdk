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
	JobList        *QueryJobListResponseBodyJobList        `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
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
	Code             *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime     *string                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FinishTime       *string                                             `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input            *QueryJobListResponseBodyJobListJobInput            `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId            *string                                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MNSMessageResult *QueryJobListResponseBodyJobListJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Output           *QueryJobListResponseBodyJobListJobOutput           `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Percent          *int64                                              `json:"Percent,omitempty" xml:"Percent,omitempty"`
	PipelineId       *string                                             `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	State            *string                                             `json:"State,omitempty" xml:"State,omitempty"`
	SubmitTime       *string                                             `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
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
	Audio                  *QueryJobListResponseBodyJobListJobOutputAudio                  `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	AudioStreamMap         *string                                                         `json:"AudioStreamMap,omitempty" xml:"AudioStreamMap,omitempty"`
	Clip                   *QueryJobListResponseBodyJobListJobOutputClip                   `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	Container              *QueryJobListResponseBodyJobListJobOutputContainer              `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	DeWatermark            *string                                                         `json:"DeWatermark,omitempty" xml:"DeWatermark,omitempty"`
	Encryption             *QueryJobListResponseBodyJobListJobOutputEncryption             `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	ExtendData             *string                                                         `json:"ExtendData,omitempty" xml:"ExtendData,omitempty"`
	M3U8NonStandardSupport *QueryJobListResponseBodyJobListJobOutputM3U8NonStandardSupport `json:"M3U8NonStandardSupport,omitempty" xml:"M3U8NonStandardSupport,omitempty" type:"Struct"`
	MergeConfigUrl         *string                                                         `json:"MergeConfigUrl,omitempty" xml:"MergeConfigUrl,omitempty"`
	MergeList              *QueryJobListResponseBodyJobListJobOutputMergeList              `json:"MergeList,omitempty" xml:"MergeList,omitempty" type:"Struct"`
	MultiSpeedInfo         *QueryJobListResponseBodyJobListJobOutputMultiSpeedInfo         `json:"MultiSpeedInfo,omitempty" xml:"MultiSpeedInfo,omitempty" type:"Struct"`
	MuxConfig              *QueryJobListResponseBodyJobListJobOutputMuxConfig              `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	OpeningList            *QueryJobListResponseBodyJobListJobOutputOpeningList            `json:"OpeningList,omitempty" xml:"OpeningList,omitempty" type:"Struct"`
	OutSubtitleList        *QueryJobListResponseBodyJobListJobOutputOutSubtitleList        `json:"OutSubtitleList,omitempty" xml:"OutSubtitleList,omitempty" type:"Struct"`
	OutputFile             *QueryJobListResponseBodyJobListJobOutputOutputFile             `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	Priority               *string                                                         `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Properties             *QueryJobListResponseBodyJobListJobOutputProperties             `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Rotate                 *string                                                         `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SubtitleConfig         *QueryJobListResponseBodyJobListJobOutputSubtitleConfig         `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Struct"`
	SuperReso              *QueryJobListResponseBodyJobListJobOutputSuperReso              `json:"SuperReso,omitempty" xml:"SuperReso,omitempty" type:"Struct"`
	TailSlateList          *QueryJobListResponseBodyJobListJobOutputTailSlateList          `json:"TailSlateList,omitempty" xml:"TailSlateList,omitempty" type:"Struct"`
	TemplateId             *string                                                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransConfig            *QueryJobListResponseBodyJobListJobOutputTransConfig            `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	UserData               *string                                                         `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Video                  *QueryJobListResponseBodyJobListJobOutputVideo                  `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
	VideoStreamMap         *string                                                         `json:"VideoStreamMap,omitempty" xml:"VideoStreamMap,omitempty"`
	WaterMarkConfigUrl     *string                                                         `json:"WaterMarkConfigUrl,omitempty" xml:"WaterMarkConfigUrl,omitempty"`
	WaterMarkList          *QueryJobListResponseBodyJobListJobOutputWaterMarkList          `json:"WaterMarkList,omitempty" xml:"WaterMarkList,omitempty" type:"Struct"`
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
	Bitrate    *string                                              `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *string                                              `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string                                              `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string                                              `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale     *string                                              `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Samplerate *string                                              `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	Volume     *QueryJobListResponseBodyJobListJobOutputAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
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
	Level  *string `json:"Level,omitempty" xml:"Level,omitempty"`
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
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Seek     *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
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
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	KeyUri  *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	SkipCnt *string `json:"SkipCnt,omitempty" xml:"SkipCnt,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Md5Support  *bool `json:"Md5Support,omitempty" xml:"Md5Support,omitempty"`
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
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MergeURL *string `json:"MergeURL,omitempty" xml:"MergeURL,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
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
	Code            *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	DowngradePolicy *string  `json:"DowngradePolicy,omitempty" xml:"DowngradePolicy,omitempty"`
	Duration        *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Enable          *string  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Message         *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RealSpeed       *float64 `json:"RealSpeed,omitempty" xml:"RealSpeed,omitempty"`
	SettingSpeed    *int32   `json:"SettingSpeed,omitempty" xml:"SettingSpeed,omitempty"`
	TimeCost        *float64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
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
	Gif     *QueryJobListResponseBodyJobListJobOutputMuxConfigGif     `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	Segment *QueryJobListResponseBodyJobListJobOutputMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	Webp    *QueryJobListResponseBodyJobListJobOutputMuxConfigWebp    `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
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
	DitherMode      *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	FinalDelay      *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	Loop            *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
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
	Height  *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Start   *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Width   *string `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Map             *string                                                                            `json:"Map,omitempty" xml:"Map,omitempty"`
	Message         *string                                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OutSubtitleFile *QueryJobListResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile `json:"OutSubtitleFile,omitempty" xml:"OutSubtitleFile,omitempty" type:"Struct"`
	Success         *bool                                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	Bitrate     *string                                                        `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration    *string                                                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileFormat  *string                                                        `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	FileSize    *string                                                        `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Format      *QueryJobListResponseBodyJobListJobOutputPropertiesFormat      `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	Fps         *string                                                        `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height      *string                                                        `json:"Height,omitempty" xml:"Height,omitempty"`
	SourceLogos *QueryJobListResponseBodyJobListJobOutputPropertiesSourceLogos `json:"SourceLogos,omitempty" xml:"SourceLogos,omitempty" type:"Struct"`
	Streams     *QueryJobListResponseBodyJobListJobOutputPropertiesStreams     `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	Width       *string                                                        `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	NumPrograms    *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	NumStreams     *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	Size           *string `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	AudioStreamList    *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList    `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	SubtitleStreamList *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	VideoStreamList    *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList    `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
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
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	AvgFPS           *string                                                                                         `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	Bitrate          *string                                                                                         `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName    *string                                                                                         `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName        *string                                                                                         `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag         *string                                                                                         `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString   *string                                                                                         `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase    *string                                                                                         `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Dar              *string                                                                                         `json:"Dar,omitempty" xml:"Dar,omitempty"`
	Duration         *string                                                                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps              *string                                                                                         `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HasBFrames       *string                                                                                         `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height           *string                                                                                         `json:"Height,omitempty" xml:"Height,omitempty"`
	Index            *string                                                                                         `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang             *string                                                                                         `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Level            *string                                                                                         `json:"Level,omitempty" xml:"Level,omitempty"`
	NetworkCost      *QueryJobListResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	NumFrames        *string                                                                                         `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	PixFmt           *string                                                                                         `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Profile          *string                                                                                         `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Sar              *string                                                                                         `json:"Sar,omitempty" xml:"Sar,omitempty"`
	StartTime        *string                                                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase         *string                                                                                         `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	Width            *string                                                                                         `json:"Width,omitempty" xml:"Width,omitempty"`
	BitsPerRawSample *string                                                                                         `json:"bitsPerRawSample,omitempty" xml:"bitsPerRawSample,omitempty"`
	ColorPrimaries   *string                                                                                         `json:"colorPrimaries,omitempty" xml:"colorPrimaries,omitempty"`
	ColorTransfer    *string                                                                                         `json:"colorTransfer,omitempty" xml:"colorTransfer,omitempty"`
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
	AvgBitrate    *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	PreloadTime   *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
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
	ExtSubtitleList *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList `json:"ExtSubtitleList,omitempty" xml:"ExtSubtitleList,omitempty" type:"Struct"`
	SubtitleList    *QueryJobListResponseBodyJobListJobOutputSubtitleConfigSubtitleList    `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty" type:"Struct"`
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
	CharEnc  *string                                                                                `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	FontName *string                                                                                `json:"FontName,omitempty" xml:"FontName,omitempty"`
	Input    *QueryJobListResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	BgColor       *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	BlendDuration *string `json:"BlendDuration,omitempty" xml:"BlendDuration,omitempty"`
	Height        *string `json:"Height,omitempty" xml:"Height,omitempty"`
	IsMergeAudio  *bool   `json:"IsMergeAudio,omitempty" xml:"IsMergeAudio,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	TailUrl       *string `json:"TailUrl,omitempty" xml:"TailUrl,omitempty"`
	Width         *string `json:"Width,omitempty" xml:"Width,omitempty"`
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
	AdjDarMethod            *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	IsCheckAudioBitrate     *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
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
	Bitrate      *string                                                  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateBnd   *QueryJobListResponseBodyJobListJobOutputVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	Bufsize      *string                                                  `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec        *string                                                  `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf          *string                                                  `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop         *string                                                  `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Degrain      *string                                                  `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	Fps          *string                                                  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop          *string                                                  `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height       *string                                                  `json:"Height,omitempty" xml:"Height,omitempty"`
	MaxFps       *string                                                  `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	Maxrate      *string                                                  `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	Pad          *string                                                  `json:"Pad,omitempty" xml:"Pad,omitempty"`
	PixFmt       *string                                                  `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Preset       *string                                                  `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Profile      *string                                                  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale       *string                                                  `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	ResoPriority *string                                                  `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	ScanMode     *string                                                  `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width        *string                                                  `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
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
	Dx                  *string                                                                  `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy                  *string                                                                  `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height              *string                                                                  `json:"Height,omitempty" xml:"Height,omitempty"`
	InputFile           *QueryJobListResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	ReferPos            *string                                                                  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	Type                *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	WaterMarkTemplateId *string                                                                  `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
	Width               *string                                                                  `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
