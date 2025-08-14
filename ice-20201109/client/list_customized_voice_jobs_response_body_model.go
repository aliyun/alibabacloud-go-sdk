// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoiceJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCustomizedVoiceJobsResponseBodyData) *ListCustomizedVoiceJobsResponseBody
	GetData() *ListCustomizedVoiceJobsResponseBodyData
	SetRequestId(v string) *ListCustomizedVoiceJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomizedVoiceJobsResponseBody
	GetSuccess() *bool
}

type ListCustomizedVoiceJobsResponseBody struct {
	// The data returned.
	Data *ListCustomizedVoiceJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomizedVoiceJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoiceJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoiceJobsResponseBody) GetData() *ListCustomizedVoiceJobsResponseBodyData {
	return s.Data
}

func (s *ListCustomizedVoiceJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomizedVoiceJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomizedVoiceJobsResponseBody) SetData(v *ListCustomizedVoiceJobsResponseBodyData) *ListCustomizedVoiceJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBody) SetRequestId(v string) *ListCustomizedVoiceJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBody) SetSuccess(v bool) *ListCustomizedVoiceJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCustomizedVoiceJobsResponseBodyData struct {
	// The queried human voice cloning jobs.
	CustomizedVoiceJobList []*ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList `json:"CustomizedVoiceJobList,omitempty" xml:"CustomizedVoiceJobList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 271
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomizedVoiceJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoiceJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoiceJobsResponseBodyData) GetCustomizedVoiceJobList() []*ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	return s.CustomizedVoiceJobList
}

func (s *ListCustomizedVoiceJobsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomizedVoiceJobsResponseBodyData) SetCustomizedVoiceJobList(v []*ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) *ListCustomizedVoiceJobsResponseBodyData {
	s.CustomizedVoiceJobList = v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyData) SetTotalCount(v int32) *ListCustomizedVoiceJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList struct {
	// 	- The time when the job was created.
	//
	// 	- The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T06:23:59Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The gender. Valid values:
	//
	// 	- female
	//
	// 	- male
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2022-06-27T02:42:28Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the human voice cloning job.
	//
	// example:
	//
	// 2245ab99a7fd4116a4fd3f499b7a56c5
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The scenario. Valid values:
	//
	// 	- story
	//
	// 	- interaction
	//
	// 	- navigation
	//
	// example:
	//
	// story
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The job state. Valid values:
	//
	// 	- Initialization
	//
	// 	- AudioDetecting
	//
	// 	- PreTraining
	//
	// 	- Training
	//
	// 	- Success
	//
	// 	- Fail
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 	- The voice type. Valid values:
	//
	//     	- Basic
	//
	//     	- Standard
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The voice description.
	//
	// 	- The description can be up to 256 characters in length.
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	//
	// 	- The name can be up to 32 characters in length.
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetGender() *string {
	return s.Gender
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetMessage() *string {
	return s.Message
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetScenario() *string {
	return s.Scenario
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetStatus() *string {
	return s.Status
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetType() *string {
	return s.Type
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetVoiceId() *string {
	return s.VoiceId
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) GetVoiceName() *string {
	return s.VoiceName
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetCreateTime(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.CreateTime = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetGender(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.Gender = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetGmtCreate(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.GmtCreate = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetJobId(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.JobId = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetMessage(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.Message = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetScenario(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.Scenario = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetStatus(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.Status = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetType(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.Type = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetVoiceDesc(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.VoiceDesc = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetVoiceId(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.VoiceId = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) SetVoiceName(v string) *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList {
	s.VoiceName = &v
	return s
}

func (s *ListCustomizedVoiceJobsResponseBodyDataCustomizedVoiceJobList) Validate() error {
	return dara.Validate(s)
}
