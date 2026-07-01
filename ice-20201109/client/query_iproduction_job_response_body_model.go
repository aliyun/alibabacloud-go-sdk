// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *QueryIProductionJobResponseBody
	GetCreateTime() *string
	SetFinishTime(v string) *QueryIProductionJobResponseBody
	GetFinishTime() *string
	SetFunctionName(v string) *QueryIProductionJobResponseBody
	GetFunctionName() *string
	SetInput(v *QueryIProductionJobResponseBodyInput) *QueryIProductionJobResponseBody
	GetInput() *QueryIProductionJobResponseBodyInput
	SetJobId(v string) *QueryIProductionJobResponseBody
	GetJobId() *string
	SetJobParams(v string) *QueryIProductionJobResponseBody
	GetJobParams() *string
	SetName(v string) *QueryIProductionJobResponseBody
	GetName() *string
	SetOutput(v *QueryIProductionJobResponseBodyOutput) *QueryIProductionJobResponseBody
	GetOutput() *QueryIProductionJobResponseBodyOutput
	SetOutputFiles(v []*string) *QueryIProductionJobResponseBody
	GetOutputFiles() []*string
	SetOutputMediaIds(v []*string) *QueryIProductionJobResponseBody
	GetOutputMediaIds() []*string
	SetOutputUrls(v []*string) *QueryIProductionJobResponseBody
	GetOutputUrls() []*string
	SetRequestId(v string) *QueryIProductionJobResponseBody
	GetRequestId() *string
	SetResult(v string) *QueryIProductionJobResponseBody
	GetResult() *string
	SetScheduleConfig(v *QueryIProductionJobResponseBodyScheduleConfig) *QueryIProductionJobResponseBody
	GetScheduleConfig() *QueryIProductionJobResponseBodyScheduleConfig
	SetStatus(v string) *QueryIProductionJobResponseBody
	GetStatus() *string
	SetTemplateId(v string) *QueryIProductionJobResponseBody
	GetTemplateId() *string
	SetUserData(v string) *QueryIProductionJobResponseBody
	GetUserData() *string
}

