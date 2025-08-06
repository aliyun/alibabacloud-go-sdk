// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobInfoList(v []*ListJobInfoResponseBodyJobInfoList) *ListJobInfoResponseBody
	GetJobInfoList() []*ListJobInfoResponseBodyJobInfoList
	SetJobType(v string) *ListJobInfoResponseBody
	GetJobType() *string
	SetMediaId(v string) *ListJobInfoResponseBody
	GetMediaId() *string
	SetRequestId(v string) *ListJobInfoResponseBody
	GetRequestId() *string
}

type ListJobInfoResponseBody struct {
	// The historical tasks of the last 6 months.
	JobInfoList []*ListJobInfoResponseBodyJobInfoList `json:"JobInfoList,omitempty" xml:"JobInfoList,omitempty" type:"Repeated"`
	// The type of the task. Valid values:
	//
	// 	- transcode
	//
	// 	- snapshot
	//
	// 	- ai
	//
	// example:
	//
	// transcode
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6708D849-F109-1A6C-AC91-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobInfoResponseBody) GetJobInfoList() []*ListJobInfoResponseBodyJobInfoList {
	return s.JobInfoList
}

func (s *ListJobInfoResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *ListJobInfoResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *ListJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobInfoResponseBody) SetJobInfoList(v []*ListJobInfoResponseBodyJobInfoList) *ListJobInfoResponseBody {
	s.JobInfoList = v
	return s
}

func (s *ListJobInfoResponseBody) SetJobType(v string) *ListJobInfoResponseBody {
	s.JobType = &v
	return s
}

func (s *ListJobInfoResponseBody) SetMediaId(v string) *ListJobInfoResponseBody {
	s.MediaId = &v
	return s
}

func (s *ListJobInfoResponseBody) SetRequestId(v string) *ListJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListJobInfoResponseBodyJobInfoList struct {
	// The time when the task was complete.
	//
	// example:
	//
	// 2024-10-14T07:39:34Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 5c9dff751ba**********59d50a967f5
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// TranscodeSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who submitted the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListJobInfoResponseBodyJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfoResponseBodyJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobInfoResponseBodyJobInfoList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListJobInfoResponseBodyJobInfoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobInfoResponseBodyJobInfoList) GetJobId() *string {
	return s.JobId
}

func (s *ListJobInfoResponseBodyJobInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListJobInfoResponseBodyJobInfoList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListJobInfoResponseBodyJobInfoList) SetCompleteTime(v string) *ListJobInfoResponseBodyJobInfoList {
	s.CompleteTime = &v
	return s
}

func (s *ListJobInfoResponseBodyJobInfoList) SetCreateTime(v string) *ListJobInfoResponseBodyJobInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListJobInfoResponseBodyJobInfoList) SetJobId(v string) *ListJobInfoResponseBodyJobInfoList {
	s.JobId = &v
	return s
}

func (s *ListJobInfoResponseBodyJobInfoList) SetStatus(v string) *ListJobInfoResponseBodyJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobInfoResponseBodyJobInfoList) SetUserId(v int64) *ListJobInfoResponseBodyJobInfoList {
	s.UserId = &v
	return s
}

func (s *ListJobInfoResponseBodyJobInfoList) Validate() error {
	return dara.Validate(s)
}
