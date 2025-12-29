// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushVoiceBoxCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommands(v []*PushVoiceBoxCommandsRequestCommands) *PushVoiceBoxCommandsRequest
	GetCommands() []*PushVoiceBoxCommandsRequestCommands
	SetHotelId(v string) *PushVoiceBoxCommandsRequest
	GetHotelId() *string
	SetRoomNo(v string) *PushVoiceBoxCommandsRequest
	GetRoomNo() *string
}

type PushVoiceBoxCommandsRequest struct {
	// This parameter is required.
	Commands []*PushVoiceBoxCommandsRequestCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// This parameter is required.
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s PushVoiceBoxCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsRequest) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsRequest) GetCommands() []*PushVoiceBoxCommandsRequestCommands {
	return s.Commands
}

func (s *PushVoiceBoxCommandsRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *PushVoiceBoxCommandsRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PushVoiceBoxCommandsRequest) SetCommands(v []*PushVoiceBoxCommandsRequestCommands) *PushVoiceBoxCommandsRequest {
	s.Commands = v
	return s
}

func (s *PushVoiceBoxCommandsRequest) SetHotelId(v string) *PushVoiceBoxCommandsRequest {
	s.HotelId = &v
	return s
}

func (s *PushVoiceBoxCommandsRequest) SetRoomNo(v string) *PushVoiceBoxCommandsRequest {
	s.RoomNo = &v
	return s
}

func (s *PushVoiceBoxCommandsRequest) Validate() error {
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PushVoiceBoxCommandsRequestCommands struct {
	// This parameter is required.
	CommandDomain *string `json:"CommandDomain,omitempty" xml:"CommandDomain,omitempty"`
	// This parameter is required.
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	Payload     *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s PushVoiceBoxCommandsRequestCommands) String() string {
	return dara.Prettify(s)
}

func (s PushVoiceBoxCommandsRequestCommands) GoString() string {
	return s.String()
}

func (s *PushVoiceBoxCommandsRequestCommands) GetCommandDomain() *string {
	return s.CommandDomain
}

func (s *PushVoiceBoxCommandsRequestCommands) GetCommandName() *string {
	return s.CommandName
}

func (s *PushVoiceBoxCommandsRequestCommands) GetPayload() *string {
	return s.Payload
}

func (s *PushVoiceBoxCommandsRequestCommands) SetCommandDomain(v string) *PushVoiceBoxCommandsRequestCommands {
	s.CommandDomain = &v
	return s
}

func (s *PushVoiceBoxCommandsRequestCommands) SetCommandName(v string) *PushVoiceBoxCommandsRequestCommands {
	s.CommandName = &v
	return s
}

func (s *PushVoiceBoxCommandsRequestCommands) SetPayload(v string) *PushVoiceBoxCommandsRequestCommands {
	s.Payload = &v
	return s
}

func (s *PushVoiceBoxCommandsRequestCommands) Validate() error {
	return dara.Validate(s)
}
