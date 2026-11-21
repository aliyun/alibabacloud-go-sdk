// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeOncallSchedule interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v map[string]interface{}) *MergeOncallSchedule
	GetDetail() map[string]interface{}
	SetGmtCreate(v string) *MergeOncallSchedule
	GetGmtCreate() *string
	SetGmtModified(v string) *MergeOncallSchedule
	GetGmtModified() *string
	SetIdentifier(v string) *MergeOncallSchedule
	GetIdentifier() *string
	SetName(v string) *MergeOncallSchedule
	GetName() *string
	SetSource(v string) *MergeOncallSchedule
	GetSource() *string
	SetWorkspace(v string) *MergeOncallSchedule
	GetWorkspace() *string
}

type MergeOncallSchedule struct {
	// The extended details.
	//
	// example:
	//
	// {
	//
	//   "oncallScheduleId": "testId",
	//
	//   "oncallScheduleName": "testName",
	//
	//   "rotations": [
	//
	//     {
	//
	//       "activeDays": [
	//
	//         1
	//
	//       ],
	//
	//       "contacts": [
	//
	//         "testId"
	//
	//       ],
	//
	//       "rotationEndTime": "24:00",
	//
	//       "rotationName": "testName",
	//
	//       "rotationStartTime": "00:00",
	//
	//       "shiftRecurrenceFrequency": "WEEKLY",
	//
	//       "startDate": "2024-01-01",
	//
	//       "timeZone": "Asia/Shanghai",
	//
	//       "shiftLength": 3
	//
	//     }
	//
	//   ],
	//
	//   "shiftRobotId": "testId",
	//
	//   "substitudes": [
	//
	//     {
	//
	//       "startTime": "2024-10-10 09:00",
	//
	//       "endTime": "2024-10-15 18:00",
	//
	//       "replacedContactId": "testId",
	//
	//       "contactId": "testId"
	//
	//     }
	//
	//   ],
	//
	//   "workspace": "default-cms-5884500827389836-cn-hangzhou"
	//
	// }
	Detail map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1700000000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1700000000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The unique identifier.
	//
	// example:
	//
	// Sample value
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The name of the schedule.
	//
	// example:
	//
	// Sample name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source.
	//
	// example:
	//
	// Sample value
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The workspace.
	//
	// example:
	//
	// Sample value
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s MergeOncallSchedule) String() string {
	return dara.Prettify(s)
}

func (s MergeOncallSchedule) GoString() string {
	return s.String()
}

func (s *MergeOncallSchedule) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *MergeOncallSchedule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MergeOncallSchedule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MergeOncallSchedule) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeOncallSchedule) GetName() *string {
	return s.Name
}

func (s *MergeOncallSchedule) GetSource() *string {
	return s.Source
}

func (s *MergeOncallSchedule) GetWorkspace() *string {
	return s.Workspace
}

func (s *MergeOncallSchedule) SetDetail(v map[string]interface{}) *MergeOncallSchedule {
	s.Detail = v
	return s
}

func (s *MergeOncallSchedule) SetGmtCreate(v string) *MergeOncallSchedule {
	s.GmtCreate = &v
	return s
}

func (s *MergeOncallSchedule) SetGmtModified(v string) *MergeOncallSchedule {
	s.GmtModified = &v
	return s
}

func (s *MergeOncallSchedule) SetIdentifier(v string) *MergeOncallSchedule {
	s.Identifier = &v
	return s
}

func (s *MergeOncallSchedule) SetName(v string) *MergeOncallSchedule {
	s.Name = &v
	return s
}

func (s *MergeOncallSchedule) SetSource(v string) *MergeOncallSchedule {
	s.Source = &v
	return s
}

func (s *MergeOncallSchedule) SetWorkspace(v string) *MergeOncallSchedule {
	s.Workspace = &v
	return s
}

func (s *MergeOncallSchedule) Validate() error {
	return dara.Validate(s)
}
