// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDataNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v string) *DeleteGlobalDataNetworkRequest
	GetNetworkId() *string
}

type DeleteGlobalDataNetworkRequest struct {
	// GDN ID
	//
	// example:
	//
	// gdn-xxx
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DeleteGlobalDataNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDataNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDataNetworkRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DeleteGlobalDataNetworkRequest) SetNetworkId(v string) *DeleteGlobalDataNetworkRequest {
	s.NetworkId = &v
	return s
}

func (s *DeleteGlobalDataNetworkRequest) Validate() error {
	return dara.Validate(s)
}
