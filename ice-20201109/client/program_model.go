// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgram interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v []*ProgramAdBreaks) *Program
	GetAdBreaks() []*ProgramAdBreaks
	SetArn(v string) *Program
	GetArn() *string
	SetChannelName(v string) *Program
	GetChannelName() *string
	SetClipRange(v string) *Program
	GetClipRange() *string
	SetGmtCreate(v string) *Program
	GetGmtCreate() *string
	SetGmtModified(v string) *Program
	GetGmtModified() *string
	SetProgramName(v string) *Program
	GetProgramName() *string
	SetSourceLocationName(v string) *Program
	GetSourceLocationName() *string
	SetSourceName(v string) *Program
	GetSourceName() *string
	SetSourceType(v string) *Program
	GetSourceType() *string
	SetTransition(v string) *Program
	GetTransition() *string
}

type Program struct {
	AdBreaks           []*ProgramAdBreaks `json:"AdBreaks,omitempty" xml:"AdBreaks,omitempty" type:"Repeated"`
	Arn                *string            `json:"Arn,omitempty" xml:"Arn,omitempty"`
	ChannelName        *string            `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ClipRange          *string            `json:"ClipRange,omitempty" xml:"ClipRange,omitempty"`
	GmtCreate          *string            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string            `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ProgramName        *string            `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName *string            `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName         *string            `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType         *string            `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Transition         *string            `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s Program) String() string {
	return dara.Prettify(s)
}

func (s Program) GoString() string {
	return s.String()
}

func (s *Program) GetAdBreaks() []*ProgramAdBreaks {
	return s.AdBreaks
}

func (s *Program) GetArn() *string {
	return s.Arn
}

func (s *Program) GetChannelName() *string {
	return s.ChannelName
}

func (s *Program) GetClipRange() *string {
	return s.ClipRange
}

func (s *Program) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Program) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Program) GetProgramName() *string {
	return s.ProgramName
}

func (s *Program) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *Program) GetSourceName() *string {
	return s.SourceName
}

func (s *Program) GetSourceType() *string {
	return s.SourceType
}

func (s *Program) GetTransition() *string {
	return s.Transition
}

func (s *Program) SetAdBreaks(v []*ProgramAdBreaks) *Program {
	s.AdBreaks = v
	return s
}

func (s *Program) SetArn(v string) *Program {
	s.Arn = &v
	return s
}

func (s *Program) SetChannelName(v string) *Program {
	s.ChannelName = &v
	return s
}

func (s *Program) SetClipRange(v string) *Program {
	s.ClipRange = &v
	return s
}

func (s *Program) SetGmtCreate(v string) *Program {
	s.GmtCreate = &v
	return s
}

func (s *Program) SetGmtModified(v string) *Program {
	s.GmtModified = &v
	return s
}

func (s *Program) SetProgramName(v string) *Program {
	s.ProgramName = &v
	return s
}

func (s *Program) SetSourceLocationName(v string) *Program {
	s.SourceLocationName = &v
	return s
}

func (s *Program) SetSourceName(v string) *Program {
	s.SourceName = &v
	return s
}

func (s *Program) SetSourceType(v string) *Program {
	s.SourceType = &v
	return s
}

func (s *Program) SetTransition(v string) *Program {
	s.Transition = &v
	return s
}

func (s *Program) Validate() error {
	if s.AdBreaks != nil {
		for _, item := range s.AdBreaks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ProgramAdBreaks struct {
	ChannelName          *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OffsetMillis         *int64  `json:"OffsetMillis,omitempty" xml:"OffsetMillis,omitempty"`
	ProgramName          *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	SourceLocationName   *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	SourceName           *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SpliceInsertSettings *string `json:"SpliceInsertSettings,omitempty" xml:"SpliceInsertSettings,omitempty"`
	TimeSignalSettings   *string `json:"TimeSignalSettings,omitempty" xml:"TimeSignalSettings,omitempty"`
}

func (s ProgramAdBreaks) String() string {
	return dara.Prettify(s)
}

func (s ProgramAdBreaks) GoString() string {
	return s.String()
}

func (s *ProgramAdBreaks) GetChannelName() *string {
	return s.ChannelName
}

func (s *ProgramAdBreaks) GetMessageType() *string {
	return s.MessageType
}

func (s *ProgramAdBreaks) GetOffsetMillis() *int64 {
	return s.OffsetMillis
}

func (s *ProgramAdBreaks) GetProgramName() *string {
	return s.ProgramName
}

func (s *ProgramAdBreaks) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ProgramAdBreaks) GetSourceName() *string {
	return s.SourceName
}

func (s *ProgramAdBreaks) GetSpliceInsertSettings() *string {
	return s.SpliceInsertSettings
}

func (s *ProgramAdBreaks) GetTimeSignalSettings() *string {
	return s.TimeSignalSettings
}

func (s *ProgramAdBreaks) SetChannelName(v string) *ProgramAdBreaks {
	s.ChannelName = &v
	return s
}

func (s *ProgramAdBreaks) SetMessageType(v string) *ProgramAdBreaks {
	s.MessageType = &v
	return s
}

func (s *ProgramAdBreaks) SetOffsetMillis(v int64) *ProgramAdBreaks {
	s.OffsetMillis = &v
	return s
}

func (s *ProgramAdBreaks) SetProgramName(v string) *ProgramAdBreaks {
	s.ProgramName = &v
	return s
}

func (s *ProgramAdBreaks) SetSourceLocationName(v string) *ProgramAdBreaks {
	s.SourceLocationName = &v
	return s
}

func (s *ProgramAdBreaks) SetSourceName(v string) *ProgramAdBreaks {
	s.SourceName = &v
	return s
}

func (s *ProgramAdBreaks) SetSpliceInsertSettings(v string) *ProgramAdBreaks {
	s.SpliceInsertSettings = &v
	return s
}

func (s *ProgramAdBreaks) SetTimeSignalSettings(v string) *ProgramAdBreaks {
	s.TimeSignalSettings = &v
	return s
}

func (s *ProgramAdBreaks) Validate() error {
	return dara.Validate(s)
}
