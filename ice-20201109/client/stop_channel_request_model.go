// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *StopChannelRequest
	GetChannelName() *string
}

type StopChannelRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
}

func (s StopChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s StopChannelRequest) GoString() string {
	return s.String()
}

func (s *StopChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *StopChannelRequest) SetChannelName(v string) *StopChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *StopChannelRequest) Validate() error {
	return dara.Validate(s)
}
