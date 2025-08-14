// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteChannelRequest
	GetChannelName() *string
}

type DeleteChannelRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
}

func (s DeleteChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteChannelRequest) SetChannelName(v string) *DeleteChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteChannelRequest) Validate() error {
	return dara.Validate(s)
}
