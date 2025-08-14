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
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-07T07:16:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2021-11-26T14:50:25Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The name of the algorithm that you want to use for the job. Valid values:
	//
	// 	- **Cover**: This algorithm intelligently generates a thumbnail image for a video.
	//
	// 	- **VideoClip**: This algorithm intelligently generates a summary for a video.
	//
	// 	- **VideoDelogo**: This algorithm removes logos from a video.
	//
	// 	- **VideoDetext**: This algorithm removes captions from a video.
	//
	// example:
	//
	// Cover
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The input file.
	Input *QueryIProductionJobResponseBodyInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the intelligent production job.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The algorithm-specific parameters. The parameters are specified as JSON objects and vary based on the algorithm.
	//
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The name of the intelligent production job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output file.
	Output *QueryIProductionJobResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The output files.
	OutputFiles    []*string `json:"OutputFiles,omitempty" xml:"OutputFiles,omitempty" type:"Repeated"`
	OutputMediaIds []*string `json:"OutputMediaIds,omitempty" xml:"OutputMediaIds,omitempty" type:"Repeated"`
	// The URLs of the output files.
	OutputUrls []*string `json:"OutputUrls,omitempty" xml:"OutputUrls,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The output of the algorithm. The output is in JSON format and varies based on the algorithm. For more information, see the "Parameters of Result" section of this topic.
	//
	// example:
	//
	// {}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The scheduling configuration.
	ScheduleConfig *QueryIProductionJobResponseBodyScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The status of the job. Valid values:
	//
	// 	- Queuing: The job is waiting in the queue.
	//
	// 	- Analysing: The job is in progress.
	//
	// 	- Fail: The job failed.
	//
	// 	- Success: The job was successful.
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
	// The user-defined data that is returned in the response.
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
	return dara.Validate(s)
}

type QueryIProductionJobResponseBodyInput struct {
	// The input file. If Type is set to OSS, set this parameter to the path of an OSS object. If Type is set to Media, set this parameter to the ID of a media asset. You can specify the path of an OSS object in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object bucket in the path specifies an OSS bucket that resides in the same region as the intelligent production job. object in the path specifies the object path in OSS.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media type. Valid values:
	//
	// 1.  OSS: Object Storage Service (OSS) object
	//
	// 2.  Media: media asset
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
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The output file. If Type is set to OSS, set this parameter to the path of an OSS object. If Type is set to Media, set this parameter to the ID of a media asset. You can specify the path of an OSS object in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object bucket in the path specifies an OSS bucket that resides in the same region as the intelligent production job. object in the path specifies the object path in OSS.
	//
	// example:
	//
	// oss://bucket/object
	Media     *string `json:"Media,omitempty" xml:"Media,omitempty"`
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media type. Valid values:
	//
	// 	- OSS: OSS object
	//
	// 	- Media: media asset
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
	// The ID of the ApsaraVideo Media Processing (MPS) queue.
	//
	// example:
	//
	// a54fdc9c9aab413caef0d1150f565e86
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job in the MPS queue to which the job is added.
	//
	// 	- A value of 10 indicates the highest priority.
	//
	// 	- Default value: **6**.
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
