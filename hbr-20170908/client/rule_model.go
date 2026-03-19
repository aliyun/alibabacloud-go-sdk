// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRule interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *Rule
	GetBackupType() *string
	SetDestinationRegionId(v string) *Rule
	GetDestinationRegionId() *string
	SetDestinationRetention(v int64) *Rule
	GetDestinationRetention() *int64
	SetDisabled(v bool) *Rule
	GetDisabled() *bool
	SetDoCopy(v bool) *Rule
	GetDoCopy() *bool
	SetRetention(v int64) *Rule
	GetRetention() *int64
	SetRuleName(v string) *Rule
	GetRuleName() *string
	SetSchedule(v string) *Rule
	GetSchedule() *string
}

type Rule struct {
	// The backup type.
	//
	// 	- COMPLETE：Full backup.
	//
	// 	- INCREMENTAL：Incremental backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The replication region id.
	//
	// example:
	//
	// cn-shenzhen
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The retention period of remote backups. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// Whether the plan is disbaled for this data source.
	//
	// - **true**: disabled
	//
	// - **false**: Not disabled
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Whether do copy.
	//
	// example:
	//
	// true
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// This parameter indicates the retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule-com-backup-20241023-163113
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **BACKUP**. This parameter indicates the backup schedule settings. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1631685600|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s Rule) String() string {
	return dara.Prettify(s)
}

func (s Rule) GoString() string {
	return s.String()
}

func (s *Rule) GetBackupType() *string {
	return s.BackupType
}

func (s *Rule) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *Rule) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *Rule) GetDisabled() *bool {
	return s.Disabled
}

func (s *Rule) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *Rule) GetRetention() *int64 {
	return s.Retention
}

func (s *Rule) GetRuleName() *string {
	return s.RuleName
}

func (s *Rule) GetSchedule() *string {
	return s.Schedule
}

func (s *Rule) SetBackupType(v string) *Rule {
	s.BackupType = &v
	return s
}

func (s *Rule) SetDestinationRegionId(v string) *Rule {
	s.DestinationRegionId = &v
	return s
}

func (s *Rule) SetDestinationRetention(v int64) *Rule {
	s.DestinationRetention = &v
	return s
}

func (s *Rule) SetDisabled(v bool) *Rule {
	s.Disabled = &v
	return s
}

func (s *Rule) SetDoCopy(v bool) *Rule {
	s.DoCopy = &v
	return s
}

func (s *Rule) SetRetention(v int64) *Rule {
	s.Retention = &v
	return s
}

func (s *Rule) SetRuleName(v string) *Rule {
	s.RuleName = &v
	return s
}

func (s *Rule) SetSchedule(v string) *Rule {
	s.Schedule = &v
	return s
}

func (s *Rule) Validate() error {
	return dara.Validate(s)
}
