// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *ChannelAssemblyChannel) *UpdateChannelResponseBody
	GetChannel() *ChannelAssemblyChannel
	SetRequestId(v string) *UpdateChannelResponseBody
	GetRequestId() *string
}

type UpdateChannelResponseBody struct {
	// The channel information.
	Channel *ChannelAssemblyChannel `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChannelResponseBody) GetChannel() *ChannelAssemblyChannel {
	return s.Channel
}

func (s *UpdateChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChannelResponseBody) SetChannel(v *ChannelAssemblyChannel) *UpdateChannelResponseBody {
	s.Channel = v
	return s
}

func (s *UpdateChannelResponseBody) SetRequestId(v string) *UpdateChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
