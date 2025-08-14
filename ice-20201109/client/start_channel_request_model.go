// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *StartChannelRequest
	GetChannelName() *string
}

type StartChannelRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
}

func (s StartChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s StartChannelRequest) GoString() string {
	return s.String()
}

func (s *StartChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *StartChannelRequest) SetChannelName(v string) *StartChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *StartChannelRequest) Validate() error {
	return dara.Validate(s)
}
