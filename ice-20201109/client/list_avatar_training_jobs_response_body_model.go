// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarTrainingJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAvatarTrainingJobsResponseBodyData) *ListAvatarTrainingJobsResponseBody
	GetData() *ListAvatarTrainingJobsResponseBodyData
	SetRequestId(v string) *ListAvatarTrainingJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvatarTrainingJobsResponseBody
	GetSuccess() *bool
}

type ListAvatarTrainingJobsResponseBody struct {
	// The data returned.
	Data *ListAvatarTrainingJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListAvatarTrainingJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarTrainingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvatarTrainingJobsResponseBody) GetData() *ListAvatarTrainingJobsResponseBodyData {
	return s.Data
}

func (s *ListAvatarTrainingJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvatarTrainingJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvatarTrainingJobsResponseBody) SetData(v *ListAvatarTrainingJobsResponseBodyData) *ListAvatarTrainingJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListAvatarTrainingJobsResponseBody) SetRequestId(v string) *ListAvatarTrainingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBody) SetSuccess(v bool) *ListAvatarTrainingJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAvatarTrainingJobsResponseBodyData struct {
	// The list of digital human training jobs.
	AvatarTrainingJobList []*ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList `json:"AvatarTrainingJobList,omitempty" xml:"AvatarTrainingJobList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvatarTrainingJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarTrainingJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvatarTrainingJobsResponseBodyData) GetAvatarTrainingJobList() []*ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	return s.AvatarTrainingJobList
}

func (s *ListAvatarTrainingJobsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAvatarTrainingJobsResponseBodyData) SetAvatarTrainingJobList(v []*ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) *ListAvatarTrainingJobsResponseBodyData {
	s.AvatarTrainingJobList = v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyData) SetTotalCount(v int64) *ListAvatarTrainingJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyData) Validate() error {
	if s.AvatarTrainingJobList != nil {
		for _, item := range s.AvatarTrainingJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList struct {
	// The description of the digital human.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// The ID of the digital human.
	//
	// example:
	//
	// Avatar-XXX
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// The name of the digital human.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The type of the digital human.
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	// 	- The time when the job was created.
	//
	// 	- The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-23T13:33:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// *****aded114489ea02e0addf93*****
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
	// *****aded114489ea02e0addf93*****
	Portrait *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	// The state of the digital human training job.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GoString() string {
	return s.String()
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetAvatarId() *string {
	return s.AvatarId
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetAvatarName() *string {
	return s.AvatarName
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetAvatarType() *string {
	return s.AvatarType
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetFirstTrainingTime() *string {
	return s.FirstTrainingTime
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetLastTrainingTime() *string {
	return s.LastTrainingTime
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetMessage() *string {
	return s.Message
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetPortrait() *string {
	return s.Portrait
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) GetStatus() *string {
	return s.Status
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetAvatarDescription(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.AvatarDescription = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetAvatarId(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.AvatarId = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetAvatarName(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.AvatarName = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetAvatarType(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.AvatarType = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetCreateTime(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.CreateTime = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetFirstTrainingTime(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.FirstTrainingTime = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetJobId(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.JobId = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetLastTrainingTime(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.LastTrainingTime = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetMessage(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.Message = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetPortrait(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.Portrait = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) SetStatus(v string) *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList {
	s.Status = &v
	return s
}

func (s *ListAvatarTrainingJobsResponseBodyDataAvatarTrainingJobList) Validate() error {
	return dara.Validate(s)
}
