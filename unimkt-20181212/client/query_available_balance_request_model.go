// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryAvailableBalanceRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *QueryAvailableBalanceRequest
	GetProxyUserId() *int64
}

type QueryAvailableBalanceRequest struct {
	// This parameter is required.
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ProxyUserId *int64  `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s QueryAvailableBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableBalanceRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableBalanceRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryAvailableBalanceRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *QueryAvailableBalanceRequest) SetChannelId(v string) *QueryAvailableBalanceRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryAvailableBalanceRequest) SetProxyUserId(v int64) *QueryAvailableBalanceRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryAvailableBalanceRequest) Validate() error {
	return dara.Validate(s)
}
