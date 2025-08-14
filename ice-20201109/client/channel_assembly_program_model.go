// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelAssemblyProgram interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v []*ChannelAssemblyProgramAdBreaks) *ChannelAssemblyProgram
	GetAdBreaks() []*ChannelAssemblyProgramAdBreaks
	SetArn(v string) *ChannelAssemblyProgram
	GetArn() *string
	SetChannelName(v string) *ChannelAssemblyProgram
	GetChannelName() *string
	SetClipRange(v string) *ChannelAssemblyProgram
	GetClipRange() *string
	SetGmtCreate(v string) *ChannelAssemblyProgram
	GetGmtCreate() *string
	SetGmtModified(v string) *ChannelAssemblyProgram
	GetGmtModified() *string
	SetProgramName(v string) *ChannelAssemblyProgram
	GetProgramName() *string
	SetSourceLocationName(v string) *ChannelAssemblyProgram
	GetSourceLocationName() *string
	SetSourceName(v string) *ChannelAssemblyProgram
	GetSourceName() *string
	SetSourceType(v string) *ChannelAssemblyProgram
	GetSourceType() *string
	SetTransition(v string) *ChannelAssemblyProgram
	GetTransition() *string
}

type ChannelAssemblyProgram struct {
	AdBreaks           []*ChannelAssemblyProgramAdBreaks `json:"AdBreaks,omitempty" xml:"AdBreaks,omitempty" type:"Repeated"`
	Arn                *string                           `json:"Arn,omitempty" xml:"Arn,omitempty"`
	ChannelName        *string                           `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ClipRange          *string                           `json:"ClipRange,omitempty" xml:"ClipRange,omitempty"`
	GmtCreate          *string                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ProgramName        *string                           `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName *string                           `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName         *string                           `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType         *string                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Transition         *string                           `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s ChannelAssemblyProgram) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyProgram) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyProgram) GetAdBreaks() []*ChannelAssemblyProgramAdBreaks {
	return s.AdBreaks
}

func (s *ChannelAssemblyProgram) GetArn() *string {
	return s.Arn
}

func (s *ChannelAssemblyProgram) GetChannelName() *string {
	return s.ChannelName
}

func (s *ChannelAssemblyProgram) GetClipRange() *string {
	return s.ClipRange
}

func (s *ChannelAssemblyProgram) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ChannelAssemblyProgram) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ChannelAssemblyProgram) GetProgramName() *string {
	return s.ProgramName
}

func (s *ChannelAssemblyProgram) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblyProgram) GetSourceName() *string {
	return s.SourceName
}

func (s *ChannelAssemblyProgram) GetSourceType() *string {
	return s.SourceType
}

func (s *ChannelAssemblyProgram) GetTransition() *string {
	return s.Transition
}

func (s *ChannelAssemblyProgram) SetAdBreaks(v []*ChannelAssemblyProgramAdBreaks) *ChannelAssemblyProgram {
	s.AdBreaks = v
	return s
}

func (s *ChannelAssemblyProgram) SetArn(v string) *ChannelAssemblyProgram {
	s.Arn = &v
	return s
}

func (s *ChannelAssemblyProgram) SetChannelName(v string) *ChannelAssemblyProgram {
	s.ChannelName = &v
	return s
}

func (s *ChannelAssemblyProgram) SetClipRange(v string) *ChannelAssemblyProgram {
	s.ClipRange = &v
	return s
}

func (s *ChannelAssemblyProgram) SetGmtCreate(v string) *ChannelAssemblyProgram {
	s.GmtCreate = &v
	return s
}

func (s *ChannelAssemblyProgram) SetGmtModified(v string) *ChannelAssemblyProgram {
	s.GmtModified = &v
	return s
}

func (s *ChannelAssemblyProgram) SetProgramName(v string) *ChannelAssemblyProgram {
	s.ProgramName = &v
	return s
}

func (s *ChannelAssemblyProgram) SetSourceLocationName(v string) *ChannelAssemblyProgram {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblyProgram) SetSourceName(v string) *ChannelAssemblyProgram {
	s.SourceName = &v
	return s
}

func (s *ChannelAssemblyProgram) SetSourceType(v string) *ChannelAssemblyProgram {
	s.SourceType = &v
	return s
}

func (s *ChannelAssemblyProgram) SetTransition(v string) *ChannelAssemblyProgram {
	s.Transition = &v
	return s
}

func (s *ChannelAssemblyProgram) Validate() error {
	return dara.Validate(s)
}

type ChannelAssemblyProgramAdBreaks struct {
	ChannelName          *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OffsetMillis         *int64  `json:"OffsetMillis,omitempty" xml:"OffsetMillis,omitempty"`
	ProgramName          *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName   *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName           *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SpliceInsertSettings *string `json:"SpliceInsertSettings,omitempty" xml:"SpliceInsertSettings,omitempty"`
	TimeSignalSettings   *string `json:"TimeSignalSettings,omitempty" xml:"TimeSignalSettings,omitempty"`
}

func (s ChannelAssemblyProgramAdBreaks) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyProgramAdBreaks) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyProgramAdBreaks) GetChannelName() *string {
	return s.ChannelName
}

func (s *ChannelAssemblyProgramAdBreaks) GetMessageType() *string {
	return s.MessageType
}

func (s *ChannelAssemblyProgramAdBreaks) GetOffsetMillis() *int64 {
	return s.OffsetMillis
}

func (s *ChannelAssemblyProgramAdBreaks) GetProgramName() *string {
	return s.ProgramName
}

func (s *ChannelAssemblyProgramAdBreaks) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ChannelAssemblyProgramAdBreaks) GetSourceName() *string {
	return s.SourceName
}

func (s *ChannelAssemblyProgramAdBreaks) GetSpliceInsertSettings() *string {
	return s.SpliceInsertSettings
}

func (s *ChannelAssemblyProgramAdBreaks) GetTimeSignalSettings() *string {
	return s.TimeSignalSettings
}

func (s *ChannelAssemblyProgramAdBreaks) SetChannelName(v string) *ChannelAssemblyProgramAdBreaks {
	s.ChannelName = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetMessageType(v string) *ChannelAssemblyProgramAdBreaks {
	s.MessageType = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetOffsetMillis(v int64) *ChannelAssemblyProgramAdBreaks {
	s.OffsetMillis = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetProgramName(v string) *ChannelAssemblyProgramAdBreaks {
	s.ProgramName = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetSourceLocationName(v string) *ChannelAssemblyProgramAdBreaks {
	s.SourceLocationName = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetSourceName(v string) *ChannelAssemblyProgramAdBreaks {
	s.SourceName = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetSpliceInsertSettings(v string) *ChannelAssemblyProgramAdBreaks {
	s.SpliceInsertSettings = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) SetTimeSignalSettings(v string) *ChannelAssemblyProgramAdBreaks {
	s.TimeSignalSettings = &v
	return s
}

func (s *ChannelAssemblyProgramAdBreaks) Validate() error {
	return dara.Validate(s)
}
