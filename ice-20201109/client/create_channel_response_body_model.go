// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *ChannelAssemblyChannel) *CreateChannelResponseBody
	GetChannel() *ChannelAssemblyChannel
	SetRequestId(v string) *CreateChannelResponseBody
	GetRequestId() *string
}

type CreateChannelResponseBody struct {
	// The channel information.
	Channel *ChannelAssemblyChannel `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChannelResponseBody) GetChannel() *ChannelAssemblyChannel {
	return s.Channel
}

func (s *CreateChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChannelResponseBody) SetChannel(v *ChannelAssemblyChannel) *CreateChannelResponseBody {
	s.Channel = v
	return s
}

func (s *CreateChannelResponseBody) SetRequestId(v string) *CreateChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
