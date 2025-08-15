// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhitelistShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelistsShrink(v string) *GetInstanceIpWhitelistShrinkRequest
	GetIpWhitelistsShrink() *string
}

type GetInstanceIpWhitelistShrinkRequest struct {
	// The  filter IP address whitelists.
	IpWhitelistsShrink *string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty"`
}

func (s GetInstanceIpWhitelistShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhitelistShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhitelistShrinkRequest) GetIpWhitelistsShrink() *string {
	return s.IpWhitelistsShrink
}

func (s *GetInstanceIpWhitelistShrinkRequest) SetIpWhitelistsShrink(v string) *GetInstanceIpWhitelistShrinkRequest {
	s.IpWhitelistsShrink = &v
	return s
}

func (s *GetInstanceIpWhitelistShrinkRequest) Validate() error {
	return dara.Validate(s)
}
