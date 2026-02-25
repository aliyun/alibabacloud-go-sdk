// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v *ListJobResponseBodyJobList) *ListJobResponseBody
	GetJobList() *ListJobResponseBodyJobList
	SetNextPageToken(v string) *ListJobResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListJobResponseBody
	GetRequestId() *string
}

type ListJobResponseBody struct {
	JobList *ListJobResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
	// The pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC860F04-778A-472F-AB39-E1BF329C1EA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobResponseBody) GetJobList() *ListJobResponseBodyJobList {
	return s.JobList
}

func (s *ListJobResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobResponseBody) SetJobList(v *ListJobResponseBodyJobList) *ListJobResponseBody {
	s.JobList = v
	return s
}

func (s *ListJobResponseBody) SetNextPageToken(v string) *ListJobResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListJobResponseBody) SetRequestId(v string) *ListJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobResponseBody) Validate() error {
	if s.JobList != nil {
		if err := s.JobList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobList struct {
	Job []*ListJobResponseBodyJobListJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobList) GetJob() []*ListJobResponseBodyJobListJob {
	return s.Job
}

func (s *ListJobResponseBodyJobList) SetJob(v []*ListJobResponseBodyJobListJob) *ListJobResponseBodyJobList {
	s.Job = v
	return s
}

func (s *ListJobResponseBodyJobList) Validate() error {
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

type ListJobResponseBodyJobListJob struct {
	Code             *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime     *string                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FinishTime       *string                                        `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input            *ListJobResponseBodyJobListJobInput            `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId            *string                                        `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MNSMessageResult *ListJobResponseBodyJobListJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Output           *ListJobResponseBodyJobListJobOutput           `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Percent          *int64                                         `json:"Percent,omitempty" xml:"Percent,omitempty"`
	PipelineId       *string                                        `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	State            *string                                        `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobResponseBodyJobListJob) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJob) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJob) GetCode() *string {
	return s.Code
}

func (s *ListJobResponseBodyJobListJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListJobResponseBodyJobListJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListJobResponseBodyJobListJob) GetInput() *ListJobResponseBodyJobListJobInput {
	return s.Input
}

func (s *ListJobResponseBodyJobListJob) GetJobId() *string {
	return s.JobId
}

func (s *ListJobResponseBodyJobListJob) GetMNSMessageResult() *ListJobResponseBodyJobListJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *ListJobResponseBodyJobListJob) GetMessage() *string {
	return s.Message
}

func (s *ListJobResponseBodyJobListJob) GetOutput() *ListJobResponseBodyJobListJobOutput {
	return s.Output
}

func (s *ListJobResponseBodyJobListJob) GetPercent() *int64 {
	return s.Percent
}

func (s *ListJobResponseBodyJobListJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListJobResponseBodyJobListJob) GetState() *string {
	return s.State
}