type QueryIProductionJobResponseBody struct {
	// The time when the job was created, in UTC.
	//
	// example:
	//
	// 2022-07-07T07:16:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was completed, in UTC.
	//
	// example:
	//
	// 2021-11-26T14:50:25Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The name of the algorithm to use. Valid values:
	//
	// - **Cover**: smart cover
	//
	// - **VideoClip**: video summary
	//
	// - **VideoDelogo**: video logo removal
	//
	// - **VideoDetext**: video text removal
	//
	// - **CaptionExtraction**: caption extraction
	//
	// - **VideoGreenScreenMatting**: green screen matting
	//
	// - **FaceBeauty**: video beautification
	//
	// - **VideoH2V**: horizontal-to-vertical video conversion
	//
	// - **MusicSegmentDetect**: chorus detection
	//
	// - **AudioBeatDetection**: beat detection
	//
	// - **AudioQualityAssessment**: audio quality assessment
	//
	// - **SpeechDenoise**: speech denoising
	//
	// - **AudioMixing**: audio mixing
	//
	// - **MusicDemix**: music source separation
	//
	// example:
	//
	// Cover
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The input media.
	Input *QueryIProductionJobResponseBodyInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// A JSON object that contains the parameters for the algorithm job. The specific parameters vary depending on the selected algorithm.
	//
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The job name.
	//
	// example:
	//
	// Test task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output media.
	Output *QueryIProductionJobResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// An array of output file paths.
	OutputFiles []*string `json:"OutputFiles,omitempty" xml:"OutputFiles,omitempty" type:"Repeated"`
	// The output media asset IDs.
	OutputMediaIds []*string `json:"OutputMediaIds,omitempty" xml:"OutputMediaIds,omitempty" type:"Repeated"`
	// An array of output file URLs.
	OutputUrls []*string `json:"OutputUrls,omitempty" xml:"OutputUrls,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The algorithm output, returned as a JSON string. The structure of the output varies based on the `FunctionName`. For more information, see the additional notes below.
	//
	// example:
	//
	// {}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The job configuration.
	ScheduleConfig *QueryIProductionJobResponseBodyScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The job status. Valid values:
	//
	// - Queuing: The job is awaiting processing.
	//
	// - Analyzing: The job is being processed.
	//
	// - Fail: The job failed to complete.
	//
	// - Success: The job completed successfully.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The user data. The system returns this value unchanged.
	//
	// example:
	//
	// {"test":1}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryIProductionJobResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryIProductionJobResponseBody) GetFunctionName() *string {
	return s.FunctionName
}

func (s *QueryIProductionJobResponseBody) GetInput() *QueryIProductionJobResponseBodyInput {
	return s.Input
}

func (s *QueryIProductionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *QueryIProductionJobResponseBody) GetJobParams() *string {
	return s.JobParams
}

func (s *QueryIProductionJobResponseBody) GetName() *string {
	return s.Name
}

func (s *QueryIProductionJobResponseBody) GetOutput() *QueryIProductionJobResponseBodyOutput {
	return s.Output
}

func (s *QueryIProductionJobResponseBody) GetOutputFiles() []*string {
	return s.OutputFiles
}

func (s *QueryIProductionJobResponseBody) GetOutputMediaIds() []*string {
	return s.OutputMediaIds
}

func (s *QueryIProductionJobResponseBody) GetOutputUrls() []*string {
	return s.OutputUrls
}

func (s *QueryIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIProductionJobResponseBody) GetResult() *string {
	return s.Result
}

func (s *QueryIProductionJobResponseBody) GetScheduleConfig() *QueryIProductionJobResponseBodyScheduleConfig {
	return s.ScheduleConfig
}

func (s *QueryIProductionJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryIProductionJobResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryIProductionJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *QueryIProductionJobResponseBody) SetCreateTime(v string) *QueryIProductionJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetFinishTime(v string) *QueryIProductionJobResponseBody {
	s.FinishTime = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetFunctionName(v string) *QueryIProductionJobResponseBody {
	s.FunctionName = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetInput(v *QueryIProductionJobResponseBodyInput) *QueryIProductionJobResponseBody {
	s.Input = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetJobId(v string) *QueryIProductionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetJobParams(v string) *QueryIProductionJobResponseBody {
	s.JobParams = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetName(v string) *QueryIProductionJobResponseBody {
	s.Name = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetOutput(v *QueryIProductionJobResponseBodyOutput) *QueryIProductionJobResponseBody {
	s.Output = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetOutputFiles(v []*string) *QueryIProductionJobResponseBody {
	s.OutputFiles = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetOutputMediaIds(v []*string) *QueryIProductionJobResponseBody {
	s.OutputMediaIds = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetOutputUrls(v []*string) *QueryIProductionJobResponseBody {
	s.OutputUrls = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetRequestId(v string) *QueryIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetResult(v string) *QueryIProductionJobResponseBody {
	s.Result = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetScheduleConfig(v *QueryIProductionJobResponseBodyScheduleConfig) *QueryIProductionJobResponseBody {
	s.ScheduleConfig = v
	return s
}

func (s *QueryIProductionJobResponseBody) SetStatus(v string) *QueryIProductionJobResponseBody {
	s.Status = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetTemplateId(v string) *QueryIProductionJobResponseBody {
	s.TemplateId = &v
	return s
}

func (s *QueryIProductionJobResponseBody) SetUserData(v string) *QueryIProductionJobResponseBody {
	s.UserData = &v
	return s
}

func (s *QueryIProductionJobResponseBody) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryIProductionJobResponseBodyInput struct {
	// The source file for the job. Set this to an OSS file URL if `Type` is `OSS`, or a media asset ID if `Type` is `Media`.
	//
	// Valid OSS URL formats:
	//
	// 1. oss\\://bucket/object
	//
	// 2. http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	//    In these formats, `bucket` is the name of the OSS bucket in the same region as the current project, and `object` is the file path.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Valid values:
	//
	// 1. OSS: An OSS file URL.
	//
	// 2. Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryIProductionJobResponseBodyInput) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponseBodyInput) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponseBodyInput) GetMedia() *string {
	return s.Media
}

func (s *QueryIProductionJobResponseBodyInput) GetType() *string {
	return s.Type
}

func (s *QueryIProductionJobResponseBodyInput) SetMedia(v string) *QueryIProductionJobResponseBodyInput {
	s.Media = &v
	return s
}

func (s *QueryIProductionJobResponseBodyInput) SetType(v string) *QueryIProductionJobResponseBodyInput {
	s.Type = &v
	return s
}

func (s *QueryIProductionJobResponseBodyInput) Validate() error {
	return dara.Validate(s)
}

type QueryIProductionJobResponseBodyOutput struct {
	// The service that the media asset belongs to.
	//
	// example:
	//
	// IMS
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The destination for the output. If the output `Type` is `OSS`, this parameter returns an OSS file URL. If the output `Type` is `Media`, it returns the specified or a newly generated media asset ID.
	//
	// Valid OSS URL formats:
	//
	// 1. oss\\://bucket/object
	//
	// 2. http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	//    In these formats, `bucket` is the name of the OSS bucket in the same region as the current project, and `object` is the file path.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The OSS URL of the output file. This value is returned only when `Type` is `Media`.
	//
	// example:
	//
	// http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media type. Valid values:
	//
	// - OSS: An OSS file URL.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryIProductionJobResponseBodyOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponseBodyOutput) GetBiz() *string {
	return s.Biz
}

func (s *QueryIProductionJobResponseBodyOutput) GetMedia() *string {
	return s.Media
}

func (s *QueryIProductionJobResponseBodyOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *QueryIProductionJobResponseBodyOutput) GetType() *string {
	return s.Type
}

func (s *QueryIProductionJobResponseBodyOutput) SetBiz(v string) *QueryIProductionJobResponseBodyOutput {
	s.Biz = &v
	return s
}

func (s *QueryIProductionJobResponseBodyOutput) SetMedia(v string) *QueryIProductionJobResponseBodyOutput {
	s.Media = &v
	return s
}

func (s *QueryIProductionJobResponseBodyOutput) SetOutputUrl(v string) *QueryIProductionJobResponseBodyOutput {
	s.OutputUrl = &v
	return s
}

func (s *QueryIProductionJobResponseBodyOutput) SetType(v string) *QueryIProductionJobResponseBodyOutput {
	s.Type = &v
	return s
}

func (s *QueryIProductionJobResponseBodyOutput) Validate() error {
	return dara.Validate(s)
}

type QueryIProductionJobResponseBodyScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// a54fdc9c9aab413caef0d1150f565e86
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job\\"s priority within the pipeline.
	//
	// - A larger value indicates a higher priority. The highest priority is 10.
	//
	// - Default: **6**.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s QueryIProductionJobResponseBodyScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryIProductionJobResponseBodyScheduleConfig) GoString() string {
	return s.String()
}

func (s *QueryIProductionJobResponseBodyScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryIProductionJobResponseBodyScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryIProductionJobResponseBodyScheduleConfig) SetPipelineId(v string) *QueryIProductionJobResponseBodyScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *QueryIProductionJobResponseBodyScheduleConfig) SetPriority(v int32) *QueryIProductionJobResponseBodyScheduleConfig {
	s.Priority = &v
	return s
}

func (s *QueryIProductionJobResponseBodyScheduleConfig) Validate() error {
	return dara.Validate(s)
}
