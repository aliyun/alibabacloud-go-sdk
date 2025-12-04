// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryHistoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *GetDeliveryHistoryJobResponseBody
	GetCreatedTime() *string
	SetEndTime(v string) *GetDeliveryHistoryJobResponseBody
	GetEndTime() *string
	SetHomeRegion(v string) *GetDeliveryHistoryJobResponseBody
	GetHomeRegion() *string
	SetJobId(v int64) *GetDeliveryHistoryJobResponseBody
	GetJobId() *int64
	SetJobStatus(v int32) *GetDeliveryHistoryJobResponseBody
	GetJobStatus() *int32
	SetRequestId(v string) *GetDeliveryHistoryJobResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetDeliveryHistoryJobResponseBody
	GetStartTime() *string
	SetStatus(v []*GetDeliveryHistoryJobResponseBodyStatus) *GetDeliveryHistoryJobResponseBody
	GetStatus() []*GetDeliveryHistoryJobResponseBodyStatus
	SetTrailName(v string) *GetDeliveryHistoryJobResponseBody
	GetTrailName() *string
	SetUpdatedTime(v string) *GetDeliveryHistoryJobResponseBody
	GetUpdatedTime() *string
}

type GetDeliveryHistoryJobResponseBody struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2021-05-27T07:15:03Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2021-05-27T07:20:03Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The home region of the trail.
	//
	// example:
	//
	// cn-hangzhou
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 16602
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task status. Valid values:
	//
	// 	- 0: The task is initializing.
	//
	// 	- 1: The task is delivering historical events.
	//
	// 	- 2: The task is complete.
	//
	// 	- 3: The task fails.
	//
	// example:
	//
	// 2
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FAFEC427-A00D-5653-B837-D0FA52220D8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2021-02-26T07:15:03Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of task statuses in each region.
	Status []*GetDeliveryHistoryJobResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The name of the trail based on which the task delivers events.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 2021-05-27T07:28:47Z
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetDeliveryHistoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetDeliveryHistoryJobResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDeliveryHistoryJobResponseBody) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *GetDeliveryHistoryJobResponseBody) GetJobId() *int64 {
	return s.JobId
}

func (s *GetDeliveryHistoryJobResponseBody) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *GetDeliveryHistoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeliveryHistoryJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDeliveryHistoryJobResponseBody) GetStatus() []*GetDeliveryHistoryJobResponseBodyStatus {
	return s.Status
}

func (s *GetDeliveryHistoryJobResponseBody) GetTrailName() *string {
	return s.TrailName
}

func (s *GetDeliveryHistoryJobResponseBody) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetDeliveryHistoryJobResponseBody) SetCreatedTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetEndTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetHomeRegion(v string) *GetDeliveryHistoryJobResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetJobId(v int64) *GetDeliveryHistoryJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetJobStatus(v int32) *GetDeliveryHistoryJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetRequestId(v string) *GetDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetStartTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetStatus(v []*GetDeliveryHistoryJobResponseBodyStatus) *GetDeliveryHistoryJobResponseBody {
	s.Status = v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetTrailName(v string) *GetDeliveryHistoryJobResponseBody {
	s.TrailName = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetUpdatedTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) Validate() error {
	if s.Status != nil {
		for _, item := range s.Status {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDeliveryHistoryJobResponseBodyStatus struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The task status in each region. Valid values:
	//
	// 	- 0: The task is initializing.
	//
	// 	- 1: The task is delivering historical events.
	//
	// 	- 2: The task is complete.
	//
	// 	- 3: The task fails.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeliveryHistoryJobResponseBodyStatus) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryHistoryJobResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) GetRegion() *string {
	return s.Region
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) SetRegion(v string) *GetDeliveryHistoryJobResponseBodyStatus {
	s.Region = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) SetStatus(v int32) *GetDeliveryHistoryJobResponseBodyStatus {
	s.Status = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) Validate() error {
	return dara.Validate(s)
}
