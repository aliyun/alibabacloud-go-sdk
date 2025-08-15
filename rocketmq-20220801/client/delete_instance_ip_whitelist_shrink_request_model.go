// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceIpWhitelistShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelist(v string) *DeleteInstanceIpWhitelistShrinkRequest
	GetIpWhitelist() *string
	SetIpWhitelistsShrink(v string) *DeleteInstanceIpWhitelistShrinkRequest
	GetIpWhitelistsShrink() *string
}

type DeleteInstanceIpWhitelistShrinkRequest struct {
	// Deprecated
	//
	// The IP address whitelist.
	//
	// example:
	//
	// 0.0.0.0/0
	IpWhitelist *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
	// The IP address whitelist.
	IpWhitelistsShrink *string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty"`
}

func (s DeleteInstanceIpWhitelistShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceIpWhitelistShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistShrinkRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *DeleteInstanceIpWhitelistShrinkRequest) GetIpWhitelistsShrink() *string {
	return s.IpWhitelistsShrink
}

func (s *DeleteInstanceIpWhitelistShrinkRequest) SetIpWhitelist(v string) *DeleteInstanceIpWhitelistShrinkRequest {
	s.IpWhitelist = &v
	return s
}

func (s *DeleteInstanceIpWhitelistShrinkRequest) SetIpWhitelistsShrink(v string) *DeleteInstanceIpWhitelistShrinkRequest {
	s.IpWhitelistsShrink = &v
	return s
}

func (s *DeleteInstanceIpWhitelistShrinkRequest) Validate() error {
	return dara.Validate(s)
}
