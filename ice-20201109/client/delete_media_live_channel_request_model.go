// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *DeleteMediaLiveChannelRequest
	GetChannelId() *string
}

type DeleteMediaLiveChannelRequest struct {
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DeleteMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DeleteMediaLiveChannelRequest) SetChannelId(v string) *DeleteMediaLiveChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DeleteMediaLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}
