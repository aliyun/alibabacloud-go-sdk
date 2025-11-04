// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAvatarTrainingJobResponseBodyData) *GetAvatarTrainingJobResponseBody
	GetData() *GetAvatarTrainingJobResponseBodyData
	SetRequestId(v string) *GetAvatarTrainingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAvatarTrainingJobResponseBody
	GetSuccess() *bool
}

type GetAvatarTrainingJobResponseBody struct {
	// The data returned if the request was successful.
	Data *GetAvatarTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvatarTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvatarTrainingJobResponseBody) GetData() *GetAvatarTrainingJobResponseBodyData {
	return s.Data
}

func (s *GetAvatarTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvatarTrainingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAvatarTrainingJobResponseBody) SetData(v *GetAvatarTrainingJobResponseBodyData) *GetAvatarTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *GetAvatarTrainingJobResponseBody) SetRequestId(v string) *GetAvatarTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBody) SetSuccess(v bool) *GetAvatarTrainingJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAvatarTrainingJobResponseBodyData struct {
	// The information about the digital human training job.
	AvatarTrainingJob *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob `json:"AvatarTrainingJob,omitempty" xml:"AvatarTrainingJob,omitempty" type:"Struct"`
}

func (s GetAvatarTrainingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvatarTrainingJobResponseBodyData) GetAvatarTrainingJob() *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	return s.AvatarTrainingJob
}

func (s *GetAvatarTrainingJobResponseBodyData) SetAvatarTrainingJob(v *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) *GetAvatarTrainingJobResponseBodyData {
	s.AvatarTrainingJob = v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyData) Validate() error {
	if s.AvatarTrainingJob != nil {
		if err := s.AvatarTrainingJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob struct {
	// The description of the digital human.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// The ID of the digital human.
	//
	// example:
	//
	// Avatar-XXXX
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The name of the digital human.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The type of the digital human.
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	// 	- The time when the first training was initiated.
	//
	// 	- The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:40Z
	FirstTrainingTime *string `json:"FirstTrainingTime,omitempty" xml:"FirstTrainingTime,omitempty"`
	// The ID of the digital human training job.
	//
	// example:
	//
	// ****55d86f7f4587943ce7734d6b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 	- The time when the last training was initiated.
	//
	// 	- The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:40Z
	LastTrainingTime *string `json:"LastTrainingTime,omitempty" xml:"LastTrainingTime,omitempty"`
	// The status description.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The media asset ID of the portrait image.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Portrait *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	// 	- The state of the digital human training job.
	//
	// 	- Valid values: Init, Queuing, Training, Success, and Fail.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/thumbnail.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// Indicates whether the input video supports alpha channels.
	//
	// example:
	//
	// true
	Transparent *bool `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
	// The ID of the video used for training.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GoString() string {
	return s.String()
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetAvatarName() *string {
	return s.AvatarName
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetAvatarType() *string {
	return s.AvatarType
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetFirstTrainingTime() *string {
	return s.FirstTrainingTime
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetJobId() *string {
	return s.JobId
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetLastTrainingTime() *string {
	return s.LastTrainingTime
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetMessage() *string {
	return s.Message
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetPortrait() *string {
	return s.Portrait
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetStatus() *string {
	return s.Status
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetTransparent() *bool {
	return s.Transparent
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) GetVideo() *string {
	return s.Video
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetAvatarDescription(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.AvatarDescription = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetAvatarId(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.AvatarId = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetAvatarName(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.AvatarName = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetAvatarType(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.AvatarType = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetFirstTrainingTime(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.FirstTrainingTime = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetJobId(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.JobId = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetLastTrainingTime(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.LastTrainingTime = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetMessage(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Message = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetPortrait(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Portrait = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetStatus(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Status = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetThumbnail(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Thumbnail = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetTransparent(v bool) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Transparent = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) SetVideo(v string) *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob {
	s.Video = &v
	return s
}

func (s *GetAvatarTrainingJobResponseBodyDataAvatarTrainingJob) Validate() error {
	return dara.Validate(s)
}
