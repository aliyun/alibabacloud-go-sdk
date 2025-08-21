// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceIdsShrink(v string) *DeleteNetworkInterfacesShrinkRequest
	GetNetworkInterfaceIdsShrink() *string
}

type DeleteNetworkInterfacesShrinkRequest struct {
	// This parameter is required.
	NetworkInterfaceIdsShrink *string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty"`
}

func (s DeleteNetworkInterfacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacesShrinkRequest) GetNetworkInterfaceIdsShrink() *string {
	return s.NetworkInterfaceIdsShrink
}

func (s *DeleteNetworkInterfacesShrinkRequest) SetNetworkInterfaceIdsShrink(v string) *DeleteNetworkInterfacesShrinkRequest {
	s.NetworkInterfaceIdsShrink = &v
	return s
}

func (s *DeleteNetworkInterfacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