func (s *ListJobResponseBodyJobListJob) SetCode(v string) *ListJobResponseBodyJobListJob {
	s.Code = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetCreationTime(v string) *ListJobResponseBodyJobListJob {
	s.CreationTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetFinishTime(v string) *ListJobResponseBodyJobListJob {
	s.FinishTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetInput(v *ListJobResponseBodyJobListJobInput) *ListJobResponseBodyJobListJob {
	s.Input = v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetJobId(v string) *ListJobResponseBodyJobListJob {
	s.JobId = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetMNSMessageResult(v *ListJobResponseBodyJobListJobMNSMessageResult) *ListJobResponseBodyJobListJob {
	s.MNSMessageResult = v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetMessage(v string) *ListJobResponseBodyJobListJob {
	s.Message = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetOutput(v *ListJobResponseBodyJobListJobOutput) *ListJobResponseBodyJobListJob {
	s.Output = v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetPercent(v int64) *ListJobResponseBodyJobListJob {
	s.Percent = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetPipelineId(v string) *ListJobResponseBodyJobListJob {
	s.PipelineId = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) SetState(v string) *ListJobResponseBodyJobListJob {
	s.State = &v
	return s
}

func (s *ListJobResponseBodyJobListJob) Validate() error {
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

type ListJobResponseBodyJobListJobInput struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListJobResponseBodyJobListJobInput) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobInput) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *ListJobResponseBodyJobListJobInput) GetLocation() *string {
	return s.Location
}

func (s *ListJobResponseBodyJobListJobInput) GetObject() *string {
	return s.Object
}

func (s *ListJobResponseBodyJobListJobInput) SetBucket(v string) *ListJobResponseBodyJobListJobInput {
	s.Bucket = &v
	return s
}

func (s *ListJobResponseBodyJobListJobInput) SetLocation(v string) *ListJobResponseBodyJobListJobInput {
	s.Location = &v
	return s
}

func (s *ListJobResponseBodyJobListJobInput) SetObject(v string) *ListJobResponseBodyJobListJobInput {
	s.Object = &v
	return s
}

func (s *ListJobResponseBodyJobListJobInput) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobMNSMessageResult struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s ListJobResponseBodyJobListJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) SetErrorCode(v string) *ListJobResponseBodyJobListJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) SetErrorMessage(v string) *ListJobResponseBodyJobListJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) SetMessageId(v string) *ListJobResponseBodyJobListJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *ListJobResponseBodyJobListJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutput struct {
	Audio                  *ListJobResponseBodyJobListJobOutputAudio                  `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	AudioStreamMap         *string                                                    `json:"AudioStreamMap,omitempty" xml:"AudioStreamMap,omitempty"`
	Clip                   *ListJobResponseBodyJobListJobOutputClip                   `json:"Clip,omitempty" xml:"Clip,omitempty" type:"Struct"`
	Container              *ListJobResponseBodyJobListJobOutputContainer              `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	DeWatermark            *string                                                    `json:"DeWatermark,omitempty" xml:"DeWatermark,omitempty"`
	Encryption             *ListJobResponseBodyJobListJobOutputEncryption             `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	M3U8NonStandardSupport *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport `json:"M3U8NonStandardSupport,omitempty" xml:"M3U8NonStandardSupport,omitempty" type:"Struct"`
	MergeConfigUrl         *string                                                    `json:"MergeConfigUrl,omitempty" xml:"MergeConfigUrl,omitempty"`
	MergeList              *ListJobResponseBodyJobListJobOutputMergeList              `json:"MergeList,omitempty" xml:"MergeList,omitempty" type:"Struct"`
	MuxConfig              *ListJobResponseBodyJobListJobOutputMuxConfig              `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	OpeningList            *ListJobResponseBodyJobListJobOutputOpeningList            `json:"OpeningList,omitempty" xml:"OpeningList,omitempty" type:"Struct"`
	OutSubtitleList        *ListJobResponseBodyJobListJobOutputOutSubtitleList        `json:"OutSubtitleList,omitempty" xml:"OutSubtitleList,omitempty" type:"Struct"`
	OutputFile             *ListJobResponseBodyJobListJobOutputOutputFile             `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	Priority               *string                                                    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Properties             *ListJobResponseBodyJobListJobOutputProperties             `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Rotate                 *string                                                    `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SubtitleConfig         *ListJobResponseBodyJobListJobOutputSubtitleConfig         `json:"SubtitleConfig,omitempty" xml:"SubtitleConfig,omitempty" type:"Struct"`
	SuperReso              *ListJobResponseBodyJobListJobOutputSuperReso              `json:"SuperReso,omitempty" xml:"SuperReso,omitempty" type:"Struct"`
	TailSlateList          *ListJobResponseBodyJobListJobOutputTailSlateList          `json:"TailSlateList,omitempty" xml:"TailSlateList,omitempty" type:"Struct"`
	TemplateId             *string                                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransConfig            *ListJobResponseBodyJobListJobOutputTransConfig            `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	UserData               *string                                                    `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Video                  *ListJobResponseBodyJobListJobOutputVideo                  `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
	VideoStreamMap         *string                                                    `json:"VideoStreamMap,omitempty" xml:"VideoStreamMap,omitempty"`
	WaterMarkConfigUrl     *string                                                    `json:"WaterMarkConfigUrl,omitempty" xml:"WaterMarkConfigUrl,omitempty"`
	WaterMarkList          *ListJobResponseBodyJobListJobOutputWaterMarkList          `json:"WaterMarkList,omitempty" xml:"WaterMarkList,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutput) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutput) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutput) GetAudio() *ListJobResponseBodyJobListJobOutputAudio {
	return s.Audio
}

func (s *ListJobResponseBodyJobListJobOutput) GetAudioStreamMap() *string {
	return s.AudioStreamMap
}

func (s *ListJobResponseBodyJobListJobOutput) GetClip() *ListJobResponseBodyJobListJobOutputClip {
	return s.Clip
}

func (s *ListJobResponseBodyJobListJobOutput) GetContainer() *ListJobResponseBodyJobListJobOutputContainer {
	return s.Container
}

func (s *ListJobResponseBodyJobListJobOutput) GetDeWatermark() *string {
	return s.DeWatermark
}

func (s *ListJobResponseBodyJobListJobOutput) GetEncryption() *ListJobResponseBodyJobListJobOutputEncryption {
	return s.Encryption
}

func (s *ListJobResponseBodyJobListJobOutput) GetM3U8NonStandardSupport() *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport {
	return s.M3U8NonStandardSupport
}

func (s *ListJobResponseBodyJobListJobOutput) GetMergeConfigUrl() *string {
	return s.MergeConfigUrl
}

func (s *ListJobResponseBodyJobListJobOutput) GetMergeList() *ListJobResponseBodyJobListJobOutputMergeList {
	return s.MergeList
}

func (s *ListJobResponseBodyJobListJobOutput) GetMuxConfig() *ListJobResponseBodyJobListJobOutputMuxConfig {
	return s.MuxConfig
}

func (s *ListJobResponseBodyJobListJobOutput) GetOpeningList() *ListJobResponseBodyJobListJobOutputOpeningList {
	return s.OpeningList
}

func (s *ListJobResponseBodyJobListJobOutput) GetOutSubtitleList() *ListJobResponseBodyJobListJobOutputOutSubtitleList {
	return s.OutSubtitleList
}

func (s *ListJobResponseBodyJobListJobOutput) GetOutputFile() *ListJobResponseBodyJobListJobOutputOutputFile {
	return s.OutputFile
}

func (s *ListJobResponseBodyJobListJobOutput) GetPriority() *string {
	return s.Priority
}

func (s *ListJobResponseBodyJobListJobOutput) GetProperties() *ListJobResponseBodyJobListJobOutputProperties {
	return s.Properties
}

func (s *ListJobResponseBodyJobListJobOutput) GetRotate() *string {
	return s.Rotate
}

func (s *ListJobResponseBodyJobListJobOutput) GetSubtitleConfig() *ListJobResponseBodyJobListJobOutputSubtitleConfig {
	return s.SubtitleConfig
}

func (s *ListJobResponseBodyJobListJobOutput) GetSuperReso() *ListJobResponseBodyJobListJobOutputSuperReso {
	return s.SuperReso
}

func (s *ListJobResponseBodyJobListJobOutput) GetTailSlateList() *ListJobResponseBodyJobListJobOutputTailSlateList {
	return s.TailSlateList
}

func (s *ListJobResponseBodyJobListJobOutput) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobResponseBodyJobListJobOutput) GetTransConfig() *ListJobResponseBodyJobListJobOutputTransConfig {
	return s.TransConfig
}

func (s *ListJobResponseBodyJobListJobOutput) GetUserData() *string {
	return s.UserData
}

func (s *ListJobResponseBodyJobListJobOutput) GetVideo() *ListJobResponseBodyJobListJobOutputVideo {
	return s.Video
}

func (s *ListJobResponseBodyJobListJobOutput) GetVideoStreamMap() *string {
	return s.VideoStreamMap
}

func (s *ListJobResponseBodyJobListJobOutput) GetWaterMarkConfigUrl() *string {
	return s.WaterMarkConfigUrl
}

func (s *ListJobResponseBodyJobListJobOutput) GetWaterMarkList() *ListJobResponseBodyJobListJobOutputWaterMarkList {
	return s.WaterMarkList
}

func (s *ListJobResponseBodyJobListJobOutput) SetAudio(v *ListJobResponseBodyJobListJobOutputAudio) *ListJobResponseBodyJobListJobOutput {
	s.Audio = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetAudioStreamMap(v string) *ListJobResponseBodyJobListJobOutput {
	s.AudioStreamMap = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetClip(v *ListJobResponseBodyJobListJobOutputClip) *ListJobResponseBodyJobListJobOutput {
	s.Clip = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetContainer(v *ListJobResponseBodyJobListJobOutputContainer) *ListJobResponseBodyJobListJobOutput {
	s.Container = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetDeWatermark(v string) *ListJobResponseBodyJobListJobOutput {
	s.DeWatermark = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetEncryption(v *ListJobResponseBodyJobListJobOutputEncryption) *ListJobResponseBodyJobListJobOutput {
	s.Encryption = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetM3U8NonStandardSupport(v *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) *ListJobResponseBodyJobListJobOutput {
	s.M3U8NonStandardSupport = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetMergeConfigUrl(v string) *ListJobResponseBodyJobListJobOutput {
	s.MergeConfigUrl = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetMergeList(v *ListJobResponseBodyJobListJobOutputMergeList) *ListJobResponseBodyJobListJobOutput {
	s.MergeList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetMuxConfig(v *ListJobResponseBodyJobListJobOutputMuxConfig) *ListJobResponseBodyJobListJobOutput {
	s.MuxConfig = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetOpeningList(v *ListJobResponseBodyJobListJobOutputOpeningList) *ListJobResponseBodyJobListJobOutput {
	s.OpeningList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetOutSubtitleList(v *ListJobResponseBodyJobListJobOutputOutSubtitleList) *ListJobResponseBodyJobListJobOutput {
	s.OutSubtitleList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetOutputFile(v *ListJobResponseBodyJobListJobOutputOutputFile) *ListJobResponseBodyJobListJobOutput {
	s.OutputFile = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetPriority(v string) *ListJobResponseBodyJobListJobOutput {
	s.Priority = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetProperties(v *ListJobResponseBodyJobListJobOutputProperties) *ListJobResponseBodyJobListJobOutput {
	s.Properties = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetRotate(v string) *ListJobResponseBodyJobListJobOutput {
	s.Rotate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetSubtitleConfig(v *ListJobResponseBodyJobListJobOutputSubtitleConfig) *ListJobResponseBodyJobListJobOutput {
	s.SubtitleConfig = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetSuperReso(v *ListJobResponseBodyJobListJobOutputSuperReso) *ListJobResponseBodyJobListJobOutput {
	s.SuperReso = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetTailSlateList(v *ListJobResponseBodyJobListJobOutputTailSlateList) *ListJobResponseBodyJobListJobOutput {
	s.TailSlateList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetTemplateId(v string) *ListJobResponseBodyJobListJobOutput {
	s.TemplateId = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetTransConfig(v *ListJobResponseBodyJobListJobOutputTransConfig) *ListJobResponseBodyJobListJobOutput {
	s.TransConfig = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetUserData(v string) *ListJobResponseBodyJobListJobOutput {
	s.UserData = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetVideo(v *ListJobResponseBodyJobListJobOutputVideo) *ListJobResponseBodyJobListJobOutput {
	s.Video = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetVideoStreamMap(v string) *ListJobResponseBodyJobListJobOutput {
	s.VideoStreamMap = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetWaterMarkConfigUrl(v string) *ListJobResponseBodyJobListJobOutput {
	s.WaterMarkConfigUrl = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) SetWaterMarkList(v *ListJobResponseBodyJobListJobOutputWaterMarkList) *ListJobResponseBodyJobListJobOutput {
	s.WaterMarkList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutput) Validate() error {
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

type ListJobResponseBodyJobListJobOutputAudio struct {
	Bitrate    *string                                         `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *string                                         `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string                                         `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string                                         `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale     *string                                         `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	Samplerate *string                                         `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	Volume     *ListJobResponseBodyJobListJobOutputAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputAudio) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputAudio) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetChannels() *string {
	return s.Channels
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetCodec() *string {
	return s.Codec
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetProfile() *string {
	return s.Profile
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetQscale() *string {
	return s.Qscale
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *ListJobResponseBodyJobListJobOutputAudio) GetVolume() *ListJobResponseBodyJobListJobOutputAudioVolume {
	return s.Volume
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetChannels(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Channels = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetCodec(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Codec = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetProfile(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Profile = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetQscale(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Qscale = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetSamplerate(v string) *ListJobResponseBodyJobListJobOutputAudio {
	s.Samplerate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) SetVolume(v *ListJobResponseBodyJobListJobOutputAudioVolume) *ListJobResponseBodyJobListJobOutputAudio {
	s.Volume = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputAudioVolume struct {
	Level  *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputAudioVolume) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputAudioVolume) GetLevel() *string {
	return s.Level
}

func (s *ListJobResponseBodyJobListJobOutputAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *ListJobResponseBodyJobListJobOutputAudioVolume) SetLevel(v string) *ListJobResponseBodyJobListJobOutputAudioVolume {
	s.Level = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudioVolume) SetMethod(v string) *ListJobResponseBodyJobListJobOutputAudioVolume {
	s.Method = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputAudioVolume) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputClip struct {
	TimeSpan *ListJobResponseBodyJobListJobOutputClipTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputClip) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputClip) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputClip) GetTimeSpan() *ListJobResponseBodyJobListJobOutputClipTimeSpan {
	return s.TimeSpan
}

func (s *ListJobResponseBodyJobListJobOutputClip) SetTimeSpan(v *ListJobResponseBodyJobListJobOutputClipTimeSpan) *ListJobResponseBodyJobListJobOutputClip {
	s.TimeSpan = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputClip) Validate() error {
	if s.TimeSpan != nil {
		if err := s.TimeSpan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputClipTimeSpan struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Seek     *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputClipTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputClipTimeSpan) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputClipTimeSpan) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputClipTimeSpan) GetSeek() *string {
	return s.Seek
}

func (s *ListJobResponseBodyJobListJobOutputClipTimeSpan) SetDuration(v string) *ListJobResponseBodyJobListJobOutputClipTimeSpan {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputClipTimeSpan) SetSeek(v string) *ListJobResponseBodyJobListJobOutputClipTimeSpan {
	s.Seek = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputClipTimeSpan) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputContainer struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputContainer) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputContainer) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputContainer) GetFormat() *string {
	return s.Format
}

func (s *ListJobResponseBodyJobListJobOutputContainer) SetFormat(v string) *ListJobResponseBodyJobListJobOutputContainer {
	s.Format = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputContainer) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputEncryption struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	KeyUri  *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	SkipCnt *string `json:"SkipCnt,omitempty" xml:"SkipCnt,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputEncryption) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputEncryption) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetId() *string {
	return s.Id
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetKey() *string {
	return s.Key
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetKeyType() *string {
	return s.KeyType
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetKeyUri() *string {
	return s.KeyUri
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetSkipCnt() *string {
	return s.SkipCnt
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) GetType() *string {
	return s.Type
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetId(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.Id = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetKey(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.Key = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetKeyType(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.KeyType = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetKeyUri(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.KeyUri = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetSkipCnt(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.SkipCnt = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) SetType(v string) *ListJobResponseBodyJobListJobOutputEncryption {
	s.Type = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputEncryption) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport struct {
	TS *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS `json:"TS,omitempty" xml:"TS,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) GetTS() *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	return s.TS
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) SetTS(v *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport {
	s.TS = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupport) Validate() error {
	if s.TS != nil {
		if err := s.TS.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS struct {
	Md5Support  *bool `json:"Md5Support,omitempty" xml:"Md5Support,omitempty"`
	SizeSupport *bool `json:"SizeSupport,omitempty" xml:"SizeSupport,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GetMd5Support() *bool {
	return s.Md5Support
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) GetSizeSupport() *bool {
	return s.SizeSupport
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) SetMd5Support(v bool) *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	s.Md5Support = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) SetSizeSupport(v bool) *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS {
	s.SizeSupport = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputM3U8NonStandardSupportTS) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputMergeList struct {
	Merge []*ListJobResponseBodyJobListJobOutputMergeListMerge `json:"Merge,omitempty" xml:"Merge,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputMergeList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMergeList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMergeList) GetMerge() []*ListJobResponseBodyJobListJobOutputMergeListMerge {
	return s.Merge
}

func (s *ListJobResponseBodyJobListJobOutputMergeList) SetMerge(v []*ListJobResponseBodyJobListJobOutputMergeListMerge) *ListJobResponseBodyJobListJobOutputMergeList {
	s.Merge = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMergeList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputMergeListMerge struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MergeURL *string `json:"MergeURL,omitempty" xml:"MergeURL,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputMergeListMerge) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMergeListMerge) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) GetMergeURL() *string {
	return s.MergeURL
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) GetStart() *string {
	return s.Start
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) SetDuration(v string) *ListJobResponseBodyJobListJobOutputMergeListMerge {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) SetMergeURL(v string) *ListJobResponseBodyJobListJobOutputMergeListMerge {
	s.MergeURL = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) SetRoleArn(v string) *ListJobResponseBodyJobListJobOutputMergeListMerge {
	s.RoleArn = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) SetStart(v string) *ListJobResponseBodyJobListJobOutputMergeListMerge {
	s.Start = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMergeListMerge) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputMuxConfig struct {
	Gif     *ListJobResponseBodyJobListJobOutputMuxConfigGif     `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	Segment *ListJobResponseBodyJobListJobOutputMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	Webp    *ListJobResponseBodyJobListJobOutputMuxConfigWebp    `json:"Webp,omitempty" xml:"Webp,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMuxConfig) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) GetGif() *ListJobResponseBodyJobListJobOutputMuxConfigGif {
	return s.Gif
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) GetSegment() *ListJobResponseBodyJobListJobOutputMuxConfigSegment {
	return s.Segment
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) GetWebp() *ListJobResponseBodyJobListJobOutputMuxConfigWebp {
	return s.Webp
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) SetGif(v *ListJobResponseBodyJobListJobOutputMuxConfigGif) *ListJobResponseBodyJobListJobOutputMuxConfig {
	s.Gif = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) SetSegment(v *ListJobResponseBodyJobListJobOutputMuxConfigSegment) *ListJobResponseBodyJobListJobOutputMuxConfig {
	s.Segment = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) SetWebp(v *ListJobResponseBodyJobListJobOutputMuxConfigWebp) *ListJobResponseBodyJobListJobOutputMuxConfig {
	s.Webp = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfig) Validate() error {
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

type ListJobResponseBodyJobListJobOutputMuxConfigGif struct {
	DitherMode      *string `json:"DitherMode,omitempty" xml:"DitherMode,omitempty"`
	FinalDelay      *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	IsCustomPalette *string `json:"IsCustomPalette,omitempty" xml:"IsCustomPalette,omitempty"`
	Loop            *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigGif) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) GetDitherMode() *string {
	return s.DitherMode
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) GetIsCustomPalette() *string {
	return s.IsCustomPalette
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) SetDitherMode(v string) *ListJobResponseBodyJobListJobOutputMuxConfigGif {
	s.DitherMode = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) SetFinalDelay(v string) *ListJobResponseBodyJobListJobOutputMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) SetIsCustomPalette(v string) *ListJobResponseBodyJobListJobOutputMuxConfigGif {
	s.IsCustomPalette = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) SetLoop(v string) *ListJobResponseBodyJobListJobOutputMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputMuxConfigSegment struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigSegment) SetDuration(v string) *ListJobResponseBodyJobListJobOutputMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputMuxConfigWebp struct {
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigWebp) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputMuxConfigWebp) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigWebp) GetLoop() *string {
	return s.Loop
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigWebp) SetLoop(v string) *ListJobResponseBodyJobListJobOutputMuxConfigWebp {
	s.Loop = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputMuxConfigWebp) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputOpeningList struct {
	Opening []*ListJobResponseBodyJobListJobOutputOpeningListOpening `json:"Opening,omitempty" xml:"Opening,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputOpeningList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOpeningList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOpeningList) GetOpening() []*ListJobResponseBodyJobListJobOutputOpeningListOpening {
	return s.Opening
}

func (s *ListJobResponseBodyJobListJobOutputOpeningList) SetOpening(v []*ListJobResponseBodyJobListJobOutputOpeningListOpening) *ListJobResponseBodyJobListJobOutputOpeningList {
	s.Opening = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOpeningList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputOpeningListOpening struct {
	Height  *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Start   *string `json:"Start,omitempty" xml:"Start,omitempty"`
	Width   *string `json:"Width,omitempty" xml:"Width,omitempty"`
	OpenUrl *string `json:"openUrl,omitempty" xml:"openUrl,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputOpeningListOpening) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOpeningListOpening) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) GetStart() *string {
	return s.Start
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) GetOpenUrl() *string {
	return s.OpenUrl
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) SetHeight(v string) *ListJobResponseBodyJobListJobOutputOpeningListOpening {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) SetStart(v string) *ListJobResponseBodyJobListJobOutputOpeningListOpening {
	s.Start = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) SetWidth(v string) *ListJobResponseBodyJobListJobOutputOpeningListOpening {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) SetOpenUrl(v string) *ListJobResponseBodyJobListJobOutputOpeningListOpening {
	s.OpenUrl = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOpeningListOpening) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputOutSubtitleList struct {
	OutSubtitle []*ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle `json:"OutSubtitle,omitempty" xml:"OutSubtitle,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleList) GetOutSubtitle() []*ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	return s.OutSubtitle
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleList) SetOutSubtitle(v []*ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) *ListJobResponseBodyJobListJobOutputOutSubtitleList {
	s.OutSubtitle = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle struct {
	Map             *string                                                                       `json:"Map,omitempty" xml:"Map,omitempty"`
	Message         *string                                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	OutSubtitleFile *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile `json:"OutSubtitleFile,omitempty" xml:"OutSubtitleFile,omitempty" type:"Struct"`
	Success         *bool                                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetMap() *string {
	return s.Map
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetMessage() *string {
	return s.Message
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetOutSubtitleFile() *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	return s.OutSubtitleFile
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetMap(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Map = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetMessage(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Message = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetOutSubtitleFile(v *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.OutSubtitleFile = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) SetSuccess(v bool) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle {
	s.Success = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitle) Validate() error {
	if s.OutSubtitleFile != nil {
		if err := s.OutSubtitleFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetLocation() *string {
	return s.Location
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetObject() *string {
	return s.Object
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetBucket(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Bucket = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetLocation(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Location = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetObject(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.Object = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) SetRoleArn(v string) *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile {
	s.RoleArn = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutSubtitleListOutSubtitleOutSubtitleFile) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputOutputFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputOutputFile) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputOutputFile) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) GetLocation() *string {
	return s.Location
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) GetObject() *string {
	return s.Object
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) SetBucket(v string) *ListJobResponseBodyJobListJobOutputOutputFile {
	s.Bucket = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) SetLocation(v string) *ListJobResponseBodyJobListJobOutputOutputFile {
	s.Location = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) SetObject(v string) *ListJobResponseBodyJobListJobOutputOutputFile {
	s.Object = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) SetRoleArn(v string) *ListJobResponseBodyJobListJobOutputOutputFile {
	s.RoleArn = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputOutputFile) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputProperties struct {
	Bitrate    *string                                               `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration   *string                                               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileFormat *string                                               `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	FileSize   *string                                               `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Format     *ListJobResponseBodyJobListJobOutputPropertiesFormat  `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	Fps        *string                                               `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height     *string                                               `json:"Height,omitempty" xml:"Height,omitempty"`
	Streams    *ListJobResponseBodyJobListJobOutputPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	Width      *string                                               `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputProperties) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputProperties) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetFileFormat() *string {
	return s.FileFormat
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetFileSize() *string {
	return s.FileSize
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetFormat() *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	return s.Format
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetFps() *string {
	return s.Fps
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetStreams() *ListJobResponseBodyJobListJobOutputPropertiesStreams {
	return s.Streams
}

func (s *ListJobResponseBodyJobListJobOutputProperties) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetDuration(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetFileFormat(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.FileFormat = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetFileSize(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.FileSize = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetFormat(v *ListJobResponseBodyJobListJobOutputPropertiesFormat) *ListJobResponseBodyJobListJobOutputProperties {
	s.Format = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetFps(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.Fps = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetHeight(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetStreams(v *ListJobResponseBodyJobListJobOutputPropertiesStreams) *ListJobResponseBodyJobListJobOutputProperties {
	s.Streams = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) SetWidth(v string) *ListJobResponseBodyJobListJobOutputProperties {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputProperties) Validate() error {
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

type ListJobResponseBodyJobListJobOutputPropertiesFormat struct {
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	NumPrograms    *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	NumStreams     *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	Size           *string `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesFormat) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesFormat) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetSize() *string {
	return s.Size
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetDuration(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetFormatLongName(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.FormatLongName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetFormatName(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.FormatName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetNumPrograms(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.NumPrograms = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetNumStreams(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.NumStreams = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetSize(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.Size = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) SetStartTime(v string) *ListJobResponseBodyJobListJobOutputPropertiesFormat {
	s.StartTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputPropertiesStreams struct {
	AudioStreamList    *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList    `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	SubtitleStreamList *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	VideoStreamList    *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList    `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreams) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreams) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) GetAudioStreamList() *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) GetSubtitleStreamList() *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) GetVideoStreamList() *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) SetAudioStreamList(v *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) *ListJobResponseBodyJobListJobOutputPropertiesStreams {
	s.AudioStreamList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) SetSubtitleStreamList(v *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) *ListJobResponseBodyJobListJobOutputPropertiesStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) SetVideoStreamList(v *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) *ListJobResponseBodyJobListJobOutputPropertiesStreams {
	s.VideoStreamList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreams) Validate() error {
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

type ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList struct {
	AudioStream []*ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) GetAudioStream() []*ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) SetAudioStream(v []*ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream struct {
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

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetChannels(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecName(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTag(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetDuration(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetIndex(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetLang(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetNumFrames(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetSamplerate(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetStartTime(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) SetTimebase(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList struct {
	SubtitleStream []*ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) GetSubtitleStream() []*ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) SetSubtitleStream(v []*ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream struct {
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList struct {
	VideoStream []*ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) GetVideoStream() []*ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) SetVideoStream(v []*ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream struct {
	AvgFPS         *string                                                                                    `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	Bitrate        *string                                                                                    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName  *string                                                                                    `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string                                                                                    `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string                                                                                    `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string                                                                                    `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string                                                                                    `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Dar            *string                                                                                    `json:"Dar,omitempty" xml:"Dar,omitempty"`
	Duration       *string                                                                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps            *string                                                                                    `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HasBFrames     *string                                                                                    `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height         *string                                                                                    `json:"Height,omitempty" xml:"Height,omitempty"`
	Index          *string                                                                                    `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string                                                                                    `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Level          *string                                                                                    `json:"Level,omitempty" xml:"Level,omitempty"`
	NetworkCost    *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	NumFrames      *string                                                                                    `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	PixFmt         *string                                                                                    `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Profile        *string                                                                                    `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Sar            *string                                                                                    `json:"Sar,omitempty" xml:"Sar,omitempty"`
	StartTime      *string                                                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string                                                                                    `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	Width          *string                                                                                    `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNetworkCost() *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecName(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTag(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDar(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetDuration(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetFps(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetHeight(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetIndex(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLang(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetLevel(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNetworkCost(v *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetNumFrames(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetPixFmt(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetProfile(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetSar(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetStartTime(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetTimebase(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) SetWidth(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	AvgBitrate    *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	PreloadTime   *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputPropertiesStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputSubtitleConfig struct {
	ExtSubtitleList *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList `json:"ExtSubtitleList,omitempty" xml:"ExtSubtitleList,omitempty" type:"Struct"`
	SubtitleList    *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList    `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfig) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfig) GetExtSubtitleList() *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList {
	return s.ExtSubtitleList
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfig) GetSubtitleList() *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList {
	return s.SubtitleList
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfig) SetExtSubtitleList(v *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) *ListJobResponseBodyJobListJobOutputSubtitleConfig {
	s.ExtSubtitleList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfig) SetSubtitleList(v *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) *ListJobResponseBodyJobListJobOutputSubtitleConfig {
	s.SubtitleList = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfig) Validate() error {
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

type ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList struct {
	ExtSubtitle []*ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle `json:"ExtSubtitle,omitempty" xml:"ExtSubtitle,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) GetExtSubtitle() []*ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	return s.ExtSubtitle
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) SetExtSubtitle(v []*ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList {
	s.ExtSubtitle = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle struct {
	CharEnc  *string                                                                           `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	FontName *string                                                                           `json:"FontName,omitempty" xml:"FontName,omitempty"`
	Input    *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetCharEnc() *string {
	return s.CharEnc
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetFontName() *string {
	return s.FontName
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) GetInput() *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	return s.Input
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetCharEnc(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.CharEnc = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetFontName(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.FontName = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) SetInput(v *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle {
	s.Input = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitle) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetBucket() *string {
	return s.Bucket
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetLocation() *string {
	return s.Location
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) GetObject() *string {
	return s.Object
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetBucket(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Bucket = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetLocation(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Location = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) SetObject(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput {
	s.Object = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigExtSubtitleListExtSubtitleInput) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList struct {
	Subtitle []*ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) GetSubtitle() []*ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle {
	return s.Subtitle
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) SetSubtitle(v []*ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList {
	s.Subtitle = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle struct {
	Map *string `json:"Map,omitempty" xml:"Map,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) GetMap() *string {
	return s.Map
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) SetMap(v string) *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle {
	s.Map = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSubtitleConfigSubtitleListSubtitle) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputSuperReso struct {
	IsHalfSample *string `json:"IsHalfSample,omitempty" xml:"IsHalfSample,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputSuperReso) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputSuperReso) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputSuperReso) GetIsHalfSample() *string {
	return s.IsHalfSample
}

func (s *ListJobResponseBodyJobListJobOutputSuperReso) SetIsHalfSample(v string) *ListJobResponseBodyJobListJobOutputSuperReso {
	s.IsHalfSample = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputSuperReso) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputTailSlateList struct {
	TailSlate []*ListJobResponseBodyJobListJobOutputTailSlateListTailSlate `json:"TailSlate,omitempty" xml:"TailSlate,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputTailSlateList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputTailSlateList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateList) GetTailSlate() []*ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	return s.TailSlate
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateList) SetTailSlate(v []*ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) *ListJobResponseBodyJobListJobOutputTailSlateList {
	s.TailSlate = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputTailSlateListTailSlate struct {
	BgColor       *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	BlendDuration *string `json:"BlendDuration,omitempty" xml:"BlendDuration,omitempty"`
	Height        *string `json:"Height,omitempty" xml:"Height,omitempty"`
	IsMergeAudio  *bool   `json:"IsMergeAudio,omitempty" xml:"IsMergeAudio,omitempty"`
	Start         *string `json:"Start,omitempty" xml:"Start,omitempty"`
	TailUrl       *string `json:"TailUrl,omitempty" xml:"TailUrl,omitempty"`
	Width         *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetBgColor() *string {
	return s.BgColor
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetBlendDuration() *string {
	return s.BlendDuration
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetIsMergeAudio() *bool {
	return s.IsMergeAudio
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetStart() *string {
	return s.Start
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetTailUrl() *string {
	return s.TailUrl
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetBgColor(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.BgColor = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetBlendDuration(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.BlendDuration = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetHeight(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetIsMergeAudio(v bool) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.IsMergeAudio = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetStart(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Start = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetTailUrl(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.TailUrl = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) SetWidth(v string) *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTailSlateListTailSlate) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputTransConfig struct {
	AdjDarMethod            *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	IsCheckAudioBitrate     *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputTransConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputTransConfig) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetAdjDarMethod(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckAudioBitrate(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckAudioBitrateFail(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckReso(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckResoFail(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckVideoBitrate(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetIsCheckVideoBitrateFail(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) SetTransMode(v string) *ListJobResponseBodyJobListJobOutputTransConfig {
	s.TransMode = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputTransConfig) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputVideo struct {
	Bitrate      *string                                             `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateBnd   *ListJobResponseBodyJobListJobOutputVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	Bufsize      *string                                             `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec        *string                                             `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf          *string                                             `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop         *string                                             `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Degrain      *string                                             `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	Fps          *string                                             `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop          *string                                             `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height       *string                                             `json:"Height,omitempty" xml:"Height,omitempty"`
	MaxFps       *string                                             `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	Maxrate      *string                                             `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	Pad          *string                                             `json:"Pad,omitempty" xml:"Pad,omitempty"`
	PixFmt       *string                                             `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Preset       *string                                             `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Profile      *string                                             `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Qscale       *string                                             `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	ResoPriority *string                                             `json:"ResoPriority,omitempty" xml:"ResoPriority,omitempty"`
	ScanMode     *string                                             `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width        *string                                             `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputVideo) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputVideo) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetBitrateBnd() *ListJobResponseBodyJobListJobOutputVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetCodec() *string {
	return s.Codec
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetCrf() *string {
	return s.Crf
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetCrop() *string {
	return s.Crop
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetFps() *string {
	return s.Fps
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetGop() *string {
	return s.Gop
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetMaxFps() *string {
	return s.MaxFps
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetPad() *string {
	return s.Pad
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetPreset() *string {
	return s.Preset
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetProfile() *string {
	return s.Profile
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetQscale() *string {
	return s.Qscale
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetResoPriority() *string {
	return s.ResoPriority
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *ListJobResponseBodyJobListJobOutputVideo) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetBitrate(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Bitrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetBitrateBnd(v *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) *ListJobResponseBodyJobListJobOutputVideo {
	s.BitrateBnd = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetBufsize(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Bufsize = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetCodec(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Codec = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetCrf(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Crf = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetCrop(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Crop = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetDegrain(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Degrain = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetFps(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Fps = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetGop(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Gop = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetHeight(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetMaxFps(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.MaxFps = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetMaxrate(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Maxrate = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetPad(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Pad = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetPixFmt(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.PixFmt = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetPreset(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Preset = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetProfile(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Profile = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetQscale(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Qscale = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetResoPriority(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.ResoPriority = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetScanMode(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.ScanMode = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) SetWidth(v string) *ListJobResponseBodyJobListJobOutputVideo {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputVideoBitrateBnd struct {
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) SetMax(v string) *ListJobResponseBodyJobListJobOutputVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) SetMin(v string) *ListJobResponseBodyJobListJobOutputVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type ListJobResponseBodyJobListJobOutputWaterMarkList struct {
	WaterMark []*ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark `json:"WaterMark,omitempty" xml:"WaterMark,omitempty" type:"Repeated"`
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkList) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkList) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkList) GetWaterMark() []*ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	return s.WaterMark
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkList) SetWaterMark(v []*ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) *ListJobResponseBodyJobListJobOutputWaterMarkList {
	s.WaterMark = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkList) Validate() error {
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

type ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark struct {
	Dx                  *string                                                             `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy                  *string                                                             `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height              *string                                                             `json:"Height,omitempty" xml:"Height,omitempty"`
	InputFile           *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	ReferPos            *string                                                             `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	Type                *string                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	WaterMarkTemplateId *string                                                             `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
	Width               *string                                                             `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetDx() *string {
	return s.Dx
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetDy() *string {
	return s.Dy
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetHeight() *string {
	return s.Height
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetInputFile() *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	return s.InputFile
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetReferPos() *string {
	return s.ReferPos
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetType() *string {
	return s.Type
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) GetWidth() *string {
	return s.Width
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetDx(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Dx = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetDy(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Dy = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetHeight(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Height = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetInputFile(v *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.InputFile = v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetReferPos(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.ReferPos = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetType(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Type = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetWaterMarkTemplateId(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) SetWidth(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark {
	s.Width = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMark) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GoString() string {
	return s.String()
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetLocation() *string {
	return s.Location
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) GetObject() *string {
	return s.Object
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetBucket(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Bucket = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetLocation(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Location = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) SetObject(v string) *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile {
	s.Object = &v
	return s
}

func (s *ListJobResponseBodyJobListJobOutputWaterMarkListWaterMarkInputFile) Validate() error {
	return dara.Validate(s)
}
