// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceIds(v []*string) *DeleteNetworkInterfacesRequest
	GetNetworkInterfaceIds() []*string
}

type DeleteNetworkInterfacesRequest struct {
	// The IDs of the elastic network interfaces (ENIs).
	//
	// This parameter is required.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
}

func (s DeleteNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *DeleteNetworkInterfacesRequest) SetNetworkInterfaceIds(v []*string) *DeleteNetworkInterfacesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DeleteNetworkInterfacesRequest) Validate() error {
	return dara.Validate(s)
}
