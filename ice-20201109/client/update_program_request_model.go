// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdBreaks(v string) *UpdateProgramRequest
	GetAdBreaks() *string
	SetChannelName(v string) *UpdateProgramRequest
	GetChannelName() *string
	SetClipRange(v string) *UpdateProgramRequest
	GetClipRange() *string
	SetProgramName(v string) *UpdateProgramRequest
	GetProgramName() *string
	SetSourceLocationName(v string) *UpdateProgramRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *UpdateProgramRequest
	GetSourceName() *string
	SetSourceType(v string) *UpdateProgramRequest
	GetSourceType() *string
	SetTransition(v string) *UpdateProgramRequest
	GetTransition() *string
}

type UpdateProgramRequest struct {
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
	// The name of the source location.
	//
	// example:
	//
	// MySourceLcation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
	// The name of the source.
	//
	// example:
	//
	// MySource
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The source type of the program. Valid values: vodSource and liveSource.
	//
	// example:
	//
	// vodSource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The program transition method.
	//
	// example:
	//
	// {"Type": "RELATIVE", "RelativePosition": "AFTER_PROGRAM", "RelativeProgram": "program2"}
	Transition *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s UpdateProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProgramRequest) GoString() string {
	return s.String()
}

func (s *UpdateProgramRequest) GetAdBreaks() *string {
	return s.AdBreaks
}

func (s *UpdateProgramRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateProgramRequest) GetClipRange() *string {
	return s.ClipRange
}

func (s *UpdateProgramRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *UpdateProgramRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *UpdateProgramRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *UpdateProgramRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateProgramRequest) GetTransition() *string {
	return s.Transition
}

func (s *UpdateProgramRequest) SetAdBreaks(v string) *UpdateProgramRequest {
	s.AdBreaks = &v
	return s
}

func (s *UpdateProgramRequest) SetChannelName(v string) *UpdateProgramRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateProgramRequest) SetClipRange(v string) *UpdateProgramRequest {
	s.ClipRange = &v
	return s
}

func (s *UpdateProgramRequest) SetProgramName(v string) *UpdateProgramRequest {
	s.ProgramName = &v
	return s
}

func (s *UpdateProgramRequest) SetSourceLocationName(v string) *UpdateProgramRequest {
	s.SourceLocationName = &v
	return s
}

func (s *UpdateProgramRequest) SetSourceName(v string) *UpdateProgramRequest {
	s.SourceName = &v
	return s
}

func (s *UpdateProgramRequest) SetSourceType(v string) *UpdateProgramRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateProgramRequest) SetTransition(v string) *UpdateProgramRequest {
	s.Transition = &v
	return s
}

func (s *UpdateProgramRequest) Validate() error {
	return dara.Validate(s)
}
