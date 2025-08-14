// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *ChannelAssemblyChannel) *GetChannelResponseBody
	GetChannel() *ChannelAssemblyChannel
	SetRequestId(v string) *GetChannelResponseBody
	GetRequestId() *string
}

type GetChannelResponseBody struct {
	// The channel information.
	Channel *ChannelAssemblyChannel `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetChannelResponseBody) GetChannel() *ChannelAssemblyChannel {
	return s.Channel
}

func (s *GetChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChannelResponseBody) SetChannel(v *ChannelAssemblyChannel) *GetChannelResponseBody {
	s.Channel = v
	return s
}

func (s *GetChannelResponseBody) SetRequestId(v string) *GetChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
