// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *GetProgramRequest
	GetChannelName() *string
	SetProgramName(v string) *GetProgramRequest
	GetProgramName() *string
}

type GetProgramRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The name of the program.
	//
	// This parameter is required.
	//
	// example:
	//
	// program1
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
}

func (s GetProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProgramRequest) GoString() string {
	return s.String()
}

func (s *GetProgramRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetProgramRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *GetProgramRequest) SetChannelName(v string) *GetProgramRequest {
	s.ChannelName = &v
	return s
}

func (s *GetProgramRequest) SetProgramName(v string) *GetProgramRequest {
	s.ProgramName = &v
	return s
}

func (s *GetProgramRequest) Validate() error {
	return dara.Validate(s)
}
