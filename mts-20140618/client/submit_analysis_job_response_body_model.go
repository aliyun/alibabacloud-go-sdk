// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisJob(v *SubmitAnalysisJobResponseBodyAnalysisJob) *SubmitAnalysisJobResponseBody
	GetAnalysisJob() *SubmitAnalysisJobResponseBodyAnalysisJob
	SetRequestId(v string) *SubmitAnalysisJobResponseBody
	GetRequestId() *string
}

type SubmitAnalysisJobResponseBody struct {
	// The information about the preset template analysis job that was submitted.
	AnalysisJob *SubmitAnalysisJobResponseBodyAnalysisJob `json:"AnalysisJob,omitempty" xml:"AnalysisJob,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B52658D4-07AB-43CD-82B0-210958A65E23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAnalysisJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBody) GetAnalysisJob() *SubmitAnalysisJobResponseBodyAnalysisJob {
	return s.AnalysisJob
}

func (s *SubmitAnalysisJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAnalysisJobResponseBody) SetAnalysisJob(v *SubmitAnalysisJobResponseBodyAnalysisJob) *SubmitAnalysisJobResponseBody {
	s.AnalysisJob = v
	return s
}

func (s *SubmitAnalysisJobResponseBody) SetRequestId(v string) *SubmitAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAnalysisJobResponseBody) Validate() error {
	if s.AnalysisJob != nil {
		if err := s.AnalysisJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJob struct {
	// The job configurations.
	AnalysisConfig *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig `json:"AnalysisConfig,omitempty" xml:"AnalysisConfig,omitempty" type:"Struct"`
	// The error code returned if the job failed. This parameter is not returned if the job was successful.
	//
	// example:
	//
	// InvalidParameter.ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2014-01-10T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the template analysis job.
	//
	// example:
	//
	// 57f6aa3f84824309bcba67231b40****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the job input.
	InputFile *SubmitAnalysisJobResponseBodyAnalysisJobInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The message sent by MNS to notify users of the job result.
	MNSMessageResult *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the job failed.
	//
	// example:
	//
	// The resource operated \\"PipelineId\\" cannot be found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The transcoding progress.
	//
	// example:
	//
	// 100
	Percent *int64 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the MPS queue to which the analysis job was submitted.
	//
	// example:
	//
	// bb558c1cc25b45309aab5be44d19****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job in the MPS queue to which the job was submitted.
	//
	// 	- Valid values: **1 to 10**. A value of 10 indicates the highest priority.
	//
	// 	- Default value: **10**.
	//
	// example:
	//
	// 10
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **Submitted**: The job is submitted.
	//
	// 	- **Analyzing**: The job is being run.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The matched preset templates.
	TemplateList *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Struct"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJob) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetAnalysisConfig() *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig {
	return s.AnalysisConfig
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetCode() *string {
	return s.Code
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetId() *string {
	return s.Id
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetInputFile() *SubmitAnalysisJobResponseBodyAnalysisJobInputFile {
	return s.InputFile
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetMNSMessageResult() *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetPercent() *int64 {
	return s.Percent
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetPriority() *string {
	return s.Priority
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetState() *string {
	return s.State
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetTemplateList() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList {
	return s.TemplateList
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetAnalysisConfig(v *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.AnalysisConfig = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetCode(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.Code = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetCreationTime(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetId(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.Id = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetInputFile(v *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.InputFile = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetMNSMessageResult(v *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.MNSMessageResult = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetMessage(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.Message = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetPercent(v int64) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.Percent = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetPipelineId(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.PipelineId = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetPriority(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.Priority = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetState(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.State = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetTemplateList(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.TemplateList = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) SetUserData(v string) *SubmitAnalysisJobResponseBodyAnalysisJob {
	s.UserData = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJob) Validate() error {
	if s.AnalysisConfig != nil {
		if err := s.AnalysisConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateList != nil {
		if err := s.TemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig struct {
	// The control on the attributes of the job output.
	PropertiesControl *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl `json:"PropertiesControl,omitempty" xml:"PropertiesControl,omitempty" type:"Struct"`
	// The quality control on the job output.
	QualityControl *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl `json:"QualityControl,omitempty" xml:"QualityControl,omitempty" type:"Struct"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) GetPropertiesControl() *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl {
	return s.PropertiesControl
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) GetQualityControl() *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl {
	return s.QualityControl
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) SetPropertiesControl(v *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig {
	s.PropertiesControl = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) SetQualityControl(v *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig {
	s.QualityControl = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfig) Validate() error {
	if s.PropertiesControl != nil {
		if err := s.PropertiesControl.Validate(); err != nil {
			return err
		}
	}
	if s.QualityControl != nil {
		if err := s.QualityControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl struct {
	// The cropping configurations of video images.
	Crop *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop `json:"Crop,omitempty" xml:"Crop,omitempty" type:"Struct"`
	// Indicates whether deinterlacing was forced to run. Valid values:
	//
	// 	- **Auto**: Deinterlacing was automatically run.
	//
	// 	- **Force**: Deinterlacing was forced to run.
	//
	// 	- **None**: Deinterlacing was forced not to run.
	//
	// example:
	//
	// Force
	Deinterlace *string `json:"Deinterlace,omitempty" xml:"Deinterlace,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) GetCrop() *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	return s.Crop
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) GetDeinterlace() *string {
	return s.Deinterlace
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) SetCrop(v *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl {
	s.Crop = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) SetDeinterlace(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl {
	s.Deinterlace = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControl) Validate() error {
	if s.Crop != nil {
		if err := s.Crop.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop struct {
	// The height of the video after the margins were cropped out.
	//
	// > This parameter is invalid if the **Mode*	- parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left margin that was cropped out.
	//
	// > This parameter is invalid if the **Mode*	- parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The cropping mode. Valid values:
	//
	// 	- **Auto**: Cropping was automatically run. This is the default value.
	//
	// 	- **Force**: Cropping was forced to run.
	//
	// 	- **None**: Cropping was forced not to run.
	//
	// example:
	//
	// Auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The top margin that was cropped out.
	//
	// > This parameter is invalid if the **Mode*	- parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width of the video after the margins were cropped out.
	//
	// > This parameter is invalid if the **Mode*	- parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GetHeight() *string {
	return s.Height
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GetLeft() *string {
	return s.Left
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GetMode() *string {
	return s.Mode
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GetTop() *string {
	return s.Top
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) GetWidth() *string {
	return s.Width
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) SetHeight(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Height = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) SetLeft(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Left = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) SetMode(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Mode = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) SetTop(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Top = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) SetWidth(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Width = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigPropertiesControlCrop) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl struct {
	// The playback mode. Valid values:
	//
	// 	- **network**: online playback
	//
	// 	- **local**: playback on local devices
	//
	// 	- Default value: **network**.
	//
	// example:
	//
	// network
	MethodStreaming *string `json:"MethodStreaming,omitempty" xml:"MethodStreaming,omitempty"`
	// The quality level of the output file.
	//
	// example:
	//
	// 50
	RateQuality *string `json:"RateQuality,omitempty" xml:"RateQuality,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) GetMethodStreaming() *string {
	return s.MethodStreaming
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) GetRateQuality() *string {
	return s.RateQuality
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) SetMethodStreaming(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl {
	s.MethodStreaming = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) SetRateQuality(v string) *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl {
	s.RateQuality = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobAnalysisConfigQualityControl) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobInputFile struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the OSS region.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input file.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobInputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobInputFile) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) SetBucket(v string) *SubmitAnalysisJobResponseBodyAnalysisJobInputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) SetLocation(v string) *SubmitAnalysisJobResponseBodyAnalysisJobInputFile {
	s.Location = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) SetObject(v string) *SubmitAnalysisJobResponseBodyAnalysisJobInputFile {
	s.Object = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobInputFile) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult struct {
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
	// The resource operated \\"PipelineId\\" cannot be found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message returned if the job was successful. This parameter is not returned if the job failed.
	//
	// example:
	//
	// 3ca84a39a9024f19853b21be9cf9****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) SetErrorCode(v string) *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) SetErrorMessage(v string) *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) SetMessageId(v string) *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateList struct {
	Template []*SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) GetTemplate() []*SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	return s.Template
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) SetTemplate(v []*SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList {
	s.Template = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateList) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate struct {
	// The audio codec configurations.
	Audio *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format configurations.
	Container *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// S00000000-00****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transmuxing configurations.
	MuxConfig *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The name of the template.
	//
	// example:
	//
	// FLV-UD
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the template.
	//
	// 	- **Normal**: The template is normal.
	//
	// 	- **Deleted**: The template is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The general transcoding configurations.
	TransConfig *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetAudio() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	return s.Audio
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetContainer() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer {
	return s.Container
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetId() *string {
	return s.Id
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetMuxConfig() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig {
	return s.MuxConfig
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetName() *string {
	return s.Name
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetState() *string {
	return s.State
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetTransConfig() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig {
	return s.TransConfig
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) GetVideo() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	return s.Video
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetAudio(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.Audio = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetContainer(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.Container = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetId(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.Id = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetMuxConfig(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.MuxConfig = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetName(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.Name = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetState(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.State = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetTransConfig(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.TransConfig = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) SetVideo(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate {
	s.Video = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplate) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
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
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: **128**.
	//
	// example:
	//
	// 8
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Default value: **2**.
	//
	// example:
	//
	// 1
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec format. Default value: **acc**.
	//
	// example:
	//
	// mp3
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The codec profile of the audio. Valid values if the **Codec*	- parameter is set to **aac**: aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of quality control on the audio.
	//
	// example:
	//
	// 10
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The sampling rate.
	//
	// 	- Unit: Hz.
	//
	// 	- Default value: **44100**.
	//
	// example:
	//
	// 32000
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetBitrate(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetChannels(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Channels = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetCodec(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Codec = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetProfile(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Profile = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetQscale(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) SetSamplerate(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateAudio) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer struct {
	// The container format.
	//
	// example:
	//
	// flv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) SetFormat(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer {
	s.Format = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig struct {
	// The transmuxing configurations for the GIF format.
	Gif *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment configurations.
	Segment *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) GetGif() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif {
	return s.Gif
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) GetSegment() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment {
	return s.Segment
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) SetGif(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) SetSegment(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfig) Validate() error {
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
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif struct {
	// The interval between two consecutive loops for the GIF format. Unit: 0.01s. For example, a value of 500 indicates 5 seconds.
	//
	// example:
	//
	// 0
	FinalDelay *string `json:"FinalDelay,omitempty" xml:"FinalDelay,omitempty"`
	// The number of loops for the GIF or WebP format. Default value: 0.
	//
	// example:
	//
	// 0
	Loop *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) SetFinalDelay(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) SetLoop(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment struct {
	// The length of the segment. Unit: seconds.
	//
	// example:
	//
	// 60
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) SetDuration(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig struct {
	// The transcoding mode. Valid values: onepass, twopass, and CBR. Default value: **onepass**.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) SetTransMode(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo struct {
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 10
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The average bitrate range of the video.
	BitrateBnd *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
	// The size of the buffer.
	//
	// 	- Unit: KB.
	//
	// 	- Default value: **6000**.
	//
	// example:
	//
	// 5000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The video codec. Default value: **H.264**.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Default value if the Codec parameter is set to H.264: **23**. Default value if the Codec parameter is set to H.265: **26**.
	//
	// 	- If this parameter is returned, the setting of the Bitrate parameter is invalid.
	//
	// example:
	//
	// 27
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The strength of the independent noise reduction algorithm.
	//
	// example:
	//
	// 5
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	// The frame rate.
	//
	// 	- The value is 60 if the frame rate of the input video exceeds 60.
	//
	// 	- Default value: the frame rate of the input video.
	//
	// example:
	//
	// 60
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between two keyframes. Default value: **250**.
	//
	// example:
	//
	// 1
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: the height of the input video.
	//
	// example:
	//
	// 1880
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The maximum bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 10
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The pixel format for video color encoding. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuvj420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The preset video algorithm. Valid values: veryfast, fast, medium, slow, and slower. Default value: **medium**.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The codec profile. Valid values:
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
	// baseline
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of quality control on the video.
	//
	// example:
	//
	// 15
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- **interlaced**
	//
	// 	- **progressive**
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: the width of the input video.
	//
	// example:
	//
	// 1990
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetBitrateBnd() *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetBitrate(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetBitrateBnd(v *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetBufsize(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetCodec(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Codec = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetCrf(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Crf = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetDegrain(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetFps(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Fps = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetGop(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Gop = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetHeight(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Height = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetMaxrate(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetPixFmt(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetPreset(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Preset = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetProfile(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Profile = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetQscale(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetScanMode(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) SetWidth(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo {
	s.Width = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd struct {
	// The upper limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 20
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The lower limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 10
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) SetMax(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) SetMin(v string) *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *SubmitAnalysisJobResponseBodyAnalysisJobTemplateListTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}
