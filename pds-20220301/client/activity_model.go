// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivity interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *Activity
	GetActivityId() *string
	SetDevice(v string) *Activity
	GetDevice() *string
	SetDriveId(v string) *Activity
	GetDriveId() *string
	SetEventType(v int32) *Activity
	GetEventType() *int32
	SetLatestEventTime(v string) *Activity
	GetLatestEventTime() *string
	SetResourceCategory(v int32) *Activity
	GetResourceCategory() *int32
	SetResourceList(v []map[string]interface{}) *Activity
	GetResourceList() []map[string]interface{}
	SetTotalResourceCount(v int64) *Activity
	GetTotalResourceCount() *int64
	SetUserId(v string) *Activity
	GetUserId() *string
}

type Activity struct {
	ActivityId         *string                  `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	Device             *string                  `json:"device,omitempty" xml:"device,omitempty"`
	DriveId            *string                  `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	EventType          *int32                   `json:"event_type,omitempty" xml:"event_type,omitempty"`
	LatestEventTime    *string                  `json:"latest_event_time,omitempty" xml:"latest_event_time,omitempty"`
	ResourceCategory   *int32                   `json:"resource_category,omitempty" xml:"resource_category,omitempty"`
	ResourceList       []map[string]interface{} `json:"resource_list,omitempty" xml:"resource_list,omitempty" type:"Repeated"`
	TotalResourceCount *int64                   `json:"total_resource_count,omitempty" xml:"total_resource_count,omitempty"`
	UserId             *string                  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s Activity) String() string {
	return dara.Prettify(s)
}

func (s Activity) GoString() string {
	return s.String()
}

func (s *Activity) GetActivityId() *string {
	return s.ActivityId
}

func (s *Activity) GetDevice() *string {
	return s.Device
}

func (s *Activity) GetDriveId() *string {
	return s.DriveId
}

func (s *Activity) GetEventType() *int32 {
	return s.EventType
}

func (s *Activity) GetLatestEventTime() *string {
	return s.LatestEventTime
}

func (s *Activity) GetResourceCategory() *int32 {
	return s.ResourceCategory
}

func (s *Activity) GetResourceList() []map[string]interface{} {
	return s.ResourceList
}

func (s *Activity) GetTotalResourceCount() *int64 {
	return s.TotalResourceCount
}

func (s *Activity) GetUserId() *string {
	return s.UserId
}

func (s *Activity) SetActivityId(v string) *Activity {
	s.ActivityId = &v
	return s
}

func (s *Activity) SetDevice(v string) *Activity {
	s.Device = &v
	return s
}

func (s *Activity) SetDriveId(v string) *Activity {
	s.DriveId = &v
	return s
}

func (s *Activity) SetEventType(v int32) *Activity {
	s.EventType = &v
	return s
}

func (s *Activity) SetLatestEventTime(v string) *Activity {
	s.LatestEventTime = &v
	return s
}

func (s *Activity) SetResourceCategory(v int32) *Activity {
	s.ResourceCategory = &v
	return s
}

func (s *Activity) SetResourceList(v []map[string]interface{}) *Activity {
	s.ResourceList = v
	return s
}

func (s *Activity) SetTotalResourceCount(v int64) *Activity {
	s.TotalResourceCount = &v
	return s
}

func (s *Activity) SetUserId(v string) *Activity {
	s.UserId = &v
	return s
}

func (s *Activity) Validate() error {
	return dara.Validate(s)
}
