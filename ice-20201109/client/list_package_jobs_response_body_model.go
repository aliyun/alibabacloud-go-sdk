// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPackageJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackageJobList(v *ListPackageJobsResponseBodyPackageJobList) *ListPackageJobsResponseBody
	GetPackageJobList() *ListPackageJobsResponseBodyPackageJobList
	SetRequestId(v string) *ListPackageJobsResponseBody
	GetRequestId() *string
}

type ListPackageJobsResponseBody struct {
	// The list of packaging jobs.
	PackageJobList *ListPackageJobsResponseBodyPackageJobList `json:"PackageJobList,omitempty" xml:"PackageJobList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPackageJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBody) GetPackageJobList() *ListPackageJobsResponseBodyPackageJobList {
	return s.PackageJobList
}

func (s *ListPackageJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPackageJobsResponseBody) SetPackageJobList(v *ListPackageJobsResponseBodyPackageJobList) *ListPackageJobsResponseBody {
	s.PackageJobList = v
	return s
}

func (s *ListPackageJobsResponseBody) SetRequestId(v string) *ListPackageJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPackageJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPackageJobsResponseBodyPackageJobList struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. The token of the next page is returned after you call this operation for the first time.
	//
	// example:
	//
	// 019daf5780f74831b0e1a767c9f1c178
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The list of packaging jobs.
	PackageJobs []*ListPackageJobsResponseBodyPackageJobListPackageJobs `json:"PackageJobs,omitempty" xml:"PackageJobs,omitempty" type:"Repeated"`
}

func (s ListPackageJobsResponseBodyPackageJobList) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBodyPackageJobList) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBodyPackageJobList) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListPackageJobsResponseBodyPackageJobList) GetPackageJobs() []*ListPackageJobsResponseBodyPackageJobListPackageJobs {
	return s.PackageJobs
}

func (s *ListPackageJobsResponseBodyPackageJobList) SetNextPageToken(v string) *ListPackageJobsResponseBodyPackageJobList {
	s.NextPageToken = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobList) SetPackageJobs(v []*ListPackageJobsResponseBodyPackageJobListPackageJobs) *ListPackageJobsResponseBodyPackageJobList {
	s.PackageJobs = v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobList) Validate() error {
	return dara.Validate(s)
}

type ListPackageJobsResponseBodyPackageJobListPackageJobs struct {
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
	// 2022-07-07T14:00:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-07-07T15:00:32Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input of the job.
	Inputs []*ListPackageJobsResponseBodyPackageJobListPackageJobsInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// 7b38a5d86f1e47838927b6e7ccb11cbe
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
	// 2022-07-07T15:00:32Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	Output *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue that is used to run the job.
	//
	// example:
	//
	// 5b40833e4c3e4d4e95a866abb9a42510
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid values: 1 to 10. The greater the value, the higher the priority. Default value: 6.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The state of the job.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was submitted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-07-07T14:00:32Z
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

func (s ListPackageJobsResponseBodyPackageJobListPackageJobs) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBodyPackageJobListPackageJobs) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetCode() *string {
	return s.Code
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetInputs() []*ListPackageJobsResponseBodyPackageJobListPackageJobsInputs {
	return s.Inputs
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetMessage() *string {
	return s.Message
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetName() *string {
	return s.Name
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetOutput() *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput {
	return s.Output
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetStatus() *string {
	return s.Status
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) GetUserData() *string {
	return s.UserData
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetCode(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Code = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetCreateTime(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.CreateTime = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetFinishTime(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.FinishTime = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetInputs(v []*ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Inputs = v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetJobId(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.JobId = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetMessage(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Message = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetModifiedTime(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.ModifiedTime = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetName(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Name = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetOutput(v *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Output = v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetPipelineId(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.PipelineId = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetPriority(v int32) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Priority = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetStatus(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.Status = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetSubmitTime(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetTriggerSource(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.TriggerSource = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) SetUserData(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobs {
	s.UserData = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobs) Validate() error {
	return dara.Validate(s)
}

type ListPackageJobsResponseBodyPackageJobListPackageJobsInputs struct {
	// The information about the input stream file.
	Input *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) GetInput() *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput {
	return s.Input
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) SetInput(v *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) *ListPackageJobsResponseBodyPackageJobListPackageJobsInputs {
	s.Input = v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputs) Validate() error {
	return dara.Validate(s)
}

type ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) GetMedia() *string {
	return s.Media
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) GetType() *string {
	return s.Type
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) SetMedia(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput {
	s.Media = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) SetType(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput {
	s.Type = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsInputsInput) Validate() error {
	return dara.Validate(s)
}

type ListPackageJobsResponseBodyPackageJobListPackageJobsOutput struct {
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
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

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) GetMedia() *string {
	return s.Media
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) GetType() *string {
	return s.Type
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) SetMedia(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput {
	s.Media = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) SetType(v string) *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput {
	s.Type = &v
	return s
}

func (s *ListPackageJobsResponseBodyPackageJobListPackageJobsOutput) Validate() error {
	return dara.Validate(s)
}
