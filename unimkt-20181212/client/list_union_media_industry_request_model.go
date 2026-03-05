// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnionMediaIndustryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *ListUnionMediaIndustryRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *ListUnionMediaIndustryRequest
	GetProxyUserId() *int64
}

type ListUnionMediaIndustryRequest struct {
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ProxyUserId *int64  `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s ListUnionMediaIndustryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnionMediaIndustryRequest) GoString() string {
	return s.String()
}

func (s *ListUnionMediaIndustryRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListUnionMediaIndustryRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *ListUnionMediaIndustryRequest) SetChannelId(v string) *ListUnionMediaIndustryRequest {
	s.ChannelId = &v
	return s
}

func (s *ListUnionMediaIndustryRequest) SetProxyUserId(v int64) *ListUnionMediaIndustryRequest {
	s.ProxyUserId = &v
	return s
}

func (s *ListUnionMediaIndustryRequest) Validate() error {
	return dara.Validate(s)
}
