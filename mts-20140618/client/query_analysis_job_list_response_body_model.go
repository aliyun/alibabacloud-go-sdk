// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnalysisJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisJobList(v *QueryAnalysisJobListResponseBodyAnalysisJobList) *QueryAnalysisJobListResponseBody
	GetAnalysisJobList() *QueryAnalysisJobListResponseBodyAnalysisJobList
	SetNonExistAnalysisJobIds(v *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) *QueryAnalysisJobListResponseBody
	GetNonExistAnalysisJobIds() *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds
	SetRequestId(v string) *QueryAnalysisJobListResponseBody
	GetRequestId() *string
}

type QueryAnalysisJobListResponseBody struct {
	// The IDs of template analysis jobs.
	AnalysisJobList *QueryAnalysisJobListResponseBodyAnalysisJobList `json:"AnalysisJobList,omitempty" xml:"AnalysisJobList,omitempty" type:"Struct"`
	// The message sent by Message Service (MNS) to notify the user of the job result.
	NonExistAnalysisJobIds *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds `json:"NonExistAnalysisJobIds,omitempty" xml:"NonExistAnalysisJobIds,omitempty" type:"Struct"`
	// The status of the job. Valid values:
	//
	// 	- **Submitted**: The job has been submitted.
	//
	// 	- **Analyzing**: The job is being run.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// 5CA6E020-4102-4FFF-AA56-5ED7ECD811A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAnalysisJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBody) GetAnalysisJobList() *QueryAnalysisJobListResponseBodyAnalysisJobList {
	return s.AnalysisJobList
}

func (s *QueryAnalysisJobListResponseBody) GetNonExistAnalysisJobIds() *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds {
	return s.NonExistAnalysisJobIds
}

func (s *QueryAnalysisJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAnalysisJobListResponseBody) SetAnalysisJobList(v *QueryAnalysisJobListResponseBodyAnalysisJobList) *QueryAnalysisJobListResponseBody {
	s.AnalysisJobList = v
	return s
}

func (s *QueryAnalysisJobListResponseBody) SetNonExistAnalysisJobIds(v *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) *QueryAnalysisJobListResponseBody {
	s.NonExistAnalysisJobIds = v
	return s
}

