// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCustomizedVoiceJobResponseBodyData) *GetCustomizedVoiceJobResponseBody
	GetData() *GetCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *GetCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type GetCustomizedVoiceJobResponseBody struct {
	// The data returned if the request was successful.
	Data *GetCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceJobResponseBody) GetData() *GetCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *GetCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomizedVoiceJobResponseBody) SetData(v *GetCustomizedVoiceJobResponseBodyData) *GetCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomizedVoiceJobResponseBody) SetRequestId(v string) *GetCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBody) SetSuccess(v bool) *GetCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomizedVoiceJobResponseBodyData struct {
	// The information about the human voice cloning job.
	CustomizedVoiceJob *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob `json:"CustomizedVoiceJob,omitempty" xml:"CustomizedVoiceJob,omitempty" type:"Struct"`
}

func (s GetCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceJobResponseBodyData) GetCustomizedVoiceJob() *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	return s.CustomizedVoiceJob
}

func (s *GetCustomizedVoiceJobResponseBodyData) SetCustomizedVoiceJob(v *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) *GetCustomizedVoiceJobResponseBodyData {
	s.CustomizedVoiceJob = v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyData) Validate() error {
	if s.CustomizedVoiceJob != nil {
		if err := s.CustomizedVoiceJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob struct {
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-07T02:27:08Z
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
	// The ID of the human voice cloning job.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status description.
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
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the human voice cloning job. Valid values:
	//
	// 	- Basic
	//
	// 	- Standard
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The voice description.
	//
	// example:
	//
	// This is an exclusive voice
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	//
	// example:
	//
	// Xiaozhuan
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetGender() *string {
	return s.Gender
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetJobId() *string {
	return s.JobId
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetMessage() *string {
	return s.Message
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetScenario() *string {
	return s.Scenario
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetStatus() *string {
	return s.Status
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetType() *string {
	return s.Type
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetVoiceId() *string {
	return s.VoiceId
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) GetVoiceName() *string {
	return s.VoiceName
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetCreateTime(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.CreateTime = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetGender(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.Gender = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetJobId(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.JobId = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetMessage(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.Message = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetScenario(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.Scenario = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetStatus(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.Status = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetType(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.Type = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetVoiceDesc(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.VoiceDesc = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetVoiceId(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.VoiceId = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) SetVoiceName(v string) *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob {
	s.VoiceName = &v
	return s
}

func (s *GetCustomizedVoiceJobResponseBodyDataCustomizedVoiceJob) Validate() error {
	return dara.Validate(s)
}
