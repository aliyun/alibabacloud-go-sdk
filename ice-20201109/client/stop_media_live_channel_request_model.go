// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *StopMediaLiveChannelRequest
	GetChannelId() *string
}

type StopMediaLiveChannelRequest struct {
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s StopMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *StopMediaLiveChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopMediaLiveChannelRequest) SetChannelId(v string) *StopMediaLiveChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *StopMediaLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}
