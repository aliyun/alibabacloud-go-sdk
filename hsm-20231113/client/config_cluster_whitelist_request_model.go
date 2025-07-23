// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConfigClusterWhitelistRequest
	GetClusterId() *string
	SetWhitelist(v string) *ConfigClusterWhitelistRequest
	GetWhitelist() *string
}

type ConfigClusterWhitelistRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-729dm40FG****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IP address whitelist of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ConfigClusterWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConfigClusterWhitelistRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *ConfigClusterWhitelistRequest) SetClusterId(v string) *ConfigClusterWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterWhitelistRequest) SetWhitelist(v string) *ConfigClusterWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *ConfigClusterWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
