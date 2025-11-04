// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartHandleJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetSmartHandleJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSmartHandleJobResponseBody
	GetErrorMessage() *string
	SetJobId(v string) *GetSmartHandleJobResponseBody
	GetJobId() *string
	SetJobResult(v *GetSmartHandleJobResponseBodyJobResult) *GetSmartHandleJobResponseBody
	GetJobResult() *GetSmartHandleJobResponseBodyJobResult
	SetOutput(v string) *GetSmartHandleJobResponseBody
	GetOutput() *string
	SetRequestId(v string) *GetSmartHandleJobResponseBody
	GetRequestId() *string
	SetSmartJobInfo(v *GetSmartHandleJobResponseBodySmartJobInfo) *GetSmartHandleJobResponseBody
	GetSmartJobInfo() *GetSmartHandleJobResponseBodySmartJobInfo
	SetState(v string) *GetSmartHandleJobResponseBody
	GetState() *string
	SetUserData(v string) *GetSmartHandleJobResponseBody
	GetUserData() *string
}

type GetSmartHandleJobResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job results.
	JobResult *GetSmartHandleJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	// The job results.
	//
	// example:
	//
	// {}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the intelligent job.
	SmartJobInfo *GetSmartHandleJobResponseBodySmartJobInfo `json:"SmartJobInfo,omitempty" xml:"SmartJobInfo,omitempty" type:"Struct"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Finished
	//
	// 	- Failed
	//
	// 	- Executing
	//
	// 	- Created
	//
	// example:
	//
	// Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The user-defined data in the JSON format.
	//
	// example:
	//
	// {"user":"data"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetSmartHandleJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSmartHandleJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSmartHandleJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetSmartHandleJobResponseBody) GetJobResult() *GetSmartHandleJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetSmartHandleJobResponseBody) GetOutput() *string {
	return s.Output
}

func (s *GetSmartHandleJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmartHandleJobResponseBody) GetSmartJobInfo() *GetSmartHandleJobResponseBodySmartJobInfo {
	return s.SmartJobInfo
}

func (s *GetSmartHandleJobResponseBody) GetState() *string {
	return s.State
}

func (s *GetSmartHandleJobResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetSmartHandleJobResponseBody) SetErrorCode(v string) *GetSmartHandleJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetErrorMessage(v string) *GetSmartHandleJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetJobId(v string) *GetSmartHandleJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetJobResult(v *GetSmartHandleJobResponseBodyJobResult) *GetSmartHandleJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetOutput(v string) *GetSmartHandleJobResponseBody {
	s.Output = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetRequestId(v string) *GetSmartHandleJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetSmartJobInfo(v *GetSmartHandleJobResponseBodySmartJobInfo) *GetSmartHandleJobResponseBody {
	s.SmartJobInfo = v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetState(v string) *GetSmartHandleJobResponseBody {
	s.State = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetUserData(v string) *GetSmartHandleJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) Validate() error {
	if s.JobResult != nil {
		if err := s.JobResult.Validate(); err != nil {
			return err
		}
	}
	if s.SmartJobInfo != nil {
		if err := s.SmartJobInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmartHandleJobResponseBodyJobResult struct {
	// The AI analysis result.
	//
	// example:
	//
	// Intelligent segmentation or tagging information
	AiResult *string `json:"AiResult,omitempty" xml:"AiResult,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId  *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// The token usage. This parameter is returned only for keyword-based text generation jobs.
	//
	// example:
	//
	// {"total_tokens":100}
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetSmartHandleJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodyJobResult) GetAiResult() *string {
	return s.AiResult
}

func (s *GetSmartHandleJobResponseBodyJobResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetSmartHandleJobResponseBodyJobResult) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *GetSmartHandleJobResponseBodyJobResult) GetUsage() *string {
	return s.Usage
}

func (s *GetSmartHandleJobResponseBodyJobResult) SetAiResult(v string) *GetSmartHandleJobResponseBodyJobResult {
	s.AiResult = &v
	return s
}

func (s *GetSmartHandleJobResponseBodyJobResult) SetMediaId(v string) *GetSmartHandleJobResponseBodyJobResult {
	s.MediaId = &v
	return s
}

func (s *GetSmartHandleJobResponseBodyJobResult) SetMediaUrl(v string) *GetSmartHandleJobResponseBodyJobResult {
	s.MediaUrl = &v
	return s
}

func (s *GetSmartHandleJobResponseBodyJobResult) SetUsage(v string) *GetSmartHandleJobResponseBodyJobResult {
	s.Usage = &v
	return s
}

func (s *GetSmartHandleJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}

type GetSmartHandleJobResponseBodySmartJobInfo struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job description.
	//
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input configurations.
	InputConfig *GetSmartHandleJobResponseBodySmartJobInfoInputConfig `json:"InputConfig,omitempty" xml:"InputConfig,omitempty" type:"Struct"`
	// The job type.
	//
	// example:
	//
	// ASR
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The output configurations.
	OutputConfig *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty" type:"Struct"`
	// The job title.
	//
	// example:
	//
	// 测试标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1974526429******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSmartHandleJobResponseBodySmartJobInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponseBodySmartJobInfo) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetDescription() *string {
	return s.Description
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetInputConfig() *GetSmartHandleJobResponseBodySmartJobInfoInputConfig {
	return s.InputConfig
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetJobType() *string {
	return s.JobType
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetOutputConfig() *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig {
	return s.OutputConfig
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetTitle() *string {
	return s.Title
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) GetUserId() *string {
	return s.UserId
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetCreateTime(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetDescription(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.Description = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetInputConfig(v *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.InputConfig = v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetJobType(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.JobType = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetModifiedTime(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetOutputConfig(v *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.OutputConfig = v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetTitle(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.Title = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetUserId(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.UserId = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) Validate() error {
	if s.InputConfig != nil {
		if err := s.InputConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OutputConfig != nil {
		if err := s.OutputConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmartHandleJobResponseBodySmartJobInfoInputConfig struct {
	// The OSS URL or the ID of the material in the media asset library.
	//
	// example:
	//
	// oss://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4 或 ******11-DB8D-4A9A-875B-275798******
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
}

func (s GetSmartHandleJobResponseBodySmartJobInfoInputConfig) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponseBodySmartJobInfoInputConfig) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) GetInputFile() *string {
	return s.InputFile
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) SetInputFile(v string) *GetSmartHandleJobResponseBodySmartJobInfoInputConfig {
	s.InputFile = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) Validate() error {
	return dara.Validate(s)
}

type GetSmartHandleJobResponseBodySmartJobInfoOutputConfig struct {
	// The OSS bucket.
	//
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS object.
	//
	// example:
	//
	// test-object
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) String() string {
	return dara.Prettify(s)
}

func (s GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) GetBucket() *string {
	return s.Bucket
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) GetObject() *string {
	return s.Object
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) SetBucket(v string) *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig {
	s.Bucket = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) SetObject(v string) *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig {
	s.Object = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoOutputConfig) Validate() error {
	return dara.Validate(s)
}