func (s *QueryAnalysisJobListResponseBody) SetRequestId(v string) *QueryAnalysisJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAnalysisJobListResponseBody) Validate() error {
	if s.AnalysisJobList != nil {
		if err := s.AnalysisJobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistAnalysisJobIds != nil {
		if err := s.NonExistAnalysisJobIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAnalysisJobListResponseBodyAnalysisJobList struct {
	AnalysisJob []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob `json:"AnalysisJob,omitempty" xml:"AnalysisJob,omitempty" type:"Repeated"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobList) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobList) GetAnalysisJob() []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	return s.AnalysisJob
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobList) SetAnalysisJob(v []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) *QueryAnalysisJobListResponseBodyAnalysisJobList {
	s.AnalysisJob = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobList) Validate() error {
	if s.AnalysisJob != nil {
		for _, item := range s.AnalysisJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob struct {
	// The job configurations.
	AnalysisConfig *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig `json:"AnalysisConfig,omitempty" xml:"AnalysisConfig,omitempty" type:"Struct"`
	// The error code returned if the job fails.
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
	// 57f6aa3f84824309bcba67231b406****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the job input.
	InputFile *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The message sent by Message Service (MNS) to notify users of the job result.
	MNSMessageResult *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
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
	// 86
	Percent *int64 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// bb558c1cc25b45309aab5be44d19****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job in the ApsaraVideo Media Processing (MPS) queue to which the job was submitted.
	//
	// 	- Valid values: **1 to 10**. The value 10 indicates the highest priority.
	//
	// 	- Default value: **10**.
	//
	// example:
	//
	// 8
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The job status.
	//
	// 	- **Submitted**: The job was submitted.
	//
	// 	- **Analyzing**: The job is in progress.
	//
	// 	- **Success**: The job was successful.
	//
	// 	- **Fail**: The job failed.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The matched preset templates.
	TemplateList *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Struct"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetAnalysisConfig() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig {
	return s.AnalysisConfig
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetCode() *string {
	return s.Code
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetId() *string {
	return s.Id
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetInputFile() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile {
	return s.InputFile
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetMNSMessageResult() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetMessage() *string {
	return s.Message
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetPercent() *int64 {
	return s.Percent
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetPriority() *string {
	return s.Priority
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetState() *string {
	return s.State
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetTemplateList() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList {
	return s.TemplateList
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetAnalysisConfig(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.AnalysisConfig = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetCode(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.Code = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetCreationTime(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.CreationTime = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetId(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.Id = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetInputFile(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.InputFile = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetMNSMessageResult(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.MNSMessageResult = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetMessage(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.Message = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetPercent(v int64) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.Percent = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetPipelineId(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.PipelineId = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetPriority(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.Priority = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetState(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.State = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetTemplateList(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.TemplateList = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) SetUserData(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob {
	s.UserData = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJob) Validate() error {
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

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig struct {
	// The control on the attributes of the job output.
	PropertiesControl *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl `json:"PropertiesControl,omitempty" xml:"PropertiesControl,omitempty" type:"Struct"`
	// The quality control on the job output.
	QualityControl *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl `json:"QualityControl,omitempty" xml:"QualityControl,omitempty" type:"Struct"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) GetPropertiesControl() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl {
	return s.PropertiesControl
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) GetQualityControl() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl {
	return s.QualityControl
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) SetPropertiesControl(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig {
	s.PropertiesControl = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) SetQualityControl(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig {
	s.QualityControl = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfig) Validate() error {
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

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl struct {
	// The cropping configuration of the video image.
	Crop *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop `json:"Crop,omitempty" xml:"Crop,omitempty" type:"Struct"`
	// Specifies whether deinterlacing was forced to run. Valid values:
	//
	// 	- **Auto**: Deinterlacing was automatically run.
	//
	// 	- **Force**: Deinterlacing was forced to run.
	//
	// 	- **None**: Deinterlacing was forced not to run.
	//
	// example:
	//
	// Auto
	Deinterlace *string `json:"Deinterlace,omitempty" xml:"Deinterlace,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) GetCrop() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	return s.Crop
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) GetDeinterlace() *string {
	return s.Deinterlace
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) SetCrop(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl {
	s.Crop = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) SetDeinterlace(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl {
	s.Deinterlace = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControl) Validate() error {
	if s.Crop != nil {
		if err := s.Crop.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop struct {
	// The height of the video image after the margins are cropped out.
	//
	// >  This parameter is invalid if the Mode parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left margin to be cropped out.
	//
	// >  This parameter is invalid if the Mode parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The cropping mode. Valid values:
	//
	// 	- **Auto**: Cropping was automatically run.
	//
	// 	- **Force**: Cropping was forced to run.
	//
	// 	- **None**: Cropping was forced not to run.
	//
	// 	- This parameter is required if the value of the Crop parameter is not an empty JSON {}.
	//
	// example:
	//
	// Auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The top margin to be cropped out.
	//
	// >  This parameter is invalid if the Mode parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width of the video image after the margins are cropped out.
	//
	// >  This parameter is invalid if the Mode parameter is set to Auto or None.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GetHeight() *string {
	return s.Height
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GetLeft() *string {
	return s.Left
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GetMode() *string {
	return s.Mode
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GetTop() *string {
	return s.Top
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) GetWidth() *string {
	return s.Width
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) SetHeight(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Height = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) SetLeft(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Left = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) SetMode(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Mode = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) SetTop(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Top = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) SetWidth(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop {
	s.Width = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigPropertiesControlCrop) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl struct {
	// The playback mode. Valid values:
	//
	// 	- **network**: online playback.
	//
	// 	- **local**: playback on on-premises devices.
	//
	// 	- Default value: **network**.
	//
	// example:
	//
	// network
	MethodStreaming *string `json:"MethodStreaming,omitempty" xml:"MethodStreaming,omitempty"`
	// The quality level of the job output. Default value: **25**.
	//
	// example:
	//
	// 25
	RateQuality *string `json:"RateQuality,omitempty" xml:"RateQuality,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) GetMethodStreaming() *string {
	return s.MethodStreaming
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) GetRateQuality() *string {
	return s.RateQuality
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) SetMethodStreaming(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl {
	s.MethodStreaming = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) SetRateQuality(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl {
	s.RateQuality = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobAnalysisConfigQualityControl) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The region in which OSS resides.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the Object Storage Service (OSS) object that is used as the input file.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) GetObject() *string {
	return s.Object
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) SetBucket(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile {
	s.Bucket = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) SetLocation(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile {
	s.Location = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) SetObject(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile {
	s.Object = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobInputFile) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult struct {
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

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) SetErrorCode(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) SetErrorMessage(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) SetMessageId(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList struct {
	Template []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) GetTemplate() []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	return s.Template
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) SetTemplate(v []*QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList {
	s.Template = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateList) Validate() error {
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

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate struct {
	// The audio codec configurations.
	Audio *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format configurations.
	Container *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The transcoding template ID.
	//
	// example:
	//
	// S00000000-00****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transmuxing configurations.
	MuxConfig *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The name of the template.
	//
	// example:
	//
	// FLV-UD
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the template. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The general transcoding configurations.
	TransConfig *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video codec configurations.
	Video *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetAudio() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	return s.Audio
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetContainer() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer {
	return s.Container
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetId() *string {
	return s.Id
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetMuxConfig() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig {
	return s.MuxConfig
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetName() *string {
	return s.Name
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetState() *string {
	return s.State
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetTransConfig() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig {
	return s.TransConfig
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) GetVideo() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	return s.Video
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetAudio(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.Audio = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetContainer(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.Container = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetId(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.Id = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetMuxConfig(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.MuxConfig = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetName(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.Name = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetState(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.State = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetTransConfig(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.TransConfig = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) SetVideo(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate {
	s.Video = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplate) Validate() error {
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

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Valid values: **8 to 1000**.
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: **128**.
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Default value: **2**.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec. Default value: **aac**. Valid values:
	//
	// 	- **aac**
	//
	// 	- **mp3**
	//
	// 	- **vorbis**
	//
	// 	- **flac**
	//
	// example:
	//
	// aac
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The codec profile of the audio. Valid values when the **Codec*	- parameter is set to **aac**:
	//
	// 	- **aac_low**
	//
	// 	- **aac_he**
	//
	// 	- **aac_he_v2**
	//
	// 	- **aac_ld**
	//
	// 	- **aac_eld**
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The level of quality control on the audio.
	//
	// example:
	//
	// 15
	Qscale *string `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// The sampling rate.
	//
	// 	- Unit: Hz.
	//
	// 	- Default value: **44100**.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetQscale() *string {
	return s.Qscale
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetBitrate(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetChannels(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Channels = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetCodec(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Codec = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetProfile(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Profile = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetQscale(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Qscale = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) SetSamplerate(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateAudio) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer struct {
	// The container format.
	//
	// example:
	//
	// flv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) SetFormat(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer {
	s.Format = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig struct {
	// The transmuxing configurations for the GIF format.
	Gif *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif `json:"Gif,omitempty" xml:"Gif,omitempty" type:"Struct"`
	// The segment configurations.
	Segment *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) GetGif() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif {
	return s.Gif
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) GetSegment() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment {
	return s.Segment
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) SetGif(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig {
	s.Gif = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) SetSegment(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfig) Validate() error {
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

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif struct {
	// The interval between two consecutive loops for the GIF format. Unit: 0.01 second. For example, a value of 500 indicates 5 seconds.
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

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) GetFinalDelay() *string {
	return s.FinalDelay
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) GetLoop() *string {
	return s.Loop
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) SetFinalDelay(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif {
	s.FinalDelay = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) SetLoop(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif {
	s.Loop = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigGif) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment struct {
	// The segment length. Unit: seconds.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) SetDuration(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig struct {
	// The transcoding mode. Default value: **onepass**. Valid values:
	//
	// 	- **onepass**
	//
	// 	- **twopass**
	//
	// 	- **CBR**
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) SetTransMode(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo struct {
	// The average bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 1000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The average bitrate range of the video.
	BitrateBnd *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd `json:"BitrateBnd,omitempty" xml:"BitrateBnd,omitempty" type:"Struct"`
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
	// The codec. Valid values: H.264 and H.265. Default value: **H.264**.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Default value when the the Codec parameter is set to H.264: **23**, default value when the the Codec parameter is set to H.265: **26**.
	//
	// 	- If this parameter is specified, the value of the Bitrate parameter becomes invalid.
	//
	// example:
	//
	// 26
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The strength of the independent noise reduction algorithm.
	//
	// example:
	//
	// 5
	Degrain *string `json:"Degrain,omitempty" xml:"Degrain,omitempty"`
	// The frame rate of the video.
	//
	// 	- The value is 60 if the frame rate of the input file exceeds 60.
	//
	// 	- Default value: **the frame rate of the input video**.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between two keyframes. Default value: **250**.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: **the height of the input video**.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The maximum bitrate of the video.
	//
	// 	- Valid values: **10 to 50000**.
	//
	// 	- Unit: Kbit/s.
	//
	// example:
	//
	// 2000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The pixel format of the video. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
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
	// high
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
	// interlaced
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the video.
	//
	// 	- Unit: pixel.
	//
	// 	- Default value: **the width of the input video**.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetBitrateBnd() *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd {
	return s.BitrateBnd
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetDegrain() *string {
	return s.Degrain
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetQscale() *string {
	return s.Qscale
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetBitrate(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetBitrateBnd(v *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.BitrateBnd = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetBufsize(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetCodec(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Codec = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetCrf(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Crf = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetDegrain(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Degrain = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetFps(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Fps = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetGop(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Gop = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetHeight(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Height = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetMaxrate(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetPixFmt(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetPreset(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Preset = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetProfile(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Profile = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetQscale(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Qscale = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetScanMode(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) SetWidth(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo {
	s.Width = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideo) Validate() error {
	if s.BitrateBnd != nil {
		if err := s.BitrateBnd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd struct {
	// The upper limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 1500
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The lower limit of the total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 800
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) GetMax() *string {
	return s.Max
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) GetMin() *string {
	return s.Min
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) SetMax(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd {
	s.Max = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) SetMin(v string) *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd {
	s.Min = &v
	return s
}

func (s *QueryAnalysisJobListResponseBodyAnalysisJobListAnalysisJobTemplateListTemplateVideoBitrateBnd) Validate() error {
	return dara.Validate(s)
}

type QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) String() string {
	return dara.Prettify(s)
}

func (s QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) GoString() string {
	return s.String()
}

func (s *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) GetString_() []*string {
	return s.String_
}

func (s *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) SetString_(v []*string) *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds {
	s.String_ = v
	return s
}

func (s *QueryAnalysisJobListResponseBodyNonExistAnalysisJobIds) Validate() error {
	return dara.Validate(s)
}
