// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelAssemblyScheduleData interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v []*ChannelAssemblyScheduleDataAdBreaks) *ChannelAssemblyScheduleData
	GetAdBreaks() []*ChannelAssemblyScheduleDataAdBreaks
	SetApproximateDurationSeconds(v int64) *ChannelAssemblyScheduleData
	GetApproximateDurationSeconds() *int64
	SetApproximateStartTime(v string) *ChannelAssemblyScheduleData
	GetApproximateStartTime() *string
	SetEntryType(v string) *ChannelAssemblyScheduleData
	GetEntryType() *string
	SetProgramName(v string) *ChannelAssemblyScheduleData
	GetProgramName() *string
	SetSourceLocationName(v string) *ChannelAssemblyScheduleData
	GetSourceLocationName() *string
	SetSourceName(v string) *ChannelAssemblyScheduleData
	GetSourceName() *string
	SetSourceType(v string) *ChannelAssemblyScheduleData
	GetSourceType() *string
}

type ChannelAssemblyScheduleData struct {
	AdBreaks                   []*ChannelAssemblyScheduleDataAdBreaks `json:"AdBreaks,omitempty" xml:"AdBreaks,omitempty" type:"Repeated"`
	ApproximateDurationSeconds *int64                                 `json:"ApproximateDurationSeconds,omitempty" xml:"ApproximateDurationSeconds,omitempty"`
	ApproximateStartTime       *string                                `json:"ApproximateStartTime,omitempty" xml:"ApproximateStartTime,omitempty"`
	EntryType                  *string                                `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	ProgramName                *string                                `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName         *string                                `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName                 *string                                `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType                 *string                                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ChannelAssemblyScheduleData) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyScheduleData) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyScheduleData) GetAdBreaks() []*ChannelAssemblyScheduleDataAdBreaks {
	return s.AdBreaks
}

func (s *ChannelAssemblyScheduleData) GetApproximateDurationSeconds() *int64 {
	return s.ApproximateDurationSeconds
}

func (s *ChannelAssemblyScheduleData) GetApproximateStartTime() *string {
	return s.ApproximateStartTime
}

func (s *ChannelAssemblyScheduleData) GetEntryType() *string {
	return s.EntryType
}

func (s *ChannelAssemblyScheduleData) GetProgramName() *string {
	return s.ProgramName
}

func (s *ChannelAssemblyScheduleData) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblyScheduleData) GetSourceName() *string {
	return s.SourceName
}

func (s *ChannelAssemblyScheduleData) GetSourceType() *string {
	return s.SourceType
}

func (s *ChannelAssemblyScheduleData) SetAdBreaks(v []*ChannelAssemblyScheduleDataAdBreaks) *ChannelAssemblyScheduleData {
	s.AdBreaks = v
	return s
}

func (s *ChannelAssemblyScheduleData) SetApproximateDurationSeconds(v int64) *ChannelAssemblyScheduleData {
	s.ApproximateDurationSeconds = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetApproximateStartTime(v string) *ChannelAssemblyScheduleData {
	s.ApproximateStartTime = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetEntryType(v string) *ChannelAssemblyScheduleData {
	s.EntryType = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetProgramName(v string) *ChannelAssemblyScheduleData {
	s.ProgramName = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetSourceLocationName(v string) *ChannelAssemblyScheduleData {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetSourceName(v string) *ChannelAssemblyScheduleData {
	s.SourceName = &v
	return s
}

func (s *ChannelAssemblyScheduleData) SetSourceType(v string) *ChannelAssemblyScheduleData {
	s.SourceType = &v
	return s
}

func (s *ChannelAssemblyScheduleData) Validate() error {
	return dara.Validate(s)
}

type ChannelAssemblyScheduleDataAdBreaks struct {
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OffsetMillis         *string `json:"OffsetMillis,omitempty" xml:"OffsetMillis,omitempty"`
	SourceLocationName   *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName           *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SpliceInsertSettings *string `json:"SpliceInsertSettings,omitempty" xml:"SpliceInsertSettings,omitempty"`
	TimeSignalSettings   *string `json:"TimeSignalSettings,omitempty" xml:"TimeSignalSettings,omitempty"`
}

func (s ChannelAssemblyScheduleDataAdBreaks) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyScheduleDataAdBreaks) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetMessageType() *string {
	return s.MessageType
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetOffsetMillis() *string {
	return s.OffsetMillis
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetSourceName() *string {
	return s.SourceName
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetSpliceInsertSettings() *string {
	return s.SpliceInsertSettings
}

func (s *ChannelAssemblyScheduleDataAdBreaks) GetTimeSignalSettings() *string {
	return s.TimeSignalSettings
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetMessageType(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.MessageType = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetOffsetMillis(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.OffsetMillis = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetSourceLocationName(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetSourceName(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.SourceName = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetSpliceInsertSettings(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.SpliceInsertSettings = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) SetTimeSignalSettings(v string) *ChannelAssemblyScheduleDataAdBreaks {
	s.TimeSignalSettings = &v
	return s
}

func (s *ChannelAssemblyScheduleDataAdBreaks) Validate() error {
	return dara.Validate(s)
}
