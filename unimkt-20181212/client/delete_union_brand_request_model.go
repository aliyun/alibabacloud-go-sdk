// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnionBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandMainId(v int64) *DeleteUnionBrandRequest
	GetBrandMainId() *int64
	SetChannelId(v string) *DeleteUnionBrandRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *DeleteUnionBrandRequest
	GetProxyUserId() *int64
}

type DeleteUnionBrandRequest struct {
	// example:
	//
	// 167332546421354
	BrandMainId *int64 `json:"BrandMainId,omitempty" xml:"BrandMainId,omitempty"`
	// example:
	//
	// QD-WXXJ-403576
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 167332546421354
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s DeleteUnionBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnionBrandRequest) GoString() string {
	return s.String()
}

func (s *DeleteUnionBrandRequest) GetBrandMainId() *int64 {
	return s.BrandMainId
}

func (s *DeleteUnionBrandRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DeleteUnionBrandRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *DeleteUnionBrandRequest) SetBrandMainId(v int64) *DeleteUnionBrandRequest {
	s.BrandMainId = &v
	return s
}

func (s *DeleteUnionBrandRequest) SetChannelId(v string) *DeleteUnionBrandRequest {
	s.ChannelId = &v
	return s
}

func (s *DeleteUnionBrandRequest) SetProxyUserId(v int64) *DeleteUnionBrandRequest {
	s.ProxyUserId = &v
	return s
}

func (s *DeleteUnionBrandRequest) Validate() error {
	return dara.Validate(s)
}
