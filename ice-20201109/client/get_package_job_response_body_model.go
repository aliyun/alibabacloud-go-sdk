// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackageJob(v *GetPackageJobResponseBodyPackageJob) *GetPackageJobResponseBody
	GetPackageJob() *GetPackageJobResponseBodyPackageJob
	SetRequestId(v string) *GetPackageJobResponseBody
	GetRequestId() *string
}

type GetPackageJobResponseBody struct {
	// The information about the packaging job.
	PackageJob *GetPackageJobResponseBodyPackageJob `json:"PackageJob,omitempty" xml:"PackageJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPackageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponseBody) GetPackageJob() *GetPackageJobResponseBodyPackageJob {
	return s.PackageJob
}

func (s *GetPackageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPackageJobResponseBody) SetPackageJob(v *GetPackageJobResponseBodyPackageJob) *GetPackageJobResponseBody {
	s.PackageJob = v
	return s
}

func (s *GetPackageJobResponseBody) SetRequestId(v string) *GetPackageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackageJobResponseBody) Validate() error {
	if s.PackageJob != nil {
		if err := s.PackageJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPackageJobResponseBodyPackageJob struct {
	// The error code returned if the job fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-08T11:34:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-08T11:44:05Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input of the job.
	Inputs []*GetPackageJobResponseBodyPackageJobInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// Resource content bad.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the job was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-08T11:44:05Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	Output *GetPackageJobResponseBodyPackageJobOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The URL of the output file.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/output.m3u8
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// 36f3fee40aa047c0b067d0fb85edc12b
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid values: 1 to 10. The greater the value, the higher the priority.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The state of the job.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was submitted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-08T11:34:05Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The source of the job. Valid values:
	//
	// 	- API
	//
	// 	- WorkFlow
	//
	// 	- Console
	//
	// example:
	//
	// API
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// {"param": "value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetPackageJobResponseBodyPackageJob) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponseBodyPackageJob) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponseBodyPackageJob) GetCode() *string {
	return s.Code
}

func (s *GetPackageJobResponseBodyPackageJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPackageJobResponseBodyPackageJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetPackageJobResponseBodyPackageJob) GetInputs() []*GetPackageJobResponseBodyPackageJobInputs {
	return s.Inputs
}

func (s *GetPackageJobResponseBodyPackageJob) GetJobId() *string {
	return s.JobId
}

func (s *GetPackageJobResponseBodyPackageJob) GetMessage() *string {
	return s.Message
}

func (s *GetPackageJobResponseBodyPackageJob) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPackageJobResponseBodyPackageJob) GetName() *string {
	return s.Name
}

func (s *GetPackageJobResponseBodyPackageJob) GetOutput() *GetPackageJobResponseBodyPackageJobOutput {
	return s.Output
}

func (s *GetPackageJobResponseBodyPackageJob) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetPackageJobResponseBodyPackageJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetPackageJobResponseBodyPackageJob) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPackageJobResponseBodyPackageJob) GetStatus() *string {
	return s.Status
}

func (s *GetPackageJobResponseBodyPackageJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetPackageJobResponseBodyPackageJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetPackageJobResponseBodyPackageJob) GetUserData() *string {
	return s.UserData
}

func (s *GetPackageJobResponseBodyPackageJob) SetCode(v string) *GetPackageJobResponseBodyPackageJob {
	s.Code = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetCreateTime(v string) *GetPackageJobResponseBodyPackageJob {
	s.CreateTime = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetFinishTime(v string) *GetPackageJobResponseBodyPackageJob {
	s.FinishTime = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetInputs(v []*GetPackageJobResponseBodyPackageJobInputs) *GetPackageJobResponseBodyPackageJob {
	s.Inputs = v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetJobId(v string) *GetPackageJobResponseBodyPackageJob {
	s.JobId = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetMessage(v string) *GetPackageJobResponseBodyPackageJob {
	s.Message = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetModifiedTime(v string) *GetPackageJobResponseBodyPackageJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetName(v string) *GetPackageJobResponseBodyPackageJob {
	s.Name = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetOutput(v *GetPackageJobResponseBodyPackageJobOutput) *GetPackageJobResponseBodyPackageJob {
	s.Output = v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetOutputUrl(v string) *GetPackageJobResponseBodyPackageJob {
	s.OutputUrl = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetPipelineId(v string) *GetPackageJobResponseBodyPackageJob {
	s.PipelineId = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetPriority(v int32) *GetPackageJobResponseBodyPackageJob {
	s.Priority = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetStatus(v string) *GetPackageJobResponseBodyPackageJob {
	s.Status = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetSubmitTime(v string) *GetPackageJobResponseBodyPackageJob {
	s.SubmitTime = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetTriggerSource(v string) *GetPackageJobResponseBodyPackageJob {
	s.TriggerSource = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) SetUserData(v string) *GetPackageJobResponseBodyPackageJob {
	s.UserData = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJob) Validate() error {
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPackageJobResponseBodyPackageJobInputs struct {
	// The information about the input stream file.
	Input *GetPackageJobResponseBodyPackageJobInputsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s GetPackageJobResponseBodyPackageJobInputs) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponseBodyPackageJobInputs) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponseBodyPackageJobInputs) GetInput() *GetPackageJobResponseBodyPackageJobInputsInput {
	return s.Input
}

func (s *GetPackageJobResponseBodyPackageJobInputs) SetInput(v *GetPackageJobResponseBodyPackageJobInputsInput) *GetPackageJobResponseBodyPackageJobInputs {
	s.Input = v
	return s
}

func (s *GetPackageJobResponseBodyPackageJobInputs) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPackageJobResponseBodyPackageJobInputsInput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, the ID of a media asset is returned.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an Object Storage Service (OSS) object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPackageJobResponseBodyPackageJobInputsInput) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponseBodyPackageJobInputsInput) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponseBodyPackageJobInputsInput) GetMedia() *string {
	return s.Media
}

func (s *GetPackageJobResponseBodyPackageJobInputsInput) GetType() *string {
	return s.Type
}

func (s *GetPackageJobResponseBodyPackageJobInputsInput) SetMedia(v string) *GetPackageJobResponseBodyPackageJobInputsInput {
	s.Media = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJobInputsInput) SetType(v string) *GetPackageJobResponseBodyPackageJobInputsInput {
	s.Type = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJobInputsInput) Validate() error {
	return dara.Validate(s)
}

type GetPackageJobResponseBodyPackageJobOutput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, the ID of a media asset is returned.
	//
	// example:
	//
	// oss://bucket/path/to/video.m3u8
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPackageJobResponseBodyPackageJobOutput) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponseBodyPackageJobOutput) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponseBodyPackageJobOutput) GetMedia() *string {
	return s.Media
}

func (s *GetPackageJobResponseBodyPackageJobOutput) GetType() *string {
	return s.Type
}

func (s *GetPackageJobResponseBodyPackageJobOutput) SetMedia(v string) *GetPackageJobResponseBodyPackageJobOutput {
	s.Media = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJobOutput) SetType(v string) *GetPackageJobResponseBodyPackageJobOutput {
	s.Type = &v
	return s
}

func (s *GetPackageJobResponseBodyPackageJobOutput) Validate() error {
	return dara.Validate(s)
}
