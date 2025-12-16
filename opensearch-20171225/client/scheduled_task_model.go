// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledTask interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSwitch(v bool) *ScheduledTask
	GetAutoSwitch() *bool
	SetCron(v string) *ScheduledTask
	GetCron() *string
	SetEnabled(v bool) *ScheduledTask
	GetEnabled() *bool
	SetFilter(v *ScheduledTaskFilter) *ScheduledTask
	GetFilter() *ScheduledTaskFilter
	SetForkedAppId(v string) *ScheduledTask
	GetForkedAppId() *string
	SetPermanent(v bool) *ScheduledTask
	GetPermanent() *bool
	SetRunNow(v bool) *ScheduledTask
	GetRunNow() *bool
	SetType(v string) *ScheduledTask
	GetType() *string
	SetVersion(v string) *ScheduledTask
	GetVersion() *string
}

type ScheduledTask struct {
	AutoSwitch  *bool                `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cron        *string              `json:"cron,omitempty" xml:"cron,omitempty"`
	Enabled     *bool                `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Filter      *ScheduledTaskFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	ForkedAppId *string              `json:"forkedAppId,omitempty" xml:"forkedAppId,omitempty"`
	Permanent   *bool                `json:"permanent,omitempty" xml:"permanent,omitempty"`
	RunNow      *bool                `json:"runNow,omitempty" xml:"runNow,omitempty"`
	Type        *string              `json:"type,omitempty" xml:"type,omitempty"`
	Version     *string              `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s ScheduledTask) GoString() string {
	return s.String()
}

func (s *ScheduledTask) GetAutoSwitch() *bool {
	return s.AutoSwitch
}

func (s *ScheduledTask) GetCron() *string {
	return s.Cron
}

func (s *ScheduledTask) GetEnabled() *bool {
	return s.Enabled
}

func (s *ScheduledTask) GetFilter() *ScheduledTaskFilter {
	return s.Filter
}

func (s *ScheduledTask) GetForkedAppId() *string {
	return s.ForkedAppId
}

func (s *ScheduledTask) GetPermanent() *bool {
	return s.Permanent
}

func (s *ScheduledTask) GetRunNow() *bool {
	return s.RunNow
}

func (s *ScheduledTask) GetType() *string {
	return s.Type
}

func (s *ScheduledTask) GetVersion() *string {
	return s.Version
}

func (s *ScheduledTask) SetAutoSwitch(v bool) *ScheduledTask {
	s.AutoSwitch = &v
	return s
}

func (s *ScheduledTask) SetCron(v string) *ScheduledTask {
	s.Cron = &v
	return s
}

func (s *ScheduledTask) SetEnabled(v bool) *ScheduledTask {
	s.Enabled = &v
	return s
}

func (s *ScheduledTask) SetFilter(v *ScheduledTaskFilter) *ScheduledTask {
	s.Filter = v
	return s
}

func (s *ScheduledTask) SetForkedAppId(v string) *ScheduledTask {
	s.ForkedAppId = &v
	return s
}

func (s *ScheduledTask) SetPermanent(v bool) *ScheduledTask {
	s.Permanent = &v
	return s
}

func (s *ScheduledTask) SetRunNow(v bool) *ScheduledTask {
	s.RunNow = &v
	return s
}

func (s *ScheduledTask) SetType(v string) *ScheduledTask {
	s.Type = &v
	return s
}

func (s *ScheduledTask) SetVersion(v string) *ScheduledTask {
	s.Version = &v
	return s
}

func (s *ScheduledTask) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScheduledTaskFilter struct {
	Days       *int32  `json:"days,omitempty" xml:"days,omitempty"`
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	Field      *string `json:"field,omitempty" xml:"field,omitempty"`
	Unit       *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ScheduledTaskFilter) String() string {
	return dara.Prettify(s)
}

func (s ScheduledTaskFilter) GoString() string {
	return s.String()
}

func (s *ScheduledTaskFilter) GetDays() *int32 {
	return s.Days
}

func (s *ScheduledTaskFilter) GetExpression() *string {
	return s.Expression
}

func (s *ScheduledTaskFilter) GetField() *string {
	return s.Field
}

func (s *ScheduledTaskFilter) GetUnit() *string {
	return s.Unit
}

func (s *ScheduledTaskFilter) SetDays(v int32) *ScheduledTaskFilter {
	s.Days = &v
	return s
}

func (s *ScheduledTaskFilter) SetExpression(v string) *ScheduledTaskFilter {
	s.Expression = &v
	return s
}

func (s *ScheduledTaskFilter) SetField(v string) *ScheduledTaskFilter {
	s.Field = &v
	return s
}

func (s *ScheduledTaskFilter) SetUnit(v string) *ScheduledTaskFilter {
	s.Unit = &v
	return s
}

func (s *ScheduledTaskFilter) Validate() error {
	return dara.Validate(s)
}
