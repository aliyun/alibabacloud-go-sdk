// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDataNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *CreateGlobalDataNetworkResponseBody
	GetChannelId() *string
	SetNetworkId(v string) *CreateGlobalDataNetworkResponseBody
	GetNetworkId() *string
	SetRequestId(v string) *CreateGlobalDataNetworkResponseBody
	GetRequestId() *string
}

type CreateGlobalDataNetworkResponseBody struct {
	// example:
	//
	// gdc-xxx
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// GDN ID
	//
	// example:
	//
	// gdn-xxx
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalDataNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDataNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalDataNetworkResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateGlobalDataNetworkResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateGlobalDataNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalDataNetworkResponseBody) SetChannelId(v string) *CreateGlobalDataNetworkResponseBody {
	s.ChannelId = &v
	return s
}

func (s *CreateGlobalDataNetworkResponseBody) SetNetworkId(v string) *CreateGlobalDataNetworkResponseBody {
	s.NetworkId = &v
	return s
}

func (s *CreateGlobalDataNetworkResponseBody) SetRequestId(v string) *CreateGlobalDataNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalDataNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
