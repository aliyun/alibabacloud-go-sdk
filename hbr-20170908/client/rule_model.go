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
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
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
