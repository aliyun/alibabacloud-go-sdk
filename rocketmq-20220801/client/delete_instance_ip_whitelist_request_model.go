// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelist(v string) *DeleteInstanceIpWhitelistRequest
	GetIpWhitelist() *string
	SetIpWhitelists(v []*string) *DeleteInstanceIpWhitelistRequest
	GetIpWhitelists() []*string
}

type DeleteInstanceIpWhitelistRequest struct {
	// Deprecated
	//
	// The IP address whitelist.
	//
	// example:
	//
	// 0.0.0.0/0
	IpWhitelist *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
	// The IP address whitelist.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
}

func (s DeleteInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *DeleteInstanceIpWhitelistRequest) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *DeleteInstanceIpWhitelistRequest) SetIpWhitelist(v string) *DeleteInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

func (s *DeleteInstanceIpWhitelistRequest) SetIpWhitelists(v []*string) *DeleteInstanceIpWhitelistRequest {
	s.IpWhitelists = v
	return s
}

func (s *DeleteInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
