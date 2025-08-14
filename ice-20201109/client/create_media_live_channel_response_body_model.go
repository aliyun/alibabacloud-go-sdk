// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *CreateMediaLiveChannelResponseBody
	GetChannelId() *string
	SetRequestId(v string) *CreateMediaLiveChannelResponseBody
	GetRequestId() *string
}

type CreateMediaLiveChannelResponseBody struct {
	// The ID of the channel.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaLiveChannelResponseBody) SetChannelId(v string) *CreateMediaLiveChannelResponseBody {
	s.ChannelId = &v
	return s
}

func (s *CreateMediaLiveChannelResponseBody) SetRequestId(v string) *CreateMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
