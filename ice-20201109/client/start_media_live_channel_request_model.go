// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *StartMediaLiveChannelRequest
	GetChannelId() *string
}

type StartMediaLiveChannelRequest struct {
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s StartMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *StartMediaLiveChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartMediaLiveChannelRequest) SetChannelId(v string) *StartMediaLiveChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *StartMediaLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}
