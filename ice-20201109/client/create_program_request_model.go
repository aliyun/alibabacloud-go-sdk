// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v string) *CreateProgramRequest
	GetAdBreaks() *string
	SetChannelName(v string) *CreateProgramRequest
	GetChannelName() *string
	SetClipRange(v string) *CreateProgramRequest
	GetClipRange() *string
	SetProgramName(v string) *CreateProgramRequest
	GetProgramName() *string
	SetSourceLocationName(v string) *CreateProgramRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *CreateProgramRequest
	GetSourceName() *string
	SetSourceType(v string) *CreateProgramRequest
	GetSourceType() *string
	SetTransition(v string) *CreateProgramRequest
	GetTransition() *string
}

type CreateProgramRequest struct {
	// The information about ad breaks.
	//
	// example:
	//
	// [{"MessageType":"SPLICE_INSERT","OffsetMillis":1000,"SourceLocationName":"MySourceLocation","SourceName":"MyAdSource","SpliceInsertSettings":{"AvailNumber":0,"AvailExpected":0,"SpliceEventID":1,"UniqueProgramID":0}}]
	AdBreaks *string `json:"AdBreaks,omitempty" xml:"AdBreaks,omitempty"`
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// Extracts a clip from the source.
	//
	// example:
	//
	// {StartOffsetMillis: 213123, EndOffsetMillis: 213134}
	ClipRange *string `json:"ClipRange,omitempty" xml:"ClipRange,omitempty"`
	// The name of the program.
	//
	// This parameter is required.
	//
	// example:
	//
	// program1
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The source location.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySourceLcation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySource
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The source type of the program.
	//
	// This parameter is required.
	//
	// example:
	//
	// vodSource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The program transition method.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Type": "RELATIVE", "RelativePosition": "AFTER_PROGRAM", "RelativeProgram": "program2"}
	Transition *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s CreateProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProgramRequest) GoString() string {
	return s.String()
}

func (s *CreateProgramRequest) GetAdBreaks() *string {
	return s.AdBreaks
}

func (s *CreateProgramRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateProgramRequest) GetClipRange() *string {
	return s.ClipRange
}

func (s *CreateProgramRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *CreateProgramRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *CreateProgramRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *CreateProgramRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateProgramRequest) GetTransition() *string {
	return s.Transition
}

func (s *CreateProgramRequest) SetAdBreaks(v string) *CreateProgramRequest {
	s.AdBreaks = &v
	return s
}

func (s *CreateProgramRequest) SetChannelName(v string) *CreateProgramRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateProgramRequest) SetClipRange(v string) *CreateProgramRequest {
	s.ClipRange = &v
	return s
}

func (s *CreateProgramRequest) SetProgramName(v string) *CreateProgramRequest {
	s.ProgramName = &v
	return s
}

func (s *CreateProgramRequest) SetSourceLocationName(v string) *CreateProgramRequest {
	s.SourceLocationName = &v
	return s
}

func (s *CreateProgramRequest) SetSourceName(v string) *CreateProgramRequest {
	s.SourceName = &v
	return s
}

func (s *CreateProgramRequest) SetSourceType(v string) *CreateProgramRequest {
	s.SourceType = &v
	return s
}

func (s *CreateProgramRequest) SetTransition(v string) *CreateProgramRequest {
	s.Transition = &v
	return s
}

func (s *CreateProgramRequest) Validate() error {
	return dara.Validate(s)
}
