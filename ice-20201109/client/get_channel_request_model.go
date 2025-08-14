// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *GetChannelRequest
	GetChannelName() *string
}

type GetChannelRequest struct {
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
}

func (s GetChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChannelRequest) GoString() string {
	return s.String()
}

func (s *GetChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetChannelRequest) SetChannelName(v string) *GetChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *GetChannelRequest) Validate() error {
	return dara.Validate(s)
}
