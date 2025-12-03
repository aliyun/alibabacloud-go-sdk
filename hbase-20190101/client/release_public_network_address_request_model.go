// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReleasePublicNetworkAddressRequest
	GetClusterId() *string
}

type ReleasePublicNetworkAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReleasePublicNetworkAddressRequest) SetClusterId(v string) *ReleasePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
