// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelists(v []*string) *GetInstanceIpWhitelistRequest
	GetIpWhitelists() []*string
}

type GetInstanceIpWhitelistRequest struct {
	// The  filter IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
}

func (s GetInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhitelistRequest) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *GetInstanceIpWhitelistRequest) SetIpWhitelists(v []*string) *GetInstanceIpWhitelistRequest {
	s.IpWhitelists = v
	return s
}

func (s *GetInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
