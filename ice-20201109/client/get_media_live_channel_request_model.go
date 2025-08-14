// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *GetMediaLiveChannelRequest
	GetChannelId() *string
}

type GetMediaLiveChannelRequest struct {
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s GetMediaLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetMediaLiveChannelRequest) SetChannelId(v string) *GetMediaLiveChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *GetMediaLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}
