// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v string) *DeleteNetworkRequest
	GetNetworkId() *string
}

type DeleteNetworkRequest struct {
	// The ID of the network.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-5***
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DeleteNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DeleteNetworkRequest) SetNetworkId(v string) *DeleteNetworkRequest {
	s.NetworkId = &v
	return s
}

func (s *DeleteNetworkRequest) Validate() error {
	return dara.Validate(s)
}
