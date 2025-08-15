// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelists(v []*string) *CreateInstanceIpWhitelistRequest
	GetIpWhitelists() []*string
}

type CreateInstanceIpWhitelistRequest struct {
	// The IP address whitelists.
	//
	// This parameter is required.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
}

func (s CreateInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistRequest) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *CreateInstanceIpWhitelistRequest) SetIpWhitelists(v []*string) *CreateInstanceIpWhitelistRequest {
	s.IpWhitelists = v
	return s
}

func (s *CreateInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
