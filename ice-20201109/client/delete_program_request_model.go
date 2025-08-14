// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteProgramRequest
	GetChannelName() *string
	SetProgramName(v string) *DeleteProgramRequest
	GetProgramName() *string
}

type DeleteProgramRequest struct {
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
	// program_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
}

func (s DeleteProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProgramRequest) GoString() string {
	return s.String()
}

func (s *DeleteProgramRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteProgramRequest) GetProgramName() *string {
	return s.ProgramName
}

func (s *DeleteProgramRequest) SetChannelName(v string) *DeleteProgramRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteProgramRequest) SetProgramName(v string) *DeleteProgramRequest {
	s.ProgramName = &v
	return s
}

func (s *DeleteProgramRequest) Validate() error {
	return dara.Validate(s)
}
