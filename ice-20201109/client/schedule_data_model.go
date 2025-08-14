// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleData interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v []*ScheduleDataAdBreaks) *ScheduleData
	GetAdBreaks() []*ScheduleDataAdBreaks
	SetApproximateDurationSeconds(v int64) *ScheduleData
	GetApproximateDurationSeconds() *int64
	SetApproximateStartTime(v string) *ScheduleData
	GetApproximateStartTime() *string
	SetEntryType(v string) *ScheduleData
	GetEntryType() *string
	SetProgramName(v string) *ScheduleData
	GetProgramName() *string
	SetSourceLocationName(v string) *ScheduleData
	GetSourceLocationName() *string
	SetSourceName(v string) *ScheduleData
	GetSourceName() *string
	SetSourceType(v string) *ScheduleData
	GetSourceType() *string
}

type ScheduleData struct {
	AdBreaks                   []*ScheduleDataAdBreaks `json:"AdBreaks,omitempty" xml:"AdBreaks,omitempty" type:"Repeated"`
	ApproximateDurationSeconds *int64                  `json:"ApproximateDurationSeconds,omitempty" xml:"ApproximateDurationSeconds,omitempty"`
	ApproximateStartTime       *string                 `json:"ApproximateStartTime,omitempty" xml:"ApproximateStartTime,omitempty"`
	EntryType                  *string                 `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	ProgramName                *string                 `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName         *string                 `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName                 *string                 `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType                 *string                 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ScheduleData) String() string {
	return dara.Prettify(s)
}

func (s ScheduleData) GoString() string {
	return s.String()
}

func (s *ScheduleData) GetAdBreaks() []*ScheduleDataAdBreaks {
	return s.AdBreaks
}

func (s *ScheduleData) GetApproximateDurationSeconds() *int64 {
	return s.ApproximateDurationSeconds
}

func (s *ScheduleData) GetApproximateStartTime() *string {
	return s.ApproximateStartTime
}

func (s *ScheduleData) GetEntryType() *string {
	return s.EntryType
}

func (s *ScheduleData) GetProgramName() *string {
	return s.ProgramName
}

func (s *ScheduleData) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ScheduleData) GetSourceName() *string {
	return s.SourceName
}

func (s *ScheduleData) GetSourceType() *string {
	return s.SourceType
}

func (s *ScheduleData) SetAdBreaks(v []*ScheduleDataAdBreaks) *ScheduleData {
	s.AdBreaks = v
	return s
}

func (s *ScheduleData) SetApproximateDurationSeconds(v int64) *ScheduleData {
	s.ApproximateDurationSeconds = &v
	return s
}

func (s *ScheduleData) SetApproximateStartTime(v string) *ScheduleData {
	s.ApproximateStartTime = &v
	return s
}

func (s *ScheduleData) SetEntryType(v string) *ScheduleData {
	s.EntryType = &v
	return s
}

func (s *ScheduleData) SetProgramName(v string) *ScheduleData {
	s.ProgramName = &v
	return s
}

func (s *ScheduleData) SetSourceLocationName(v string) *ScheduleData {
	s.SourceLocationName = &v
	return s
}

func (s *ScheduleData) SetSourceName(v string) *ScheduleData {
	s.SourceName = &v
	return s
}

func (s *ScheduleData) SetSourceType(v string) *ScheduleData {
	s.SourceType = &v
	return s
}

func (s *ScheduleData) Validate() error {
	return dara.Validate(s)
}

type ScheduleDataAdBreaks struct {
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OffsetMillis         *string `json:"OffsetMillis,omitempty" xml:"OffsetMillis,omitempty"`
	SourceLocationName   *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName           *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SpliceInsertSettings *string `json:"SpliceInsertSettings,omitempty" xml:"SpliceInsertSettings,omitempty"`
	TimeSignalSettings   *string `json:"TimeSignalSettings,omitempty" xml:"TimeSignalSettings,omitempty"`
}

func (s ScheduleDataAdBreaks) String() string {
	return dara.Prettify(s)
}

func (s ScheduleDataAdBreaks) GoString() string {
	return s.String()
}

func (s *ScheduleDataAdBreaks) GetMessageType() *string {
	return s.MessageType
}

func (s *ScheduleDataAdBreaks) GetOffsetMillis() *string {
	return s.OffsetMillis
}

func (s *ScheduleDataAdBreaks) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ScheduleDataAdBreaks) GetSourceName() *string {
	return s.SourceName
}

func (s *ScheduleDataAdBreaks) GetSpliceInsertSettings() *string {
	return s.SpliceInsertSettings
}

func (s *ScheduleDataAdBreaks) GetTimeSignalSettings() *string {
	return s.TimeSignalSettings
}

func (s *ScheduleDataAdBreaks) SetMessageType(v string) *ScheduleDataAdBreaks {
	s.MessageType = &v
	return s
}

func (s *ScheduleDataAdBreaks) SetOffsetMillis(v string) *ScheduleDataAdBreaks {
	s.OffsetMillis = &v
	return s
}

func (s *ScheduleDataAdBreaks) SetSourceLocationName(v string) *ScheduleDataAdBreaks {
	s.SourceLocationName = &v
	return s
}

func (s *ScheduleDataAdBreaks) SetSourceName(v string) *ScheduleDataAdBreaks {
	s.SourceName = &v
	return s
}

func (s *ScheduleDataAdBreaks) SetSpliceInsertSettings(v string) *ScheduleDataAdBreaks {
	s.SpliceInsertSettings = &v
	return s
}

func (s *ScheduleDataAdBreaks) SetTimeSignalSettings(v string) *ScheduleDataAdBreaks {
	s.TimeSignalSettings = &v
	return s
}

func (s *ScheduleDataAdBreaks) Validate() error {
	return dara.Validate(s)
}
